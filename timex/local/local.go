package local

import (
	"github.com/smartwalle/stdx/timex"
	"time"
)

func Location() *time.Location {
	return time.Local
}

type Timezone struct {
}

func (Timezone) Location() *time.Location {
	return time.Local
}

type Time = timex.Time[Timezone]

func Now() Time {
	return timex.Now[Timezone]()
}

func Date(year int, month time.Month, day, hour, min, sec, nsec int) Time {
	return timex.Date[Timezone](year, month, day, hour, min, sec, nsec)
}

func Parse(layout, value string) (Time, error) {
	return timex.Parse[Timezone](layout, value)
}

func Unix(sec int64, nsec int64) Time {
	return timex.Unix[Timezone](sec, nsec)
}

func UnixMicro(usec int64) Time {
	return timex.UnixMicro[Timezone](usec)
}

func UnixMilli(msec int64) Time {
	return timex.UnixMilli[Timezone](msec)
}

func FromTime(t timex.UTCTime) Time {
	return timex.FromTime[Timezone](t)
}
