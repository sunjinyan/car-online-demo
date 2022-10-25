package str_obj

import (
	"coolcar/shared/id"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FromID(id fmt.Stringer) (primitive.ObjectID,error) {
	return primitive.ObjectIDFromHex(id.String())
}

func MustFromId(id fmt.Stringer) primitive.ObjectID { //这不是有点脱裤子放屁么？在外边判断err有那么困难么？
	oid,err := FromID(id)
	if err != nil {
		panic(err)
	}
	return oid
}


func ToAccountID(objId primitive.ObjectID) id.AccountID {
			return id.AccountID(objId.Hex())
}

func ToTripID(objId primitive.ObjectID) id.TripID {
	return id.TripID(objId.Hex())
}