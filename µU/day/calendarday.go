package day

// (c) Christian Maurer   v. 250508 - license see µU.go

import (
  . "µU/ker"
  "µU/time"
  . "µU/obj"
  "µU/rand"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/font"
  "µU/pbox"
  "µU/N"
)
const (
  emptyYear = uint(1879)
  startYear = emptyYear + 1
  limitYear = uint(2030) // Yy: 0..20 -> 2000..2020; 21..99 -> 1921..1999
  endYear   = uint(2058) // emptyYear + 179, // got to change that, if I am 113 years old
                         // 179 Jahre < MAX (uint16) Tage < 180 Jahre
  wdDay     = 10
  wdMonth   =  9
  maxMonth  = uint(12)
  maxDay    = uint(31)
  maxCode   = uint16(65379) // 31.12.2058
)
type
  calendarday struct {
                 day,
               month,
                year uint
                     Format
              cF, cB col.Colour
                     font.Font
                     }
var (
  today = New()
  currentCentury uint
  todayCode uint16
  nameMonth = [maxMonth+1]string {"         ",
                                  "Januar   ", "Februar  ", "März     ",
                                  "April    ", "Mai      ", "Juni     ",
                                  "Juli     ", "August   ", "September",
                                  "Oktober  ", "November ", "Dezember "}
/*/
                                  "January  ", "February ", "March    ",
                                  "April    ", "Mai      ", "June     ",
                                  "July     ", "August   ", "September",
                                  "October  ", "November ", "December "}
/*/
  WdText = [NWeekdays]string {"Montag    ", "Dienstag  ", "Mittwoch  ",
                              "Donnerstag", "Freitag   ", "Sonnabend ", "Sonntag   "}
  WdShorttext = [NWeekdays]string {"Mo", "Di", "Mi", "Do", "Fr", "Sa", "So"}
  wd = []uint {
    Dd:         2,
    Dd_mm_:     2 + 1 + 2 + 1,
    Dd_mm_yy:   2 + 1 + 2 + 1 + 2,
    Dd_mm_yyyy: 2 + 1 + 2 + 1 + 4,
    Yymmdd:     2 + 2 + 2,
    Yyyymmdd:   4 + 2 + 2,
    Dd_M:       3 + wdMonth,
    D_M:        3 + wdMonth,
    Dd_M_yyyy:  2 + 1 + 1 + wdMonth + 1 + 4,
    D_M_yyyy:   2 + 1 + 1 + wdMonth + 1 + 4,
    Yy:         2,
    Yyyy:       4,
    Wd:         2,
    WD:         wdDay,
    Mmm:        3,
    M:          wdMonth,
    Myyyy:      wdMonth + 1 + 4,
    Wn:         2,
    WN:         2 + 1 + 1 + 5,
    WNyyyy:     2 + 1 + 1 + 5 + 1 + 4,
    Qu:         6,
  }
  bx = box.New()
  pbx = pbox.New()
  actualDay, actualMonth, actualYear = maxDay, maxMonth, emptyYear
  Codeyear = emptyYear
  yearcode = uint16(0)
  actualHolidayYear = emptyYear
  holiday [maxDay+1][maxMonth+1]bool
  op = attribute
//  carnival *calendarday
)

func init() {
  WeekdayF, WeekdayB = col.StartCols()
  HolidayF, HolidayB = col.StartColsA()
  YearnumberF, YearnumberB = col.FlashWhite(), col.Magenta()
  WeekdayNameF, WeekdayNameB = col.Magenta(), WeekdayB
  MonthF, MonthB = YearnumberF, YearnumberB
  today.Update()
  currentCentury = 100 * (today.(*calendarday).year / 100)
  todayCode = today.(*calendarday).internalCode()
//  carnival = new_()
}

func (x *calendarday) imp (Y any) *calendarday {
  y, ok := Y.(*calendarday)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Calendarday {
  x := new (calendarday)
  x.day, x.month, x.year = maxDay, maxMonth, emptyYear
  x.Format = Dd_mm_yy
  x.cF, x.cB = col.StartCols()
  x.Font = font.Roman
  return x
}

func new3 (d, m, y uint) Calendarday {
  x := new_().(*calendarday)
  x.Set (d, m, y)
  return x
}

func (x *calendarday) Randomize() {
  const n = uint16(14976) // Code of 1.1.1921
  x.Decode (Encode (n + uint16(rand.Natural (uint(todayCode + uint16(1) - n)))))
}

func (x *calendarday) Empty() bool {
  return x.year == emptyYear
}

func (x *calendarday) Clr() {
  x.day, x.month, x.year = maxDay, maxMonth, emptyYear
}

func (x *calendarday) SetMin() {
  x.day, x.month, x.year = 1, 1, startYear
}

func (x *calendarday) SetMax() {
  x.day, x.month, x.year = maxDay, maxMonth, endYear
}

func (x *calendarday) Update() {
  x.day, x.month, x.year = time.ActDate()
}

func (x *calendarday) Set (d, m, y uint) bool {
  if x.defined (d, m, y) {
    return true
  }
  x.Clr()
  return false
}

func (x *calendarday) Copy (Y any) {
  y := x.imp (Y)
  x.day, x.month, x.year = y.day, y.month, y.year
  x.Format = y.Format
  x.cF, x.cB = y.cF, y.cB
}

func (x *calendarday) Clone() any {
  y := new_().(*calendarday)
  y.Copy (x)
  return y
}

func (x *calendarday) Eq (Y any) bool {
  y := x.imp (Y)
  if x.year == emptyYear { return y.year == emptyYear }
  if y.year == emptyYear { return false }
  return x.day == y.day && x.month == y.month && x.year == y.year
}

func (x *calendarday) Less (Y any) bool {
  y := x.imp (Y)
  if x.year == emptyYear {
    return y.year != emptyYear
  }
  if y.year == emptyYear { return false }
  if x.year == y.year {
    if x.month == y.month {
      return x.day < y.day
    } else {
      return x.month < y.month
    }
  }
  return x.year < y.year
}

func (x *calendarday) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *calendarday) LessInYear (Y Calendarday) bool {
  y := x.imp (Y)
  if x.year == emptyYear {
    return x.year != emptyYear
  }
  if x.year == emptyYear { return false }
  if x.month == y.month {
    if x.day == y.day {
      return x.year < y.year
    } else {
      return x.day < y.day
    }
  }
  return x.month < y.month
}

