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

func MustLoadLocation(name string) *time.Location {
	var loc, err = LoadLocation(name)
	if err != nil {
		panic(err)
	}
	return loc
}
