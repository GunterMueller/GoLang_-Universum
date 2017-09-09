package img

// (c) Christian Maurer   v. 140130 - license see murus.go

// The pixels in the rectangle on the screen with top left
// corner (x,y) of width w and height h are written to files
// in the pnm- and png-format under the names "n.ppm"/"n.png".
func Put (n string, x, y, w, h uint) { put(n,x,y,w,h) }

// The pixels of the whole screen are written to files
// in the pnm- and png-format under the names "n.ppm"/"n.png".
func Put1 (n string) { put1(n) }

// Returns (0, 0), if there is no file with the name "n.ppm"
// containing an image in the ppm-format; returns otherwise
// width and height of this image.
func Size (n string) (uint, uint) { return size_(n) }

// If there is no file with the name "n.ppm" containing an image
// in the ppm-format, nothing has happened. Otherwise, that
// image is drawn to the screen with (x, y) as top left corner.
func Get (n string, x, y uint) { get(n,x,y) }

// The pixels in the rectangle on the screen with top left
// corner (x,y) of width w and height h are written to a file
// under the name "tmp.ppm" and are printed.
func Print (x, y, w, h uint) { print_(x,y,w,h) }

// All pixels on the screen are written to a file
// under the name "tmp.ppm" and are printed.
func Print1 () { print1() }
