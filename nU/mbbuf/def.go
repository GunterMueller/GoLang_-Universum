package mbbuf

// (c) Christian Maurer   v. 171127 - license see nU.go

import . "nU/obj"

type MBoundedBuffer interface {
  Ins (a Any)
  Get() Any
}

func New (a Any, n uint) MBoundedBuffer { return new_(a,n) }
func New1 (a Any, n uint) MBoundedBuffer { return new1(a,n) }
func NewM (a Any, n uint) MBoundedBuffer { return newM(a,n) }
func NewGo (a Any, n uint) MBoundedBuffer { return newGo(a,n) }
func NewCh (a Any, n uint) MBoundedBuffer { return newCh(a,n) }
func NewCh1 (a Any, n uint) MBoundedBuffer { return newCh1(a,n) }
func NewGS (a Any, n uint) MBoundedBuffer { return newgs(a,n) }
func NewFM (a Any, n uint, h string, p uint16, s bool) MBoundedBuffer { return newfm(a,n,h,p,s) }
