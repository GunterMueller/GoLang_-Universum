package xwin

// (c) Christian Maurer   v. 180804 - license see µU.go

// #cgo LDFLAGS: -lX11 -lXext -lGL
// #include <stdio.h>
// #include <stdlib.h>
// #include <string.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xatom.h>
// #include <GL/glx.h>
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

void navi (Display *d, Window w, Atom a) {
  XEvent e;
  e.type = ClientMessage;
  e.xclient.display = d;
  e.xclient.window = w;
  e.xclient.message_type = a;
  e.xclient.send_event = False;
  e.xclient.format = 16; // doesn't matter
  if (XSendEvent (d, w, False, 0L, &e) < 0) ;
  if (XSync (d, False) < 0) ;
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

int lookupString (XEvent e) {
  XKeyEvent kev = e.xkey;
  char buffer[10];
  KeySym key;
  XComposeStatus cs;
  int g;
  g = XLookupString (&kev, buffer, 20, &key, &cs);
//  printf("key %d %  %s\n", key, buffer, &cs.compose_ptr);
  printf("key %d %s %s\n", key, buffer, ";");
  return g;
}
*/
import
  "C"
//  char *buffer;
import (
  "unsafe"
  "strconv"
  "sync"
  "time"
  . "µU/shape"
  "µU/ptr"
  ."µU/linewd"
  "µU/env"
  "µU/str"
  "µU/mode"
  "µU/col"
  "µU/font"
  "µU/navi"
)
const ( // standards.freedesktop.org/wm-spec:
  _NET_WM_STATE_REMOVE = C.int(0) // remove/unset property
  _NET_WM_STATE_ADD    = C.int(1) // add/set property
  _NET_WM_STATE_TOGGLE = C.int(2) // toggle property
)
type
  vector = [3]float64
type
  xwindow struct {
             win C.Window // C.XID = C.ulong
            x, y int
          wd, ht C.uint // window
              gc C.GC
  buffer, shadow C.Pixmap
            buff bool
cF, cB, cFA, cBA,
      scrF, scrB col.Colour
          lineWd Linewidth
             fsp *C.XFontStruct
        fontsize font.Size
   wd1, ht1, bl1 C.uint // font baseline
     transparent bool
     cursorShape,
    consoleShape,
      blinkShape Shape
  blinkX, blinkY uint
      blinkMutex sync.Mutex
        blinking bool
     firstExpose bool
         mouseOn bool
//       pointer ptr.Pointer
          xM, yM int
      subWindows []C.Window
             eye, // only for openGL
          eyeOld,
           focus vector
             vec [3]vector // coord system of eye
           delta float64 // invariant: delta == distance (origin, focus)
           angle [3]float64
                 }
var (
  dspl string = env.Val ("DISPLAY")
  initialized = false
  dpy *C.struct__XDisplay
  rootWin C.Window
  screen C.int
  monitorWd, monitorHt C.uint // full screen
  fullScreen mode.Mode
  planes C.uint
  fdNavi uint
  naviAtom C.Atom
  navipipe chan navi.Command
  actual *xwindow
  first bool = true // to start goSendEvents only once
  winList []*xwindow
  txt = []string { "", "",
                   "KeyPress", "KeyRelease", "ButtonPress", "ButtonRelease", "MotionNotify",
                   "EnterNotify", "LeaveNotify", "FocusIn", "FocusOut", "KeymapNotify",
                   "Expose", "GraphicsExpose", "NoExpose", "VisibilityNotify",
                   "CreateNotify", "DestroyNotify", "UnmapNotify", "MapNotify", "MapRequest",
                   "ReparentNotify", "ConfigureNotify", "ConfigureRequest", "GravityNotify",
                   "ResizeRequest", "CirculateNotify", "CirculateRequest",
                   "PropertyNotify", "SelectionClear", "SelectionRequest", "SelectionNotify",
                   "ColormapNotify", "ClientMessage", "MappingNotify", "GenericEvent", "LASTEvent"}
  startSendEvents = make(chan int)
)

func underX() bool {
  return dspl != ""
}

func far() bool {
  return dspl[0] == 'l' // localhost
}

func init() {
  go sendEvents()
}

func initX() {
  if initialized { return }
  if C.XInitThreads() == 0 { panic ("XKern.XInitThreads error") }
  d := C.CString(dspl); defer C.free (unsafe.Pointer(d))
  dpy = C.XOpenDisplay (d); if dpy == nil { panic ("dpy == nil") }
  screen := C.XDefaultScreen (dpy)
  monitorWd, monitorHt = C.uint(C.XDisplayWidth (dpy, screen)),
                         C.uint(C.XDisplayHeight (dpy, screen))
  fullScreen = mode.ModeOf (uint(monitorWd), uint(monitorHt))
  planes = C.uint(C.XDefaultDepth (dpy, screen))
  col.SetDepth (uint(planes))
  initialized = true
}

func (X *xwindow) Display() C.struct__XDisplay {
  return *dpy
}

func cc (c col.Colour) C.ulong {
  return C.ulong(c.Code())
}

func (X *xwindow) Name (n string) {
  s := C.CString (str.Lat1(n)); defer C.free(unsafe.Pointer(s))
  C.XStoreName (dpy, X.win, s)
}

func (X *xwindow) clear() {
  C.XSetForeground (dpy, X.gc, cc (X.scrB))
  C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, 0, 0, X.wd, X.ht)
  C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, 0, 0, X.wd, X.ht)
  C.XFillRectangle (dpy, C.Drawable(X.shadow), X.gc, 0, 0, X.wd, X.ht)
  C.XSetForeground (dpy, X.gc, cc (X.scrF))
}