func (x *calendarday) Equiv (Y Calendarday, p Period) bool {
  y := x.imp (Y)
  if x.year == emptyYear {
    return y.year == emptyYear
  }
  if y.year == emptyYear { return false }
  switch p {
  case Daily:
    return x.Eq (y)
  case Weekly:
    c := x.internalCode()
    w := (c + 2) % 7
    c1 := y.internalCode()
    w1 := (c1 + 2) % 7
    if c <= c1 {
      if c1 - c < 7 {
        return w <= w1
      }
    } else if c - c1 < 7 {
      return w > w1
    }
    break
  case Monthly:
    return x.month == y.month && x.year == y.year
  case Quarterly:
    return ((x.month - 1) / 3 == (y.month - 1) / 3) && x.year == y.year
  case HalfYearly:
    return ((x.month - 1) / 6 == (y.month - 1) / 6) && x.year == y.year
  case Yearly:
    return x.year == y.year
  case Decadic:
    return x.year % 10 == y.year % 10
  }
  return false
}

func (x *calendarday) IsBeginning (p Period) bool {
  switch p {
  case Daily:
    break
  case Weekly:
    return x.Weekday (Daily) == Monday
  case Monthly:
    return x.day == 1
  case Quarterly:
    return x.day == 1 && 3 * (x.month - 1) / 3 + 1 == x.month
  case HalfYearly:
    return x.day == 1 && 6 * (x.month - 1) / 6 + 1 == x.month
  case Yearly:
    return uint8 (x.day) * uint8 (x.month) == 1
  case Decadic:
    return uint8 (x.day) * uint8 (x.month) == 1 && x.year % 10 == 0
  }
  return true
}

func (x *calendarday) SetBeginning (p Period) {
  if x.year == emptyYear { return }
  switch p {
  case Daily:
    return
  case Weekly:
    for w := x.Weekday (Daily); w > Monday; w-- {
      x.Dec (Daily)
    }
  case Monthly:
    x.day = 1
  case Quarterly:
    x.day, x.month = 1, 3 * (x.month - 1) / 3 + 1
  case HalfYearly:
    x.day, x.month = 1, 6 * (x.month - 1) / 6 + 1
  case Yearly:
    x.day, x.month = 1, 1
  case Decadic:
    if 10 * x.year / 10 > emptyYear {
      x.day, x.month, x.year = 1, 1, 10 * x.year / 10
    }
  }
}

func (x *calendarday) SetEnd (p Period) {
  if x.year == emptyYear { return }
  switch p {
  case Daily:
    return
  case Weekly:
    for w := x.Weekday (Daily); w < Sunday; w++ {
      x.Inc (Daily)
    }
  case Monthly:
    x.day = x.daysInMonth()
  case Quarterly:
    x.month = 3 * (((x.month - 1) / 3) + 1)
    x.day = x.daysInMonth()
  case HalfYearly:
    x.month = 6 * (((x.month - 1) / 6) + 1)
    x.day = x.daysInMonth()
  case Yearly:
    x.day = maxDay
    x.month = maxMonth
  case Decadic:
    if x.year + 9 <= endYear + x.year % 10 {
      x.day = maxDay
      x.month = maxMonth
      x.year += 9 - x.year % 10
    }
  }
}

func isLeapYear (y uint) bool {
/*
  if y % 400 == 0 {
    return true
  else if y % 100 != 0 {
    for the range of this implementation we only need:
*/
  if y == 1900 {
    return false
  }
  return (y % 4) == 0 // emptyYear: false
}

func (x *calendarday) daysInMonth() uint {
  if x.year == actualYear {
    if x.month == actualMonth {
      return actualDay
    } else {
      actualMonth = x.month
    }
  } else {
    actualMonth, actualYear = x.month, x.year
  }
  if x.month == 2 {
    actualDay = 28
    if isLeapYear (x.year) { actualDay ++ }
  } else if x.month / 8 == x.month % 2 { // Fingerknöchelprinzip!
    actualDay = 30 // maxDay - 1
  } else {
    actualDay = maxDay
  }
  return actualDay
}

func daysInYear (y uint) uint16 {
  d := uint16(365)
  if isLeapYear (y) {
    d ++
  }
  return d
}

func (x *calendarday) internalCode() uint16 {
  if x.year == emptyYear { return 0 }
  code := uint16(x.day)
  m := x.month
  for x.month > 1 {
    x.month--
    code += uint16(x.daysInMonth())
  }
  x.month = m
  if x.year != Codeyear {
    Codeyear = x.year
    yearcode = 0
    y := x.year
    for x.year > startYear {
      x.year--
      yearcode += daysInYear (x.year)
    }
    x.year = y
  }
  code += yearcode
  return code
}

func (x *calendarday) Actual() bool {
  return x.internalCode() == todayCode
}

func (x *calendarday) Elapsed() bool {
  if x.year == emptyYear {
    return false
  }
  return x.internalCode() < todayCode
}

func (x *calendarday) Distance (Y Calendarday) uint {
  y := x.imp (Y)
  if x.year == emptyYear || y.year == emptyYear { return MaxNat }
  c, c1 := x.internalCode(), y.internalCode()
  if c < c1 {
    return uint(c1 - c)
  }
  return uint(c - c1)
}

func (x *calendarday) NumberOfDays() uint {
  if x.year == emptyYear { return 0 }
  return uint(daysInYear (x.year))
}

func (x *calendarday) OrdDay() uint {
  n := uint(0)
  if x.year != emptyYear {
    y := new (calendarday)
    y.day, y.month, y.year = 1, 1, x.year
    for y.month < x.month {
      n += uint(y.daysInMonth())
      y.month++
    }
    n += uint(x.day)
  }
  return n
}

func (x *calendarday) Inc (p Period) {
  if x.year == emptyYear { return }
  t := x.daysInMonth()
  d, m, y := x.day, x.month, x.year
  switch p {
  case Daily:
    if d < t {
      d ++
    } else {
      d = 1
      if m < maxMonth {
        m ++
      } else if y < endYear {
        y ++
        m = 1
      } else {
        return
      }
    }
  case Weekly:
    if d + 7 <= t {
      d += 7
    } else if m < maxMonth {
      m ++
      d -= t - 7
    } else if y < endYear {
      y ++
      m = 1
      d -= 24
    } else {
      return
    }
  case Monthly:
    if m < maxMonth {
      m++
    } else if y < endYear {
      m = 1
      y ++
    } else {
      return
    }
  case Quarterly:
    if m < 10 {
      m += 3
    } else if y < endYear {
      m -= 9
      y ++
    } else {
      return
    }
  case HalfYearly:
    if m < 7 {
      m += 6
    } else if y < endYear {
      m -= 6
      y ++
    } else {
      return
    }
  case Yearly:
    if y < endYear {
      y ++
    } else {
      return
    }
  case Decadic:
    if y <= endYear - 10 {
      y += 10
    } else {
      return
    }
  }
  x.day, x.month, x.year = d, m, y
  t = x.daysInMonth()
  if x.day > t { x.day = t }
}

