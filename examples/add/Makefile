build-all: linux-arm64 windows-amd64 darwin-arm64

linux-arm64:
	zig build-obj -O ReleaseFast -target aarch64-linux-gnu --library c -femit-bin=core_linux_arm64.syso add.zig

windows-amd64:
	zig build-obj -O ReleaseFast -target x86_64-windows-gnu --library c -femit-bin=core_windows_amd64.syso add.zig

darwin-arm64:
	zig build-obj -O ReleaseFast -target aarch64-macos-none --library c -femit-bin=core_darwin_arm64.syso add.zig
