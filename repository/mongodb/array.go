package mongodb

import (
	"context"
	"fmt"

	"blueprint/service/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (repo *Repository) Push(ctx context.Context, param *util.SetOpParam) (err error) {
	filters := repo.makeFilters([]string{fmt.Sprintf("id:eq:%s", param.ID)})
	update := bson.M{
		"$addToSet": bson.M{
			param.SetFieldName: param.Item,
		},
	}
	_, err = repo.Coll.UpdateOne(ctx, filters, update)
	return err
}

func (repo *Repository) Pop(ctx context.Context, param *util.SetOpParam) (err error) {
	filters := repo.makeFilters([]string{fmt.Sprintf("id:eq:%s", param.ID)})
	update := bson.M{
		"$pop": bson.M{
			param.SetFieldName: -1,
		},
	}
	_, err = repo.Coll.UpdateOne(ctx, filters, update)
	return err
}

func (repo *Repository) IsFirst(ctx context.Context, param *util.SetOpParam) (is bool, err error) {
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"id": param.ID},
		},
		bson.M{
			"$project": bson.M{
				"id":    1,
				"first": bson.M{"$arrayElemAt": bson.A{fmt.Sprintf("$%s", param.SetFieldName), 0}},
			},
		},
	}

	cursor, err := repo.Coll.Aggregate(ctx, pipeline)
	if err != nil {
		return false, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	var items []struct {
		ID    primitive.ObjectID `bson:"_id,omitempty"`
		First string             `bson:"first"`
	}

	err = cursor.All(ctx, &items)
	if err != nil {
		return false, err
	}

	if len(items) < 1 {
		return false, err
	}

	return items[0].First == param.Item.(string), nil
}

func (repo *Repository) CountArray(ctx context.Context, param *util.SetOpParam) (total int, err error) {
	pipeline := bson.A{
		bson.M{
			"$match": bson.M{"id": param.ID},
		},
		bson.M{
			"$project": bson.M{
				"id":    1,
				"total": bson.M{"$size": fmt.Sprintf("$%s", param.SetFieldName)},
			},
		},
	}

	cursor, err := repo.Coll.Aggregate(ctx, pipeline)
	if err != nil {
		return 0, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	var items []struct {
		ID    string `bson:"id"`
		Total int    `bson:"total"`
	}

	err = cursor.All(ctx, &items)
	if err != nil {
		return 0, err
	}

	if len(items) < 1 {
		return 0, err
	}

	return items[0].Total, nil
}

func (repo *Repository) ClearArray(ctx context.Context, param *util.SetOpParam) (err error) {
	filters := repo.makeFilters([]string{fmt.Sprintf("id:eq:%s", param.ID)})
	_, err = repo.Coll.UpdateOne(ctx, filters, bson.M{"$set": bson.M{param.SetFieldName: param.Item}})
	return err
}
