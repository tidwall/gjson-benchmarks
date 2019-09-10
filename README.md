# `gjson-benchmarks`

> Ran: 2019-09-10 on a Intel(R) Core(TM) i7-4790K CPU @ 4.00GHz (cap/boost 4.6GHz), 4 cores. Go version 1.13

Benchmarks for [GJSON](https://github.com/tidwall/gjson)
alongside [encoding/json](https://golang.org/pkg/encoding/json/),
[ffjson](https://github.com/pquerna/ffjson),
[EasyJSON](https://github.com/mailru/easyjson),
[jsonparser](https://github.com/buger/jsonparser),
and [json-iterator](https://github.com/json-iterator/go)

```
name                    time/op        bytes/op       allocs/op
GJSONGet-8               302ns ± 1%        0B           0.0     
GJSONUnmarshalMap-8     3.55µs ± 3%    1.92kB ± 0%     26.0 ± 0%
GJSONUnmarshalStruct-8  3.31µs ± 2%      992B ± 0%      4.0 ± 0%
JSONUnmarshalMap-8      7.12µs ± 1%    2.98kB ± 0%     69.0 ± 0%
JSONUnmarshalStruct-8   4.84µs ± 2%      912B ± 0%     12.0 ± 0%
JSONDecoder-8           10.9µs ± 1%    4.03kB ± 0%    160.0 ± 0%
FFJSONLexer-8           2.69µs ± 7%      896B ± 0%      8.0 ± 0%
EasyJSONLexer-8          846ns ± 2%      501B ± 0%      5.0 ± 0%
JSONParserGet-8          332ns ±10%        0B           0.0     
JSONIterator-8           914ns ± 1%      677B ± 0%     14.0 ± 0%
ConvertNone-8           18.7µs ± 3%        0B           0.0     
ConvertGet-8            29.8µs ± 2%    49.2kB ± 0%      1.0 ± 0%
ConvertGetBytes-8       18.7µs ± 2%     48.0B ± 0%      1.0 ± 0% 
```

Benchmarks for the `GetMany` function:

```
name                     time/op        bytes/op        allocs/op
GJSONGetMany4Paths-8     313ns ± 2%     56.0B ± 0%      0.00     
GJSONGetMany8Paths-8     324ns ± 1%     56.0B ± 0%      0.00     
GJSONGetMany16Paths-8    337ns ± 1%     56.0B ± 0%      0.00     
GJSONGetMany32Paths-8    335ns ± 2%     56.0B ± 0%      0.00     
GJSONGetMany64Paths-8    340ns ± 3%     64.0B ± 0%      0.00     
GJSONGetMany128Paths-8   351ns ± 1%     64.0B ± 0%      0.00     
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

## Usage

> benchstat must be installed

> See output in pretty-results.txt

```sh
go mod download
./run.sh
```
