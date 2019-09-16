package withdrawal

import (
	"fmt"
	"log"
	"net/url"

	jpaypp "github.com/jpaypp/jpaypp-go/jpaypp"
)

type Client struct {
	B   jpaypp.Backend
	Key string
}

func getC() Client {
	return Client{jpaypp.GetBackend(jpaypp.APIBackend), jpaypp.Key}
}

func New(appId string, params *jpaypp.WithdrawalParams) (*jpaypp.Withdrawal, error) {
	return getC().New(appId, params)
}

func (c Client) New(appId string, params *jpaypp.WithdrawalParams) (*jpaypp.Withdrawal, error) {
	paramsString, _ := jpaypp.JsonEncode(params)
	withdrawal := &jpaypp.Withdrawal{}

	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/withdrawals", appId), c.Key, nil, paramsString, withdrawal)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("Balance Withdrawal error error: %v\n", err)
		}
	}
	return withdrawal, err
}

func Get(appId, withdrawalId string) (*jpaypp.Withdrawal, error) {
	return getC().Get(appId, withdrawalId)
}

func (c Client) Get(appId, withdrawalId string) (*jpaypp.Withdrawal, error) {
	withdrawal := &jpaypp.Withdrawal{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), c.Key, nil, nil, withdrawal)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("Get BalanceWithdrawal error: %v\n", err)
		}
	}
	return withdrawal, err
}

//用户的余额提现列表
func List(appId string, params *jpaypp.PagingParams) (*jpaypp.WithdrawalList, error) {
	return getC().List(appId, params)
}

func (c Client) List(appId string, params *jpaypp.PagingParams) (*jpaypp.WithdrawalList, error) {
	var body *url.Values
	body = &url.Values{}
	params.Filters.AppendTo(body)
	withdrawalList := &jpaypp.WithdrawalList{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/withdrawals", appId), c.Key, body, nil, withdrawalList)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("Get Withdrawal List error: %v\n", err)
		}
	}
	return withdrawalList, err
}

func Cancel(appId, withdrawalId string) (*jpaypp.Withdrawal, error) {
	return getC().Cancel(appId, withdrawalId)
}
func (c Client) Cancel(appId, withdrawalId string) (*jpaypp.Withdrawal, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}
	paramsString, _ := jpaypp.JsonEncode(cancelParams)

	withdrawal := &jpaypp.Withdrawal{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), c.Key, nil, paramsString, withdrawal)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("Balance Withdrawal error error: %v\n", err)
		}
	}
	return withdrawal, err
}

func Confirm(appId, withdrawalId string) (*jpaypp.Withdrawal, error) {
	return getC().Confirm(appId, withdrawalId)
}
func (c Client) Confirm(appId, withdrawalId string) (*jpaypp.Withdrawal, error) {
	confirmParams := struct {
		Status string `json:"status"`
	}{
		Status: "pending",
	}
	paramsString, _ := jpaypp.JsonEncode(confirmParams)

	withdrawal := &jpaypp.Withdrawal{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/withdrawals/%s", appId, withdrawalId), c.Key, nil, paramsString, withdrawal)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("Balance Withdrawal error: %v\n", err)
		}
	}
	return withdrawal, err
}
