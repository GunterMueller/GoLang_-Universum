package kbd

// (c) Christian Maurer   v. 220206 - license see µU.go
//
// >>> Pre: The preconditions of mouse are met.

/* We distinguish between three groups of keys to operate and control a system
   with keyboard and mouse:
   - character-keys (with echo in form of an alphanumerical character on the screen)
     to enter texts and numbers,
   - command-keys
     to induce particular reactions of the system and
   - mouse-buttons and -movements
     to navigate on the screen.
   In order to abstract from concrete keyboards or mouses,
   the following commands are provided for the last two groups: */

type
  Comm byte; const (
  None = Comm(iota)       // to distinguish between character- and command-keys,
                          // see specification of "Read"
  Esc                     // to leave the system (or a part of it) or to reject
  Enter                   // to confirm or to leave an input
  Back                    // to move backwards in the system
  Left; Right; Up; Down   // to move the cursor on the screen and
  PgLeft; PgRight; PgUp; PgDown // to move in the system, e.g. in a screen mask,
  Pos1; End               // in the corresponding direction
  Tab                     // for special purposes
  Del; Ins                // to remove or insert objects
  Help; Search            // to induce context dependent reactions of the system
  Act; Cfg;               // and for special purposes
  Mark; Unmark            // to mark and unmark objects
  Cut; Copy; Paste        // cut buffer operations
  Red; Green; Blue        // to handle colours
  Print; Roll; Pause      // for special purposes
  OnOff; Lower; Louder    // loudspeaker
  Go                      // to move the mouse
  Here; This; That        // to click on objects and
  Drag; Drop; Move        // to move them around with a mouse
  To; There; Thither      // and to drag and drop them
  ScrollUp; ScrollDown    // for the mouse wheel
//  Nav                     // to navigate in space with a 3d-mouse
  NComms                  // number of commands
)

/* Commands may be enforced in the "depth" of their "impact":
   Every command is associated with a natural number as its depth
   (0 as basic version, bigger numbers for greater depths).
   So we allow for commands with conceptionally equal effects
   but variable ranges of "move depth", as e.g. the movement
   in a text to the next character, word, sentence, paragraph or page,
   or in a calendar to the next day, week, month, year, decade.

   Commands of depth 0 are implemented by keys (without metakeys)
   or mouse-actions with system independent semantics:
   - Enter:                   input-key "Enter"/"Return"
   - Esc:                     stop-/break-key "Esc"
   - Back:                    backspace-key "<-"
   - Left, Right, Up, Down:   corresponding arrow-keys
   - PgUp, PgDown, Pos1, End: corresponding keys
   - Tab:                     Tabkey "|<- ->|"
   - Del, Ins:                corresponding keys
   - Help, Search:            F1-, F2-key
   - Act, Cfg:                F3-, F4-key
   - Mark, Unmark:            F5-, F6-key
   - Cut, Copy, Paste:        F7-, F8-, F9-key
   - Red, Green, Blue:        F10-, F11-, F12-key
   - Print, Roll, Pause:      corresponding keys
   - OnOff, Lower, Louder:    corresponding keys on laptops
   - Go:                      mouse moved with no button pressed
   - Here, This, That:        left, right, middle button pressed
   - Drag, Drop, Move:        mouse moved with corresponding button pressed
   - To, There, Thither:      corresponding button released
//   - Navigate:                3d-mouse used

   commands of depth > 0 by combination with metakeys:
   - depth 1:                 Shift- or Strg-key,
   - depth 2:                 Alt-key,
   - depth 3:                 Ctrl- and Alt-key.

   >>> Under X some meta-key/key-combinations are eaten up
       by the window-manager, e.g. with Esc, Tab and Back. */

// The calling process was blocked, until the keyboard buffer was not empty.
// Returns a tripel (b, c, d) with the following properties:
// Either c == None and the first object from the keyboard buffer is the byte b
// or b == 0 and the first object of the keyboard buffer is the command c of depth d.
// This object is now removed from the keyboard buffer.
// If there is no mouse, then c < Go.
func Read() (byte, Comm, uint) { return read() }

// The calling process was blocked, until there is a byte in the keyboard buffer.
// Returns the first byte from the keyboard buffer.
// This byte is deleted from the keyboard buffer.
func Byte() byte { return byte_() }

// The calling process is blocked, until there is a command in the keyboard buffer.
// Returns the first command and its depth from the keyboard buffer.
// This command is deleted from the keyboard buffer.
func Command() (Comm, uint) { return command() }

// Returns a string, describing the calling Command.
func (c Comm) String() string { return text[c] }

// Returns the movement- and rotationvalues of the last command Navigate.
// The velues are all 0, if there is no 3d-mouse or this command was not given,.
// func ReadNavi() (spc.GridCoord, spc.GridCoord) { return readNavi() }

// Precondition: A byte or command was read.
// Returns the last read byte, if there is one, otherwise 0.
func LastByte() byte { return lastByte() }

// Precondition: A byte or command was read.
// Returns the last read command, if one was read, otherwise None.
// In the first case, d is the depth of the command, otherwise d = 0.
func LastCommand() (Comm, uint) { return lastCommand() }

// c is stored as last read command.
func DepositCommand (c Comm) { depositCommand(c) }

// b is stored as last read byte.
func DepositByte (b byte) { depositByte(b) }

// The calling process was blocked, until until the keyboard buffer contained
// one of the commands Enter (for b = true) resp. Esc or Back (for b = false).
// This command is now removed from the keyboard buffer.
// Returns true, iff the depth of the command was == 0. 
func Wait (b bool) bool { return wait(b) }

// The calling process was blocked,
// until until the keyboard buffer contained command c with depth d.
// This command is now removed from the keyboard buffer.
func WaitFor (c Comm, d uint) { waitFor(c,d) }

// The calling process was blocked, until until the keyboard buffer
// contained one of the commands Enter, Esc or Back.
// This command is now removed from the keyboard buffer.
func Quit() { quit() }

// Returns true, if the keyboard buffer contained one of the commands Enter or Here,
// and false, if it contained one of the commands // Back or There,
// for b = false of any depth and for b = true of a depth > 0.
// The calling process was blocked, until the keyboard buffer contained
// one of these commands; this command is now deleted from it.
func Confirmed (b bool) bool { return confirmed(b) }
