package xwin

// (c) Christian Maurer   v. 170816 - license see murus.go

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

func (X *xwindow) Start (a, b, c, dx, dy, dz float64) {
  X.A, X.B, X.C = C.GLfloat(a), C.GLfloat(b), C.GLfloat(c)
  X.dx, X.dy, X.dz = C.GLfloat(dx), C.GLfloat(dy), C.GLfloat(dz)
}

var
  glmode GLmode

func setMode (m GLmode) {
  glmode = m
}

func (X *xwindow) Look (draw func()) {
  for i, s := range help { help[i] = str.Lat1 (s) }
  var xev C.XEvent
  delta := C.GLfloat(1.5)
  loop: for {
    for C.XPending (dpy) > 0 {
      C.XNextEvent (dpy, &xev)
      redraw := true
      eventtype := C.etyp (&xev)
      switch eventtype {
      case C.KeyPress:
        code, state := uint(C.kCode (&xev)), uint(C.kState (&xev))
// println (code, state)
        switch code {
        case 9: // Escape
          break loop
        case 36: // Enter
          X.dz--
        case 22: // Back
          X.dz++
        case 113: // Left
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
        case 114: // Right
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
        case 111: // Up
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
        case 116: // Down
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
      case C.ConfigureNotify:
        C.glViewport (C.GLint(0), C.GLint(0), C.GLsizei(C.confWd(&xev)),
                                              C.GLsizei(C.confHt(&xev)))
      case C.FocusIn, C.FocusOut:
        redraw = false
      case C.Expose, C.VisibilityNotify:
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
