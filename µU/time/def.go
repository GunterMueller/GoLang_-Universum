package time

// (c) Christian Maurer   v. 220713 - license see µU.go

func Sleep (s uint) { sleep(s) }
func Msleep (s uint) { msleep(s) }
func Usleep (s uint) { usleep(s) }

// func Mess0() { mess0() }
// func Mess (s string) { mess(s) }

func Secmsec() uint { return secmsec() }
func Secµsec() uint { return secµsec() }
func Secnsec() uint { return secnsec() }

func UpdateTime() (uint, uint, uint) { return uTime() }
func UpdateDate() (uint, uint, uint) { return uDate() }

func SecondsSinceUnix() (s, us uint) { return sSU() }
