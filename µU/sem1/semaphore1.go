package sem1

// (c) Christian Maurer   v. 240930 - license see µU.go

// >>> Implementation with asynchronous message passing

type
  semaphore1 struct {
                 ch chan int
                    }

func new_() Semaphore1 {
  x := new(semaphore1)
  x.ch = make(chan int, 1)
  x.ch <- 0
  return x
}

func (x *semaphore1) P() {
  <-x.ch
}

func (x *semaphore1) V() {
  x.ch <- 0
}
