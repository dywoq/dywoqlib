#include "textflag.h"
#include "go_asm.h"

TEXT ·String(SB), NOSPLIT, $0-24
	MOVQ ptr+0(FP), AX
    MOVQ len+8(FP), CX
    MOVQ AX, ret+16(FP)
    MOVQ CX, ret+24(FP)
	RET
	