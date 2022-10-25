package dao

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"coolcar/shared/mgoutil"
	"coolcar/shared/str_obj"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	tripField = "trip"
	accountIdField = tripField + ".accountid"
	statusField = tripField + ".status"
)

type Mongo struct {
	col *mongo.Collection
	//newstr_obj func() primitive.ObjectID
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col:      db.Collection("trip"),
		//newstr_obj: primitive.NewObjectID,
	}
}

type TripRecord struct {
	mgoutil.IDField       `bson:"inline"`
	mgoutil.UpdateAtField `bson:"inline"`
	Trip                 *rentalpb.Trip `bson:"trip"`
}

//todo 同一个account 最多只能有一个进行中的Trip
//todo 强类型化tripID
//todo 表格驱动测试

func (m *Mongo)CreateTrip(c context.Context,trip *rentalpb.Trip)(*TripRecord,error)  {
	//var t TripRecord
	r := &TripRecord{
		Trip:          trip,
	}
	r.ID = mgoutil.NewObjId()
	r.UpdateAt = mgoutil.UpdateAt()
	_, err := m.col.InsertOne(c, r)
	if err != nil {
		return nil, err
	}
	return r,nil
}

func (m *Mongo)GetTrip(c context.Context,id id.TripID,accountId id.AccountID) (*TripRecord,error) {
	ObjId,err := str_obj.FromID(id)
	if err != nil {
		return nil, err
	}
	one := m.col.FindOne(c, bson.M{
		mgoutil.IDFieldName: ObjId,
		accountIdField:     accountId,
	})
	if err := one.Err(); err != nil {
		return nil, err
	}

	var tr TripRecord
	err = one.Decode(&tr)
	if err != nil {
		return nil, err
	}
	return &tr,nil
}

func (m *Mongo)GetTrips(c context.Context,accountId id.AccountID,status rentalpb.TripStatus) ([]*TripRecord,error) {
	filter := bson.M{
		accountIdField: accountId.String(),
	}
	if status != rentalpb.TripStatus_TS_NOT_SPECIFIED {
		filter[statusField] = status
	}
	res ,err := m.col.Find(c,filter,&options.FindOptions{
		Sort: bson.M{
			mgoutil.IDFieldName: -1,//-1从大到小
		},
	})
	//m.col.Find(c,filter,options.Find().SetSort(bson.M{
	//	mgoutil.IDFieldName: -1,//-1从大到小
	//}))

	if err != nil {
		return nil, err
	}
	var trips []*TripRecord
	for res.Next(c) {
		var trip TripRecord
		err := res.Decode(&trip)
		if err != nil {
			return  nil,err
		}
		trips = append(trips,&trip)
	}
	return trips,nil
}


//Update Trip updates a trip ,乐观锁更新
func (m *Mongo)UpdateTrip(c context.Context,tid id.TripID,aid id.AccountID,updateAt int64,trip *rentalpb.Trip) error {
	objId,err := str_obj.FromID(tid)
	if err != nil {
		return err
	}

	newUpdateAt := mgoutil.UpdateAt()

	res, err := m.col.UpdateOne(c, bson.M{
		mgoutil.IDFieldName:       objId,
		accountIdField:           aid.String(),
		mgoutil.UpdatedAtFieldName: updateAt,
	}, mgoutil.Set(bson.M{
		tripField:                trip,
		mgoutil.UpdatedAtFieldName: newUpdateAt,
	}))

	if res == nil {
		return fmt.Errorf("update faild")
	}

	if err != nil {
		return err
	}

	if res.MatchedCount == 0 {
		return  mongo.ErrNoDocuments
	}
	return nil
}