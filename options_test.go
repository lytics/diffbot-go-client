// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"net/url"
	"testing"
	"time"
)

func TestOptions(t *testing.T) {
	for i, v := range testOptionsList {
		if s := v.opt.MethodParamString(v.method); s.Encode() != v.str {
			t.Fatalf("%d: expect = %q, got = %q", i, v.str, s.Encode())
		}
	}
}

var testOptionsList = []struct {
	method string
	opt    Options
	str    string
}{
	// empty
	{
		method: "",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "",
	},

	// case "article", "image", "product":
	{
		method: "article",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "callback=abc&discussion=false&fields=" + url.QueryEscape("meta,querystring,images(*)") + "&timeout=5000",
	},
	{
		method: "image",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "callback=abc&discussion=false&fields=" + url.QueryEscape("meta,querystring,images(*)") + "&timeout=5000",
	},
	{
		method: "product",
		opt: Options{
			Fields:   "meta,querystring,images(*)",
			Timeout:  time.Second * 5,
			Callback: "abc",
		},
		str: "callback=abc&discussion=false&fields=" + url.QueryEscape("meta,querystring,images(*)") + "&timeout=5000",
	},

	// case "frontpage":
	{
		method: "frontpage",
		opt: Options{
			Fields:       "meta,querystring,images(*)",
			Timeout:      time.Second * 5,
			Callback:     "abc",
			FrontpageAll: "*",
		},
		str: "all=%2A&timeout=5000",
	},

	// case "analyze":
	{
		method: "analyze",
		opt: Options{
			Fields:          "meta,querystring,images(*)",
			Timeout:         time.Second * 5,
			Callback:        "abc",
			FrontpageAll:    "*",
			ClassifierMode:  "frontpage",
			ClassifierStats: "abc",
		},
		str: "discussion=false&fields=" + url.QueryEscape("meta,querystring,images(*)") + "&mode=frontpage&stats=abc",
	},

	// case "bulk":
	// case "crawl":
	// case "batch":
	// default: // Custom APIs
}
