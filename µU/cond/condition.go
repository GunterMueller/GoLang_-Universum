package cond

// (c) Christian Maurer    v. 241001 - license see  mU.go

type
  condition struct {
                   }

func new_() Condition {
  x := new(condition)

  return x
}

func (x *condition) Wait() {

}

func (x *condition) Signal() {

}

func (x *condition) Awaited() bool {
  return false
}