func pause() {
  time.Sleep (time.Duration (1e7))
}

func (X *xwindow) mousePointer (on bool) {
  if on {
//    C.XUndefineCursor (dpy, X.win)
    X.SetPointer (ptr.Gumby)
  } else {
    var c C.XColor; c.red, c.green, c.blue = C.ushort(0), C.ushort(0), C.ushort(0)
    s := C.CString(string([]byte{ 0, 8, 0, 0, 0, 0 })); defer C.free (unsafe.Pointer(s))
    m := C.XCreateBitmapFromData (dpy, C.Drawable(X.win), s, C.uint(8), C.uint(8));
    cursor := C.XCreatePixmapCursor (dpy, m, m, &c, &c, C.uint(0), C.uint(0))
    C.XDefineCursor (dpy, X.win, cursor)
    C.XFreeCursor (dpy, cursor)
    C.XFreePixmap (dpy, m)
  }
  X.mouseOn = on
//  X.Flush()
}

func (X *xwindow) mousePointerOn() bool {
  return X.mouseOn
}

func new_(x, y uint, m mode.Mode) XWindow {
  initX()
  X := new(xwindow)
  actual = X
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = C.uint(mode.Wd(m)), C.uint(mode.Ht(m))
  if X.wd > monitorWd || X.ht > monitorHt { panic ("win too large: " +
                         strconv.Itoa(int(X.wd)) + " x " + strconv.Itoa(int(X.ht))) }
  X.scrF, X.scrB = col.StartCols()
  X.cF, X.cB = col.StartCols()
  X.cFA, X.cBA = col.StartColsA()
  a := [12]C.int{ 4, 8, 1, 9, 1, 10, 1, 12, 16, 5, 0 }
  vi := C.glXChooseVisual (dpy, screen, &a[0] )
  cx := C.glXCreateContext (dpy, vi, C.GLXContext(nil), C.True)
  rootWin = C.XRootWindow (dpy, vi.screen)
  var swa C.XSetWindowAttributes
  swa.colormap = C.XCreateColormap (dpy, rootWin, vi.visual, C.AllocNone)
  swa.border_pixel = 8
  swa.event_mask = C.KeyPressMask | C.KeyReleaseMask | C.ButtonPressMask | C.ButtonReleaseMask |
                   C.FocusChangeMask | C.ExposureMask | C.VisibilityChangeMask | C.StructureNotifyMask
//  swa.override_redirect = C.True
  X.win = C.XCreateWindow (dpy, rootWin, C.int(x), C.int(y), X.wd, X.ht, 0, vi.depth, C.InputOutput,
                           vi.visual, C.CWBorderPixel | C.CWColormap | C.CWEventMask, &swa)
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
  X.buffer = C.XCreatePixmap (dpy, C.Drawable(X.win), X.wd, X.ht, planes)
  X.shadow = C.XCreatePixmap (dpy, C.Drawable(X.win), X.wd, X.ht, planes)
  const mask = (C.KeyPressMask + // C.KeyReleaseMask +
                C.ButtonPressMask + C.ButtonReleaseMask + // C.ButtonMotionMask +
                C.PointerMotionMask + // C.PointerMotionHintMask +
//                C.EnterWindowMask + C.LeaveWindowMask +
                C.FocusChangeMask + C.ExposureMask + C.VisibilityChangeMask +
//                C.ResizeRedirectMask + C.PropertyChangeMask +
                C.StructureNotifyMask)
  C.XSelectInput (dpy, X.win, mask)
  X.mousePointer (true)
  if X.wd == monitorWd && X.ht == monitorHt { C.fullscreen (dpy, X.win, rootWin, _NET_WM_STATE_ADD) }
  winList = append (winList, X)
  defer X.doBlink()
  X.firstExpose = true
  if first {
    first = false
/*
    p := C.CString ("WM_PROTOCOLS"); defer C.free (unsafe.Pointer(p))
    wm_protocols := C.XInternAtom (dpy, p, C.False)
    C.XSetWMProtocols (dpy, X.win, &wm_protocols, 1)
*/
/*
    m := C.CString ("navi"); defer C.free (unsafe.Pointer(m))
    naviAtom = C.XInternAtom (dpy, m, C.False)
    if navi.Exists() {
      navipipe = navi.Channel()
      go X.catchNavi()
    }
*/
    startSendEvents <- 0
  }
  C.XMapWindow (dpy, X.win)

/*
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
    println ("at initializing xwin:" + txt [et])
  }
  X.clear()
*/
  return X
}

