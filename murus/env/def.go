package env

// (c) murus.org  v. 170814 - license see murus.go

// Pre: v does not contain characters, that must not be contained
//      in an environment variable, as e.g. ' ' or '='.
// v is an environment variable with the value c.
func Set (v string, c *string) { set(v, c) }

// Returns the value of the environment variable v,
// if that is defined; otherwise "".
func Val (v string) string { return val(v) }

// Returns the name of the user.
func User() string { return user() }

// Returns the name of the user's home directory.
func Home() string { return home() }

// Returns the first byte of the 1st parameter of the programm call,
// if that was given; otherwise 0.
func Par1() byte { return par1() }

// Returns the first byte of the 2nd parameter of the programm call,
// if that was given; otherwise 0.
func Par2() byte { return par2() }

// Returns the number of parameters of the programm call.
func NPars() uint { return nPars() }

// Returns the i-th parameter of the programm call,
// if that was given; otherwise "".
func Par (i uint) string { return par(i) }

// Returns the name of the call.
func Call() string { return call() }
