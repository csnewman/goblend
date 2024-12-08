#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT fwrite(SB),NOSPLIT,$0-0
  MOVQ __imp_fwrite(SB), AX
  JMP AX

TEXT abort(SB),NOSPLIT,$0-0
  MOVQ __imp_abort(SB), AX
  JMP AX

TEXT fprintf(SB),NOSPLIT,$0-0
  MOVQ __imp_fprintf(SB), AX
  JMP AX

TEXT malloc(SB),NOSPLIT,$0-0
  MOVQ __imp_malloc(SB), AX
  JMP AX

TEXT free(SB),NOSPLIT,$0-0
  MOVQ __imp_free(SB), AX
  JMP AX
