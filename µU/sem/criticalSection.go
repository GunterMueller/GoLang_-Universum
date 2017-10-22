package sem

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> Implementation with a critical section

import
  "µU/cs"
type
  semaphoreCS struct {
                         cs.CriticalSection
                         }

func newCS (n uint) Semaphore {
  val := n
  x := new(semaphoreCS)
  c := func (i uint) bool {
         return val > 0
       }
  f := func (i uint) uint {
         val--
         return val
       }
  l := func (i uint) {
         val++
       }
  x.CriticalSection = cs.New (1, c, f, l)
  return x
}

func (x *semaphoreCS) P() {
  x.Enter (0)
}

func (x *semaphoreCS) V() {
  x.Leave (0)
}
