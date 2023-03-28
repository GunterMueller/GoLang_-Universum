package main

// (c) Christian Maurer   v. 230303 - license see µU.go

import (
  "os/exec"
  "math"
  "µU/env"
//  "µU/time"
  "µU/ker"
  "µU/kbd"
  "µU/str"
  "µU/col"
  "µU/linewd"
  "µU/scr"
  "µU/errh"
  "µU/N"
  "µU/R"
  "µU/spc"
  "µU/pseq"
  "µU/files"
  g "lsys/grammar"
  . "lsys/symstk"
)
const (
  wd = uint(800)
  ht = wd
)
var (
  x, y, z, x0, y0, z0, xmin, xmax, ymin, ymax, zmin, zmax, xm, ym, zm, mm, my float64
  alpha, delta float64
  noMeasure = false
  colour col.Colour
  gofile = pseq.New (Symbol(0))
  screen scr.Screen
  threedimensional = false
  eps = 1e-10
  ls = ""
  colF, colB col.Colour
)

func init() {
  colF, colB = col.LightWhite(), col.Black()
  colF, colB = col.Black(), col.LightWhite()
}

func arctan (x, y float64) float64 {
  if math.Abs(x) < eps { if y > 0 { return 90 } else { return 270 } }
  if math.Abs(y) < eps { if x > 0 { return  0 } else { return 180 } }
  a := math.Atan (y / x)
  if math.Abs(a) < eps { return 0 }
  alpha := a * 180 / math.Pi
  if x < 0 { alpha += 180 } else { if y < 0 { alpha += 360 } }
  return alpha
}

func step (s Symbol) {
  switch s {
  case g.Step:
    x0, y0, z0 = spc.GetOrigin()
    zero (&x0, &y0, &z0)
    spc.Move1Front (1)
    x, y, z = spc.GetOrigin()
    zero (&x, &y, &z)
    if threedimensional { // measure
      if x < xmin { xmin = x }; if x > xmax { xmax = x }
      if y < ymin { ymin = y }; if y > ymax { ymax = y }
      if z < zmin { zmin = z }; if z > zmax { zmax = z }
    }
  case g.YetiStep:
    x0, y0, z0 = spc.GetOrigin()
    zero (&x0, &y0, &z0)
    spc.Move1Front (1)
    x, y, z = spc.GetOrigin()
    zero (&x, &y, &z)
  case g.TurnLeft:
    spc.Turn (delta)
    ox, oy, _ := spc.GetOrigin(); if math.Abs(ox) < eps {ox = 0}; if math.Abs(oy) < eps {oy = 0}
    alpha = arctan (ox, oy)
  case g.TurnRight:
    spc.Turn (-delta)
    ox, oy, _ := spc.GetOrigin(); if math.Abs(ox) < eps {ox = 0}; if math.Abs(oy) < eps {oy = 0}
    alpha = arctan (ox, oy)
  case g.Invert:
    spc.Turn (180)
  case g.TiltDown:
    spc.Tilt (delta)
  case g.TiltUp:
    spc.Tilt (-delta)
  case g.RollLeft:
    spc.Roll (delta)
  case g.RollRight:
    spc.Roll (-delta)
  case g.BranchStart:
    spc.Push()
  case g.BranchEnd:
    spc.Pop()
  default:
    if noMeasure {
      var ok bool
      if colour, ok = g.IsColour (s); ok {
        scr.ColourF (colour)
      } else {
        scr.ColourF (colF)
      }
    }
  }
}

func zero (x, y, z *float64) {
  if math.Abs (*x) < eps { *x = 0 }
  if math.Abs (*y) < eps { *y = 0 }
  if math.Abs (*z) < eps { *z = 0 }
}

