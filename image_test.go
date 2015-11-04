// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestImage_type(t *testing.T) {
	var a imageResponse
	a.Objects = append(a.Objects, &Image{})
	_ = a
}

func TestImage_parseJson(t *testing.T) {
	var result1 imageResponse
	if err := json.Unmarshal([]byte(testJsonDataImage), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 imageResponse
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	if !reflect.DeepEqual(testGoldenImage, result1) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", testGoldenImage, result1)
	}
}

var testGoldenImage = imageResponse{
	Request: &Request{
		PageUrl: "http://www.diffbot.com/products",
		Options: []string{},
		API:     "image",
		Version: 3,
	},
	Objects: []*Image{
		&Image{
			Type:          "image",
			Url:           "http://www.diffbot.com/images/image_diffy_sample.png",
			Title:         "Diffy, climbing a mountain",
			NaturalHeight: 1158,
			NaturalWidth:  950,
			HumanLanguage: "en",
			PageUrl:       "http://www.diffbot.com/products",
			XPath:         "/HTML/BODY/DIV[@class='main']/DIV[@id='primaryImage']/IMG",
			DiffbotUri:    "image|3|-1897071612",
		},
		&Image{
			Type:          "image",
			Url:           "http://www.diffbot.com/images/image_atopmountain_sample.png",
			Title:         "Diffy atop said mountain",
			NaturalHeight: 1120,
			NaturalWidth:  920,
			HumanLanguage: "en",
			AnchorUrl:     "http://www.diffbot.com",
			PageUrl:       "http://www.diffbot.com/products",
			XPath:         "/HTML/BODY/DIV[@class='main']/DIV[@id='secondaryImage']/A/IMG",
			DiffbotUri:    "image|3|-1221792290",
		},
	},
}

const testJsonDataImage = `
{
  "request": {
    "pageUrl": "http://www.diffbot.com/products",
    "api": "image",
    "options": [],
    "fields": "",
    "version": 3
  },
  "objects": [
    {
      "title": "Diffy, climbing a mountain",
      "naturalHeight": 1158,
      "diffbotUri": "image|3|-1897071612",
      "pageUrl": "http://www.diffbot.com/products",
      "humanLanguage": "en",
      "naturalWidth": 950,
      "date": "Oct 19, 2013",
      "type": "image",
      "url": "http://www.diffbot.com/images/image_diffy_sample.png",
      "xpath": "/HTML/BODY/DIV[@class='main']/DIV[@id='primaryImage']/IMG"
    },
    {
      "title": "Diffy atop said mountain",
      "naturalHeight": 1120,
      "diffbotUri": "image|3|-1221792290",
      "pageUrl": "http://www.diffbot.com/products",
      "humanLanguage": "en",
      "naturalWidth": 920,
      "anchorUrl": "http://www.diffbot.com",
      "date": "Oct 21, 2013",
      "type": "image",
      "url": "http://www.diffbot.com/images/image_atopmountain_sample.png",
      "xpath": "/HTML/BODY/DIV[@class='main']/DIV[@id='secondaryImage']/A/IMG"
    }
  ]
}
`
