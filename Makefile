build-png2bmp:
	mkdir -p bin
	go build -o bin/png2bmp cmd/png2bmp/main.go

build-jpeg2bmp:
	mkdir -p bin
	go build -o bin/jpeg2bmp cmd/jpeg2bmp/main.go

build: build-png2bmp build-jpeg2bmp