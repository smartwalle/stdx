package timex

import (
	"database/sql/driver"
	"fmt"
	"time"
)

type Time[T Timezone] struct {
	utc time.Time
}

func Now[T Timezone]() Time[T] {
	return Time[T]{utc: time.Now()}
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
	return Time[T]{utc: time.UnixMicro(usec)}
}

func UnixMilli[T Timezone](msec int64) Time[T] {
	return Time[T]{utc: time.UnixMilli(msec)}
}

func FromTime[T Timezone](t time.Time) Time[T] {
	return Time[T]{utc: t.UTC()}
}

func (t Time[T]) IsZero() bool {
	return t.utc.IsZero()
}

func (t Time[T]) After(u Time[T]) bool {
	return t.utc.After(u.utc)
}

func (t Time[T]) Before(u Time[T]) bool {
	return t.utc.Before(u.utc)
}

func (t Time[T]) Compare(u Time[T]) int {
	return t.utc.Compare(u.utc)
}

func (t Time[T]) Equal(u Time[T]) bool {
	return t.utc.Equal(u.utc)
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

func (t Time[T]) Sub(u Time[T]) time.Duration {
	return t.utc.Sub(u.utc)
}

func (t Time[T]) AddDate(years int, months int, days int) Time[T] {
	return Time[T]{utc: t.utc.AddDate(years, months, days)}
}

func (t Time[T]) UTC() time.Time {
	return t.utc.UTC()
}

func (t Time[T]) Time() time.Time {
	return t.utc.In(t.Location())
}

func (t Time[T]) Local() time.Time {
	return t.utc.Local()
}

func (t Time[T]) In(loc *time.Location) time.Time {
	return t.utc.In(loc)
}

func (t Time[T]) Location() *time.Location {
	var zone T
	return zone.Location()
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
	return t.utc.UnmarshalBinary(data)
}

func (t Time[T]) GobEncode() ([]byte, error) {
	return t.Time().GobEncode()
}

func (t *Time[T]) GobDecode(data []byte) error {
	return t.utc.GobDecode(data)
}

func (t Time[T]) MarshalJSON() ([]byte, error) {
	return t.Time().MarshalJSON()
}

func (t *Time[T]) UnmarshalJSON(data []byte) error {
	return t.utc.UnmarshalJSON(data)
}

func (t Time[T]) MarshalText() ([]byte, error) {
	return t.Time().MarshalText()
}

func (t *Time[T]) UnmarshalText(data []byte) error {
	return t.utc.UnmarshalText(data)
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
	switch val := value.(type) {
	case time.Time:
		t.utc = val.UTC()
		return nil
	case *time.Time:
		t.utc = (*val).UTC()
		return nil
	case Time[T]:
		t.utc = val.utc
		return nil
	case *Time[T]:
		t.utc = val.utc
		return nil
	case nil:
		return nil
	default:
		return fmt.Errorf("timex: scanning unsupported type: %T", value)
	}
}

// Previous 获取当前日期的前一天（昨天）
func (t Time[T]) Previous() Time[T] {
	return Time[T]{utc: t.utc.Add(time.Hour * -24)}
}

// Next 获取当前日期的后一天（明天）
func (t Time[T]) Next() Time[T] {
	return Time[T]{utc: t.utc.Add(time.Hour * 24)}
}

// BeginningOfMinute 获取当前分钟的开始时间
func (t Time[T]) BeginningOfMinute() Time[T] {
	return Date[T](t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 0, 0)
}

// EndOfMinute 获取当前分钟的结束时间
func (t Time[T]) EndOfMinute() Time[T] {
	return Date[T](t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), 59, int(time.Second-time.Nanosecond))
}

// BeginningOfHour 获取当前小时的开始时间
func (t Time[T]) BeginningOfHour() Time[T] {
	return Date[T](t.Year(), t.Month(), t.Day(), t.Hour(), 0, 0, 0)
}

// EndOfHour 获取当前小时的结束时间
func (t Time[T]) EndOfHour() Time[T] {
	return Date[T](t.Year(), t.Month(), t.Day(), t.Hour(), 59, 59, int(time.Second-time.Nanosecond))
}

// BeginningOfDay 获取当前天的开始时间
func (t Time[T]) BeginningOfDay() Time[T] {
	return Date[T](t.Year(), t.Month(), t.Day(), 0, 0, 0, 0)
}

