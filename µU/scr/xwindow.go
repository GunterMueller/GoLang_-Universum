package scr

// (c) Christian Maurer   v. 230223 - license see µU.go

// #cgo LDFLAGS: -lX11 -lXext -lGL -lGLU
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xlibint.h>
// #include <X11/Xutil.h>
// #include <X11/Xatom.h>
// #include <X11/cursorfont.h>
// #include <GL/gl.h>
// #include <GL/glx.h>
// #include <GL/glu.h>
/*
int evtyp (XEvent *e) { return (*e).type; }

void wait (Display *d, XEvent *e) {
  while (XCheckTypedEvent (d, Expose, e)) ;
}

void waitForLastContExpose (Display *d, XEvent *e) {
//  if XCheckMaskEvent (d, ExposureMask + StructureNotifyMask, e) == True {
  while (XCheckTypedWindowEvent (d, (*e).xexpose.window, Expose, e)) ;
}

void fullscreen (Display *d, Window w, Window w0, int on) {
  Bool only_if_it_exists = True;
  XEvent e;
  memset (&e, 0, sizeof (e)); // superfluous with "for ..."
  e.type = ClientMessage;
  e.xclient.display = d;
  e.xclient.window = w;
  e.xclient.message_type = XInternAtom (d, "_NET_WM_STATE", only_if_it_exists);
  e.xclient.format = 32;
  e.xclient.data.l[0] = on;
  e.xclient.data.l[1] = XInternAtom (d, "_NET_WM_STATE_FULLSCREEN", only_if_it_exists);
//  e.xclient.data.l[2] = 0; // only 1 state is going to be changed
  e.xclient.data.l[3] = 1; // souce indication
//  int i; for (i = 4; i <= 8; i++) { e.xclient.data.l[i] = 0; }
  XSendEvent (d, w0, False, SubstructureRedirectMask | SubstructureNotifyMask, &e);
  XFlush (d);
}

// void navi (Display *d, Window w, Atom a) {
//   XEvent e;
//   e.type = ClientMessage;
//   e.xclient.display = d;
//   e.xclient.window = w;
//   e.xclient.message_type = a;
//   e.xclient.send_event = False;
//   e.xclient.format = 16; // doesn't matter
//   if (XSendEvent (d, w, False, 0L, &e) < 0) ;
//   if (XSync (d, False) < 0) ;
// }

void getPixelColour (Display *d, XImage *i, Colormap m, int x, int y, int16_t *r, int16_t *g, int16_t *b) {
  XColor c;
  c.pixel = XGetPixel (i, x, y);
  XFree (i);
  XQueryColor (d, m, &c);
  *r = c.red;
  *g = c.green;
  *b = c.blue;
}

Window keyWin (XEvent *e) { return (*e).xkey.window; }

unsigned int keyState (XEvent *e) { return (*e).xkey.state; }

unsigned int keyCode (XEvent *e) { return (*e).xkey.keycode; }

Window buttonWin (XEvent *e) { return (*e).xbutton.window; }

unsigned int buttonState (XEvent *e) { return (*e).xbutton.state; }

unsigned int buttonButt (XEvent *e) { return (*e).xbutton.button; }

int buttonX (XEvent *e) { return (*e).xbutton.x; }

int buttonY (XEvent *e) { return (*e).xbutton.y; }

Window motionWin (XEvent *e) { return (*e).xmotion.window; }

unsigned int motionState (XEvent *e) { return (*e).xmotion.state; }

unsigned int motionH (XEvent *e) { return (*e).xmotion.is_hint; }

int motionX (XEvent *e) { return (*e).xmotion.x; }

int motionY (XEvent *e) { return (*e).xmotion.y; }

Window enterLeaveWin (XEvent *e) { return (*e).xcrossing.window; }

Window focusWin (XEvent *e) { return (*e).xfocus.window; }

Window exposeWin (XEvent *e) { return (*e).xexpose.window; }

unsigned int exposeX (XEvent *e) { return (*e).xexpose.x; }

unsigned int exposeY (XEvent *e) { return (*e).xexpose.y; }

unsigned int exposeWd (XEvent *e) { return (*e).xexpose.width; }

unsigned int exposeHt (XEvent *e) { return (*e).xexpose.height; }

unsigned int exposeC (XEvent *e) { return (*e).xexpose.count; }

Window visibilityWin (XEvent *e) { return (*e).xvisibility.window; }

Window mapWin (XEvent *e) { return (*e).xmap.window; }

Window unmapWin (XEvent *e) { return (*e).xunmap.window; }

Window configureWin (XEvent *e) { return (*e).xconfigure.window; }

int configureX (XEvent *e) { return (*e).xconfigure.x; }

int configureY (XEvent *e) { return (*e).xconfigure.y; }

int configureWd (XEvent *e) { return (*e).xconfigure.width; }

int configureHt (XEvent *e) { return (*e).xconfigure.height; }

Window resizeWin (XEvent *e) { return (*e).xresizerequest.window; }

unsigned int resizeWd (XEvent *e) { return (*e).xresizerequest.width; }

unsigned int resizeHt (XEvent *e) { return (*e).xresizerequest.height; }

Window circulateWin (XEvent *e) { return (*e).xcirculaterequest.window; }

Window mappingWin (XEvent *e) { return (*e).xmapping.window; }

Atom mT (XEvent *e) { return (*e).xclient.message_type; }

XKeyEvent keyEvent (XEvent *e) { return (*e).xkey; }

unsigned long xGetPixel (XImage *i, int x, int y) { return ((*((i)->f.get_pixel))((i), (x), (y))); }

void xPutPixel (XImage *i, int x, int y, unsigned long p) { ((*((i)->f.put_pixel))((i), (x), (y), (p))); }

void xDestroyImage (XImage *i) { ((*((i)->f.destroy_image))((i))); }

int etyp (XEvent *e) { return (*e).type; }

unsigned int kState (XEvent *e) { return (*e).xkey.state; }

unsigned int kCode (XEvent *e) { return (*e).xkey.keycode; }
*/
import
  "C"
import (
  "runtime"
  "unsafe"
  "strconv"
  "sync"
  "math"
  "µU/obj"
  "µU/ker"
  "µU/time"
  "µU/char"
  "µU/env"
  "µU/str"
  "µU/fontsize"
  "µU/font"
  "µU/col"
  "µU/mode"
  "µU/linewd"
  "µU/scr/shape"
//  "µU/navi"
  "µU/vect"
  "µU/gl"
  "µU/glu"
  "µU/spc"
)
const ( // see standards.freedesktop.org/wm-spec:
  _NET_WM_STATE_REMOVE = C.int(0) // remove/unset property
  _NET_WM_STATE_ADD    = C.int(1) // add/set property
  _NET_WM_STATE_TOGGLE = C.int(2) // toggle property
)
type
  xwindow struct {
             win C.Window // C.XID = C.ulong
                 mode.Mode
            x, y int
          wd, ht uint // window
      proportion float64
              gc C.GC
          buffer C.Pixmap
         shadows []C.Pixmap
            buff bool
          cF, cB,
        cFA, cBA,
      scrF, scrB col.Colour
          lineWd linewd.Linewidth
             fsp *C.XFontStruct
                 font.Font
                 fontsize.Size
   wd1, ht1, bl1 uint // font baseline
     transparent bool
         pointer uint
     cursorShape,
    consoleShape,
      blinkShape shape.Shape
  blinkX, blinkY uint
      blinkMutex sync.Mutex
        blinking bool
     firstExpose bool
         mouseOn bool
          xM, yM int
      subWindows []C.Window
          origin,
           focus,
             top vect.Vector
 polygonW, doneW [][]bool // to fill polygons
             swa C.XSetWindowAttributes // XXX
                 }
