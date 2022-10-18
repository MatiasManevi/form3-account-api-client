package form3Client

import (
	"net/http"
	"fmt"
	"encoding/json"
	"bytes"
)

const (
	Endpoint = "organisation/accounts"
)

type Account struct {
	Attributes     *AccountAttributes `json:"attributes,omitempty"`
	ID             string             `json:"id,omitempty"`
	OrganisationID string             `json:"organisation_id,omitempty"`
	Type           string             `json:"type,omitempty"`
	Version        *int64             `json:"version,omitempty"`
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

func (c *Client) GetAccount(id string) (*Account, error) {
	url := fmt.Sprintf("%s/%s/%s", c.Host, Endpoint, id)
	// HTTP request creation
	req, err := http.NewRequest(http.MethodGet, url, nil)
	// Handle HTTP request creation errors
	if err != nil {
		return nil, err
	}

	res := Account{}
	err = c.doRequest(req, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) CreateAccount(account *Account) (*Account, error) {
	data, err := json.Marshal(account)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%s/%s", c.Host, Endpoint)	
	// HTTP request creation
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	// Handle HTTP request creation errors
	if err != nil {
		return nil, err
	}
	
	res := Account{}
	err = c.doRequest(req, &res)

	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) DeleteAccount(id string) (error) {
	url := fmt.Sprintf("%s/%s/%s", c.Host, Endpoint, id)
	// HTTP request creation
	req, err := http.NewRequest(http.MethodDelete, url, nil)
	// Handle HTTP request creation errors
	if err != nil {
		return err
	}

	err = c.doRequest(req, nil)

	if err != nil {
		return err
	}

	return nil
}