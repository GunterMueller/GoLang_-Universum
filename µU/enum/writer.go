package enum

// (c) Christian Maurer   v. 170r19 - license see µU.go

import
  "µU/str"
var (
  lWrit, sWrit[]string =
  []string { "",
    "Gellert, Christian Fürchtegott (1715-1769)",
    "Lessing, Gotthold Ephraim (1729-1781)",
    "Claudius, Matthias (1740-1815)",
    "Goethe, Johann Wolfgang (1749-1832)",
    "Lenz, Jakob Michael Reinhold (1751-1792)",
    "Schiller, Friedrich (1759-1805)",
    "Paul, Jean (1763-1825)",
    "Hölderlin, Friedrich (1770-1843)",
    "Hoffmann, Ernst Theodor Amadeus (1776-1822)",
    "Kleist, Heinrich von (1777-1811)",
    "Chamisso, Adelbert von (1781-1838)",
    "Eichendorff, Joseph von (1788-1857)",
    "Droste-Hülshoff, Annette von (1797-1848)",
    "Grabbe, Christian Dietrich (1801-1836)",
    "Mörike, Eduard (1804-1875)",
    "Stifter, Adalbert (1805-1868)",
    "Büchner, Georg (1813-1837)",
    "Storm, Theodor (1817-1888)",
    "Keller, Gottfried (1819-1890)",
    "Fontane, Theodor (1819-1898)",
    "Wedekind, Frank (1864-1918)",
    "Hauptmann, Gerhart (1864-1946)",
    "Lasker-Schüler, Else (1869-1945)",
    "Morgenstern, Christian (1871-1914)",
    "Mann, Heinrich (1871-1950)",
    "Rilke, Rainer Maria (1875-1926)",
    "Mann, Thomas (1875-1955)",
    "Zweig, Stefan (1881-1942)",
    "Kafka, Franz (1883-1924)",
    "Ringelnatz, Joachim  (1883-1934)",
    "Feuchtwanger, Lion (1884-1958)",
    "Tucholsky, Kurt (1890-1935)",
    "Werfel, Franz (1890-1945)",
    "Brecht, Berthold (1898-1956)",
    "Hausmann, Manfred (1898-1986)",
    "Kästner, Erich (1904-1974)",
    "Frisch, Max (1911-1991)",
    "Andersch, Alfred (1914-1980)",
    "Weiss, Peter (1916-1982)",
    "Böll, Heinrich (1917-1985)",
    "Dürrenmatt, Friedrich (1921-1990)",
    "Bachmann, Ingeborg (1926-1973)",
    "Lenz, Siegfried (1926-)",
    "Grass, Günter (1927-)",
    "Walser, Martin (1927-)",
/*
    ",  (18-19)",
*/
  },
  lWrit
  NWriters = uint8(len(lWrit))
)

func init() {
  for i:= 1; i < len(lWrit); i++ {
    p, _:= str.Pos (lWrit[i], ',')
    sWrit[i] = str.Part (lWrit[i], 0, p)
  }
  sWrit[21] = "Bach, Ph.E."
  l[Writer], s[Writer] = lWrit, sWrit
}
