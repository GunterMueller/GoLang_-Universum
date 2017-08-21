package obj

// (c) murus.org  v. 170701 - license see murus.go

type
  Coder interface {

// Returns the number of bytes, that are needed
// to serialize x uniquely revertibly.
  Codelen() uint

// x.Eq (x.Decode (x.Encode())
  Encode() []byte // []byte

// Pre: b is a result of y.Encode() for some y
//      of the same type as x.
// x.Eq(y); x.Encode() == b, i.e. those slices coincide.
  Decode (b []byte) // []byte)
}
var
  C0 = Codelen(uint(0))

// Pre: a is atomic or implements Object.
// Returns the codelength of a.
func Codelen (a Any) uint { return codelen(a) }

// Pre: a is atomic or implements Object.
// Returns a as encoded byte sequence.
func Encode (a Any) []byte { return encode(a) }

// Pre: a is atomic or implements Object.
//      bs is a encoded byte sequence.
// Returns the object, that was encoded.
func Decode (a Any, bs []byte) Any { return decode(a,bs) }

// Returns a byte sequence of length 16,
// that encodes a, b, c, d.
func Encode4 (a, b, c, d uint32) []byte { return encode4 (a,b,c,d) }

// Pre: len(bs) == 16;
//      bs encodes 4 numbers of type uint32.
// Returns those 4 numbers.
func Decode4 (bs []byte) (uint32, uint32, uint32, uint32) { return decode4(bs) }
