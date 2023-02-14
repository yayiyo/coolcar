package profile

import (
	"context"
	"encoding/base64"
	"fmt"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
	"google.golang.org/protobuf/proto"
)

type Fetcher interface {
	GetProfile(ctx context.Context, req *rentalpb.GetProfileRequest) (*rentalpb.Profile, error)
}

type Manager struct {
	Fetcher
}

func (m *Manager) Verify(ctx context.Context, aid id.AccountID) (id.IdentityID, error) {
	nilID := id.IdentityID("")
	profile, err := m.Fetcher.GetProfile(ctx, &rentalpb.GetProfileRequest{})
	if err != nil {
		return nilID, err
	}

	if profile.IdentityStatus != rentalpb.IdentityStatus_VERIFIED {
		return nilID, fmt.Errorf("invalid indentity status")
	}

	b, err := proto.Marshal(profile.Identity)
	if err != nil {
		return nilID, fmt.Errorf("marshal identity error: %+v", err)
	}

	return id.IdentityID(base64.StdEncoding.EncodeToString(b)), nil
}