var (
  dspl string = env.Val ("DISPLAY")
  initializedW = false
  dpy *C.struct__XDisplay
  rootWin C.Window
  screen C.int
  monitorWd, monitorHt uint // full screen
  fullScreenW mode.Mode
  planes C.uint
/*/
  fdNavi uint
  naviAtom C.Atom
  navipipe chan navi.Command
/*/
  actualW *xwindow
  first bool = true // to start goSendEvents only once
  winList []*xwindow
  cutbuffer []string
  txt = []string {"", "",
                  "KeyPress", "KeyRelease", "ButtonPress", "ButtonRelease", "MotionNotify",
                  "EnterNotify", "LeaveNotify", "FocusIn", "FocusOut", "KeymapNotify",
                  "Expose", "GraphicsExpose", "NoExpose", "VisibilityNotify",
                  "CreateNotify", "DestroyNotify", "UnmapNotify", "MapNotify", "MapRequest",
                  "ReparentNotify", "ConfigureNotify", "ConfigureRequest", "GravityNotify",
                  "ResizeRequest", "CirculateNotify", "CirculateRequest",
                  "PropertyNotify", "SelectionClear", "SelectionRequest", "SelectionNotify",
                  "ColormapNotify", "ClientMessage", "MappingNotify", "GenericEvent", "LASTEvent"}
  startSendEvents = make(chan int)
  xMouse, yMouse int
)

func init() {
  cutbuffer = make([]string, 0)
  if env.UnderX() {
    go sendEvents()
  }
}

func initX() {
  if initializedW { return }
  if C.XInitThreads() == 0 { ker.Panic ("XKern.XInitThreads error") }
  d := C.CString(dspl); defer C.free (unsafe.Pointer(d))
  dpy = C.XOpenDisplay (d); if dpy == nil { ker.Panic ("dpy == nil") }
  screen = C.XDefaultScreen (dpy)
  monitorWd, monitorHt = uint(C.XDisplayWidth (dpy, screen)),
                         uint(C.XDisplayHeight (dpy, screen))
  fullScreenW = mode.ModeOf (uint(monitorWd), uint(monitorHt))
  planes = C.uint(C.XDefaultDepth (dpy, screen))
  initializedW = true
}

func cu (c col.Colour) C.ulong {
  return C.ulong(c.Code())
}

func (x *xwindow) Fin() {
  C.XFreePixmap (dpy, x.buffer)
  C.XUnmapWindow (dpy, x.win)
//  C.XFreeGC (dpy, x.gc)
//  C.XCloseDisplay (dpy) // XXX "BadDrawable (invalid .. Window parameter)"
  C.XDestroyWindow (dpy, x.win)
  initializedW = false
}

func err (e C.int) { // see /usr/incluode/X11/X.h
  switch e {
    case  1: ker.Panic ("BadRequeset")
    case  2: ker.Panic ("BadValue")
    case  3: ker.Panic ("BadWindow")
    case  4: ker.Panic ("BadPixmap")
    case  5: ker.Panic ("BadAtom")
    case  6: ker.Panic ("BadCursor")
    case  7: ker.Panic ("BadFont")
    case  8: ker.Panic ("BadMatch")
    case  9: ker.Panic ("BadDrawable")
    case 10: ker.Panic ("BadAccess")
    case 11: ker.Panic ("BadAlloc")
    case 12: ker.Panic ("BadColor")
    case 13: ker.Panic ("BadGC")
    case 14: ker.Panic ("BadIDChoice")
    case 15: ker.Panic ("BadName")
    case 16: ker.Panic ("BadLength")
    case 17: ker.Panic ("BadImplementation")
  }
}

func (X *xwindow) Flush() {
  C.XFlush (dpy)
}

func (X *xwindow) Name (n string) {
  s := C.CString (str.Lat1(n)); defer C.free(unsafe.Pointer(s))
  C.XStoreName (dpy, X.win, s)
}

func (X *xwindow) ActMode() mode.Mode {
  return X.Mode
}

func (w *xwindow) X() uint {
  return uint(w.x)
}

func (w *xwindow) Y() uint {
  return uint(w.y)
}

func (w *xwindow) Wd() uint {
  return w.wd
}

func (w *xwindow) Ht() uint {
  return w.ht
}

func (w *xwindow) Proportion() float64 {
  return float64(w.wd) / float64(w.ht)
}

////////////////////////////////////////////////////////////////////////

func (X *xwindow) ok() bool {
  return uint(X.x) + X.wd <= width && uint(X.y) + X.ht <= height
}

func (X *xwindow) OnFocus() bool {
  return actualW == X
}

func (X *xwindow) OffFocus() bool {
  return actualW != X
}

// colours /////////////////////////////////////////////////////////////

func (X *xwindow) ScrColours (f, b col.Colour) {
  X.scrF, X.scrB = f, b
}

func (X *xwindow) ScrColourF (f col.Colour) {
  X.scrF = f
}

func (X *xwindow) ScrColourB (b col.Colour) {
  X.scrB = b
}

func (X *xwindow) ScrCols() (col.Colour, col.Colour) {
  return X.scrF, X.scrB
}

func (X *xwindow) ScrColF() col.Colour {
  return X.scrF
}

func (X *xwindow) ScrColB() col.Colour {
  return X.scrB
}

func (X *xwindow) Colours (f, b col.Colour) {
  if ! initializedW { ker.Panic ("xwin.Colours: ! initializedW"); return }
  X.cF, X.cB = f, b
  X.ColourF (X.cF)
  X.ColourB (X.cB)
}

func (X *xwindow) ColourF (f col.Colour) {
  X.cF = f
  C.XSetForeground (dpy, X.gc, cu (X.cF))
}

func (X *xwindow) ColourB (b col.Colour) {
  X.cB = b
  C.XSetBackground (dpy, X.gc, cu (X.cB))
}

func (X *xwindow) Cols() (col.Colour, col.Colour) {
  return X.cF, X.cB
}

func (X *xwindow) ColF() col.Colour {
  return X.cF
}

func (X *xwindow) ColB() col.Colour {
  return X.cB
}

func (X *xwindow) Colour (x, y uint) col.Colour {
  var r, g, b C.int16_t
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x), C.int(y), 1, 1, C.AllPlanes, C.XYPixmap)
  C.getPixelColour (dpy, ximg, X.swa.colormap, C.int(0), C.int(0), &r, &g, &b)
  return col.New3 (byte(r), byte(g), byte(b))
}

// ranges //////////////////////////////////////////////////////////////

func (X *xwindow) clr (x, y int, w, h uint) {
  C.XSetForeground (dpy, X.gc, cu (X.scrB))
  C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XSetForeground (dpy, X.gc, cu (X.scrF))
  C.XFlush (dpy)
}

func (X *xwindow) Clr (l, c, w, h uint) {
  x, y := int(c) * int(X.wd1), int(l) * int(X.ht1)
  X.ClrGr (x, y, w * X.wd1, h * X.ht1)
}

func (X *xwindow) ClrGr (x, y int, w, h uint) {
  X.clr (x, y, w, h)
}

func (X *xwindow) Cls() {
  X.clr (0, 0, X.wd, X.ht)
}

func (X *xwindow) Buffered () bool {
  return X.buff
}

func (X *xwindow) Buf (on bool) {
  if X.buff == on { return }
  X.buff = on
  if on {
    C.XSetForeground (dpy, X.gc, cu (X.scrB))
    C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, 0, 0, C.uint(X.wd), C.uint(X.ht))
    C.XSetForeground (dpy, X.gc, cu (X.scrF))
    C.XFlush (dpy)
  } else {
    X.buf2win()
  }
}

func natord (x, y, x1, y1 *uint) {
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
}

func (X *xwindow) Save (l, c, w, h uint) {
  x, y := int(X.wd1) * int(c), int(X.ht1) * int(l)
  X.SaveGr (x, y, X.wd1 * w, X.ht1 * h)
}

