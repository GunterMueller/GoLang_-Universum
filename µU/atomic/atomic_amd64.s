// package atomic

// (c) Christian Maurer   v. 190331 - license see µU.go

#include "textflag.h"

TEXT ·TestAndSet(SB),NOSPLIT,$0
  MOVQ a+0(FP), BP
  MOVL $1, AX
  LOCK
  XCHGL AX, 0(BP)
  MOVL AX, ret+8(FP)
  RET

TEXT ·Exchange(SB),NOSPLIT,$0
  MOVQ n+0(FP), BX
  MOVQ k+8(FP), AX
  XCHGQ	AX, 0(BX)
  MOVQ AX, ret+16(FP)
  RET

TEXT ·CompareAndSwap(SB),NOSPLIT,$0
  MOVQ n+0(FP), BX
  MOVQ k+8(FP), AX
  MOVQ m+16(FP), CX
  LOCK
  CMPXCHGQ CX, 0(BX)
  SETEQ ret+24(FP)
  RET

TEXT ·FetchAndIncrement(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  MOVQ $1, AX
  LOCK
  XADDQ AX, 0(BP)
  MOVQ AX, ret+8(FP)
  RET

TEXT ·Add(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  MOVQ k+8(FP), AX
  LOCK
  XADDQ AX, 0(BP)
  RET

TEXT ·Inc(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  MOVQ $1, AX
  LOCK
  XADDQ AX, 0(BP)
  RET

TEXT ·FetchAndAdd(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  MOVQ k+8(FP), AX
  LOCK
  XADDQ AX, 0(BP)
  MOVQ AX, ret+16(FP)
  RET

TEXT ·Decrement(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  LOCK
  DECQ 0(BP)
  SETMI ret+8(FP)
  RET

TEXT ·Dec(SB),NOSPLIT,$0
  MOVQ n+0(FP), BP
  LOCK
  DECQ 0(BP)
  RET

TEXT ·Store(SB),NOSPLIT,$0
  MOVQ n+0(FP), BX
  MOVQ k+8(FP), AX
  LOCK
  XCHGQ	AX, 0(BX)
  RET
