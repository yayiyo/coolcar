package sim

import (
	"context"
	"time"

	carpb "coolcar/car/api/gen/v1"
	"coolcar/car/mq"
	"go.uber.org/zap"
)

type Controller struct {
	CarService carpb.CarServiceClient
	mq.Subscriber
	Logger *zap.Logger
}

func (c *Controller) RunSimulations(ctx context.Context) {
	var cars []*carpb.CarEntity
	for {
		time.Sleep(5 * time.Second)
		res, err := c.CarService.GetCars(ctx, &carpb.GetCarsRequest{})
		if err != nil {
			c.Logger.Error("failed to get cars", zap.Error(err))
			continue
		}
		cars = res.Cars
		break
	}

	msgCh, cleanUp, err := c.Subscribe(ctx)
	if err != nil {
		panic(err)
	}
	defer cleanUp()

	carChs := make(map[string]chan *carpb.Car)
	for _, car := range cars {
		ch := make(chan *carpb.Car)
		carChs[car.Id] = ch
		go c.SimulateCar(ctx, car, ch)
	}

	for carUpdate := range msgCh {
		ch := carChs[carUpdate.Id]
		if ch != nil {
			ch <- carUpdate.Car
		}
	}
}

func (c *Controller) SimulateCar(ctx context.Context, initial *carpb.CarEntity, ch chan *carpb.Car) {
	carID := initial.Id

	for update := range ch {
		if update.Status == carpb.CarStatus_UNLOCKING {
			_, err := c.CarService.UpdateCar(ctx, &carpb.UpdateCarRequest{
				Id:     carID,
				Status: carpb.CarStatus_UNLOCKED,
			})
			if err != nil {
				c.Logger.Error("Error updating CarStatus_UNLOCKED", zap.Error(err))
			}
		} else if carpb.CarStatus_LOCKING == update.Status {
			_, err := c.CarService.UpdateCar(ctx, &carpb.UpdateCarRequest{
				Id:     carID,
				Status: carpb.CarStatus_LOCKED,
			})
			if err != nil {
				c.Logger.Error("Error updating CarStatus_LOCKED", zap.Error(err))
			}
		}
	}
}
