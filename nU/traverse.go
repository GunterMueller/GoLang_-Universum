package main

import (
  "nU/ego"
  "nU/scr"
  "nU/dgra"
)

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8(me)
  g.SetRoot(4)
/*/
  g.Bfs(); scr.Write (g.ParentChildren(), 0, 0)
  g.Bfsfm(); scr.Write (g.ParentChildren(), 0, 0)
  g.Awerbuch(); scr.Write (g.ParentChildren(), 0, 0)
  g.HelaryRaynal(); scr.Write (g.ParentChildren(), 0, 0)

  g.Dfs(); scr.Write (g.ParentChildren(), 0, 0)
           scr.Write ("arrival    / departure", 1, 0)
      scr.WriteNat (g.Time(), 1, 8); scr.WriteNat (g.Time1(), 1, 23)
  g.Dfs1(); scr.Write ("spanning DFS-tree", 7, 0)
  g.Dfsfm(); scr.Write ("spanning DFS-tree", 7, 0)
  g.Bfsfm1(); scr.Write ("spanning BFS-tree", 7, 0)
  g.Awerbuch1(); scr.Write ("spanning DFS-tree", 7, 0)

  g.Ring(); scr.Write ("   is number    in the ring.", 0, 0)
         scr.WriteNat (g.Me(), 0, 0); scr.WriteNat (g.Time(), 0, 13)
  g.Ring1(); scr.Write ("ring made of DFS-tree", 7, 0)
/*/
  g.Ring(); scr.Write ("   is number    in the ring.", 0, 0)
         scr.WriteNat (g.Me(), 0, 0); scr.WriteNat (g.Time(), 0, 13)
}