// EndOfDay 获取当前天的结束时间
func (t Time[T]) EndOfDay() Time[T] {
	return Date[T](t.Year(), t.Month(), t.Day(), 23, 59, 59, int(time.Second-time.Nanosecond))
}

// BeginningOfWeek 获取当前日期所在周的开始时间
func (t Time[T]) BeginningOfWeek() Time[T] {
	var w = t.Weekday()
	var d = int(w - time.Sunday)
	return Date[T](t.Year(), t.Month(), t.Day()-d, 0, 0, 0, 0)
}

// EndOfWeek 获取当前日期所在周的结束时间
func (t Time[T]) EndOfWeek() Time[T] {
	var w = t.Weekday()
	var d = int(time.Saturday - w)
	return Date[T](t.Year(), t.Month(), t.Day()+d, 23, 59, 59, int(time.Second-time.Nanosecond))
}

// BeginningOfMonth 获取当前日期所在月的开始时间
func (t Time[T]) BeginningOfMonth() Time[T] {
	return Date[T](t.Year(), t.Month(), 1, 0, 0, 0, 0)
}

// EndOfMonth 获取当前日期所在月的结束时间
func (t Time[T]) EndOfMonth() Time[T] {
	return Date[T](t.Year(), t.Month(), DaysInMonth(t.Year(), t.Month()), 23, 59, 59, int(time.Second-time.Nanosecond))
}

// BeginningOfQuarter 获取当前日期所在季度的开始时间
func (t Time[T]) BeginningOfQuarter() Time[T] {
	var m = int(t.Month()-1)/3*3 + 1
	return Date[T](t.Year(), time.Month(m), 1, 0, 0, 0, 0)
}

// EndOfQuarter 获取当前日期所在季度的结束时间
func (t Time[T]) EndOfQuarter() Time[T] {
	var m = time.Month(int(t.Month()-1)/3*3 + 3)
	return Date[T](t.Year(), m, DaysInMonth(t.Year(), m), 23, 59, 59, int(time.Second-time.Nanosecond))
}

// BeginningOfYear 获取当前日期所在年的开始时间
func (t Time[T]) BeginningOfYear() Time[T] {
	return Date[T](t.Year(), time.January, 1, 0, 0, 0, 0)
}

// EndOfYear 获取当前日期所在年的结束时间
func (t Time[T]) EndOfYear() Time[T] {
	return Date[T](t.Year(), time.December, DaysInMonth(t.Year(), time.December), 23, 59, 59, int(time.Second-time.Nanosecond))
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

// BeginningOfDay 获取指定日期的开始时间
func BeginningOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.Local)
}

// EndOfDay 获取指定日期的结束时间
func EndOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
}

// BeginningOfWeek 获取指定日期所在周的开始时间
func BeginningOfWeek(year int, month time.Month, day int) time.Time {
	var t = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	var w = t.Weekday()
	var d = int(w - time.Sunday)
	return time.Date(year, month, day-d, 0, 0, 0, 0, time.Local)
}

// EndOfWeek 获取指定日期所在周的结束时间
func EndOfWeek(year int, month time.Month, day int) time.Time {
	var t = time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	var w = t.Weekday()
	var d = int(time.Saturday - w)
	return time.Date(year, month, day+d, 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
}

// BeginningOfMonth 获取指定月份的开始时间
func BeginningOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, time.Local)
}

// EndOfMonth 获取指定月份的结束时间
func EndOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, DaysInMonth(year, month), 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
}

// BeginningOfQuarter 获取指定季度的开始时间
func BeginningOfQuarter(year int, quarter int) time.Time {
	var m = time.Month((quarter-1)*3 + 1)
	return time.Date(year, m, 1, 0, 0, 0, 0, time.Local)
}

// EndOfQuarter 获取指定季度的结束时间
func EndOfQuarter(year int, quarter int) time.Time {
	var m = time.Month((quarter-1)*3 + 3)
	return time.Date(year, m, DaysInMonth(year, m), 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
}

// BeginningOfYear 获取指定年份的开始时间
func BeginningOfYear(year int) time.Time {
	return time.Date(year, time.January, 1, 0, 0, 0, 0, time.Local)
}

// EndOfYear 获取指定年份的结束时间
func EndOfYear(year int) time.Time {
	return time.Date(year, time.December, DaysInMonth(year, time.December), 23, 59, 59, int(time.Second-time.Nanosecond), time.Local)
}