func (x *calendarday) Inc1 (n uint) {
  if x.year == emptyYear { return }
  d, m, y := x.day, x.month, x.year
  for n > 0 {
    x.Inc (Daily)
    if x.Empty() {
      x.day, x.month, x.year = d, m, y
      return
    }
    n--
  }
}

func (x *calendarday) Dec (p Period) {
  if x.year == emptyYear { return }
  t := x.daysInMonth()
  d, m, y := x.day, x.month, x.year
  switch p {
  case Daily:
    if d > 1 {
      d--
    } else if m > 1 {
      m--
      x.month--
      d = x.daysInMonth()
      x.month ++
    } else {
      y--
      m = maxMonth
      d = maxDay
      if y == emptyYear {
        return
      }
    }
  case Weekly:
    if d > 7 {
      d -= 7
    } else if m > 1 {
      m--
      x.month--
      t = x.daysInMonth()
      x.month ++
      d += t - 7
    } else {
      y--
      m = maxMonth
      d += 24
      if y == emptyYear { return }
    }
  case Monthly:
    if m > 1 {
      m--
    } else {
      y--
      m = maxMonth
      if y == emptyYear { return }
    }
  case Quarterly:
    if m > 3 {
      m -= 3
    } else {
      y--
      m += 9
      if y == emptyYear { return }
    }
  case HalfYearly:
    if m > 6 {
      m -= 6
    } else {
      y--
      m += 6
      if y == emptyYear { return }
    }
  case Yearly:
    y--
    if y == emptyYear { return }
  case Decadic:
    if y > emptyYear + 10 {
      y -= 10
    } else {
      return
    }
  }
  x.day, x.month, x.year = d, m, y
  t = x.daysInMonth()
  if x.day > t { x.day = t }
}

func (x *calendarday) Dec1 (n uint) {
  if x.year == emptyYear { return }
  d, m, y := x.day, x.month, x.year
  for n > 0 {
    x.Dec (Daily)
    if x.Empty() {
      x.day, x.month, x.year = d, m, y
      return
    }
    n--
  }
}

func (x *calendarday) Change (c kbd.Comm, d uint) {
  if x.year == emptyYear {
    return
  }
  switch c {
  case kbd.Enter, kbd.Esc:
    return
  case kbd.Right, kbd.Down:
    p := Monthly
    switch d {
    case 0:
      p = Daily
    case 1:
      p = Weekly
    }
    x.Inc (p)
  case kbd.Left, kbd.Up:
    p := Monthly
    switch d {
    case 0:
      p = Daily
    case 1:
      p = Weekly
    }
    x.Dec (p)
  case kbd.PgRight, kbd.PgDown:
    p := Yearly
    if d == 0 {
      p = Monthly
    }
    x.Inc (p)
  case kbd.PgLeft, kbd.PgUp:
    p := Yearly
    if d == 0 {
      p = Monthly
    }
    x.Dec (p)
  case kbd.Pos1, kbd.End:
    p := Yearly
    switch d {
    case 1:
      p = Weekly
    case 2:
      p = Monthly
    }
    if c == kbd.Pos1 {
      x.SetBeginning (p)
    } else {
      x.SetEnd (p)
    }
  }
}

func ggg (m uint) uint {
  if m <= 2 { return 1 }
  return 0
}

func (x *calendarday) wochentag() uint { // I have no idea about the sources, but this algorithm seems to be ok
  d, m, y := x.day, x.month, x.year
  y -= ggg(m)
  return (d + y + y / 4 - y / 100 + y / 400 + (31 * (m + 12 * ggg(m) - 2)) / 12 + 6) % 7
}

func (x *calendarday) weekday() Weekday {
  w := Weekday((x.internalCode() + uint16(Wednesday)) % 7)
  if uint(w) != x.wochentag() { panic("bug in day/wochentag") }
  return w
 // The day with code 0, 31.12.1879, was a Wednesday
}

func (x *calendarday) Weekday (p Period) Weekday {
  d := x.Clone().(*calendarday)
  switch p {
  case Daily:
    ;
  case Weekly:
    return Monday
  case Monthly:
    d.day = 1
  case Quarterly:
    d.day, d.month = 1, 3 * (d.month - 1) / 3 + 1
  case HalfYearly:
    d.day, d.month = 1, 6 * (d.month - 1) / 6 + 1
  case Yearly:
    d.day, d.month = 1, 1
  case Decadic:
    if 10 * d.year / 10 > emptyYear {
      d.day, d.month, d.year = 1, 1, 10 * d.year / 10
    }
  }
  return Weekday(d.weekday())
}

