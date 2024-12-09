build-all:
	cd runtime && make build-all
	cd examples/add && make build-all
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o add_linux_arm64 ./examples/add
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o add_darwin_arm64 ./examples/add
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o add_windows_amd64.exe ./examples/add
