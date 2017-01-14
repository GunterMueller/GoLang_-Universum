package xker

// (c) murus.org  v. 161216 - license see murus.go

// #cgo LDFLAGS: -lX11 -lGL
// #include <stdlib.h>
// #include <string.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// #include <X11/Xatom.h>
// #include <GL/gl.h>
// #include <GL/glx.h>
/*
int typ (XEvent *e) { return (*e).type; }

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

Window unmapWin (XEvent *e) { return (*e).xunmap.window; }

Window mapWin (XEvent *e) { return (*e).xmap.window; }

Window configureWin (XEvent *e) { return (*e).xconfigure.window; }
int configureX (XEvent *e) { return (*e).xconfigure.x; }
int configureY (XEvent *e) { return (*e).xconfigure.y; }
int configureWd (XEvent *e) { return (*e).xconfigure.width; }
int configureHt (XEvent *e) { return (*e).xconfigure.height; }

Window resizeWin (XEvent *e) { return (*e).xresizerequest.window; }
unsigned int resizeWd (XEvent *e) { return (*e).xresizerequest.width; }
unsigned int resizeHt (XEvent *e) { return (*e).xresizerequest.height; }

Window circulateWin (XEvent *e) { return (*e).xcirculaterequest.window; }

Atom mT (XEvent *e) { return (*e).xclient.message_type; }

void initialize (Display *d, int s, Window w) {
  int a[11];
  a[0] = GLX_RED_SIZE;     a[1] = 1;
  a[2] = GLX_GREEN_SIZE;   a[3] = 1;
  a[4] = GLX_BLUE_SIZE;    a[5] = 1;
  a[6] = GLX_DOUBLEBUFFER; a[7] = 1;
  a[8] = GLX_DEPTH_SIZE;   a[9] = 1;
  a[10] = 0;
  int n;
  GLXFBConfig config = *(glXChooseFBConfig (d, s, a, &n));
  GLXContext c = glXCreateNewContext (d, config, GLX_RGBA_TYPE, NULL, 1);
  glXMakeContextCurrent (d, w, w, c);
}
*/
import
  "C"
import (
  "unsafe"; "strconv"; "sync"
  "time"
  . "murus/shape"; "murus/ptr"; . "murus/linewd"
  "murus/env"; "murus/str"
  "murus/mode"; "murus/col"; "murus/font"
  "murus/navi"
)
const (
  nul = C.int(0)
  unul = C.uint(0)
  // standards.freedesktop.org/wm-spec:
  _NET_WM_STATE_REMOVE = nul      // remove/unset property
  _NET_WM_STATE_ADD    = C.int(1) // add/set property
  _NET_WM_STATE_TOGGLE = C.int(2) // toggle property
)
type
  window struct {
           x, y int
         wd, ht C.uint // window
            win C.Window // C.XID = C.ulong
             gc C.GC
         buffer,
         shadow C.Pixmap
           buff bool
       wd1, ht1,
            bl1 C.uint // font baseline
         cF, cB,
     scrF, scrB col.Colour
         lineWd Linewidth
            fsp *C.XFontStruct
       fontsize font.Size
    transparent bool
    cursorShape,
   consoleShape,
     blinkShape Shape
 blinkX, blinkY uint
     blinkMutex sync.Mutex
       blinking bool
    firstExpose bool
        mouseOn bool
//      pointer ptr.Pointer
         xM, yM int
     subWindows []C.Window // TODO tree
                }
