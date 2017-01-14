package cons

// (c) murus.org  v. 140728 - license see murus.go

import (
  "syscall"; "strconv"
  . "murus/shape"; . "murus/mode"; "murus/ker"; "murus/col"
)
const (
  esc1 = "\x1b["
  ClearScreen = esc1 + "H" + esc1 + "J"
  home = esc1 + "?25h" + esc1 + "?0c"
)
var (
  fbmemsize uint
  fbmem, fbcop,
  emptyBackground []byte
  visible bool // only for console switching
)

func consoleOn() {
  ker.ActivateConsole()
  n:= width * height * uint(colourdepth)
  copy (fbmem[:n], fbcop[:n])
  visible = true
  c:= actual
  c.Warp (c.ht1 * c.blinkY, c.wd1 * c.blinkX, c.consoleShape)
}

func consoleOff() {
  visible = false
  c:= actual
  c.consoleShape = c.blinkShape
  c.Warp (c.ht1 * c.blinkY, c.wd1 * c.blinkX, Off)
  ker.DeactivateConsole()
}

func consoleFin() {
// TODO wait (blink())
// TODO fin (blink())
  c:= actual
  finished = true
  ker.Msleep (250) // provisorial
  c.cursorShape = Off
  print (ClearScreen + home)
}

var
  initialized bool

func resMaxConsole() (uint, uint) {
  consoleInit()
  return width, height
}

func consoleInit() {
  if initialized { return }
  initialized = true
  colbits:= uint(0)
  width, height, colbits, fbmem = ker.Framebuffer()
  if fbmem == nil {
    ker.Panic ("framebuffer was not initialized ! (Is /dev/fb0 crw-rw-rw ?)")
  }
  fullScreen = ModeOf (width, height)
  if Wd(fullScreen) != width || Ht(fullScreen) != height { ker.Panic ("absolute Katastrophe, Katastrophe pur ...") }
  col.SetDepth (colbits)
  colourdepth = col.Depth
  fbmemsize = width * height * colourdepth
  if uint(len (fbmem)) != fbmemsize {
    ker.Panic ("len (fbmem) == " + strconv.Itoa(len(fbmem)) + " != fbmemsize == " + strconv.Itoa(int(fbmemsize)))
  }
  fbcop = make ([]byte, fbmemsize)
  emptyBackground = make ([]byte, fbmemsize) // TODO each X needs it's own emptyBackground ?
  ker.ConsoleInit()
  ker.SetAction (syscall.SIGUSR1, consoleOff)
  ker.SetAction (syscall.SIGUSR2, consoleOn)
  ker.InstallTerm (consoleFin)
  go ker.CatchSignals()
  initConsoleFonts()
  print (esc1 + "2J" + esc1 + "?1c" + esc1 + "?25l")
  visible = true
}
