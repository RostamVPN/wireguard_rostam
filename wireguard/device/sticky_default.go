//go:build !linux

package device

import (
	"github.com/wwwiretap/amnezia-wg/conn"
	"github.com/wwwiretap/amnezia-wg/rwcancel"
)

func (device *Device) startRouteListener(bind conn.Bind) (*rwcancel.RWCancel, error) {
	return nil, nil
}
