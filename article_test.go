// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestArticle_type(t *testing.T) {
	a := ArticleResponse{
		Objects: []*Article{&Article{}},
	}
	a.Objects[0].Images = append(a.Objects[0].Images, &articleImageType{})
	a.Objects[0].Videos = append(a.Objects[0].Videos, &articleVideoType{})
	_ = a
}

func TestArticle_parseJson(t *testing.T) {
	var result1 ArticleResponse
	if err := json.Unmarshal([]byte(testJsonDataArticle), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 ArticleResponse
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %v, got = %v", result1, result2)
	}

	if !reflect.DeepEqual(testArticleValue, result1) {
		t.Fatalf("not equal, expect = \n%v, got = \n%v", testArticleValue, result1)
	}
}

var (
	testArticleValue = ArticleResponse{
		Request: &Request{
			PageUrl: "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online",
			API:     "article",
			Version: 3,
		},
		Objects: []*Article{
			&Article{
				Type:             "article",
				Title:            "Diffbot's New Product API Teaches Robots to Shop Online",
				Text:             "Diffbot's human wranglers are proud today to announce the release of our newest product: an API for\u2026 products!\nThe Product API can be used for extracting clean, structured data from any e-commerce product page. It automatically makes available all the product data you'd expect: price, discount/savings amount, shipping cost, product description, any relevant product images, SKU and/or other product IDs.\nEven cooler: pair the Product API with Crawlbot, our intelligent site-spidering tool, and let Diffbot determine which pages are products, then automatically structure the entire catalog. Here's a quick demonstration of Crawlbot at work:\nWe've developed the Product API over the course of two years, building upon our core vision technology that's extracted structured data from billions of web pages, and training our machine learning systems using data from tens of thousands of unique shopping sites. We can't wait for you to try it out.\nWhat are you waiting for? Check out the Product API documentation and dive on in! If you need a token, check out our pricing and plans (including our Free plan).\nQuestions? Hit us up at support@diffbot.com.",
				Html:             "<p>Diffbot&rsquo;s human wranglers are proud today to announce the release of our newest product: an API for&hellip; products!</p>\n<p>The <a href=\"http://www.diffbot.com/products/automatic/product\">Product API</a> can be used for extracting clean, structured data from any e-commerce product page. It automatically makes available all the product data you&rsquo;d expect: price, discount/savings amount, shipping cost, product description, any relevant product images, SKU and/or other product IDs.</p>\n<p>Even cooler: pair the Product API with <a href=\"http://www.diffbot.com/products/crawlbot\">Crawlbot</a>, our intelligent site-spidering tool, and let Diffbot determine which pages are products, then automatically structure the entire catalog. Here&rsquo;s a quick demonstration of Crawlbot at work:</p>\n<figure><iframe frameborder=\"0\" src=\"http://www.youtube.com/embed/lfcri5ungRo?feature=oembed\"></iframe></figure>\n<p>We&rsquo;ve developed the Product API over the course of two years, building upon our core vision technology that&rsquo;s extracted structured data from billions of web pages, and training our machine learning systems using data from tens of thousands of unique shopping sites. We can&rsquo;t wait for you to try it out.</p>\n<p>What are you waiting for? Check out the <a href=\"http://www.diffbot.com/products/automatic/product\">Product API documentation</a> and dive on in! If you need a token, check out our <a href=\"http://www.diffbot.com/pricing\">pricing and plans</a> (including our Free plan).</p>\n<p>Questions? Hit us up at <a href=\"mailto:support@diffbot.com\">support@diffbot.com</a>.</p>",
				Date:             "Wed, 31 Jul 2013 00:00:00 GMT",
				EstimatedDate:    "Wed, 31 Jul 2013 00:00:00 GMT",
				Author:           "John Davi",
				AuthorUrl:        "http://blog.diffbot.com/author/johndavi/",
				HumanLanguage:    "en",
				SiteName:         "Diffbot",
				PublisherRegion:  "North America",
				PublisherCountry: "Diffbot HQ",
				PageUrl:          "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online",
				ResolvedUrl:      "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/",
				Tags: []*articleTag{
					&articleTag{
						Score: 0.48,
						Count: 1,
						Label: "Online and offline",
						Uri:   "http://dbpedia.org/resource/Online_and_offline",
					},
					&articleTag{
						Score: 0.45,
						Count: 1,
						Label: "Software release life cycle",
						Uri:   "http://dbpedia.org/resource/Software_release_life_cycle",
					},
					&articleTag{
						Score: 0.51,
						Count: 2,
						Label: "Structured content",
						Uri:   "http://dbpedia.org/resource/Structured_content",
					},
					&articleTag{
						Score: 0.5,
						Count: 3,
						Label: "Data",
						Uri:   "http://dbpedia.org/resource/Data",
					},
					&articleTag{
						Score: 0.78,
						Count: 5,
						Label: "Application programming interface",
						Uri:   "http://dbpedia.org/resource/Application_programming_interface",
					},
				},
				Videos: []*articleVideoType{
					&articleVideoType{
						Url:        "http://www.youtube.com/embed/lfcri5ungRo?feature=oembed",
						Primary:    true,
						DiffbotUri: "video|3|-761237582",
					},
				},
				DiffbotUri: "article|3|-820542508",
			},
		},
	}
)

