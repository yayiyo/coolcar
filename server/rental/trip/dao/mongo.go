package dao

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"
	mgo "coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/mongo"
)

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