func newMax() XWindow {
  return new_(0, 0, mode.ModeOf (maxRes()))
}

// W == imp (w) iff w == W.win
func imp (w C.Window) *xwindow {
  for _, x := range winList {
    if x.win == w {
      return x
    }
  }
  panic ("µU/xwin.imp: there is no xwindow there")
}

/*
func Win (i uint) *xwindow {
  if int(i) < len(winList) {
    return winList[i]
  }
  panic ("µU/xwin.Win: there is no xwindow there with number" + strconv.Itoa(int(i)))
}
*/

func (X *xwindow) Flush() {
  C.XFlush (dpy)
}

func (X *xwindow) OnFocus() bool {
  return actual == X
}

func (X *xwindow) OffFocus() bool {
  return actual != X
}

func (X *xwindow) Subwindow (x, y int, w, h uint) {
  s := C.XCreateSimpleWindow (dpy, X.win, C.int(x), C.int(y), C.uint(w), C.uint(h),
                              0, cc (X.scrF), cc (X.scrB))
  X.subWindows = append (X.subWindows, s)
}

func (X *xwindow) Sub (n uint) C.Window {
  if int(n) < len (X.subWindows) {
    return X.subWindows[n]
  }
  return C.Window(0)
}

func (X *xwindow) catchNavi() {
  for {
    C.navi (dpy, X.win, naviAtom)
  }
}

// TODO
func (x *xwindow) Fin() {
  C.XFreePixmap (dpy, x.buffer)
  C.XFreePixmap (dpy, x.shadow)
  C.XUnmapWindow (dpy, x.win)
  C.XDestroyWindow (dpy, x.win)
/*
// mouse.showCursor (true)
  C.glXMakeCurrent (dpy, C.None, nil)
  C.glXDestroyContext (dpy, X.cx)
  X.cx = nil
*/
  C.XCloseDisplay (dpy)
}

