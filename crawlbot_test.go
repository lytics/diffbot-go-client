package diffbot

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestCrawlbotCreate_parseJson(t *testing.T) {
	var result1 CrawlResponse
	if err := json.Unmarshal([]byte(testJsonCrawlCreate), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 CrawlResponse
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	req := &CrawlResponse{
		Response: "Successfully added urls for spidering.",
	}

	if !reflect.DeepEqual(req.Response, result1.Response) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", req.Response, result1.Response)
	}
	if !reflect.DeepEqual(req.Jobs, result1.Jobs) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", req.Jobs, result1.Jobs)
	}
}

const testJsonCrawlCreate = `
{
  "response": "Successfully added urls for spidering."
}
`

func TestCrawlbotView_parseJson(t *testing.T) {
	var result1 CrawlResponse
	if err := json.Unmarshal([]byte(testJsonCrawlView), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 CrawlResponse
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	req := &CrawlResponse{
		Jobs: []*Crawl{
			&Crawl{
				Name:                 "crawlJob",
				Type:                 "crawl",
				JobCreationTimeUTC:   1427410692,
				JobCompletionTimeUTC: 1427410798,
				JobStatus: &jobStatusType{
					Status:  9,
					Message: "Job has completed and no repeat is scheduled.",
				},
				SentJobDoneNotification:       1,
				ObjectsFound:                  177,
				UrlsHarvested:                 2152,
				PageCrawlAttempts:             367,
				PageCrawlSuccesses:            365,
				PageCrawlSuccessesThisRound:   365,
				PageProcessAttempts:           210,
				PageProcessSuccesses:          210,
				PageProcessSuccessesThisRound: 210,
				MaxRounds:                     0,
				Repeat:                        0.0,
				CrawlDelay:                    0.25,
				ObeyRobots:                    1,
				MaxToCrawl:                    100000,
				MaxToProcess:                  100000,
				OnlyProcessIfNew:              1,
				Seeds:                         "http://support.diffbot.com",
				RoundsCompleted:               0,
				RoundStartTime:                0,
				CurrentTime:                   1443822683,
				CurrentTimeUTC:                1443822683,
				ApiUrl:                        "http://api.diffbot.com/v3/analyze",
				UrlCrawlPattern:               "",
				UrlProcessPattern:             "",
				PageProcessPattern:            "",
				UrlCrawlRegEx:                 "",
				UrlProcessRegEx:               "",
				MaxHops:                       -1,
				DownloadJson:                  "http://api.diffbot.com/v3/crawl/download/sampletoken-crawlJob_data.json",
				DownloadUrls:                  "http://api.diffbot.com/v3/crawl/download/sampletoken-crawlJob_urls.csv",
				NotifyEmail:                   "support@diffbot.com",
				NotifyWebhook:                 "http://www.diffbot.com",
			},
		},
	}

	if !reflect.DeepEqual(req.Response, result1.Response) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", req.Response, result1.Response)
	}
	if !reflect.DeepEqual(req.Jobs, result1.Jobs) {
		t.Fatalf("not equal, expect = \n%q, got = \n%q", req.Jobs, result1.Jobs)
	}
}

const testJsonCrawlView = `
{
  "jobs": [
    {
      "name": "crawlJob",
      "type": "crawl",
      "jobCreationTimeUTC": 1427410692,
      "jobCompletionTimeUTC": 1427410798,
      "jobStatus": {
        "status": 9,
        "message": "Job has completed and no repeat is scheduled."
      },
      "sentJobDoneNotification": 1,
      "objectsFound": 177,
      "urlsHarvested": 2152,
      "pageCrawlAttempts": 367,
      "pageCrawlSuccesses": 365,
      "pageCrawlSuccessesThisRound": 365,
      "pageProcessAttempts": 210,
      "pageProcessSuccesses": 210,
      "pageProcessSuccessesThisRound": 210,
      "maxRounds": 0,
      "repeat": 0.0,
      "crawlDelay": 0.25,
      "obeyRobots": 1,
      "maxToCrawl": 100000,
      "maxToProcess": 100000,
      "onlyProcessIfNew": 1,
      "seeds": "http://support.diffbot.com",
      "roundsCompleted": 0,
      "roundStartTime": 0,
      "currentTime": 1443822683,
      "currentTimeUTC": 1443822683,
      "apiUrl": "http://api.diffbot.com/v3/analyze",
      "urlCrawlPattern": "",
      "urlProcessPattern": "",
      "pageProcessPattern": "",
      "urlCrawlRegEx": "",
      "urlProcessRegEx": "",
      "maxHops": -1,
      "downloadJson": "http://api.diffbot.com/v3/crawl/download/sampletoken-crawlJob_data.json",
      "downloadUrls": "http://api.diffbot.com/v3/crawl/download/sampletoken-crawlJob_urls.csv",
      "notifyEmail": "support@diffbot.com",
      "notifyWebhook": "http://www.diffbot.com"
    }
  ]
}
`

