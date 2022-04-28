package rpc

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  RPC interface {

// Pre: i < number of RPC functions of x.
// The value of a is sent to the server with this call;
// the function f(_, i) of x is executed on the server.
// F returns the value computed by the remote server.
  F (a any, i uint) any

// All net channels used by x are closed.
  Fin()
}

// Pre: f is defined in its second argument for all i < n.
//      h is contained in /etc/hosts or denotes a DNS-resolvable host,
//      p is not used by any network service.
//      If s == false, New was called in a process on the host h.
// Returns a new rpc with n functions.
// Its input type is the type of a and its output type is the type of b,
// i.e. objects of type a are passed from clients to the server as arguments
// and objects of type b are passed from the server to the clients as results.
// f(_,i) for i < n is the i-th RPC function,
// h is the server executing the remote calls and p is the port
// used by the TCP-IP connection between the RPC server and the clients;
// the needed net channels are opened.
// The rpc runs as server, iff s == true; otherwise as client.
func New (a, b any, n uint, h string, p uint16, s bool, f FuncSpectrum) RPC {
  return new_(a, b, n, h, p, s, f)
}
