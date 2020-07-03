package main

import (
  "runtime"
  "time"
  "math/rand"
  "nU/lockn"
)
const (
  ncpu = 2
  m = 8
  N = 10 * 1000
)
var (
  n uint
/*
  x = lockn.NewDijkstra (m)
  x = lockn.NewHabermann (m)
  x = lockn.NewBakery (m)
  x = lockn.NewBakery1 (m)
  x = lockn.NewTicket (m)
  x = lockn.NewTiebreaker (m)
  x = lockn.NewKessels (m) // Pre: m is a power of 2.
  x = lockn.NewSzymanski (m)
  x = lockn.NewKnuth (m)
  x = lockn.NewDeBruijn (m)
  x = lockn.NewChannel (m)
*/
  x = lockn.NewKessels (m) // Pre: m is a power of 2.
  c = make(chan int)
)

func v (p uint, b bool) {
  t := int64 (1e6)
  if b && p > 1 { t *= int64(p) / 2 }
  time.Sleep (time.Duration (rand.Int63n(t)))
}

func count (p uint) {
  for i := 0; i < N; i++ {
    x.Lock (p)
    accu := n
    v (p, false)
    accu++
    v (p, false)
    n = accu
if n % 100 == 0 { print (n, " ") }
    x.Unlock (p)
    v (p, true)
  }
  c <- 0
}

func main() {
  runtime.GOMAXPROCS (ncpu)
  for p := uint(0); p < m; p++ { go count(p) }
  for i := 0; i < m; i++ { <-c }
  println(); println ("locktest: counter =", n, "<=", m * N)
}
