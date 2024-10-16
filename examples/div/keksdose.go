package main // keksdose

// (c) Christian Maurer   v. 150222 - license see murus.go

import (
	"math/rand"
	"sync"
	"time"
)
const
  AnzahlMonster = 6
type (
  schöneKekse bool
  Monster uint // < AnzahlMonster
  Kekse uint
)
var (
  hatGebacken,
  hatHunger [AnzahlMonster]chan Kekse
  darfEssen [AnzahlMonster]chan schöneKekse
  KekseInDose Kekse
  einigeKekse [AnzahlMonster]Kekse
  Name [AnzahlMonster]string = [...]string {"Krümelmonster", "Oscar", "Elmo", "Grobi", "Kermit", "Lulatsch"}
  m sync.Mutex
)

func PfotenWeg() {
  m.Lock()
}

func PfotenDürfenWieder() {
  m.Unlock()
  time.Sleep(time.Duration(1e9))
}

func leckereKekse() Kekse {
  return 1 + Kekse(rand.Uint32()) % 8
}

func nimmtHeraus(alleKekse *Kekse, einPaarKekse Kekse) {
  *alleKekse -= einPaarKekse
}

func legtHinein(alleKekse *Kekse, soVieleKekse Kekse) {
  *alleKekse += soVieleKekse
}

func nachzählen(soviele Kekse) {
  PfotenWeg()
  println("In der Keksdose sind", soviele, "Kekse")
  PfotenDürfenWieder()
}

func (vomMonster Monster) willHaben (etliche Kekse) {
  PfotenWeg()
  println (Name[vomMonster], "will", etliche, "Kekse")
  PfotenDürfenWieder()
}

func (Bäcker Monster) danktFür (neue Kekse) {
  PfotenWeg()
  println(Name[Bäcker], "hat", neue, "Kekse gebacken")
  PfotenDürfenWieder()
}

func (Krümel Monster) essenLassen (seine *Kekse) {
  PfotenWeg()
  println (Name[Krümel], "isst", *seine, "Kekse")
  *seine = 0
  PfotenDürfenWieder()
}

func (dasMonster Monster) darfGreifen(dieseKekse schöneKekse) {
  nimmtHeraus (&KekseInDose, einigeKekse[dasMonster])
  darfEssen[dasMonster] <- dieseKekse
  dasMonster.essenLassen (&einigeKekse[dasMonster])
  nachzählen (KekseInDose)
}

func genugKekseDaFür (dasMonster Monster) bool {
  return 0 < einigeKekse[dasMonster] && einigeKekse[dasMonster] <= KekseInDose
}

const erstesMonster = Monster(0)

func Keksdose() {
  var inDieKeksdose schöneKekse
  for {
    for einMonster := erstesMonster; einMonster < AnzahlMonster; einMonster++ {
      select {
      case einigeKekse[einMonster] = <-hatHunger[einMonster]:
        einMonster.willHaben (einigeKekse[einMonster])
        if genugKekseDaFür (einMonster) {
          einMonster.darfGreifen (inDieKeksdose)
        }
      case neueKekse := <-hatGebacken[einMonster]:
        einMonster.danktFür (neueKekse)
        legtHinein (&KekseInDose, neueKekse)
        nachzählen (KekseInDose)
        for dasMonster := erstesMonster; dasMonster < AnzahlMonster; dasMonster++ {
          if genugKekseDaFür (dasMonster) {
            dasMonster.darfGreifen (inDieKeksdose)
          }
        }
      }
    }
  }
}

func (dasMonster Monster) istFleissig() {
  for {
    hatGebacken[dasMonster] <- leckereKekse()
  }
}

func (dasMonster Monster) istVerfressen() {
  for {
    sovieleKekse := leckereKekse()
    hatHunger[dasMonster] <- sovieleKekse
    einigeKekse := <-darfEssen[dasMonster]
    if einigeKekse {
      dasMonster.essenLassen (&sovieleKekse)
    }
  }
}

func main() {
  for m := erstesMonster; m < AnzahlMonster; m++ {
    hatGebacken[m], hatHunger[m] = make(chan Kekse), make(chan Kekse)
    darfEssen[m] = make(chan schöneKekse)
  }
  go Keksdose()
  for dasMonster := erstesMonster; dasMonster < AnzahlMonster; dasMonster++ {
    go dasMonster.istFleissig()
    go dasMonster.istVerfressen()
  }
  for { } // Abbruch mit Strg C
}
