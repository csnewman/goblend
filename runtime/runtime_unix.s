#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT goblend_cgo_setenv(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_setenv(SB)
  RET

TEXT goblend_cgo_unsetenv(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_unsetenv(SB)
  RET
