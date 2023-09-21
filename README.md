# peg-playground

Playground for learning PEG grammar

## Query Parser

PEG grammar written following [Building a simple query parser using PEG in Go](https://prasanthmj.github.io/go/peg-parser-in-go/) article.

The syntax of the query language parsed is:

```
item.name=laptop item.spec.ram > 8gb item.spec.ssd=yes item.spec.ssd.capacity > 512gb sort_by:price/asc
```

### Getting started

Install the [pigeon](https://pkg.go.dev/github.com/mna/pigeon) binary:

```
go get -u github.com/mna/pigeon
```

Generate the parser from the peg grammar:

```
pigeon -o query/parser.go query/parser.peg
```

Run the maingo file:

```
go run ./main.go
```

Run the unit tests:

```
go test -v ./query
```
