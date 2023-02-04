package dao

import (
	"context"
	"fmt"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var accountIDField = "trip.accountid"

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("trip"),
	}
}

type TripRecord struct {
	mgo.IDField       `bson:"inline"`
	mgo.UpdateAtField `bson:"inline"`
	Trip              *rentalpb.Trip `bson:"trip"`
}

// CreateTrip creates a new trip record.
func (m *Mongo) CreateTrip(ctx context.Context, trip *rentalpb.Trip) (*TripRecord, error) {
	r := &TripRecord{
		Trip: trip,
	}
	r.ID = mgo.NewObjID()
	r.UpdateAt = mgo.UpdatedAt()
	_, err := m.col.InsertOne(ctx, r)
	if err != nil {
		return nil, err
	}

	return r, nil
}

// GetTrip returns a new trip record
func (m *Mongo) GetTrip(ctx context.Context, tripID id.TripID, accountID id.AccountID) (*TripRecord, error) {
	objID, err := objid.FromID(tripID)
	if err != nil {
		return nil, fmt.Errorf("id invalid")
	}

	r := &TripRecord{}
	err = m.col.FindOne(ctx, bson.M{mgo.IDFieldName: objID, accountIDField: accountID}).Decode(r)
	if err != nil {
		return nil, err
	}

	return r, nil
}
