package rw

// (c) Christian Maurer   v. 210123 - license see nU.go

type ReaderWriter interface { // Protokolle für das Leser-Schreiber-Problem

// Vor.: Der aufrufende Prozess ist weder aktiver Leser
//       noch aktiver Schreiber.
// Er ist aktiver Leser; kein anderer ist aktiver Schreiber.
// Wenn es zum Zeitpunkt des Aufrufs einen aktiven Schreiber
// gab, war er solange verzögert, bis das nicht der Fall war.
  ReaderIn()

// Vor.: Der aufrufende Prozess ist aktiver Leser.
// Er ist es nicht mehr.
  ReaderOut()

// Vor.: Der aufrufende Prozess ist weder aktiver Leser
//       noch aktiver Schreiber.
// Er ist aktiver Schreiber und kein anderer Prozess ist
// aktiver Leser oder Schreiber. Wenn es zum Zeitpunkt des
// Aufrufs aktive Leser oder einen aktiven Schreiber gab,
// war er solange verzögert, bis das nicht der Fall war.
  WriterIn()

// Vor.: Der aufrufende Prozess ist aktiver Schreiber.
// Er ist es nicht mehr.
  WriterOut()
}

// Der Parameter m bedeutet die maximal zulässige Anzahl
// der nebenläufigen Leser.
func New1() ReaderWriter { return new1() }
func New2() ReaderWriter { return new2() }
func NewSemaphore() ReaderWriter { return newS() }
func NewAddSemaphore (m uint) ReaderWriter { return newAS(m) }
func NewGo() ReaderWriter { return newG() }
func NewCriticalSection1() ReaderWriter { return newCS1() }
func NewCriticalSection2() ReaderWriter { return newCS2() }
func NewCriticalSectionBounded (m uint) ReaderWriter { return newCSB(m) }
func NewCriticalSectionFair() ReaderWriter { return newCSF() }
func NewCriticalResource (m uint) ReaderWriter { return newCR(m) }
func NewMonitor1() ReaderWriter { return newM1() }
func NewMonitor2() ReaderWriter { return newM2() }
func NewConditionedMonitor() ReaderWriter { return newCM() }
func NewConditionedMonitorBounded (m uint) ReaderWriter { return newCMB(m) }
func NewChannel() ReaderWriter { return newCh() }
func NewGuardedSelect() ReaderWriter { return newGS() }
