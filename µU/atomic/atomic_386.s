// package atomic

// (c) Christian Maurer   v. 171024 - license see µU.go

#include "textflag.h"

TEXT ·TestAndSet(SB),NOSPLIT,$0
  MOVL a+0(FP), BP
  MOVL $1, AX
  LOCK
  XCHGL AX, 0(BP)
  MOVL AX, b+4(FP)
  RET

TEXT ·FetchAndAdd(SB),7,$0
  MOVL n+0(FP), BP
  MOVL k+4(FP), AX
  LOCK
  XADDL AX, (BP)
  MOVL AX, m+8(FP)
  RET

TEXT ·Decrement(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  LOCK
  DECL 0(BP)
  SETMI b+4(FP)
  RET
