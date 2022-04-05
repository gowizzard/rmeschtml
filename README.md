# Remove escape html

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
