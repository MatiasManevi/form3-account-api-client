package form3Client

import (
	"testing"
	"encoding/json"
	"io/ioutil"
	"github.com/stretchr/testify/assert"
)

func fixture(path string) string {
    data, err := ioutil.ReadFile("testdata/" + path)
    if err != nil {
        panic(err)
    }
    return string(data)
}

func TestCreateAccount(t *testing.T) {
	file := fixture("account_1.json")

	NewAccount := Account{}
	err := json.Unmarshal([]byte(file), &NewAccount)
	if err != nil {
        panic(err)
    }
	
	response, err := NewClient(nil).CreateAccount(NewAccount)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, response.Data.ID, "expecting non-nil result")
}

func TestGetAccount(t *testing.T) {
	file := fixture("account_1.json")

	acc := Account{}
	err := json.Unmarshal([]byte(file), &acc)
	if err != nil {
        panic(err)
    }
	
	response, err := NewClient(nil).GetAccount(acc.ID)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, response.Data.ID, "expecting non-nil result")
}

func TestDeleteAccount(t *testing.T) {
	file := fixture("account_1.json")

	acc := Account{}
	err := json.Unmarshal([]byte(file), &acc)
	if err != nil {
        panic(err)
    }
	
	err = NewClient(nil).DeleteAccount(acc.ID)

	assert.Nil(t, err, "expecting nil error")
	assert.NotNil(t, acc, "expecting non-nil result")
}

func TestAccountDuplicatedError(t *testing.T) {
	file := fixture("account_2.json")
	client := NewClient(nil);

	NewAccount := Account{}
	err := json.Unmarshal([]byte(file), &NewAccount)
	if err != nil {
        panic(err)
    }
	
	client.CreateAccount(NewAccount)
	response, err := client.CreateAccount(NewAccount)

	assert.Nil(t, response, "expecting nil result")
	assert.NotNil(t, err, "expecting non-nil error")
	assert.Equal(t, err.Error(), "Account cannot be created as it violates a duplicate constraint", "expecting duplicate constraint violation message")
}

func TestInvalidDataOnCreateAccount(t *testing.T) {
	file := fixture("invalid_account.json")

	NewAccount := Account{}
	err := json.Unmarshal([]byte(file), &NewAccount)
	if err != nil {
        panic(err)
    }
	
	response, err := NewClient(nil).CreateAccount(NewAccount)

	assert.NotNil(t, err, "expecting non-nil error")
	assert.Nil(t, response, "expecting nil result")
	assert.Equal(t, err.Error(), "validation failure list:\nvalidation failure list:\nattributes in body is required\ntype in body is required", "expecting validation failure due to incomplete request payload")
}

func TestGetAccountNotFound(t *testing.T) {
	response, err := NewClient(nil).GetAccount("1d209d7f-d07a-4542-947f-5885fddddaa2")

	assert.NotNil(t, err, "expecting non-nil error")
	assert.Nil(t, response, "expecting nil result")
	assert.Equal(t, err.Error(), "record 1d209d7f-d07a-4542-947f-5885fddddaa2 does not exist", "expecting not found account error")
}

func TestGetAccountInvalidUUID(t *testing.T) {
	response, err := NewClient(nil).GetAccount("non-uuid-string")

	assert.NotNil(t, err, "expecting non-nil error")
	assert.Nil(t, response, "expecting nil result")
	assert.Equal(t, err.Error(), "id is not a valid uuid", "expecting invalid uuid id when trying to get an account")
}

func TestUnkownErrorOnDeleteUnexistentAccount(t *testing.T) {
	err := NewClient(nil).DeleteAccount("1d209d7f-d07a-4542-947f-5885fddddaa2")

	assert.NotNil(t, err, "expecting non-nil error")
	assert.Equal(t, err.Error(), "unknown error, status code: 404", "expecting unkown error when trying to delete an account that does not exists")
}