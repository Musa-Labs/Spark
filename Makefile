build:
	go build -o spark cmd/... main.go

run: build
	./spark

clean:
	rm -f spark
