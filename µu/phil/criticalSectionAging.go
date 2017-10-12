package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> Speisende Philosophen mit universellen kritischen Abschnitten
//     Beseitigung der Unfairness durch "aging"

import (
  . "µu/obj"
  . "µu/lockp"
  . "µu/cs"
)
type
  criticalSectionAging struct {
                              CriticalSection
                              }

func newCSA() LockerP {
  var (
    eating = make ([]bool, NPhilos)
    hungry = make ([]uint, NPhilos)
  )
  var s CriticalSection
  c := func (i uint) bool {
         starving := true
         for j := uint(0); j < NPhilos; j++ {
           starving = starving && hungry[i] >= hungry[j]
         }
         return starving &&
                ! eating[left(i)] &&
                ! eating[right(i)]
      }
  e := func (a Any, i uint) {
         eating[i] = true
       }
  a := func (a Any, i uint) {
         eating[i], hungry[i] = false, 0
         for j := uint(0); j < NPhilos; j++ {
           if s.Blocked(j) {
             hungry[j]++
           }
         }
       }
  s = New (NPhilos, c, e, a)
  return &criticalSectionAging { s }
}

func (x *criticalSectionAging) Lock (i uint) {
  changeStatus (i, hungry)
  x.Enter (i, nil)
  changeStatus (i, dining)
}

func (x *criticalSectionAging) Unlock (i uint) {
  changeStatus (i, satisfied)
  x.Leave (i, nil)
}
