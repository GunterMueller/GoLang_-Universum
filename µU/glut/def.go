package glut

// (c) Christian Maurer   v. 191004 - license see µU.go
//
// from freeglut_std.h; (c) 1999-2000 Pawel W. Olszta. All Rights Reserved.
//
// Specs see books on openGL

// #cgo LDFLAGS: -lglut
// #include <GL/gl.h>
// #include <GL/glu.h>
// #include <GL/glut.h>
// int getWindow() { int w = glutGetWindow(); return w; }
import
  "C"
const (
  KEY_F1  = 1
  KEY_F2  = 2
  KEY_F3  = 3
  KEY_F4  = 4
  KEY_F5  = 5
  KEY_F6  = 6
  KEY_F7  = 7
  KEY_F8  = 8
  KEY_F9  = 9
  KEY_F10 = 10
  KEY_F11 = 11
  KEY_F12 = 12
  KEY_LEFT      = 100
  KEY_UP        = 101
  KEY_RIGHT     = 102
  KEY_DOWN      = 103
  KEY_PAGE_UP   = 104
  KEY_PAGE_DOWN = 105
  KEY_HOME      = 106
  KEY_END       = 107
  KEY_INSERT    = 108
  LEFT_BUTTON   = 0
  MIDDLE_BUTTON = 1
  RIGHT_BUTTON  = 2
  DOWN    = 0
  UP      = 1
  LEFT    = 0
  ENTERED = 1
  RGB     = 0
  RGBA    = 0
  INDEX   = 1
  SINGLE  = 0
  DOUBLE  = 2
  ACCUM   = 4
  ALPHA   = 8
  DEPTH   = 16
  STENCIL = 32
  MULTISAMPLE = 128
  STEREO      = 256
  LUMINANCE   = 512
  GLUT_MENU_NOT_IN_USE    = 0
  GLUT_MENU_IN_USE        = 1
  GLUT_NOT_VISIBLE        = 0
  GLUT_VISIBLE            = 1
  GLUT_HIDDEN             = 0
  GLUT_FULLY_RETAINED     = 1
  GLUT_PARTIALLY_RETAINED = 2
  GLUT_FULLY_COVERED      = 3

  WINDOW_X                = 0x0064
  WINDOW_Y                = 0x0065
  WINDOW_WIDTH            = 0x0066
  WINDOW_HEIGHT           = 0x0067
  WINDOW_BUFFER_SIZE      = 0x0068
  WINDOW_STENCIL_SIZE     = 0x0069
  WINDOW_DEPTH_SIZE       = 0x006A
  WINDOW_RED_SIZE         = 0x006B
  WINDOW_GREEN_SIZE       = 0x006C
  WINDOW_BLUE_SIZE        = 0x006D
  WINDOW_ALPHA_SIZE       = 0x006E
  WINDOW_ACCUM_RED_SIZE   = 0x006F
  WINDOW_ACCUM_GREEN_SIZE = 0x0070
  WINDOW_ACCUM_BLUE_SIZE  = 0x0071
  WINDOW_ACCUM_ALPHA_SIZE = 0x0072
  WINDOW_DOUBLEBUFFER     = 0x0073
  WINDOW_RGBA             = 0x0074
  WINDOW_PARENT           = 0x0075
  WINDOW_NUM_CHILDREN     = 0x0076
  WINDOW_COLORMAP_SIZE    = 0x0077
  WINDOW_NUM_SAMPLES      = 0x0078
  WINDOW_STEREO           = 0x0079
  WINDOW_CURSOR           = 0x007A
  SCREEN_WIDTH            = 0x00C8
  SCREEN_HEIGHT           = 0x00C9
  SCREEN_WIDTH_MM         = 0x00CA
  SCREEN_HEIGHT_MM        = 0x00CB
  MENU_NUM_ITEMS          = 0x012C
  DISPLAY_MODE_POSSIBLE   = 0x0190
  INIT_WINDOW_X           = 0x01F4
  INIT_WINDOW_Y           = 0x01F5
  INIT_WINDOW_WIDTH       = 0x01F6
  INIT_WINDOW_HEIGHT      = 0x01F7
  INIT_DISPLAY_MODE       = 0x01F8
  ELAPSED_TIME            = 0x02BC
  WINDOW_FORMAT_ID        = 0x007B

  HAS_KEYBOARD            = 0x0258
  HAS_MOUSE               = 0x0259
  HAS_SPACEBALL           = 0x025A
  HAS_DIAL_AND_BUTTON_BOX = 0x025B
  HAS_TABLET              = 0x025C
  NUM_MOUSE_BUTTONS       = 0x025D
  NUM_SPACEBALL_BUTTONS   = 0x025E
  NUM_BUTTON_BOX_BUTTONS  = 0x025F
  NUM_DIALS               = 0x0260
  NUM_TABLET_BUTTONS      = 0x0261
  DEVICE_IGNORE_KEY_REPEAT= 0x0262
  DEVICE_KEY_REPEAT       = 0x0263
  HAS_JOYSTICK            = 0x0264
  OWNS_JOYSTICK           = 0x0265
  JOYSTICK_BUTTONS        = 0x0266
  JOYSTICK_AXES           = 0x0267
  JOYSTICK_POLL_RATE      = 0x0268

  OVERLAY_POSSIBLE  = 0x0320
  LAYER_IN_USE      = 0x0321
  HAS_OVERLAY       = 0x0322
  TRANSPARENT_INDEX = 0x0323
  NORMAL_DAMAGED    = 0x0324
  OVERLAY_DAMAGED   = 0x0325

  VIDEO_RESIZE_POSSIBLE     = 0x0384
  VIDEO_RESIZE_IN_USE       = 0x0385
  VIDEO_RESIZE_X_DELTA      = 0x0386
  VIDEO_RESIZE_Y_DELTA      = 0x0387
  VIDEO_RESIZE_WIDTH_DELTA  = 0x0388
  VIDEO_RESIZE_HEIGHT_DELTA = 0x0389
  VIDEO_RESIZE_X            = 0x038A
  VIDEO_RESIZE_Y            = 0x038B
  VIDEO_RESIZE_WIDTH        = 0x038C
  VIDEO_RESIZE_HEIGHT       = 0x038D

  NORMAL  = 0
  OVERLAY = 1

  ACTIVE_SHIFT = 1
  ACTIVE_CTRL  = 2
  ACTIVE_ALT   = 4

  CURSOR_RIGHT_ARROW         = 0x0000
  CURSOR_LEFT_ARROW          = 0x0001
  CURSOR_INFO                = 0x0002
  CURSOR_DESTROY             = 0x0003
  CURSOR_HELP                = 0x0004
  CURSOR_CYCLE               = 0x0005
  CURSOR_SPRAY               = 0x0006
  CURSOR_WAIT                = 0x0007
  CURSOR_TEXT                = 0x0008
  CURSOR_CROSSHAIR           = 0x0009
  CURSOR_UP_DOWN             = 0x000A
  CURSOR_LEFT_RIGHT          = 0x000B
  CURSOR_TOP_SIDE            = 0x000C
  CURSOR_BOTTOM_SIDE         = 0x000D
  CURSOR_LEFT_SIDE           = 0x000E
  CURSOR_RIGHT_SIDE          = 0x000F
  CURSOR_TOP_LEFT_CORNER     = 0x0010
  CURSOR_TOP_RIGHT_CORNER    = 0x0011
  CURSOR_BOTTOM_RIGHT_CORNER = 0x0012
  CURSOR_BOTTOM_LEFT_CORNER  = 0x0013
  CURSOR_INHERIT             = 0x0064
  CURSOR_NONE                = 0x0065
  CURSOR_FULL_CROSSHAIR      = 0x0066

  RED   = 0
  GREEN = 1
  BLUE  = 2

  KEY_REPEAT_OFF     = 0
  KEY_REPEAT_ON      = 1
  KEY_REPEAT_DEFAULT = 2
  JOYSTICK_BUTTON_A  = 1
  JOYSTICK_BUTTON_B  = 2
  JOYSTICK_BUTTON_C  = 4
  JOYSTICK_BUTTON_D  = 8
)

