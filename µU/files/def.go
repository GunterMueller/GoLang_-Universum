package files

// (c) Christian Maurer   v. 210329 - license see µU.go

import
  . "µU/obj"
const
  maxN = 255
type
  Type = byte
const (// Types of entries in the file system tree:
  Unknown = Type(iota)
  Fifo
  Device
  Dir
  File
  Link
  Socket
  NTypes
)
// Paths are sequences of names of directories in the file system tree
// - separated from each other by the separator "/" - in the direction
// from the root of the file system tree towards its leaves,
// where always the next directory is contained in the previous one.
// If the first character of a path is the separator, it starts
// in the root directory and is called an absolute Path,
// otherwise it starts in the actual directory.
// A separator at the end of a path does not matter,
// if the path is not only the root directory "/".
// The last directory of a path is called its base directory;
// the base directory of the path "/" is the root directory.
// Every calling process at any time is in exactly one directory,
// which is called the actual directory; the sequence of the
// directories from the root directory to the actual one
// is called the actual path, so the actual directory
// is always the base directory of the actual path.
//
// The empty string "" is also a path; its base directory
// is the actual directory. It is equivalent with "./".
// For a path p and a directory d in p, "p/d/.." is
// equivalent with "p/": both denote the same path.

// Returns the actual path.
// The actual directory is not changed.
func ActualPath() string { return actualPath() }

// Returns the actual directory.
func ActualDir () string { return actualDir() }

// Returns the string denoting the type t
// (e.g. "dir" for DIR, "files" for File etc.)
func TypeString (t Type) string { return typeString(t) }

// Returns true, iff p is a path in the file system tree.
func IsPath (p string) bool { return isPath(p) }

// Returns true, iff d is a directory in the file system tree.
func IsDir (d string) bool { return isDir(d) }

// Returns true, iff f is a file in the file system tree.
func IsFile (f string) bool { return isFile(f) }

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

// Spec: The source is the doc.
func Cdp() { cdp() }

// Spec: The source is the doc.
func Cd0() { cd0() }

// Pre: d is a name of a directory entry.
//      The calling process has the necessary access rights.
// d ist now contained as subdirectory of the actual directory
// in the file system tree.
func Sub (d string) { sub(d) }

// Pre: p is a path and d is a name of a directory entry.
//      The calling process has the necessary access rights.
// Returns true, iff d is now contained as directory
// in the file system tree and appended to p.
// Otherwise nothing has happened.
// The actual path of the calling process is not changed.
func Ins (p, d string) bool { return ins(p,d) }

// Pre: p is a path and n is either the name of an empty directory
//      (not the actual directory of any process) or of a file,
//      that is contained in the base directory of p.
//      The calling process has the necessary rights.
// Returns true, iff n is deleted from the base directory of p.
// Otherwise nothing has happend.
// The actual path of the calling process is not changed.
func Del (p, n string) bool { return del(p,n) }

// Pre: f is a file and d a subdirectory in the actual directory.
// f is moved into d.
func Move (f, d string) { move(f,d) }

// Returns the number of entries in the actual directory.
func Num() uint { return num() }

// Returns the sequence of all names of entries in the actual directory.
func Names() []string { return names() }

// Returns the number of entries of type t in the actual directory.
func Num1 (t Type) uint { return num1(t) }

// Returns the sequence of all names of entries of type t in the actual directory.
func Names1 (t Type) []string { return names() }

// Returns the type of the entry and true, iff there is an entry
// with name n in the actual directory; otherwise Unknown and false.
func Typ (n string) Type { return typ(n) }

// Returns the number of nonempty files in the actual directory,
// that have names, for which p returns true.
func NumPred (p Pred) uint { return numPred(p) }

// Pre: NumPred(p) was called after the last change of the actual directory,
//      i < result of this call.
// Returns the name of the i-th of the files in the actual directory,
// that have names, for which p returns true.
func NamePred (p Pred, i uint) string { return namePred(p,i) }
