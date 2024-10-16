package time

// (c) Christian Maurer   v. 240903 - license see µU.go

// The calling process is delayed by s seconds.
func Sleep (s uint) { sleep(s) }

// The calling process is delayed by s milliseconds.
func Msleep (s uint) { msleep(s) }

// The calling process is delayed by s nanoseconds.
func Usleep (s uint) { usleep(s) }

// func Mess0() { mess0() }
// func Mess (s string) { mess(s) }

// TODO Spec
func Secmsec() uint { return secmsec() }

// TODO Spec
func Secµsec() uint { return secµsec() }

// TODO Spec
func Secnsec() uint { return secnsec() }

// Returns the time at the moment of the call in form of hour, minute and second.
func ActTime() (uint, uint, uint) { return actTime() }

// Returns the time at the moment of the call in form of h, m, s and millisecond.
func ActTimeM() (uint, uint, uint, uint) { return actTimeM() }

// Returns the date at the moment of the call in form of day, month and year.
func ActDate() (uint, uint, uint) { return actDate() }

// Returns the numbers of secondes and nanoseconds
// elapsed since the 1st of January 1970.
func SecSinceUnix() (s, ns uint) { return secSinceUnix() }
