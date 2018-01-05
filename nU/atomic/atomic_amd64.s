// package atomic

// (c) Christian Maurer   v. 171231 - license see nU.go

#include "textflag.h"

TEXT ·TestAndSet(SB),NOSPLIT,$0
  MOVQ a+0(FP), BP
  MOVL $1, AX
  LOCK
  XCHGL AX, 0(BP)
  MOVL AX, b+8(FP)
  RET

TEXT ·FetchAndAdd(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  MOVL k+8(FP), AX
  LOCK
  XADDL AX, 0(BP)
  MOVL AX, m+16(FP)
  RET

TEXT ·Decrement(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  LOCK
  DECL 0(BP)
  SETMI b+8(FP)
  RET
