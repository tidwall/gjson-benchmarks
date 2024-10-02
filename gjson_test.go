//lint:file-ignore SA3001 calcs against three paths
package gjson_benchmarks

import (
	"bytes"
	"encoding/json"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/buger/jsonparser"
	jsoniter "github.com/json-iterator/go"
	"github.com/mailru/easyjson/jlexer"
	fflib "github.com/pquerna/ffjson/fflib/v1"
	"github.com/tidwall/gjson"
)

var exampleJSON = `{
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
}`

type BenchStruct struct {
	Widget struct {
		Window struct {
			Name string `json:"name"`
		} `json:"window"`
		Image struct {
			HOffset int `json:"hOffset"`
		} `json:"image"`
		Text struct {
			OnMouseUp string `json:"onMouseUp"`
		} `json:"text"`
	} `json:"widget"`
}

var benchPaths = []string{
	"widget.window.name",
	"widget.image.hOffset",
	"widget.text.onMouseUp",
}

var benchManyPaths = []string{
	"widget.window.name",
	"widget.image.hOffset",
	"widget.text.onMouseUp",
	"widget.window.title",
	"widget.image.alignment",
	"widget.text.style",
	"widget.window.height",
	"widget.image.src",
	"widget.text.data",
	"widget.text.size",
}

