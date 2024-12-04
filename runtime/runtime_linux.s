#include "go_asm.h"
#include "funcdata.h"
#include "textflag.h"

TEXT goblend_cgo_libc_setegid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setegid(SB)
  RET

TEXT goblend_cgo_libc_seteuid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_seteuid(SB)
  RET

TEXT goblend_cgo_libc_setregid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setregid(SB)
  RET

TEXT goblend_cgo_libc_setresgid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setresgid(SB)
  RET

TEXT goblend_cgo_libc_setresuid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setresuid(SB)
  RET

TEXT goblend_cgo_libc_setreuid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setreuid(SB)
  RET

TEXT goblend_cgo_libc_setgroups(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setgroups(SB)
  RET

TEXT goblend_cgo_libc_setgid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setgid(SB)
  RET

TEXT goblend_cgo_libc_setuid(SB),NOSPLIT|NOFRAME,$0-0
  JMP _cgo_libc_setuid(SB)
  RET

TEXT goblend_cgo_mmap(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_mmap(SB)
  RET

TEXT goblend_cgo_munmap(SB),NOSPLIT|NOFRAME,$0-0
  JMP x_cgo_munmap(SB)
  RET
