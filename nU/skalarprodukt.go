package main

type vector [3]int

func scalarproduct (v, w vector, p *int, d chan int) {
  for j := 0; j < 3; j++ {
    *p += v[j] * w[j]
  }
  d <- 0
}

type matrix [3]vector

func column (a matrix, k int) (s vector) {
  for j := 0; j < 3; j++ {
    s[j] = a[j][k]
  }
  return s
}

func product (a, b matrix) (p matrix) {
  done := make (chan int)
  for i := 0; i < 3; i++ {
    for k := 0; k < 3; k++ {
      go scalarproduct (a[i], column(b, k), &p[i][k], done)
    }
  }
  for j := 0; j < 9; j++ {
    <-done
  }
  return p
}

func main() {
  a := matrix { vector{1,2,3}, vector{4,5,6}, vector{7,8,9} }
  b := matrix { vector{9,8,7}, vector{6,5,4}, vector{3,2,1} }
  c := product (a, b)
  for i := 0; i < 3; i++ {
    for k := 0; k < 3; k++ {
      print(c[i][k], " ")
    }
    println()
  }
}
