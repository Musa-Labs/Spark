build:
	go build -o spark main.go

run: build
	./spark

clean:
	rm -f spark
