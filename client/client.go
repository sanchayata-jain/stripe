package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const (
	Version         = "2018-09-24"
	DefaultCurrency = "gbp"
)

type Customer struct {
	ID            string `json:"id"`
	DefaultSource string `json:"default_source"`
	Email         string `json:"email"`
}

type Client struct {
	Key string
}

type Charge struct {
	ID             string `json:"id"`
	Amount         int    `json:"amount"`
	FailureCode    string `json:"failure_code"`
	FailureMessage string `json:"failure_message"`
	Paid           bool   `json:"paid"`
	Status         string `json:"status"`
}

func (c *Client) Charge(customerID string, amount int) (*Charge, error) {
	endpoint := "https://api.stripe.com/v1/charges"
	v := url.Values{}
	v.Set("customer", customerID)
	v.Set("amount", strconv.Itoa(amount))
	v.Set("currency", DefaultCurrency)
	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Stripe-Version", Version)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.Key, "")
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	var chg Charge
	err = json.Unmarshal(body, &chg)
	if err != nil {
		return nil, err
	}

	return &chg, nil
}

func (c *Client) Customer(token string, email string) (*Customer, error) {
	endpoint := "https://api.stripe.com/v1/customers"
	v := url.Values{}
	v.Set("source", token)
	v.Set("email", email)
	req, err := http.NewRequest(http.MethodPost, endpoint, strings.NewReader(v.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Stripe-Version", Version)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.SetBasicAuth(c.Key, "")
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	// fmt.Println(string(body))
	var cus Customer
	err = json.Unmarshal(body, &cus)
	if err != nil {
		return nil, err
	}

	return &cus, nil
}
