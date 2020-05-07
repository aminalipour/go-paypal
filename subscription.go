package paypal

import (
	"fmt"
	"net/http"
	"time"
)

type (
	// SubscriptionDetailResp struct
	SubscriptionDetailResp struct {
		ID               string         `json:"id,omitempty"`
		PlanID           string         `json:"plan_id,omitempty"`
		StartTime        time.Time      `json:"start_time,omitempty"`
		Quantity         string         `json:"quantity,omitempty"`
		ShippingAmount   ShippingAmount `json:"shipping_amount,omitempty"`
		Subscriber       Subscriber     `json:"subscriber,omitempty"`
		BillingInfo      BillingInfo    `json:"billing_info,omitempty"`
		CreateTime       time.Time      `json:"create_time,omitempty"`
		UpdateTime       time.Time      `json:"update_time,omitempty"`
		Links            []Link         `json:"links,omitempty"`
		Status           string         `json:"status,omitempty"`
		StatusUpdateTime time.Time      `json:"status_update_time,omitempty"`
	}
	SubscriptionTransaction struct {
		Status              string              `json:"status,"`
		ID                  string              `json:"id,omitempty"`
		AmountWithBreakDown AmountWithBreakDown `json:"amount_with_break_down,omitempty"`
		PayerName           PayerName           `json:"payer_name,omitempty"`
		PayerEmail          string              `json:"payer_email,omitempty"`
		Time                time.Time           `json:"time,omitempty"`
	}
	AmountWithBreakDown struct {
		GrossAmount    *Money `json:"gross_amount,omitempty"`
		FeeAmount      *Money `json:"fee_amount,omitempty"`
		ShippingAmount *Money `json:"shipping_amount,omitempty"`
		TaxAmount      *Money `json:"tax_amount,omitempty"`
		NetAmount      *Money `json:"net_amount,omitempty"`
	}
	PayerName struct {
		Prefix            string `json:"prefix,omitempty"`
		GivenName         string `json:"given_name,omitempty"`
		Surname           string `json:"surname,omitempty"`
		MiddleName        string `json:"middle_name,omitempty"`
		Suffix            string `json:"suffix,omitempty"`
		AlternateFullName string `json:"alternate_full_name,omitempty"`
		FullName          string `json:"full_name,omitempty"`
	}
	SubscriptionTransactions struct {
		Transactions []SubscriptionTransaction `json:"transactions,omitempty"`
		Links        []Link        `json:"links"`
		TotalItems   int           `json:"total_items"`
		TotalPages   int           `json:"total_pages"`
	}
)

// GetSubscriptionDetails shows details for a subscription, by ID.
// Endpoint: GET /v1/billing/subscriptions/
func (c *Client) GetSubscriptionDetails(subscriptionID string) (*SubscriptionDetailResp, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/billing/subscriptions/%s", c.APIBase, subscriptionID), nil)
	response := &SubscriptionDetailResp{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}

// Create Subscription a billing plan, by PlanID.
// Endpoint: POST /v1/billing/subscriptions/
func (c *Client) CreateSubscription(subscription Subscription) (*SubscriptionDetailResp, error) {
	req, err := c.NewRequest("POST", fmt.Sprintf("%s%s", c.APIBase, "/v1/billing/subscriptions"), subscription)
	response := &SubscriptionDetailResp{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}



// GetSubscriptionTransactions shows transactions for a subscription, by ID.
// Endpoint: GET /v1/billing/subscriptions/
func (c *Client) GetSubscriptionTransactions(subscriptionID string,filter string) (*SubscriptionTransactions, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/v1/billing/subscriptions/%s/transactions%s", c.APIBase, subscriptionID,filter), nil)
	response := &SubscriptionTransactions{}
	if err != nil {
		return response, err
	}
	err = c.SendWithAuth(req, response)
	return response, err
}