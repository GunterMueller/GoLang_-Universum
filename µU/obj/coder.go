package obj

// (c) Christian Maurer   v. 171104 - license see µU.go

type
  Coder interface {

// Returns the number of bytes, that are needed
// to serialize x uniquely revertibly.
  Codelen() uint

// x.Eq (x.Decode (x.Encode())
  Encode() Stream

// Pre: b is a result of y.Encode() for some y
//      of the same type as x.
// x.Eq(y); x.Encode() == b, i.e. those slices coincide.
  Decode (Stream)
}
var
  C0 = Codelen(uint(0))

// Pre: a is atomic or implements Object.
// Returns the codelength of a.
func Codelen (a Any) uint { return codelen(a) }

// Pre: a is atomic or implements Object.
// Returns a as encoded byte sequence.
func Encode (a Any) Stream { return encode(a) }

// Pre: a is atomic or implements Object.
//      bs is a encoded byte sequence.
// Returns the object, that was encoded.
func Decode (a Any, bs Stream) Any { return decode(a,bs) }

// Returns a byte sequence of length 16,
// that encodes a, b, c, d.
func Encode4 (a, b, c, d uint32) Stream { return encode4 (a,b,c,d) }

// Pre: len(bs) == 16;
//      bs encodes 4 numbers of type uint32.
// Returns those 4 numbers.
func Decode4 (bs Stream) (uint32, uint32, uint32, uint32) { return decode4(bs) }

// Returns true, iff a implements Coder.
func IsCoder (a Any) bool { return isCoder(a) }

// Returns true, iff a is atomic or implements Coder.
func AtomicOrCoder (a Any) bool { return atomic(a) || isCoder(a) }