func measure (s Symbol) { // for 2-dimensional L-Systems
  step (s)
  switch s {
  case g.Step, g.YetiStep:
    if x < xmin { xmin = x }; if x > xmax { xmax = x }
    if y < ymin { ymin = y }; if y > ymax { ymax = y }
    if z < zmin { zmin = z }; if z > zmax { zmax = z }
  }
}

func ins (f pseq.PersistentSequence, s string) {
  for i := 0; i < len(s); i++ {
    f.Ins (s[i])
  }
}

func draw2 (s Symbol) {
  const d = 8
  step (s)
  if s == g.Step {
    xa, ya := int((x0 - xmin) * mm), int((y0 - ymin) * mm)
    xb, yb := int((x  - xmin) * mm), int((y  - ymin) * mm)
    xa += d; xb += d; ya += d; yb += d
    scr.Line (xa, ya, xb, yb)
  }
}

func r (x float64) string {
  return R.String(x)
}

func k (x float64) string {
  return r(x) + ", "
}

func draw3 (s Symbol) {
  step (s)
  if x < xmin { xmin = x }; if x > xmax { xmax = x } // measure for 3-dimensional L-Systems
  if y < ymin { ymin = y }; if y > ymax { ymax = y }
  if z < zmin { zmin = z }; if z > zmax { zmax = z }
  switch s {
  case g.Step:
    ins (gofile, "  x0, y0, z0 = " + k(x0) + k(y0) + r(z0)+ "\n")
    ins (gofile, "  x, y, z = " + k(x) + k(y) + r(z) + "\n")
    ins (gofile, "  gl.Line (x0, y0, z0, x, y, z)\n")
  case g.YetiStep, g.TurnLeft, g.TurnRight, g.Invert, g.TiltDown, g.TiltUp,
       g.RollLeft, g.RollRight, g.BranchStart, g.BranchEnd, g.PolygonStart, g.PolygonEnd:
    // ignore
  default:
    if f, ok := g.IsColour (s); ok {
      ins (gofile, "  gl.Colour (col." + f.String() + "())\n")
    }
  }
}

func execute (op func (Symbol)) {
  for i := len(g.Startword) - 1; i >= 0; i-- {
    Push (g.Startword[i], 0)
  }
  s, i, a, left := Symbol(0), uint(0), 0, ""
  for ! Empty() {
    s, i = Pop()
    t := string(s)
    if i == g.NumIterations || ! g.ExRule (t) {
      op (s)
    } else {
      a = 0
      left = string(s)
      a++
      for {
        if Empty() || a >= g.MaxL {
          break
        }
        i0 := i
        s, i = Pop()
        left += string(s)
        if g.ExRule (left) {
          // test for longer rule
        } else {
          Push (s, i)
          left = left[:len(left)-1]
          i = i0
          break
        }
      }
      right := g.Derivation (left)
      n := len(right)
      switch n {
      case 0:
        op (left[0])
      case 1:
        op (right[0])
      default:
        for j := n - 1; j > 0; j-- {
          Push (right[j], i + 1)
        }
        Push (right[0], i + 1)
      }
    }
  }
}

