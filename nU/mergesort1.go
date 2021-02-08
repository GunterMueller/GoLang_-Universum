package main // mergesort with message passing

import "math/rand"

const (
  N = 128 // power of 2 !
  max = 2 * N - 2
)
var (
  c [max+1]chan int
  done chan bool
)

func generate() {
  for s := 0; s < N; s++ {
    n := rand.Intn(1000)
    c[max] <- n // unsortiert an die Ausgabe
    c[s] <- n // und zu den Sortierern
  }
}

func nAwaitedNumbers (i int) int {
  e, d, m := N / 2, 1, N - 2
  for m > i {
    d *= 2; m -= d; e /= 2
  }
  return e
}

func sort (i int) { // i = number of sorting processes
  rL, rR := 2 * i, 2 * i + 1 // number of left reso,
  t := N + i // right receive- and send-channels
  e := nAwaitedNumbers(i) // number of the messages
  e0 := 1 // to be awaited, also from the left
  nL := <-c[rL]
  e1 := 1 // and from the right
  nR := <-c[rR]
  for e0 <= e && e1 <= e {
    if nL <= nR {
      c[t] <- nL
      e0++
      if e0 <= e {
        nL = <-c[rL]
      }
    } else {
      c[t] <- nR
      e1++
      if e1 <= e {
        nR = <-c[rR]
      }
    }
  }
  for e0 <= e {
    c[t] <- nL
    e0++; if e0 <= e {
      nL = <-c[rL]
    }
  }
  for e1 <= e {
    c[t] <- nR
    e1++
    if e1 <= e {
      nR = <-c[rR]
    }
  }
}

func write() {
  println("randomly generated:")
  for i := 0; i < N; i++ {
    print(<-c[max], " ")
  }
  println(); println("sorted:")
  for i := 0; i < N; i++ {
    print(<-c[max], " ")
  }
  println(); done <- true
}

func main() {
  for i := 0; i <= max; i++ {
    c[i] = make(chan int)
  }
  done = make(chan bool)
  go generate()
  for i := 0; i <= N - 2; i++ {
    go sort (i)
  }
  go write()
  <-done
}
