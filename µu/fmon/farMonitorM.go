package fmon

// (c) Christian Maurer   v. 170923 - license see mu.

import (
  . "µu/ker"
  . "µu/obj"
  "µu/host"
  "µu/nchan"
)
type
  farMonitorM struct {
                     Any "type of objects the monitor functions operate on"
                     uint "number of monitor functions"
                  nr []uint
                  ok [][]bool
                  ch [][]nchan.NetChannel
             in, out [][]chan Any
                     FuncSpectrum; PredSpectrum
                host host.Host
                     bool "true iff the monitor is a server"
                     }
func newM(a Any, n, j uint, nr []uint, fs FuncSpectrum, ps PredSpectrum,
          h host.Host, p []uint16, s bool) FarMonitorM {
  return newMS(a, n, j, nr, fs, ps, h, p, s, Null)
}

func newMS(a Any, n, j uint, nr []uint, fs FuncSpectrum, ps PredSpectrum,
           h host.Host, p []uint16, s bool, stmt Stmt) FarMonitorM {
  if n == 0 { Panic ("fmon.NewM must be called with 2nd arg > 0") }
  x := new(farMonitorM)
  x.Any = Clone(a)
  x.uint = n
  x.nr = nr
  nn := uint(len(nr))
  if uint(len(p)) != nn { Oops() }
  x.host = h
  x.bool = s
  x.ok = make([][]bool, x.uint)
  x.ch = make([][]nchan.NetChannel, x.uint)
  x.in, x.out = make([][]chan Any, x.uint), make([][]chan Any, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.ok[i] = make([]bool, nn)
    x.ch[i] = make([]nchan.NetChannel, nn)
    x.in[i], x.out[i] = make([]chan Any, nn), make([]chan Any, nn)
    for k, nr := range x.nr { // nr == x.nr[j]
      x.ok[i][k] = x.bool
      if ! x.bool { // client
        x.ok[i][k] = x.nr[j] == nr
      }
      if x.ok[i][k] {
        x.ch[i][k] = nchan.NewF (x.Any, h, p[k], s)
        x.in[i][k], x.out[i][k] = x.ch[i][k].Chan()
      }
    }
  }
  if ! x.bool {
    return x
  }
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  stmt()
  for ii := uint(0); ii < x.uint; ii++ {
    for kk, _ := range x.nr {
      if x.ok[ii][kk] {
        go func (i uint, k int) {
          for {
            select {
            case any, ok := <-When (x.PredSpectrum (x.Any, i), x.in[i][k]):
              if ok {
                any = x.FuncSpectrum (any, i)
                x.out[i][k] <- any
              } else {
//                println("client off")
              }
            default:
//              println("client off")
            }
            Msleep(100)
          }
          Msleep(100)
        }(ii, kk)
      }
    }
  }
  return x
}

func (x *farMonitorM) Fm (a Any, i, k uint) Any {
  if i > x.uint { panic("no monitor function"); return a }
  x.ch[i][k].Send (a)
  return x.ch[i][k].Recv()
}

func (x *farMonitorM) Fin() {
  for i := uint(0); i < x.uint; i++ {
    for k, _ := range x.nr {
      if x.ok[i][k] {
        close(x.in[i][k]); close(x.out[i][k])
        x.ch[i][k].Fin()
      }
    }
  }
}
