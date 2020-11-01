package sort

// (c) Christian Maurer   v. 121125 - license see µU.go

import
  . "µU/obj"


func quicksort (a []Any, min, max int) {
  i, j := min, max
  m := a [(i + j) / 2]
  for i <= j {
    for Less (a[i], m) {
      if i < max { i++ }
    }
    for Less (m, a[j]) {
      if j > min { j-- }
    }
    if i <= j {
      if i < j {
        a[i], a[j] = a[j], a[i]
      }
      if i < max { i++ }
      if j > min { j-- }
    }
  }
  if min < j { quicksort (a, min, j) }
  if i < max { quicksort (a, i, max) }
}


func sort_ (a []Any) {
  n := len (a)
  if n <= 1 { return }
  quicksort (a, 0, n - 1)
}
