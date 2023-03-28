package rob

// (c) Christian Maurer   v. 210325 - license see µU.go

func init() {
  for y := 1; y < z - 1; y++ {
    bild[0][y] = "x                              x"
  }
  bild[0][0] = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  bild[0][z-1] = bild[0][0]
/*/
  bild [1] = [z]string {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "x             ++++             x",
    "x            ++++++            x",
    "x           +oo++oo+           x",
    "x           +oo++oo+           x",
    "x           ++++++++           x",
    "x            +oooo+            x",
    "x             ++++             x",
    "x             +oo+             x",
    "x           +oooooo+           x",
    "x         +oooooooooo+         x",
    "x        +ooo+oooo+ooo+        x",
    "x       +ooo +oooo+ ooo+       x",
    "x      +ooo  +oooo+  ooo+      x",
    "x      +ooo  +oooo+   ooo+     x",
    "x      +oo   +oooo+    oo+     x",
    "x     +oo    +oooo+    oo+     x",
    "x     +oo   ++++++++   oo+     x",
    "x    o+o    +oooooo+  o+o      x",
    "x    o+o    +oooooo+  o+o      x",
    "x          +ooo  ooo+          x",
    "x         +ooo    ooo+         x",
    "x         +ooo    ooo+         x",
    "xkkkkkkkk +ooo    ooo+   mmmm  x",
    "xkkkkkkkk +ooo    ooo+  mmmmmm x",
    "xkkkkkkkk +oo     oo+  mmm  mmmx",
    "xkkkkkkkk +oo     oo+  mm    mmx",
    "xkkkkkkkk +oo     oo+  mm    mmx",
    "xkkkkkkkk +oo     oo+  mmm  mmmx",
    "xkkkkkkkooooo     ooooo mmmmmm x",
    "xkkkkkkkooooo     ooooo  mmmm  x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  }
/*/
  bild[1] = [z]string {
    "xxxxxxxxxxxooxxxxxxxxxxxxxxxxxxx",
    "x          oo                  x",
    "x          oo                  x",
    "x          oo ++++             x",
    "x          oo++++++            x",
    "x          o++++++++           x",
    "x      oooo++oo++oo++oooo      x",
    "x        oo+ooo++ooo+oo        x",
    "x         o++oo++oo++o         x",
    "x         ++++++++++++         x",
    "x        ++++oo++oo++++        x",
    "x        +++++oooo+++++        x",
    "x             ++++             x",
    "x       +oo++++++++++oo+       x",
    "x      +ooo+oooooooo+ooo+      x",
    "x      +ooo+oooooooo+ooo+      x",
    "x     +ooo++oooooooo++ooo+     x",
    "x     +oo+ +oooooooo+ +oo+     x",
    "x    +ooo+ +oooooooo+ +ooo+    x",
    "x   +ooo+  ++++++++++  +ooo+   x",
    "x   oooo   oooooooooo   oooo   x",
    "xkkko+o+kkkoooooooooo   +o+o   x",
    "xkkko+o+kkk+oooooooo+   +o+o   x",
    "xkkko+o+kkk+oo++++oo+   +o+om  x",
    "xkkkkkkkkkk+oo+  +oo+   mmmmmm x",
    "xkkkkkkkkkk+oo+  +oo+  mmm  mmmx",
    "xkkkkkkkkkk+oo+  +oo+  mm    mmx",
    "xkkkkkkkkkk+oo+  +oo+  mm    mmx",
    "xkkkkkkkkoooooo  oooooommm  mmmx",
    "xkkkkkkkooooooo  ooooooommmmmm x",
    "xkkkkkkoooooooo  oooooooommmm  x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  }
  bild[2] = [z]string {
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            ",
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            ",
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            ",
    "oooooooooooooooooooooooooooooooo",
    "           o               o    ",
    "           o               o    ",
    "           o               o    ",
    "oooooooooooooooooooooooooooooooo",
    "   o               o            ",
    "   o               o            ",
    "   o               o            ",
  }
  bild[3] = [z]string {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "x                              x",
    "x           oo                 x",
    "x          o  oooo             x",
    "x         o       oooo         x",
    "x        o   mmmmmm   oooo     x",
    "x       o  mmmmmmmmmm     ooo  x",
    "x      o  mmmmm  mmmmm     oo  x",
    "x     o   mmmm     mmmm   o o  x",
    "x    o     mmmmm  mmmmm  o  o  x",
    "x   oo      mmmmmmmmmm  o   o  x",
    "x   oooooo    mmmmmm   o  o o  x",
    "x   oo o ooooo        o     o  x",
    "x   o o o o o oooo   o  o   o  x",
    "x   oo o o o o o oooo     o o  x",
    "x   o o o o o o o o o o     o  x",
    "x   oo o o o o o o oo   o   o  x",
    "x   o o o o o o o o o     o o  x",
    "x   oo o o o o o o oo o     o  x",
    "x   o o o o o o o o o   o   o  x",
    "x   oo o o o o o o oo     o o  x",
    "x   o o o o o o o o o o     o  x",
    "x   oo o o o o o o oo   o  o   x",
    "x   o o o o o o o o o     o    x",
    "x   oo o o o o o o oo o  o     x",
    "x   ooo o o o o o o o   o      x",
    "x     oooo o o o o oo  o       x",
    "x         ooooo o o o o        x",
    "x             oooo ooo         x",
    "x                 ooo          x",
    "x                              x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  }
  bild[4] = [z]string {
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x            oooooo            x",
    "x          oooooooooo          x",
    "x         ooooo  ooooo         x",
    "x         oooo     oooo        x",
    "x          ooooo  ooooo        x",
    "x           oooooooooo         x",
    "x             oooooo           x",
    "x                              x",
    "x                              x",
    "x                              x",
    "x                              x",
    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
  }
}