package scr

// (c) Christian Maurer   v. 210107 - license see µU.go

/* Pre: For use in a (tty)-console:
          The framebuffer is usable, i.e. one of the options "vga=..."
          is contained in the line "kernel ..." of /boot/grub/menu.lst
          (posible values can be found at the end of imp.go).
          Users are in the group video (or world has the rights "r" and "w"
          in /dev/fb0) and world has the right "r" in /dev/input/mice.
        For use in a screen on a graphical user interface:
          X is installed.
        Programs for execution on far hosts are only called under X.
   Fore-/background colour of the screen and actual fore-/backgroundcolour
   are White and Black. The screen is cleared and the cursor is off.
   In a console SIGUSR1 and SIGUSR2 are used internally and not any more available.
   No process is in the exclusive possession of the screen. */

import (
  "µU/obj"
  . "µU/shape"
  . "µU/linewd"
  "µU/mode"
  "µU/xwin"
  "µU/ptr"
  "µU/col"
  "µU/font"
)
const (
  Look = xwin.Look
  Walk = xwin.Walk
  Fly  = xwin.Fly
)
type
  Screen interface {

// The keyboard is switched back to normal mode.
  Fin()

// Under X, the screen is newly written.
  Flush()

// Under X, in the title bar of the window framing the screen 
// the string n appears, unless the screen was initialized by a call of NewMax.
  Name (n string)

// Returns the actual mode.
  ActMode() mode.Mode

// Returns the coordinates of the top left corner of the screen.
  X() uint
  Y() uint

// Returns - depending on the actual fontsize -
// the number of textlines and -columns of the actual mode.
  NLines() uint
  NColumns() uint

// Returns the pixelwidth/-height of the screen in the actual mode.
  Wd() uint
  Ht() uint

// Returns the pixel distance between two textlines
// = charheight/-width of the actual fontsize (s. below).
  Wd1() uint
  Ht1() uint

// Return the quotient Pixelwidth / Pixelheight of the actual mode.
  Proportion() float64

// colours /////////////////////////////////////////////////////////////

// The colours of the screen are set to f and b (fore-/background).
  ScrColours (f, b col.Colour)
  ScrColourF (f col.Colour)
  ScrColourB (b col.Colour)

// Returns the fore-/backgroundcolour of the screen.
  ScrCols() (col.Colour, col.Colour)
  ScrColF() col.Colour
  ScrColB() col.Colour

// The actual foregroundcolour is f, the actual backgroundcolour is b
// resp. that of the screen.
// The colours of the screen are not changed.
  Colours (f, b col.Colour)
  ColourF (f col.Colour)
  ColourB (b col.Colour)

// Returns the actual fore-/backgroundcolour.
  Cols() (col.Colour, col.Colour)
  ColF() col.Colour
  ColB() col.Colour

// Returns the colour of the pixel at (x, y).
  Colour (x, y uint) col.Colour

// ranges //////////////////////////////////////////////////////////////

// Pre: c + w <= NColumns, l + h <= NLines resp.
//      x <= x1 < Wd, y <= y1 < Ht.

// The screen is cleared between line l and l+h and column c and c+w
// (both including) in its backgroundcolour.
  Clr (l, c, w, h uint)

// The pixels in the rectangle between (x, y) and (x1, y1)
// (both including) have the backgroundcolour of the screen.
  ClrGr (x, y, x1, y1 int)

// The screen is cleared in its backgroundcolour.
// The cursor has the position (0, 0) and is off. 
// // If there exists a mouse, its cursor has the position (?, ?) and is off.
  Cls()

// If on, then the screen buffer is cleared and
// all further output is only going to the screen buffer,
// otherwise, the screen contains the content of the screen buffer
// and all further output is going to the screen.
  Buf (on bool)

// Returns true, iff the output goes only to the screen buffer.
  Buffered() bool

// The content of the screen between line l and l+h and column c and c+w
// is copied into the archive (the former content of the archive is lost).
  Save (l, c, w, h uint)
  SaveGr (x, y, x1, y1 int)
  Save1() // full screen

// The content of the screen between line l and l+h and column c and c+w
// is restored from the archive.
  Restore (l, c, w, h uint)
  RestoreGr (x, y, x1, y1 int)
  Restore1() // full screen

// cursor //////////////////////////////////////////////////////////////

// Pre: l < NLines, c < NColumns.
// The cursor has the position (line, coloumn) == (l, c)
// and the shape s. (0, 0) is the top left top corner.
  Warp (l, c uint, s Shape)

// Pre: x <= NColumsGr - Columnwidth, y <= Ht - Lineheight.
// The cursor has the graphics position (column, line) = (x, y)
// and the shape s. (0, 0) is the top left top corner.
  WarpGr (x, y uint, s Shape)

// text ////////////////////////////////////////////////////////////////

// The position (0, 0) is the top left corner of the screen.
// The pixels of the characters have the actual foregroundcolour,
// the pixels in the rectangles around them have the actual backgroundcolour
// (if transparency is switched on, those pixels are not changed).

// Pre: 32 <= b < 127, l < NLines, c + 1 < NColumns. 
// b is written to the screen at position (line, colum) = (l, c). 
  Write1 (b byte, l, c uint)

// Pre: l < NLines, c + len(s) < NColumns. 
// s is written to the screen starting at position (line, column) == (l, c).
  Write (s string, l, c uint)

// Pre: x + Columnwidth < Wd resp.
//      x + Columnwidth * Länge (s) < Wd,
//      y + Lineheight < Ht.
// b resp. s is written to the screen within the rectangle
// with the top left corner (x, y).
  Write1Gr (b byte, x, y int)
  WriteGr (s string, x, y int)

// Pre: c + number of digits of n < NColumns, l < NLines.
// n is written to the screen starting at position (line, column) == (l, c).
  WriteNat (n, l, c uint)
  WriteNatGr (n uint, x, y int)

// TODO Spec
  Write1InvGr (b byte, x, y int)
  WriteInvGr (s string, x, y int)

// Returns true, iff transparency is set.
  Transparent() bool

// Transparence is switched on, iff t == true.
  Transparence (t bool)

// font ////////////////////////////////////////////////////////////////

// Returns the actual fontsize; at the beginning normal.
  ActFontsize() font.Size

// f is the actual fontsize.
// NColumns and NLines are changed accordingly.
  SetFontsize (f font.Size)

// graphics ////////////////////////////////////////////////////////////

// Position (0, 0) is the top left corner of the screen.
// All output is done in the actual foregroundcolour;
// For operations with name ...Inv all pixels have the complementary
// colour of the fgcolour; for operations with name ...Full
// also all pixels in the interior have these colours.
// The actual linewidth at the beginning is Thin.

// Returns the actual linewidth.
  ActLinewidth() Linewidth

// The actual linewidth is w.
  SetLinewidth (w Linewidth)

// Pre: See above.
// A pixel in the actual foregroundcolour is set at position (x, y)
// on the screen resp. the colour of that pixel is inverted.
  Point (x, y int)
  PointInv (x, y int)

  OnPoint (x, y, a, b int, d uint) bool

// Pre: See above.
// At (xs[i], ys[i]) (i < len(xs) == len(ys)) a pixel is set in the actual
// foregroundcolour resp. that pixel is inverted in its colour.
  Points (xs, ys []int)
  PointsInv (xs, ys []int)

// Pre: See above.
// The part of the line segment between (x, y) and (x1, y1)
// visible on the screen is drawn in the actual foregroundcolour resp.
// the pixels on that part are inverted in their colour.
  Line (x, y, x1, y1 int)
  LineInv (x, y, x1, y1 int)

// Pre: See above.
// Returns true, iff the point at (x, y) has a distance of
// at most d pixels from the line segment between (x, y) to (x1, y1).
  OnLine (x, y, x1, y1, a, b int, d uint) bool

// Pre: See above.
//      If the calling process runs under X:
//        -1<<15 <= x[i], x1[i], y[i], y1[i] < 1<<15
//        for all i < n:= len(x) == len(y).
//      Otherwise:
//        0 <= x[i], x1[i] < Wd and
//        0 <= y[i], y1[i] < Ht for all i < N.
// For all i < n the parts of the line segments between (x[i], y[i]) and (x1[i], y1[i]),
// that are visible on the screen, are drawn in the actual foregroundcolour
// resp. all points on them are inverted.
  Lines (x, y, x1, y1 []int)
  LinesInv (x, y, x1, y1 []int)

// Pre: See above.
// TODO Spec
  OnLines (x, y, x1, y1 []int, a, b int, d uint) bool

// Pre: See above.
//      x[i] < Wd, y[i] < Ht für alle i < n:= len(x) == len(y).
// From (x[0], y[0]) over (x[1], y[1]), ... until (x[n-1], y[n-1])
// a sequence of line segments is drawn resp. all points on it are inverted.
  Segments (x, y []int)
  SegmentsInv (x, y []int)

// Returns true, iff the point at (a, b) has a distance of at most d pixels
// from one of the sequence of line segments defined by x and y.
  OnSegments (x, y []int, a, b int, d uint) bool

// Pre: See above.
// A line through (x, y) and (x1, y1) is drawn resp. all points on it are inverted.
  InfLine (x, y, x1, y1 int)
  InfLineInv (x, y, x1, y1 int)

// Returns true, iff the point at (a, b) has a distance of at most d pixels
// from the line through (x, y) and (x1, y1).
  OnInfLine (x, y, x1, y1, a, b int, d uint) bool

// Pre: See above.
// Between (x, y), (x1, y1) and (x2, y2) a triangle is drawn
// in the actual foregroundcolour resp. all points on it are inverted
// resp. all its interior points (including its borders) are drawn / inverted.
  Triangle (x, y, x1, y1, x2, y2 int)
  TriangleInv (x, y, x1, y1, x2, y2 int)
  TriangleFull (x, y, x1, y1, x2, y2 int)
  TriangleFullInv (x, y, x1, y1, x2, y2 int)

// Pre: See above.
// Between (x, y) and (x1, y1) a rectangle (with horizontal and vertical borders)
// is drawn in the actual foregroundcolour resp. all points on it are inverted
// resp. all its interior points (including its borders) are drawn / inverted.
  Rectangle (x, y, x1, y1 int)
  RectangleInv (x, y, x1, y1 int)
  RectangleFull (x, y, x1, y1 int)
  RectangleFullInv (x, y, x1, y1 int)

// Pre: See above.
// Returns true, iff the point at (a, b) has a distance of at most d pixels
// from the border of the rectangle between (x, y) and (x1, y1).
  OnRectangle (x, y, x1, y1, a, b int, d uint) bool

// Returns true, iff the point at (a, b) is not outside the rectangle
// between (x, y) and (x1, y1) up to tolerance of t pixels.
  InRectangle (x, y, x1, y1, a, b int, t uint) bool

// Pre: See above. For n:= len(x) == len(y): n > 2 and
//      PolygonFull:
//        The calling process runs under X;
//        the polygon defined by x and y is convex and drawn in the same colour.
//      PolygonFull1:
//        (x0, y0) lies in the interior of the polygon defined by x and y.
//        The polygon is drawn in the same colour.
// A polygon is drawn between (x[0], y[0]), (x[1], y[1]), ... (x[n-1], y[n-1), (x[0], y[0])
// resp. all pixels on it are inverted resp. the polygon is filled.
  Polygon (x, y []int)
  PolygonInv (x, y []int)
  PolygonFull (x, y []int)
  PolygonFullInv (x, y []int)

// Returns true, iff the point at (a, b) has a distance of at most d pixels
// from the polyon defined by x and y.
  OnPolygon (x, y []int, a, b int, d uint) bool

// Pre: See above. r <= x, x + r < Wd, r <= y, y + r < Ht. 
// Around (x, y) a circle with radius r is drawn / inverted
// resp. all points in its interior are set / inverted.
  Circle (x, y int, r uint)
  CircleInv (x, y int, r uint)
  CircleFull (x, y int, r uint)
  CircleFullInv (x, y int, r uint)

// Pre: See above. r <= x, x + r < Wd, r <= y, y + r < Ht,
//      a and b given in degrees.
// Around (x, y) an arc with radius r is drawn / inverted
// resp. all points in its interior are set / inverted
// from angle a to angle a+b, starting at vertical upright position
// with a and b signed in mathematical orientation (counterclockwise).
  Arc (x, y int, r uint, a, b float64)
  ArcInv (x, y int, r uint, a, b float64)
  ArcFull (x, y int, r uint, a, b float64)
  ArcFullInv (x, y int, r uint, a, b float64)

// Returns true, iff the point at (x, y) has a distance of at most d pixels
// from the border of the circle around (a, b) with radius r.
  OnCircle (x, y int, r uint, a, b int, d uint) bool
//  InCircle (x, y int, r uint, a, b int) bool // TODO

// Pre: See above. a <= x, x + a < Wd, b <= y, y + b < Ht. 
// Around (x, y) an ellipse with horizontal / vertical semiaxis a / b
// is drawn / inverted resp. all points in its interior are set / inverted.
  Ellipse (x, y int, a, b uint)
  EllipseInv (x, y int, a, b uint)
  EllipseFull (x, y int, a, b uint)
  EllipseFullInv (x, y int, a, b uint)

// Returns true, iff the point at (A, B) has a distance of at most d pixels
// from the border of the ellipse around (x, y) with semiaxis a and b.
  OnEllipse (x, y int, a, b uint, A, B int, d uint) bool
// func InEllipse (x, y int, a, b uint, A, B int) bool // TODO

// Pre: See above. n:= len(xs) == len(ys).
// From (xs[0], ys[0]) to (xs[n], ys[n]) a Beziercurve of order n
// with (xs[1], ys[1]) .. (xs[n-1], ys[n-1]) as nodes is drawn to the screen
// resp. all points on that curve are inverted.
// (For n == 0 the curve is the point (xs[0], ys[0]),
// for n == 1 the line between (xs[0], ys[0]) and (xs[1], ys[1]).
  Curve (xs, ys []int)
  CurveInv (xs, ys []int)

// Returns true, iff the point at (x, y) has a distance of at most d pixels
// from the curve defined by xs and ys.
  OnCurve (xs, ys []int, a, b int, d uint) bool

// mouse ///////////////////////////////////////////////////////////////

// Returns true, iff a mouse is installed.
  MouseEx() bool

// Spec TODO
  SetPointer (p ptr.Pointer)

// Returns the position of the mouse cursor.
// For the result (l, c) holds 0 <= l < NLines and 0 <= c < NColumns.
  MousePos() (uint, uint)

// Returns the position of the mouse cursor.
// For the result (x, y) holds 0 <= x < Wd and 0 <= y < Ht.
  MousePosGr() (int, int)

// Pre: The calling process does not run under X.
// If no mouse exists, nothing has happened.
// Otherwise, the mouse cursor is switched on, iff b (otherwise off).
  MousePointer (b bool)

// Pre: The calling process does not run under X.
// Returns true, iff the mouse cursor is switched on.
  MousePointerOn() bool

// Pre: l < NLines, c < NColumns.
// The mouse cursor has the position (line, column) = (l, c).
  WarpMouse (l, c uint)

// Pre: 0 <= x < Wd, 0 <= y < Ht.
// The mouse cursor has the position (row, line) = (x, y).
  WarpMouseGr (x, y int)

// Pre: c + w <= NColumns, l + h <= NLines.
// Returns false, if there is no mouse; returns otherwise true,
// iff the the mouse cursor is in the interior of the rectangle
// defined by l, c, w, h.
  UnderMouse (l, c, w, h uint) bool

// Pre: 0 <= x <= x1 < Wd, 0 <= y <= y1 < Ht.
// Returns false, if there is no mouse; returns otherwise true,
// iff the mouse cursor is inside the rectangle between (x, y) and (x1, y1)
// or has a distance of at most d pixels from its boundary.
  UnderMouseGr (x, y, x1, y1 int, d uint) bool

// Pre: 0 <= x < Wd, 0 <= y < Ht.
// Returns false, if there is no mouse; returns otherwise true,
// iff the mouse cursor has a distance of at most d pixels from (x, y).
  UnderMouse1 (x, y int, d uint) bool

// serialisation ///////////////////////////////////////////////////////

// Pre: 0 < w <= Wd, 0 < h <= Ht.
// Returns the number of bytes, that are needed to serialize the pixels
// of the rectangle between (0, 0) and (w, h) uniquely invertibly.
  Codelen (w, h uint) uint

// Pre: 0 < w, x + w < Wd, 0 < h, y + h < Ht.
// Returns the byte sequence, that serializes the pixels
// in the rectangle between (x, y) and (x + w, y + h).
  Encode (x, y, w, h uint) obj.Stream

// Pre: s is the result of a call of Encode for some rectangle.
// The pixels of that rectangle are drawn to the screen;
// the rest of the screen is not changed.
  Decode (s obj.Stream)

// ppm-serialization ///////////////////////////////////////////////////

// Returns a string of the form "P6 w h 255" with numbers w and m,
// with only one space between the strings and a line feed at the end.
  PPMHeader (w, h uint) string

// Pre: s is a raw ppm-file generated by a call to PPMEncoder.
// Returns the length of that file.
  PPMCodelen (w, h uint) uint

// Returns the rectangle between upper left corner (x, y) and lower right corner
// (x+w, y+h) of the screen as raw ppm-file with the corresponding header (see above).
  PPMEncode (x, y, w, h uint) obj.Stream

// Pre: s is a raw ppm-file generated by a call to PPMEncoder.
//      The screen used by a constructor-call is big enough.
// This rectangle defined by s is written to the screen
// with upper left corner (x, y).
  PPMDecode (s obj.Stream, x, y uint)

// Pre: s is a raw ppm-file generated by a call to PPMEncoder.
// Returns width and height of the corresponding raw ppm-file.
  PPMSize (s obj.Stream) (uint, uint)

// cut buffer //////////////////////////////////////////////////////////

// The content of the cutbuffer is the former *s and *s is now empty.
  Cut (s *string)

// The content of the cutbuffer is s.
  Copy (s string)

// Returns the content of the cutbuffer.
  Paste() string

// openGL //////////////////////////////////////////////////////////////

// Pre: m <= Fly
  Go (m int, draw func(), ex, ey, ez, fx, fy, fz, nx, ny, nz float64)
}

