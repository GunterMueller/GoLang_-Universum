package barr

// (c) Christian Maurer   v. 170410 - license see nU.go

// >>> implementation with semaphores

import
  "nU/sem"

type barrier struct {
               uint "Anzahl der beteiligten Prozesse"
            waiting uint
           mutex, s sem.Semaphore
                    }

func new_(m uint) Barrier {
  if m < 2 { return nil }
  x := new(barrier)
  x.uint = m
  x.mutex = sem.New(1)
  x.s = sem.New(0)
  return x
}

func (x *barrier) Wait() {
  x.mutex.P()
  x.waiting++
  if x.waiting < x.uint {
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
