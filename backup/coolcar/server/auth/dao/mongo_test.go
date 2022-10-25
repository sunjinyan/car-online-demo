package dao

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"testing"
	"time"
)
var  db   *mongo.Database
var  ctx   context.Context
func TestMongo_ResolveAccountID(t *testing.T) {
	var timer   time.Duration = 5*time.Second
	var MaxPoolSize   uint64 = 100
	var MinPoolSize   uint64 = 50
	pref, err := readpref.New(readpref.PrimaryMode)
	if err != nil {
		t.Fatalf("cannot new readpref %v",err)
		//panic(err)
	}
	//options.Client().ApplyURI("mongodb://47.93.20.75:27017/coolcar?readPreference=primary&ssl=false")
	con, err := mongo.Connect(ctx, &options.ClientOptions{
		ConnectTimeout:           &timer,
		Hosts:                    []string{
			"47.93.20.75:27017",
		},
		MaxPoolSize:              &MaxPoolSize,
		MinPoolSize:              &MinPoolSize,
		ReadPreference:  pref,
	})
	if err != nil {

		t.Fatalf("cannot Connect %v",err)
	}

	db = con.Database("coolcar")

	mon := NewMongo(db)

	id, err := mon.ResolveAccountID(ctx, "123")
	if err != nil {

		t.Fatalf("cannot ResolveAccountID %v",err)
	}

	t.Logf("get account id : %q\n",id)//%q将输出加上双引号
}
