package ker

// (c) Christian Maurer   v. 240413 - license see µU.go

import
  "os"

func Binom (n, k uint) uint { return binom(n,k) }
func Bezier (x, y []int, m, n, i uint) (int, int) { return bezier(x,y,m,n,i) }
func ArcLen (xs, ys []int) uint { return arcLen(xs,ys) }

func Fin() { fin() }
func Panic (s string) { panic_(s) }
func Panic1 (s string, n uint) { panic1(s,n) }
func Panic2 (s string, n uint, s1 string, n1 uint) { panic2(s,n,s1,n1) }
func Panic3 (s string, n uint, s1 string, n1 uint, s2 string, n2 uint) { panic3(s,n,s1,n1,s2,n2) }
func PrePanic() { prePanic() }
func Oops() { oops() }
func Halt (s int) { halt(s) }
func InstallTerm (h func()) { installTerm(h) }

func SetAction (s os.Signal, a func()) { setAction(s,a) }
func CatchSignals() { catchSignals() }

func ConsoleInit() { init_() }
func SwitchConsole (forward bool) { switch_(forward) }
func Console (a uint8) { console(a) }
func ActualConsole() uint { return actual() }
func DeactivateConsole() { deactivate() }
func ActivateConsole() { activate() }
