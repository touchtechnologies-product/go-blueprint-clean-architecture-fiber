package mongodb

import (
	"github.com/touchtechnologies-product/go-blueprint-clean-architecture/service/util"
	"fmt"
)

func (suite *MongoDBTestSuite) TestPush() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &util.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         "c",
	}
	err = suite.repo.Push(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 3)
}

func (suite *MongoDBTestSuite) TestPop() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &util.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         []string{},
	}
	err = suite.repo.Pop(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 1)
}

func (suite *MongoDBTestSuite) TestIsFirst() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &util.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         "a",
	}
	err = suite.repo.Pop(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 1)
}

func (suite *MongoDBTestSuite) TestCountArray() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &util.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         "a",
	}
	total, err := suite.repo.CountArray(suite.ctx, param)
	suite.NoError(err)
	suite.Equal(total, 2)
}

func (suite *MongoDBTestSuite) TestClearArray() {
	test := suite.makeTestStruct("test", "a", "b")
	_, err := suite.repo.Create(suite.ctx, test)
	suite.NoError(err)

	param := &util.SetOpParam{
		Filters:      []string{"title:eq:test"},
		SetFieldName: "list",
		Item:         []string{},
	}
	err = suite.repo.ClearArray(suite.ctx, param)
	suite.NoError(err)

	filters := []string{fmt.Sprintf("title:eq:%s", test.Title)}
	err = suite.repo.Read(suite.ctx, filters, test)
	suite.NoError(err)
	suite.Len(test.List, 0)
}