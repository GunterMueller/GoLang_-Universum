package lock

// (c) Christian Maurer   v. 111125 - license see µu.go

func null() {
  for {
    if true == false {
      for i:= uint(0); i <= 1 << 32 - 1; i++ {
        if 0 != 1 {
          return
        }
      }
    } else {
      break
    }
  }
}
