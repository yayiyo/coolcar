package car

import (
	"context"
	"fmt"

	carpb "coolcar/car/api/gen/v1"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
)

type Manager struct {
	CarService carpb.CarServiceClient
}

func (m *Manager) Lock(ctx context.Context, carID id.CarID) error {
	_, err := m.CarService.LockCar(ctx, &carpb.LockCarRequest{
		Id: carID.String(),
	})
	if err != nil {
		return fmt.Errorf("error locking car: %v", err)
	}
	return nil
}

func (m *Manager) Verify(ctx context.Context, carID id.CarID, location *rentalpb.Location) error {
	car, err := m.CarService.GetCar(ctx, &carpb.GetCarRequest{
		Id: carID.String(),
	})
	if err != nil {
		return fmt.Errorf("can not get car: %v", err)
	}

	if car.Status != carpb.CarStatus_LOCKED {
		return fmt.Errorf("can not unlock, car status is %v", car.Status)
	}
	return nil
}

func (m *Manager) Unlock(ctx context.Context, carID id.CarID, aid id.AccountID, tid id.TripID, avatar string) error {
	_, err := m.CarService.UnlockCar(ctx, &carpb.UnlockCarRequest{
		Id: carID.String(),
		Driver: &carpb.Driver{
			Id:        aid.String(),
			AvatarUrl: avatar,
		},
		TripId: tid.String(),
	})
	if err != nil {
		return fmt.Errorf("unlock car %s failed: %v", carID.String(), err)
	}
	return nil
}
