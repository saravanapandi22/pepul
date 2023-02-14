package services

import (
	"context"
	"errors"
	"example/pepel/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UserImplementation struct {
	mongoCollection *mongo.Collection
	ctx context.Context
}

func NewUserService(mongoCollection *mongo.Collection, ctx context.Context) UserService {
	return &UserImplementation{
		mongoCollection: mongoCollection,
		ctx: ctx,
	}
}

func (ui *UserImplementation) CreateUser(UserModel *models.User) error {
	_, err := ui.mongoCollection.InsertOne(ui.ctx, UserModel)
	return err
}

func (ui *UserImplementation) GetAll() ([]*models.User, error) {
	var UsersData []*models.User
	gctx := gin.Context{
		Request:  nil,
		Writer:   nil,
		Params:   nil,
		Keys:     nil,
		Errors:   nil,
		Accepted: nil,
	}
	var Pagination = AddPagination(&gctx)
	opts := options.Find().SetLimit(int64(Pagination.Limit))
	cursor, err := ui.mongoCollection.Find(ui.ctx, bson.D{{}}, opts)
	if err != nil {
		return nil, err
	}
	for cursor.Next(ui.ctx) {
		var User models.User
		err := cursor.Decode(&User)
		if err != nil {
			return nil, err
		}

		UsersData = append(UsersData, &User)

		if err := cursor.Err(); err != nil {
			return nil, err
		}

		if len(UsersData) == 0 {
			return nil, errors.New("no data found")
		}

		return UsersData, nil
	}
	return nil, nil
}

func (ui *UserImplementation) UpdateUser(UserModel *models.User) error {
	filterCond := bson.D{bson.E{Key: "full_name", Value: UserModel.FullName}}
	updateCond := bson.D{bson.E{Key: "$set", Value: bson.D{bson.E{Key: "email", Value: UserModel.Email}, bson.E{Key: "phone_number", Value: UserModel.PhoneNumber}, bson.E{Key: "picture_upload", Value: UserModel.PictureUpload}, bson.E{Key: "update_at", Value: time.Now()}}}}
	updateRes, _ := ui.mongoCollection.UpdateOne(ui.ctx, filterCond, updateCond)
	if updateRes.MatchedCount != 1 {
		return errors.New("no data affected")
	}
	return nil
}

func (ui *UserImplementation) DeleteUser(Name *string) error {
	filterCond := bson.D{bson.E{Key: "full_name", Value: Name}}
	delRes, _ := ui.mongoCollection.DeleteOne(ui.ctx, filterCond)
	if delRes.DeletedCount != 1 {
		return errors.New("no data affected")
	}
	return nil
}


func AddPagination(gctx *gin.Context) models.Pagination {
	limit, _ := strconv.Atoi(gctx.DefaultQuery("limit", "10"))
	page, _ := strconv.Atoi(gctx.DefaultQuery("page", "0"))
	return models.Pagination{
		Limit: limit,
		Page: page,
	}
}
