all: generate

clean:
	rm -rf crude.toml crude.go

parse: clean
	go run ../cmd/crude_parser -i entities.go -o crude.toml

generate: parse
	go run ../cmd/crude -o crude.go
