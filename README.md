# Hello Go

## Build

```
$ go build
$ ls
README.md   go.mod      hello       hello.go    morestrings
```
## Install
  
```
$ go install example.com/user/hello
or
$ go install .
or
$ go install
```
## Add the install directory to our PATH

```
$ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
```
## Running

```
$ hello
Hello, Go!
```

## Reference
- https://golang.org/doc/code