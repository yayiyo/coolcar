package dao

import (
	"context"
	"fmt"

	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"coolcar/shared/mongo/objid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("blob"),
	}
}

type BlobRecord struct {
	mgo.IDField `bson:"inline"`
	AccountID   string `bson:"account_id"`
	Path        string `bson:"path"`
}

func (m *Mongo) CreateBlob(ctx context.Context, aid id.AccountID) (*BlobRecord, error) {
	br := &BlobRecord{
		AccountID: aid.String(),
	}

	objID := mgo.NewObjID()
	br.ID = objID
	br.Path = fmt.Sprintf("%s/%s", aid.String(), objID.Hex())

	_, err := m.col.InsertOne(ctx, br)
	if err != nil {
		return nil, err
	}
	return br, nil
}

func (m *Mongo) GetBlob(ctx context.Context, bid id.BlobID) (*BlobRecord, error) {
	objID, err := objid.FromID(bid)
	if err != nil {
		return nil, fmt.Errorf("invalid blob ID: %v", err)
	}

	br := &BlobRecord{}
	err = m.col.FindOne(ctx, bson.M{
		mgo.IDFieldName: objID,
	}).Decode(br)
	if err != nil {
		return nil, err
	}
	return br, nil
}
