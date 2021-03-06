package fmon

// (c) Christian Maurer   v. 170508 - license see murus.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt,
//     Kapitel 8, insbesondere Abschnitt 8.3

import (
  . "murus/obj"
  "murus/host"
)
type
  FarMonitor interface { // x always denotes the calling object.

// Pre: i < number of monitor functions of x.
// The calling process was delayed, if the predicate of x
// for the i-th monitor function is not true.
// The monitor function fs(_, i) is executed on the server;
// if necessary, the calling process was delayed until ps(a, i) == true.
// The value of a is sent to the server with this call.
// It returns the value (of the type of x), that the server returns.
  F (a Any, i uint) Any

// All net channels used by x are closed.
  Fin()
}

// Pre: fs and ps are defined in their second argument for all i < n.
//      h is by an entry in /etc/hosts or DNS-lookup reachable,
//      p is not used by any network service.
//      If s == true, New is called in a process on the host h.
// Returns a new far monitor with n monitor operations.
// Its type is the type of a, i.e. objects of this type are passed
// between server and clients.
// For a == nil the server returns upon a monitor call a byte stream
// (object of type []byte); in this case the caller is responsible
// for decoding that stream into an object of the type of the object,
// that he sent as first argument in his monitor function call F.
// fs(_,i) for i < n is the i-th monitor function and ps(_, i) is the
// corresponding predicate determining whether the monitor functions
// can be executed.
// h is the server executing the monitor calls and p is the port
// used by the TCP-IP connection between the server and the clients;
// the needed net channels are opened.
// The far monitor runs as server, iff s == true; otherwise as client.
func New (a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h host.Host, p uint16, s bool) FarMonitor {
  return new_(a, n, fs, ps, h, p, s)
}

// See above. Additionally, st is executed by the server before it starts serving.
func NewS (a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
           h host.Host, p uint16, s bool, stmt Stmt) FarMonitor {
  return newS(a, n, fs, ps, h, p, s, stmt)
}

type
  FarMonitorM interface {
  Fm (a Any, i, k uint) Any
  Fin()
}

// Pre: j is the id of the calling process,
//      nr are the ids of the neighbours of the calling process.
// See above. Additionally, for each server-client-pair there is an own channel.
func NewM (a Any, n, j uint, nr []uint, fs FuncSpectrum, ps PredSpectrum,
           h host.Host, p []uint16, s bool) FarMonitorM {
  return newM (a, n, j, nr, fs, ps, h, p, s)
}
