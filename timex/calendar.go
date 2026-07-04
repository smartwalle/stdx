package timex

import "time"

type Calendar struct {
	loc *time.Location
}

func NewCalendar(loc *time.Location) Calendar {
	if loc == nil {
		loc = time.Local
	}
	return Calendar{loc: loc}
}

func (c Calendar) Location() *time.Location {
	if c.loc == nil {
		return time.Local
	}
	return c.loc
}

// BeginningOfDay 获取当前日历时区指定日期的开始时间
func (c Calendar) BeginningOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, c.Location())
}

// EndOfDay 获取当前日历时区指定日期的结束时间
func (c Calendar) EndOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// BeginningOfWeek 获取当前日历时区指定日期所在周的开始时间
func (c Calendar) BeginningOfWeek(year int, month time.Month, day int) time.Time {
	var t = time.Date(year, month, day, 0, 0, 0, 0, c.Location())
	var w = t.Weekday()
	var d = int(w - time.Sunday)
	return time.Date(year, month, day-d, 0, 0, 0, 0, c.Location())
}

// EndOfWeek 获取当前日历时区指定日期所在周的结束时间
func (c Calendar) EndOfWeek(year int, month time.Month, day int) time.Time {
	var t = time.Date(year, month, day, 0, 0, 0, 0, c.Location())
	var w = t.Weekday()
	var d = int(time.Saturday - w)
	return time.Date(year, month, day+d, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// BeginningOfMonth 获取当前日历时区指定月份的开始时间
func (c Calendar) BeginningOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, c.Location())
}

// EndOfMonth 获取当前日历时区指定月份的结束时间
func (c Calendar) EndOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month+1, 0, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// BeginningOfQuarter 获取当前日历时区指定季度的开始时间
func (c Calendar) BeginningOfQuarter(year int, quarter int) time.Time {
	var m = time.Month((quarter-1)*3 + 1)
	return time.Date(year, m, 1, 0, 0, 0, 0, c.Location())
}

// EndOfQuarter 获取当前日历时区指定季度的结束时间
func (c Calendar) EndOfQuarter(year int, quarter int) time.Time {
	var m = time.Month((quarter-1)*3 + 3)
	return time.Date(year, m+1, 0, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// BeginningOfYear 获取当前日历时区指定年份的开始时间
func (c Calendar) BeginningOfYear(year int) time.Time {
	return time.Date(year, time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfYear 获取当前日历时区指定年份的结束时间
func (c Calendar) EndOfYear(year int) time.Time {
	return time.Date(year, time.December+1, 0, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// BeginningOfDay 获取本地时区指定日期的开始时间
func BeginningOfDay(year int, month time.Month, day int) time.Time {
	return NewCalendar(time.Local).BeginningOfDay(year, month, day)
}

// EndOfDay 获取本地时区指定日期的结束时间
func EndOfDay(year int, month time.Month, day int) time.Time {
	return NewCalendar(time.Local).EndOfDay(year, month, day)
}

// BeginningOfWeek 获取本地时区指定日期所在周的开始时间
func BeginningOfWeek(year int, month time.Month, day int) time.Time {
	return NewCalendar(time.Local).BeginningOfWeek(year, month, day)
}

// EndOfWeek 获取本地时区指定日期所在周的结束时间
func EndOfWeek(year int, month time.Month, day int) time.Time {
	return NewCalendar(time.Local).EndOfWeek(year, month, day)
}

// BeginningOfMonth 获取本地时区指定月份的开始时间
func BeginningOfMonth(year int, month time.Month) time.Time {
	return NewCalendar(time.Local).BeginningOfMonth(year, month)
}

// EndOfMonth 获取本地时区指定月份的结束时间
func EndOfMonth(year int, month time.Month) time.Time {
	return NewCalendar(time.Local).EndOfMonth(year, month)
}

// BeginningOfQuarter 获取本地时区指定季度的开始时间
func BeginningOfQuarter(year int, quarter int) time.Time {
	return NewCalendar(time.Local).BeginningOfQuarter(year, quarter)
}

// EndOfQuarter 获取本地时区指定季度的结束时间
func EndOfQuarter(year int, quarter int) time.Time {
	return NewCalendar(time.Local).EndOfQuarter(year, quarter)
}

// BeginningOfYear 获取本地时区指定年份的开始时间
func BeginningOfYear(year int) time.Time {
	return NewCalendar(time.Local).BeginningOfYear(year)
}

// EndOfYear 获取本地时区指定年份的结束时间
func EndOfYear(year int) time.Time {
	return NewCalendar(time.Local).EndOfYear(year)
}