func (X *xwindow) SaveGr (x, y int, w, h uint) {
  pixmap := C.XCreatePixmap (dpy, C.Drawable(X.win), C.uint(X.wd), C.uint(X.ht), planes)
  X.shadows = append (X.shadows, pixmap)
  n := len(X.shadows) - 1
  C.XCopyArea (dpy, C.Drawable(X.buffer), C.Drawable(X.shadows[n]), X.gc,
               C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
}

func (X *xwindow) Save1() {
  X.SaveGr (0, 0, X.wd, X.ht)
}

func (X *xwindow) Restore (l, c, w, h uint) {
  x, y := int(X.wd1) * int(c), int(X.ht1) * int(l)
  RestoreGr (x, y, X.wd1 * w, X.ht1 * h)
}

func (X *xwindow) RestoreGr (x, y int, w, h uint) {
  n := len(X.shadows)
  if n == 0 { return }
  n--
  C.XCopyArea (dpy, C.Drawable(X.shadows[n]), C.Drawable(X.win),
               X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer),
               X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
  C.XFreePixmap (dpy, X.shadows[n])
  X.shadows = X.shadows[:n]
}

func (X *xwindow) Restore1() {
  X.RestoreGr (0, 0, X.wd, X.ht)
}

// cursor //////////////////////////////////////////////////////////////

var
  finishedW bool

func (w *xwindow) blink() {
  var s shape.Shape
  for {
    w.blinkMutex.Lock()
    if w.cursorShape == shape.Off {
      s = w.blinkShape
    } else {
      s = shape.Off
    }
    w.cursor (w.blinkX, w.blinkY, s)
    w.blinkMutex.Unlock()
    if finishedW {
      break
    }
    time.Msleep (250)
  }
  runtime.Goexit()
}

func (w *xwindow) doBlink() {
  if w.blinking { return }
  w.blinking = true
  go w.blink()
}

func (w *xwindow) cursor (x, y uint, s shape.Shape) {
  y0, y1 := shape.Cursor (x, y, w.ht1, w.cursorShape, s)
  if y0 + y1 == 0 { return }
  w.cursorShape = s
//  Lock()
  w.RectangleFullInv (int(x), int(y + y0), int(x + w.wd1), int(y + y1))
  w.Flush()
//  Unlock()
}

func (w *xwindow) Warp (l, c uint, s shape.Shape) {
  w.WarpGr (w.wd1 * c, w.ht1 * l, s)
}

func (w *xwindow) WarpGr (x, y uint, s shape.Shape) {
  w.blinkMutex.Lock()
  w.blinkX, w.blinkY = x, y
  w.blinkShape = s
  w.cursor (x, y, w.blinkShape)
  w.blinkMutex.Unlock()
}

// text ////////////////////////////////////////////////////////////////

func (X *xwindow) Write1 (b byte, l, c uint) {
  X.Write (char.String(b), l, c)
}

func (X *xwindow) Write (s string, l, c uint) {
  X.WriteGr (s, int(c) * int(X.wd1), int(l) * int(X.ht1))
}

func (X *xwindow) WriteNat (n uint, l, c uint) {
  x, y := int(c * X.wd1), int(l * X.ht1)
  X.WriteNatGr (n, x, y)
}

func (X *xwindow) WriteNatGr (n uint, x, y int) {
  s := ""
  if n == 0 {
    s = "0"
  } else {
    for n > 0 {
      s = string(byte (n % 10 + '0')) + s
      n /= 10
    }
  }
  X.WriteGr (s, x, y)
}

func (X *xwindow) WriteInt (n int, l, c uint) {
  x, y := int(c * X.wd1), int(l * X.ht1)
  X.WriteIntGr (n, x, y)
}

func (X *xwindow) WriteIntGr (n, x, y int) {
  if n < 0 {
    n = -n
    Write1Gr ('-', x, y)
    x += int(X.wd1)
  }
  X.WriteNatGr (uint(n), x, y)
}

func (X *xwindow) Write1Gr (b byte, x, y int) {
  X.WriteGr (char.String(b), x, y)
}

func (X *xwindow) WriteGr (s string, x, y int) {
  C.XSetFont (dpy, X.gc, C.Font(X.fsp.fid)) // TODO TODO TODO
  n := C.uint(len (s))
  if ! X.transparent {
    C.XSetForeground (dpy, X.gc, cu (X.cB))
    if ! X.buff { C.XFillRectangle (dpy, C.Drawable(X.win), X.gc,
                                    C.int(x), C.int(y), n * C.uint(X.wd1), C.uint(X.ht1)) }
    C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc,
                      C.int(x), C.int(y), n * C.uint(X.wd1), C.uint(X.ht1))
    C.XSetForeground (dpy, X.gc, cu (X.cF))
  }
  cs := C.CString (s)
  defer C.free (unsafe.Pointer (cs))
  if ! X.buff { C.XDrawString (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y) + C.int(X.bl1), cs, C.int(n)) }
  C.XDrawString (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y) + C.int(X.bl1), cs, C.int(n))
  C.XFlush (dpy)
}

func (X *xwindow) Write1InvGr (b byte, x, y int) {
  X.WriteInvGr (char.String(b), x, y)
}

func (X *xwindow) WriteInvGr (s string, x, y int) {
  C.XSetFunction (dpy, X.gc, C.GXinvert)
  X.WriteGr (s, x, y)
  C.XSetFunction (dpy, X.gc, C.GXcopy)
}

func (X *xwindow) Transparent() bool {
  return X.transparent
}

func (X *xwindow) Transparence (t bool) {
  X.transparent = t
}

// font ////////////////////////////////////////////////////////////////

func (X *xwindow) ActFont() font.Font {
  return X.Font
}

func (X *xwindow) SetFont (f font.Font) {
 X.Font = f
}

func (X *xwindow) ActFontsize() fontsize.Size {
  return X.Size
}

func (X *xwindow) SetFontsize (s fontsize.Size) {
  X.Size = s
  name := "-xos4-terminus-bold"
  h := int(font.Ht (s))
  if s < fontsize.Normal {
    name = "-misc-fixed-medium"
  }
  name += "-r-*-*-" + strconv.Itoa(h) + "-*-*-*-*-*-*-*"
  f := C.CString (name); defer C.free (unsafe.Pointer(f))
  if dpy == nil { ker.Panic ("xwin.SetFontsize: dpy == nil") }
  X.fsp = C.XLoadQueryFont (dpy, f)
  if X.fsp == nil { ker.Panic ("terminus-bitmap-fonts are not installed !") }
  X.ht1 = uint(h)
  X.wd1, X.bl1 = uint(X.fsp.max_bounds.width), uint(X.fsp.max_bounds.ascent)
  if X.bl1 + uint(X.fsp.max_bounds.descent) != X.ht1 { ker.Panic ("xwin: font bl + d != ht") }
  C.XSetFont (dpy, X.gc, C.Font(X.fsp.fid))
}

func (X *xwindow) Wd1() uint {
  return uint(X.wd1)
}

func (X *xwindow) Ht1() uint {
  return X.ht1
}

func (X *xwindow) NLines() uint {
  return uint(X.ht / X.ht1)
}

func (X *xwindow) NColumns() uint {
  return uint(X.wd / X.wd1)
}

// graphics ////////////////////////////////////////////////////////////

func (X *xwindow) ActLinewidth() linewd.Linewidth {
  return X.lineWd
}

func (X *xwindow) SetLinewidth (w linewd.Linewidth) {
  X.lineWd = w
  cw := C.uint(0)
  switch w {
  case linewd.Thin:
    cw = C.uint(1)
  case linewd.Thick:
    cw = C.uint(2)
  case linewd.Thicker:
    cw = C.uint(3)
  case linewd.VeryThick:
    cw = C.uint(4)
  case linewd.Fat:
    cw = C.uint(5)
  }
  C.XSetLineAttributes (dpy, X.gc, cw, C.LineSolid, C.CapRound, C.JoinRound)
}

func (X *xwindow) point (x, y int, n bool) {
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff {
    C.XDrawPoint (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y))
  }
  C.XDrawPoint (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y))
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Point (x, y int) {
  X.point (x, y, true)
}

func near (x, y, a, b int, d uint) bool {
  dx, dy := x - a, y - b
  return dx * dx + dy * dy <= int(d * d)
}

func (X *xwindow) PointInv (x, y int) {
  X.point (x, y, false)
}

func (X *xwindow) OnPoint (x, y, a, b int, d uint) bool {
  return near (x, y, a, b, d)
}

func ok2W (xs, ys []int) bool {
  return len (xs) == len (ys)
}

func ok4W (xs, ys, xs1, ys1 []int) bool {
  return len (xs) == len (ys) &&
         len (xs1) == len (ys1) &&
         len (xs) == len (xs1)
}