func computeHolidays() { // Quelle: S. Deschauer, Die Osterfestberechnung. DdM 14 (1986), 68-84
  x := new_().(*calendarday)
  x.day, x.month, x.year = 1, 1, actualHolidayYear
  Wochentag := x.weekday()
  for m := uint(1); m <= maxMonth; m++ {
    x.month = uint(m)
    t1 := x.daysInMonth()
    for t := 1; uint(t) <= t1; t++ {
      holiday [t][m] = Wochentag == Sunday
      if Wochentag == Sunday {
        Wochentag = Monday
      } else {
        Wochentag++
      }
    }
  }
  holiday [1][1] = true // Neujahr
  if actualHolidayYear >= 1890 { // Tag der Arbeit
    holiday [1][5] = true
  }
  if actualHolidayYear == 2025 { // Gedenktag an das Ende des 2. Weltkriegs
    holiday [8][5] = true
  }
  if actualHolidayYear > 1953 { // Tag der deutschen Einheit
    if actualHolidayYear < 1990 { // 17.6.1990 ein Sonntag
      holiday [17][6] = true
    } else {
      holiday [3][10] = true
    }
  }
  holiday [25][12] = true // Weihnachten
  holiday [26][12] = true
/* >>> 1583..6199:
  s = J / 100 - J / 400 - 2
   für 1900..2099: 13, 1800..1899: 12
  m = (J - 100 * (J / 4200)) / 300 - 2
    für 1800..2000: 4
  M = (15 + s - m) % 30
    für 1900, 2000: 24, 1800..1899: 23
  N = (6 + s) % 7
    für 1900..2099: 5, 1800..1899: 4
>>> für 1800..2099 reicht also: */
  mm := uint(24)
  nn := uint(5)
  if actualHolidayYear < 1900 {
    mm--
    nn--
  }
  d := (mm + 19 * (actualHolidayYear % 19)) % 30
  e := (nn + 2 * (actualHolidayYear % 4) + 4 * (actualHolidayYear % 7) + 6 * d) % 7
  t := 22 + d + e
  if e == 6 { // Sonntag
    if d == 29 || actualHolidayYear % 19 >= 11 && d == 28 {
      t -= 7
    }
  }
  // 22 <= t <= 56, t. März ist Ostersonntag
  if t <= 30 { // Ostermontag
    holiday [t + 1][3] = true
  } else {
    holiday [t - 30][4] = true
  }
/*
  var f, ft uint
  carnival.year = actualHolidayYear
  // carnival = 7 Wochen vor Osterdienstag = 48 Tage vor Ostermontag
  // 2. Februar <= carnival <= 8. Mürz
  ft = t
  carnival.month = 3, // März
  if ft <= 48 {
    ft += 28
    if isLeapYear (actualHolidayYear) {
      ft++
    }
    carnival.month-- // Februar
  }
  ft -= 48
  carnival.day = ft
*/
  if t <= 33 { // Karfreitag
    holiday [t - 2][3] = true
  } else {
    holiday [t - 33][4] = true
  }
  t -= 11 // Ostersonntag + 50 Tage - April - Mai: 11 <= t <= 46
  if t <= maxDay { // Pfingstmontag
    holiday [t][5] = true
  } else {
    holiday [t - maxDay][6] = true
  }
  t -= 11 // Pfingstmontag - 11 Tage: 0 <= t <= 35
  if t == 0 { // Himmelfahrt
    holiday [30][4] = true
  } else if t <= maxDay {
    holiday [t][5] = true
  } else {
    holiday [t - maxDay][6] = true
  }
  if actualHolidayYear <= 1994 { // Bußtag
    x.day = 20 // wenn das ein Mo ist, ist der 22. Bußtag
    x.month = 11
    x.year = actualHolidayYear
    x.day -= uint(x.weekday())
    x.day += 2 // Spanne von Montag bis Mittwoch, s.o.
    holiday [x.day][11] = true
  }
  if actualHolidayYear == 2017 { // 500. Reformationstag
    holiday [31][10] = true
  }
  if actualHolidayYear >= 2019 { // Weltfrauentag (Berlin)
    holiday [8][3] = true
  }
  if actualHolidayYear == 2020 { // einmaliger Feiertag in Berlin
    holiday [8][5] = true
  }
}

func (x *calendarday) IsHoliday() bool {
  if x.year == emptyYear { return false }
  if x.year != actualHolidayYear {
    actualHolidayYear = x.year
    computeHolidays()
  }
  return holiday [x.day][x.month]
}

func (x *calendarday) SetEaster() {
  if x.year == emptyYear { return }
  if x.year != actualHolidayYear {
    actualHolidayYear = x.year
    computeHolidays()
  }
  x.day = 24
  x.month = 3 // earliest possible Eastermonday
  w := x.Weekday (Daily)
  if w != Monday {
    x.day += 7 - uint(w)
  } // the first monday after
  for ! holiday [x.day][x.month] {
    x.day += 7
    if x.day > maxDay {
      x.day -= maxDay
      x.month = 4
    }
  } // Eastermonday
  if x.day > 1 {
    x.day--
  } else {
    x.day = maxDay
    x.month = 3
  }
}

func (x *calendarday) SetCarnival() {
  if x.year == emptyYear { return }
  x.SetEaster()
  for i := 0; i < 47; i++ {
    x.Dec (Daily)
  }
}

func (x *calendarday) SetCasetta() {
  x.SetEaster()
  for i := 0; i < 9; i++ {
    x.Inc (Weekly)
  }
  x.Dec (Daily)
}

func (x *calendarday) LastSunday (a bool) Calendarday {
  y := new_().(*calendarday)
  if x.year == emptyYear {
    return y
  }
  y.year = x.year
  if a { // October
    y.month = 10
  } else {
    y.month = 3
  }
  y.day = maxDay
  wd := y.weekday()
  if wd != Sunday {
    y.day -= 1 + uint(wd)
  }
  return y
}

func (x *calendarday) Normal() bool {
  if x.year == emptyYear { return false }
  oct, mar := x.LastSunday (true), x.LastSunday (false)
  return x.Less (mar) || oct.Eq (x) || oct.Less (x)
}

func (x *calendarday) Normal1() bool {
  if x.year == emptyYear { return false }
  oct, mar := x.LastSunday (true), x.LastSunday (false)
  if x.Normal() {
    return x.Eq (oct)
  }
  return x.Eq (mar)
}

func (x *calendarday) IsWorkday() bool {
  switch x.Weekday (Daily) {
  case Saturday, Sunday:
    return false
  }
  return ! x.IsHoliday()
}

func (x *calendarday) NWorkdays (Y Calendarday) uint {
  y := x.imp (Y)
  if x.Empty() || y.Empty() { return 0 } // MaxNat ?
  a := uint(0)
  if x.Less (y) {
    z := x.Clone().(*calendarday)
    for {
      if z.IsWorkday() {
        a++
      }
      z.Inc (Daily)
      if y.Less (z) { break }
    }
  } else {
    z := y.Clone().(*calendarday)
    for {
      if z.IsWorkday() {
        a++
      }
      z.Inc (Daily)
      if x.Less (z) { break }
    }
  }
  return a
}

func (x *calendarday) GetFormat() Format {
  return x.Format
}

func (x *calendarday) SetFormat (f Format) {
  if f < NFormats {
    x.Format = f
  }
}

func (x *calendarday) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *calendarday) Cols() (col.Colour, col.Colour) {
  return x.cF, x.cB
}

func (x *calendarday) Weeknumber() uint {
  if x.year == emptyYear { return 0 }
  const Stichtag = Thursday // DIN 8601 (1975), entspricht ISO-Entwurf
  y := x.Clone().(*calendarday)
  y.day, y.month = 1, 1
  n := uint(x.internalCode() - y.internalCode())
  wd := y.weekday()
  n += uint(wd)
  if wd <= Stichtag { n += 7 }
  return n / 7
}

