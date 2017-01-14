package barr

// (c) murus.org  v. 140330 - license see murus.go

// >>> implementation with semaphores
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 102

import
  "murus/sem"
type
  barrier struct {
        involved,
         waiting uint
           mutex,
               s sem.Semaphore
                 }

func New (n uint) Barrier {
  if n < 2 {
    return nil
  }
  x:= new (barrier)
  x.involved = n
  x.mutex = sem.New(1)
  x.s = sem.New(0)
  return x
}

func (x *barrier) Wait() {
  x.mutex.P()
  x.waiting++
  if x.waiting < x.involved {
    x.mutex.V()
    x.s.P()
    // x.mutex ist übernommen
    x.waiting--
    if x.waiting == 0 {
      x.mutex.V()
    } else {
      x.s.V()
    }
  } else { // waiting > 1
    x.waiting--
    x.s.V()
    // x.mutex wird übergeben
  }
}
