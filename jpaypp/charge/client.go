package charge

import (
	"log"
	"net/url"
	"strconv"
	"time"

	jpaypp "github.com/jpaypp/jpaypp-go/jpaypp"
)

type Client struct {
	B   jpaypp.Backend
	Key string
}

func getC() Client {
	return Client{jpaypp.GetBackend(jpaypp.APIBackend), jpaypp.Key}
}

// 发送 charge 请求
func New(pars *jpaypp.ChargePars) (*jpaypp.Charge, error) {
	return getC().New(pars)
}

func (c Client) New(params *jpaypp.ChargeParams) (*jpaypp.Charge, error) {
	start := time.Now()
	paramsString, errs := jpaypp.JsonEncode(params)
	if errs != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("ChargeParams Marshall Errors is : %q\n", errs)
		}
	}
	if jpaypp.LogLevel > 2 {
		log.Printf("params of charge request to jpaypp is :\n %v\n ", string(paramsString))
	}

	charge := &jpaypp.Charge{}
	errch := c.B.Call("POST", "/charges", c.Key, nil, paramsString, charge)
	if errch != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("%v\n", errch)
		}
		return nil, errch
	}
	if jpaypp.LogLevel > 2 {
		log.Println("Charge completed in ", time.Since(start))
	}
	return charge, errch

}

// 撤销charge，此接口仅接受线下 isv_scan、isv_wap、isv_qr 渠道的订单调用
func Reverse(id string) (*jpaypp.Charge, error) {
	return getC().Reverse(id)
}

func (c Client) Reverse(id string) (*jpaypp.Charge, error) {
	var body *url.Values
	body = &url.Values{}
	charge := &jpaypp.Charge{}
	err := c.B.Call("POST", "/charges/"+id+"/reverse", c.Key, body, nil, charge)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("Reverse Charge error: %v\n", err)
		}
	}
	return charge, err
}

//查询指定 charge 对象
func Get(id string) (*jpaypp.Charge, error) {
	return getC().Get(id)
}

func (c Client) Get(id string) (*jpaypp.Charge, error) {
	var body *url.Values
	body = &url.Values{}
	charge := &jpaypp.Charge{}
	err := c.B.Call("GET", "/charges/"+id, c.Key, body, nil, charge)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("Get Charge error: %v\n", err)
		}
	}
	return charge, err
}

// 查询 charge 列表
func List(appId string, params *jpaypp.ChargeListParams) *Iter {
	return getC().List(appId, params)
}

func (c Client) List(appId string, params *jpaypp.ChargeListParams) *Iter {
	type chargeList struct {
		jpaypp.ListMeta
		Values []*jpaypp.Charge `json:"data"`
	}

	var body *url.Values
	var lp *jpaypp.ListParams

	if params == nil {
		params = &jpaypp.ChargeListParams{}
	}
	params.Filters.AddFilter("app[id]", "", appId)
	body = &url.Values{}
	if params.Created > 0 {
		body.Add("created", strconv.FormatInt(params.Created, 10))
	}
	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{jpaypp.GetIter(lp, body, func(b url.Values) ([]interface{}, jpaypp.ListMeta, error) {
		list := &chargeList{}
		err := c.B.Call("GET", "/charges", c.Key, &b, nil, list)

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

func (i *Iter) Charge() *jpaypp.Charge {
	return i.Current().(*jpaypp.Charge)
}
