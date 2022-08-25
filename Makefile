build:
	@go build -o ApiTest ./main.go

clean:
	@rm -rf ApiTest

h:
	@echo "\t-s: 并发数"