func (X *xwindow) points (xs, ys []int, b bool) {
  n := len (xs)
  if n == 0 { return }
  if ! ok2W (xs, ys) { return }
  if n == 1 { X.point (xs[0], ys[0], b) }
  p := make ([]C.XPoint, n)
  for i := 0; i < n; i++ {
    p[i].x, p[i].y = C.short(xs[i]), C.short(ys[i])
  }
  if ! b { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawPoints (dpy, C.Drawable(X.win), X.gc, &p[0], C.int(n), C.CoordModeOrigin) }
  C.XDrawPoints (dpy, C.Drawable(X.buffer), X.gc, &p[0], C.int(n), C.CoordModeOrigin)
  if ! b { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Points (xs, ys []int) {
  X.points (xs, ys, true)
}

func (X *xwindow) PointsInv (xs, ys []int) {
  X.points (xs, ys, false)
}

func (X *xwindow) OnPoints (xs, ys []int, a, b int, d uint) bool {
  n := len(xs)
  if n == 0 { return false }
  for i := 0; i < n; i++ {
    if near (xs[i], ys[i], a, b, d) {
      return true
    }
  }
  return false
}

func (X *xwindow) line (x, y, x1, y1 int, n bool) {
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawLine (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.int(x1), C.int(y1)) }
  C.XDrawLine (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.int(x1), C.int(y1))
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Line (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, true)
}

func (X *xwindow) LineInv (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, false)
}

func (X *xwindow) OnLine (x, y, x1, y1, a, b int, t uint) bool {
  if x > x1 { x, x1 = x1, x; y, y1 = y1, y }
  if x == x1 {
    return between (x, x, a, int(t)) && between (y, y1, b, int(t))
  }
  if y == y1 {
    return between (y, y, b, int(t)) && between (x, x1, a, int(t))
  }
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  if near (x, y, a, b, t) || near (x1, y1, a, b, t) { return true }
  m := float64(y1 - y) / float64(x1 - x)
  return near (a, b, a, y + int(m * float64(a - x) + 0.5), t)
}

func (X *xwindow) lines (xs, ys, xs1, ys1 []int, n bool) {
  l := len (xs); if len (ys) != l { return }
  s := make ([]C.XSegment, l)
  for i := 0; i < l; i++ {
    s[i].x1, s[i].y1, s[i].x2, s[i].y2 = C.short(xs[i]), C.short(ys[i]), C.short(xs1[i]), C.short(ys1[i])
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawSegments (dpy, C.Drawable(X.win), X.gc, &s[0], C.int(l)) }
  C.XDrawSegments (dpy, C.Drawable(X.buffer), X.gc, &s[0], C.int(l))
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Lines (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, true)
}

func (X *xwindow) LinesInv (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, false)
}

func (X *xwindow) OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool {
  if len (xs) == 0 { return false }
  if ! ok4W (xs, ys, xs1, ys1) { return false }
  for i := 0; i < len (xs); i++ {
    if X.OnLine (xs[i], ys[i], xs1[i], ys1[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *xwindow) segments (xs, ys []int, n bool) {
  l := len (xs); if len (ys) != l { return }
  p := make ([]C.XPoint, l)
  for i := 0; i < l; i++ {
    p[i].x, p[i].y = C.short(xs[i]), C.short(ys[i])
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawLines (dpy, C.Drawable(X.win), X.gc, &p[0], C.int(l), C.CoordModeOrigin) }
  C.XDrawLines (dpy, C.Drawable(X.buffer), X.gc, &p[0], C.int(l), C.CoordModeOrigin)
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Segments (xs, ys []int) {
  X.segments (xs, ys, true)
}

func (X *xwindow) SegmentsInv (xs, ys []int) {
  X.segments (xs, ys, false)
}

func (X *xwindow) OnSegments (xs, ys []int, a, b int, t uint) bool {
  if ! ok2W (xs, ys) { return false }
  if len (xs) == 1 { return xs[0] == a && ys[0] == b }
  for i := 1; i < len (xs); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *xwindow) border (x, y, x1, y1 *int) {
  if *x > *x1 { *x, *x1 = *x1, *x; *y, *y1 = *y1, *y }
  for *x > 0 {
    *x -= *x1 - *x
    *y -= *y1 - *y
  }
  for *x1 < int(X.wd) {
    *x1 += *x1 - *x
    *y1 += *y1 - *y
  }
}

func (X *xwindow) InfLine (x, y, x1, y1 int) {
  if x == x1 {
    if y == y1 { return }
    X.Line (x, 0, x, int(X.ht) - 1)
    return
  }
  if y == y1 {
    X.Line (0, y, int(X.wd) - 1, y)
    return
  }
  X.border (&x, &y, &x1, &y1)
  X.Line (x, y, x1, y1)
}

func (X *xwindow) InfLineInv (x, y, x1, y1 int) {
  if x == x1 {
    if y == y1 { return }
    X.LineInv (x, 0, x, int(X.ht) - 1)
    return
  }
  if y == y1 {
    X.LineInv (0, y, int(X.wd) - 1, y)
    return
  }
  X.border (&x, &y, &x1, &y1)
  X.LineInv (x, y, x1, y1)
}

func (X *xwindow) OnInfLine (x, y, x1, y1, a, b int, t uint) bool {
  if x > x1 { x, x1 = x1, x; y, y1 = y1, y }
  if x == x1 {
    return between (x, x, a, int(t))
  }
  if y == y1 {
    return between (y, y, b, int(t))
  }
  if near (x, y, a, b, t) || near (x1, y1, a, b, t) { return true }
  X.border (&x, &y, &x1, &y1)
  m := float64(y1 - y) / float64(x1 - x)
  return near (a, b, a, y + int(m * float64(a - x) + 0.5), t)
}

func (X *xwindow) Triangle (x, y, x1, y1, x2, y2 int) {
  X.Polygon ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *xwindow) TriangleInv (x, y, x1, y1, x2, y2 int) {
  X.PolygonInv ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *xwindow) TriangleFull (x, y, x1, y1, x2, y2 int) {
  X.PolygonFull ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *xwindow) TriangleFullInv (x, y, x1, y1, x2, y2 int) {
  X.PolygonFullInv ([]int{x, x1, x2}, []int{y, y1, y2})
}

func (X *xwindow) rectangle (x, y, x1, y1 int, n, f bool) {
  w, h := x1 - x, y1 - y
  if f {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) } // C.GXcopyInverted ? 
    if ! X.buff { C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h)) }
    C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  } else {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
    if ! X.buff { C.XDrawRectangle (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h)) }
    C.XDrawRectangle (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Rectangle (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1, y1, true, false)
}

func (X *xwindow) RectangleInv (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1, y1, false, false)
}

func (X *xwindow) RectangleFull (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1, y1, true, true)
}

func (X *xwindow) RectangleFullInv (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1, y1, false, true)
}

func (X *xwindow) OnRectangle (x, y, x1, y1, a, b int, t uint) bool {
  if ! X.InRectangle (x, y, x1, y1, a, b, t) { return false }
  return between (x, x, a, int(t)) || between (x1, x1, a, int(t)) ||
         between (y, y, b, int(t)) || between (y1, y1, b, int(t))
}

func (X *xwindow) InRectangle (x, y, x1, y1, a, b int, t uint) bool {
  return between (x, x1, a, int(t)) && between (y, y1, b, int(t))
}

func (X *xwindow) CopyRectangle (x0, y0, x1, y1, x, y int) {
  intord (&x0, &y0, &x1, &y1)
  w, h := x1 - x0, y1 - y0
  if w == 0 || h == 0 { return }
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.win), X.gc, C.int(x0), C.int(y0),
               C.uint(w), C.uint(h), C.int(x), C.int(y))
  X.Flush()
}

func (X *xwindow) Polygon (xs, ys []int) {
  n := len(xs)
  if n < 2 { return }
  n--
  X.segments (xs, ys, true)
  X.line (xs[0], ys[0], xs[n], ys[n], true)
}

