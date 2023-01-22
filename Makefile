TARGET=go-cmd-sample

go-cmd-sample: main.go
	go build -v -x -ldflags "-X 'main.version=$(shell cat VERSION)'" -o $(TARGET) main.go
clean:
	rm -f $(TARGET)
