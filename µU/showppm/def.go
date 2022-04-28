package showppm

// (c) Christian Maurer   v. 220320 - license see µU.go

// Pre: There is a ppm-file with the name n+".ppm" in the actual directory.
// This file is converted into a file with the name n+".dat".
func Ppm2dat (n string) { ppm2dat (n) }

// Pre: Under X:
//        There is a file with the name n+"dat" in the actual directory,
//        that was constructed by a call of Ppm2dat.
//      In a tty-console:
//        There is a ppm-file with the name n+".ppm" in the actual directory.
// The ppm-Image in the ppm-file is written to the screen, starting at (x, y).
func ShowImage (n string, x, y int) { showImage (n,x,y) }
