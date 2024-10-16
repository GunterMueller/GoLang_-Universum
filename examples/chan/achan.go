package main

// (c) Christian Maurer   v. 241005 - license see µU.go

import
  "µU/achan"

func main() {
  const n = 1000
  x := achan.New (0)
  x.Send (n)
  println (x.Recv().(int))
  for i := 0; i < n; i++ { x.Send (i) }
  for i := 0; i < n; i++ { println (x.Recv().(int)) }
  println (x.Recv().(int))
  return
  c := make(chan int, n)
  c <- n
  println (<-c)
  for i := 0; i < n; i++ { c <- i }
  for i := 0; i < n; i++ { println (<-c) }
  println (<-c)
}
