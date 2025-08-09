package netx_test

import (
	"github.com/smartwalle/stdx/netx"
	"testing"
)

func TestGetHardwareAddr(t *testing.T) {
	t.Log(netx.HardwareAddr())
	t.Log(netx.HardwareAddrs())
}
