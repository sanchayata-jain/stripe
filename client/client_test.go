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

func TestClientCharge(t *testing.T) {
	if apiKey == "" {
		t.Skip("No API key provided")
	}

	c := client.Client{
		Key: apiKey,
	}
	// Create a customer for the test
	tok := "tok_amex"
	email := "test@testwithgo.com"
	cus, err := c.Customer(tok, email)
	if err != nil {
		t.Fatalf("Customer() err = %s; want %v", err, nil)
	}
	amount := 1234
	charge, err := c.Charge(cus.ID, amount)
	if err != nil {
		t.Errorf("Charge() err = %s; want %v", err, nil)
	}
	if charge == nil {
		t.Fatal("Charge() = nil; want non-nil value")
	}

	if charge.Amount != amount {
		t.Errorf("Charge() Amount = %d; want %d", charge.Amount, amount)
	}
}