func (x *calendarday) String() string {
  if x.year == emptyYear {
    return str.New (wd[x.Format])
  }
  if x.day == 0 { Panic ("day.String: x.day == 0") }
  s := ""
  switch x.Format {
  case Dd, Dd_mm_, Dd_mm_yy, Dd_mm_yyyy, Dd_M, Dd_M_yyyy:
    s = N.StringFmt (x.day, 2, true)
    if x.Format == Dd {
      return s
    }
    s += "."
    switch x.Format {
    case Dd_M, Dd_M_yyyy:
      s += " " + nameMonth[x.month]
      str.OffSpc (&s)
      if x.Format == Dd_M { return s }
      s += " "
    case Dd_mm_, Dd_mm_yy, Dd_mm_yyyy:
      s += N.StringFmt (x.month, 2, true) + "."
    }
    switch x.Format {
    case Dd_mm_:
      ;
    case Dd_mm_yy:
      s += N.StringFmt (x.year % 100, 2, true)
    case Dd_mm_yyyy, Dd_M_yyyy:
      s += N.StringFmt (x.year, 4, false)
    }
  case Yymmdd:
    s = N.StringFmt (x.year % 100, 2, true) +
        N.StringFmt (x.month, 2, true) +
        N.StringFmt (x.day, 2, true)
  case Yyyymmdd:
    s = N.StringFmt (x.year, 4, true) +
        N.StringFmt (x.month, 2, true) +
        N.StringFmt (x.day, 2, true)
  case Yy:
    s = N.StringFmt (x.year, 2, true)
  case Yyyy:
    s = N.StringFmt (x.year, 4, false)
  case Wd:
    s = WdShorttext [x.Weekday (Daily)]
  case WD:
    s = WdText [x.Weekday (Daily)]
  case Mmm, M:
    s = nameMonth [x.month]
    if x.Format == Mmm { s = str.Part (s, 0, 3) }
  case Myyyy:
    s = nameMonth [x.month] + " " +
        N.StringFmt (x.year, 4, false)
  case Wn, WN, WNyyyy:
    s = N.StringFmt (x.Weeknumber(), 2, false)
    if x.Format > Wn { s += ". Woche" }
    if x.Format == WNyyyy {
      s += " " + N.StringFmt (x.year, 4, false)
    }
  case Qu:
    switch (x.month - 1) / 3 {
    case 0:
      s = "  I"
    case 1:
      s = " II"
    case 2:
      s = "III"
    case 3:
      s = " IV"
    }
    s += "/" + N.StringFmt (x.year, 2, true)
  }
  return s
}

func (x *calendarday) Day() uint {
  if x.year == emptyYear { return 0 }
  return x.day
}

func (x *calendarday) Month() uint {
  if x.year == emptyYear { return 0 }
  return x.month
}

func (x *calendarday) Year() uint {
  if x.year == emptyYear { return 0 }
  return x.year
}

func (x *calendarday) Val() uint {
  return uint(x.internalCode())
}

func (x *calendarday) SetVal (n uint) {
  if n < 1<<16 {
    x.decode(uint16(n))
  }
}

func (X *calendarday) WriteGr (x, y int) {
//  switch x.Format {
//  case W, M, Dd_M_yyyy:
//    var e calendarday; e.day, e.month, e.year = maxDay, maxMonth, emptyYear
//    e.Format = x.Format
//    e.cF, e.cB = x.cF, x.cB
//    e.Write (l, c)
//  }
//   default:
  bx.Wd (wd[X.Format])
  bx.Colours (X.cF, X.cB)
  bx.WriteGr (X.String(), x, y)
}

func (x *calendarday) Write (l, c uint) {
  bx.Wd (wd[x.Format])
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}

func (x *calendarday) PosInWeek (vertical bool, a uint) (uint, uint) {
  if x.year == emptyYear { return 0, 0 }
  l := uint(0)
  S := a * uint(x.weekday())
  if vertical { l = S; S = 0 }
  return l, S
}

func (x *calendarday) WriteWeek (vertical bool, a, l, c uint) {
  if x.year == emptyYear { return }
// oldF, oldB := x.cF, x.cB
  if vertical {
    if a == 0 { a = 1 }
  } else {
    if a == 0 { a = wd[x.Format] + 1 }
  }
  y := x.Clone().(*calendarday)
  y.SetBeginning (Weekly)
  l1, c1 := uint(0), uint(0)
  for i := 0; i <= 6; i++ {
    if vertical {
      l1 = a * uint(i)
    } else {
      c1 = a * uint(i)
    }
    op (y)
    y.Write (l + l1, c + c1)
    y.Inc (Daily)
  }
// x.Colours (oldF, oldB)
}

/*
func (x *calendarday) changeWithMouseInWeek (vertical bool, a, l, c uint) {
  x.changeWithMouse (Weekly, vertical, a, 0, 0, l, c)
} */

func (x *calendarday) PosInMonth (vertical bool, n, z, s uint) (uint, uint) {
  if x.year == emptyYear { return 0, 0 }
  if n == 0 { n = 1 }
  n = 7 * n
  if vertical {
    if z == 0 { z = 1 }
  } else {
    if s == 0 { s = wd[x.Format] + 1 }
  }
  d := x.Clone().(*calendarday)
  d.day = 1
  i := uint (d.weekday()) + x.day - 1
  if vertical {
    return z * (i % n), s * (i / n)
  }
  return z * (i / n), s * (i % n)
}

func (x *calendarday) WriteMonth (vertical bool, n, l0, c0, l, c uint) {
  if x.year == emptyYear { return }
  t := x.daysInMonth()
  if n == 0 { n = 1 }
  n = 7 * n
  if vertical {
    if l0 == 0 { l0 = 1 }
  } else {
    if c0 == 0 { c0 = wd[x.Format] + 1 }
  }
  max := int((maxDay - 2) / n + 2)
  max *= int(n)
  y := x.Clone().(*calendarday)
  y.day = 1
  w := int(y.weekday())
  var e calendarday
  e.day, e.month, e.year = maxDay, maxMonth, emptyYear
  e.Format = x.Format
  e.Colours (col.Blue(), WeekdayB)
  var l1, c1 uint
  for i := 0; i < max; i++ {
    if vertical {
      l1, c1 = l0 * (uint(i) % n), c0 * (uint(i) / n)
    } else {
      l1, c1 = l0 * (uint(i) / n), c0 * (uint(i) % n)
    }
    if i < w || i >= w + int(t) {
      e.Write (l + l1, c + c1)
    } else {
      op (y)
      y.Write (l + l1, c + c1)
      if y.day < t {
        y.day ++
      }
    }
  }
}

