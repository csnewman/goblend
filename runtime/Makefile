DARWIN_ARM64_FILES := gcc_libinit.c gcc_stack_darwin.c gcc_darwin_arm64.c gcc_arm64.S gcc_util.c gcc_setenv.c gcc_sigaction.c
LINUX_COMMON_FILES := gcc_libinit.c gcc_stack_unix.c  gcc_util.c gcc_setenv.c linux_syscall.c gcc_mmap.c gcc_sigaction.c
LINUX_ARM64_FILES := $(LINUX_COMMON_FILES) gcc_linux_arm64.c gcc_arm64.S
WINDOWS_COMMON_FILES := gcc_libinit_windows.c gcc_util.c gcc_stack_windows.c gcc_windows_amd64.c
WINDOWS_AMD64_FILES := $(WINDOWS_COMMON_FILES) gcc_amd64.S gcc_windows_amd64.c

build-all: linux-arm64 windows-amd64 darwin-arm64

linux-arm64:
	zig build-obj -O ReleaseFast -target aarch64-linux-gnu --library c -femit-bin=runtime_linux_arm64.syso \
		$(LINUX_ARM64_FILES)

windows-amd64:
	rm -rf ./btmp
	mkdir btmp
	zig build-lib -static -O ReleaseSmall -target x86_64-windows-gnu -fno-stack-protector -fno-stack-check \
		-fno-sanitize-c --library c -femit-bin=btmp/lib.a $(WINDOWS_AMD64_FILES)
	cd btmp && x86_64-w64-mingw32-ar x lib.a
	cd btmp && x86_64-w64-mingw32-ld -o ../runtime_windows_amd64.syso -r *.obj
	rm -rf ./btmp

darwin-arm64:
	rm -rf ./btmp
	mkdir btmp
	zig build-lib -static -O ReleaseSmall -target aarch64-macos-none -fno-stack-protector -fno-stack-check \
		-fno-sanitize-c --library c -femit-bin=btmp/lib.a $(DARWIN_ARM64_FILES)
	cd btmp && zig ar x -U lib.a
	cd btmp && chmod +rw *.o
	cd btmp && find . -name '*.o' -exec sh -c 'cp "$$0" "../runtime_$$(basename $${0%.o})_darwin_arm64.syso"' {} \;
	rm -rf ./btmp
