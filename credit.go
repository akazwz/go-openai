package openai

import "net/http"

const (
	CreditURL = "/dashboard/billing/credit_grants"
)

func (c *Client) GetCredit() (*CreditResponse, error) {
	var credit CreditResponse
	req, err := http.NewRequest(http.MethodGet, c.apiURL+CreditURL, nil)
	err = c.doRequest(req, &credit)
	if err != nil {
		return nil, err
	}
	return &credit, nil
}