func TestCrawlbotData_parseJson(t *testing.T) {
	var result1 CrawlData
	if err := json.Unmarshal([]byte(testJsonCrawlData), &result1); err != nil {
		t.Fatal(err)
	}
	var result2 CrawlData
	if err := json.Unmarshal([]byte(result1.String()), &result2); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(result1, result2) {
		t.Fatalf("not equal, expect = %q, got = %q", result1, result2)
	}

	req := []*Classification{
		&Classification{
			Type:            "article",
			ResolvedPageUrl: "http://techcrunch.com/2012/05/31/diffbot-raises-2-million-seed-round-for-web-content-extraction-technology/",
			PageUrl:         "http://tcrn.ch/Jw7ZKw",
			HumanLanguage:   "en",
			Text:            "Diffbot , the super-geeky/awesome visual learning robot technology which aims to \"see\" the web the way that people do, is today announcing a new infusion of capital. The company has closed $2 million in funding from a number of technology veterans, including EarthLink founder Sky Dayton ; Andy Bechtolsheim , co-founder of Sun Microsystems; Joi Ito , Director of MIT Media Lab; Brad Garlinghouse , CEO of YouSendIt ( and formerly of TechCrunch parent company AOL ), Maynard Webb , Chairman of the Board at LiveOps, formerly eBay COO; Elad Gil , VP of Corporate Strategy at Twitter; Jonathan Heiliger , former VP of Technical Operations at Facebook; Redbeacon co-founder Aaron Lee ; and founder of VitalSigns Montgomery Kersten .\nMatrix Partners also participated in the round. Of the new investors, Sky Dayton will be the first to join Diffbot's board and will be taking an active role in the company, including plans to go hands-on with various Diffbot projects.\nLast August, the company publicly debuted its first APIs , which allow developers to build apps that can automatically extract meaning from web pages. For example, the Front Page API is able to analyze site homepages, and understands the difference between article text, headlines, bylines, ads, etc. The Article API can then extract clean article text, images and videos. Another example of Diffbot in action is the \"follow API,\" which can track the changes made to a website.\nToday, Diffbot has categorized the web into about 20 different page types, including homepages and article pages, which are the first two types it can now identity. Going forward, Diffbot plans train its bots to recognize all the other types of pages, including product pages, social networking profiles, recipe pages, review pages, and more.\nIts APIs have been put to use by AOL (again: disclosure, TC parent) in its news magazine AOL Editions , as well as by companies like Nuance , SocMetrics , and others. Diffbot says it's now processing 100 million API calls per month on behalf of its customers. Thousands of developers are using the APIs, the company notes, but paying customers are only in the \"tens.\" Correction: we're now told they have \"a lot more!\"\nDiffbot founder and CEO Michael Tung (aka \"Diffbot Mike\") says the new funding will be put towards new hires and expanding its resources. \"More than that, we're receiving a huge vote of confidence from veterans who have built massive companies and understand the fine points of building for scale, maintaining uptime and delivering the absolute highest standards of service.\"\nTung is a patent attorney and Stanford PhD student who left the doctoral program to pursue Diffbot, thanks to seed funding from Stanford's incubator, StartX . Diffbot was StartX's first investment. With today's funding, Diffbot total raise is $2 million and change.",
			Title:           "Diffbot Raises $2 Million Angel Round For Web Content Extraction Technology",
			Images: []*articleImageType{
				&articleImageType{
					Primary: true,
					Url:     "http://tctechcrunch2011.files.wordpress.com/2012/05/diffbot_9.png?w=300",
				},
			},
			Date: "Thu, 31 May 2012 07:00:00 GMT",
		},
	}

	if !reflect.DeepEqual(req[0], result1[0]) {
		t.Fatalf("not equal, expect = \n%+v, got = \n%+v", req[0], result1[0])
	}
}

const testJsonCrawlData = `
[
 {
    "type": "article",
    "resolvedPageUrl": "http://techcrunch.com/2012/05/31/diffbot-raises-2-million-seed-round-for-web-content-extraction-technology/",
    "pageUrl": "http://tcrn.ch/Jw7ZKw",
    "humanLanguage": "en",
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
`
