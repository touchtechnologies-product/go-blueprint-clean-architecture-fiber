package mongodb

import (
	"context"

	"blueprint/service/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (repo *Repository) List(ctx context.Context, opt *util.PageOption, itemType interface{}) (total int, items []interface{}, err error) {
	var filters bson.M
	var optFilter []string
	var opts *options.FindOptions
	if opt != nil {
		opts = repo.makePagingOpts(opt.Page, opt.PerPage)
		if opt.Filters != nil && len(opt.Filters) > 0 {
			optFilter = opt.Filters
			filters = repo.makeFilters(nil)
		}
	}

	total, err = repo.Count(ctx, optFilter)
	if err != nil {
		return 0, nil, err
	}

	cursor, err := repo.Coll.Find(ctx, filters, opts)
	if err != nil {
		return 0, nil, err
	}
	defer func() { _ = cursor.Close(ctx) }()

	for cursor.Next(ctx) {
		item, err := repo.clone(itemType)
		if err != nil {
			return 0, nil, err
		}
		err = cursor.Decode(item)
		if err != nil {
			return 0, nil, err
		}
		items = append(items, item)
	}

	return total, items, nil
}

func (repo *Repository) Create(ctx context.Context, ent interface{}) (ID string, err error) {
	res, err := repo.Coll.InsertOne(ctx, ent)
	if err != nil {
		return "", err
	}
	return res.InsertedID.(primitive.ObjectID).Hex(), nil
}

func (repo *Repository) Read(ctx context.Context, filters []string, out interface{}) (err error) {
	conditions := repo.makeFilters(filters)
	return repo.Coll.FindOne(ctx, conditions).Decode(out)
}

func (repo *Repository) Update(ctx context.Context, filters []string, ret interface{}) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.UpdateOne(ctx, conditions, bson.M{"$set": ret})
	return err
}

func (repo *Repository) Delete(ctx context.Context, filters []string) (err error) {
	conditions := repo.makeFilters(filters)
	_, err = repo.Coll.DeleteOne(ctx, conditions)
	return err
}

func (repo *Repository) Count(ctx context.Context, filters []string) (total int, err error) {
	cnt, err := repo.Coll.CountDocuments(ctx, repo.makeFilters(filters))
	if err != nil {
		return 0, err
	}
	return int(cnt), nil
}
