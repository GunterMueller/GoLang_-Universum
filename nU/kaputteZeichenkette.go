package main

import ("strings"; "time")

var (
  a, b string
  done = make(chan int)
)

func pause() {
  time.Sleep (1e6)
}

func write() {
  for n := 0; n < len(a); n++ {
    switch n {
    case 5, 11, 14, 22:
      pause()
    }
    a = strings.Replace (a, string(byte(n % 10) + '0'),
                            string(byte(n) + 'a'), 1)
  }
  done <- 0
}

func read() {
  for n := 0; n < len(a); n++ {
    switch n {
    case 3, 10, 17, 23:
      pause()
    }
    b += string(a[n])
  }
  done <- 0
}

func main () {
  a = "01234567890123456789012345"
  go write()
  go read()
  <-done
  <-done
  println(b)
}
