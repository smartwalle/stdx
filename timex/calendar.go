package timex

import "time"

type Calendar struct {
	loc       *time.Location
	weekStart time.Weekday
}

func NewCalendar(loc *time.Location) Calendar {
	if loc == nil {
		loc = time.Local
	}
	return Calendar{loc: loc, weekStart: time.Sunday}
}

func (c Calendar) Location() *time.Location {
	if c.loc == nil {
		return time.Local
	}
	return c.loc
}

func (c Calendar) WeekStart() time.Weekday {
	if c.weekStart < time.Sunday || c.weekStart > time.Saturday {
		return time.Sunday
	}
	return c.weekStart
}

func (c Calendar) UseWeekStart(weekStart time.Weekday) Calendar {
	if weekStart < time.Sunday || weekStart > time.Saturday {
		weekStart = time.Sunday
	}
	c.weekStart = weekStart
	return c
}

// BeginningOfMinuteAt 获取指定时间在当前日历时区所在分钟的开始时间
func (c Calendar) BeginningOfMinuteAt(t time.Time) time.Time {
	t = t.In(c.Location())
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()
	return time.Date(year, month, day, hour, minute, 0, 0, c.Location())
}

// EndOfMinuteAt 获取指定时间在当前日历时区所在分钟的结束时间
func (c Calendar) EndOfMinuteAt(t time.Time) time.Time {
	t = t.In(c.Location())
	year, month, day := t.Date()
	hour, minute, _ := t.Clock()
	return time.Date(year, month, day, hour, minute, 59, int(time.Second-time.Nanosecond), c.Location())
}

// MinuteRangeAt 获取指定时间在当前日历时区所在分钟的开始时间和结束时间
func (c Calendar) MinuteRangeAt(t time.Time) (start, end time.Time) {
	return c.BeginningOfMinuteAt(t), c.EndOfMinuteAt(t)
}

// BeginningOfHourAt 获取指定时间在当前日历时区所在小时的开始时间
func (c Calendar) BeginningOfHourAt(t time.Time) time.Time {
	t = t.In(c.Location())
	year, month, day := t.Date()
	hour, _, _ := t.Clock()
	return time.Date(year, month, day, hour, 0, 0, 0, c.Location())
}

// EndOfHourAt 获取指定时间在当前日历时区所在小时的结束时间
func (c Calendar) EndOfHourAt(t time.Time) time.Time {
	t = t.In(c.Location())
	year, month, day := t.Date()
	hour, _, _ := t.Clock()
	return time.Date(year, month, day, hour, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// HourRangeAt 获取指定时间在当前日历时区所在小时的开始时间和结束时间
func (c Calendar) HourRangeAt(t time.Time) (start, end time.Time) {
	return c.BeginningOfHourAt(t), c.EndOfHourAt(t)
}

// BeginningOfDay 获取当前日历时区指定日期的开始时间
func (c Calendar) BeginningOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, c.Location())
}

// EndOfDay 获取当前日历时区指定日期的结束时间
func (c Calendar) EndOfDay(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// DayRange 获取当前日历时区指定日期的开始时间和结束时间
func (c Calendar) DayRange(year int, month time.Month, day int) (start, end time.Time) {
	return c.BeginningOfDay(year, month, day), c.EndOfDay(year, month, day)
}

// BeginningOfDayAt 获取指定时间在当前日历时区所在日期的开始时间
func (c Calendar) BeginningOfDayAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.BeginningOfDay(t.Date())
}

// EndOfDayAt 获取指定时间在当前日历时区所在日期的结束时间
func (c Calendar) EndOfDayAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.EndOfDay(t.Date())
}

// DayRangeAt 获取指定时间在当前日历时区所在日期的开始时间和结束时间
func (c Calendar) DayRangeAt(t time.Time) (start, end time.Time) {
	return c.BeginningOfDayAt(t), c.EndOfDayAt(t)
}

// BeginningOfWeek 获取当前日历时区指定日期所在周的开始时间
func (c Calendar) BeginningOfWeek(year int, month time.Month, day int) time.Time {
	var t = time.Date(year, month, day, 0, 0, 0, 0, c.Location())
	var w = t.Weekday()
	var d = int((w - c.WeekStart() + 7) % 7)
	return time.Date(year, month, day-d, 0, 0, 0, 0, c.Location())
}

