package schol

// (c) Christian Maurer   v. 210308 - license see µU.go

// >>> TODO: simplify it into a Molecule

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/box"
//  "µU/errh"
  "µU/font"
  "µU/pbox"
  "µU/text"
  "µU/pers"
  "µU/cntry"
  "µU/addr"
  "µU/langs"
  "µU/enum"
//  "µU/atom"
//  "µU/mol"
//  "µU/masks"
)
const ( // Order
  nameOrder = iota
  ageOrder
  nOrders
)
type
  scholar struct {
                 pers.Person
                 text.Text "place of birth"
                 cntry.Country "nationality"
                 addr.Address
        guardian pers.Person "legal guardian"
        addressG addr.Address
                 langs.LanguageSequence
                 enum.Enumerator "religion"
                 byte "not used"
                 Format
           field []Any
              cl []uint
// ^ superfluous, instead:
//               mol.Molecule
             }
const
  lenPlace = uint(22)
var (
  bx = box.New()
  cF, cB col.Colour = col.White(), col.Black()
  pbx = pbox.New()
  temp, temp1 = new_().(*scholar), new_().(*scholar)
)

func new_() Scholar {
  x := new (scholar)
  x.Person = pers.New()
  x.Text = text.New (lenPlace)
  x.Country = cntry.New()
  x.Address = addr.New()
  x.guardian = pers.New()
  x.guardian.SetFormat (pers.LongT)
  x.addressG = addr.New()
  x.LanguageSequence = langs.New()
  x.Enumerator = enum.New (enum.Religion)
  x.field = []Any { x.Person, x.Text, x.Country, x.Address, x.guardian, x.addressG, x.LanguageSequence, x.Enumerator, x.byte }
  x.cl = []uint { x.Person.Codelen(), lenPlace, x.Country.Codelen(), x.Address.Codelen(), x.guardian.Codelen(), x.addressG.Codelen(), x.LanguageSequence.Codelen(), x.Enumerator.Codelen(), 1 }
/*
  x.Molecule = mol.New()
  a := atom.New (pers.New())
  a.SetFormat (pers.LongTB)
  x.Ins (a, 0, 0)

  a = atom.New (text.New (lenPlace)) // place of birth
  x.Ins (a, _, _)

  a = atom.New (cntry.New()) // nationality
  x.Ins (a, _, _)

  a = atom.New (addr.New()) // address
  x.Ins (a, _, _)

  a = atom.New (pers.New()) // legal guardian
  a.SetFormat (pers.LongT)
  x.Ins (a, _, _)

  a = atom.New (addr.New()) // guardian's address
  x.Ins (a, _, _)

  a = atom.New (langs.New())
  x.Ins (a, _, _)

  a = atom.New (enumerator.New (enum.Religion))
  x.Ins (a, _, _)

  a = atom.New (char.New()) // not used
  x.Ins (a, _, _)
*/
  x.Format = Short
  x.Colours (col.LightCyan(), col.Black())
// TODO masks
  return x
}

