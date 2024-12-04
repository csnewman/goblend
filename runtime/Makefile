linux-arm64:
	zig build-obj -O ReleaseFast -target aarch64-linux-gnu --library c -femit-bin=runtime_linux_arm64.syso \
		gcc_libinit.c gcc_stack_unix.c gcc_linux_arm64.c gcc_arm64.S gcc_util.c gcc_setenv.c linux_syscall.c \
		gcc_mmap.c gcc_sigaction.c
