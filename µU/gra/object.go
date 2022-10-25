package gra

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/errh"
)

func (x *graph) numEdges (n *vertex) uint {
  c := uint(0)
  for nb := n.nbPtr; nb.nextNb != n.nbPtr; nb = nb.nextNb {
    c++
  }
  return c
}

func (x *graph) Eq (Y any) bool { // disgusting complexity
  y := x.imp (Y)
  if x.nVertices != y.nVertices || x.nEdges != y.nEdges ||
     ! TypeEq (x.vAnchor.any, y.vAnchor.any) ||
     ! TypeEq (x.eAnchor.any, y.eAnchor.any) {
    return false
  }
  ya := y.local // save
  eq := true
  loop:
  for xv := x.vAnchor.nextV; xv != x.vAnchor; xv = xv.nextV {
    if ! y.Ex (xv.any) {
      eq = false
      break
    }
    yv := y.local // y.local was changed
    if x.numEdges (xv) != y.numEdges (yv) {
      eq = false
      break
    }
    for xn := xv.nbPtr; xn.nextNb != xv.nbPtr; xn = xn.nextNb {
      for yn := yv.nbPtr; yn.nextNb != yv.nbPtr; yn = yn.nextNb {
        if yn.to == xn.to {
          aa := true
          if x.eAnchor.any != nil {
            if xn.edgePtr == nil { break }
            if yn.edgePtr == nil { break }
            aa = Eq (xn.edgePtr.any, yn.edgePtr.any)
          }
          if aa {
            break // next xnb
          } else {
            eq = false
            break loop
          }
        }
      }
    }
  }
  y.local = ya // restore
  return eq
}

// XXX The actual path is not copied.
func (x *graph) Copy (Y any) {
  y := x.imp(Y)
  x.Decode (y.Encode())
  x.SetWrite (y.Writes())
}

func (x *graph) Clone() any {
  y := new_(x.bool, x.vAnchor.any, x.eAnchor.any)
  y.Copy (x)
  return y
}

func (x *graph) Less (Y any) bool {
  return false
}

func (x *graph) Leq (Y any) bool {
  return false
}

func (x *graph) Empty() bool {
  return x.vAnchor.nextV == x.vAnchor
}

func delEdge (e *edge) {
  if e.nbPtr0 == nil { ker.Panic("gra.delEdge: e.nbPtr0 == nil") }
  e.prevE.nextE, e.nextE.prevE = e.nextE, e.prevE
  e.nbPtr0.prevNb.nextNb, e.nbPtr0.nextNb.prevNb = e.nbPtr0.nextNb, e.nbPtr0.prevNb // bug
  e.nbPtr1.prevNb.nextNb, e.nbPtr1.nextNb.prevNb = e.nbPtr1.nextNb, e.nbPtr1.prevNb
}

func delVertex (v *vertex) {
  n := v.nbPtr.nextNb
  for n != v.nbPtr {
    n = v.nbPtr
    n.to.predecessor = nil
    v.nbPtr = v.nbPtr.nextNb
  }
  v.prevV.nextV, v.nextV.prevV = v.nextV, v.prevV
  v = v.nextV
}

func (x *graph) Clr() {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    delEdge (e)
  }
  x.nEdges = 0
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    delVertex (v)
  }
  x.nVertices = 0
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.path, x.eulerPath = nil, nil
}

func (x *graph) Codelen() uint {
  c := uint(1) + 4
  if x.nVertices > 0 {
    for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
      c += 4 + Codelen(v.any) + 1
    }
    c += 3 * 4
    if x.nEdges > 0 {
      for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
        c += 4 + Codelen(e.any) + 1 + 2 * (4 + 1)
      }
    }
  }
  return c
}

