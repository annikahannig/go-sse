
all: test

test:
	go test -v

coverage:
	go test -coverprofile=coverage.out
	go tool cover -func=coverage.out

coverinspect: coverage
	go tool cover -html=coverage.out

