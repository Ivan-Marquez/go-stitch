build:
	export GO111MODULE=on
	go build -o bin/go-stitch auth.go constants.go fetchLogs.go
clean:
	rm -rf ./bin