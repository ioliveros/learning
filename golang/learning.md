## Golang cheatsheet

ref: https://devhints.io/go

## 07-17-2021

1. Download latest stable release from `golang` website `https://golang.org/doc/install#download`
2. Installing `go` is simple, just use this tutorial `https://golang.org/doc/install` 

```sh
# Go to your root directory
ioliveros@devBox ~: cd

# Exstract the archive you downloaded into /usr/local, creating a Go tree in /usr/local/go
ioliveros@devBox ~: rm -rf /usr/local/go && tar -C /usr/local -xzf go1.16.6.linux-amd64.tar.gz

# Add /usr/local/go/bin to the PATH environment variable
ioliveros@devBox ~: export PATH=$PATH:/usr/local/go/bin

# Verify that you've installed Go by opening a command prompt and typing the following command:
go version
# Should see something like:
ioliveros@devBox ~: go version
go version go1.16.6 linux/amd64
```

## 02-20-2021

`go mod init`
- this command is similar to pythons `__init__.py` to recognize that the directory is a module

ref: https://golang.org/doc/tutorial/create-module

## to check for `go path`
```
ioliveros@devBox ~: go env GOPATH
/home/ioliveros/go
```
## to check `GOPATH` env
```
ioliveros@devBox ~: go env $GOPATH
GO111MODULE=""
GOARCH="amd64"
GOBIN=""
GOCACHE="/home/ioliveros/.cache/go-build"
GOENV="/home/ioliveros/.config/go/env"
GOEXE=""
GOFLAGS=""
GOHOSTARCH="amd64"
GOHOSTOS="linux"
GOINSECURE=""
GOMODCACHE="/home/ioliveros/go/pkg/mod"
GONOPROXY=""
GONOSUMDB=""
GOOS="linux"
GOPATH="/home/ioliveros/go"
GOPRIVATE=""
GOPROXY="https://proxy.golang.org,direct"
GOROOT="/usr/local/go"
GOSUMDB="sum.golang.org"
GOTMPDIR=""
GOTOOLDIR="/usr/local/go/pkg/tool/linux_amd64"
GOVCS=""
GOVERSION="go1.16.6"
GCCGO="gccgo"
AR="ar"
CC="gcc"
CXX="g++"
CGO_ENABLED="1"
GOMOD="/dev/null"
CGO_CFLAGS="-g -O2"
CGO_CPPFLAGS=""
CGO_CXXFLAGS="-g -O2"
CGO_FFLAGS="-g -O2"
CGO_LDFLAGS="-g -O2"
PKG_CONFIG="pkg-config"
GOGCCFLAGS="-fPIC -m64 -pthread -fmessage-length=0 -fdebug-prefix-map=/tmp/go-build2483726277=/tmp/go-build -gno-record-gcc-switches"
```

## TO GET ALL GOLANGS SUPPORTED PACKAGES
https://pkg.go.dev/search?
https://pkg.go.dev/golang.org/x/tools


## TO INSTALL PACKAGE 
```
ioliveros@devBox ~: go get github.com/aws/aws-sdk-go
go: downloading github.com/aws/aws-sdk-go v1.40.1
go: downloading github.com/jmespath/go-jmespath v0.4.0
```