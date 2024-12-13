package day

// (c) Christian Maurer   v. 241205 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
)
type
  Period byte; const (
  Daily = Period(iota); Weekly; Monthly; Quarterly; HalfYearly; Yearly; Decadic; NPeriods
)
type
  Weekday byte; const (
  Monday = Weekday(iota); Tuesday; Wednesday; Thursday; Friday; Saturday; Sunday; nWeekdays
)
type
  DayOp = func (d Calendarday)
const
  NWeekdays = uint(nWeekdays)
const ( // Format
  Dd = iota  // e.g. "05"
  Dd_mm_     // e.g. "05.12."
  Dd_mm_yy   // e.g. "05.12.24" // default Format
  Dd_mm_yyyy // e.g. "05.12.2024"
  Yymmdd     // e.g. "241205"
  Yyyymmdd   // e.g. "20241205"
  Dd_M       // e.g. "05. Dezember"
  Dd_M_yyyy  // e.g. "05. Dezember 2024"
  D_M        // e.g. "5. Dezember"
  D_M_yyyy   // e.g. "5. Dezember 2024"
  Yy         // e.g. "24"
  Yyyy       // e.g. "2024"
  Wd         // e.g. "Do"
  WD         // e.g. "Donnerstag"
  Mmm        // e.g. "Dez"
  M          // e.g. "Dezember"
  Myyyy      // e.g. "Dezember 2024"
  Wn         // e.g. "49" (.Woche)
  WN         // e.g. "49. Woche"
  WNyyyy     // e.g. "49. Woche 2024"
  Qu         // e.g. "IV/24"
  NFormats
)
// Every day has - depending on being a holiday or not
// and on its format - the corresponding colours,
// where the suffix F/B means foreground/background.
var (
  WeekdayF, WeekdayB, HolidayF, HolidayB,
  YearnumberF, YearnumberB, WeekdayNameF, WeekdayNameB,
  MonthF, MonthB col.Colour
)
type
  Calendarday interface {

  Editor
  EditorGr
  col.Colourer
  Valuator
  Formatter
  Stringer
  Printer

// x is the first day in the range of the implementation.
  SetMin()

// x is the last day in the range of the implementation.
  SetMax()

// x is the day of the system date.
  Update()

// Returns true, iff (d, m, y) defines a day within the range of the
// implementation. In this case, x is that day. Otherwise, x is empty.
  Set (d, m, y uint) bool

// Returns true, iff x is empty and y is not empty or
// if x and y are not empty and x within one year is before y;
// in case day and month of x and y coincide,
// iff the year of x is before that of y.
  LessInYear (y Calendarday) bool

// Returns false, if x is empty.
// Returns otherwise true, iff x is in the same period p as y.
  Equiv (y Calendarday, p Period) bool

// Returns false, if x is empty.
// Returns otherwise true, iff x is the first day within the period p.
  IsBeginning (b Period) bool

// If x is empty, nothing has happened. Otherwise
// x is now the first day within the period p of x before.
  SetBeginning (p Period)

// If x is empty, nothing has happened. Otherwise
// x is now the last day within the period p of x before.
  SetEnd (p Period)

// Returns true, iff x is the day of the system date.
  Actual() bool

// Returns true, iff x is before the actual day.
  Elapsed() bool

// Returns max(uint), if x or y is empty. Returns otherwise
// the absolute value of the number of days between x and y.
  Distance (y Calendarday) uint

// Returns 0, if x is empty.
// Returns otherwise the number of days in the year of x.
  NumberOfDays() uint

// Returns 0, if x is empty. Returns otherwise,
// the ordinal number of the day in the year of x (starting with 1 on the 1st January)
  OrdDay() uint

// If x is empty or the effect would lead outside the range
// of the implemenation, nothing has happened. Otherwise
// x is increased by the number of days of p.
  Inc (p Period)

// If x is empty or the effect would lead outside the range
// of the implemenation, nothing has happened.
// Otherwise x is increased by d days.
  Inc1 (d uint)

// If x is empty or the effect would lead outside the range
// of the implemenation, nothing has happened. Otherwise
// x is decreased by the number of days of p.
  Dec (p Period)

// If x is empty or the effect would lead outside the range
// of the implemenation, nothing has happened.
// Otherwise x is decreased by d days.
  Dec1 (d uint)

// TODO Spec
  Change (c kbd.Comm, d uint)

// Returns the weekday of the first day within the period p of x.
  Weekday (p Period) Weekday

// Returns false, if x is empty.
// Returns otherwise true, iff x is a holiday by law in Germany.
  IsHoliday() bool

// If x is empty, nothing has happened.
// Otherwise, x is now the easter sunday in the year of x.
  SetEaster()

// If x is empty, nothing has happened.
// Otherwise, x is now the carnival tuesday in the year of x.
  SetCarnival()

// Spec is top secret.
  SetCasetta()

// Returns false, if x is empty.
// Returns otherwise true, iff x is neither saturday nore sunday.
  IsWorkday() bool

// Returns 0, if x is empty.
// Returns otherwise the number of weekdays in the year of x.
  NWorkdays (y Calendarday) uint

// Returns false, if x is empty. Returns otherwise true,
// iff last Sunday in October <= x < last Sunday in March.
  Normal() bool

// Returns false, if x is empty. Returns otherwise true,
// iff x.Normal && x == last Sunday in October
// or ! x.Normal && x == last Sunday in March.
  Normal1() bool

// Returns the empty Calendarday, if x is empty; returns otherwise
// the last Sunday for n in October; for !n in March in the year of x.
  LastSunday (n bool) Calendarday

// TODO Spec
  SetAttribute (p DayOp)

// Returns 0, if x is empty. Returns otherwise the number
// of the week of x in the year of x due to DIN 1355/ISO 8601.
  Weeknumber() uint

// Returns the day, month resp. year of x.
  Day() uint
  Month() uint
  Year() uint

// TODO Spec
  PosInWeek (v bool, a uint) (uint, uint)

// TODO Spec
  WriteWeek (v bool, a, l, c uint)

// TODO Spec
  PosInMonth (v bool, n, l, c uint) (uint, uint)

// TODO Spec
  WriteMonth (v bool, n, z, s, l, c uint)

// TODO Spec
  PrintMonth (v bool, n, z, s, l, c uint)

// TODO Spec
  PosInYear() (uint, uint)

// TODO Spec
  WriteYear (l, c uint)

// TODO Spec
  PrintYear (l, c uint)

// TODO Spec
  Selected (o CondOp) bool

// TODO Spec
  Randomize()
}

func New() Calendarday { return new_() }

// Pre: d, m and y have "reasonable" values.
func New3 (d, m, y uint) Calendarday { return new3 (d, m, y) }
