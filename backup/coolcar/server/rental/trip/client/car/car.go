package car

import (
	"context"
	rentalpb "coolcar/rental/api/gen/v1"
	"coolcar/shared/id"
)

type Manager struct {

}

func (m *Manager)Verify(ctx context.Context,id id.CarID,lc *rentalpb.Location) error {


	return nil
}

func (m *Manager)Unlock(context.Context,id.CarID,id.AccountID,id.TripID, string) error  {

	return nil
}

func (m *Manager)Lock(c context.Context,cid id.CarID) error  {


	return nil
}