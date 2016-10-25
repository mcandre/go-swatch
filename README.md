# go-swatch - a Go port of Internet Time AKA Swatch Beats

# EXAMPLE

```
$ swatch
@721.48
```

# REQUIREMENTS

* [Go](https://golang.org/) 1.7+

## Optional

* [Git](https://git-scm.com)
* [Make](https://www.gnu.org/software/make/)
* [Bash](https://www.gnu.org/software/bash/)

# INSTALL FROM REMOTE GIT REPOSITORY

```
$ go get github.com/mcandre/go-swatch/...
```

(Yes, include the ellipsis as well, it's the magic Go syntax for downloading, building, and installing all components of a package, including any libraries and command line tools.)

# INSTALL FROM LOCAL GIT REPOSITORY

```
$ mkdir -p $GOPATH/src/github.com/mcandre
$ git clone git@github.com:mcandre/go-swatch.git $GOPATH/src/github.com/mcandre/go-swatch
$ sh -c "cd $GOPATH/src/github.com/mcandre/go-swatch/cmd/swatch && go install"
```

# LINT

Keep the code tidy:

```
$ make lint
```
