# Gomon

  
This package aim to reproduce the behavior of [nodemon](https://github.com/remy/nodemon) for go.
I made this for training purpose so it's probably not really usable.

## Installation guide

Install the package
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
