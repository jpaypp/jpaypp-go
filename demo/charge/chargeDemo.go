
package charge

import (
	"math/rand"
	"strconv"
	"time"
	jpaypp "github.com/jpaypp/jpaypp-go/jpaypp"
	"github.com/jpaypp/jpaypp-go/jpaypp/charge"
)

var Demo = new(ChargeDemo)

type ChargeDemo struct {
	demoAppID   string
	demoChannel string
	demoCharge  string
}

func (c *ChargeDemo) Setup(app string) {
	c.demoAppID = app
	c.demoChannel = "alipay"
	c.demoCharge = "ch_L8qn10mLmr1GS8e5OODmHaL4"
}

func (c *ChargeDemo) New() (*jpaypp.ChargePars, error) {
	//针对metadata字段，可以在每一个 charge 对象中加入订单的一些详情，如颜色、型号等属性
	metadata := make(map[string]interface{})
	metadata["color"] = "red"

	product := make(map[string]interface{})
	product["subject"] = "测试商品"
	product["body"] = "测试商品"
	product["amount"] = "1"
	product["quantity"] = "1"

	extra := make(map[string]interface{})
	extra["mode"] = "mweb"
	extra["format"] = "json"

	//这里是随便设置的随机数作为订单号，仅作示例，该方法可能产生相同订单号，商户请自行生成，不要纠结该方法。
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	orderno := r.Intn(999999999999999)

	chargePars := &jpaypp.ChargePars{
		OutOrderNo:  strconv.Itoa(orderno),
		Product: product,
		Channel:   "901",
		Extra:  extra,
		ClientIp: "127.0.0.1",
		Metadata:   metadata,
		Description: "Your Body",
		Notify:     "https://127.0.0.1",
		Callback:  "https://127.0.0.1",
	}

	//返回的第一个参数是 charge 对象，你需要将其转换成 json 给客户端，或者客户端接收后转换。
	return charge.New(chargePars)
}

// 查询 charge 对象
func (c *ChargeDemo) Get() (*jpaypp.Charge, error) {
	return charge.Get(c.demoCharge)
}

// 撤销charge，此接口仅接受线下 isv_scan、isv_wap、isv_qr 渠道的订单调用
func (c *ChargeDemo) Reverse() (*jpaypp.Charge, error) {
	return charge.Reverse(c.demoCharge)
}

// 查询 charge 对象列表
func (c *ChargeDemo) List() *charge.Iter {
	params := &jpaypp.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "3")
	//设置是不是只需要之前设置的 limit 这一个查询参数
	params.Single = true
	return charge.List(c.demoAppID, params)
}

func (c *ChargeDemo) Run() {
	c.New()
	c.Get()
	c.List()
}
