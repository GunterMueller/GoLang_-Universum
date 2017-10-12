package files

// (c) Christian Maurer   v. 170814 - license see µu.go

// >>> alpha-version

import
  . "µu/obj"
const
  maxN = 255
type
  Type byte
const
// Types of entries in the file system tree:
  (Unknown = Type(iota); Fifo; Device; Dir; File; Link; Socket; NTypes)
// Entries of type Dir, File etc. are also simply denoted by
// "directory", "file" etc.

// Paths in the file system tree are sequences of names of directories
// - separated from each other by "/" - in the direction
// from the root of the file system tree towards its leaves, s.t.
// >>>jeweils<<< the next directory is contained in the previous one.
// If the first character of a path is the separator, it starts
// in the root directory, otherwise in the actual directory.
// A separator at the end of a path does not matter,
// if the path is not only the root directory "/".
// (The name of a directory is therefore also a path.)
// The last directory of a path is denoted as its target directory;
// its target directory of the path "/" is the root directory.
// Every calling process befindet sich at any time in exactly one
// directory, which is called its actual directory; the sequence
// of the directories from the root directory to the actual directory
// is called the actual path; i.e. the actual directory is always
// the target directory of the actual path.
// The empty string "" is also a path; its target directory
// is the actual directory. It is equivalent with "./".
// For paths p and q and a directory d, "p/d/../q" is
// equivalent with "p/q": both denote the same path.

// Returns the actual directory.
func ActualDir () string { return actualDir() }

// Returns the actual path.
// The actual directory is not changed.
func ActualPath () string { return actualPath() }

// Pre: p is a path. The calling process has the appropriate rights.
// Returns true, iff p is contained in the file system tree, i.e.
// if all directories in p are contained in the file system tree, s.t.
// the >>jeweils<< next directory of p is contained in the previous one.
func Ex (p string) bool { return ex(p) }

// Returns true, iff n is the name of an entry (i.e. a string
// without leading or trailing spaces and without "/") or
// of an entry with leading path and a separator at its end.
func Defined (n string) bool { return defined(n) }

// Pre: p is a path name.
// Returns true, iff the calling process has access rights to
// all directories in n and there is a non empty file with the name n.
func Contained (p string) bool { return contained(p) }

// Pre: p is a path or the string "..".
//      The calling process has access rights to all directories in p resp.
//      (for p = "..") to that directory, in which the actual directory is contained.
// If p == ".." and the actual directory is the root directory, nothing has happened.
// Otherwise the actual directory has changed as follows:
// For p == ".." the actual directory is now that one,
// in which the former actual directory was contained;
// i.e. the last directory is removed from the actual path.
// If p starts with "/", p is the actual path,
// otherwise p is appended to the former actual path.
func Cd (p string) { cd(p) }

// The package furthermore maintains the program variable.
// Its value at the beginning is the name of the called program
// (without the pathname, if called with such).

// Pre: v does not contain spaces and no "=".
// v is the program variable.
//  func set (v string) // TODO name

// The effect is secret (the source is the doc).
func Cd0() { cd0() }

// Pre: d enthält ggf. einen path und am Ende einen zulässigen Namen für einen
//      Eintrag vom Typ Datei; die Länge dieses Eintrags ist <= maxN - 11.
// d ist mit einem Präfix aus elf Zeichen versehen.
//  func Temp (d *string)

// Pre: p is contained in the file system tree.
//      d is a name of a directory entry.
//      The calling process has the necessary access rights.
// If there is an entry with name d in the target directory of p,
// nothing has happened. Otherwise, d is now contained as directory
// in the file system tree and appended to p.
// The actual path of the calling process is not changed.
func Ins (p, d string) { ins(p,d) }

// Pre: p is contained in the file system tree (s. Defined).
//      d is the name of a directory, that is contained in the
//      target directory of p and that directory is empty.
//      d is not the actual directory of any process.
//      The calling process has the necessary rights.
// d is deleted from the target directory of p, hence
// not any more contained in the file system tree.
// The actual path of the calling process is not changed.
func Del (p, d string) { del(p,d) }

// Returns true, iff there are no entries in the actual directory.
func Empty () bool { return empty() }

// Returns the number of entries in the actual directory.
func Num () uint { return num() }

// Returns the type of the entry and true, iff there is an entry
// with name n in the actual directory; otherwise Unknown and false.
func Typ (n string) (Type, bool) { return typ(n) }

// Pre: i < number of entries in the actual directory.
// Returns name and type of the i-th entry (starting with 0)
// in the actual directory and, if this entry is a file,
// its length, i.e. the number of bytes in it, otherwise 0.
// (otherwise "", Unknown and 0)
func Entry (i uint) (string, Type, uint64) { return entry(i) }

// Returns true, iff there is no entry of type t in the actual directory.
func Empty1 (t Type) bool { return empty1(t) }

// Returns the number of entries of type t in the actual directory.
func Num1 (t Type) uint { return num1(t) }

// Returns true, iff there is an entry with name n of type t in the actual directory.
func Contained1 (n string, t Type) bool { return contained1(n,t) }

// Pre: i < number of entries of type t in the actual directory.
// Returns the name of the i-th entry (starting with 0) of type t
// in the actual directory, and, if this entry is a file, its length, otherwise 0.
func Name1 (t Type, i uint) (string, uint64) { return name1(t,i) }

// Returns the number of nonempty files in the actual directory,
// that have names, for which p returns true.
func NumPred (p Pred) uint { return numPred(p) }

// Pre: NumPred(p) was called after the last Wechsel of the actual directory,
// i < result of this call.
// Returns the name of the i-th of the files in the actual directory,
// that have names, for which p returns true.
func NamePred (p Pred, i uint) string { return namePred(p,i) }