func Init (s string) { a, cs := C.int(1), C.CString(s); C.glutInit (&a, &cs) } // call C.free(unsafe.Pointer(s))
func WindowPosition (x, y int) { C.glutInitWindowPosition (C.int(x), C.int(y)) }
func WindowSize (w, h uint) { C.glutInitWindowSize (C.int(w), C.int(h)) }
func DisplayMode (m uint) { C.glutInitDisplayMode (C.uint(m)) }
func InitDisplayString (s string) { C.glutInitDisplayString (C.CString(s)); } // call C.free(unsafe.Pointer(s))
func MainLoop() { C.glutMainLoop(); }

func CreateWindow (s string) { C.glutCreateWindow (C.CString(s)) }
func CreateSubWindow (window, x, y, w, h int) { C.glutCreateSubWindow (C.int(window), C.int(x), C.int(y), C.int(w), C.int(h)); }
func DestroyWindow (w int) { C.glutDestroyWindow(C.int(w)); }
func SetWindow (w int) { C.glutSetWindow (C.int(w)); }
func GetWindow() int { w := C.getWindow(); return int(w) }
func SetWindowTitle (s string) { C.glutSetWindowTitle (C.CString(s)); }
func SetIconTitle (s string) { C.glutSetIconTitle (C.CString(s)); }
func ReshapeWindow (w, h int) { C.glutReshapeWindow (C.int(w), C.int(h)); }
func PositionWindow (x, y int) { C.glutPositionWindow (C.int(x), C.int(y)); }
func ShowWindow() { C.glutShowWindow(); }
func HideWindow() { C.glutHideWindow(); }
func IconifyWindow() { C.glutIconifyWindow(); }
func PushWindow() { C.glutPushWindow(); }
func PopWindow() { C.glutPopWindow(); }
func FullScreen() { C.glutFullScreen(); }

