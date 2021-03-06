package nchan

// (c) Christian Maurer   v. 170507 - license see µu.go

import (
  . "µu/obj" // see also µu/host
  "µu/host"
)
const
  Port0 = uint16(50000) // uint16(1<<16 - 1<<14) // first private port (= 49152)
type
  NetChannel interface { // Channels for passing objects over the net

  Chan() (chan Any, chan Any)

// Pre: a is of the type of x.
// a is sent on x (resp. if x is a 1:n channel and
// the calling process is a server, on the actual subchannel of x)
// to the communication partner of the calling process.
  Send (a Any)
//  SendM (a ...Any)

// Returns a slice of bytes, if x was created by New with nil as first argument.
// In this case, the client is responsible for decoding that slice,
// according to the type of what was sent.
// Otherwise, i.e. if x is bound to a type:
// Returns the object of the type of x, that was received on x (resp. if
// x is a 1:n channel and the calling process is a server, on the actual
// subchannel of x) from the communication partner, if that was received;
// returns an empty object otherwise.
// The calling process was blocked, until an object was received.
  Recv() Any
//  RecvM() []Any

// The port used by x is not used by a network service on the calling host.
  Fin()
}

// h0 always denotes the calling host (running the calling process).

// For all constructors for the first parameter a the following holds:
//      For a == nil, arbitrary objects of Codelen <= 1<<32 can be passed.
//      In this case, calls of Recv() return only slices of bytes, so
//      it is up to the receiver to Decode the object wanted to receive.
//
// Pre: h is in /etc/hosts or resolvable per DNS (! h.Empty());
//      h is different from h0.
//      p > 0; p is not used on h0 or h by a network service.
//      The communication partner calls New with
//      - an object of the same type as the type of a as 1st argument,
//      - with h0 as 2nd argument and
//      - an identical value of the 3rd argument.
// Returns a asynchronous 1:1 channel for messages of the type of a
// between h0 and h over port p.
// p is now used on h0 and h by a network service.
func NewD (a Any, h host.Host, p uint16) NetChannel { return newd(a,h,p) }

// Pre: h is in /etc/hosts or resolvable per DNS (! h.Empty()).
//      me != i; me is the identity of h0 and i is the identity of h
//      (me, i < number of hosts involved).
//      p > 0; p is not used on h0 or h by a network service.
//      The communication partner calls New with
//      - an object of the same type as the type of a as 1st argument,
//      - with the host of the calling process as 4th argument.
//      - with the values of me and i reversed, i.e. me/i as 3rd/2nd argument and
//      - an identical value of the 5th argument.
// Returns a asynchronous 1:1 channel for messages of the type of a
// between h0 and h over port p.
// p is now used on h0 and h by a network service.
func New (a Any, me, i uint, h host.Host, p uint16) NetChannel { return new_(a,me,i,h,p) }

// See above function. To be called in the constructor of a far monitor.
// h is the server; s == true, if the calling process is the serving monitor.
// i and o are the in- and outgoing channels for the internal communication
// of the calling far monitor.
func NewF (a Any, h host.Host, p uint16, s bool /*, i, o chan Any */) NetChannel { return newf(a,h,p,s /*,i,o*/) }

// Note: In case of consecutive calls of New you have to keep
//       the correct pairing in both programs to avoid deadlocks!

// Pre: i, j < n, a < 2.
// Returns different port numbers 0..a * NPorts(n)-1.
func Port (n, i, j, a uint) uint16 { return port(n,i,j,a) }
func NPorts (n, a uint) uint { return nPorts(n,a) }
