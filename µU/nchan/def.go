package nchan

// (c) Christian Maurer   v. 220420 - license see µU.go

const
//  Port0 = uint16(1<<16 - 1<<14) // == 49152 (first private port)
  Port0 = uint16(50000)
type
  NetChannel interface { // Channels for passing objects over the net

  Chan() (chan any, chan any)

// Pre: a is of the type of x.
// Returns nil, if a is sent on x (resp. if x is a 1:n channel and
// the calling process is a server, on the actual subchannel of x)
// to the communication partner of the calling process.
// Returns otherwise an appropriate error.
  Send (a any) error

// Returns a slice of bytes, if x was created by New with nil as first argument.
// In this case, the client is responsible for decoding that slice,
// according to the type of what was sent.
// Otherwise, i.e. if x is bound to a type:
// Returns the object of the type of x, that was received on x (resp. if
// x is a 1:n channel and the calling process is a server, on the actual
// subchannel of x) from the communication partner, if that was received;
// returns an empty object otherwise.
// The calling process was blocked, until an object was received.
  Recv() any

// The port used by x is not used by a network service on the calling host.
  Fin()
}

// h0 always denotes the name of the host running the calling process.

// For both constructors for the first parameter a the following holds:
//      For a == nil, arbitrary objects of Codelen <= 1<<32 can be passed.
//      In this case, calls of Recv() return only slices of bytes, so
//      it is up to the receiver to Decode the object wanted to receive.

// Pre: h is contained in /etc/hosts or resolvable per DNS.
//      me != i; me is the identity of h0 and i is the identity of h
//      (me, i < number of hosts involved).
//      p > 0; p is not used on h0 or h by a network service.
//      The communication partner calls New with
//      - an object of the same type as the type of a as 1st argument
//        and an identical value of the port,
//      - the values of me and i reversed and h0 as 4th argument.
// Returns a asynchronous 1:1 channel for messages of the type of a
// between h0 and h over port p.
// p is now used on h0 and h by a network service.
func New (a any, me, i uint, h string, p uint16) NetChannel { return new_(a,me,i,h,p) }

// See above function. To be called in the constructor of a far monitor.
// h is the server (s = true), if the calling process is the serving monitor.
func NewN (a any, h string, p uint16, s bool) NetChannel { return newn(a,h,p,s) }

// Note: In case of consecutive calls of New you have to keep
//       the correct pairing in both programs to avoid deadlocks!

// Pre: i, j < n; a < 2.
// Returns different port numbers 0..NPorts(n,a).
func Port (n, i, j, a uint) uint16 { return port(n,i,j,a) }
func NPorts (n, a uint) uint { return nPorts(n,a) }