/*
func (x *calendarday) changeWithMouseInMonth (vertical bool, n, l, c, l1, c1 uint) {
  x.changeWithMouse (Monthly, vertical, n, l, c, l1, c1)
}
*/

func (x *calendarday) PrintMonth (vertical bool, n, z, s, l, S uint) {
  if x.year == emptyYear { return }
  if n == 0 { n = 1 }
  n = 7 * n
  if vertical {
    if z == 0 { z = 1 }
  } else {
    if s == 0 { s = wd[x.Format] + 1 }
  }
  max := (maxDay - 2) / n + 2
  max = n * max
  y := x.Clone().(*calendarday)
  y.SetFormat (Dd)
  y.day = 1
  W := uint(y.weekday())
  t := x.daysInMonth()
  for i := uint(0); i < max; i++ {
    var l1, S1 uint
    if vertical {
      l1 = z * (i % n)
      S1 = s * (i / n)
    } else {
      l1 = z * (i / n)
      S1 = s * (i % n)
    }
    if i < W || i >= W + t {
      // pbx.Clr (l + l1, S + S1) // TODO
    } else {
      op (y)
      if y.IsHoliday() {
        y.SetFont (font.Bold)
      } else {
        y.SetFont (font.Roman)
      }
      y.Print (l + l1, S + S1)
      if y.day < t {
        y.day++
      }
    }
  }
}

const (
  monthsHorizontally = 4
  leftMargin = uint(5)) // mindestens 3, höchstens 5

func (x *calendarday) shift (l, c *uint) {
  *l += (7 + 1) * ((x.month - 1) / monthsHorizontally)
  *c += (7 - 1) * (2 + 1) * ((x.month - 1) % monthsHorizontally)
}

func (x *calendarday) PosInYear() (uint, uint) {
  l, c := x.PosInMonth (true, 1, 1, 3)
  x.shift (&l, &c)
  l ++
  c += leftMargin
  return l, c
}

func (x *calendarday) changeWithMouse (p Period, vertical bool, a, l, c, l0, c0 uint) {
  switch p {
  case Daily, Decadic:
    return;
  }
  lm, cm := scr.MousePos()
//  if p == Yearly { SM += leftMargin }
  y := x.Clone().(*calendarday)
  y.SetBeginning (p)
  A := y.Clone().(*calendarday)
  y.SetFormat (Dd_mm_yy)
  n := wd[x.Format]
  var lpos, cpos uint
  for {
    n = wd[x.Format]
    if ! y.Equiv (A, p) {
      break
    }
    switch p {
    case Weekly:
      lpos, cpos = y.PosInWeek (vertical, a)
      lpos += l0; cpos += c0
    case Monthly:
      lpos, cpos = y.PosInMonth (vertical, a, l, c)
      lpos += l0; cpos += c0
    case Quarterly:
      errh.Error ("in day not yet implemented", 3)
      return
    case HalfYearly:
      errh.Error ("in day not yet implemented", 6)
      return
    case Yearly:
      n = 2
      lpos, cpos = y.PosInYear()
    default:
    }
    if lm == lpos && cpos <= cm && cm < cpos + n {
      y.Copy (x)
      break
    } else {
      y.Inc (Daily)
    }
  }
}

func (x *calendarday) writeYearMask (l, c uint) {
  const X = 80
  scr.Clr (l, c, X, 25 - 1)
  y := x.Clone().(*calendarday)
  y.Format = Yyyy
  y.Colours (YearnumberF, YearnumberB)
  y.Write (l, c)
  y.Write (l, c + 80 - 4)
  bx.Colours (MonthF, MonthB)
  bx.Wd (X - 4 /* - c */)
  T1 := str.New (X - 4 /* - c */)
  bx.Write (T1, l, c + 4)
  bx.Wd (X /* - c */)
  T1 = str.New (X /* - c */)
  bx.Write (T1, l + 8, c)
  bx.Write (T1, l + 16, c)
  y.Format = M
  var l1, c1 uint
  for m := uint(1); m <= maxMonth; m++ {
    y.month = m
    l1, c1 = 0, leftMargin
    y.shift (&l1, &c1)
    y.Colours (MonthF, MonthB)
    y.Write (l + l1, c + c1 + 3)
  }
  bx.Colours (WeekdayNameF, WeekdayNameB)
  y.Format = Wd
  c2 := leftMargin + monthsHorizontally * 6 * 3 // 6: columsn per month, 3: tt-Format + 1 spaces
  bx.Wd (2) // len (WdShorttext)
  for m := uint(1); m <= maxMonth; m += monthsHorizontally {
    for w := Monday; w <= Sunday; w++ {
      for i := uint(0); i <= 2; i++ {
        l1 = 1 + 8 * i + uint(w)
        bx.Write (WdShorttext [w], l + l1, c + 1)
        bx.Write (WdShorttext [w], l + l1, c + c2)
      }
    }
  }
}

func (x *calendarday) WriteYear (l, c uint) {
  if x.year == emptyYear { return }
  var l1, c1 uint
  x.writeYearMask (l, c)
  y := x.Clone().(*calendarday)
  y.Format = Dd
  for m := uint(1); m <= maxMonth; m++ {
    y.day, y.month = 1, uint(m)
    l1, c1 = 1, leftMargin
    y.shift (&l1, &c1)
    y.WriteMonth (true, 1, 1, 3, l + l1, c + c1)
  }
}

func (x *calendarday) EditInYear (l, c uint) {
  for {
    c, _ := kbd.Command()
    if c == kbd.Here {
      x.changeWithMouse (Yearly, false, 0, 0, 0, 0, 0)
      break
    }
  }
}

