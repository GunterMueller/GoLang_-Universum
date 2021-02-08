package main

/* (c) Christian Maurer   v. 210123

  Dieses Paket - das n(ano)Universum - enthält die Quelltexte aus meinem Lehrbuch
  "Nichtsequentielle und Verteilte Programmierung mit Go" (Springer Vieweg 2018).

  THIS SOFTWARE  IS PROVIDED BY THE AUTHOR  "AS IS"  AND ANY EXPRESS OR IMPLIED WARRANTIES,
  INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS
  FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL the authors BE LIABLE FOR ANY
  DIRECT, INDIRECT,  INCIDENTAL, SPECIAL,  EXEMPLARY, OR CONSEQUENTIAL DAMAGES  (INCLUDING,
  BUT NOT LIMITED TO,  PROCUREMENT OF SUBSTITUTE GOODS  OR SERVICES;  LOSS OF USE, DATA, OR
  PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER
  IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY
  WAY OUT OF THE USE OF THIS SOFTWARE,  EVEN IF ADVISED  OF THE POSSIBILITY OF SUCH DAMAGE.

  Die Quelltexte von nU sind äußerst sorgfältig entwickelt und werden laufend überarbeitet.
  ABER:  Es gibt keine fehlerfreie Software - dies gilt natürlich auch für _diese_ Quellen.
  Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zum Abfackeln von Rechnern,
  zur Entgleisung von Eisenbahnen, zum GAU in Atomkraftwerken  oder zum Absturz des Mondes.
  Deshalb wird vor der Verwendung irgendwelcher Quellen von µU in Programmen zu ernsthaften
  Zwecken AUSDRÜCKLICH GEWARNT! (Ausgenommen sind Demo-Programme zum Einsatz in der Lehre.)

  Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden sehr dankbar angemommen. */

import ("nU/lock"; "nU/lock2"; "nU/lockn";
        "nU/mbuf"; "nU/mbbuf"; "nU/macc"; "nU/rw"; "nU/lr"; "nU/dgra")

func main() {
  lock.NewTAS()
  lock2.NewDekker()
  lockn.NewDijkstra(3)
  mbuf.New(3)
  mbbuf.New(0,3)
  macc.New()
  rw.New1()
  lr.New1()
  dgra.Touch()
}
