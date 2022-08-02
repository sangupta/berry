# berry

Utility functions in Go.

# Hacking

* To build the Go docs locally:
  - `$ godoc -http=:6060`
  - Open http://localhost:6060/pkg/github.com/sangupta/berry

* To run all tests along with code coverage report
  - `$ go test ./... -v -coverprofile coverage.out`
  - `$ go tool cover -html=coverage.out`

* To publish the Go module:
  - `$ git tag v0.x.0`
  - `$ git push origin v0.x.0`
  - `$ GOPROXY=proxy.golang.org go list -m github.com/sangupta/berry@v0.x.0`


# License

MIT License. Copyright (c) 2022, Sandeep Gupta.
