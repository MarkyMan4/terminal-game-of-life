BIN=gol

build:
	go build -o gol main.go

clean:
	rm $(BIN)