func (x *calendarday) printYearMask (l, c uint) {
  const X = 80
  y := x.Clone().(*calendarday)
  y.Format, y.Font = Yyyy, font.Bold
  y.Print (l, c)
  y.Print (l, X - 4)
  pbx.Print (str.New (X - 4 - c), l, c + 4)
  pbx.Print (str.New (X - c), l + 8, c)
  pbx.Print (str.New (X - c), l + 16, c)
  y.Format, y.Font = M, font.Italic
  var l1, c1 uint
  for m := uint(1); m <= maxMonth; m++ {
    y.month = m
    l1, c1 = 0, leftMargin
    y.shift (&l1, &c1)
    y.Print (l + l1, c + c1 + 3)
  }
  y.Format = Wd
  c2 := leftMargin + monthsHorizontally * 6 * 3
                // 6 Spalten pro Monat     tt-Format + 1 Zwischenraum
  pbx.SetFont (font.Italic)
  for m := uint(1); m <= maxMonth; m += monthsHorizontally {
    for w := Monday; w <= Sunday; w++ {
      for i := uint(0); i <= 2; i++ {
        l1 = 1 + 8 * i + uint(w)
        pbx.Print (WdShorttext [w], l + l1, c + 1)
        pbx.Print (WdShorttext [w], l + l1, c + c2)
      }
    }
  }
}

func (x *calendarday) PrintYear (l, c uint) {
  if x.year == emptyYear { return }
  x.printYearMask (l, c)
  y := new_().(*calendarday)
  y.Format = Dd
  for m := uint(1); m <= maxMonth; m++ {
    y.day, y.month, y.year = 1, m, x.year
    y.year = x.year
    dl, dc := uint(1), leftMargin
    y.shift (&dl, &dc)
    y.PrintMonth (true, 1, 1, 3, l + dl, c + dc)
  }
}

func isYear (y *uint) bool {
  if *y < uint(100) {
    *y += currentCentury
    if *y > limitYear {
      *y -= 100
    }
  }
  return startYear <= *y && *y <= endYear
}

func isMonth (m *uint, s string) bool {
  n := str.ProperLen (s)
  if n > 0 {
    for i := uint(1); i <= maxMonth; i++ {
      t := str.Part (nameMonth[i], 0, n)
      if s == t { // str.QuasiEq (s, t) {
        *m = uint(i)
        return true
      }
    }
  }
  return false
}

func (x *calendarday) defined (d, m, y uint) bool {
  if d == 0 || d > maxDay { return false }
  if m == 0 || m > maxMonth { return false }
  if ! isYear (&y) { return false }
  x.day, x.month, x.year = d, m, y
  return x.day <= x.daysInMonth()
}

func (x *calendarday) Defined (s string) bool {
  if str.Empty (s) { x.Clr(); return true }
  d := x.Clone().(*calendarday)
  var T string
  var l, p uint
  k, ss, P, L := N.DigitSequences (s)
  ok := false
  switch x.Format {
  case Dd, // e.g. " 8"
       Dd_mm_, // e.g. " 8.10."
       Dd_mm_yy, // e.g. " 8.10.07"
       Dd_mm_yyyy: // e.g. " 8.10.2007" *):
    switch k {
    case 1:
      l = 2
    case 2, 3:
      l = L[0]
    default:
      return false
    } // see below
  case Dd_M, // e.g. "8. Oktober"
       Dd_M_yyyy: // e.g. "8. Oktober 2007"
    if x.Format == Dd_M {
      if k != 1 { return false }
    } else {
      if k != 2 { return false }
    }
    if _, ok := str.Pos (s, '.'); ! ok { return false }
    if x.Format == Dd_M_yyyy {
//      l = str.ProperLen (s)
//      T = str.Part (s, p, l - p)
      T = ss[1]
      if d.year, ok = N.Natural (T); ! ok { return false }
    }
    T = ss[0]
    str.Move (&T, true)
    if d.day, ok = N.Natural (T); ! ok { return false }
    T = str.Part (s, p + 1, P[1] - p - 1)
    str.Move (&T, true)
    if ! isMonth (&d.month, T) { return false }
    return x.defined (d.day, d.month, d.year)
  case Yymmdd: // e.g. "090418"
    if d.year, ok = N.Natural (str.Part (s, 0, 2)); ! ok { return false }
    if d.month, ok = N.Natural (str.Part (s, 2, 2)); ! ok { return false }
    if d.day, ok = N.Natural (str.Part (s, 4, 2)); ! ok { return false }
    return x.defined (d.day, d.month, d.year)
  case Yyyymmdd: // e.g. "20090418"
    if d.year, ok = N.Natural (str.Part (s, 0, 4)); ! ok { return false }
    if d.month, ok = N.Natural (str.Part (s, 4, 2)); ! ok { return false }
    if d.day, ok = N.Natural (str.Part (s, 6, 2)); ! ok { return false }
    return x.defined (d.day, d.month, d.year)
  case Yy, // e.g. "08"
       Yyyy: // e.g. "2007"
    if k != 1 { return false }
    if d.year, ok = N.Natural (ss[0]); ok {
      return x.defined (d.day, d.month, d.year)
    } else {
      return false
    }
  case Wd, // e.g. "Mo"
       WD: // e.g. "Monday"
    return false // Fall noch nicht erledigt
  case Mmm, // e.g. "Mon"
       M: // e.g. "Oktober"
    if ! isMonth (&d.month, s) { return false }
    return x.defined (d.day, d.month, d.year)
  case Myyyy: // e.g. "Oktober 2007"
    if k != 1 { return false }
    if d.year, ok = N.Natural (ss[0]); ! ok { return false }
    if _, ok := str.Pos (s, ' '); ! ok { return false }
    if ! isMonth (&d.month, str.Part (s, 0, p)) { return false }
    return x.defined (d.day, d.month, d.year)
  case Wn, // e.g. "1" (.Woche)
       WN: // e.g. "1.Woche"
    if k != 1 { return false }
    k, ok = N.Natural (T)
    if ok {
      if 0 < k && k <= 3 {
        d.day, d.month, d.year = 1, 1, x.year
        c := d.internalCode()
        w := d.weekday()
        if w > Thursday { c += 7 } // see Weeknumber
        if c < uint16(w) { return false }
        c -= uint16(w) // so c is a Monday
        d.Decode (Encode (uint(c) + 7 * k))
        if d.year == x.year {
          x.day = d.day
          x.month = d.month
          return true
        }
      }
      return false
    }
  case WNyyyy: // e.g. "1.Woche 2007"
    return false // not yet implemented
  case Qu: // e.g. "  I/06"
    if k != 1 { return false }
    if _, ok := str.Pos (s, '/'); ! ok { return false }
    if d.year, ok = N.Natural (ss[0]); ! ok { return false }
    T = str.Part (s, 0, p)
    str.Move (&T, true)
    k = str.ProperLen (T)
    if T[0] != 'I' { return false }
    switch k {
    case 1:
      d.month = 1
    case 2:
      switch T [1] {
      case 'I':
        d.month = 4
      case 'V':
        d.month = 10
      default:
        return false
      }
    case 3:
      if T [1] == 'I' && T [2] == 'I' { d.month = 7 }
    default:
      return false
    }
    return x.defined (d.day, d.month, d.year)
  }
  if d.day, ok = N.Natural (str.Part (s, P[0], l)); ! ok { return false }
  if k == 1 {
    if L [0] > 8 { return false } // maximal "Dd_mm_yyyy"
    if L [0] > 2 {
      if d.month, ok = N.Natural (str.Part (s, P [0] + 2, 2)); ! ok { return false }
    }
    if L [0] > 4 {
      if d.year, ok = N.Natural (str.Part (s, P [0] + 4, L [0] - 4)); ! ok { return false }
    }
  } else { // n == 2, 3
    if d.month, ok = N.Natural (ss[1]); ! ok { return false }
    if k == 2 && x.Empty() {
      d.year = today.(*calendarday).year
    }
    if k == 3 {
      if d.year, ok = N.Natural (ss[2]); ! ok { return false }
    }
  }
  return x.defined (d.day, d.month, d.year)
}

