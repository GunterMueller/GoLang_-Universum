package prt

// (c) murus.org  v. 151116 - license see murus.go

import (
  "os/exec"
  "murus/ker"
  "murus/str"
  . "murus/font"
  "murus/nat"
  "murus/pseq"
  "murus/files"
)
const (
  maxC = 160
  maxL = 128
)
var (
  size Size = Big // Normal ??
  tex, dvi, log, ps pseq.PersistentSequence
  texname, dviname, logname, psname, patternname string
  page [][]string // only 1 page TODO more pages
  nC = [NSizes]uint { 136, 112, 92, 80, 64 }
  nL = [NSizes]uint { 108,  80, 60, 48, 40 }
  dH, dW [NSizes]string
  code, cm [NFonts][NSizes]string
  initialized bool
)

func init() {
  cm[Roman]        = [NSizes]string { "cmtt8 scaled 725",    "cmtt8",               "cmtt10",   "cmtt12",               "cmtt12 scaled 1200" }
  cm[Bold]         = [NSizes]string { "cmbtt10 scaled 600",  "cmbtt8",              "cmbtt10",  "cmbtt10 scaled 1200",  "cmbtt10 scaled 1440" }
  cm[Slanted]      = [NSizes]string { "cmsltt10 scaled 600", "cmsltt10 scaled 800", "cmsltt10", "cmsltt10 scaled 1200", "cmsltt10 scaled 1440" }
  cm[Italic]       = [NSizes]string { "cmitt10 scaled 600",  "cmitt10 scaled 800",  "cmitt10",  "cmitt10 scaled 1200",  "cmitt10 scaled 1440" }
  dH               = [NSizes]string { "6.75",                "9.2",                 "12",       "14.4",                 "17.28" }
  dW               = [NSizes]string { "3.661",               "4.446",               "5.412",    "6.224",                "7.78" } // 175 / nC * 72.27 / 25.4
}

func _switch (s Size) {
  size = s
  if ! initialized {
    _init()
    initialized = true
  }
  startPage()
}

func fontsize() Size {
  return size
}

func print1 (b byte, l, c uint, f Font) {
  if l >= nL[size] || c >= nC[size] { return }
  if ! initialized {
    _init()
    initialized = true
    startPage()
  }
  page[l][c] = Code (f, size) + "  "
  str.Replace (&page[l][c], 3, b)
}

func print (s string, l, c uint, f Font) {
  if l >= nL[size] {
    return // TODO more than one page
  }
  t:= str.Lat1 (s)
  n:= uint(len (t))
  for i:= uint(0); i < n; i++ {
    print1 (t[i], l, c + i, f)
  }
}

func ins (s string) {
  for i:= 0; i < len(s); i++ {
    tex.Ins (byte(s[i]))
  }
}

func _init() {
  N:= files.Tmp()
  null:= byte(0)
  patternname = N + "*"
  tex = pseq.New (null)
  texname = N + "tex"
  tex.Name (texname)
  dvi = pseq.New (null)
  dviname = N + "dvi"
  dvi.Name (dviname)
  log = pseq.New (null)
  logname = N + "log"
  log.Name (logname)
  ps = pseq.New (null)
  psname = N + "ps"
  ps.Name (psname)
}

func voffset (mm uint) {
  ins ("\\voffset " + nat.String (mm) + "mm\n")
}

func footline (s string) {
  ins ("\\footline {\\rmf " + s + "\\hfil}\n")
}

