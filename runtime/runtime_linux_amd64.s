#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT goblend_cgo_sigaction(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_sigaction(SB)
  RET
