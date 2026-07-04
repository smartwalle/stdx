package timex

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type UTCTime interface {
	UTC() time.Time
}

type Time[T Timezone] struct {
	utc time.Time
}

func Now[T Timezone]() Time[T] {
	return Time[T]{utc: time.Now().UTC()}
}

func Date[T Timezone](year int, month time.Month, day, hour, min, sec, nsec int) Time[T] {
	var zone T
	var rTime = time.Date(year, month, day, hour, min, sec, nsec, zone.Location())
	return Time[T]{utc: rTime.UTC()}
}

func Parse[T Timezone](layout, value string) (Time[T], error) {
	var zone T
	rTime, err := time.ParseInLocation(layout, value, zone.Location())
	if err != nil {
		return Time[T]{}, err
	}
	return Time[T]{utc: rTime.UTC()}, nil
}

func Unix[T Timezone](sec int64, nsec int64) Time[T] {
	return Time[T]{utc: time.Unix(sec, nsec).UTC()}
}

func UnixMicro[T Timezone](usec int64) Time[T] {
	return Time[T]{utc: time.UnixMicro(usec).UTC()}
}

func UnixMilli[T Timezone](msec int64) Time[T] {
	return Time[T]{utc: time.UnixMilli(msec).UTC()}
}

func FromTime[T Timezone](t UTCTime) Time[T] {
	return Time[T]{utc: t.UTC()}
}

func (t Time[T]) IsZero() bool {
	return t.utc.IsZero()
}

func (t Time[T]) After(u UTCTime) bool {
	return t.utc.After(u.UTC())
}

func (t Time[T]) Before(u UTCTime) bool {
	return t.utc.Before(u.UTC())
}

func (t Time[T]) Compare(u UTCTime) int {
	return t.utc.Compare(u.UTC())
}

func (t Time[T]) Equal(u UTCTime) bool {
	return t.utc.Equal(u.UTC())
}

func (t Time[T]) InRange(u1 UTCTime, u2 UTCTime) bool {
	var start = u1.UTC()
	var end = u2.UTC()
	if start.After(end) {
		start, end = end, start
	}
	var current = t.UTC()
	return !current.Before(start) && !current.After(end)
}

func (t Time[T]) Date() (year int, month time.Month, day int) {
	return t.Time().Date()
}

func (t Time[T]) Year() int {
	return t.Time().Year()
}

func (t Time[T]) Quarter() int {
	return int(t.Time().Month()-1)/3 + 1
}

func (t Time[T]) Month() time.Month {
	return t.Time().Month()
}

func (t Time[T]) Day() int {
	return t.Time().Day()
}

func (t Time[T]) Weekday() time.Weekday {
	return t.Time().Weekday()
}

func (t Time[T]) ISOWeek() (year, week int) {
	return t.Time().ISOWeek()
}

func (t Time[T]) Clock() (hour, min, sec int) {
	return t.Time().Clock()
}

func (t Time[T]) Hour() int {
	return t.Time().Hour()
}

func (t Time[T]) Minute() int {
	return t.Time().Minute()
}

func (t Time[T]) Second() int {
	return t.Time().Second()
}

func (t Time[T]) Nanosecond() int {
	return t.Time().Nanosecond()
}

func (t Time[T]) YearDay() int {
	return t.Time().YearDay()
}

func (t Time[T]) String() string {
	return t.Time().String()
}

func (t Time[T]) Add(d time.Duration) Time[T] {
	return Time[T]{utc: t.utc.Add(d)}
}

func (t Time[T]) Sub(u UTCTime) time.Duration {
	return t.utc.Sub(u.UTC())
}

func (t Time[T]) AddDate(years int, months int, days int) Time[T] {
	return Time[T]{utc: t.Time().AddDate(years, months, days).UTC()}
}

func (t Time[T]) UTC() time.Time {
	return t.utc.UTC()
}

func (t Time[T]) Time() time.Time {
	if t.utc.IsZero() {
		return t.utc
	}
	return t.utc.In(t.Location())
}

