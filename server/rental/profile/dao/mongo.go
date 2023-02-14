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
	AccountID string            `bson:"account_id"`
	Profile   *rentalpb.Profile `bson:"profile"`
}

func (m *Mongo) GetProfile(ctx context.Context, aid id.AccountID) (*rentalpb.Profile, error) {
	p := &ProfileRecord{}
	err := m.col.FindOne(ctx, byAccountID(aid)).Decode(p)
	if err != nil {
		return nil, err
	}
	return p.Profile, nil
}

func (m *Mongo) UpdateProfile(ctx context.Context, aid id.AccountID, prevState rentalpb.IdentityStatus, p *rentalpb.Profile) error {
	_, err := m.col.UpdateOne(ctx, bson.M{
		accountIDField:      aid.String(),
		identityStatusField: prevState,
	}, mgo.Set(bson.M{
		accountIDField: aid.String(),
		profileField:   p,
	}), options.Update().SetUpsert(true))

	if err != nil {
		return err
	}

	return nil
}

func byAccountID(aid id.AccountID) bson.M {
	return bson.M{accountIDField: aid.String()}
}
