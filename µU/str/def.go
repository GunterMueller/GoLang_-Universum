package str

// (c) Christian Maurer   v. 230311 - license see µU.go

//     latin-1-strings (without any UTF-8-stuff)

// Returns true, iff s contains UTF8-runes.
func DevilsDung (s string) bool { return devilsDung(s) }

// Returns a string of n spaces.
func New (n uint) string { return new_(n) }

// s consists only of spaces (in its original length).
func Clr (s *string) { clr(s) }

// Returns a string, then coincides with s up to certain UTF-8-runes
// (see constants in char.go), that are changed to latin1-bytes.
func Lat1 (s string) string { return lat1(s) }

// Invers to Lat1.
func UTF8 (s string) string { return utf8(s) }

// Returns true, iff s has the form x{x|y}, where x = [A-Z]|[a-z], y = x|[0-9]|
func Lit (s string) bool { return lit(s) }

// Returns true, iff s contains no bytes or only spaces.
func Empty (s string) bool { return empty(s) }

// Returns a string of n 'c's.
func Const (c byte, n uint) string { return const_(c,n) }

// Returns the number of bytes of s without considering trailing spaces.
func ProperLen (s string) uint { return properLen(s) }

// Returns true, iff s and t coincide up to trailing spaces,
// i.e. if they contains the same bytes in the same order,
// where spaces at their ends are not considered.
func Eq (s, t string) bool { return eq(s,t) }

// All/the first / small letter(s)/capital(s) of s are/is
// replaced by the corresponding capital(s)/small letter(s).
func ToUpper0 (s *string) { toUpper0(s) }
func ToLower0 (s *string) { toLower0(s) }
func ToUpper (s *string) { toUpper(s) }
func ToLower (s *string) { toLower(s) }

// Returns true, iff the first character of s is no letter or a capital 'A'..'Z'.
func Cap0 (s string) bool { return cap0(s) }

// Returns true, iff s and t are equal up to trailing spaces and
// up to the difference between small letters and corresponding capitals.
func Equiv (s, t string) bool { return equiv(s,t) }

// func QuasiEquiv (s, t string) <-- deprecated, use Equiv

// Returns true, iff s is lexicographically before t
// (particularly an empty string is Less than a nonempty one).
// Capitals are less than the corresponding small letters.
// TODO: The problem of the equivalences 'ä'/"ae", 'ö'/"oe",
// 'ü'/"ue", 'Ä'/'Ae', 'Ö'/"", 'Ü'/"Ue" and 'ß'/"ss" and 
// that of characters with "deadkeys" is not yet solved.
func Less (s, t string) bool { return less(s,t) }

// Returns true, iff s is lexicographically before t, where
// capitals and corresponding small letters are identified.
func EquivLess (s, t string) bool { return equivLess(s,t) }

// Returns (p, true), if c occurs at position p in s,
// returns otherwise (uint(len (s)), false).
func Pos (s string, c byte) (uint, bool) { return pos(s,c) }

// Returns (p, true), iff c occurs at position p in s
// up to the difference between small letters and corresponding capitals
// (hence e.g. for c = 'x', if 'X' occurs in s, and for c = 'X',
// if 'x' occurs in s). In this case p is the index of the first occurrence
// of the corresponding byte in s; otherwise (uint(len(s)), false).
func EquivPos (s string, c byte) (uint, bool) { return equivPos(s,c) }

// Returns true, iff s is contained as trailing part in t.
func Sub0 (s, t string) bool { return sub0(s,t) }

// Returns (p, true), iff s is contained as connected part in t.
// In this case p is the position in t, at which s starts;
// otherwise p == len(t).
func Sub (s, t string) (uint, bool) { return sub(s,t) }

// Returns (n, p), iff s is n-times contained as connected part in t,
// where the p's are the start positions of s;
// returns otherwise 0, []uint{}.
func SubAll (s, t string) (uint, []uint) { return subAll(s,t) }

// Returns (p, true), iff s is up to trailing spaces and
// up to the difference between small letters and corresponding capitals
// contained as connected part in t.
// In this case p is the index in s, at which t starts;
// otherwise p == len(t).
func EquivSub (s, t string) (uint, bool) { return equivSub(s,t) }

// func QuasiEquivSub (s, t string) bool <-- deprecated, use EquivSub

