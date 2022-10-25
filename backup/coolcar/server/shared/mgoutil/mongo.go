package mgoutil

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)


//都可能是公用的
const (
	IDFieldName = "_id"
	UpdatedAtFieldName = "updatedat"
)

type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

type UpdateAtField struct {
	UpdateAt int64  `bson:"updatedat"`
}

var UpdateAt = func() int64{
	return  time.Now().UnixNano()
}

var NewObjId = primitive.NewObjectID


//Set return a $set update document
func Set(v interface{}) bson.M {
	//primitive.ObjectIDFromHex()
	return bson.M{
		"$set": v,
	}
}


func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert":v,
	}
}