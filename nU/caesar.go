package main

import ("os"; "bufio")

const eol = 13
var input = bufio.NewReader(os.Stdin)

func cap (b byte) byte {
  if b >= 'a' {
    return b - 'a' + 'A'
  }
  return b
}

func accepted (b byte) bool {
  B := cap(b)
  if 'A' <= B && B <= 'Z' || b == ' ' {
    return true
  }
  return false
}

func dictate (t chan byte) {
  for {
    b, _ := input.ReadByte()
    if accepted (b) {
      t <- b
    }
    if b == eol {
      break
    }
  }
}

func encrypt (t, c chan byte) {
  for {
    b := <-t
    if b == ' ' || b == eol {
      c <- b
    } else if cap(b) < 'X' {
      c <- b + 3
    } else {
      c <- b - 23
    }
  }
}

func send (c, m chan byte) {
  b:= byte(0)
  for b != eol {
    b = <-c
    print(string(b))
  }
  println()
  m <- 0
}

func main () {
  input = bufio.NewReader(os.Stdin)
  textchan := make(chan byte)
  cryptchan := make(chan byte)
  messagerchan := make(chan byte)
  go dictate (textchan) // Caesar
  go encrypt (textchan, cryptchan) // Offizier
  go send (cryptchan, messagerchan) // Bote
  <-messagerchan
}
