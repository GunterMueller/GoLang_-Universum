package chanm

// (c) Christian Maurer   v. 161223 - license see µu.go

//     Nichtsequentielle Programmierung mit Go 1 kompakt, 7.7, S. 185 ff.

import
  . "µu/obj"
type
  ChannelModel interface { // "models" of channels (i.e. working only within one process)

// a is contained in x.
  Send (a Any)

// Returns true, iff there are no messages in x.
  Empty() bool

// Returns the message, that was sent to x; the message is removed from x.
// The calling process might have been blocked, until x contained a message.
  Recv() Any
}

// Returns an new empty channel model.
func New() ChannelModel { return new_() }