func (X *xwindow) PolygonInv (xs, ys []int) {
  n := len(xs)
  if n < 3 { return }
  n--
  X.segments (xs, ys, false)
  X.line (xs[0], ys[0], xs[n], ys[n], false)
}

func (X *xwindow) polygonFull (xs, ys []int, shape int, n bool) {
  l := len (xs); if len (ys) != l { return }
  p := make ([]C.XPoint, l)
  for i := 0; i < l; i++ {
    p[i].x, p[i].y = C.short(xs[i]), C.short(ys[i])
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopyInverted) }
  if ! X.buff { C.XFillPolygon (dpy, C.Drawable(X.win), X.gc, &p[0], C.int(l), C.int(shape), C.CoordModeOrigin) }
  C.XFillPolygon (dpy, C.Drawable(X.buffer), X.gc, &p[0], C.int(l), C.int(shape), C.CoordModeOrigin)
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) PolygonFull (xs, ys []int) {
  shape := C.Nonconvex
  if convex (xs, ys) {
    shape = C.Convex
  } else {
    shape = C.Complex
  }
  X.polygonFull (xs, ys, shape, true)
}

func (X *xwindow) PolygonFullInv (xs, ys []int) {
  shape := C.Nonconvex
  if convex (xs, ys) {
    shape = C.Convex
  }
  X.polygonFull (xs, ys, shape, false)
}

func (X *xwindow) PolygonFull1 (xs, ys []int, a, b int) {
  X.polygonFull (xs, ys, C.Nonconvex, true)
}

func (X *xwindow) PolygonFullInv1 (xs, ys []int, a, b int) {
  X.polygonFull (xs, ys, C.Nonconvex, false)
}

func (X *xwindow) OnPolygon (xs, ys []int, a, b int, t uint) bool {
  n := len (xs)
  if n == 0 { return false }
  if ! ok2W (xs, ys) { return false }
  if n == 1 { return xs[0] == a && ys[0] == b }
  for i := 1; i < int(n); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return X.OnLine (xs[n-1], ys[n-1], xs[0], ys[0], a, b, t)
}

func (X *xwindow) ellipse (x, y int, a, b uint, n, f bool) {
  x0, y0 := C.int(x) - C.int(a), C.int(y) - C.int(b)
  aa, bb := C.uint(2 * a), C.uint(2 * b)
  const a0 = C.int(0)
  if f {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) } // C.GXcopyInverted ?
    if ! X.buff { C.XFillArc (dpy, C.Drawable(X.win), X.gc, x0, y0, aa, bb, 0, 64 * 360) }
    C.XFillArc (dpy, C.Drawable(X.buffer), X.gc, C.int(x0), y0, aa, bb, 0, 64 * 360)
  } else {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
    if ! X.buff { C.XDrawArc (dpy, C.Drawable(X.win), X.gc, x0, y0, aa, bb, 0, 64 * 360) }
    C.XDrawArc (dpy, C.Drawable(X.buffer), X.gc, C.int(x0), y0, aa, bb, 0, 64 * 360)
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Circle (x, y int, r uint) {
  X.ellipse (x, y, r, r, true, false)
}

func (X *xwindow) CircleInv (x, y int, r uint) {
  X.ellipse (x, y, r, r, false, false)
}

func (X *xwindow) CircleFull (x, y int, r uint) {
  X.ellipse (x, y, r, r, true, true)
}

func (X *xwindow) CircleFullInv (x, y int, r uint) {
  X.ellipse (x, y, r, r, false, true)
}

func (X *xwindow) OnCircle (x, y int, r uint, a, b int, t uint) bool {
  d := uint(dist2 (x, y, a, b))
  if d >= r {
    return d <= t + r
  }
  return d + t > r
}

func (X *xwindow) InCircle (x, y int, r uint, a, b int, t uint) bool {
  return uint(dist2 (x, y, a, b)) <= r + t
}

func (X *xwindow) arc (x, y int, r uint, a, b float64, n, f bool) {
  for a >= 360 { a -= 360 }
  for a <= -360 { a += 360 }
  x0, y0 := C.int(x) - C.int(r), C.int(y) - C.int(r)
  rr, aa, bb := C.uint(2 * r), C.int(64 * a + 0.5), C.int(64 * b + 0.5)
  if f {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) } // C.GXcopyInverted ?
    if ! X.buff { C.XFillArc (dpy, C.Drawable(X.win), X.gc, x0, y0, rr, rr, aa, bb) }
    C.XFillArc (dpy, C.Drawable(X.buffer), X.gc, x0, y0, rr, rr, aa, bb)
  } else {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
    if ! X.buff { C.XDrawArc (dpy, C.Drawable(X.win), X.gc, x0, y0, rr, rr, aa, bb) }
    C.XDrawArc (dpy, C.Drawable(X.buffer), X.gc, x0, y0, rr, rr, aa, bb)
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Arc (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, true, false)
}

func (X *xwindow) ArcInv (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, false, false)
}

func (X *xwindow) ArcFull (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, true, true)
}

func (X *xwindow) ArcFullInv (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, false, true)
}

func (X *xwindow) Ellipse (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, true, false)
}

func (X *xwindow) EllipseInv (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, false, false)
}

func (X *xwindow) EllipseFull (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, true, true)
}

func (X *xwindow) EllipseFullInv (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, false, true)
}

func dist2 (x, y, x1, y1 int) int {
  return int((math.Sqrt(float64((x1 - x) * (x1 - x) + (y1 - y) * (y1 - y))) + 0.5))
}

// work around Bresenham ellipse
func (X *xwindow) OnEllipse (x, y int, a, b uint, A, B int, t uint) bool {
  e := int(math.Sqrt(math.Abs(float64(a * a) - float64(b * b))) + 0.5)
  r := 2 * int(a); z:= 2 * dist2 (x, y, A, B) // if a == b
  if a > b {
    z = dist2 (x - e, y, A, B) + dist2 (x + e, y, A, B)
  }
  if a < b {
    z = dist2 (x, y - e, A, B) + dist2 (x, y + e, A, B)
    r = 2 * int(b)
  }
  return between (r, r, z, int(t))
}

func (X *xwindow) InEllipse (x, y int, a, b uint, A, B int, t uint) bool {
  return inEllipse (x, y, a, b, A, B, t)
}

func (X *xwindow) Curve (xs, ys []int) {
  var xs1, ys1 []int
  curve (xs, ys, &xs1, &ys1)
  X.Points (xs1, ys1)
}

func (X *xwindow) CurveInv (xs, ys []int) {
  var xs1, ys1 []int
  curve (xs, ys, &xs1, &ys1)
  X.PointsInv (xs1, ys1)
}

func (X *xwindow) OnCurve (xs, ys []int, a, b int, t uint) bool {
  var xs1, ys1 []int
  curve (xs, ys, &xs1, &ys1)
  n := len(xs1)
  for i := 0; i < n; i++ {
    if near (xs1[i], ys1[i], a, b, t) {
      return true
    }
  }
  return false
}

// mouse ///////////////////////////////////////////////////////////////

/*/
func (X *xwindow) clear() { // XXX
  C.XSetForeground (dpy, X.gc, cu (X.scrB))
  C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, 0, 0, C.uint(X.wd), C.uint(X.ht))
  C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, 0, 0, C.uint(X.wd), C.uint(X.ht))
  C.XFillRectangle (dpy, C.Drawable(X.shadow1), X.gc, 0, 0, C.uint(X.wd), C.uint(X.ht))
  C.XSetForeground (dpy, X.gc, cu (X.scrF))
}
/*/

func (X *xwindow) MousePos() (uint, uint) {
  X.xM, X.yM = xMouse, yMouse
  return uint(X.yM) / uint(X.ht1), uint(X.xM) / uint(X.wd1)
}

func (X *xwindow) MousePosGr() (int, int) {
  X.xM, X.yM = xMouse, yMouse
  return X.xM, X.yM
}

func (X *xwindow) WriteMousePos (l, c uint) {
  x, y := int(l * X.wd1), int(c * X.ht1)
  X.WriteMousePosGr (x, y)
}