func (x *graph) Encode() Stream {
  s := make (Stream, x.Codelen())
  s[0] = 0; if x.bool { s[0] = 1 }
  i, a := uint32(1), uint32(4)
  copy (s[i:i+a], Encode (x.nVertices))
  if x.nVertices == 0 { return s }
  i += a
  z := uint32(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    k := uint32(Codelen (v.any))
    copy (s[i:i+a], Encode (k))
    i += a
    s[i] = 0; if v.bool { s[i] = 1 }
    i++
    copy (s[i:i+k], Encode (v.any))
    i += k
    v.dist = z
    z++
  }
  copy (s[i:i+a], Encode (x.colocal.dist))
  i += a
  copy (s[i:i+a], Encode (x.local.dist))
  i += a
  copy (s[i:i+a], Encode (x.nEdges))
  if x.nEdges == 0 { return s }
  i += a
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if x.eAnchor.any == nil { ker.Oops() }
    k := uint32(Codelen (e.any))
    copy (s[i:i+a], Encode (k))
    i += a
    copy (s[i:i+k], Encode (e.any))
    i += k
    s[i] = 0; if e.bool { s[i] = 1 }
    i++
    copy (s[i:i+a], Encode (e.nbPtr0.from.dist))
    i += a
    s[i] = 0; if e.nbPtr0.outgoing { s[i] = 1 }
    i++
    copy (s[i:i+a], Encode (e.nbPtr1.from.dist))
    i += a
    s[i] = 0; if e.nbPtr1.outgoing { s[i] = 1 }
    i++
  }
  return s
}

func (x *graph) check (s string, i, a uint32, bs Stream) {
  n := uint32(len(bs))
  if i >= n {
    errh.Error2(s + ": i =", uint(i), ">= len(bs) =", uint(n))
    as := bs[i:i+a]
    m := uint32(len(as))
    if m != a {
      errh.Error2("a =", uint(a), "!= len =", uint(m))
    }
  }
}

func (x *graph) Decode (s Stream) {
  if len(s) == 0 { panic("gra.Decode: len(s) == 0") }
  x.Clr()
  x.bool = s[0] == 1
  i, a := uint32(1), uint32(4)
  x.nVertices = Decode (uint32(0), s[i:i+a]).(uint32)
  if x.nVertices == 0 {
    return
  }
  i += a
  for n := uint32(0); n < x.nVertices; n++ {
    k := Decode (uint32(0), s[i:i+a]).(uint32)
    i += a
    marked := s[i] == 1
    i++
    vc := Clone (x.vAnchor.any)
    content := Decode (x.vAnchor.any, s[i:i+k])
    x.vAnchor.any = Clone (vc)
    x.insertedVertex (content, marked)
    i += k
  }
  p := Decode (uint32(0), s[i:i+a]).(uint32)
  i += a
  c := Decode (uint32(0), s[i:i+a]).(uint32)
  i += a
  for v, z := x.vAnchor.nextV, uint32(0); v != x.vAnchor; v, z = v.nextV, z + 1 {
    if z == p {
      x.colocal = v
    }
    if z == c {
      x.local = v
    }
  }
  x.nEdges = Decode (uint32(0), s[i:i+a]).(uint32)
  if x.nEdges == 0 {
    return
  }
  i += a
  for z := uint32(0); z < x.nEdges; z++ {
    k := Decode (uint32(0), s[i:i+a]).(uint32)
    i += a
    attrib := Decode (x.eAnchor.any, s[i:i+k])
    e := newEdge (attrib)
    i += k
    e.bool = s[i] == 1
    i++
    fromdist := Decode (uint32(0), s[i:i+a]).(uint32)
    i += a
    v0 := x.vAnchor.nextV
    for fromdist > 0 {
      v0 = v0.nextV
      fromdist--
    }
    e.nbPtr0 = newNeighbour (e, v0, nil, s[i] == 1) // e.nbPtr0.to see below
    i++
    insertNeighbour (e.nbPtr0, v0)
    fromdist = Decode (uint32(0), s[i:i+a]).(uint32)
    i += a
    v0 = x.vAnchor.nextV
    for fromdist > 0 {
      v0 = v0.nextV
      fromdist--
    }
    e.nbPtr0.to = v0
    dir := s[i] == 1

    d := e.nbPtr0.outgoing != dir
    if d != x.bool {
      t := "decoded Graph is "; if x.bool { t += "not " }; ker.Panic (t + "directed")
    }
    e.nbPtr1 = newNeighbour (e, v0, e.nbPtr0.from, dir)
    insertNeighbour (e.nbPtr1, v0)
    e.nextE = x.eAnchor
    e.prevE = x.eAnchor.prevE
    e.prevE.nextE = e
    x.eAnchor.prevE = e
  }
  x.path, x.eulerPath = nil, nil
}