// Returns a new screen with the size of the physical screen.
// The keyboard is switched to raw mode.
func New (x, y uint, m mode.Mode) Screen { return new_(x,y,m) }

// Returns a new screen of the size given by the mode m.
// The keyboard is switched to raw mode.
func NewMax() Screen { return newMax() }

// Pre: The size of the screen given by x, y, w, h
//      fits into the available physical screen.
// Returns a new screen with upper left corner (x, y),
// width w and height h. The keyboard is switched to raw mode.
func NewWH (x, y, w, h uint) Screen { return newWH(x,y,w,h) }

// Returns the (X, Y)-resolution of the screen in pixels.
func MaxRes() (uint, uint) { return maxRes() }

// Returns true, iff mode.Res(m) <= MaxRes().
func Ok (m mode.Mode) bool { return ok(m) }

// Returns true, if the calling process runs under X.
func UnderX() bool { return underX }

// Returns the colours at the start of the system.
func StartCols() (col.Colour, col.Colour) { return startCols() }
func StartColsA() (col.Colour, col.Colour) { return startColsA() }

// Lock / Unlock guarantee the mutual exclusion when writing on the screen
// (e.g. to avoid, that a process after having set its colours
// is interrupted in a subsequent draw and later resumes its drawing
// in another colour, that was meanwhile changed by another process).
func Lock() { lock() }
func Unlock() { unlock() }

// Pre: s is the stream of a PPM-Header generated by a call to PPMEncoder.
// Returns width and height of the corresponding rectangle and
// the length of that stream encoded in the PPM-Header (see above).
func P6HeaderData (s obj.Stream) (uint, uint, uint, int) { return ppmHeaderData(s) }
