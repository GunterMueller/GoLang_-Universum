// package atomic

// (c) Christian Maurer   v. 190328 - license see µU.go

#include "textflag.h"

TEXT ·TestAndSet(SB),NOSPLIT,$0
  MOVL a+0(FP), BP
  MOVL $1, AX
  LOCK
  XCHGL AX, 0(BP)
  MOVL AX, ret+4(FP)
  RET

TEXT ·Exchange(SB),NOSPLIT,$0
  MOVL n+0(FP), BX
  MOVL k+4(FP), AX
  XCHGL	AX, 0(BX)
  MOVL AX, ret+8(FP)
  RET

TEXT ·CompareAndSwap(SB),NOSPLIT,$0
  MOVL n+0(FP), BX
  MOVL k+4(FP), AX
  MOVL m+8(FP), CX
  LOCK
  CMPXCHGL  CX, 0(BX)
  SETEQ ret+12(FP)
  RET

TEXT ·FetchAndIncrement(SB),NOSPLIT,$0
  MOVL n+0(FP), BP
  MOVL $1, AX
  LOCK
  XADDL AX, 0(BP)
  MOVL AX, ret+4(FP)
  RET

TEXT ·Add(SB),NOSPLIT,$0
  MOVL n+0(FP), BP
  MOVL k+4(FP), AX
  LOCK
  XADDL AX, 0(BP)
  RET

TEXT ·FetchAndAdd(SB),NOSPLIT,$0
  MOVL n+0(FP), BP
  MOVL k+4(FP), AX
  LOCK
  XADDL AX, 0(BP)
  MOVL AX, ret+8(FP)
  RET

TEXT ·Decrement(SB),NOSPLIT,$0
  MOVL n+0(FP), BP
  LOCK
  DECL 0(BP)
  SETMI ret+4(FP)
  RET

TEXT ·Store(SB),NOSPLIT,$0
  MOVL n+0(FP), BX
  MOVL k+4(FP), AX
  LOCK
  XCHGL AX, 0(BX)
  RET