func (X *xwindow) WriteMousePosGr (x, y int) {
  xm, ym := X.MousePos()
  X.Colours (X.ScrCols())
  X.WriteGr ("        ", x, y)
  X.WriteNatGr (uint(xm), x, y)
  X.WriteNatGr (uint(ym), x + 5 * int(X.wd1), y)
}

func (X *xwindow) WarpMouse (l, c uint) {
  X.WarpMouseGr (int(c) * int(X.wd1), int(l) * int(X.ht1))
}

func (X *xwindow) WarpMouseGr (x, y int) {
  C.XWarpPointer (dpy, C.None, X.win, 0, 0, 0, 0, C.int(x), C.int(y))
  C.XFlush (dpy)
}

func (X *xwindow) SetPointer (p uint) {
  switch p {
  case Crosshair, Gumby, Standard:
    X.pointer = p
    C.XDefineCursor (dpy, X.win, C.XCreateFontCursor (dpy, C.uint(p)))
  }
}

func (X *xwindow) MousePointer (on bool) {
  X.mouseOn = on
  X.WriteMousePointer()
}

func (X *xwindow) mousePointerOff() {
  var c C.XColor
  c.red, c.green, c.blue = C.ushort(0), C.ushort(0), C.ushort(0)
  s := C.CString(string(obj.Stream {0, 8, 0, 0, 0, 0})); defer C.free (unsafe.Pointer(s))
  m := C.XCreateBitmapFromData (dpy, C.Drawable(X.win), s, C.uint(8), C.uint(8));
  cursor := C.XCreatePixmapCursor (dpy, m, m, &c, &c, C.uint(0), C.uint(0))
  C.XDefineCursor (dpy, X.win, cursor)
  C.XFreeCursor (dpy, cursor)
  C.XFreePixmap (dpy, m)
}

func (X *xwindow) MousePointerOn() bool {
  return X.mouseOn
}

func (X *xwindow) WriteMousePointer() {
  if X.mouseOn {
    X.SetPointer (X.pointer)
  } else {
    X.mousePointerOff()
  }
}

func (X *xwindow) UnderMouse (l, c, w, h uint) bool {
  xm, ym := X.MousePos()
  return l <= xm && xm < l + h && c <= ym && ym < c + w
}

func (X *xwindow) UnderMouseGr (x, y, x1, y1 int, t uint) bool {
  intord (&x, &y, &x1, &y1)
  xm, ym := X.MousePosGr()
  return x <= xm + int(t) && xm <= x1 + int(t) && y <= ym + int(t) && ym <= y1 + int(t)
}

func (X *xwindow) UnderMouse1 (x, y int, d uint) bool {
  xm, ym := X.MousePosGr()
  return (x - xm) * (x - xm) + (y - ym) * (y - ym) <= int(d * d)
}

// serialisation ///////////////////////////////////////////////////////

func (X *xwindow) Codelen (w, h uint) uint {
  return 16 + 3 * w * h
}

func (X *xwindow) Encode (x, y int, w, h uint) obj.Stream {
  if w == 0 || h == 0 { ker.Panic ("xwin.Encode: w == 0 or h == 0") }
  if w > uint(X.wd) { ker.Panic ("xwin.Encode: w > X.wd") }
  if h > uint(X.ht) { ker.Panic ("xwin.Encode: h > X.ht") }
  s := make (obj.Stream, X.Codelen (w, h))
  n := 16
  copy (s[:n], obj.Encode2 (int(w), int(h)))
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x), C.int(y),
                       C.uint(w), C.uint(h), C.AllPlanes, C.XYPixmap)
  var pixel C.ulong
  for j := 0; j < int(h); j++ {
    for i := 0; i < int(w); i++ {
      pixel = C.xGetPixel (ximg, C.int(i), C.int(j))
      e := obj.Encode(uint32(pixel))
      copy (s[n:n+3], e)
      n += 3
    }
  }
  C.xDestroyImage (ximg)
  return s
}

func (X *xwindow) Decode (s obj.Stream, x, y int) {
  if s == nil { return }
  n := 16
  w, h := obj.Decode2 (s[:n])
  ximg := C.XGetImage (dpy, C.Drawable(X.win), C.int(x), C.int(y),
                       C.uint(w), C.uint(h), C.AllPlanes, C.XYPixmap)
  var pixel C.ulong
  c := col.New()
  for j := 0; j < int(h); j++ {
    for i := 0; i < int(w); i++ {
      c.Set (s[n+2], s[n+1], s[n+0])
      pixel = (C.ulong)(c.Code())
      C.xPutPixel (ximg, C.int(i), C.int(j), pixel)
      n += 3
    }
  }
  C.XPutImage (dpy, C.Drawable(X.win), X.gc, ximg,
               0, 0, C.int(x), C.int(y), C.uint(w), C.uint(h))
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc,
               C.int(x), C.int(y), C.uint(w), C.uint(h), C.int(x), C.int(y))
  C.XFlush (dpy)
}

// image-operations ////////////////////////////////////////////////////

func (X *xwindow) WriteImage (c [][]col.Colour, x, y int) {
  w, h := len(c[0]), len(c)
  for i := 0; i < w; i++ {
    for j := 0; j < h; j++ {
      if x + i < int(X.Wd()) && y + j < int(X.Ht()) {
        X.ColourF (c[j][i])
        X.Point (x + i, y + j)
      }
    }
  }
}

func (X *xwindow) Screenshot (x, y int, w, h uint) obj.Stream {
  _ = C.XGetImage (dpy, C.Drawable(X.win),
                   0, 0, C.uint(w), C.uint(h), C.AllPlanes, C.XYPixmap)
  return X.Encode (x, y, w, h)
}

////////////////////////////////////////////////////////////////////////

const (
  um = math.Pi / 180
  epsilon = 1e-6
  Esc = 9; Enter = 36; Back = 22
  Left = 113; Right = 114; Up = 111; Down = 116
  F1 = 67; F9 = 75; F10 = 76; F11 = 95; F12 = 96
  Shift = 1; Strg = 4; Alt = 8; AltGr = 128
)
type
  d = C.GLdouble
/*/
var
  help = []string{ // TODO
// 0         1         2         3         4         5         6         7
// 01234567890123456789012345678901234567890123456789012345678901234567890123456789
  "                                                                                ",
  "                                                                                ",
  "                                                                                ")
                 }
/*/

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
//  gl.ShowLight (true) // XXX
  spc.Set (ox, oy, oz, fx, fy, fz, tx, ty, tz)
  phi, delta := 1.0, 0.1
  var xev C.XEvent
  redraw := true
