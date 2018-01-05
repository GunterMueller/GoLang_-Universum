package main

// (c) Christian Maurer   v. 171230 - license see nU.go

func merge (a []int, m int) {
  n := len(a)
  b := make([]int, len(a))
  for j, i, k := 0, 0, m; j < n; j++ {
    if i < m && k < n {
      if a[i] < a[k] {
        b[j] = a[i]; i++
      } else {
        b[j] = a[k]; k++
      }
    } else if i < m {
      b[j] = a[i]; i++
    } else if k < n {
      b[j] = a[k]; k++
    }
  }
  copy (a, b)
}

func mergesort (a []int, e chan int) {
  if len(a) > 1 {
    m := len(a)/2
    c, d := make(chan int), make(chan int)
    go mergesort (a[:m], c)
    go mergesort (a[m:], d)
    <-c
    <-d
    merge (a, m)
  }
  e <- 0
}

func main() {
  done := make(chan int)
  is := []int{7,9,1,4,0,6,8,2,5,3}
  go mergesort (is, done)
  <-done
  for _ , i := range is {
    print(i, " ")
  }
  println()
}
