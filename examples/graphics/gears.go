package main

// #cgo LDFLAGS: -lGL -lglut
// #include <stdlib.h> 
// #include <GL/gl.h> 
// #include <GL/freeglut.h>
import
  "C"
import (
  "os"
  "glut"
  "math"
)
/**
  Draw a gear wheel.  You'll probably want to call this function when
  building a display list since we do a lot of trig here.
 
  Input:  inner_radius - radius of hole at center
          outer_radius - radius at center of teeth
          width - width of gear
          teeth - number of teeth
          tooth_depth - depth of tooth
 **/

func gear (inner_radius, outer_radius, width C.GLfloat, teeth C.GLint, tooth_depth C.GLfloat) {
  r0 := float64(inner_radius)
  r1 := float64(outer_radius - tooth_depth) / 2.0
  r2 := float64(outer_radius + tooth_depth) / 2.0
  da := 2.0 * math.Pi / float64(teeth) / 4.0
  C.glShadeModel (C.GL_FLAT)
  C.glNormal3f(0.0, 0.0, 1.0)
  C.glBegin(C.GL_QUAD_STRIP)
  for i := C.GLint(0); i <= teeth; i++ {
    angle := (float64(i) * 2.0 * math.Pi) / float64(teeth)
    C.glVertex3f (C.GLfloat(r0 * math.Cos(angle)), C.GLfloat(r0 * math.Sin(angle)), width * 0.5)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle)), C.GLfloat(r1 * math.Sin(angle)), width * 0.5)
    C.glVertex3f (C.GLfloat(r0 * math.Cos(angle)), C.GLfloat(r0 * math.Sin(angle)), width * 0.5)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle + 3 * da)),
                  C.GLfloat(r1 * math.Sin(angle + 3 * da)), width * 0.5)
  }
  C.glEnd()
  C.glBegin(C.GL_QUADS)
  da = 2.0 * math.Pi / float64(teeth) / 4.0
  for i := C.GLint(0); i < teeth; i++ {
    angle := (float64(i) * 2.0 * math.Pi) / float64(teeth)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle)), C.GLfloat(r1 * math.Sin(angle)), width * 0.5)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + da)),
                  C.GLfloat(r2 * math.Sin(angle + da)), width * 0.5)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + 2 * da)),
                  C.GLfloat(r2 * math.Sin(angle + 2 * da)), width * 0.5)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle + 3 * da)),
                  C.GLfloat(r1 * math.Sin(angle + 3 * da)), width * 0.5)
  }
  C.glEnd()
  C.glNormal3f(0.0, 0.0, -1.0)
  C.glBegin(C.GL_QUAD_STRIP)
  for i := C.GLint(0); i <= teeth; i++ {
    angle := float64(i) * 2.0 * math.Pi / float64(teeth)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle)), C.GLfloat(r1 * math.Sin(angle)), -width * 0.5)
    C.glVertex3f (C.GLfloat(r0 * math.Cos(angle)), C.GLfloat(r0 * math.Sin(angle)), -width * 0.5)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle + 3 * da)),
                  C.GLfloat(r1 * math.Sin(angle + 3 * da)), -width * 0.5)
    C.glVertex3f (C.GLfloat(r0 * math.Cos(angle)),
                  C.GLfloat(r0 * math.Sin(angle)), -width * 0.5)
  }
  C.glEnd()
  C.glBegin(C.GL_QUADS)
  da = 2.0 * math.Pi / float64(teeth) / 4.0
  for i := C.GLint(0); i < teeth; i++ {
    angle := float64(i) * 2.0 * math.Pi / float64(teeth)

    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle + 3 * da)),
                  C.GLfloat(r1 * math.Sin(angle + 3 * da)), -width * 0.5)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + 2 * da)),
                  C.GLfloat(r2 * math.Sin(angle + 2 * da)), -width * 0.5)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + da)),
                  C.GLfloat(r2 * math.Sin(angle + da)), -width * 0.5)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle)),
                  C.GLfloat(r1 * math.Sin(angle)), -width * 0.5)
  }
  C.glEnd()
  C.glBegin(C.GL_QUAD_STRIP)
  for i := C.GLint(0); i < teeth; i++ {
    angle := float64(i) * 2.0 * math.Pi / float64(teeth)

    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle)),
                  C.GLfloat(r1 * math.Sin(angle)), width * 0.5)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle)),
                  C.GLfloat(r1 * math.Sin(angle)), -width * 0.5)
    u := r2 * math.Cos(angle + da) - r1 * math.Cos(angle)
    v := r2 * math.Sin(angle + da) - r1 * math.Sin(angle)
    lenn := math.Sqrt(u * u + v * v)
    u /= lenn
    v /= lenn
    C.glNormal3f (C.GLfloat(v), C.GLfloat(-u), 0.0)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + da)),
                  C.GLfloat(r2 * math.Sin(angle + da)), width * 0.5)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + da)),
                  C.GLfloat(r2 * math.Sin(angle + da)), -width * 0.5)
    C.glNormal3f (C.GLfloat(math.Cos(angle)), C.GLfloat(math.Sin(angle)), 0.0)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + 2 * da)),
                  C.GLfloat(r2 * math.Sin(angle + 2 * da)), width * 0.5)
    C.glVertex3f (C.GLfloat(r2 * math.Cos(angle + 2 * da)),
                  C.GLfloat(r2 * math.Sin(angle + 2 * da)), -width * 0.5)
    u = r1 * math.Cos(angle + 3 * da) - r2 * math.Cos(angle + 2 * da)
    v = r1 * math.Sin(angle + 3 * da) - r2 * math.Sin(angle + 2 * da)
    C.glNormal3f (C.GLfloat(v), C.GLfloat(-u), 0.0)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle + 3 * da)),
                  C.GLfloat(r1 * math.Sin(angle + 3 * da)), width * 0.5)
    C.glVertex3f (C.GLfloat(r1 * math.Cos(angle + 3 * da)),
                  C.GLfloat(r1 * math.Sin(angle + 3 * da)), -width * 0.5)
    C.glNormal3f (C.GLfloat(math.Cos(angle)), C.GLfloat(math.Sin(angle)), 0.0)
  }
  C.glVertex3f (C.GLfloat(r1 * math.Cos(0)), C.GLfloat(r1 * math.Sin(0)), width * 0.5)
  C.glVertex3f (C.GLfloat(r1 * math.Cos(0)), C.GLfloat(r1 * math.Sin(0)), -width * 0.5)
  C.glEnd()
  C.glShadeModel(C.GL_SMOOTH)
  C.glBegin(C.GL_QUAD_STRIP)
  for i := C.GLint(0); i <= teeth; i++ {
    angle := float64(i) * 2.0 * math.Pi / float64(teeth)
    C.glNormal3f (C.GLfloat(-math.Cos(angle)), C.GLfloat(-math.Sin(angle)), 0.0)
    C.glVertex3f (C.GLfloat(r0 * math.Cos(angle)), C.GLfloat(r0 * math.Sin(angle)), -width * 0.5)
    C.glVertex3f (C.GLfloat(r0 * math.Cos(angle)), C.GLfloat(r0 * math.Sin(angle)), width * 0.5)
  }
  C.glEnd()
}

