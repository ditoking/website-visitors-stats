# Website Visitors Stats

> Project must be imported under {GO_SRC_DIR}/github.com/ditoking (eg: `~/go/src/github.com/ditoking`)

### To run local:
####Go is a dependency, and this project uses ***go modules*** to handle package dependencies.
> Please run (from project root) `go mod vendor` in order to get all dependencies in a local vendor dir.
> 
> To run the app, please run (from project root) `cd ./cmd/website-visitor-stats && go build . && go run main.go`
> 
> Make sure port 8080 is available.

### To run with docker:
> Please run `docker build -t <NAME>:<TAG> && docker run -p 8080:8080 <NAME>:<TAG>`

### To run tests:
`go clean -testcache && go test ./...`