func (x *scholar) imp (Y Any) *scholar {
  y, ok := Y.(*scholar)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

/*
func (x *scholar) imp (Y Any) mol.Molecule {
  y, ok := Y.(*scholar)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y.Molecule
}
*/

func (x *scholar) Empty() bool { // superfluous
  return x.Person.Empty()
}

func (x *scholar) Clr() { // superfluous
  x.Person.Clr()
  x.Text.Clr()
  x.Country.Clr()
  x.Address.Clr()
  x.guardian.Clr()
  x.addressG.Clr()
  x.LanguageSequence.Clr()
  x.Enumerator.Clr()
  x.byte = 0
}

func (x *scholar) Copy (Y Any) {
  y := x.imp (Y)
  y.Person.Copy (x.Person)
  y.Text.Copy (x.Text)
  y.Country.Copy (x.Country)
  y.Address.Copy (x.Address)
  y.guardian.Copy (x.guardian)
  y.addressG.Copy (x.addressG)
  y.LanguageSequence.Copy (x.LanguageSequence)
  y.Enumerator.Copy (x.Enumerator)
  y.byte = x.byte
// ^ superfluous, instead:
//  x.Molecule.Copy (x.imp (Y))
}

func (x *scholar) Clone() Any {
  y := new_()
  y.Copy (x)
  return x
}

func (x *scholar) Eq (Y Any) bool {
  y := x.imp (Y)
  switch x.Format {
  case Minimal, VeryShort:
    return x.Person.Eq (y.Person)
  case Short:
    return x.Person.Eq (y.Person) &&
           x.LanguageSequence.Eq (y.LanguageSequence)
  } // Long:
  return x.Person.Eq (y.Person) &&
         x.LanguageSequence.Eq (y.LanguageSequence) &&
         x.Text.Eq (y.Text) &&
         x.Country.Eq (y.Country) &&
         x.Address.Eq (y.Address) &&
         x.guardian.Eq (y.guardian) &&
         x.addressG.Eq (y.addressG) &&
         x.Enumerator.Eq (y.Enumerator)
// ^ superfluous, instead:
//  return x.Molecule.Eq (x.imp (Y))
}

func (x *scholar) Equiv (Y Any) bool {
  return x.Person.Equiv (x.imp (Y).Person)
}

func (x *scholar) Less (Y Any) bool {
  return x.Person.Less (x.imp (Y).Person)
//  return x.Molecule.Less (x.imp (Y))
}

func (x *scholar) String() string {
  return x.Person.String()
}

/*
func (x *scholar) Num (l []*enum.scholar, v, b []uint) uint {
  return x.LanguageSequence.Num (l, v, b)
}
*/

func (x *scholar) FullAged() bool {
  return x.Person.FullAged()
}

func (x *scholar) SetFormat (f Format) {
  if f < NFormats {
    x.Format = f
    switch f {
    case Minimal:
      x.Person.SetFormat (pers.ShortB)
    case Short:
      x.LanguageSequence.SetFormat (langs.Short)
    case Long:
      x.LanguageSequence.SetFormat (langs.Long)
    }
  }
}

func (x *scholar) GetFormat() Format {
  return x.Format
}

func (x *scholar) Colours (f, b col.Colour) { // superfluous
  x.Person.Colours (f, b)
  x.Text.Colours (f, b)
  x.Country.Colours (f, b)
  x.Address.Colours (f, b)
  x.guardian.Colours (cF, cB)
  x.addressG.Colours (cF, cB)
  x.LanguageSequence.Colours (f, b)
  x.Enumerator.Colours (f, b)
}

var
  lLs, lPb, lNa, lAd, lLg, lAg, lRe,
  cLs, cPb, cNa, cAd, cLg, cAg, cRe uint

func (x *scholar) writeMask (l, c uint) { // complicated, see TODO top
  switch x.Format {
  case Minimal:
    lLs = 0; cLs = 0
  case VeryShort:
    lLs = 0; cLs = 0
  case Short:
    lLs = 1; cLs = 16
  case Long:
/*        1         2         3         4         5         6         7
0123456789012345678901234567890123456789012345678901234567890123456789012345
Geburtsort: ______________________ Staatsangehörigkeit: ____________________

gesetzl. Vertreter(in):
Person, Anschrift
Sprachenfolge: Sprachenfolge
 
Religionszugehörigkeit: ______________________
*/
    lPb = 1; cPb = 12; lNa = 1; cNa = 56
    lAd = 3; cAd = 0; lLg = 7; cLg = 0; lAg = 9; cAg = 0
    lLs = 12; cLs = 15; lRe = 17; cRe = 25
  }
  bx.ScrColours()
  switch x.Format {
    case Minimal, VeryShort:
  default:
    bx.Wd (14)
    bx.Write ("Sprachenfolge:", l + lLs, c + cLs - 15)
  }
  if x.Format == Long {
    bx.Wd (11)
    bx.Write ("Geburtsort:", l + lPb, c + cPb - 12)
    bx.Wd (20)
    bx.Write ("Staatsangehörigkeit:", l + lNa, c + cNa - 21)
    bx.Wd (24)
    bx.Write ("gesetzl. Vertreter(in)", l + lLg - 1, c + 1)
    bx.Wd (23)
    bx.Write ("Religionszugehörigkeit:", l + lRe, c + cRe - 24)
  }
}

func (x *scholar) Write (l, c uint) {
  x.writeMask (l, c)
  x.Person.Write (l, c)
  switch x.Format {
  case Minimal, VeryShort:
  default:
    x.LanguageSequence.Write (l + lLs, c + cLs)
  }
  if x.Format == Long {
    x.Text.Write (l + lPb, c + cPb)
    x.Country.Write (l + lNa, c + cNa)
    x.Address.Write (l + lAd, c + cAd)
    x.guardian.Write (l + lLg, c + cLg)
    x.addressG.Write (l + lAg, c + cAg)
    x.Enumerator.Write (l + lRe, c + cRe)
  }
}

func (x *scholar) Edit0 (l, c uint) {
  x.Person.Edit (l, c)
}

func (x *scholar) Edit (l, c uint) {
  const nKomponenten = 8
  x.Write (l, c)
  i := 1
  loop:
  for {
    switch i {
    case 1:
    x.Person.Edit (l, c)
    case 2:
      if x.Format == Long {
        x.Text.Edit (l + lPb, c + cPb)
      }
    case 3:
      if x.Format == Long {
        x.Country.Edit (l + lNa, c + cNa)
      }
    case 4:
      if x.Format == Long {
        x.Address.Edit (l + lAd, c + cAd)
//        if ! x.Person.FullAged() {
//          x.Address.Copy (x.addressG)
//        }
      }
    case 5:
      if x.Format == Long {
        if x.Person.FullAged() {
          // x.guardian.Clr()
        } else {
          x.guardian.Edit (l + lLg, c + cLg)
        }
      }
    case 6:
      if x.Format == Long {
        if x.Person.FullAged() {
          // x.legalGuardian.Clr()
        } else {
          x.addressG.Edit (l + lAg, c + cAg)
        }
      }
    case 7:
      if x.Format != Minimal && x.Format != VeryShort {
        x.LanguageSequence.Edit (l + lLs, c + cLs)
      }
    case nKomponenten:
      if x.Format == Long {
        x.Enumerator.Edit (l + lRe, c + cRe)
      }
    }
    switch C, d := kbd.LastCommand(); C {
    case kbd.Enter:
      if d == 0 /* aufpassen bei i == 0 ! */ {
        if i < nKomponenten { i++ } else { break loop }
      } else {
        break loop
      }
    case kbd.Esc:
      break loop
    case kbd.Down:
      if i < nKomponenten { i++ } else { break loop }
    case kbd.Up:
      if i > 1 { i -- }
    }
//    if ! x.Person.Identifiable() {
//      errh.Error0("Name, Vorname, Geb.-Datum ?")
//    }
  }
}

func (x *scholar) SetFont (f font.Font) {
  pbx.SetFont (f)
}

func (x *scholar) printMask (l, c uint) {
  switch x.Format {
  case Minimal:
    lLs = 0; cLs = 0
  case VeryShort:
    lLs = 0; cLs = 0
  case Short:
    lLs = 1; cLs = 16
  case Long:
/*        1         2         3         4         5         6         7
0123456789012345678901234567890123456789012345678901234567890123456789012345
Person
Geburtsort: ______________________ Staatsangehörigkeit: ____________________

Anschrift
gesetzl. Vertreter(in):
Person, Anschrift

Sprachenfolge: ___________ von Klasse __ bis Klasse __
               ___________ von Klasse __ bis Klasse __
               ___________ von Klasse __ bis Klasse __
               ___________ von Klasse __ bis Klasse __

Religionszugehörigkeit: ______________________
*/
    lPb = 1; cPb = 12; lNa = 1; cNa = 56
    lAd = 3; cAd = 0; lLg = 7; cLg = 0; lAg = 9; cAg = 0
    lLs = 12; cLs = 15; lRe = 17; cRe = 25
  }
  switch x.Format {
  case Minimal, VeryShort:
  default:
    pbx.Print ("langSeq:", l + lLs, c + cLs - 15)
  }
  if x.Format == Long {
    pbx.Print ("Geburtsort:", l + lPb, c + cPb - 12)
    pbx.Print ("Staatsangehörigkeit:", l + lNa, c + cNa - 21)
    pbx.Print ("gesetzl. Vertreter(in):", l + lLg - 1, c + 1)
    pbx.Print ("Religionszugehörigkeit:", l + lRe, c + cRe - 24)
  }
}

func (x *scholar) Print (l, c uint) {
  x.printMask (l, c)
  x.Person.Print (l, c)
  switch x.Format {
  case Minimal, VeryShort:
  default:
    x.LanguageSequence.Print (l + lLs, c + cLs)
  }
  if x.Format == Long {
    x.Text.Print (l + lPb, c + cPb)
    x.Country.Print (l + lNa, c + cNa)
    x.Address.Print (l + lAd, c + cAd)
    x.guardian.Print (l + lLg, c + cLg)
    x.addressG.Print (l + lAg, c + cAg)
    x.Enumerator.Print (l + lRe, c + cRe)
  }
}

func (x *scholar) Codelen() uint {
  var c uint; for _, n := range x.cl { c += n }
  return c
}

func (x *scholar) Encode() Stream {
  return Encodes (x.field, x.cl)
}

func (x *scholar) Decode (bs Stream) {
  Decodes (bs, x.field, x.cl)
  x.Person = x.field[0].(pers.Person)
  x.Text = x.field[1].(text.Text)
  x.Country = x.field[2].(cntry.Country)
  x.Address = x.field[3].(addr.Address)
  x.guardian = x.field[4].(pers.Person)
  x.addressG = x.field[5].(addr.Address)
  x.LanguageSequence = x.field[6].(langs.LanguageSequence)
  x.Enumerator = x.field[7].(enum.Enumerator)
}

func (x *scholar) Index() Func {
  return func (a Any) Any {
    x, ok := a.(*scholar)
    if ! ok { TypeNotEqPanic (x, a) }
    return x.Person
  }
}

func (x *scholar) rotOrder() {
  x.Person.Rotate()
}