const testJsonDataArticle = `
{
  "request": {
    "pageUrl": "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online",
    "api": "article",
    "version": 3
  },
  "objects": [
    {
      "date": "Wed, 31 Jul 2013 00:00:00 GMT",
      "author": "John Davi",
      "estimatedDate": "Wed, 31 Jul 2013 00:00:00 GMT",
      "publisherRegion": "North America",
      "diffbotUri": "article|3|-820542508",
      "siteName": "Diffbot",
      "videos": [
        {
          "diffbotUri": "video|3|-761237582",
          "url": "http://www.youtube.com/embed/lfcri5ungRo?feature=oembed",
          "primary": true
        }
      ],
      "type": "article",
      "title": "Diffbot's New Product API Teaches Robots to Shop Online",
      "tags": [
        {
          "score": 0.48,
          "count": 1,
          "label": "Online and offline",
          "uri": "http://dbpedia.org/resource/Online_and_offline"
        },
        {
          "score": 0.45,
          "count": 1,
          "label": "Software release life cycle",
          "uri": "http://dbpedia.org/resource/Software_release_life_cycle"
        },
        {
          "score": 0.51,
          "count": 2,
          "label": "Structured content",
          "uri": "http://dbpedia.org/resource/Structured_content"
        },
        {
          "score": 0.5,
          "count": 3,
          "label": "Data",
          "uri": "http://dbpedia.org/resource/Data"
        },
        {
          "score": 0.78,
          "count": 5,
          "label": "Application programming interface",
          "uri": "http://dbpedia.org/resource/Application_programming_interface"
        }
      ],
      "publisherCountry": "Diffbot HQ",
      "humanLanguage": "en",
      "authorUrl": "http://blog.diffbot.com/author/johndavi/",
      "pageUrl": "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online",
      "html": "<p>Diffbot&rsquo;s human wranglers are proud today to announce the release of our newest product: an API for&hellip; products!</p>\n<p>The <a href=\"http://www.diffbot.com/products/automatic/product\">Product API</a> can be used for extracting clean, structured data from any e-commerce product page. It automatically makes available all the product data you&rsquo;d expect: price, discount/savings amount, shipping cost, product description, any relevant product images, SKU and/or other product IDs.</p>\n<p>Even cooler: pair the Product API with <a href=\"http://www.diffbot.com/products/crawlbot\">Crawlbot</a>, our intelligent site-spidering tool, and let Diffbot determine which pages are products, then automatically structure the entire catalog. Here&rsquo;s a quick demonstration of Crawlbot at work:</p>\n<figure><iframe frameborder=\"0\" src=\"http://www.youtube.com/embed/lfcri5ungRo?feature=oembed\"></iframe></figure>\n<p>We&rsquo;ve developed the Product API over the course of two years, building upon our core vision technology that&rsquo;s extracted structured data from billions of web pages, and training our machine learning systems using data from tens of thousands of unique shopping sites. We can&rsquo;t wait for you to try it out.</p>\n<p>What are you waiting for? Check out the <a href=\"http://www.diffbot.com/products/automatic/product\">Product API documentation</a> and dive on in! If you need a token, check out our <a href=\"http://www.diffbot.com/pricing\">pricing and plans</a> (including our Free plan).</p>\n<p>Questions? Hit us up at <a href=\"mailto:support@diffbot.com\">support@diffbot.com</a>.</p>",
      "text": "Diffbot's human wranglers are proud today to announce the release of our newest product: an API for\u2026 products!\nThe Product API can be used for extracting clean, structured data from any e-commerce product page. It automatically makes available all the product data you'd expect: price, discount/savings amount, shipping cost, product description, any relevant product images, SKU and/or other product IDs.\nEven cooler: pair the Product API with Crawlbot, our intelligent site-spidering tool, and let Diffbot determine which pages are products, then automatically structure the entire catalog. Here's a quick demonstration of Crawlbot at work:\nWe've developed the Product API over the course of two years, building upon our core vision technology that's extracted structured data from billions of web pages, and training our machine learning systems using data from tens of thousands of unique shopping sites. We can't wait for you to try it out.\nWhat are you waiting for? Check out the Product API documentation and dive on in! If you need a token, check out our pricing and plans (including our Free plan).\nQuestions? Hit us up at support@diffbot.com.",
      "resolvedPageUrl": "http://blog.diffbot.com/diffbots-new-product-api-teaches-robots-to-shop-online/"
    }
  ]
}
`