func main() {
  ls = env.Arg(1)
  if n, ok := str.Sub (".ls", ls); ok {
    str.Rem (&ls, n, 3)
  }
  files.Cd (env.Gosrc() + "/lsys")
  t := ls + ".ls"
  lsfile := pseq.New(Symbol(0))
  lsfile.Name (t)
  if lsfile.Empty() {
    files.Del (".", t)
    ker.Fin()
    return
  }
  t = ""
  for i := uint(0); i < lsfile.Num(); i++ {
    lsfile.Seek (i)
    s := lsfile.Get().(Symbol)
    t += string (s)
  }
  g.Initialize (ls)
  screen = scr.NewWH (0, 0, wd, ht); defer screen.Fin()
  screen.ScrColours (colF, colB)
  screen.Cls()
  scr.SetLinewidth (linewd.VeryFat)
  scr.SetLinewidth (linewd.Thin)
  scr.SetLinewidth (linewd.VeryThick)
  alpha, delta = g.Startangle, g.Turnangle
  a := float64(alpha) * math.Pi / 180
  cos, sin := math.Cos(a), math.Sin(a)
  if math.Abs(cos) < eps { cos = 0 }; if math.Abs(sin) < eps { sin = 0 }
  spc.Set (0, 0, 0, cos, sin, 0, 0, 0, 1)
  colour = g.StartColour
  execute (measure)
  symstk3 := []Symbol {g.TiltDown, g.TiltUp, g.RollLeft, g.RollRight}
  for _, symbol := range(symstk3) {
    if _, ok := str.Pos (t, symbol); ok {
      threedimensional = true
      break
    }
  }
  if threedimensional {
    goto dim3
  }
  alpha, delta = g.Startangle, g.Turnangle
  a = float64(alpha) * math.Pi / 180
  cos, sin = math.Cos(a), math.Sin(a)
  if math.Abs(cos) < eps { cos = 0. }; if math.Abs(sin) < eps { sin = 0. }
  spc.Set (0, 0, 0, cos, sin, 0, 0, 0, 1)
  x, y = 0, 0
  mm = float64(scr.Wd()) / (xmax - xmin)
  my = float64(scr.Ht()) / (ymax - ymin)
  if my < mm { mm = my }
  mm *= .98
  colour = g.StartColour
  scr.Name (ls)
  noMeasure = true
  spc.Set (0, 0, 0, cos, sin, 0, 0, 0, 1)
  scr.SetLinewidth (linewd.Thick)
  execute (draw2)
  kbd.Wait (false)
//  time.Sleep (3)
  return
dim3:
  if env.UnderC() { errh.Error0 ("a GUI is required for 3-dimensional L-Systems"); return }
  gofile = pseq.New (Symbol(0))
  filename, suffix := "zeig", ".go"
  gofile.Name (filename + suffix)
  gofile.Clr()
  gf := string('"')
  colour = g.StartColour
  ins (gofile, "package main\n\n")
  ins (gofile, "import (\n")
  ins (gofile, "  " + gf + "µU/col" + gf + "\n")
  ins (gofile, "  " + gf + "µU/gl" + gf + "\n")
  ins (gofile, "  " + gf + "µU/scr" + gf + "\n")
  ins (gofile, ")\n\n")
  ins (gofile, "func draw() {\n")
  ins (gofile, "  scr.Name (" + gf + ls + gf + ")\n")
  ins (gofile, "  gl.ClearColour (col.LightWhite())\n")
//    ins (gofile, "  gl.ClearColour (col.Black())\n")
  ins (gofile, "  gl.Clear()\n")
  ins (gofile, "  gl.Colour (col.Black())\n")
//    ins (gofile, "  gl.Colour (col.LightWhite())\n")
  ins (gofile, "  var x0, y0, z0, x, y, z float64\n")
  colour = g.StartColour
  ins (gofile, "  gl.Linewidth (5)\n")
  execute (draw3)
  xm, ym, zm = (xmin + xmax) / 2, (ymin + ymax) / 2, (zmin + zmax) / 2
  zero (&xm, &ym, &zm)
  ins (gofile, "}\n\n")
  ins (gofile, "func main() {\n")
  d := 0.5
  y = ymin - d * (xmax - xmin)
  ins (gofile, "  s := scr.NewWH (0, 0, " + N.String(wd) + ", " + N.String(ht) +
                                "); defer s.Fin()\n")
  ins (gofile, "  s.Go (scr.Look, draw, " + k(xm)+k(y)+k(zm) + k(xm)+k(ym)+k(zm) + "1, 0, 0)\n")
  ins (gofile, "}\n")
  exec.Command ("go", "build", filename + suffix).Run()
  files.Move (filename, env.Val ("GOBIN"))
  scr.Name (ls)
  defer exec.Command (filename).Run()
}
