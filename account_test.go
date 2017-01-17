package diffbot

import (
	"encoding/json"
	"testing"
)

func TestAccount(t *testing.T) {
	account := &Account{}
	err := json.Unmarshal([]byte(testAccountResponse), account)
	if err != nil {
		t.Fatalf("could not unmarshal response into account: %v", err)
	}
	if account.Name == "" {
		t.Fatalf("account name should not be empty. account=%+v", account)
	}
	if account.Plan == "" {
		t.Fatalf("account plan should not be empty. account=%+v", account)
	}
	if account.Status == "" {
		t.Fatalf("account status should not be empty. account=%+v", account)
	}
	if len(account.ApiCalls) == 0 {
		t.Fatalf("account apiCalls should not be empty. account=%+v", account)
	}
	if account.Email == "" {
		t.Fatalf("account email should not be empty. account=%+v", account)
	}
}

var testAccountResponse = `
{
  "planCalls": 250000,
  "lastBilling": "2017-01-05",
  "name": "John Doe",
  "plan": "startup",
  "apiCalls": [
    {
      "date": "2016-12-19",
      "calls": 100,
      "proxyCalls": 10,
      "giCalls": 0
    },
    {
      "date": "2016-12-18",
      "calls": 200,
      "proxyCalls": 0,
      "giCalls": 0
    }
  ],
  "email": "email@lol.com",
  "token": "<fake_token>",
  "status": "active"
}
`
