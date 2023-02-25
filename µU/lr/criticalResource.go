package lr

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> bounded left/right problem

import
  "µU/cr"
type
  criticalResource struct {
                          cr.CriticalResource
                          }

func newCR (mL, mR uint) LeftRight {
  const nc = 2
  x := &criticalResource { cr.New (nc, 1) }
  m := make([][]uint, nc)
  for i := uint(0); i < nc; i++ { m[i] = make([]uint, 1) }
  m[0][0], m[1][0] = mL, mR
  x.Limit (m)
  return x
}

func (x *criticalResource) LeftIn() {
  x.Enter (left)
}

func (x *criticalResource) LeftOut() {
  x.Leave (left)
}

func (x *criticalResource) RightIn() {
  x.Enter (right)
}

func (x *criticalResource) RightOut() {
  x.Leave (right)
}

func (x *criticalResource) Fin() {
}