func BenchmarkGJSONGet(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			if gjson.Get(exampleJSON, benchPaths[j]).Type == gjson.Null {
				t.Fatal("did not find the value")
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkGJSONUnmarshalMap(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			parts := strings.Split(benchPaths[j], ".")
			m, _ := gjson.Parse(exampleJSON).Value().(map[string]interface{})
			var v interface{}
			for len(parts) > 0 {
				part := parts[0]
				if len(parts) > 1 {
					m = m[part].(map[string]interface{})
					if m == nil {
						t.Fatal("did not find the value")
					}
				} else {
					v = m[part]
					if v == nil {
						t.Fatal("did not find the value")
					}
				}
				parts = parts[1:]
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

// func BenchmarkGJSONUnmarshalStruct(t *testing.B) {
// 	t.ReportAllocs()
// 	t.ResetTimer()
// 	for i := 0; i < t.N; i++ {
// 		for j := 0; j < len(benchPaths); j++ {
// 			var s BenchStruct
// 			if err := gjson.Unmarshal([]byte(exampleJSON), &s); err != nil {
// 				t.Fatal(err)
// 			}
// 			switch benchPaths[j] {
// 			case "widget.window.name":
// 				if s.Widget.Window.Name == "" {
// 					t.Fatal("did not find the value")
// 				}
// 			case "widget.image.hOffset":
// 				if s.Widget.Image.HOffset == 0 {
// 					t.Fatal("did not find the value")
// 				}
// 			case "widget.text.onMouseUp":
// 				if s.Widget.Text.OnMouseUp == "" {
// 					t.Fatal("did not find the value")
// 				}
// 			}
// 		}
// 	}
// 	t.N *= len(benchPaths) // because we are running against 3 paths
// }

func BenchmarkJSONUnmarshalMap(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			parts := strings.Split(benchPaths[j], ".")
			var m map[string]interface{}
			if err := json.Unmarshal([]byte(exampleJSON), &m); err != nil {
				t.Fatal(err)
			}
			var v interface{}
			for len(parts) > 0 {
				part := parts[0]
				if len(parts) > 1 {
					m = m[part].(map[string]interface{})
					if m == nil {
						t.Fatal("did not find the value")
					}
				} else {
					v = m[part]
					if v == nil {
						t.Fatal("did not find the value")
					}
				}
				parts = parts[1:]
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkJSONUnmarshalStruct(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			var s BenchStruct
			if err := json.Unmarshal([]byte(exampleJSON), &s); err != nil {
				t.Fatal(err)
			}
			switch benchPaths[j] {
			case "widget.window.name":
				if s.Widget.Window.Name == "" {
					t.Fatal("did not find the value")
				}
			case "widget.image.hOffset":
				if s.Widget.Image.HOffset == 0 {
					t.Fatal("did not find the value")
				}
			case "widget.text.onMouseUp":
				if s.Widget.Text.OnMouseUp == "" {
					t.Fatal("did not find the value")
				}
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkJSONDecoder(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			dec := json.NewDecoder(bytes.NewBuffer([]byte(exampleJSON)))
			var found bool
		outer:
			for {
				tok, err := dec.Token()
				if err != nil {
					if err == io.EOF {
						break
					}
					t.Fatal(err)
				}
				switch v := tok.(type) {
				case string:
					if found {
						// break out once we find the value.
						break outer
					}
					switch benchPaths[j] {
					case "widget.window.name":
						if v == "name" {
							found = true
						}
					case "widget.image.hOffset":
						if v == "hOffset" {
							found = true
						}
					case "widget.text.onMouseUp":
						if v == "onMouseUp" {
							found = true
						}
					}
				}
			}
			if !found {
				t.Fatal("field not found")
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkFFJSONLexer(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			l := fflib.NewFFLexer([]byte(exampleJSON))
			var found bool
		outer:
			for {
				t := l.Scan()
				if t == fflib.FFTok_eof {
					break
				}
				if t == fflib.FFTok_string {
					b, _ := l.CaptureField(t)
					v := string(b)
					if found {
						// break out once we find the value.
						break outer
					}
					switch benchPaths[j] {
					case "widget.window.name":
						if v == "\"name\"" {
							found = true
						}
					case "widget.image.hOffset":
						if v == "\"hOffset\"" {
							found = true
						}
					case "widget.text.onMouseUp":
						if v == "\"onMouseUp\"" {
							found = true
						}
					}
				}
			}
			if !found {
				t.Fatal("field not found")
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func skipCC(l *jlexer.Lexer, n int) {
	for i := 0; i < n; i++ {
		l.Skip()
		l.WantColon()
		l.Skip()
		l.WantComma()
	}
}
func skipGroup(l *jlexer.Lexer, n int) {
	l.WantColon()
	l.Delim('{')
	skipCC(l, n)
	l.Delim('}')
	l.WantComma()
}
func easyJSONWindowName(t *testing.B, l *jlexer.Lexer) {
	if l.String() == "window" {
		l.WantColon()
		l.Delim('{')
		skipCC(l, 1)
		if l.String() == "name" {
			l.WantColon()
			if l.String() == "" {
				t.Fatal("did not find the value")
			}
		}
	}
}
func easyJSONImageHOffset(t *testing.B, l *jlexer.Lexer) {
	if l.String() == "image" {
		l.WantColon()
		l.Delim('{')
		skipCC(l, 1)
		if l.String() == "hOffset" {
			l.WantColon()
			if l.Int() == 0 {
				t.Fatal("did not find the value")
			}
		}
	}
}
func easyJSONTextOnMouseUp(t *testing.B, l *jlexer.Lexer) {
	if l.String() == "text" {
		l.WantColon()
		l.Delim('{')
		skipCC(l, 5)
		if l.String() == "onMouseUp" {
			l.WantColon()
			if l.String() == "" {
				t.Fatal("did not find the value")
			}
		}
	}
}
func easyJSONWidget(t *testing.B, l *jlexer.Lexer, j int) {
	l.WantColon()
	l.Delim('{')
	switch benchPaths[j] {
	case "widget.window.name":
		skipCC(l, 1)
		easyJSONWindowName(t, l)
	case "widget.image.hOffset":
		skipCC(l, 1)
		if l.String() == "window" {
			skipGroup(l, 4)
		}
		easyJSONImageHOffset(t, l)
	case "widget.text.onMouseUp":
		skipCC(l, 1)
		if l.String() == "window" {
			skipGroup(l, 4)
		}
		if l.String() == "image" {
			skipGroup(l, 4)
		}
		easyJSONTextOnMouseUp(t, l)
	}
}
func BenchmarkEasyJSONLexer(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			l := &jlexer.Lexer{Data: []byte(exampleJSON)}
			l.Delim('{')
			if l.String() == "widget" {
				easyJSONWidget(t, l, j)
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

func BenchmarkJSONParserGet(t *testing.B) {
	data := []byte(exampleJSON)
	keys := make([][]string, 0, len(benchPaths))
	for i := 0; i < len(benchPaths); i++ {
		keys = append(keys, strings.Split(benchPaths[i], "."))
	}
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j, k := range keys {
			if j == 1 {
				// "widget.image.hOffset" is a number
				v, _ := jsonparser.GetInt(data, k...)
				if v == 0 {
					t.Fatal("did not find the value")
				}
			} else {
				// "widget.window.name",
				// "widget.text.onMouseUp",
				v, _ := jsonparser.GetString(data, k...)
				if v == "" {
					t.Fatal("did not find the value")
				}
			}
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}
func jsoniterWindowName(t *testing.B, iter *jsoniter.Iterator) {
	var v string
	for {
		key := iter.ReadObject()
		if key != "window" {
			iter.Skip()
			continue
		}
		for {
			key := iter.ReadObject()
			if key != "name" {
				iter.Skip()
				continue
			}
			v = iter.ReadString()
			break
		}
		break
	}
	if v == "" {
		t.Fatal("did not find the value")
	}
}

func jsoniterTextOnMouseUp(t *testing.B, iter *jsoniter.Iterator) {
	var v string
	for {
		key := iter.ReadObject()
		if key != "text" {
			iter.Skip()
			continue
		}
		for {
			key := iter.ReadObject()
			if key != "onMouseUp" {
				iter.Skip()
				continue
			}
			v = iter.ReadString()
			break
		}
		break
	}
	if v == "" {
		t.Fatal("did not find the value")
	}
}
func jsoniterImageOffset(t *testing.B, iter *jsoniter.Iterator) {
	var v int
	for {
		key := iter.ReadObject()
		if key != "image" {
			iter.Skip()
			continue
		}
		for {
			key := iter.ReadObject()
			if key != "hOffset" {
				iter.Skip()
				continue
			}
			v = iter.ReadInt()
			break
		}
		break
	}
	if v == 0 {
		t.Fatal("did not find the value")
	}
}
func jsoniterWidget(t *testing.B, iter *jsoniter.Iterator, j int) {
	for {
		key := iter.ReadObject()
		if key != "widget" {
			iter.Skip()
			continue
		}
		switch benchPaths[j] {
		case "widget.window.name":
			jsoniterWindowName(t, iter)
		case "widget.image.hOffset":
			jsoniterImageOffset(t, iter)
		case "widget.text.onMouseUp":
			jsoniterTextOnMouseUp(t, iter)
		}
		break
	}
}

func BenchmarkJSONIterator(t *testing.B) {
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		for j := 0; j < len(benchPaths); j++ {
			iter := jsoniter.ParseString(jsoniter.ConfigDefault, exampleJSON)
			jsoniterWidget(t, iter, j)
		}
	}
	t.N *= len(benchPaths) // because we are running against 3 paths
}

var massiveJSON = func() string {
	var buf bytes.Buffer
	buf.WriteString("[")
	for i := 0; i < 100; i++ {
		if i > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(exampleJSON)
	}
	buf.WriteString("]")
	return buf.String()
}()

func BenchmarkConvertNone(t *testing.B) {
	json := massiveJSON
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		gjson.Get(json, "50.widget.text.onMouseUp")
	}
}
func BenchmarkConvertGet(t *testing.B) {
	data := []byte(massiveJSON)
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		gjson.Get(string(data), "50.widget.text.onMouseUp")
	}
}
func BenchmarkConvertGetBytes(t *testing.B) {
	data := []byte(massiveJSON)
	t.ReportAllocs()
	t.ResetTimer()
	for i := 0; i < t.N; i++ {
		gjson.GetBytes(data, "50.widget.text.onMouseUp")
	}
}

var twitterMedium = `{
	"statuses": [
	  {
		"coordinates": null,
		"favorited": false,
		"truncated": false,
		"created_at": "Mon Sep 24 03:35:21 +0000 2012",
		"id_str": "250075927172759552",
		"entities": {
		  "urls": [
   
		  ],
		  "hashtags": [
			{
			  "text": "freebandnames",
			  "indices": [
				20,
				34
			  ]
			}
		  ],
		  "user_mentions": [
   
		  ]
		},
		"in_reply_to_user_id_str": null,
		"contributors": null,
		"text": "Aggressive Ponytail #freebandnames",
		"metadata": {
		  "iso_language_code": "en",
		  "result_type": "recent"
		},
		"retweet_count": 0,
		"in_reply_to_status_id_str": null,
		"id": 250075927172759552,
		"geo": null,
		"retweeted": false,
		"in_reply_to_user_id": null,
		"place": null,
		"user": {
		  "profile_sidebar_fill_color": "DDEEF6",
		  "profile_sidebar_border_color": "C0DEED",
		  "profile_background_tile": false,
		  "name": "Sean Cummings",
		  "profile_image_url": "https://a0.twimg.com/profile_images/2359746665/1v6zfgqo8g0d3mk7ii5s_normal.jpeg",
		  "created_at": "Mon Apr 26 06:01:55 +0000 2010",
		  "location": "LA, CA",
		  "follow_request_sent": null,
		  "profile_link_color": "0084B4",
		  "is_translator": false,
		  "id_str": "137238150",
		  "entities": {
			"url": {
			  "urls": [
				{
				  "expanded_url": null,
				  "url": "",
				  "indices": [
					0,
					0
				  ]
				}
			  ]
			},
			"description": {
			  "urls": [
   
			  ]
			}
		  },
		  "default_profile": true,
		  "contributors_enabled": false,
		  "favourites_count": 0,
		  "url": null,
		  "profile_image_url_https": "https://si0.twimg.com/profile_images/2359746665/1v6zfgqo8g0d3mk7ii5s_normal.jpeg",
		  "utc_offset": -28800,
		  "id": 137238150,
		  "profile_use_background_image": true,
		  "listed_count": 2,
		  "profile_text_color": "333333",
		  "lang": "en",
		  "followers_count": 70,
		  "protected": false,
		  "notifications": null,
		  "profile_background_image_url_https": "https://si0.twimg.com/images/themes/theme1/bg.png",
		  "profile_background_color": "C0DEED",
		  "verified": false,
		  "geo_enabled": true,
		  "time_zone": "Pacific Time (US & Canada)",
		  "description": "Born 330 Live 310",
		  "default_profile_image": false,
		  "profile_background_image_url": "https://a0.twimg.com/images/themes/theme1/bg.png",
		  "statuses_count": 579,
		  "friends_count": 110,
		  "following": null,
		  "show_all_inline_media": false,
		  "screen_name": "sean_cummings"
		},
		"in_reply_to_screen_name": null,
		"source": "<a href=\"//itunes.apple.com/us/app/twitter/id409789998?mt=12%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for Mac</a>",
		"in_reply_to_status_id": null
	  },
	  {
		"coordinates": null,
		"favorited": false,
		"truncated": false,
		"created_at": "Fri Sep 21 23:40:54 +0000 2012",
		"id_str": "249292149810667520",
		"entities": {
		  "urls": [
   
		  ],
		  "hashtags": [
			{
			  "text": "FreeBandNames",
			  "indices": [
				20,
				34
			  ]
			}
		  ],
		  "user_mentions": [
   
		  ]
		},
		"in_reply_to_user_id_str": null,
		"contributors": null,
		"text": "Thee Namaste Nerdz. #FreeBandNames",
		"metadata": {
		  "iso_language_code": "pl",
		  "result_type": "recent"
		},
		"retweet_count": 0,
		"in_reply_to_status_id_str": null,
		"id": 249292149810667520,
		"geo": null,
		"retweeted": false,
		"in_reply_to_user_id": null,
		"place": null,
		"user": {
		  "profile_sidebar_fill_color": "DDFFCC",
		  "profile_sidebar_border_color": "BDDCAD",
		  "profile_background_tile": true,
		  "name": "Chaz Martenstein",
		  "profile_image_url": "https://a0.twimg.com/profile_images/447958234/Lichtenstein_normal.jpg",
		  "created_at": "Tue Apr 07 19:05:07 +0000 2009",
		  "location": "Durham, NC",
		  "follow_request_sent": null,
		  "profile_link_color": "0084B4",
		  "is_translator": false,
		  "id_str": "29516238",
		  "entities": {
			"url": {
			  "urls": [
				{
				  "expanded_url": null,
				  "url": "https://bullcityrecords.com/wnng/",
				  "indices": [
					0,
					32
				  ]
				}
			  ]
			},
			"description": {
			  "urls": [
   
			  ]
			}
		  },
		  "default_profile": false,
		  "contributors_enabled": false,
		  "favourites_count": 8,
		  "url": "https://bullcityrecords.com/wnng/",
		  "profile_image_url_https": "https://si0.twimg.com/profile_images/447958234/Lichtenstein_normal.jpg",
		  "utc_offset": -18000,
		  "id": 29516238,
		  "profile_use_background_image": true,
		  "listed_count": 118,
		  "profile_text_color": "333333",
		  "lang": "en",
		  "followers_count": 2052,
		  "protected": false,
		  "notifications": null,
		  "profile_background_image_url_https": "https://si0.twimg.com/profile_background_images/9423277/background_tile.bmp",
		  "profile_background_color": "9AE4E8",
		  "verified": false,
		  "geo_enabled": false,
		  "time_zone": "Eastern Time (US & Canada)",
		  "description": "You will come to Durham, North Carolina. I will sell you some records then, here in Durham, North Carolina. Fun will happen.",
		  "default_profile_image": false,
		  "profile_background_image_url": "https://a0.twimg.com/profile_background_images/9423277/background_tile.bmp",
		  "statuses_count": 7579,
		  "friends_count": 348,
		  "following": null,
		  "show_all_inline_media": true,
		  "screen_name": "bullcityrecords"
		},
		"in_reply_to_screen_name": null,
		"source": "web",
		"in_reply_to_status_id": null
	  },
	  {
		"coordinates": null,
		"favorited": false,
		"truncated": false,
		"created_at": "Fri Sep 21 23:30:20 +0000 2012",
		"id_str": "249289491129438208",
		"entities": {
		  "urls": [
   
		  ],
		  "hashtags": [
			{
			  "text": "freebandnames",
			  "indices": [
				29,
				43
			  ]
			}
		  ],
		  "user_mentions": [
   
		  ]
		},
		"in_reply_to_user_id_str": null,
		"contributors": null,
		"text": "Mexican Heaven, Mexican Hell #freebandnames",
		"metadata": {
		  "iso_language_code": "en",
		  "result_type": "recent"
		},
		"retweet_count": 0,
		"in_reply_to_status_id_str": null,
		"id": 249289491129438208,
		"geo": null,
		"retweeted": false,
		"in_reply_to_user_id": null,
		"place": null,
		"user": {
		  "profile_sidebar_fill_color": "99CC33",
		  "profile_sidebar_border_color": "829D5E",
		  "profile_background_tile": false,
		  "name": "Thomas John Wakeman",
		  "profile_image_url": "https://a0.twimg.com/profile_images/2219333930/Froggystyle_normal.png",
		  "created_at": "Tue Sep 01 21:21:35 +0000 2009",
		  "location": "Kingston New York",
		  "follow_request_sent": null,
		  "profile_link_color": "D02B55",
		  "is_translator": false,
		  "id_str": "70789458",
		  "entities": {
			"url": {
			  "urls": [
				{
				  "expanded_url": null,
				  "url": "",
				  "indices": [
					0,
					0
				  ]
				}
			  ]
			},
			"description": {
			  "urls": [
   
			  ]
			}
		  },
		  "default_profile": false,
		  "contributors_enabled": false,
		  "favourites_count": 19,
		  "url": null,
		  "profile_image_url_https": "https://si0.twimg.com/profile_images/2219333930/Froggystyle_normal.png",
		  "utc_offset": -18000,
		  "id": 70789458,
		  "profile_use_background_image": true,
		  "listed_count": 1,
		  "profile_text_color": "3E4415",
		  "lang": "en",
		  "followers_count": 63,
		  "protected": false,
		  "notifications": null,
		  "profile_background_image_url_https": "https://si0.twimg.com/images/themes/theme5/bg.gif",
		  "profile_background_color": "352726",
		  "verified": false,
		  "geo_enabled": false,
		  "time_zone": "Eastern Time (US & Canada)",
		  "description": "Science Fiction Writer, sort of. Likes Superheroes, Mole People, Alt. Timelines.",
		  "default_profile_image": false,
		  "profile_background_image_url": "https://a0.twimg.com/images/themes/theme5/bg.gif",
		  "statuses_count": 1048,
		  "friends_count": 63,
		  "following": null,
		  "show_all_inline_media": false,
		  "screen_name": "MonkiesFist"
		},
		"in_reply_to_screen_name": null,
		"source": "web",
		"in_reply_to_status_id": null
	  },
	  {
		"coordinates": null,
		"favorited": false,
		"truncated": false,
		"created_at": "Fri Sep 21 22:51:18 +0000 2012",
		"id_str": "249279667666817024",
		"entities": {
		  "urls": [
   
		  ],
		  "hashtags": [
			{
			  "text": "freebandnames",
			  "indices": [
				20,
				34
			  ]
			}
		  ],
		  "user_mentions": [
   
		  ]
		},
		"in_reply_to_user_id_str": null,
		"contributors": null,
		"text": "The Foolish Mortals #freebandnames",
		"metadata": {
		  "iso_language_code": "en",
		  "result_type": "recent"
		},
		"retweet_count": 0,
		"in_reply_to_status_id_str": null,
		"id": 249279667666817024,
		"geo": null,
		"retweeted": false,
		"in_reply_to_user_id": null,
		"place": null,
		"user": {
		  "profile_sidebar_fill_color": "BFAC83",
		  "profile_sidebar_border_color": "615A44",
		  "profile_background_tile": true,
		  "name": "Marty Elmer",
		  "profile_image_url": "https://a0.twimg.com/profile_images/1629790393/shrinker_2000_trans_normal.png",
		  "created_at": "Mon May 04 00:05:00 +0000 2009",
		  "location": "Wisconsin, USA",
		  "follow_request_sent": null,
		  "profile_link_color": "3B2A26",
		  "is_translator": false,
		  "id_str": "37539828",
		  "entities": {
			"url": {
			  "urls": [
				{
				  "expanded_url": null,
				  "url": "https://www.omnitarian.me",
				  "indices": [
					0,
					24
				  ]
				}
			  ]
			},
			"description": {
			  "urls": [
   
			  ]
			}
		  },
		  "default_profile": false,
		  "contributors_enabled": false,
		  "favourites_count": 647,
		  "url": "https://www.omnitarian.me",
		  "profile_image_url_https": "https://si0.twimg.com/profile_images/1629790393/shrinker_2000_trans_normal.png",
		  "utc_offset": -21600,
		  "id": 37539828,
		  "profile_use_background_image": true,
		  "listed_count": 52,
		  "profile_text_color": "000000",
		  "lang": "en",
		  "followers_count": 608,
		  "protected": false,
		  "notifications": null,
		  "profile_background_image_url_https": "https://si0.twimg.com/profile_background_images/106455659/rect6056-9.png",
		  "profile_background_color": "EEE3C4",
		  "verified": false,
		  "geo_enabled": false,
		  "time_zone": "Central Time (US & Canada)",
		  "description": "Cartoonist, Illustrator, and T-Shirt connoisseur",
		  "default_profile_image": false,
		  "profile_background_image_url": "https://a0.twimg.com/profile_background_images/106455659/rect6056-9.png",
		  "statuses_count": 3575,
		  "friends_count": 249,
		  "following": null,
		  "show_all_inline_media": true,
		  "screen_name": "Omnitarian"
		},
		"in_reply_to_screen_name": null,
		"source": "<a href=\"//twitter.com/download/iphone%5C%22\" rel=\"\\\"nofollow\\\"\">Twitter for iPhone</a>",
		"in_reply_to_status_id": null
	  }
	],
	"search_metadata": {
	  "max_id": 250126199840518145,
	  "since_id": 24012619984051000,
	  "refresh_url": "?since_id=250126199840518145&q=%23freebandnames&result_type=mixed&include_entities=1",
	  "next_results": "?max_id=249279667666817023&q=%23freebandnames&count=4&include_entities=1&result_type=mixed",
	  "count": 4,
	  "completed_in": 0.035,
	  "since_id_str": "24012619984051000",
	  "query": "%23freebandnames",
	  "max_id_str": "250126199840518145"
	}
  }`

// this json block is poorly formed on purpose.
var basicJSON = `  {"age":100, "name":{"here":"B\\\"R"},
	"noop":{"what is a wren?":"a bird"},
	"happy":true,"immortal":false,
	"items":[1,2,3,{"tags":[1,2,3],"points":[[1,2],[3,4]]},4,5,6,7],
	"arr":["1",2,"3",{"hello":"world"},"4",5],
	"vals":[1,2,3,{"sadf":sdf"asdf"}],"name":{"first":"tom","last":null},
	"created":"2014-05-16T08:28:06.989Z",
	"loggy":{
		"programmers": [
    	    {
    	        "firstName": "Brett",
    	        "lastName": "McLaughlin",
    	        "email": "aaaa",
				"tag": "good"
    	    },
    	    {
    	        "firstName": "Jason",
    	        "lastName": "Hunter",
    	        "email": "bbbb",
				"tag": "bad"
    	    },
    	    {
    	        "firstName": "Elliotte",
    	        "lastName": "Harold",
    	        "email": "cccc",
				"tag":, "good"
    	    },
			{
				"firstName": 1002.3,
				"age": 101
			}
    	]
	},
	"lastly":{"end...ing":"soon","yay":"final"}
}`

var twitterLarge = func() string {
	data, err := os.ReadFile("twitter.json")
	if err != nil {
		panic(err)
	}
	return string(data)
}()

func BenchmarkGetComplexPath(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = gjson.Get(basicJSON, `loggy.programmers.#[tag="good"]#.firstName`)
		}
	})
	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = gjson.Get(twitterMedium, `statuses.#[friends_count>100]#.id`)
		}
	})
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = gjson.Get(twitterLarge, `statuses.#[friends_count>100]#.id`)
		}
	})
}

func BenchmarkGetSimplePath(b *testing.B) {
	b.Run("small", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = gjson.Get(basicJSON, `loggy.programmers.0.firstName`)
		}
	})
	b.Run("medium", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_ = gjson.Get(twitterMedium, `statuses.3.id`)
		}
	})
	b.Run("large", func(b *testing.B) {
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			x := gjson.Get(twitterLarge, `statuses.50.id`)
			if !x.Exists() {
				b.Fatal()
			}
		}
	})
}
