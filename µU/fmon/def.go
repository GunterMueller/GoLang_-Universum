package fmon

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  FarMonitor interface { // x always denotes the calling object.

// Pre: i < number of monitor functions of x.
// The monitor function fs(_, i) of x is executed on the server;
// if necessary, the calling process was delayed until ps(a, i) == true.
// The value of a is sent to the server with this call.
// It returns the value (of the type of x), that the server returns.
  F (a any, i uint) any

// All net channels used by x are closed.
  Fin()
}

// Pre: fs and ps are defined in their second argument for all i < n.
//      h is contained in /etc/hosts or denotes a DNS-resolvable host,
//      p is not used by any network service.
//      If s == true, New is called in a process on the host h.
// Returns a new far monitor with n monitor operations.
// Its type is the type of a, i.e. objects of this type are passed
// between server and clients.
// For a == nil the server returns upon a monitor call a byte stream
// (object of type Stream); in this case the caller is responsible
// for decoding that stream into an object of the type of the object,
// that he sent as first argument in his monitor function call F.
// fs(_,i) for i < n is the i-th monitor function and ps(_, i) is the
// corresponding predicate determining whether the monitor functions
// can be executed.
// h is the server executing the monitor calls and p is the port
// used by the TCP-IP connection between the server and the clients;
// the needed net channels are opened.
// The far monitor runs as server, iff s == true; otherwise as client.
func New (a any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h string, p uint16, s bool) FarMonitor {
  return new_(a,n,fs,ps,h,p,s)
}

// See above. Additionally, st is executed by the server before it starts serving.
func New1 (a any, n uint, fs FuncSpectrum, ps PredSpectrum,
           h string, p uint16, s bool, stmt Stmt) FarMonitor {
  return new1(a, n, fs, ps, h, p, s, stmt)
}

// Spec is trade secret.
func New2 (a, b any, n uint, fs FuncSpectrum, ps PredSpectrum,
           h string, p uint16, s bool) FarMonitor {
  return new2(a, b, n, fs, ps, h, p, s)
}
