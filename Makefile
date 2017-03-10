
build: clean
	go build -o bin/codecover-reader src/lucassabreu/cmd/codecover-reader/main.go

clean:
	rm -rf bin

