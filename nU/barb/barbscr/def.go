package barbscr

// (c) Christian Maurer   v. 180101 - license see nU.go

const (
  NumCustomers = 10
  Deltat = 1 * NumCustomers
)

func Init() { go snore() }

func TakeSeatInWaitingRoom () { takeSeatInWaitingRoom() }

func GetNextCustomer () { getNextCustomer() }
