package dao

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	mgo "coolcar/shared/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	accountIDField      = "account_id"
	profileField        = "profile"
	identityStatusField = profileField + ".identitystatus"
	photoBlobUDField    = "photo_blob_id"
)

type Mongo struct {
	col *mongo.Collection
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("profile"),
	}
}

type ProfileRecord struct {
	AccountID   string            `bson:"account_id"`
	Profile     *rentalpb.Profile `bson:"profile"`
	PhotoBlobID string            `bson:"photo_blob_id"`
}

func (m *Mongo) GetProfile(ctx context.Context, aid id.AccountID) (*ProfileRecord, error) {
	p := &ProfileRecord{}
	err := m.col.FindOne(ctx, byAccountID(aid)).Decode(p)
	if err != nil {
		return nil, err
	}
	return p, nil
}

func (m *Mongo) UpdateProfile(ctx context.Context, aid id.AccountID, prevState rentalpb.IdentityStatus, p *rentalpb.Profile) error {
	filter := bson.M{
		identityStatusField: prevState,
	}

	if prevState == rentalpb.IdentityStatus_NOT_SUBMITTED {
		filter = mgo.ZeroOrNotExists(identityStatusField, 0)
	}
	filter[accountIDField] = aid.String()

	_, err := m.col.UpdateOne(ctx, filter, mgo.Set(bson.M{
		accountIDField: aid.String(),
		profileField:   p,
	}), options.Update().SetUpsert(true))

	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) UpdateProfilePhoto(ctx context.Context, aid id.AccountID, bid id.BlobID) error {
	_, err := m.col.UpdateOne(ctx, bson.M{
		accountIDField: aid.String(),
	}, mgo.Set(bson.M{
		accountIDField:   aid.String(),
		photoBlobUDField: bid.String(),
	}), options.Update().SetUpsert(true))

	if err != nil {
		return err
	}

	return nil
}

func byAccountID(aid id.AccountID) bson.M {
	return bson.M{accountIDField: aid.String()}
}