// If p < len (s), the c is inserted in s as p-th character;
// otherwise c is appended at the end of s.
func Ins1 (s *string, c byte, p uint) { ins(s,string(c),p) }

// If p < len(s), then t is inserted into s, startig at position p,
// i.e. s consists of the first p bytes of s before, then t, and then
// the bytes of s starting at p; otherwise t is appended to s.
func Ins (s *string, t string, p uint) { ins(s,t,p) }

// Every occurence of v in *s is replaced by t.
func InsAll (s *string, v, t string) { insAll(s,v,t) }

// If p < len(s), the byte at position p of s is replaced by c.
// Otherwise s is not changed.
func Replace1 (s *string, p uint, c byte) { replace1 (s,p,c) }

// All bytes b in s are replaced by []byte(t).
func ReplaceAll (s *string, b byte, t string) { replaceAll (s,b,t) }

// If p + len(t) <= len(s), the part of s starting at position p is replaced by t.
// Otherwise s is not changed.
func Replace (s *string, p uint, t string) { replace (s,p,t) }

// TODO Spec
func Append (s *string, b byte) { app(s,b) }

// If s contains at least p + n bytes, then n bytes, beginning
// at position p, otherwise all bytes, are removed from s.
// Otherwise s is unchanged.
func Rem (s *string, p, n uint) { rem(s,p,n) }

// If s contains bytes equal to b, the first of these bytes
// is removed from s; return s.
func Del (s string, b byte) string { return del(s,b) }

// Returns the string consisting of n bytes of s,
// beginning at position p, if s contains at least p + n bytes.
// If not, the result is correspondingly shorter.
func Part (s string, p, n uint) string { return part(s,p,n) }

// If len(s) > n, the bytes from position n on are removed from s;
// if len(s) < n, s is filled up with trailing spaces to length n.
// In case len(s) == n, nothing is changed.
func Norm (s *string, n uint) { norm(s,n) }

// As far as existent, all b's are removed from s.
func OffBytes (s *string, b byte) { offBytes(s,b) }

// As far as existent, for left resp. !left leading resp. trailing spaces
// are removed from s and s is filled up with spaces on the other side
// to its former length.
func Move (s *string, left bool) { move(s,left) }

// As far as existent, all spaces are removed from s.
func OffSpc (s *string) { offSpc(s) }

// As far as existent, all trailing spaces are removed from s.
func OffSpc1 (s *string) { offSpc1(s) }

// If p < len(s), a space is inserted in at position p
// (if the last byte of s was not a space, it is lost).
// Otherwise nothing has happened.
func InsSpace (s *string, p uint) { insSpace(s,p) }

// If len(s) <= 1 or p + 1 >= len() or the last byte of s is not a space,
// nothing has happened. Otherwise, a space is inserted in s at position p
// (so no non-space byte is lost).
func Shift (s *string, p uint) { shift(s,p) }

// If n >= len(*s), then *s is on its left and right equally filled up with spaces,
// otherwise *s is cut down to length n.
func Center (s *string, n uint) { center(s,n) }

// All nondigits are removed from s; the digits are shifted to the left
// and s is filled up with spaces to its original length.
func OffNondigits (s *string) { offNondigits(s) }

// Returns the number of all connected strings (= strings without spaces) in s,
// those strings and their start positions in s.
func Split (s string) (uint, []string, []uint) { return split(s) }

// A linefeed (byte(10)) is appended to s.
func AppendLF (s *string) { appendLF(s) }

// Returns the leading part of *s until (excluding) the first byte b with b < ' ';
// from *s now that part and all leading bytes b with b < ' ' are removed.
func SplitLine (s *string) string { return splitLine(s) }

// Pre: b is one of '(', '[', '{', '<'
// My work is so secret that even I do not now what I am doing.
func SplitBrackets (s string, sep, b byte) []string { return splitBrackets (s,sep,b) }

// TODO Spec
func SplitByte (s string, b byte) ([]string, uint) { return splitByte(s,b) }

// Pre: s[0] == '('.
// TODO Spec
func RightBr (s string) uint { return rightBr (s) }

// TODO Spec
func StartsWithVar (s string) (string, uint, bool) { return startsWithVar(s) }

// TODO Spec
func TeX (s string) string { return tex(s) }

// Returns true, iff s contains only letters, digits and spaces.
func Alphanumeric (s string) bool { return alphanumeric(s) }
