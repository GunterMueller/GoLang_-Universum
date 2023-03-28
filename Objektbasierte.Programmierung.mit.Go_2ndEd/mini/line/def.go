package line

// (c) Christian Maurer   v. 210405 - license see µU.go

import
  . "µU/obj"
const
  EmptyLabel = byte(' ')
type
  Instruction byte; const (
  NOP = Instruction (iota)
  LDA; STA; LDB; STB; EXA; EXB // Argument: Register
  INA; DEA                     // kein Argument
  INC; DEC                     // Argument: Register
  SHL; SHR                     // Argument: Register
  ADD; ADC; SUB; MUL; DIV      // Argument: Register
  CMP                          // Argument: Register
  JMP; JE; JNE; JC; JNC        // Argument: Marke
  PUSH; POP                    // Argument: Register
  CLC; STC; CMC                // kein Argument
  CALL; RET                    // kein Argument
  nInstructions
)
type
  Line interface { // lines of mini programs, consisting of a label,
                   // an instruction, a register and a target label.
  Clearer
  Equaler
  Stringer

  Write (l, c uint)
  Edit (l, c uint)

// Returns (M, true), iff x starts with label M;
// returns otherwise (EmptyLabel, false).
  Marked() (byte, bool)

// Returns true, iff x contains the instruction CALL.
  IsCall() bool

// Returns true, iff x contains the instuction RET.
  IsRet() bool

// The instruction in x is executed.
  Run() byte
}

func New() Line { return new_() }

// The state of the processor is written to the screen,
// starting at position (line l, column c).
func WriteStatus (l, c uint) { writeStatus(l,c) }

func WriteStack (l, c uint) { writeStack(l,c) }
