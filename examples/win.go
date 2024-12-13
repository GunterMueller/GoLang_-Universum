package main

// #cgo LDFLAGS: -lX11
// #include <X11/Xlib.h>
// #include <X11/Xutil.h>
// int typ (XEvent *e) { return (*e).type; }
// void wait (Display *d, XEvent *e) { while (XCheckTypedEvent (d, Expose, e)); }
import "C"

func main() {
	d := C.XOpenDisplay (C.CString(""))
	s := C.XDefaultScreen (d)
	w := C.XCreateSimpleWindow (d, C.XRootWindow(d, s), 0, 0, 160, 90, 0,
                              C.ulong(0), C.XWhitePixel (d, s))
	c := C.XDefaultGC(d, s)
	C.XSelectInput (d, w, C.ExposureMask + C.KeyPressMask + C.StructureNotifyMask)
	C.XMapWindow (d, w)
	var e C.XEvent
	loop:
  for {
		C.XNextEvent (d, &e)
    switch C.typ (&e) {
		case C.Expose:
			C.wait (d, &e)
			C.XDrawString (d, C.Drawable (w), c, 55, 50, C.CString ("a window"), 8)
		case C.KeyPress:
			break loop
		}
	}
}
