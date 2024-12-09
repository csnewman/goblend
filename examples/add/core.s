#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT goblend_add(SB),NOSPLIT|NOFRAME,$0-0
  JMP add(SB)
  RET

#ifdef GOOS_windows

TEXT WriteFile(SB),NOSPLIT,$0-0
  MOVQ __imp_WriteFile(SB), AX
  JMP AX

#endif
