package client_test

import (
	"strings"
	"testing"

	"github.com/sanchayata-jain/stripe/client"
)

func TestClientCustomer(t *testing.T) {
	c := client.Client{
		Key: "sk_test_4eC39HqLyjWDarjtT1zdp7dc",
	}
	tok := "tok_amex"
	cus, err := c.Customer(tok)
	if err != nil {
		t.Errorf("Customer() err = %s; want %v", err, nil)
	}
	if cus == nil {
		t.Fatal("Customer() = nil; want non-nil value")
	}
	if !strings.HasPrefix(cus.ID, "cus_") {
		t.Errorf("Customer() ID = %s; want prefix %q", cus.ID, "cus_")
	}
}
