package client_test

import (
	"flag"
	"strings"
	"testing"

	"github.com/sanchayata-jain/stripe/client"
)

// key=sk_test_4eC39HqLyjWDarjtT1zdp7dc
var (
	apiKey string
)

func init() {
	flag.StringVar(&apiKey, "key", "", "Your TEST secret key for the Stripe API, if present, integration tests will be run using this key")
}

func TestClientCustomer(t *testing.T) {
	if apiKey == "" {
		t.Skip("No API key provided")
	}
	c := client.Client{
		Key: apiKey,
	}
	tok := "tok_amex"
	email := "test@testwithgo.com"
	cus, err := c.Customer(tok, email)
	if err != nil {
		t.Errorf("Customer() err = %s; want %v", err, nil)
	}
	if cus == nil {
		t.Fatal("Customer() = nil; want non-nil value")
	}
	if !strings.HasPrefix(cus.ID, "cus_") {
		t.Errorf("Customer() ID = %s; want prefix %q", cus.ID, "cus_")
	}
	if !strings.HasPrefix(cus.DefaultSource, "card_") {
		t.Errorf("Customer() ID = %s; want prefix %q", cus.DefaultSource, "card_")
	}
	if cus.Email != email {
		t.Errorf("Customer() ID = %s; want prefix %q", cus.Email, email)
	}
}
