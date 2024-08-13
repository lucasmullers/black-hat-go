run:
	go run hello-world.go
build:
	go build -ldflags "-w -s" hello-world.go
build-run:
	go build -ldflags "-w -s" hello-world.go && ./hello-world
cross-compiling:
	GOOS="linux" GOARCH="amd64" go build -ldflags "-w -s" hello-world.go