var (
  dspl string = env.Val ("DISPLAY")
//  dpy *C.struct_Display // up to Go 1.2.2
  dpy *C.struct__XDisplay // since Go 1.3
  rootWindow C.Window
  screen C.int
  fullScreen mode.Mode
  width, height C.uint // full screen
  planes C.uint
  fdNavi uint
  naviAtom C.Atom
  navipipe chan navi.Command
  initialized bool
  actual *window
  first bool = true
  winList []*window // TODO tree
  txt = []string { "", "",
          "KeyPress", "KeyRelease", "ButtonPress", "ButtonRelease", "MotionNotify",
          "EnterNotify", "LeaveNotify", "FocusIn", "FocusOut", "KeymapNotify",
          "Expose", "GraphicsExpose", "NoExpose", "VisibilityNotify",
          "CreateNotify", "DestroyNotify", "UnmapNotify", "MapNotify", "MapRequest",
          "ReparentNotify", "ConfigureNotify", "ConfigureRequest", "GravityNotify",
          "ResizeRequest", "CirculateNotify", "CirculateRequest",
          "PropertyNotify", "SelectionClear", "SelectionRequest", "SelectionNotify",
          "ColormapNotify", "ClientMessage", "MappingNotify", "GenericEvent", "LASTEvent" }
)

func underX() bool {
  return dspl != ""
}

func Far() bool {
  return dspl[0] == 'l' // localhost
}

func initX() {
  if initialized { return }
  if C.XInitThreads() == nul { panic ("XKern.XInitThreads error") }
  d:= C.CString(dspl); defer C.free (unsafe.Pointer(d))
  dpy = C.XOpenDisplay (d)
  if dpy == nil { panic ("dpy == nil") }
  rootWindow = C.XDefaultRootWindow (dpy)
  screen = C.XDefaultScreen (dpy)
  width, height = C.uint(C.XDisplayWidth (dpy, screen)), C.uint(C.XDisplayHeight (dpy, screen))
  fullScreen = mode.ModeOf (uint(width), uint(height))
  planes = C.uint(C.XDefaultDepth (dpy, screen))
  col.SetDepth (uint(planes))
  initialized = true
}

func cc (c col.Colour) C.ulong {
  return C.ulong(col.Code (c))
}

func goSendEvents() {
  go sendEvents()
}

func (X *window) warpMouse() {
  X.WarpMouseGr (int(X.wd) / 2, int(X.ht) / 2)
}

func (X *window) clear() {
  C.XSetForeground (dpy, X.gc, cc (X.scrB))
  C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, nul, nul, X.wd, X.ht)
  C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, nul, nul, X.wd, X.ht)
  C.XFillRectangle (dpy, C.Drawable(X.shadow), X.gc, nul, nul, X.wd, X.ht)
  C.XSetForeground (dpy, X.gc, cc (X.scrF))
}

func wait() {
  time.Sleep (time.Duration (1e7))
}

func (X *window) mousePointer (on bool) {
  if on {
//    C.XUndefineCursor (dpy, X.win)
    X.SetPointer (ptr.Gumby)
  } else {
    var c C.XColor; c.red, c.green, c.blue = C.ushort(0), C.ushort(0), C.ushort(0)
    s := C.CString(string([]byte{ 0, 0, 0, 8, 0, 0, 0, 0 })); defer C.free (unsafe.Pointer(s))
    m := C.XCreateBitmapFromData (dpy, C.Drawable(X.win), s, C.uint(8), C.uint(8));
    cursor := C.XCreatePixmapCursor (dpy, m, m, &c, &c, C.uint(0), C.uint(0))
    C.XDefineCursor (dpy, X.win, cursor)
    C.XFreeCursor (dpy, cursor)
    C.XFreePixmap (dpy, m)
  }
  X.mouseOn = on
//  X.Flush()
}

func (X *window) mousePointerOn() bool {
  return X.mouseOn
}

