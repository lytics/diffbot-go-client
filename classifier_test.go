// Copyright 2014 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestClassification_type(t *testing.T) {
	var a classificationResponse
	_ = a
}

func TestClassification_parseJson(t *testing.T) {
	var result1 classificationResponse
	if err := json.Unmarshal([]byte(testJsonDataClassification), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 classificationResponse
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	req := &Request{
		PageUrl:         "http://tcrn.ch/Jw7ZKw",
		ResolvedPageUrl: "http://techcrunch.com/2012/05/31/diffbot-raises-2-million-seed-round-for-web-content-extraction-technology/",
		Options:         []string{},
		API:             "analyze",
		Version:         3,
	}
	if !reflect.DeepEqual(req, result1.Request) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", result1.Request)
	}

	if !reflect.DeepEqual(1, len(result1.Objects)) {
		t.Fatalf("not equal, expected len(Objects)=1, got: %v", len(result1.Objects))
	}
}

const testJsonDataClassification = `
{
  "request": {
    "pageUrl": "http://tcrn.ch/Jw7ZKw",
    "resolvedPageUrl": "http://techcrunch.com/2012/05/31/diffbot-raises-2-million-seed-round-for-web-content-extraction-technology/",
    "api": "analyze",
    "options": [],
    "fields": "",
    "version": 3
  },
  "objects": [
    {
      "type": "article",
      "resolvedPageUrl": "http://techcrunch.com/2012/05/31/diffbot-raises-2-million-seed-round-for-web-content-extraction-technology/",
      "pageUrl": "http://tcrn.ch/Jw7ZKw",
      "human_language": "en",
      "text": "Diffbot , the super-geeky/awesome visual learning robot technology which aims to \"see\" the web the way that people do, is today announcing a new infusion of capital. The company has closed $2 million in funding from a number of technology veterans, including EarthLink founder Sky Dayton ; Andy Bechtolsheim , co-founder of Sun Microsystems; Joi Ito , Director of MIT Media Lab; Brad Garlinghouse , CEO of YouSendIt ( and formerly of TechCrunch parent company AOL ), Maynard Webb , Chairman of the Board at LiveOps, formerly eBay COO; Elad Gil , VP of Corporate Strategy at Twitter; Jonathan Heiliger , former VP of Technical Operations at Facebook; Redbeacon co-founder Aaron Lee ; and founder of VitalSigns Montgomery Kersten .\nMatrix Partners also participated in the round. Of the new investors, Sky Dayton will be the first to join Diffbot's board and will be taking an active role in the company, including plans to go hands-on with various Diffbot projects.\nLast August, the company publicly debuted its first APIs , which allow developers to build apps that can automatically extract meaning from web pages. For example, the Front Page API is able to analyze site homepages, and understands the difference between article text, headlines, bylines, ads, etc. The Article API can then extract clean article text, images and videos. Another example of Diffbot in action is the \"follow API,\" which can track the changes made to a website.\nToday, Diffbot has categorized the web into about 20 different page types, including homepages and article pages, which are the first two types it can now identity. Going forward, Diffbot plans train its bots to recognize all the other types of pages, including product pages, social networking profiles, recipe pages, review pages, and more.\nIts APIs have been put to use by AOL (again: disclosure, TC parent) in its news magazine AOL Editions , as well as by companies like Nuance , SocMetrics , and others. Diffbot says it's now processing 100 million API calls per month on behalf of its customers. Thousands of developers are using the APIs, the company notes, but paying customers are only in the \"tens.\" Correction: we're now told they have \"a lot more!\"\nDiffbot founder and CEO Michael Tung (aka \"Diffbot Mike\") says the new funding will be put towards new hires and expanding its resources. \"More than that, we're receiving a huge vote of confidence from veterans who have built massive companies and understand the fine points of building for scale, maintaining uptime and delivering the absolute highest standards of service.\"\nTung is a patent attorney and Stanford PhD student who left the doctoral program to pursue Diffbot, thanks to seed funding from Stanford's incubator, StartX . Diffbot was StartX's first investment. With today's funding, Diffbot total raise is $2 million and change.",
      "title": "Diffbot Raises $2 Million Angel Round For Web Content Extraction Technology",
      "images": [
        {
        "primary": true,
        "url": "http://tctechcrunch2011.files.wordpress.com/2012/05/diffbot_9.png?w=300"
        }
      ],
      "date": "Thu, 31 May 2012 07:00:00 GMT"
    }
  ]
}
`
