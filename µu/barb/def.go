package barb

// (c) Christian Maurer   v. 170731 - license see µu.go

type
  Barber interface {

  Customer()
  Barber()
}

func NewDir() Barber { return newD() }
func NewAndrews() Barber { return newA() }
func NewMon() Barber { return newM() }
func NewSem() Barber { return newS() }
func NewCondMon() Barber { return newCM() }
