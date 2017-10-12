package barb

// (c) Christian Maurer   v. 170731 - license see µu.go

import
  . "sync"
/*
var
  cinchair, bavail, bbusy int
// invariants: bbusy <= cinchair <= bavail

func Customer() {
  await cinchair < bavail
  cinchair ++
}

func Barber() {
  bavail ++
  await bbusy < cinchair
  bbusy ++
}
*/

type
  barberAndrews struct {
            barberfree,      // == bavail - cinchair
                     n int   // == cinchair - bbusy
                 mutex Mutex
       barberAvailable,      // signaled when barber > 0
         chairOccupied *Cond // signaled when n > 0
                       }

func newA() *barberAndrews {
  x := new (barberAndrews)
  x.barberAvailable = NewCond (&x.mutex)
  x.chairOccupied = NewCond (&x.mutex)
  return x
}

func (x *barberAndrews) Customer() {
  x.mutex.Lock()
  for x.barberfree == 0 { x.barberAvailable.Wait() }
  x.barberfree--
  x.n++
  x.chairOccupied.Signal()
  x.mutex.Unlock()
}

func (x *barberAndrews) Barber() {
  x.mutex.Lock()
  x.barberfree++
  x.barberAvailable.Signal()
  for x.n == 0 { x.chairOccupied.Wait() }
  x.n--
  x.mutex.Unlock()
}
