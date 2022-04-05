# Remove escape html

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/gowizzard/rmeschtml.svg)](https://golang.org/) [![Go](https://github.com/gowizzard/rmeschtml/actions/workflows/go.yml/badge.svg)](https://github.com/gowizzard/rmeschtml/actions/workflows/go.yml) [![Go Report Card](https://goreportcard.com/badge/github.com/gowizzard/rmeschtml)](https://goreportcard.com/report/github.com/gowizzard/rmeschtml) [![Go Reference](https://pkg.go.dev/badge/github.com/gowizzard/rmeschtml.svg)](https://pkg.go.dev/github.com/gowizzard/rmeschtml) [![GitHub issues](https://img.shields.io/github/issues/gowizzard/rmeschtml)](https://github.com/gowizzard/rmeschtml/issues) [![GitHub forks](https://img.shields.io/github/forks/gowizzard/rmeschtml)](https://github.com/gowizzard/rmeschtml/network) [![GitHub stars](https://img.shields.io/github/stars/gowizzard/rmeschtml)](https://github.com/gowizzard/rmeschtml/stargazers) [![GitHub license](https://img.shields.io/github/license/gowizzard/rmeschtml)](https://github.com/gowizzard/rmeschtml/blob/master/LICENSE)

This small library should help to replace the html escape characters with plain text. I used [this page](https://www.web2generators.com/html-based-tools/online-html-entities-encoder-and-decoder) as a starting point for the list.

## Install

First you have to install the package. You can do this as follows:

```console
go get github.com/gowizzard/rmeschtml
```

## How to use

Here is a small example how to remove the escaped characters from a string.

```go
text := "&lt;mail@gowizzard.de&gt;&nbsp;&quot;Jonas Kwiedor&quot;"
fmt.Println(rmeschtml.Replace(text))
```
