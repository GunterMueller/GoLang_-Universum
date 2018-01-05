package gra

// (c) Christian Maurer   v. 171122 - license see nU.go

import . "nU/obj"

func (x *graph) numEdges (n *vertex) uint {
  c := uint(0)
  for nb := n.nbPtr; nb.nextNb != n.nbPtr; nb = nb.nextNb {
    c++
  }
  return c
}

func (x *graph) Eq (Y Any) bool { // disgusting complexity
  y := x.imp (Y)
  if x.nVertices != y.nVertices || x.nEdges != y.nEdges {
    return false
  }
  if ! TypeEq (x.vAnchor.Any, y.vAnchor.Any) || ! TypeEq (x.eAnchor.Any, y.eAnchor.Any) {
    panic ("oops")
  }
  ya := y.local // save
  eq := true
  loop:
  for xv := x.vAnchor.nextV; xv != x.vAnchor; xv = xv.nextV {
    if ! y.Ex (xv.Any) {
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
          if x.eAnchor.Any != nil {
            if xn.edgePtr == nil { break }
            if yn.edgePtr == nil { break }
            aa = Eq (xn.edgePtr.Any, yn.edgePtr.Any)
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

func (x *graph) Copy (Y Any) {
  y := x.imp(Y)
  x.Decode (y.Encode())
  x.write, x.write2 = y.write, y.write2
}

func (x *graph) Clone() Any {
  y := new_(x.bool, x.vAnchor.Any, x.eAnchor.Any)
  y.Copy (x)
  return y
}

func (x *graph) Less (Y Any) bool {
  return false
}

func (x *graph) Empty() bool {
  return x.vAnchor.nextV == x.vAnchor
}

func delEdge (e *edge) {
  e.prevE.nextE, e.nextE.prevE = e.nextE, e.prevE
  e.nbPtr0.prevNb.nextNb, e.nbPtr0.nextNb.prevNb = e.nbPtr0.nextNb, e.nbPtr0.prevNb // bug
  e.nbPtr1.prevNb.nextNb, e.nbPtr1.nextNb.prevNb = e.nbPtr1.nextNb, e.nbPtr1.prevNb
}

func delVertex (v *vertex) {
  n := v.nbPtr.nextNb
  for n != v.nbPtr {
    n = v.nbPtr
//    n.to.predecessor = nil
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
}

func (x *graph) Codelen() uint {
  c := uint(1) + 4
  if x.nVertices > 0 {
    for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
      c += 4 + Codelen(v.Any) + 1
    }
    c += 3 * 4
    if x.nEdges > 0 {
      for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
        c += 4 + Codelen(e.Any) + 1 + 2 * (4 + 1)
      }
    }
  }
  return c
}

func (x *graph) Encode() Stream {
  bs := make (Stream, x.Codelen())
  bs[0] = 0; if x.bool { bs[0] = 1 }
  i, a := uint32(1), uint32(4)
  copy (bs[i:i+a], Encode (x.nVertices))
  if x.nVertices == 0 { return bs }
  i += a
  z := uint32(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    k := uint32(Codelen (v.Any))
    copy (bs[i:i+a], Encode (k))
    i += a
    bs[i] = 0; if v.bool { bs[i] = 1 }
    i++
    copy (bs[i:i+k], Encode (v.Any))
    i += k
    v.dist = z
    z++
  }
  copy (bs[i:i+a], Encode (x.colocal.dist))
  i += a
  copy (bs[i:i+a], Encode (x.local.dist))
  i += a
  copy (bs[i:i+a], Encode (x.nEdges))
  if x.nEdges == 0 { return bs }
  i += a
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    k := uint32(Codelen (e.Any))
    copy (bs[i:i+a], Encode (k))
    i += a
    copy (bs[i:i+k], Encode (e.Any))
    i += k
    bs[i] = 0; if e.bool { bs[i] = 1 }
    i++
    copy (bs[i:i+a], Encode (e.nbPtr0.from.dist))
    i += a
    bs[i] = 0; if e.nbPtr0.outgoing { bs[i] = 1 }
    i++
    copy (bs[i:i+a], Encode (e.nbPtr1.from.dist))
    i += a
    bs[i] = 0; if e.nbPtr1.outgoing { bs[i] = 1 }
    i++
  }
  return bs
}

func (x *graph) Decode (bs Stream) {
  if len(bs) == 0 { panic("gra.Decode: len(bs) == 0") }
  x.Clr()
  x.bool = bs[0] == 1
  i, a := uint32(1), uint32(4)
  x.nVertices = Decode (uint32(0), bs[i:i+a]).(uint32)
  if x.nVertices == 0 {
    return
  }
  i += a
  for n := uint32(0); n < x.nVertices; n++ {
    k := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    marked := bs[i] == 1
    i++
    vc := Clone (x.vAnchor.Any)
    content := Decode (x.vAnchor.Any, bs[i:i+k])
    x.vAnchor.Any = Clone (vc)
    x.insertedVertex (content, marked)
    i += k
  }
  p := Decode (uint32(0), bs[i:i+a]).(uint32)
  i += a
  c := Decode (uint32(0), bs[i:i+a]).(uint32)
  i += a
  for v, z := x.vAnchor.nextV, uint32(0); v != x.vAnchor; v, z = v.nextV, z + 1 {
    if z == p {
      x.colocal = v
    }
    if z == c {
      x.local = v
    }
  }
  x.nEdges = Decode (uint32(0), bs[i:i+a]).(uint32)
  if x.nEdges == 0 {
    return
  }
  i += a
  for z := uint32(0); z < x.nEdges; z++ {
    k := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    attrib := Decode (x.eAnchor.Any, bs[i:i+k])
    e := newEdge (attrib)
    i += k
    e.bool = bs[i] == 1
    i++
    fromdist := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    v0 := x.vAnchor.nextV
    for fromdist > 0 {
      v0 = v0.nextV
      fromdist--
    }
    e.nbPtr0 = newNeighbour (e, v0, nil, bs[i] == 1) // e.nbPtr0.to see below
    i++
    insertNeighbour (e.nbPtr0, v0)
    fromdist = Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    v0 = x.vAnchor.nextV
    for fromdist > 0 {
      v0 = v0.nextV
      fromdist--
    }
    e.nbPtr0.to = v0
    dir := bs[i] == 1
    i++
    e.nbPtr1 = newNeighbour (e, v0, e.nbPtr0.from, dir)
    insertNeighbour (e.nbPtr1, v0)
    e.nextE = x.eAnchor
    e.prevE = x.eAnchor.prevE
    e.prevE.nextE = e
    x.eAnchor.prevE = e
  }
}
