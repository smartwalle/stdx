package timex

import (
	"time"
)

type Timezone interface {
	Location() *time.Location
}

var UTC = time.UTC

var Local = time.Local

func LoadLocation(name string) (*time.Location, error) {
	return time.LoadLocation(name)
}
