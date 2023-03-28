package netz

// (c) Christian Maurer   v. 230305 - license see µU.go

import
  "µU/env"

func init() {
  if env.NArgs() == 0 {
    MeinBahnhof = Server
  } else {
    MeinBahnhof = env.N(1)
    MeinName = name (MeinBahnhof)
  }
}

func name (n uint) string {
  switch n {
  case 0:
    return "Bahnheim"
  case 1:
    return "Bahnhausen"
  case 2:
    return "Bahnstadt"
  case 3:
    return "Eisenheim"
  case 4:
    return "Eisenstadt"
  case 5:
    return "Eisenhausen"
  }
  return "Server"
}

func nachbarn (n uint) []uint {
  nb := make([]uint, 0)
  switch n {
  case Bahnheim:
    nb = append (nb, Bahnhausen)
  case Bahnhausen:
    nb = append (nb, Bahnheim)
    nb = append (nb, Bahnstadt)
  case Bahnstadt:
    nb = append (nb, Bahnhausen)
    nb = append (nb, Eisenheim)
    nb = append (nb, Eisenstadt)
  case Eisenheim:
    nb = append (nb, Bahnstadt)
  case Eisenstadt:
    nb = append (nb, Bahnstadt)
    nb = append (nb, Eisenhausen)
  case Eisenhausen:
    nb = append (nb, Eisenstadt)
  }
  return nb
}

func anzahlNachbarn (n uint) uint {
  return uint(len(nachbarn(n)))
}

func nachbar (n, i uint) uint {
  switch n {
  case Bahnheim:
    switch i {
    case 0: return Bahnhausen
    }
  case Bahnhausen:
    switch i {
    case 0: return Bahnheim
    case 1: return Bahnstadt
    }
  case Bahnstadt:
    switch i {
    case 0: return Bahnhausen 
    case 1: return Eisenheim
    case 2: return Eisenstadt
    }
  case Eisenheim:
    switch i {
    case 0: return Bahnstadt
    }
  case Eisenstadt:
    switch i {
    case 0: return Bahnstadt
    case 1: return Eisenhausen
    }
  case Eisenhausen:
    switch i {
    case 0: return Eisenstadt
    }
  }
  return Server
}
