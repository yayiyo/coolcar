package poi

import (
	"context"

	rentalpb "coolcar/rental/api/gen/v1"
)

type Manager struct {
}

func (p *Manager) Resolve(ct context.Context, location *rentalpb.Location) (string, error) {
	return "北京天安门", nil
}
