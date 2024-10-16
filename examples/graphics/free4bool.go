package main

// (c) Christian Maurer   v. 230326 - license see µU.go

// >>> the free Boolean algebra over two elements a and b
//     ("+" means union and "·" means intersection)

import (
  "µU/scr"
  "µU/gl"
  "µU/col"
)

func draw() {
  gl.ClearColour (col.FlashWhite())
  gl.Clear()

  const r = 0.075
  const h = 2*r

  gl.Colour (col.Black())
  gl.Sphere ( 0, 0, 2, r)      // 1
  gl.Colour (col.Red())

  gl.Colour (col.Cyan())
  gl.Cone   ( 1,  1, 1, r, -h) //  a+ b
  gl.Colour (col.Magenta())
  gl.Cone   (-1,  1, 1, r, -h) // ¬a+ b
  gl.Colour (col.Orange())
  gl.Cone   (-1, -1, 1, r, -h) // ¬a+¬b
  gl.Colour (col.LightGreen())
  gl.Cone   ( 1, -1, 1, r, -h) //  a+¬b

  gl.Colour (col.Green())
  gl.Sphere ( 2,  0, 0, r)     //  a
  gl.Colour (col.Blue())
  gl.Sphere ( 0,  2, 0, r)     //  b
  gl.Colour (col.Red())
  gl.Sphere (-2,  0, 0, r)     // ¬a
  gl.Colour (col.Yellow())
  gl.Sphere ( 0, -2, 0, r)     // ¬b

  gl.Colour (col.Cyan())
  gl.Cone ( 1,  1, -1, r, h)   //  a· b
  gl.Colour (col.Magenta())
  gl.Cone (-1,  1, -1, r, h)   // ¬a· b
  gl.Colour (col.Orange())
  gl.Cone (-1, -1, -1, r, h)   // ¬a·¬b
  gl.Colour (col.LightGreen())
  gl.Cone ( 1, -1, -1, r, h)   //  a·¬b

  gl.Colour (col.Gray())
  gl.Cube (.25,-.25, 0, r)     // (a+b)·(¬a+¬b)
  gl.Colour (col.Pink())
  gl.Cube (-.25,.25, 0, r)     // (a·b)+(¬a·¬b)

  gl.Colour (col.Black())
  gl.Sphere (0, 0, -2, r)      // 0

  gl.Colour (col.Black())

  gl.Line (0, 0, 2,  1,  1,  1) // 1 >  a+ b
  gl.Line (0, 0, 2, -1,  1,  1) // 1 > ¬a+ b
  gl.Line (0, 0, 2, -1, -1,  1) // 1 > ¬a+¬b
  gl.Line (0, 0, 2,  1, -1,  1) // 1 >  a+¬b

  gl.Line ( 1,  1,  1,  2,  0,  0) //  a+ b >  a
  gl.Line ( 1,  1,  1,  0,  2,  0) //  a+ b >  b
  gl.Line (-1,  1,  1, -2,  0,  0) // ¬a+ b > ¬a
  gl.Line (-1,  1,  1,  0,  2,  0) // ¬a+ b >  b
  gl.Line (-1, -1,  1,  0, -2,  0) // ¬a+¬b > ¬a
  gl.Line (-1, -1,  1, -2,  0,  0) // ¬a+¬b > ¬b
  gl.Line ( 1, -1,  1,  2,  0,  0) //  a+¬b >  a
  gl.Line ( 1, -1,  1,  0, -2,  0) //  a+¬b > ¬b

  gl.Line ( 1,  1, -1,  2,  0,  0) //  a· b <  a
  gl.Line ( 1,  1, -1,  0,  2,  0) //  a· b <  b
  gl.Line (-1,  1, -1, -2,  0,  0) // ¬a· b < ¬a
  gl.Line (-1,  1, -1,  0,  2,  0) // ¬a· b <  b
  gl.Line (-1, -1, -1, -2,  0,  0) // ¬a·¬b < ¬a
  gl.Line (-1, -1, -1,  0, -2,  0) // ¬a·¬b < ¬b
  gl.Line ( 1, -1, -1,  2,  0,  0) //  a·¬b <  a
  gl.Line ( 1, -1, -1,  0, -2,  0) //  a·¬b < ¬b

  gl.Line ( 0.25,-.25,  0,  1,  1,  1) // (a+b)·(¬a+¬b) <  a+ b
  gl.Line ( 0.25,-.25,  0, -1, -1,  1) // (a+b)·(¬a+¬b) < ¬a+¬b
  gl.Line ( 0.25,-.25,  0,  1, -1, -1) // (a+b)·(¬a+¬b) >  a·¬b
  gl.Line ( 0.25,-.25,  0, -1,  1, -1) // (a+b)·(¬a+¬b) > ¬a· b

  gl.Line (-0.25, .25,  0,  1, -1,  1) // (a·b)+(¬a·¬b) <  a+¬b
  gl.Line (-0.25, .25,  0, -1,  1,  1) // (a·b)+(¬a·¬b) < ¬a+ b
  gl.Line (-0.25, .25,  0,  1,  1, -1) // (a·b)+(¬a·¬b) >  a· b
  gl.Line (-0.25, .25,  0, -1, -1, -1) // (a·b)+(¬a·¬b) > ¬a·¬b

  gl.Line (0, 0, -2,  1,  1, -1) // 0 <  a· b
  gl.Line (0, 0, -2, -1,  1, -1) // 0 < ¬a· b
  gl.Line (0, 0, -2, -1, -1, -1) // 0 < ¬a·¬b
  gl.Line (0, 0, -2,  1, -1, -1) // 0 <  a·¬b
}

func main() {
  s := scr.NewWH (0, 0, 1150, 1150); defer scr.Fin()
  s.Cls()
  s.Go (draw, 3, -4, 0, 0, 0, 0, 0, 0, 1)
}
