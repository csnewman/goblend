#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT goblend_cgo_init(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_init(SB)
  RET

TEXT goblend_cgo_thread_start(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_thread_start(SB)
  RET

TEXT goblend_cgo_bindm(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_bindm(SB)
  RET

TEXT goblend_cgo_notify_runtime_init_done(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_notify_runtime_init_done(SB)
  RET

TEXT goblend_cgo_set_context_function(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_set_context_function(SB)
  RET

TEXT goblend_cgo_getstackbound(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_getstackbound(SB)
  RET
