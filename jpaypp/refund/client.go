// Package refund provides the /refunds APIs
package refund

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

func New(ch string, params *jpaypp.RefundParams) (*jpaypp.Refund, error) {
	return getC().New(ch, params)
}

func (c Client) New(ch string, params *jpaypp.RefundParams) (*jpaypp.Refund, error) {

	paramsString, errs := jpaypp.JsonEncode(params)

	if errs != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("RefundParams Marshall Errors is : %q\n", errs)
		}
		return nil, errs
	}
	if jpaypp.LogLevel > 2 {
		log.Printf("params of refund request to jpaypp is :\n %v\n ", string(paramsString))
	}
	refund := &jpaypp.Refund{}
	err := c.B.Call("POST", fmt.Sprintf("/charges/%v/refunds", ch), c.Key, nil, paramsString, refund)
	return refund, err
}

func Get(chid string, reid string) (*jpaypp.Refund, error) {
	return getC().Get(chid, reid)
}

func (c Client) Get(chid string, reid string) (*jpaypp.Refund, error) {
	var body *url.Values
	body = &url.Values{}
	refund := &jpaypp.Refund{}
	err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds/%v", chid, reid), c.Key, body, nil, refund)
	return refund, err
}

func List(chid string, params *jpaypp.RefundListParams) *Iter {
	return getC().List(chid, params)
}

func (c Client) List(chid string, params *jpaypp.RefundListParams) *Iter {
	body := &url.Values{}
	var lp *jpaypp.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{jpaypp.GetIter(lp, body, func(b url.Values) ([]interface{}, jpaypp.ListMeta, error) {
		list := &jpaypp.RefundList{}
		err := c.B.Call("GET", fmt.Sprintf("/charges/%v/refunds", chid), c.Key, &b, nil, list)

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

func (i *Iter) Refund() *jpaypp.Refund {
	return i.Current().(*jpaypp.Refund)
}

func getC() Client {
	return Client{jpaypp.GetBackend(jpaypp.APIBackend), jpaypp.Key}
}