func startPage() {
  page = make ([][]string, nL[size])
  for l:= uint(0); l < nL[size]; l++ {
    page[l] = make ([]string, nC[size])
    for c:= uint(0); c < nC[size]; c++ {
      page[l][c] = "    " // str.Clr (4)
    }
  }
  tex.Clr(); dvi.Clr()
  ins ("\\newcount\\nL \\newcount\\nC \\newdimen\\dH \\newdimen\\dW\n")
  ins ("\\nopagenumbers\n")
  ins ("\\catcode`\\^^c4=13 \\def ^^c4{\\\"A} \\catcode`\\^^e4=13 \\def ^^e4{\\\"a}\n")
  ins ("\\catcode`\\^^d6=13 \\def ^^d6{\\\"O} \\catcode`\\^^f6=13 \\def ^^f6{\\\"o}\n")
  ins ("\\catcode`\\^^dc=13 \\def ^^dc{\\\"U} \\catcode`\\^^fc=13 \\def ^^fc{\\\"u}\n")
  ins ("\\catcode`\\^^df=13 \\def ^^df{{\\ss}}\n")
  ins ("\\lccode`\\^^c4=`\\^^e4 \\uccode`\\^^e4=`\\^^c4\n")
  ins ("\\lccode`\\^^d6=`\\^^f6 \\uccode`\\^^f6=`\\^^d6\n")
  ins ("\\lccode`\\^^dc=`\\^^fc \\uccode`\\^^fc=`\\^^dc\n")
//  ins ("\\font\\rmf cmr8 \\font\\ttf cmtt8\n") // for footlines
  for f:= Font(0); f < NFonts; f++ {
    for s:= Size(0); s < NSizes; s++ {
      ins ("\\font\\" + Code (f, s) + " " + cm[f][s] + " ")
    }
    ins ("\n")
  }
  ins ("\\" + Code (Roman, size) + "\n")
  ins ("\\nL " + nat.String (nL[size]) + "\n")
  ins ("\\nC " + nat.String (nC[size]) + "\n")
  ins ("\\dH " + dH[size] + "pt\n")
  ins ("\\dW " + dW[size] + "pt\n")
  ins ("\\voffset -5.4mm\n") // top margin: 1in - 5.4mm = 2cm
  ins ("\\vsize\\nL\\dH \\advance\\vsize by 15.6pt\n") // because of \interlineskip
  ins ("\\baselineskip\\dH \n")
  ins ("\\hsize\\nC\\dW\n") // about 175mm
  ins ("\\def\\E{\\hbox to\\dW{\\hfil}}\n")
  ins ("\\def\\U{\\hbox to\\dW{\\hrulefill}}\n")
  ins ("\\newdimen\\vh\\vh\\baselineskip \\advance\\vh by-3pt\n") // 3pt is no good solution, TODO
  ins ("\\def\\V{\\hbox to\\dW{\\hss\\vrule height\\vh depth 3pt\\hss}}\n") // because absolute
  ins ("\\def\\C#1#2{\\hbox to\\dW{#1\\hss #2\\hss}}\n")
  ins ("\\def\\do#1{\\catcode`#1=12 }\\do\\$\\do\\&\\do\\#\\do\\^\\do\\_\\do\\%\\do\\~\\do\\@\\do\\<\\do\\>\n")
}

func goPrint() {
  for l:= uint(0); l < nL[size]; l++ {
    ins ("\\line{") // \\strut"); <-- changes the line height ! causes trouble
// we have to construct our own strut, depending on the Font, with height and depth TODO
    for c:= uint(0); c < nC[size]; c++ {
      switch page[l][c][3] { case ' ':
        ins ("\\E")
      case '_':
        ins ("\\U")
      case '|':
        ins ("\\V")
      default:
        ins ("\\C\\" + page[l][c])
      }
    }
    ins ("}\n")
  }
  ins ("\\bye\n")
  exec.Command ("tex", "-output-directory", files.TmpDir(), texname).Run()
  ker.Msleep (100)
  exec.Command ("dvips", dviname, "-o", psname).Run()
  ker.Msleep (100)
  exec.Command (PrintCommand, psname, "-o", "fit-to-page").Run()
  tex.Clr(); log.Clr(); dvi.Clr(); ps.Clr()
//  pseq.Erase (texname); pseq.Erase (logname); pseq.Erase (dviname); pseq.Erase (psname) // TODO
}
