TARGET=go-cmd-sample

go-cmd-sample: main.go
	go build -o $(TARGET) -ldflags "-X main.version=$(shell cat VERSION)" main.go
clean:
	rm -f $(TARGET)
