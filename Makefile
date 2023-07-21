bin: bin/sprox_darwin_amd64 bin/sprox_darwin_arm64 bin/sprox_linux_amd64 
bin: bin/sprox_linux_arm bin/sprox_linux_arm64

bin/sprox_darwin_amd64:
	GOOS=darwin GOARCH=amd64 go build -o bin/sprox_darwin_amd64 cmd/sprox/*.go

bin/sprox_darwin_arm64:
	GOOS=darwin GOARCH=arm64 go build -o bin/sprox_darwin_arm64 cmd/sprox/*.go

bin/sprox_linux_amd64:
	GOOS=linux GOARCH=amd64 go build -o bin/sprox_linux_amd64 cmd/sprox/*.go

bin/sprox_linux_arm:
	GOOS=linux GOARCH=arm go build -o bin/sprox_linux_arm cmd/sprox/*.go

bin/sprox_linux_arm64:
	GOOS=linux GOARCH=arm64 go build -o bin/sprox_linux_arm64 cmd/sprox/*.go

.PHONY: bin