//  if m == Fly { go X.fly() }
  for {
    if redraw {
      gl.Clear()
      ex, ey, ez := spc.GetOrigin()
      fx, fy, fz := spc.GetFocus()
      nx, ny, nz := spc.GetTop()
      if math.Abs(fx) < epsilon { fx = 0 }; if math.Abs(fy) < epsilon { fy = 0 }
      gl.MatrixMode (gl.Projection)
      gl.LoadIdentity()
      gl.Viewport (0, 0, X.wd, X.ht)
      glu.Perspective (60, X.proportion, 0.1, 10000.)
      C.gluLookAt (d(ex), d(ey), d(ez), d(fx), d(fy), d(fz), d(nx), d(ny), d(nz))
      draw()
      C.glXSwapBuffers (dpy, C.GLXDrawable(X.win))
      C.glFinish()
    }
    redraw = true
    C.XNextEvent (dpy, &xev)
    et := C.etyp (&xev)
    switch et {
    case C.KeyPress:
      c, t := C.kCode (&xev), C.kState (&xev)
      depth := uint(0)
      switch t {
      case 0:
        depth = 0
      case Shift:
        depth = 1
      case Strg, Shift + Strg:
        depth = 1
      case Alt, AltGr:
        depth = 2
      case Shift + Alt, Strg + Alt, Shift + Strg + Alt,
           Shift + AltGr, Strg + AltGr, Shift + Strg + AltGr:
        depth = 3
      }
      switch c {
      case Esc:
        return
      case Enter:
        switch m {
        case Look:
          if depth == 0 {
            spc.MoveFront (delta)
          } else {
            spc.MoveFront (-delta)
          }
        case Walk:
          if depth == 0 {
            spc.MoveFront (delta)
          } else {
            spc.MoveFront (-delta)
          }
        case Fly:
          // TODO increase speed
        }
      case Back:
        switch m {
        case Look:
          if depth == 0 {
            spc.Turn (180)
          } else {
            d := X.origin.Distance (X.focus)
            spc.MoveFront (2 * d)
            spc.Turn (180)
          }
        case Walk:
          if depth == 0 {
            spc.Turn (180)
          } else {
            spc.MoveFront (delta)
          }
        case Fly:
          // TODO decrease speed
        }
      case Left:
        switch m {
        case Look:
          switch depth {
          case 0:
            spc.Turn (-phi)
          case 1:
            spc.MoveRight (delta)
          case 2:
            spc.TurnAroundFocusT (phi)
          case 3:
            spc.Roll (phi)
          }
        case Walk:
          switch depth {
          case 0:
            spc.Turn (phi)
          case 1:
            spc.MoveRight (-phi)
          case 2:
            spc.Roll (phi)
          }
        case Fly:
          if depth == 0 {
            spc.Turn (phi)
          } else {
            spc.Roll (phi)
          }
        }
      case Right:
        switch m {
        case Look:
          switch depth {
          case 0:
            spc.Turn (phi)
          case 1:
            spc.MoveRight (-delta)
          case 2:
            spc.TurnAroundFocusT (-phi)
          case 3:
            spc.Roll (-phi)
          }
        case Walk:
          switch depth {
          case 0:
            spc.Turn (-phi)
          case 1:
            spc.MoveRight (phi)
          case 2:
            spc.Roll (-phi)
          }
        case Fly:
          if depth == 0 {
            spc.Turn (-phi)
          } else {
            spc.Roll (-phi)
          }
        }
      case Up:
        switch m {
        case Look:
          switch depth {
          case 0:
            spc.Tilt (phi)
          case 1:
            spc.MoveTop (delta)
          case 2:
            spc.TurnAroundFocusR (phi)
          case 3:
            spc.Tilt (phi)
          }
        case Walk:
          switch depth {
          case 0:
            spc.Tilt (phi)
          case 1:
            spc.MoveTop (phi)
          }
        case Fly:
          spc.Tilt (phi)
        }
      case Down:
        switch m {
        case Look:
        switch depth {
          case 0:
            spc.Tilt (-phi)
          case 1:
            spc.MoveTop (-phi / 10)
          case 2:
            spc.TurnAroundFocusR (-phi)
          case 3:
            spc.Tilt (-phi)
          }
        case Walk:
          switch depth {
          case 0:
            spc.Tilt (-phi)
          case 1:
            spc.MoveTop (-phi)
          }
        case Fly:
          spc.Tilt (-phi)
        }
      case F1:
        // TODO help
//      case F9:
//        spc.SetLight (0)
      default:
        redraw = false
      }
    default:
      redraw = false
    }
  }
}

////////////////////////////////////////////////////////////////////////

func NewW (x, y uint, m mode.Mode) Screen {
  return NewWHW (x, y, mode.Wd (m), mode.Ht (m))
}

func NewWHW (x, y, w, h uint) Screen {
  initX()
  X := new(xwindow)
  X.Mode = mode.ModeOf (w, h)
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = w, h
  if X.wd > monitorWd || X.ht > monitorHt { ker.Panic ("win too large: " +
                         strconv.Itoa(int(X.wd)) + " x " + strconv.Itoa(int(X.ht))) }
  X.polygonW = make([][]bool, X.wd)
  X.doneW = make([][]bool, X.wd)
  for i := 0; i < int(X.wd); i++ {
    X.polygonW[i] = make([]bool, X.ht)
    X.doneW[i] = make([]bool, X.ht)
  }
  X.proportion = float64(X.wd)/float64(X.ht)
  X.Transparence (false)
  X.lineWd = linewd.Thin
  X.pointer = Standard
  X.scrF, X.scrB = col.StartCols()
  X.cF, X.cB = col.StartCols()
  X.cFA, X.cBA = col.StartColsA()
  a := [11]C.int{C.GLX_RGBA, C.GLX_DOUBLEBUFFER, C.GLX_DEPTH_SIZE, 16,
                 C.GLX_RED_SIZE, 1, C.GLX_GREEN_SIZE, 1, C.GLX_BLUE_SIZE, 1, C.None}
  vi := C.glXChooseVisual (dpy, screen, &a[0] )
  cx := C.glXCreateContext (dpy, vi, C.GLXContext(nil), C.True)
  rootWin = C.XRootWindow (dpy, vi.screen)
  X.swa.colormap = C.XCreateColormap (dpy, rootWin, vi.visual, C.AllocNone)
  X.swa.border_pixel = 8
  X.swa.event_mask = C.KeyPressMask | C.KeyReleaseMask | C.ButtonPressMask | C.ButtonReleaseMask |
                   C.FocusChangeMask | C.ExposureMask | C.VisibilityChangeMask |
                   C.StructureNotifyMask
//  X.swa.override_redirect = C.True
  X.win = C.XCreateWindow (dpy, rootWin, C.int(x), C.int(y), C.uint(X.wd), C.uint(X.ht), 0,
                           vi.depth, C.InputOutput, vi.visual,
                           C.CWBorderPixel | C.CWColormap | C.CWEventMask, &X.swa)
  X.gc = C.XDefaultGC (dpy, screen)
  C.XSetGraphicsExposures (dpy, X.gc, C.False)
  s := C.CString(env.Call()); defer C.free(unsafe.Pointer(s))
  var sh C.XSizeHints
  sh.flags = C.PPosition | C.PSize | C.PMinSize | C.PMaxSize
  sh.x, sh.y, sh.width, sh.height = C.int(x), C.int(y), C.int(X.wd), C.int(X.ht)
  sh.min_width, sh.min_height = sh.width, sh.height
  sh.max_width, sh.max_height = sh.width, sh.height
  C.XSetStandardProperties (dpy, X.win, s, s, C.None, &s, 1, &sh)
//  C.glXMakeContextCurrent (dpy, C.GLXDrawable(X.win), C.GLXDrawable(X.win), cx)
  C.glXMakeCurrent (dpy, C.GLXDrawable(X.win), cx)
  X.buffer = C.XCreatePixmap (dpy, C.Drawable(X.win), C.uint(X.wd), C.uint(X.ht), planes)
  const mask = C.KeyPressMask + // C.KeyReleaseMask +
               C.ButtonPressMask + C.ButtonReleaseMask +
               C.EnterWindowMask + C.LeaveWindowMask +
               C.PointerMotionMask + // C.PointerMotionHintMask +
//               C.Button1MotionMask + C.Button2MotionMask + C.Button3MotionMask +
//               C.Button4MotionMask + C.Button5MotionMask + C.ButtonMotionMask +
               C.KeymapStateMask +
               C.ExposureMask +
               C.VisibilityChangeMask +
               C.StructureNotifyMask +
               C.ResizeRedirectMask +
               C.SubstructureNotifyMask + C.SubstructureRedirectMask +
               C.FocusChangeMask +
               C.PropertyChangeMask // +
//               C.ColormapChangeMask + C.OwnerGrabButtonMask
  C.XSelectInput (dpy, X.win, mask)
  X.MousePointer (true)
  if X.wd == monitorWd && X.ht == monitorHt {
    C.fullscreen (dpy, X.win, rootWin, _NET_WM_STATE_ADD)
  }
  winList = append (winList, X)
  defer X.doBlink()
//  X.firstExpose = true
  if first {
    first = false
/*/
    p := C.CString ("WM_PROTOCOLS"); defer C.free (unsafe.Pointer(p))
    wm_protocols := C.XInternAtom (dpy, p, C.False)
    C.XSetWMProtocols (dpy, X.win, &wm_protocols, 1)
/*/
/*/
    m := C.CString ("navi"); defer C.free (unsafe.Pointer(m))
    naviAtom = C.XInternAtom (dpy, m, C.False)
    if navi.Exists() {
      navipipe = navi.Channel()
      go X.catchNavi()
    }
/*/
    startSendEvents <- 0
  }
  C.XMapWindow (dpy, X.win)
  X.Name (env.Arg(0))
/*/
  var ev C.XEvent
  C.XNextEvent (dpy, &ev)
  et := C.evtyp (&ev)
  switch et {
  case C.Expose, C.ConfigureNotify: // zur Erstausgabe
    for C.XCheckTypedEvent (dpy, C.int(et), &ev) == C.True { }
  case C.KeyPress, C.KeyRelease, C.ButtonPress, C.ButtonRelease, C.MotionNotify:
    C.XPutBackEvent (dpy, &ev)
  case C.ReparentNotify: // at Switch (?)
    // ignore
  default: // for test purposes
    println ("at initializing xwin:" + txt[et])
  }
  X.clear()
/*/
  X.origin, X.focus, X.top = vect.New(), vect.New(), vect.New()
  X.SetLinewidth (linewd.Thin)
  mr, _ := MaxResW()
  if mr >= mode.Wd (mode.UHD) {
    X.SetFontsize (fontsize.Huge)
  } else {
    X.SetFontsize (fontsize.Normal)
  }
  X.Colours (col.StartCols())
  X.ScrColours (col.StartCols())
  X.Cls()
  actualW = X
  return X
}

