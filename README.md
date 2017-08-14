# `gjson-benchmarks`

Benchmarks for [GJSON](https://github.com/tidwall/gjson-benchmarks)
alongside [encoding/json](https://golang.org/pkg/encoding/json/), 
[ffjson](https://github.com/pquerna/ffjson), 
[EasyJSON](https://github.com/mailru/easyjson),
[jsonparser](https://github.com/buger/jsonparser),
and [json-iterator](https://github.com/json-iterator/go)

```
BenchmarkGJSONGet-8                  3000000        372 ns/op          0 B/op         0 allocs/op
BenchmarkGJSONUnmarshalMap-8          900000       4154 ns/op       1920 B/op        26 allocs/op
BenchmarkJSONUnmarshalMap-8           600000       9019 ns/op       3048 B/op        69 allocs/op
BenchmarkJSONUnmarshalStruct-8        600000       9268 ns/op       1832 B/op        69 allocs/op
BenchmarkJSONDecoder-8                300000      14120 ns/op       4224 B/op       184 allocs/op
BenchmarkFFJSONLexer-8               1500000       3111 ns/op        896 B/op         8 allocs/op
BenchmarkEasyJSONLexer-8             3000000        887 ns/op        613 B/op         6 allocs/op
BenchmarkJSONParserGet-8             3000000        499 ns/op         21 B/op         0 allocs/op
BenchmarkJSONIterator-8              3000000        812 ns/op        544 B/op         9 allocs/op
```

Benchmarks for the `GetMany` function:

```
BenchmarkGJSONGetMany4Paths-8        4000000       303 ns/op         112 B/op         0 allocs/op
BenchmarkGJSONGetMany8Paths-8        8000000       208 ns/op          56 B/op         0 allocs/op
BenchmarkGJSONGetMany16Paths-8      16000000       156 ns/op          56 B/op         0 allocs/op
BenchmarkGJSONGetMany32Paths-8      32000000       127 ns/op          64 B/op         0 allocs/op
BenchmarkGJSONGetMany64Paths-8      64000000       117 ns/op          64 B/op         0 allocs/op
BenchmarkGJSONGetMany128Paths-8    128000000       109 ns/op          64 B/op         0 allocs/op
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

*These benchmarks were run on a MacBook Pro 15" 2.8 GHz Intel Core i7 using Go 1.8.*

Last run: May 10, 2017

