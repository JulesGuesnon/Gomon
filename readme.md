# Gomon

  
This package aim to reproduce [nodemon](https://github.com/remy/nodemon) for go

## Installation guide

First get the package

```
go get github.com/julesguesnon/gomon
```

Install it
```
go install github.com/julesguesnon/gomon
```
There you go !

## How to use it ?

For now you can only watch a file, nothing else
```
gomon path/to/my/file.go
```

## Possible issue

If you face this issue:

```
gomon: command not found
```

You may need to add GOPATH to your PATH (you may need to set your GOPATH)

```
export PATH=$PATH:$GOPATH/bin
```
