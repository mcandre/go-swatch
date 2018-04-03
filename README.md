# go-swatch - a Go port of Internet Time AKA Swatch Beats

# EXAMPLE

```console
$ swatch
@721.48
```

# DOWNLOAD

https://github.com/mcandre/go-swatch/releases

# DOCUMENTATION

https://godoc.org/github.com/mcandre/go-swatch

# RUNTIME REQUIREMENTS

(None)

# BUILDTIME REQUIREMENTS

* [Go](https://golang.org/) 1.9+

## Recommended

* [Docker](https://www.docker.com/)
* [Mage](https://magefile.org/) (e.g., `go get github.com/magefile/mage`)
* [goimports](https://godoc.org/golang.org/x/tools/cmd/goimports) (e.g. `go get golang.org/x/tools/cmd/goimports`)
* [golint](https://github.com/golang/lint) (e.g. `go get github.com/golang/lint/golint`)
* [errcheck](https://github.com/kisielk/errcheck) (e.g. `go get github.com/kisielk/errcheck`)
* [nakedret](https://github.com/alexkohler/nakedret) (e.g. `go get github.com/alexkohler/nakedret`)
* [goxcart](https://github.com/mcandre/goxcart) (e.g., `github.com/mcandre/goxcart/...`)
* [zipc](https://github.com/mcandre/zipc) (e.g. `go get github.com/mcandre/zipc/...`)

# INSTALL FROM REMOTE GIT REPOSITORY

```console
$ go get github.com/mcandre/go-swatch/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```console
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone https://github.com/mcandre/go-swatch.git $GOPATH/src/github.com/mcandre/go-swatch
$ cd "$GOPATH/src/github.co/mcandre/go-swatch"
$ git submodule update --init --recursive
$ go install ./...
```

# TEST

```console
$ mage test
```

# PORT

```console
$ mage port
```

# LINT

Keep the code tidy:

```console
$ mage lint
```
