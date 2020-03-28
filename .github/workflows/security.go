on: [push, pull_request]
name: Gosec
jobs:
  test:
    strategy:
      matrix:
        go-version: [1.13.x, 1.14.x]
        platform: [ubuntu-latest]
    runs-on: ${{ matrix.platform }}
    steps:
    - name: Install Go
      uses: actions/setup-go@v1
      with:
        go-version: ${{ matrix.go-version }}
    - name: Checkout code
      uses: actions/checkout@v2
    - name: Security
      run: go get github.com/securego/gosec/cmd/gosec; `go env GOPATH`/bin/gosec ./...
