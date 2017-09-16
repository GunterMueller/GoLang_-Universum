package xwin

// (c) Christian Maurer   v. 170909 - license see murus.go

// #include <GL/glx.h>
/*
int etyp (XEvent *e) { return (*e).type; }

unsigned int kState (XEvent *e) { return (*e).xkey.state; }
unsigned int kCode (XEvent *e) { return (*e).xkey.keycode; }

int confWd (XEvent *e) { return (*e).xconfigure.width; }
int confHt (XEvent *e) { return (*e).xconfigure.height; }
*/
import
  "C"
import
  "murus/str"
var
  help = []string { "drehen / neigen: Pfeiltasten       ",
                    "         kippen: Strg-/Pfeiltasten ",
                    " näher / weiter: Eingabe-/Rücktaste",
                    "Licht von vorne: Tabulatortaste    ",
                    "        drucken: Drucktaste        ",
                    "        beenden: Abbruchtaste (Esc)" }

func (X *xwindow) WriteGlx() {
  C.glXSwapBuffers (dpy, C.GLXDrawable(X.win))
  X.win2buf()
}

func (X *xwindow) Start (x, y, z, x1, y1, z1 float64) {
  X.dx, X.dy, X.dz = C.GLfloat(x), C.GLfloat(y), C.GLfloat(y)
}

var
  glmode GLmode

func setMode (m GLmode) {
  glmode = m
}

const (
  shl = 50; shr = 62; ctll = 37; ctlr = 105; alt = 64; altgr = 108
  doofr = 135 //  doofl eaten by window manager
  esc = 9; enter = 36; back = 22
  left = 113; right = 114; up = 111; down = 116
  left1 = 166; right1 = 167; up1 = 112; down1 = 117
  pos1 = 110; end = 115
  ins = 118; del = 119
  tab = 23
  spc = 65
  hlp = 67; look = 68
  f3 = 69; f4 = 70; f5 = 71; f6 = 72; f7 = 73; f8 = 74; f9 = 75; f10 = 76
  f11 = 95; f12 = 96
  prt = 107; roll = 78; paus = 127
)

func (X *xwindow) Look (draw func()) {
  for i, s := range help { help[i] = str.Lat1 (s) }
  var xev C.XEvent
  delta := C.GLfloat(1.5)
  loop: for {
    for C.XPending (dpy) > 0 {
// println ("wait for event")
      C.XNextEvent (dpy, &xev)
      redraw := true
      eventtype := C.etyp (&xev)
      switch eventtype {
      case C.KeyPress:
// println ("keypress")
        code, state := uint(C.kCode (&xev)), uint(C.kState (&xev))
// println (code, state)
        switch code {
        case esc:
          break loop
        case enter:
          X.dz--
        case back:
          X.dz++
        case left:
          if glmode == Show {
            if state == 0 {
              X.C += delta
            } else {
              X.B -= delta
            }
          } else { // Walk
            if state == 0 {
              X.dx -= 0.1
            } else {
              X.dx -= 0.1
            }
          }
        case left1:

        case right:
          if glmode == Show {
            if state == 0 {
              X.C -= delta
            } else {
              X.B += delta
            }
          } else { // Walk
            if state == 0 {
              X.dx += 0.1
            } else {
              X.dx += 0.1
            }
          }
        case right1:

        case up:
          if glmode == Show {
            if state == 0 {
              X.A += delta
            } else {
              X.A += delta
            }
          } else { // Walk
            if state == 0 {
              X.dz += 0.1
            } else {
              X.dz += 0.1
            }
          }
        case up1:

        case down:
          if glmode == Show {
            if state == 0 {
              X.A -= delta
            } else {
              X.A -= delta
            }
          } else { // Walk
            if state == 0 {
              X.dz -= 0.1
            } else {
              X.dz -= 0.1
            }
          }
        }
        case down1:

      case C.ConfigureNotify:
println ("configure")
        C.glViewport (C.GLint(0), C.GLint(0), C.GLsizei(C.confWd(&xev)),
                                              C.GLsizei(C.confHt(&xev)))
      case C.FocusIn, C.FocusOut:
        redraw = false
      case C.Expose :
//         redraw = false
      case C.VisibilityNotify:
        redraw = false
      default:
        redraw = false
      }
      if redraw {
        C.glMatrixMode (C.GL_MODELVIEW)
        C.glLoadIdentity()
        C.glTranslatef (X.dx, X.dy, -X.dz)
        C.glRotatef (X.A, 1, 0, 0)
        C.glRotatef (X.B, 0, 1, 0)
        C.glRotatef (X.C, 0, 0, 1)
        draw()
        C.glXSwapBuffers (dpy, C.GLXDrawable(X.win))
      }
    }
  }
}


// #cgo LDFLAGS: -lX11 -lXext -lGL
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xatom.h>
// #include <GL/gl.h>
// #include <GL/glx.h>
