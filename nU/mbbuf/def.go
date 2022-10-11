package mbbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

type MBoundedBuffer interface {
  Ins (a any)
  Get() any
}

func New (a any, n uint) MBoundedBuffer { return new_(a,n) }
func New1 (a any, n uint) MBoundedBuffer { return new1(a,n) }
func NewM (a any, n uint) MBoundedBuffer { return newM(a,n) }
func NewGo (a any, n uint) MBoundedBuffer { return newGo(a,n) }
func NewCh (a any, n uint) MBoundedBuffer { return newCh(a,n) }
func NewCh1 (a any, n uint) MBoundedBuffer { return newCh1(a,n) }
func NewGS (a any, n uint) MBoundedBuffer { return newgs(a,n) }
func NewFM (a any, n uint, h string, p uint16, s bool) MBoundedBuffer { return newfm(a,n,h,p,s) }
