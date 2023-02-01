package dao

import (
	"context"

	mgo "coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	openIDFiled = "open_id"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("account"),
	}
}

func (m *Mongo) ResolveAccountID(ctx context.Context, openID string) (string, error) {
	res := m.col.FindOneAndUpdate(ctx,
		bson.M{openIDFiled: openID},
		mgo.SetOnInsert(bson.M{
			mgo.IDFieldName: mgo.NewObjID(),
			openIDFiled:     openID,
		}), options.FindOneAndUpdate().SetUpsert(true).
			SetReturnDocument(options.After))
	if err := res.Err(); err != nil {
		return "", err
	}

	var row mgo.IDField
	err := res.Decode(&row)
	if err != nil {
		return "", err
	}
	return row.ID.Hex(), nil
}
