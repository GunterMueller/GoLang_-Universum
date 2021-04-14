package reg

// (c) Christian Maurer   v. 210331 - license see murus.go

type
  register struct {
                  uint "Wert des Registers"
                  }

func null() Register {
  return new(register)
}

func fail (s string) {
  panic ("Vor. von " + s + "() nicht eingehalten")
}

func (a *register) Inc() {
  a.uint++
}

func (a *register) Dec() {
  if a.uint <= 0 { fail("Dec") }
  a.uint--
}

func (a *register) Gt0() bool {
  return a.uint > 0
}

func (a *register) Add (b Register) Register {
  c := null().(*register)
  c.uint = a.uint + b.(*register).uint
  return c
}

func (a *register) Mul (b Register) Register {
  c := null().(*register)
  c.uint = a.uint * b.(*register).uint
  return c
}

func (a *register) Write() {
  z := a.uint
  if z == 0 {
    println ("0")
    return
  }
  s := ""
  if z < 0 {
    s = "-"
    z = -z
  }
  n := z
  var t string
  for t = ""; n > 0; n /= 10 {
    t = string(n % 10 + '0') + t
  }
  println (s + t)
}

type
  registers struct {
              regs []uint
                   }

func new_(a ...Register) Registers {
  x := new(registers)
  x.regs = make([]uint, len(a))
  for i, r := range a {
    x.regs[i] = r.(*register).uint
  }
  return x
}

func (x *registers) NotEmpty() bool {
  return len(x.regs) > 0
}

func (x *registers) Num() Register {
  n := new(register)
  n.uint = uint(len(x.regs))
  return n
}

func (x *registers) Head() Register {
  if len(x.regs) == 0 { fail("Head") }
  return &register { x.regs[0] }
}

func (x *registers) Tail() Registers {
  n := len(x.regs)
  if n == 0 { fail("Tail") }
  y := new(registers)
  y.regs = make([]uint, n - 1)
  for i := 0; i < n - 1; i++ {
    y.regs[i] = x.regs[i+1]
  }
  return y
}

func (x *registers) Cons (r Register) Registers {
  y := new(registers)
  n := len(x.regs)
  y.regs = make([]uint, n + 1)
  y.regs[0] = r.(*register).uint
  for i := 1; i < n + 1; i++ {
    y.regs[i] = x.regs[i-1]
  }
  return y
}
