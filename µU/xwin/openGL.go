package xwin

// (c) Christian Maurer   v. 201030 - license see µU.go

// #cgo LDFLAGS: -lX11 -lGL -lGLU
// #include <stdlib.h>
// #include <stdio.h>
// #include <X11/Xlib.h>
// #include <GL/gl.h>
// #include <GL/glx.h>
// #include <GL/glu.h>
// int etyp (XEvent *e) { return (*e).type; }
// unsigned int kState (XEvent *e) { return (*e).xkey.state; }
// unsigned int kCode (XEvent *e) { return (*e).xkey.keycode; }
import
  "C"
import (
//  "fmt"
  "math"
//  "time"
//  "µU/ker"
  "µU/gl"
  "µU/glu"
//  "µU/col"
  "µU/spc"
)
const (
  um = math.Pi / 180
  epsilon = 1e-6
)
const (
  right = 0; front = 1; top = 2
  Esc = 9; Enter = 36; Back = 22; Tab = 23
  Left = 113; Right = 114; Up = 111; Down = 116; PgUp = 112; PgDown = 117
  Pos1 = 110; End = 115; Ins = 118; Del = 119
  F1 = 67; F2 = 68; F3 = 69; F4 = 70; F9 = 75; F10 = 76; F11 = 96; F12 = 96
  Shift = 1; Strg = 4; Alt = 8; AltGr = 128
)
type
  d = C.GLdouble
var
  firstWrite = true

/*/
func (X *xwindow) fly() {
  for {
    time.Sleep (1e8)
    spc.Move (1, 1)
    X.write()
  }
}
/*/