func newWin (x, y uint, m mode.Mode) Window {
  initX()
  X:= new(window)
  actual = X
  X.x, X.y = int(x), int(y)
  X.wd, X.ht = C.uint(mode.Wd(m)), C.uint(mode.Ht(m))
  if X.wd > width || X.ht > height {
    panic ("new window too large: " +
           strconv.Itoa(int(X.wd)) + " > " + strconv.Itoa (int(width)) + " or " +
           strconv.Itoa(int(X.ht)) + " > " + strconv.Itoa (int(height)))
  }
  X.cF, X.cB = col.StartCols()
  X.scrF, X.scrB = col.StartCols()
  X.win = C.XCreateSimpleWindow (dpy, C.Window(rootWindow), C.int(x), C.int(y), X.wd, X.ht, unul, C.ulong(0), cc (X.scrB))
  X.gc = C.XDefaultGC (dpy, screen)
  C.XSetGraphicsExposures (dpy, X.gc, C.False)
  X.buffer = C.XCreatePixmap (dpy, C.Drawable(X.win), X.wd, X.ht, planes)
  X.shadow = C.XCreatePixmap (dpy, C.Drawable(X.win), X.wd, X.ht, planes)
  X.cF, X.cB = col.White, col.Black
  C.initialize (dpy, screen, X.win)
  var sh C.XSizeHints
  sh.flags = C.PPosition | C.PSize | C.PMinSize | C.PMaxSize
  sh.x, sh.y = C.int(x), C.int(y)
  sh.min_width, sh.min_height = C.int(X.wd), C.int(X.ht)
  sh.max_width, sh.max_height = C.int(X.wd), C.int(X.ht)
  C.XSetNormalHints (dpy, X.win, &sh)
  C.XMapRaised (dpy, X.win)
  const mask = (C.KeyPressMask + // C.KeyReleaseMask +
                C.ButtonPressMask + C.ButtonReleaseMask + // C.ButtonMotionMask +
                C.PointerMotionMask + // C.PointerMotionHintMask +
//                C.EnterWindowMask + C.LeaveWindowMask +
                C.FocusChangeMask +
                C.ExposureMask + C.VisibilityChangeMask +
//                C.ResizeRedirectMask +
//                C.PropertyChangeMask +
                C.StructureNotifyMask)
  C.XSelectInput (dpy, X.win, mask)
  X.mousePointer (true)
  X.warpMouse()
  X.Name (env.Par (0))
  if X.wd == width && X.ht == height { C.fullscreen (dpy, X.win, rootWindow, _NET_WM_STATE_ADD) }
  winList = append (winList, X)
  defer X.warpMouse()
  defer X.doBlink()
  defer X.clear()
  X.firstExpose = true
  defer wait()
  if first {
    first = false
/*
    p:= C.CString ("WM_PROTOCOLS"); defer C.free (unsafe.Pointer(p))
    wm_protocols:= C.XInternAtom (dpy, p, C.False)
    C.XSetWMProtocols (dpy, X.win, &wm_protocols, 1)
*/
/*
    m:= C.CString ("navi"); defer C.free (unsafe.Pointer(m))
    naviAtom = C.XInternAtom (dpy, m, C.False)
    if navi.Exists() {
      navipipe = navi.Channel()
      go X.catchNavi()
    }
*/
    defer goSendEvents()
  }
  return X
}

func newMax() Window {
  return newWin (0, 0, mode.ModeOf (maxRes()))
}

// X == imp (w) iff X.win == w
func imp (w C.Window) *window {
  for _, x:= range winList {
    if x.win == w {
      return x
    }
  }
  panic ("murus/xker.imp: there is no window there")
}

/*
func Win (i uint) *window {
  if int(i) < len(winList) {
    return winList[i]
  }
  panic ("murus/xker.Win: there is no window there with number" + strconv.Itoa(int(i)))
}
*/

func (X *window) Flush() {
  C.XFlush (dpy)
}

func (X *window) OnFocus() bool {
  return actual == X
}

func (X *window) OffFocus() bool {
  return actual != X
}

func (X *window) Subwindow (x, y int, w, h uint) {
  s:= C.XCreateSimpleWindow (dpy, X.win, C.int(x), C.int(y), C.uint(w), C.uint(h), unul, C.ulong(0), cc (X.scrB))
  X.subWindows = append (X.subWindows, s)
}

func (X *window) Sub (n uint) C.Window {
  if int(n) < len (X.subWindows) {
    return X.subWindows[n]
  }
  return C.Window(0)
}