var (
  view_rotx = C.GLfloat(20.0)
  view_roty = C.GLfloat(30.0)
  view_rotz = C.GLfloat(0.0)
  gear1, gear2, gear3 C.GLint
  angle = C.GLfloat(0.0)
  limit C.GLuint
  count = C.GLuint(1)
)

func draw() {
  C.glClear (C.GL_COLOR_BUFFER_BIT | C.GL_DEPTH_BUFFER_BIT)
  C.glPushMatrix()
  C.glRotatef (view_rotx, 1.0, 0.0, 0.0)
  C.glRotatef (view_roty, 0.0, 1.0, 0.0)
  C.glRotatef (view_rotz, 0.0, 0.0, 1.0)
  C.glPushMatrix()
  C.glTranslatef (-3.0, -2.0, 0.0)
  C.glRotatef (angle, 0.0, 0.0, 1.0)
  C.glCallList (C.GLuint(gear1))
  C.glPopMatrix()
  C.glPushMatrix()
  C.glTranslatef(3.1, -2.0, 0.0)
  C.glRotatef (-2.0 * angle - 9.0, 0.0, 0.0, 1.0)
  C.glCallList (C.GLuint(gear2))
  C.glPopMatrix()
  C.glPushMatrix()
  C.glTranslatef (-3.1, 4.2, 0.0)
  C.glRotatef (-2.0 * angle - 25.0, 0.0, 0.0, 1.0)
  C.glCallList (C.GLuint(gear3))
  C.glPopMatrix()
  C.glPopMatrix()
  glut.SwapBuffers()
  count++
  if count == limit {
    os.Exit(0)
  }
}