func (t Time[T]) Local() time.Time {
	if t.utc.IsZero() {
		return t.utc
	}
	return t.utc.Local()
}

func (t Time[T]) In(loc *time.Location) time.Time {
	if t.utc.IsZero() {
		return t.utc
	}
	return t.utc.In(loc)
}

func (t Time[T]) Location() *time.Location {
	var zone T
	return zone.Location()
}

func (t Time[T]) calendar() Calendar {
	return NewCalendar(t.Location())
}

func (t Time[T]) fromTime(v time.Time) Time[T] {
	return Time[T]{utc: v.UTC()}
}

func (t Time[T]) Zone() (name string, offset int) {
	return t.Time().Zone()
}

func (t Time[T]) ZoneBounds() (start, end time.Time) {
	return t.Time().ZoneBounds()
}

func (t Time[T]) Unix() int64 {
	return t.utc.Unix()
}

func (t Time[T]) UnixMilli() int64 {
	return t.utc.UnixMilli()
}

func (t Time[T]) UnixMicro() int64 {
	return t.utc.UnixMicro()
}

func (t Time[T]) UnixNano() int64 {
	return t.utc.UnixNano()
}

func (t Time[T]) MarshalBinary() ([]byte, error) {
	return t.Time().MarshalBinary()
}

func (t *Time[T]) UnmarshalBinary(data []byte) error {
	if err := t.utc.UnmarshalBinary(data); err != nil {
		return err
	}
	t.utc = t.utc.UTC()
	return nil
}

func (t Time[T]) GobEncode() ([]byte, error) {
	return t.Time().GobEncode()
}

func (t *Time[T]) GobDecode(data []byte) error {
	if err := t.utc.GobDecode(data); err != nil {
		return err
	}
	t.utc = t.utc.UTC()
	return nil
}

func (t Time[T]) MarshalJSON() ([]byte, error) {
	return t.Time().MarshalJSON()
}

func (t *Time[T]) UnmarshalJSON(data []byte) error {
	if err := t.utc.UnmarshalJSON(data); err != nil {
		return err
	}
	t.utc = t.utc.UTC()
	return nil
}

func (t Time[T]) MarshalText() ([]byte, error) {
	return t.Time().MarshalText()
}

func (t *Time[T]) UnmarshalText(data []byte) error {
	if err := t.utc.UnmarshalText(data); err != nil {
		return err
	}
	t.utc = t.utc.UTC()
	return nil
}

func (t Time[T]) IsDST() bool {
	return t.Time().IsDST()
}

func (t Time[T]) Truncate(d time.Duration) Time[T] {
	return Time[T]{utc: t.utc.Truncate(d)}
}

func (t Time[T]) Round(d time.Duration) Time[T] {
	return Time[T]{utc: t.utc.Round(d)}
}

func (t Time[T]) Format(layout string) string {
	return t.Time().Format(layout)
}

func (t Time[T]) AppendFormat(b []byte, layout string) []byte {
	return t.Time().AppendFormat(b, layout)
}

func (t Time[T]) GoString() string {
	return t.Time().GoString()
}

func (t Time[T]) Value() (driver.Value, error) {
	if t.IsZero() {
		return nil, nil
	}
	return t.utc, nil
}

func (t *Time[T]) Scan(value interface{}) (err error) {
	if t == nil {
		return fmt.Errorf("timex: scan on nil Time")
	}
	switch val := value.(type) {
	case time.Time:
		t.utc = val.UTC()
		return nil
	case *time.Time:
		if val == nil {
			t.utc = time.Time{}
			return nil
		}
		t.utc = (*val).UTC()
		return nil
	case Time[T]:
		t.utc = val.UTC()
		return nil
	case *Time[T]:
		if val == nil {
			t.utc = time.Time{}
			return nil
		}
		t.utc = val.UTC()
		return nil
	case nil:
		t.utc = time.Time{}
		return nil
	default:
		return fmt.Errorf("timex: scanning unsupported type: %T", value)
	}
}

