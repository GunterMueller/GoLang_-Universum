package lock

// (c) murus.org  v. 111125 - license see murus.go

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
