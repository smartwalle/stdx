package netx_test

import (
	"github.com/smartwalle/stdx/netx"
	"testing"
)

func TestExternalIP(t *testing.T) {
	t.Log(netx.ExternalIPv4())
	t.Log(netx.ExternalIPv6())
}

func TestInternalIP(t *testing.T) {
	t.Log(netx.InternalIPv4())
	t.Log(netx.InternalIPv4s())
}
