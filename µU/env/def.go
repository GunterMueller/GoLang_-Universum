package env

// (c) Christian Maurer   v. 171125 - license see µU.go

// Pre: v does not contain characters, that must not be contained
//      in an environment variable, as e.g. ' ' or '='.
// v is an environment variable with the value c.
func Set (v string, c *string) { set(v, c) }

// Returns the value of the environment variable v,
// if that is defined; otherwise "".
func Val (v string) string { return val(v) }

// Returns the name of the local host.
func Localhost() string { return localhost() }

// Returns the name of the user (value of the environment variable USER).
func User() string { return user() }

// Returns the name of the user's home directory (value of the environment variable HOME).
func Home() string { return home() }

// Returns the value of the environment variable $GOSRC.
// If that is not $HOME/go/src (where $HOME denotes your home directory),
// you have to set that e.g. in /etc/profile.d/go.sh !
func Gosrc() string { return gosrc() }

// Returns the first byte of the 1st parameter of the program call,
// if that was given; otherwise 0.
func Par1() byte { return par1() }

// Returns the first byte of the 2nd parameter of the program call,
// if that was given; otherwise 0.
func Par2() byte { return par2() }

// Returns the number of parameters of the program call
// (the program call itself not counting as parameter).
func NPars() uint { return nPars() }

// Returns the i-th parameter of the program call,
// if that was given; otherwise "".
func Par (i uint) string { return par(i) }

// Returns the name of the call.
func Call() string { return call() }