func PostWindowRedisplay (w int) { C.glutPostWindowRedisplay(C.int(w)); }
func PostRedisplay() { C.glutPostRedisplay() }
func SwapBuffers() { C.glutSwapBuffers(); }

func WarpMouse (x, y int) { C.glutWarpPointer (C.int(x), C.int(y)); }
func SetMousePointer (c int) { C.glutSetCursor (C.int(c)); }

func Cube (a float64) { C.glutSolidCube (C.GLdouble(a)); }
func WireCube (a float64) { C.glutWireCube (C.GLdouble(a)); }
func Sphere (r float64, slices, stacks uint) { C.glutSolidSphere (C.GLdouble(r), C.GLint (slices), C.GLint (stacks)); }
func WireSphere (r float64, slices, stacks uint) { C.glutWireSphere (C.GLdouble(r), C.GLint (slices), C.GLint (stacks)); }
func Cone (b, h, slices, stacks float64) { C.glutSolidCone (C.double(b), C.double(h), C.GLint (slices), C.GLint(stacks) ); }
func WireCone (b, h, slices, stacks float64) { C.glutWireCone (C.double(b), C.double(h), C.GLint (slices), C.GLint(stacks) ); }
func Torus (ir, or, s, r float64) { C.glutSolidTorus (C.double(ir), C.double(or), C.GLint(s), C.GLint(r)); }
func WireTorus (ir, or, s, r float64) { C.glutWireTorus (C.double(ir), C.double(or), C.GLint(s), C.GLint(r)); }
func Dodecahedron() { C.glutSolidDodecahedron(); }
func WireDodecahedron() { C.glutWireDodecahedron(); }
func Octahedron() { C.glutSolidOctahedron(); }
func WireOctahedron() { C.glutWireOctahedron(); }
func Tetrahedron() { C.glutSolidTetrahedron(); }
func WriteTetrahedron() { C.glutWireTetrahedron(); }
func Icosahedron() { C.glutSolidIcosahedron(); }
func WireIcosahedron() { C.glutWireIcosahedron(); }
func Teapot (a float64) { C.glutSolidTeapot (C.double(a)); }
func WireTeapot (a float64) { C.glutWireTeapot (C.double(a)); }
