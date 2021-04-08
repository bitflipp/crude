all: clean parse generate

clean:
	rm -rf crude.toml crude.go

parse:
	go run ../cmd/crude_parser -i entities.go -o crude.toml

generate:
	go run ../cmd/crude -o crude.go
