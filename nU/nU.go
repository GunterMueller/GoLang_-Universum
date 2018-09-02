package main

/* (c) Christian Maurer   v. 180902
       christian.maurer-berlin.eu proprietary - all rights reserved

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

  APART FROM THIS  THE TEXT IN GERMAN ABOVE AND BELOW  IS A MANDATORY PART  OF THE LICENSE.

  Alle Quelltexte aus diesem Pakets nU  sind äußerst sorgfältig entwickelt und ausgetestet.
  ABER:  Es gibt keine fehlerfreie Software - dies gilt natürlich auch für _diese_ Quellen.
  Ihre Verwendung in Programmen könnte zu SCHÄDEN führen, z. B. zum Abfackeln von Rechnern,
  zur Entgleisung von Eisenbahnen, zum GAU in Atomkraftwerken  oder zum Absturz des Mondes.
  Darum wird vor der Verwendung irgendwelcher Quellen von  nU  in Programmen zu ernsthaften
  Zwecken AUSDRÜCKLICH GEWARNT!

  Meldungen entdeckter Fehler und Hinweise auf Unklarheiten werden sehr dankbar angenommen. */

import ("nU/lockn"; "nU/mbuf"; "nU/mbbuf"; "nU/macc"; "nU/rw"; "nU/lr"; "nU/dgra")

func main() {
  lockn.NewDekker()
  mbuf.New(3)
  mbbuf.New(0,3)
  macc.New()
  rw.New1()
  lr.New1()
  dgra.Touch()
}
