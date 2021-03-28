package obj

// (c) Christian Maurer   v. 210106 - license see µU.go

const
  C0 = uint(8) // == Codelen(int(0)); == Codelen(uint(0))
type
  Coder interface {

// Returns the number of bytes, that are needed
// to serialise x uniquely revertibly.
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

// Returns a byte sequence of length 8, that encodes a, b, c, d.
func Encode4 (a, b, c, d uint16) Stream { return encode4(a,b,c,d) }

// Pre: len(s) == 8; s encodes 4 numbers of type uint16.
// Returns these 4 numbers.
func Decode4 (s Stream) (uint16, uint16, uint16, uint16) { return decode4(s) }

// Returns true, iff a is atomic or implements Coder.
func AtomicOrCoder (a Any) bool { return Atomic(a) || isCoder(a) }

// Pre: For each i < len(a): c[i] == Codelen(a[i]).
// Returns a encoded as Stream.
func Encodes (a AnyStream, c []uint) Stream { return encodes (a,c) }

// Pre: For each i < len(a): c[i] == Codelen(a[i]).
//      s is an encoded AnyStream.
// a is the Anystream, that was encoded in s.
func Decodes (s Stream, a AnyStream, c []uint) { decodes (s,a,c) }
