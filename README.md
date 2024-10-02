# `gjson-benchmarks`

Benchmarks for [GJSON](https://github.com/tidwall/gjson)
alongside [encoding/json](https://golang.org/pkg/encoding/json/),
[ffjson](https://github.com/pquerna/ffjson),
[EasyJSON](https://github.com/mailru/easyjson),
[jsonparser](https://github.com/buger/jsonparser),
and [json-iterator](https://github.com/json-iterator/go)

```
BenchmarkGJSONGet-10             14919366    240.9 ns/op      0 B/op     0 allocs/op
BenchmarkGJSONUnmarshalMap-10     1663548   2157 ns/op     1920 B/op    26 allocs/op
BenchmarkJSONUnmarshalMap-10       832236   4279 ns/op     2920 B/op    68 allocs/op
BenchmarkJSONUnmarshalStruct-10   1076475   3219 ns/op      920 B/op    12 allocs/op
BenchmarkJSONDecoder-10            585729   6126 ns/op     3845 B/op   160 allocs/op
BenchmarkFFJSONLexer-10           2508573   1391 ns/op      880 B/op     8 allocs/op
BenchmarkEasyJSONLexer-10         3000000    537.9 ns/op    501 B/op     5 allocs/op
BenchmarkJSONParserGet-10        13707510    263.9 ns/op     21 B/op     0 allocs/op
BenchmarkJSONIterator-10          3000000    561.2 ns/op    693 B/op    14 allocs/op
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

*These benchmarks were run on a MacBook Pro M1 Max using Go 1.22*

Last run: Oct 1, 2024

## Usage

```sh
go test bench .
```
