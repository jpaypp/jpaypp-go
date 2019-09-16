package transfer

import (
	"log"
	"net/url"
	"strconv"

	jpaypp "github.com/jpaypp/jpaypp-go/jpaypp"
)

type Client struct {
	B   jpaypp.Backend
	Key string
}

func New(params *jpaypp.TransferParams) (*jpaypp.Transfer, error) {
	return getC().New(params)
}

func (c Client) New(params *jpaypp.TransferParams) (*jpaypp.Transfer, error) {
	paramsString, errs := jpaypp.JsonEncode(params)
	if errs != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("ChargeParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if jpaypp.LogLevel > 2 {
		log.Printf("params of redEnvelope request to jpaypp is :\n %v\n ", string(paramsString))
	}
	transfer := &jpaypp.Transfer{}
	err := c.B.Call("POST", "/transfers", c.Key, nil, paramsString, transfer)
	return transfer, err
}

func Update(id string) (*jpaypp.Transfer, error) {
	return getC().Update(id)
}

func (c Client) Update(id string) (*jpaypp.Transfer, error) {
	cancelParams := struct {
		Status string `json:"status"`
	}{
		Status: "canceled",
	}

	paramsString, _ := jpaypp.JsonEncode(cancelParams)
	transfer := &jpaypp.Transfer{}
	err := c.B.Call("PUT", "/transfers/"+id, c.Key, nil, paramsString, transfer)
	return transfer, err
}

// Get returns the details of a redenvelope.
func Get(id string) (*jpaypp.Transfer, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*jpaypp.Transfer, error) {
	var body *url.Values
	body = &url.Values{}
	transfer := &jpaypp.Transfer{}
	err := c.B.Call("GET", "/transfers/"+id, c.Key, body, nil, transfer)
	return transfer, err
}

// List returns a list of transfer.
func List(params *jpaypp.TransferListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *jpaypp.TransferListParams) *Iter {
	type transferList struct {
		jpaypp.ListMeta
		Values []*jpaypp.Transfer `json:"data"`
	}

	var body *url.Values
	var lp *jpaypp.ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}
		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{jpaypp.GetIter(lp, body, func(b url.Values) ([]interface{}, jpaypp.ListMeta, error) {
		list := &transferList{}
		err := c.B.Call("GET", "/transfers", c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

type Iter struct {
	*jpaypp.Iter
}

func (i *Iter) Transfer() *jpaypp.Transfer {
	return i.Current().(*jpaypp.Transfer)
}

func getC() Client {
	return Client{jpaypp.GetBackend(jpaypp.APIBackend), jpaypp.Key}
}
