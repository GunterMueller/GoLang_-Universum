package main

/* (c) Christian Maurer   v. 200908
       christian.maurer-berlin.eu proprietary - all rights reserved

  This package - the n(ano)Universe - contains the source texts from my textbooks
  "Nichtsequentielle und Verteilte Programmierung mit Go" (Springer Vieweg 2018)
  "Nonsequential and Distributed Programming with Go" (Springer Nature 2028).

  THIS SOFTWARE  IS PROVIDED BY THE AUTHOR  "AS IS"  AND ANY EXPRESS OR IMPLIED WARRANTIES,
  INCLUDING,  BUT NOT LIMITED TO,  THE IMPLIED WARRANTIES  OF MERCHANTABILITY  AND  FITNESS
  FOR A PARTICULAR PURPOSE ARE DISCLAIMED.  IN NO EVENT SHALL the authors BE LIABLE FOR ANY
  DIRECT, INDIRECT,  INCIDENTAL, SPECIAL,  EXEMPLARY, OR CONSEQUENTIAL DAMAGES  (INCLUDING,
  BUT NOT LIMITED TO,  PROCUREMENT OF SUBSTITUTE GOODS  OR SERVICES;  LOSS OF USE, DATA, OR
  PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER
  IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY
  WAY OUT OF THE USE OF THIS SOFTWARE,  EVEN IF ADVISED  OF THE POSSIBILITY OF SUCH DAMAGE.

  APART FROM THIS  THE TEXT IN GERMAN ABOVE AND BELOW  IS A MANDATORY PART  OF THE LICENSE.

  All source texts in this package nU are developed extremely carefully and tested.
  HOWEVER:  There does not exist any software without errors - this is of course also true
  for these sources.  Their use in programs could lead to DAMAGES, i.e. to burn computers,
  derail trains, to worst case scenarios in nuclear power plants or the crash of the moon.
  Therefore, you are warned of using sources of nU in programs for serious uses.

  Hints to discovered discrepancies or errors are of course very gratefully accepted!   */


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
