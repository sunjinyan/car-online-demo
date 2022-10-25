package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"os"
	"os/signal"
	"time"
)

var  singleCh  chan  os.Signal
var  mongoCol   *mongo.Collection
var  ctx   context.Context

func init() {
	singleCh = make(chan os.Signal)
	ctx = context.Background()
	mongoConn()
}

func main() {

	go func() {
		findRows(ctx,mongoCol)
	}()


	signal.Notify(singleCh,os.Interrupt,os.Kill)
	<-singleCh
}

func mongoConn()  {
		var timer   time.Duration = 5*time.Second
		var MaxPoolSize   uint64 = 100
		var MinPoolSize   uint64 = 50
		pref, err := readpref.New(readpref.PrimaryMode)
		if err != nil {
			fmt.Println(err.Error())
			singleCh <- os.Interrupt
			os.Exit(2)
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
			fmt.Println(err.Error())
			singleCh <- os.Interrupt
			os.Exit(2)
		}

		mongoCol = con.Database("coolcar").Collection("account")


		//one := col.FindOne(ctx, bson.M{
		//	"open_id": "123789",
		//})
}

func findRows(c context.Context, col *mongo.Collection) {
	res := col.FindOne(c, bson.M{
		"open_id": "123789",
	})

	if err := res.Err(); err  != nil {
		fmt.Println(err.Error(),1111)
		singleCh <- os.Interrupt
		return
	}

	var row struct{
		ID primitive.ObjectID `bson:"_id"`
		OpenId string  `bson:"open_id"`
	}
	err := res.Decode(&row)
	if err != nil {
		fmt.Println(err.Error(),2222)
		singleCh <- os.Interrupt
		return
	}
	fmt.Printf("%+v",row)
}


func insertRows(c context.Context,col *mongo.Collection) {
	res, err := col.InsertMany(c, []interface{}{
		bson.M{
			"open_id": "123abv",
		},
		bson.M{
			"open_id": "123abvqqq",
		},
		bson.M{
			"open_id": "123abvqqq",
		},
	})

	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v",res)
}