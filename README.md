# `gjson-benchmarks`

Benchmarks for [GJSON](https://github.com/tidwall/gjson)
alongside [encoding/json](https://golang.org/pkg/encoding/json/),
[ffjson](https://github.com/pquerna/ffjson),
[EasyJSON](https://github.com/mailru/easyjson),
[jsonparser](https://github.com/buger/jsonparser),
and [json-iterator](https://github.com/json-iterator/go)

```
BenchmarkGJSONGet-16                11644512       311 ns/op       0 B/op	       0 allocs/op
BenchmarkGJSONUnmarshalMap-16        1122678      3094 ns/op    1920 B/op	      26 allocs/op
BenchmarkJSONUnmarshalMap-16          516681      6810 ns/op    2944 B/op	      69 allocs/op
BenchmarkJSONUnmarshalStruct-16       697053      5400 ns/op     928 B/op	      13 allocs/op
BenchmarkJSONDecoder-16               330450     10217 ns/op    3845 B/op	     160 allocs/op
BenchmarkFFJSONLexer-16              1424979      2585 ns/op     880 B/op	       8 allocs/op
BenchmarkEasyJSONLexer-16            3000000       729 ns/op     501 B/op	       5 allocs/op
BenchmarkJSONParserGet-16            3000000       366 ns/op      21 B/op	       0 allocs/op
BenchmarkJSONIterator-16             3000000       869 ns/op     693 B/op	      14 allocs/op
```

JSON document used:

```json
{
  "widget": {
    "debug": "on",
    "window": {
      "title": "Sample Konfabulator Widget",
      "name": "main_window",
      "width": 500,
      "height": 500
    },
    "image": {
      "src": "Images/Sun.png",
      "hOffset": 250,
      "vOffset": 250,
      "alignment": "center"
    },
    "text": {
      "data": "Click Here",
      "size": 36,
      "style": "bold",
      "vOffset": 100,
      "alignment": "center",
      "onMouseUp": "sun1.opacity = (sun1.opacity / 100) * 90;"
    }
  }
}    
```

Each operation was rotated though one of the following search paths:

```
widget.window.name
widget.image.hOffset
widget.text.onMouseUp
```

For the `GetMany` benchmarks these paths are used:

```
widget.window.name
widget.image.hOffset
widget.text.onMouseUp
widget.window.title
widget.image.alignment
widget.text.style
widget.window.height
widget.image.src
widget.text.data
widget.text.size
```

*These benchmarks were run on a MacBook Pro 16" 2.4 GHz Intel Core i9 using Go 1.17.*

Last run: Oct 22, 2021

## Usage

```sh
go test -v bench .
```
