// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestProduct_type(t *testing.T) {
	var a productResponse
	a.Objects = append(a.Objects, &Product{})
	a.Objects[0].Images = append(a.Objects[0].Images, &productImageType{})
	_ = a
}

func TestProduct_parseJson(t *testing.T) {
	var result1 productResponse
	if err := json.Unmarshal([]byte(testJsonDataProduct), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 productResponse
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	if !reflect.DeepEqual(testGoldenProduct, result1) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", testGoldenProduct, result1)
	}
}

var testGoldenProduct = productResponse{
	Request: &Request{
		PageUrl: "http://store.livrada.com/collections/all/products/before-i-go-to-sleep",
		Options: []string{},
		API:     "product",
		Fields:  "title,text,offerPrice,regularPrice,saveAmount,pageUrl,images",
		Version: 3,
	},
	Objects: []*Product{
		&Product{
			Type:         "product",
			Title:        "Before I Go To Sleep",
			Text:         "Memories define us. So what if you lost yours every time you went to sleep? Your name, your identity, your past, even the people you love -- all forgotten overnight. And the one person you trust may be telling you only half the story. Before I Go To Sleep is a disturbing psychological thriller in which an amnesiac desperately tries to uncover the truth about who she is and who she can trust.",
			OfferPrice:   "$7.99",
			RegularPrice: "$9.99",
			SaveAmount:   "$2.00",
			PageUrl:      "http://store.livrada.com/collections/all/products/before-i-go-to-sleep",
			Images: []*productImageType{
				&productImageType{
					Title:      "Before I Go to Sleep cover",
					URL:        "http://cdn.shopify.com/s/files/1/0184/6296/products/BeforeIGoToSleep_large.png?946",
					XPath:      "/HTML[@class='no-js']/BODY[@id='page-product']/DIV[@class='content-frame']/DIV[@class='content']/DIV[@class='content-shop']/DIV[@class='row']/DIV[@class='span5']/DIV[@class='product-thumbs']/UL/LI[@class='first-image']/A[@class='single_image']/IMG",
					DiffbotUri: "image|1|768070723",
				},
			},
			DiffbotUri: "product|1|937176621",
		},
	},
}

const testJsonDataProduct = `
{
  "request": {
    "pageUrl": "http://store.livrada.com/collections/all/products/before-i-go-to-sleep",
    "api": "product",
    "options": [],
    "fields": "title,text,offerPrice,regularPrice,saveAmount,pageUrl,images",
    "version": 3
  },
  "objects": [
    {
      "type": "product",
      "title": "Before I Go To Sleep",
      "text": "Memories define us. So what if you lost yours every time you went to sleep? Your name, your identity, your past, even the people you love -- all forgotten overnight. And the one person you trust may be telling you only half the story. Before I Go To Sleep is a disturbing psychological thriller in which an amnesiac desperately tries to uncover the truth about who she is and who she can trust.",
      "offerPrice": "$7.99",
      "regularPrice": "$9.99",
      "saveAmount": "$2.00",
      "pageUrl": "http://store.livrada.com/collections/all/products/before-i-go-to-sleep",
      "images": [
        {
          "title": "Before I Go to Sleep cover",
          "url": "http://cdn.shopify.com/s/files/1/0184/6296/products/BeforeIGoToSleep_large.png?946",
          "xpath": "/HTML[@class='no-js']/BODY[@id='page-product']/DIV[@class='content-frame']/DIV[@class='content']/DIV[@class='content-shop']/DIV[@class='row']/DIV[@class='span5']/DIV[@class='product-thumbs']/UL/LI[@class='first-image']/A[@class='single_image']/IMG",
          "diffbotUri": "image|1|768070723"
        }
      ],
      "diffbotUri": "product|1|937176621"
    }
  ]
}
`