func fin() {
//  C.XFreeGC (dpy, X.gc)
  C.XDestroyWindow (dpy, rootWin)
  C.XCloseDisplay (dpy)
  initialized = false
}

func (X *xwindow) Win2buf() {
  X.win2buf()
}

func (X *xwindow) win2buf() {
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc, 0, 0, X.wd, X.ht, 0, 0)
  C.XFlush (dpy)
}

func (X *xwindow) buf2win() {
  C.XCopyArea (dpy, C.Drawable(X.buffer), C.Drawable(X.win), X.gc, 0, 0, X.wd, X.ht, 0, 0)
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
  pause()
  for {
    for C.XPending (dpy) > 0 {
      C.XNextEvent (dpy, &xev)
      event.C, event.S = 0, 0
      eventtype = C.evtyp (&xev)
      event.T = uint(eventtype)
// print (txt[eventtype] + "  ")
      switch eventtype {
      case C.Expose:
        w = C.exposeWin (&xev)
//        W.buf2win()
				W.win2buf() // XXX
        if W.firstExpose {
          W.firstExpose = false
//          C.waitForLastContExpose (dpy, &xev) // XXX
//          C.wait (dpy, &xev) // XXX
        }
//      W.win2buf() // XXX
      case C.KeyPress:
// println ("#chars in buffer:", C.lookupString(xev))
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
// print (txt[eventtype] + "  ")
        w = C.motionWin (&xev)
        W = imp (w) // w == W.win
        event.C, event.S = uint(0), uint(C.motionState (&xev))
        W.xM, W.yM = int(C.motionX (&xev)), int(C.motionY (&xev))
      case C.EnterNotify: case C.LeaveNotify:
// print (txt[eventtype] + "  ")
        w = C.enterLeaveWin (&xev)
        W = imp (w) // w == W.win
      case C.FocusIn:
// print (txt[eventtype] + "  ")
        w = C.buttonWin (&xev)
        W = imp (w) // w == W.win
        actual = W
        W.buf2win() // XXX
      case C.FocusOut:
// print (txt[eventtype] + "  ")
        w = C.buttonWin (&xev)
        W = imp (w) // w == W.win
//        xwindow = imp (rootWin) // XXX
      case C.KeymapNotify:
        ;
      case C.GraphicsExpose:
        ;
      case C.NoExpose:
        ;
      case C.VisibilityNotify:
        W.win2buf()
        ;
      case C.CreateNotify:
        ;
      case C.DestroyNotify:
        ;
      case C.UnmapNotify:
        w = C.unmapWin (&xev)
//        println ("unmapped xwindow", int(w))
        W = imp (w) // w == W.win
        W.win2buf()
      case C.MapNotify:
        w = C.mapWin (&xev)
//        println ("mapped xwindow", int(w))
        W = imp (w) // w == W.win
        W.buf2win()
      case C.MapRequest:
        ;
      case C.ReparentNotify:
        ;
      case C.ConfigureNotify:
//          C.glViewport (C.GLint(0), C.GLint(0), C.GLsizei(C.configureWd(&xev)),
//                                                C.GLsizei(C.configureHt(&xev)))
      case C.ConfigureRequest:
        ;
      case C.GravityNotify:
        ;
      case C.ResizeRequest:
/*
        w = C.resizeWin (&xev)
        W = imp (w)
        W.buf1win()
*/
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
      case C.ClientMessage:
        mT := C.mT (&xev)
        if mT != naviAtom { println ("unknown xclient.message_type ", uint32(mT)) }
      case C.MappingNotify:
        ;
      case C.GenericEvent:
        ;
      default:
        panic ("xwin.sendEvents got " + txt[eventtype] + " Xevent")
      }
      switch eventtype {
      case C.KeyRelease: // ignore
      case C.KeyPress:
        if event.C != 64 {
          Eventpipe <- event
        }
      case C.ButtonPress, C.ButtonRelease, C.MotionNotify:
        Eventpipe <- event
      default:
//        println (txt[typ])
      }
    }
  }
}