func NewMaxW() Screen {
  return NewW (0, 0, mode.ModeOf (MaxResW()))
}

func MaxModeW() mode.Mode {
  initX()
  return fullScreenW
}

func MaxResW() (uint, uint) {
  initX()
  return uint(monitorWd), uint(monitorHt)
}

func OkW (m mode.Mode) bool {
  fullScreen = MaxModeW()
  return mode.Wd (m) <= mode.Wd (fullScreenW) &&
         mode.Ht (m) <= mode.Ht (fullScreenW)
}

func actW() Screen {
  return actualW
}

// W == imp (w) iff w == W.win
func imp (w C.Window) *xwindow {
  for _, x := range winList {
    if x.win == w {
      return x
    }
  }
  ker.Panic ("µU/xwin.imp: there is no xwindow there")
  return nil
}

/*/
func Win (i uint) *xwindow {
  if int(i) < len(winList) {
    return winList[i]
  }
  ker.Panic ("µU/xwin.Win: there is no xwindow there with number" + strconv.Itoa(int(i)))
}
/*/

func (X *xwindow) Subwindow (x, y int, w, h uint) {
  s := C.XCreateSimpleWindow (dpy, X.win, C.int(x), C.int(y), C.uint(w), C.uint(h),
                              0, cu (X.scrF), cu (X.scrB))
  X.subWindows = append (X.subWindows, s)
}

func (X *xwindow) Sub (n uint) C.Window {
  if int(n) < len (X.subWindows) {
    return X.subWindows[n]
  }
  return C.Window(0)
}

// func (X *xwindow) catchNavi() {
//   for {
//     C.navi (dpy, X.win, naviAtom)
//   }
// }

func (X *xwindow) Win2buf() {
  X.win2buf()
}

func (X *xwindow) win2buf() {
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer),
               X.gc, 0, 0, C.uint(X.wd), C.uint(X.ht), 0, 0)
  C.XFlush (dpy)
}

func (X *xwindow) buf2win() {
  C.XCopyArea (dpy, C.Drawable(X.buffer), C.Drawable(X.win),
               X.gc, 0, 0, C.uint(X.wd), C.uint(X.ht), 0, 0)
  C.XFlush (dpy)
}

type
  ComposeStatus struct {
           compose_ptr C.XPointer
         chars_matched C.int
                       }

func sendEvents() {
  var (
    xev C.XEvent
    eventtype C.int
    w C.Window
    W *xwindow
    event Event
  )
  <-startSendEvents
  startSendEvents = nil
  time.Msleep(1e2)
  for {
    for C.XPending (dpy) > 0 {
      C.XNextEvent (dpy, &xev)
      event.C, event.S = 0, 0
      eventtype = C.evtyp (&xev)
      event.T = uint(eventtype)
// println (txt[eventtype])
      switch eventtype {
      case C.KeyPress:
        w = C.keyWin (&xev)
        W = imp (w)
        event.C, event.S = uint(C.keyCode (&xev)), uint(C.keyState (&xev))
      case C.KeyRelease:
        w = C.keyWin (&xev)
        W = imp (w)
        event.C, event.S = uint(C.keyCode (&xev)), uint(C.keyState (&xev))
      case C.ButtonPress:
        w = C.buttonWin (&xev)
        W = imp (w) // w == W.win
        event.C, event.S = uint(C.buttonButt (&xev)), uint(C.buttonState (&xev))
        W.xM, W.yM = int(C.buttonX (&xev)), int(C.buttonY (&xev))
      case C.ButtonRelease:
        w = C.buttonWin (&xev)
        W = imp (w) // w == W.win
        event.C, event.S = uint(C.buttonButt (&xev)), uint(C.buttonState (&xev))
        W.xM, W.yM = int(C.buttonX (&xev)), int(C.buttonY (&xev))
      case C.MotionNotify:
        w = C.motionWin (&xev)
        W = imp (w) // w == W.win
        event.C, event.S = uint(0), uint(C.motionState (&xev))
        W.xM, W.yM = int(C.motionX (&xev)), int(C.motionY (&xev))
        xMouse, yMouse = W.xM, W.yM
      case C.EnterNotify, C.LeaveNotify:
        w = C.enterLeaveWin (&xev)
        W = imp (w) // w == W.win
      case C.FocusIn:
        w = C.buttonWin (&xev)
        W = imp (w) // w == W.win
        actualW = W
      case C.FocusOut:
        w = C.buttonWin (&xev)
        W = imp (w) // w == W.win
      case C.KeymapNotify:
        ;
      case C.Expose:
        w = C.exposeWin (&xev)
        W = imp(w)
        if W.firstExpose {
          W.firstExpose = false
          C.waitForLastContExpose (dpy, &xev)
          C.wait (dpy, &xev)
        }
        W.buf2win()
/*/
        event.C, event.S = uint(C.keyCode (&xev)), uint(C.keyState (&xev))
        Eventpipe <- event
/*/
      case C.GraphicsExpose:
        ;
      case C.NoExpose:
        ;
      case C.VisibilityNotify:
        w = C.visibilityWin (&xev)
        W = imp (w) // w == W.win
        W.buf2win()
      case C.CreateNotify:
        ;
      case C.DestroyNotify:
        ;
      case C.UnmapNotify:
        w = C.unmapWin (&xev)
        W = imp (w) // w == W.win
        W.win2buf()
      case C.MapNotify:
        w = C.mapWin (&xev)
        W = imp (w) // w == W.win
        W.buf2win()
      case C.MapRequest:
        ;
      case C.ReparentNotify:
        ;
      case C.ConfigureNotify:
        C.glViewport (C.GLint(0), C.GLint(0),
                      C.GLsizei(C.configureHt(&xev)), C.GLsizei(C.configureWd(&xev)))
      case C.ConfigureRequest:
        ;
      case C.GravityNotify:
        ;
      case C.ResizeRequest: // ignored !
/*/
        w = C.resizeWin (&xev)
        W = imp (w)
        W.buf2win()
/*/
      case C.CirculateNotify:
        ;
      case C.CirculateRequest:
        ;
      case C.PropertyNotify:
        ;
      case C.SelectionClear:
        ;
      case C.SelectionRequest:
        ;
      case C.SelectionNotify:
        ;
      case C.ColormapNotify:
        ;
/*/
      case C.ClientMessage:
        mT := C.mT (&xev)
        if mT != naviAtom { println ("unknown xclient.message_type ", uint32(mT)) }
/*/
      case C.MappingNotify:
        ;
      case C.GenericEvent:
        ;
      case C.LASTEvent:
        ;
      }
      switch eventtype {
      case C.KeyRelease: // ignore
      case C.KeyPress:
        Eventpipe <- event
      case C.ButtonPress, C.ButtonRelease, C.MotionNotify:
        Eventpipe <- event
      default:
//        println (txt[eventtype])
      }
    }
  }
}