/*
func Date (m, d, y int) *calendarday {
  y := new_()
  if y.Defined3 (m, d, y) { }
  return y
} */

func (x *calendarday) Selected (cop CondOp) bool {
  loop:
  for {
    cop (x, true) // colour auffallend
    c, t := kbd.Command()
    cop (x, false) // colour normal
    switch c {
    case kbd.Enter:
      return true
    case kbd.Esc:
      break loop
    default:
      if x.Format == Yy || x.Format == Yyyy {
        if t == 0 { t = 3 } else { t = 5 }
      }
      x.Change (c, t)
    }
  }
  return false
}

func (X *calendarday) EditGr (x, y int) {
  bx.Wd (wd[X.Format])
  bx.Colours (X.cF, X.cB)
  s := X.String()
  nErr := 0
  for {
    bx.EditGr (&s, x, y)
    if X.Defined (s) {
      X.WriteGr (x, y)
      return
/*
      var t uint
      if kbd.LastCommand (&t) == kbd.Search && t == 0 {
        if x.year == emptyYear { x.Copy (today) }
        errh.Hint ("Datum ändern: Kursortasten, Datum auswählen: Enter, Eingabe stornieren: Esc")
        loop: for {
          cm, _ := kbd.Command()
          x.Change (cm, t)
          x.Write (l, c)
          switch cm {
          case kbd.Enter, kbd.Here:
            errh.DelHint(); return
          case kbd.Back, kbd.There:
            errh.DelHint(); break loop
          }
        }
      } else {
        x.Write (l, c)
        return
      }
*/
    } else {
      l, c := 16 * uint(y), 8 * uint(x)
      nErr ++
      switch nErr {
      case 1:
        errh.Error0Pos ("Die Eingabe stellt kein Datum dar!", l + 1, c)
      case 2:
        errh.Error0Pos ("Das ist auch kein Datum!", l + 1, c)
      case 3:
        errh.Error0Pos ("Jetzt passen Sie doch mal auf!", l + 1, c)
      case 4:
        errh.Error0Pos ("Sind Sie zu doof, ein Datum einzugeben?", l + 1, c)
      default:
        errh.Error0Pos ("Was soll der Quatsch?", l + 1, c)
        X.Update()
        s = X.String()
      }
    }
  }
}

func (x *calendarday) Edit (l, c uint) {
  x1, y1 := scr.Wd1(), scr.Ht1()
  x.EditGr (int(x1 * c), int(y1 * l))
}

func (x *calendarday) Codelen() uint {
  return Codelen(uint16(0))
}

func (x *calendarday) Encode() Stream {
  B := make (Stream, Codelen(uint16(0)))
  copy (B, Encode (uint16(x.internalCode())))
  return B
}

func (x *calendarday) decode (n uint16) {
  var d uint16
  if n == 0 {
    x.year = emptyYear
  } else {
    x.year = startYear
    for {
      d = daysInYear (x.year)
      if n > d {
        x.year++
        n -= d
      } else {
        break
      }
    }
    x.month = 1
    for {
      d = uint16(x.daysInMonth())
      if n > d {
        x.month++
        n -= d
      } else {
        break
      }
    }
    x.day = uint(n)
  }
}

func (x *calendarday) Decode (B Stream) {
  c := Decode (uint16(0), B).(uint16)
  if c <= uint16(maxCode) {
    x.decode (c)
  } else {
    x.Clr()
  }
}

/*
func decode (B Stream) Calendarday {
  x  := new_().(*calendarday)
  c := Decode (uint16(0), B).(uint16)
  if c <= uint16(maxCode) {
    x.decode (c)
  } else {
    x.Clr()
  }
  return x
}
*/

func (x *calendarday) SetFont (f font.Font) {
  x.Font = f
}

func (x *calendarday) Print (l, c uint) {
/*
  if x.IsHoliday() && x.Format <= Dd_M_yyyy {
    x.Font = font.Bold
  } else {
    x.Font = font.Roman
  }
  if x.Format == Yyyy || x.Format == M {
    x.Font = font.Italic
  }
*/
  pbx.SetFont (x.Font)
  pbx.Print (x.String(), l, c)
}

func (x *calendarday) SetAttribute (p DayOp) {
  op = p
}

func attribute (D Calendarday) {
  d := D.(*calendarday)
  switch d.Format {
  case Dd, Dd_mm_, Dd_mm_yy, Yymmdd, Yyyymmdd, Dd_mm_yyyy, Dd_M, Dd_M_yyyy:
    if d.IsHoliday() {
      d.Colours (HolidayF, HolidayB)
      d.SetFont (font.Bold)
    } else {
      d.Colours (WeekdayF, WeekdayB)
      d.SetFont (font.Roman)
    }
  case Yy, Yyyy:
    d.Colours (YearnumberF, YearnumberB)
    d.SetFont (font.Bold)
  case Wd, WD:
    d.Colours (WeekdayNameF, WeekdayNameB)
    d.SetFont (font.Italic)
//  case Mmm:
  case M, Myyyy:
    d.Colours (MonthF, MonthB)
    d.SetFont (font.Italic)
  default:
    d.Colours (WeekdayF, MonthB)
    d.SetFont (font.Italic)
  }
}
