package xwin

// (c) Christian Maurer   v. 191103 - license see µU.go

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
  "µU/show"
  "µU/gl"
  "µU/glu"
  "µU/col"
  "µU/spc"
)
const (
  um = math.Pi / 180
  epsilon = 1e-6
)
type
  d = C.GLdouble
var
  firstWrite = true

func (X *xwindow) Start (m show.Mode, draw func(), ex, ey, ez, fx, fy, fz, nx, ny, nz float64) {
  X.mode = m
  X.draw = draw
  X.eye.Set3 (ex, ey, ez)
  X.focus.Set3 (fx, fy, fz)
  X.delta = X.eye.Distance (X.focus)
  X.normal.Set3 (nx, ny, nz)
  X.normal.Norm()
}

/*/
func (X *xwindow) fly() {
  for {
    time.Sleep (1e8)
    spc.Move (1, 1)
    X.write()
  }
}
/*/

func (X *xwindow) write() {
  gl.ClearColor (col.White())
  gl.Clear()
  gl.Enable (gl.Depthtest)
  gl.ShadeModel (gl.Flat)
  ex, ey, ez, fx, fy, fz, nx, ny, nz := spc.Get()
  if math.Abs(fx) < epsilon { fx = 0 }; if math.Abs(fy) < epsilon { fy = 0 }
  gl.MatrixMode (gl.Projection)
  gl.LoadIdentity()
//  gl.Viewport (0, 0, X.wd, X.ht)
  glu.Perspective (60, X.proportion, 0.1, 1000.)
//  gl.MatrixMode (gl.Modelview) // implies, that gluLookAt does not work
  C.gluLookAt (d(ex), d(ey), d(ez), d(fx), d(fy), d(fz), d(nx), d(ny), d(nz))
  X.draw()
// if err := C.glGetError(); err != C.GL_NO_ERROR {println (err-1280)); ker.Panic ("openGL error")}
  C.glXSwapBuffers (dpy, C.GLXDrawable(X.win))
  C.glFinish()
//  gl.MatrixMode (gl.Modelview) // obviously superfluous
//  print("eye    "); fmt.Println (ex, ey, ez)
//  print("focus  "); fmt.Println (fx, fy, fz)
//  print("normal "); fmt.Println (nx, ny, nz)
}

func (X *xwindow) Go() {
  const (
    right = 0; front = 1; top = 2
    Esc = 9; Enter = 36; Back = 22; Tab = 23
    Left = 113; Right = 114; Up = 111; Down = 116; PgUp = 112; PgDown = 117
    Pos1 = 110; End = 115; Ins = 118; Del = 119
    F1 = 67; F2 = 68; F3 = 69; F4 = 70; F9 = 75; F10 = 76; F11 = 96; F12 = 96
    Shift = 1; Strg = 4; Alt = 8; AltGr = 128
  )
//  gl.ShowLight (true)
  ex, ey, ez := X.eye.Coord3()
  fx, fy, fz := X.focus.Coord3()
  nx, ny, nz := X.normal.Coord3()
  spc.Set (ex, ey, ez, fx, fy, fz, nx, ny, nz)
//  dfe := X.eye.Distance (X.focus)
//  delta := dfe / 500.
//  const nSteps = 3
//  Phi, Delta := [nSteps]float64 { 1, 9, 90 }, [nSteps]float64 { 1, 10, 100 }
  phi, delta0 := 3., 0.1
//  step := 1
  var xev C.XEvent
  redraw := true
//  if X.mode == show.Fly { go X.fly() }

  for {
//    phi, delta := float64 (Phi[step]), delta0 * Delta[step]
    if redraw {
      X.write()
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
        switch X.mode {
        case show.Look:
          switch t {
          case 0:
            spc.TurnAroundFocus (top, phi) // pan
          case Shift:
            spc.Move (right, delta0) // move
          case Strg:
            spc.Turn (front, phi) // roll
          case Alt, AltGr:
            // TODO
          }
        case show.Walk:
          switch t {
          case 0:
            spc.Turn (top, phi) // pan
          case Shift:
            spc.Move (right, -delta0) // move
          case Strg:
            spc.Turn (front, phi) // roll
          }
        case show.Fly:
          if t == 0 {
            spc.Turn (top, phi) // pan
          } else {
            spc.Turn (front, phi) // roll
          }
        }
      case Right:
        switch X.mode {
        case show.Look:
          switch t {
          case 0:
            spc.TurnAroundFocus (top, -phi) // pan
          case Shift:
            spc.Move (right, -delta0) // move
          case Strg:
            spc.Turn (front, -phi) // roll
          case Alt, AltGr:
            // TODO
          }
        case show.Walk:
          switch t {
          case 0:
            spc.Turn (top, -phi) // pan
          case Shift:
            spc.Move (right, delta0) // move
          case Strg:
            spc.Turn (front, -phi) // roll
          }
        case show.Fly:
          if t == 0 {
            spc.Turn (top, -phi) // pan
          } else {
            spc.Turn (front, -phi) // roll
          }
        }
      case Up:
        switch X.mode {
        case show.Look:
          switch t {
          case 0:
            spc.TurnAroundFocus (right, phi) // tilt
          case Shift:
            spc.Move (top, delta0) // move
          case Strg:

          case Alt, AltGr:

          }
        case show.Walk:
          switch t {
          case 0:
            spc.Turn (right, phi) // tilt
          case Shift:
            spc.Move (top, delta0) // move
          }
        case show.Fly:
          spc.Turn (right, phi) // tilt
        }
      case Down:
        switch X.mode {
        case show.Look:
        switch t {
          case 0:
            spc.TurnAroundFocus (right, -phi) // tilt
          case Shift:
            spc.Move (top, -delta0) // move
          case Strg:

          case Alt, AltGr:

          }
        case show.Walk:
          switch t {
          case 0:
            spc.Turn (right, -phi) // tilt
          case Shift:
            spc.Move (top, -delta0) // move
          }
        case show.Fly:
          spc.Turn (right, -phi) // tilt
        }
      case Enter:
        switch X.mode {
        case show.Look:
          if t == 0 {
            spc.Move (front, delta0) // move ahead
          } else {
            spc.Move (front, -delta0) // move back
          }
        case show.Walk:
          if t == 0 {
            spc.Move (front, delta0) // move ahead
          } else {
            spc.Move (front, -delta0) // move back
            // spc.Move (front, delta0) // move TODO translate focus ?
          }
        case show.Fly:
          // TODO increase speed
        }
      case Back:
        switch X.mode {
        case show.Look:
          spc.Move (front, -delta0) // move
        case show.Walk:
          if t == 0 {
            spc.Move (front, -delta0) // move
          } else {
            spc.Move (front, -delta0) // move TODO translate focus ?
          }
        case show.Fly:
          // TODO decrease speed
        }
      case Tab:
        spc.Invert() // pan 180°
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
        if X.mode == show.Walk {
          if t == 0 {
            spc.Invert() // pan 180°
          } else { // TODO
            dfe = X.eye.Distance (X.focus)
            spc.Move (front, 2 * dfe)
            spc.Invert()
          }
        }
      case F4:
        dfe = X.eye.Distance (X.focus)
        spc.Move (front, dfe)
        if t == 0 {
          spc.Move (right, dfe)
          spc.Turn (top, 90.)
        } else {
          spc.Move (right, -dfe)
          spc.Turn (top, -90.)
        }
      case Del: // TODO
        dfe = X.eye.Distance (X.focus)
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
