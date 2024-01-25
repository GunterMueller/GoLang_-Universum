package sem

// (c) Christian Maurer   v. 170121 - license see nU.go

type
  semaphore struct {
                 c chan int
                   }

func new_(n uint) Semaphore {
  x := new(semaphore)
  x.c = make(chan int, n)
  for i := uint(0); i < n; i++ {
    x.c <- 0
  }
  return x
}

func (x *semaphore) P() {
  <-x.c
}

func (x *semaphore) V() {
  x.c <- 0
}
