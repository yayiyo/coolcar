package car

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
)

type Manager struct {
}

func (m *Manager) Verify(ctx context.Context, carID id.CarID, location *rentalpb.Location) error {
	return nil
}

func (m *Manager) Unlock(ctx context.Context, carID id.CarID) error {
	return nil
}
