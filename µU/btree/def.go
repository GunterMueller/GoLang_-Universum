package btree

// (c) Christian Maurer   v. 201001 - license see µU.go
//
// >>> Probably completely useless;
//     apart from that there is some bloody bug in this implementation

import
  . "µU/obj"
type
  BTree interface {

  Root() Any
  Left() BTree
  Right() BTree
  Num() uint
  NumPred (p Pred) uint
  Contained (a Any) (BTree, bool)
  All (p Pred) bool
  ExPred (p Pred) BTree
  Trav (o Op)
  Ins (a Any) (BTree, BTree)
  Del (a Any) (BTree, bool)
  First (a Any) BTree
  Write (x0, x1, y, dy uint, f func (Any) string)
  Write1 (d uint, f func (Any) string)
}

func New (a Any) BTree { return new_(a) }
