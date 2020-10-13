package enum

// (c) Christian Maurer   v. 201011 - license see µU.go

import
  "µU/str"

var (
  lComposer = []string {"",
                       "Palestrina, Claudio da (1525-1594)",
    "Monteverdi, Claudio (1567-1643)",
    "Frescobaldi, Girolamo (1583-1643)",
    "Schütz, Heinrich (1585-1672)",
    "Lully, Jean Baptiste (1632-1687)",
    "Buxtehude, Dietrich (1637-1707)",
    "Corelli, Arcangelo (1653-1713)",
    "Purcell, Henry (1659-1695)",
    "Scarlatti, Alessandro (1659-1725)",
    "Couperin, Francois (1668-1733)",
    "Caldara, Antonio (1670-1736)",
    "Albinoni, Tomaso (1671-1751)",
    "Vivaldi, Antonio (1680-1743)",
    "Telemann, Georg Philipp (1681-1767)",
    "Rameau, Jean-Philippe (1683-1764)  ",
    "Bach, Johann Sebastian (1685-1750) ",
    "Scarlatti, Domenico (1685-1757)",
    "Händel, Georg Friedrich (1685-1759)",
    "Geminiani, Francesco (1687-1762)",
    "Locatelli, Pietro (1695-1764)",
    "Pergolesi, Giovanni Battista (1710-1736)",
    "Gluck, Christoph Willibald (1714-1787)",
    "Bach, Philip Emanuel (1714-1788)", // 23, see 16
    "Haydn, Joseph (1732-1809)",
    "Boccherini, Luigi (1743-1805)",
    "Mozart, Wolfgang Amadeus (1756-1791)",
    "Cherubini, Luigi (1760-1842)",
    "Beethoven, Ludwig van (1770-1827)",
    "Spohr, Ludwig (1774-1859)",
    "Hummel, Johann Nepomuk (1778-1837)",
    "Paganini, Niccolo (1782-1840)",
    "Weber, Carl Maria von (1786-1826)",
    "Rossini, Gioacchino (1792-1868)",
    "Schubert, Franz (1797-1828)",
    "Donizetti, Gaetano (1797-1848)",
    "Lortzing, Albert (1801-1851)",
    "Berlioz, Hector (1803-1869)",
    "Glinka, Michael (1804-1857)",
    "Mendelssohn-Bartholdy, Felix (1809-1874)",
    "Schumann, Robert (1810-1856)",
    "Chopin, Frederic (1810-1849)",
    "Liszt, Franz (1811-1886)",
    "Wagner, Richard (1813-1883)",
    "Verdi, Giuseppe (1813-1901)",
    "Franck, Cesar (1822-1890)",
    "Lalo, Edouard (1823-1892)",
    "Smetana, Friedrich (1824-1884)",
    "Bruckner, Anton (1824-1896)",
    "Strauss, Johann (1825-1899)",
    "Brahms, Johannes (1833-1897)",
    "Borodin, Alexander (1834-1887)",
    "Saint-Saens, Camille (1835-1921)",
    "Bizet, Georges (1836-1875)",
    "Mussorgski, Modest (1839-1881)",
    "Tschaikowskij, Peter (1840-1893)",
    "Dvorak, Antonin (1841-1904)",
    "Grieg, Edward (1843-1907)",
    "Rimskij-Korssakow, Nikolai (1844-1908)",
    "Janacek, Leos (1854-1928)",
    "Mahler, Gustav (1860-1911)",
    "Debussy, Claude (1862-1918)",
    "Strauß, Richard (1864-1949)",
    "Sibelius, Jean (1865-1957)",
    "Pfitzner, Hans (1869-1949)",
    "Scriabin, Alexander (1872-1915)",
    "Reger, Max (1873-1916)",
    "Rachmaninow, Sergej (1873-1947)",
    "Schönberg, Arnold (1874-1951)",
    "Ravel, Maurice (1875-1937)",
    "Falla, Manuel de (1876-1946)",
    "Bartok, Bela (1881-1945)",
    "Strawinsky, Igor (1882-1971)",
    "Webern, Anton von (1883-1945)",
    "Berg, Alban (1885-1935)",
    "Furtwängler, Wilhelm (1886-1954)",
    "Prokofieff, Serge (1891-1953)",
    "Honegger, Arthur (1892-1955)",
    "Hindemith, Paul (1895-1963)",
    "Orff, Carl (1895-1982)",
    "Blacher, Boris (1903-1975)",
    "Chatschaturian, Aram (1903-1978)",
    "Schostakowitsch, Dimitri (1906-1975)",
    "Fortner, Wolfgang (1907-1987)",
    "Britten, Benjamin (1913-1976)",
    "Boulez, Pierre (1925-2016)",
    "Henze, Hans Werner (1926-2012)",
  }
)

func init() {
  l[Composer] = lComposer
  s[Composer] = lComposer
  for i := 1; i < len(l[Composer]); i++ {
    p, _:= str.Pos (l[Composer][i], ',')
    s[Composer][i] = str.Part (l[Composer][i], 0, p)
  }
  s[Composer][23] = "Bach, Ph.E."
}