func idle() {
  angle += 2.0
  glut.PostRedisplay()
}

func key (k byte, x, y int) {
  switch k {
  case 'z':
    view_rotz += 5.0
    break
  case 'Z':
    view_rotz -= 5.0
    break
  case 27:
    os.Exit(0)
    break
  default:
    return
  }
  glut.PostRedisplay()
}

func special (k, x, y int) {
  switch k {
  case C.GLUT_KEY_UP:
    view_rotx += 5.0
    break
  case C.GLUT_KEY_DOWN:
    view_rotx -= 5.0
    break
  case C.GLUT_KEY_LEFT:
    view_roty += 5.0
    break
  case C.GLUT_KEY_RIGHT:
    view_roty -= 5.0
    break
  default:
    return
  }
  glut.PostRedisplay()
}

func null() { }

func reshape (width, height int) {
  h := C.GLdouble(C.GLfloat(height) / C.GLfloat(width))
  C.glViewport (0, 0, C.GLsizei(width), C.GLsizei(height))
  C.glMatrixMode (C.GL_PROJECTION)
  C.glLoadIdentity()
  C.glFrustum (-1.0, 1.0, -h, h, 5.0, 60.0)
  C.glMatrixMode (C.GL_MODELVIEW)
  C.glLoadIdentity()
  C.glTranslatef (0.0, 0.0, -40.0)
}

func init_() {
  pos := [4]C.GLfloat {5.0, 5.0, 10.0, 0.0}
  red := [4]C.GLfloat {0.8, 0.1, 0.0, 1.0}
  green := [4]C.GLfloat {0.0, 0.8, 0.2, 1.0}
  blue := [4]C.GLfloat {0.2, 0.2, 1.0, 1.0}
  C.glLightfv (C.GL_LIGHT0, C.GL_POSITION, &pos[0])
  C.glEnable (C.GL_CULL_FACE)
  C.glEnable (C.GL_LIGHTING)
  C.glEnable (C.GL_LIGHT0)
  C.glEnable (C.GL_DEPTH_TEST)
  gear1 = C.GLint(C.glGenLists(1))
  C.glNewList (C.GLuint(gear1), C.GL_COMPILE)
  C.glMaterialfv (C.GL_FRONT, C.GL_AMBIENT_AND_DIFFUSE, &red[0])
  gear (1.0, 4.0, 1.0, 20, 0.7)
  C.glEndList()
  gear2 = C.GLint(C.glGenLists(1))
  C.glNewList (C.GLuint(gear2), C.GL_COMPILE)
  C.glMaterialfv (C.GL_FRONT, C.GL_AMBIENT_AND_DIFFUSE, &green[0])
  gear (0.5, 2.0, 2.0, 10, 0.7)
  C.glEndList()
  gear3 = C.GLint(C.glGenLists(1))
  C.glNewList (C.GLuint(gear3), C.GL_COMPILE)
  C.glMaterialfv (C.GL_FRONT, C.GL_AMBIENT_AND_DIFFUSE, &blue[0])
  gear (1.3, 2.0, 0.5, 10, 0.7)
  C.glEndList()
  C.glEnable (C.GL_NORMALIZE)
}

func visible (vis int) {
  if vis == C.GLUT_VISIBLE {
    glut.IdleFunc (idle)
  } else {
    glut.IdleFunc (null)
  }
}

func main() {
  limit = 0
  glut.InitDisplayMode (C.GLUT_RGB | C.GLUT_DEPTH | C.GLUT_DOUBLE)
  glut.CreateWindow("Gears")
  init_()
  glut.DisplayFunc (draw)
  glut.ReshapeFunc (reshape)
  glut.KeyboardFunc (key)
  glut.SpecialFunc (special)
  glut.VisibilityFunc (visible)
  glut.MainLoop()
}
