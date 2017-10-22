package chanm

// (c) Christian Maurer   v. 161223 - license see µU.go

import (
  "sync"
  . "µU/obj"
)
type
  channelModel struct {
                      Any "content in channel"
                      bool "sender first at rendezvous"
               mutex,
    sender, receiver,
           rendezvous sync.Mutex
                      }

func new_() ChannelModel {
  x := new(channelModel)
  x.Any = nil
  x.bool = true
  x.rendezvous.Lock()
  return x
}

func (x *channelModel) Send (a Any) {
  x.sender.Lock()
  x.mutex.Lock()
  x.Any = Clone(a)
  if x.bool {
    x.bool = false
    x.mutex.Unlock()
    x.rendezvous.Lock()
    x.mutex.Unlock()
  } else { // receiver first at rendezvous
    x.bool = true
    x.rendezvous.Unlock()
  }
  x.sender.Unlock()
}

func (x *channelModel) Empty() bool {
  return x.Any == nil
}

func (x *channelModel) Recv() Any {
  var b Any
  x.receiver.Lock()
  x.mutex.Lock()
  if ! x.bool { // sender first at rendezvous
    x.bool = true
    b = Clone (x.Any)
    x.Any = nil
    x.rendezvous.Unlock()
  } else { // receiver first at rendezvous
    x.bool = false
    x.mutex.Unlock()
    x.rendezvous.Lock()
    b = Clone(x.Any)
    x.Any = nil
    x.mutex.Unlock()
  }
  x.receiver.Unlock()
  return b
}
