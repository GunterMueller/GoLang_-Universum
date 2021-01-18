package ppm

// (c) Christian Maurer   v. 210104 - license see µU.go

// The pixels in the screen are written to a file
// in the ppm-format under the name n.ppm.
func Put (n string) { put(n) }

// Returns (0, 0), if there is no file with the name n.ppm
// containing an image in the ppm-format; returns otherwise
// width and height of this image.
func Size (n string) (uint, uint) { return size_(n) }

// If there is no file with the name n.ppm containing an image
// in the ppm-format, nothing has happened. Otherwise, that
// image is drawn to the screen with (x, y) as top left corner.
func Get (n string, x, y uint) { get(n,x,y) }

// All pixels in the screen are written to a file
// with the name n.ppm and are printed.
func Print (n string) { print_(n) }