// Previous 获取当前日期的前一天（昨天）
func (t Time[T]) Previous() Time[T] {
	return t.AddDate(0, 0, -1)
}

// Next 获取当前日期的后一天（明天）
func (t Time[T]) Next() Time[T] {
	return t.AddDate(0, 0, 1)
}

// BeginningOfMinute 获取当前分钟的开始时间
func (t Time[T]) BeginningOfMinute() Time[T] {
	return t.fromTime(t.calendar().BeginningOfMinuteAt(t.Time()))
}

// EndOfMinute 获取当前分钟的结束时间
func (t Time[T]) EndOfMinute() Time[T] {
	return t.fromTime(t.calendar().EndOfMinuteAt(t.Time()))
}

// BeginningOfHour 获取当前小时的开始时间
func (t Time[T]) BeginningOfHour() Time[T] {
	return t.fromTime(t.calendar().BeginningOfHourAt(t.Time()))
}

// EndOfHour 获取当前小时的结束时间
func (t Time[T]) EndOfHour() Time[T] {
	return t.fromTime(t.calendar().EndOfHourAt(t.Time()))
}

// BeginningOfDay 获取当前天的开始时间
func (t Time[T]) BeginningOfDay() Time[T] {
	return t.fromTime(t.calendar().BeginningOfDayAt(t.Time()))
}

// EndOfDay 获取当前天的结束时间
func (t Time[T]) EndOfDay() Time[T] {
	return t.fromTime(t.calendar().EndOfDayAt(t.Time()))
}

// BeginningOfWeek 获取当前日期所在周的开始时间
func (t Time[T]) BeginningOfWeek() Time[T] {
	return t.fromTime(t.calendar().BeginningOfWeekAt(t.Time()))
}

// EndOfWeek 获取当前日期所在周的结束时间
func (t Time[T]) EndOfWeek() Time[T] {
	return t.fromTime(t.calendar().EndOfWeekAt(t.Time()))
}

// BeginningOfMonth 获取当前日期所在月的开始时间
func (t Time[T]) BeginningOfMonth() Time[T] {
	return t.fromTime(t.calendar().BeginningOfMonthAt(t.Time()))
}

// EndOfMonth 获取当前日期所在月的结束时间
func (t Time[T]) EndOfMonth() Time[T] {
	return t.fromTime(t.calendar().EndOfMonthAt(t.Time()))
}

// BeginningOfQuarter 获取当前日期所在季度的开始时间
func (t Time[T]) BeginningOfQuarter() Time[T] {
	return t.fromTime(t.calendar().BeginningOfQuarterAt(t.Time()))
}

// EndOfQuarter 获取当前日期所在季度的结束时间
func (t Time[T]) EndOfQuarter() Time[T] {
	return t.fromTime(t.calendar().EndOfQuarterAt(t.Time()))
}

// BeginningOfYear 获取当前日期所在年的开始时间
func (t Time[T]) BeginningOfYear() Time[T] {
	return t.fromTime(t.calendar().BeginningOfYearAt(t.Time()))
}

// EndOfYear 获取当前日期所在年的结束时间
func (t Time[T]) EndOfYear() Time[T] {
	return t.fromTime(t.calendar().EndOfYearAt(t.Time()))
}

func LeapYear(year int) bool {
	mask := 0xf
	if year%25 != 0 {
		mask = 3
	}
	return year&mask == 0
}

// DaysInMonth 获取指定月份有多少天
func DaysInMonth(year int, month time.Month) (number int) {
	if month == time.February {
		if LeapYear(year) {
			return 29
		}
		return 28
	}
	return 30 + int((month+month>>3)&1)
}

// InRange 判断时间是否在指定时间范围内(包含边界值)
func InRange(t UTCTime, u1 UTCTime, u2 UTCTime) bool {
	var start = u1.UTC()
	var end = u2.UTC()
	if start.After(end) {
		start, end = end, start
	}
	var current = t.UTC()
	return !current.Before(start) && !current.After(end)
}