func (X *window) Name (n string) {
  s:= C.CString (str.Lat1(n))
  C.XStoreName (dpy, X.win, s)
  C.free (unsafe.Pointer(s))
}

func (X *window) catchNavi() {
  for {
    C.navi (dpy, X.win, naviAtom)
  }
}

// TODO
func (x *window) Fin() {
  C.XFreePixmap (dpy, x.buffer)
  C.XFreePixmap (dpy, x.shadow)
  C.XUnmapWindow (dpy, x.win)
  C.XDestroyWindow (dpy, x.win)
}

func fin() {
//  C.XFreeGC (dpy, X.gc)
  C.XDestroyWindow (dpy, rootWindow)
  C.XCloseDisplay (dpy)
  initialized = false
}

func (X *window) Win2buf() {
  X.win2buf()
}

func (X *window) win2buf() {
  C.XCopyArea (dpy, C.Drawable(X.win), C.Drawable(X.buffer), X.gc, nul, nul, X.wd, X.ht, nul, nul)
  C.XFlush (dpy)
}

func (X *window) buf2win() {
  C.XCopyArea (dpy, C.Drawable(X.buffer), C.Drawable(X.win), X.gc, nul, nul, X.wd, X.ht, nul, nul)
  C.XFlush (dpy)
}

func sendEvents() {
  var (
    xev C.XEvent
    typ C.int
    w C.Window
    W *window
    event Event
  )
  for {
//    println ("waiting for an XEvent ...") 
    C.XNextEvent (dpy, &xev)
    event.C, event.S = 0, 0
    typ = C.typ (&xev)
    event.T = uint(typ)
    switch typ {
    case C.Expose:
      w = C.exposeWin (&xev)
      W = imp (w) // w == W.win
      if W.firstExpose {
        W.firstExpose = false
        C.waitForLastContExpose (dpy, &xev)
      }
      W.buf2win()
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
    case C.EnterNotify: case C.LeaveNotify:
      w = C.enterLeaveWin (&xev)
      W = imp (w) // w == W.win
    case C.FocusIn:
      w = C.buttonWin (&xev)
      W = imp (w) // w == W.win
      actual = W
    case C.FocusOut:
      w = C.buttonWin (&xev)
      W = imp (w) // w == W.win
//      window = imp (rootWindow)
    case C.KeymapNotify:
      ;
    case C.GraphicsExpose: case C.NoExpose:
      ;
    case C.VisibilityNotify:
      ;
    case C.CreateNotify: case C.DestroyNotify:
      ;
    case C.UnmapNotify:
      w = C.unmapWin (&xev)
//      println ("unmapped window", int(w))
      W = imp (w) // w == W.win
      W.win2buf()
    case C.MapNotify:
      w = C.mapWin (&xev)
//      println ("mapped window", int(w))
      W = imp (w) // w == W.win
      W.buf2win()
    case C.MapRequest:
    case C.ReparentNotify: case C.ConfigureNotify: case C.ConfigureRequest:
    case C.GravityNotify:
    case C.ResizeRequest:
/*
      w = C.resizeWin (&xev)
      W = imp (w)
      W.buf2win()
*/
    case C.CirculateNotify: case C.CirculateRequest:
    case C.PropertyNotify: case C.SelectionClear: case C.SelectionRequest: case C.SelectionNotify:
    case C.ColormapNotify:
    case C.ClientMessage:
      mT:= C.mT (&xev)
      if mT != naviAtom { println ("unknown xclient.message_type ", uint32(mT)) }
    case C.MappingNotify: case C.GenericEvent:
    default:
      panic ("xker.sendEvents got " + txt[typ] + " Xevent")
    }
    switch typ {
    case C.KeyRelease: // ignore
    case C.KeyPress, C.ButtonPress, C.ButtonRelease, C.MotionNotify:
      Eventpipe <- event
    default:
//      println (txt[typ])
    }
  }
}