// EndOfWeek 获取当前日历时区指定日期所在周的结束时间
func (c Calendar) EndOfWeek(year int, month time.Month, day int) time.Time {
	var t = time.Date(year, month, day, 0, 0, 0, 0, c.Location())
	var w = t.Weekday()
	var d = int((c.WeekStart() + 6 - w + 7) % 7)
	return time.Date(year, month, day+d, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// WeekRange 获取当前日历时区指定日期所在周的开始时间和结束时间
func (c Calendar) WeekRange(year int, month time.Month, day int) (start, end time.Time) {
	return c.BeginningOfWeek(year, month, day), c.EndOfWeek(year, month, day)
}

// BeginningOfWeekAt 获取指定时间在当前日历时区所在周的开始时间
func (c Calendar) BeginningOfWeekAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.BeginningOfWeek(t.Date())
}

// EndOfWeekAt 获取指定时间在当前日历时区所在周的结束时间
func (c Calendar) EndOfWeekAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.EndOfWeek(t.Date())
}

// WeekRangeAt 获取指定时间在当前日历时区所在周的开始时间和结束时间
func (c Calendar) WeekRangeAt(t time.Time) (start, end time.Time) {
	return c.BeginningOfWeekAt(t), c.EndOfWeekAt(t)
}

// BeginningOfMonth 获取当前日历时区指定月份的开始时间
func (c Calendar) BeginningOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month, 1, 0, 0, 0, 0, c.Location())
}

// EndOfMonth 获取当前日历时区指定月份的结束时间
func (c Calendar) EndOfMonth(year int, month time.Month) time.Time {
	return time.Date(year, month+1, 0, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// MonthRange 获取当前日历时区指定月份的开始时间和结束时间
func (c Calendar) MonthRange(year int, month time.Month) (start, end time.Time) {
	return c.BeginningOfMonth(year, month), c.EndOfMonth(year, month)
}

// BeginningOfMonthAt 获取指定时间在当前日历时区所在月份的开始时间
func (c Calendar) BeginningOfMonthAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.BeginningOfMonth(t.Year(), t.Month())
}

// EndOfMonthAt 获取指定时间在当前日历时区所在月份的结束时间
func (c Calendar) EndOfMonthAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.EndOfMonth(t.Year(), t.Month())
}

// MonthRangeAt 获取指定时间在当前日历时区所在月份的开始时间和结束时间
func (c Calendar) MonthRangeAt(t time.Time) (start, end time.Time) {
	return c.BeginningOfMonthAt(t), c.EndOfMonthAt(t)
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

// QuarterRange 获取当前日历时区指定季度的开始时间和结束时间
func (c Calendar) QuarterRange(year int, quarter int) (start, end time.Time) {
	return c.BeginningOfQuarter(year, quarter), c.EndOfQuarter(year, quarter)
}

// BeginningOfQuarterAt 获取指定时间在当前日历时区所在季度的开始时间
func (c Calendar) BeginningOfQuarterAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.BeginningOfQuarter(t.Year(), Quarter(t.Month()))
}

// EndOfQuarterAt 获取指定时间在当前日历时区所在季度的结束时间
func (c Calendar) EndOfQuarterAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.EndOfQuarter(t.Year(), Quarter(t.Month()))
}

// QuarterRangeAt 获取指定时间在当前日历时区所在季度的开始时间和结束时间
func (c Calendar) QuarterRangeAt(t time.Time) (start, end time.Time) {
	return c.BeginningOfQuarterAt(t), c.EndOfQuarterAt(t)
}

// BeginningOfYear 获取当前日历时区指定年份的开始时间
func (c Calendar) BeginningOfYear(year int) time.Time {
	return time.Date(year, time.January, 1, 0, 0, 0, 0, c.Location())
}

// EndOfYear 获取当前日历时区指定年份的结束时间
func (c Calendar) EndOfYear(year int) time.Time {
	return time.Date(year, time.December+1, 0, 23, 59, 59, int(time.Second-time.Nanosecond), c.Location())
}

// YearRange 获取当前日历时区指定年份的开始时间和结束时间
func (c Calendar) YearRange(year int) (start, end time.Time) {
	return c.BeginningOfYear(year), c.EndOfYear(year)
}

// BeginningOfYearAt 获取指定时间在当前日历时区所在年份的开始时间
func (c Calendar) BeginningOfYearAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.BeginningOfYear(t.Year())
}

// EndOfYearAt 获取指定时间在当前日历时区所在年份的结束时间
func (c Calendar) EndOfYearAt(t time.Time) time.Time {
	t = t.In(c.Location())
	return c.EndOfYear(t.Year())
}

// YearRangeAt 获取指定时间在当前日历时区所在年份的开始时间和结束时间
func (c Calendar) YearRangeAt(t time.Time) (start, end time.Time) {
	return c.BeginningOfYearAt(t), c.EndOfYearAt(t)
}

// Quarter 获取指定月份所在季度
func Quarter(month time.Month) int {
	return int(month-1)/3 + 1
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
