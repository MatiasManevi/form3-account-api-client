package form3Client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	// Points to form3 account API accounts endpoint
	Endpoint = "organisation/accounts"

	// Version of created accounts used on deletion
	AccountVersion = "0"
)

type NewAccountRequest struct {
	Data Account `json:"data"`
}

type AccountResponse struct {
	Data  Account     `json:"data"`
	Links interface{} `json:"links"`
}

type Account struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
}

type AccountAttributes struct {
	AccountClassification   *string  `json:"account_classification,omitempty"`
	AccountMatchingOptOut   *bool    `json:"account_matching_opt_out,omitempty"`
	AccountNumber           string   `json:"account_number,omitempty"`
	AlternativeNames        []string `json:"alternative_names,omitempty"`
	BankID                  string   `json:"bank_id,omitempty"`
	BankIDCode              string   `json:"bank_id_code,omitempty"`
	BaseCurrency            string   `json:"base_currency,omitempty"`
	Bic                     string   `json:"bic,omitempty"`
	Country                 *string  `json:"country,omitempty"`
	Iban                    string   `json:"iban,omitempty"`
	JointAccount            *bool    `json:"joint_account,omitempty"`
	Name                    []string `json:"name,omitempty"`
	SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	Status                  *string  `json:"status,omitempty"`
	Switched                *bool    `json:"switched,omitempty"`
}

func (c *Client) GetAccount(id string) (*AccountResponse, error) {
	url := fmt.Sprintf("%s/%s/%s", c.Host, Endpoint, id)
	req, err := buildRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	res := AccountResponse{}
	err = c.doRequest(req, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateAccount(account Account) (*AccountResponse, error) {
	newAccountReq := NewAccountRequest{
		Data: account,
	}

	// Converting go struct into []byte
	data, err := json.Marshal(newAccountReq)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.Host, Endpoint)
	req, err := buildRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, err
	}

	res := AccountResponse{}

	err = c.doRequest(req, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteAccount(id string) error {
	url := fmt.Sprintf("%s/%s/%s?version=%s", c.Host, Endpoint, id, AccountVersion)
	req, err := buildRequest(http.MethodDelete, url, nil)
	if err != nil {
		return err
	}

	err = c.doRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}
