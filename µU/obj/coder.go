package obj

// (c) Christian Maurer   v. 200908 - license see µU.go

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

// Returns true, iff a implements Coder.
func IsCoder (a Any) bool { return isCoder(a) }

// Returns Codelen(int(0)), equal to Codelen(uint(0)).
func C0() uint { return c0 }

// Pre: a is atomic or implements Object.
// Returns the codelength of a.
func Codelen (a Any) uint { return codelen(a) }

// Pre: a is atomic or implements Object.
// Returns a as encoded byte sequence.
func Encode (a Any) Stream { return encode(a) }

// Pre: a is atomic or streamic or implements Object.
//      b is a encoded byte sequence.
// Returns the object, that was encoded in b.
func Decode (a Any, b Stream) Any { return decode(a,b) }

// Returns a byte sequence of length 16, that encodes a, b, c, d.
func Encode4 (a, b, c, d uint32) Stream { return encode4 (a,b,c,d) }

// Pre: len(s) == 16; b encodes 4 numbers of type uint32.
// Returns those 4 numbers.
func Decode4 (s Stream) (uint32, uint32, uint32, uint32) { return decode4(s) }

// Returns true, iff a is atomic or implements Coder.
func AtomicOrCoder (a Any) bool { return Atomic(a) || isCoder(a) }

// Pre: For each i < len(a): c[i] == Codelen(a[i]).
// Returns a encoded as Stream.
func Encodes (a AnyStream, c []uint) Stream { return encodes (a,c) }

// Pre: For each i < len(a): c[i] == Codelen(a[i]).
//      s is an encoded AnyStream.
// a is the Anystream, that was encoded in s.
func Decodes (s Stream, a AnyStream, c []uint) { decodes (s,a,c) }
