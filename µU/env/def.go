package env

// (c) Christian Maurer   v. 220203 - license see µU.go

// Returns true, iff the calling process runs under X.
func UnderX() bool { return underX() }

// Returns true, iff the calling process runs on a tty-console.
func UnderC() bool { return underC() }

// Returns true, iff 
// func Far() bool { return far() }

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

// Returns the first byte of the 1st argument of the program call,
// if that was given; otherwise 0.
func Arg1() byte { return arg1() }

// Returns the first byte of the 2nd argument of the program call,
// if that was given; otherwise 0.
func Arg2() byte { return arg2() }

// Returns the number of arguments of the program call
// (the program call itself not counting as argument).
func NArgs() uint { return nArgs() }

// Returns the i-th argument of the program call,
// if that was given; otherwise "".
func Arg (i uint) string { return arg(i) }

// Returns the value of the i-th argument of the program call,
// if that was given as natural number; otherwise 0.
func N (i uint) uint { return n(i) }

// Returns the name of the call.
func Call() string { return call() }
