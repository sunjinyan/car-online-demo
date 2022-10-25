package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"coolcar/shared/mgoutil"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	accountIdField = "account_id"
)

type Mongo struct {
	col *mongo.Collection

}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{col: db.Collection("profile")}
}

type ProfileRecord struct {
	AccountID string `bson:"accountid"`
	Profile *rentalpb.Profile `bson:"profile"`
}

func (s *Mongo)GetProfile(ctx context.Context,aid id.AccountID) (*ProfileRecord,error)  {
	res := s.col.FindOne(ctx, bson.M{
		accountIdField: aid.String(),
	})
	err := res.Err()
	if err != nil {
		return nil, err
	}

	var p ProfileRecord
	err = res.Decode(&p)
	if err := res.Err(); err != nil {
		return nil, err
	}

	return &p, nil
}

func (s *Mongo)UpdateProfile(ctx context.Context,aid id.AccountID,up *rentalpb.Profile) (*ProfileRecord,error)  {
	var upsert bool = true
	res, err := s.col.UpdateOne(ctx, bson.M{
		accountIdField: aid.String(),
	}, mgoutil.Set(bson.M{
		accountIdField: aid.String(),
		"profile": up,
	}), &options.UpdateOptions{
		Upsert: &upsert,
	})
	if err != nil {
		return nil, err
	}

	if res.MatchedCount == 0 {
		return nil, fmt.Errorf("can not update profile,mathced count is %n returned",0)
	}

	return &ProfileRecord{
		AccountID: aid.String(),
		Profile:   up,
	},nil
}