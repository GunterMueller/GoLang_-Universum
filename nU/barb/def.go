package barb

// (c) Christian Maurer   v. 171019 - license see nU.go

type
  Barber interface {

  Customer()
  Barber()
}

func NewDir() Barber { return newD() }
func NewAndrews() Barber { return newA() }
func NewSem() Barber { return newS() }
func NewCS() Barber { return newCS() }
func NewMon() Barber { return newM() }
func NewCondMon() Barber { return newCM() }