func (X *xwindow) Go (m int, draw func(), ox, oy, oz, fx, fy, fz, tx, ty, tz float64) {
  X.origin.Set3 (ox, oy, oz)
  X.focus.Set3 (fx, fy, fz)
  X.top.Set3 (tx, ty, tz)
  X.top.Norm()
  gl.Enable (gl.Depthtest)
  gl.ShadeModel (gl.Flat)
//  gl.ShowLight (true)
  spc.Set (ox, oy, oz, fx, fy, fz, tx, ty, tz)
//  dfe := X.origin.Distance (X.focus)
//  delta := dfe / 500.
//  const nSteps = 3
//  Phi, Delta := [nSteps]float64 { 1, 9, 90 }, [nSteps]float64 { 1, 10, 100 }
  phi, delta0 := 3., 0.1
//  step := 1
  var xev C.XEvent
  redraw := true
//  if m == Fly { go X.fly() }
  for {
//    phi, delta := float64 (Phi[step]), delta0 * Delta[step]
    if redraw {
//      gl.Clear()
      ex, ey, ez, fx, fy, fz, nx, ny, nz := spc.Get()
      if math.Abs(fx) < epsilon { fx = 0 }; if math.Abs(fy) < epsilon { fy = 0 }
      gl.MatrixMode (gl.Projection)
      gl.LoadIdentity()
//      gl.Viewport (0, 0, X.wd, X.ht)
      glu.Perspective (60, X.proportion, 0.1, 10000.)
      C.gluLookAt (d(ex), d(ey), d(ez), d(fx), d(fy), d(fz), d(nx), d(ny), d(nz))
      draw()
      if e := gl.Error(); e != "" { println ("openGL error: " + e) }
      C.glXSwapBuffers (dpy, C.GLXDrawable(X.win))
      C.glFinish()
//      gl.MatrixMode (gl.Modelview) // obviously superfluous
//      print("origin    "); fmt.Println (ex, ey, ez)
//      print("focus  "); fmt.Println (fx, fy, fz)
//      print("normal "); fmt.Println (nx, ny, nz)
    }
    redraw = true
    C.XNextEvent (dpy, &xev)
    et := C.etyp (&xev)
    switch et {
    case C.KeyPress:
      c, t := C.kCode (&xev), C.kState (&xev)
// println (c, t)
      switch c {
      case Esc:
        return
      case Left:
        switch m {
        case Look:
          switch t {
          case 0:
            spc.TurnAroundFocus (top, phi) // turn
          case Shift:
            spc.Move (right, delta0) // move
          case Strg:
            spc.Turn (front, phi) // roll
          case Alt, AltGr:
            // TODO
          }
        case Walk:
          switch t {
          case 0:
            spc.Turn (top, phi) // turn
          case Shift:
            spc.Move (right, -delta0) // move
          case Strg:
            spc.Turn (front, phi) // roll
          }
        case Fly:
          if t == 0 {
            spc.Turn (top, phi) // turn
          } else {
            spc.Turn (front, phi) // roll
          }
        }
      case Right:
        switch m {
        case Look:
          switch t {
          case 0:
            spc.TurnAroundFocus (top, -phi) // turn
          case Shift:
            spc.Move (right, -delta0) // move
          case Strg:
            spc.Turn (front, -phi) // roll
          case Alt, AltGr:
            // TODO
          }
        case Walk:
          switch t {
          case 0:
            spc.Turn (top, -phi) // turn
          case Shift:
            spc.Move (right, delta0) // move
          case Strg:
            spc.Turn (front, -phi) // roll
          }
        case Fly:
          if t == 0 {
            spc.Turn (top, -phi) // turn
          } else {
            spc.Turn (front, -phi) // roll
          }
        }
      case Up:
        switch m {
        case Look:
          switch t {
          case 0:
            spc.TurnAroundFocus (right, phi) // tilt
          case Shift:
            spc.Move (top, delta0) // move
          case Strg:

          case Alt, AltGr:

          }
        case Walk:
          switch t {
          case 0:
            spc.Turn (right, phi) // tilt
          case Shift:
            spc.Move (top, delta0) // move
          }
        case Fly:
          spc.Turn (right, phi) // tilt
        }
      case Down:
        switch m {
        case Look:
        switch t {
          case 0:
            spc.TurnAroundFocus (right, -phi) // tilt
          case Shift:
            spc.Move (top, -delta0) // move
          case Strg:

          case Alt, AltGr:

          }
        case Walk:
          switch t {
          case 0:
            spc.Turn (right, -phi) // tilt
          case Shift:
            spc.Move (top, -delta0) // move
          }
        case Fly:
          spc.Turn (right, -phi) // tilt
        }
      case Enter:
        switch m {
        case Look:
          if t == 0 {
            spc.Move (front, delta0) // move ahead
          } else {
            spc.Move (front, -delta0) // move back
          }
        case Walk:
          if t == 0 {
            spc.Move (front, delta0) // move ahead
          } else {
            spc.Move (front, -delta0) // move back
            // spc.Move (front, delta0) // move TODO translate focus ?
          }
        case Fly:
          // TODO increase speed
        }
      case Back:
        switch m {
        case Look:
          spc.Move (front, -delta0) // move
        case Walk:
          if t == 0 {
            spc.Move (front, -delta0) // move
          } else {
            spc.Move (front, -delta0) // move TODO translate focus ?
          }
        case Fly:
          // TODO decrease speed
        }
      case Pos1:
        spc.Turn (front, phi)
      case End:
        spc.Turn (front, -phi)
/*/
      case F1: // quicker
        if step + 1 < nSteps {
          step++
        }
      case F2: // slow down
        if step > 0 {
          step --
        }
      case F3:
        if m == Walk {
          if t == 0 {
            spc.Invert() // turn 180°
          } else { // TODO
            dfe = X.origin.Distance (X.focus)
            spc.Move (front, 2 * dfe)
            spc.Invert()
          }
        }
      case F4:
        dfe = X.origin.Distance (X.focus)
        spc.Move (front, dfe)
        if t == 0 {
          spc.Move (right, dfe)
          spc.Turn (top, 90.)
        } else {
          spc.Move (right, -dfe)
          spc.Turn (top, -90.)
        }
      case Del: // TODO
        dfe = X.origin.Distance (X.focus)
        spc.Move (top, dfe)
        spc.Move (front, dfe)
        spc.Turn (right, -90) // x -> rechts, y -> oben
        if t != 0 { // y -> rechts, x -> unten
          spc.Turn (front, -90)
        }
      case F9:
        spc.SetLight (0)
      case F10:
        spc.SetLight (1)
      case F11:
        spc.SetLight (2)
      case F12:
        spc.SetLight (3)
/*/
      default:
        redraw = false
      }
    default:
      redraw = false
    }
  }
}
