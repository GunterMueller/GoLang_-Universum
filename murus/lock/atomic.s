// package lock

// (c) Christian Maurer   v. 111111 - license see murus.go


TEXT ·TestAndSet(SB),7,$0
  MOVL valptr+0(FP), BP
  MOVL $1, AX
  LOCK
  XCHGL AX, 0(BP)
  MOVL AX, ret+4(FP)
  RET


TEXT ·ExchangeUint32(SB),7,$0
  MOVL valptr+0(FP), BP
  MOVL valptr+4(FP), AX
  LOCK
  XCHGL AX, (BP)
  MOVL AX, ret+8(FP)
  RET


TEXT ·ExchangeInt32(SB),7,$0
  JMP ·ExchangeUint32(SB)


TEXT ·FetchAndAddUint32(SB),7,$0
  MOVL valptr+0(FP), BP
  MOVL valptr+4(FP), AX
  LOCK
  XADDL AX, (BP)
  MOVL AX, ret+8(FP)
  RET


TEXT ·FetchAndAddInt32(SB),7,$0
  JMP ·FetchAndAddUint32(SB)


TEXT ·DecrementInt32(SB),7,$0
  MOVL valptr+0(FP), BP
  LOCK
  DECL 0(BP)
  SETMI ret+4(FP)
  RET
