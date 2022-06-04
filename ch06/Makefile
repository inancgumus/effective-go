compile:
	# compile it for linux
	GOOS=linux GOARCH=amd64 go build -o ./bin/hit_linux_amd64 ./cmd/hit
	# compile it for macOS
	GOOS=darwin GOARCH=amd64 go build -o ./bin/hit_darwin_amd64 ./cmd/hit
	# compile it for Apple M1
	GOOS=darwin GOARCH=arm64 go build -o ./bin/hit_darwin_arm64 ./cmd/hit
	# compile it for Windows
	GOOS=windows GOARCH=amd64 go build -o ./bin/hit_win_amd64.exe ./cmd/hit
