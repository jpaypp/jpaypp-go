
package jpaypp

// Data 数据对象
type Data struct {
	Object map[string]interface{} `json:"object"`
}

// ListMeta 数据列表元数据类型
type ListMeta struct {
	Object string `json:"object"`
	More   bool   `json:"has_more"`
	URL    string `json:"url"`
}


/*支付相关数据类型*/
type (
	// ChargePars 支付请求数据类型
	ChargePars struct {
		Channel     string                 `json:"channel"`
		OutOrderNo  string                 `json:"out_order_no"`
		Product     map[string]interface{} `json:"product"`
		Extra       map[string]interface{} `json:"extra"`
		ClientIp    string                 `json:"client_ip"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Description string                 `json:"description,omitempty"`
		Notify      string                 `json:"notify"`
		Callback    string                 `json:"callback"`
	}

	// ChargeListParams 列表查询请求 数据类型
	ChargeListParams struct {
		ListParams
		Created int64
	}

	// Charge 数据类型
	Charge struct {
		ID              string                 `json:"id"`
		Mode            string                 `json:"mode"`
		Paid        	bool                   `json:"paid"`
		Reversed        bool                   `json:"reversed"`
		Refunded        bool                   `json:"refunded"`
		Channel         string                 `json:"channel"`
		ChannelName     string                 `json:"channel_name"`
		PayMode         string                 `json:"pay_mode"`
		OrderNo         string                 `json:"order_no"`
		OutOrderNo   	string                 `json:"out_order_no"`
		ClientIp        string                 `json:"client_ip"`
		Amount          string                 `json:"amount"`
		Settle          string                 `json:"settle"`
		Currency        string 				   `json:"currency"`
		Subject         string                 `json:"subject"`
		Body            string                 `json:"body"`
		Extra           string                 `json:"extra"`
		TimeCreate      uint64                 `json:"time_create"`
		TimeExpire      uint64                 `json:"time_expire"`
		TimePaid  	    uint64                 `json:"time_paid"`
		TimeSettle      uint64                 `json:"time_settle"`
		TimeClose       uint64                 `json:"time_close"`
		TransactionNo   string        			`json:"transaction_no"`
		Refunds         *RefundList            `json:"refunds"`
		Description     string                 `json:"description"`
	}

	// ChargeList 列表数据类型
	ChargeList struct {
		ListMeta
		Values []*Charge `json:"data"`
	}
)

/*退款数据类型*/
type (
	// RefundParams 退款请求数据类型
	RefundParams struct {
		Amount         uint64                 `json:"amount,omitempty"`
		Description    string                 `json:"description"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Funding_source string                 `json:"funding_source,omitempty"`
	}

	// Refund 付款退款数据类型
	Refund struct {
		ID              string                 `json:"id"`
		Object          string                 `json:"object"`
		Order_no        string                 `json:"order_no"`
		Amount          uint64                 `json:"amount"`
		Succeed         bool                   `json:"succeed"`
		Status          string                 `json:"status"`
		Created         uint64                 `json:"created"`
		Time_succeed    uint64                 `json:"time_succeed"`
		Description     string                 `json:"description"`
		Failure_code    string                 `json:"failure_code"`
		Failure_msg     string                 `json:"failure_msg"`
		Metadata        map[string]interface{} `json:"metadata"`
		Charge_id       string                 `json:"charge"`
		Charge_order_no string                 `json:"charge_order_no"`
		Transaction_no  string                 `json:"transaction_no"`
		Funding_source  string                 `json:"funding_source,omitempty"`
		Extra           map[string]interface{} `json:"extra"`
	}
	// RefundList 付款查询结果列表数据类型
	RefundList struct {
		ListMeta
		Values []*Refund `json:"data"`
	}
)

/*webhooks 相关数据类型*/
type (

	// Event webhooks 反馈数据类型
	Event struct {
		Id               string `json:"id"`
		Created          int64  `json:"created"`
		Livemode         bool   `json:"livemode"`
		Type             string `json:"type"`
		Data             Data   `json:"data"`
		Object           string `json:"object"`
		Pending_webhooks int    `json:"pending_webhooks"`
		Request          string `json:"request"`
	}

	// Summary webhooks 汇总数据
	Summary struct {
		Acct_id           string `json:"acct_id,omitempty"`
		App_id            string `json:"app_id,omitempty"`
		Acct_display_name string `json:"acct_display_name"`
		App_display_name  string `json:"app_display_name"`
		Summary_from      uint64 `json:"summary_from"`
		Summary_to        uint64 `json:"summary_to"`
		Charges_amount    uint64 `json:"charges_amount"`
		Charges_count     uint64 `json:"charges_count"`
	}
)

type (
	Withdrawal struct {
		Id                  string                 `json:"id"`
		Object              string                 `json:"object"`
		App                 string                 `json:"app"`
		Amount              int64                  `json:"amount"`
		Asset_transaction   string                 `json:"asset_transaction"`
		Balance_transaction string                 `json:"balance_transaction"`
		Channel             string                 `json:"channel"`
		Created             int64                  `json:"created"`
		Description         string                 `json:"description"`
		Extra               map[string]interface{} `json:"extra"`
		Failure_msg         string                 `json:"failure_msg"`
		Fee                 int64                  `json:"fee"`
		Livemode            bool                   `json:"livemode"`
		Metadata            map[string]interface{} `json:"metadata"`
		Operation_url       string                 `json:"operation_url"`
		Order_no            string                 `json:"order_no"`
		Source              string                 `json:"source"`
		Status              string                 `json:"status"`
		Time_canceled       int64                  `json:"time_canceled"`
		Time_succeeded      int64                  `json:"time_succeeded"`
		User                string                 `json:"user"`
		User_fee            int64                  `json:"user_fee"`
		Settle_account      string                 `json:"settle_account"`
	}

	WithdrawalParams struct {
		User           string                 `json:"user,omitempty"`
		Amount         int64                  `json:"amount"`
		Channel        string                 `json:"channel,omitempty"`
		User_fee       int64                  `json:"user_fee"`
		Description    string                 `json:"description,omitempty"`
		Extra          map[string]interface{} `json:"extra,omitempty"`
		Metadata       map[string]interface{} `json:"metadata,omitempty"`
		Order_no       string                 `json:"order_no,omitempty"`
		Settle_account string                 `json:"settle_account,omitempty"`
	}

	WithdrawalList struct {
		ListMeta
		Values                   []*Withdrawal `json:"data"`
		Total_withdrawals_amount int64         `json:"total_withdrawals_amount"`
	}
)

type (
	BatchWithdrawal struct {
		Id               string                 `json:"id"`
		Object           string                 `json:"object"`
		App              string                 `json:"app"`
		Created          int64                  `json:"created"`
		Livemode         bool                   `json:"livemode"`
		Amount           int64                  `json:"amount"`
		Amount_succeeded int64                  `json:"amount_succeeded"`
		Amount_failed    int64                  `json:"amount_failed"`
		Amount_canceled  int64                  `json:"amount_canceled"`
		Count            int64                  `json:"count"`
		Count_succeeded  int64                  `json:"count_succeeded"`
		Count_failed     int64                  `json:"count_failed"`
		Count_canceled   int64                  `json:"count_canceled"`
		Fee              int64                  `json:"fee"`
		Metadata         map[string]interface{} `json:"metadata"`
		Operation_url    string                 `json:"operation_url"`
		Source           string                 `json:"source"`
		Status           string                 `json:"status"`
		TimeFinished     int64                  `json:"time_finished"`
		User_fee         int64                  `json:"user_fee"`
		Withdrawals      struct {
			ListMeta
			Values []*Withdrawal `json:"data"`
		} `json:"withdrawals"`
	}

	BatchWithdrawalParams struct {
		Withdrawals []string               `json:"withdrawals,omitempty"`
		Metadata    map[string]interface{} `json:"metadata,omitempty"`
		Status      string                 `json:"status,omitempty"`
	}

	BatchWithdrawalList struct {
		ListMeta
		Values []*BatchWithdrawal `json:"data"`
	}
)



//多级商户相关数据结构 V1.3 add
type (
	Channel struct {
		Object      string                 `json:"object"`
		Created     int64                  `json:"created"`
		Channel     string                 `json:"channel"`
		Params      map[string]interface{} `json:"params"`
		Banned      bool                   `json:"banned"`
		BannedMsg   string                 `json:"banned_msg"`
		Description string                 `json:"description"`
	}

	ChannelParams struct {
		Channel     string                 `json:"channel"`
		Params      map[string]interface{} `json:"params"`
		Banned      bool                   `json:"banned,omitempty"`
		BannedMsg   string                 `json:"banned_msg,omitempty"`
		Description string                 `json:"description,omitempty"`
	}

	ChannelUpdateParams struct {
		Params      map[string]interface{} `json:"params,omitempty"`
		Banned      bool                   `json:"banned,omitempty"`
		BannedMsg   string                 `json:"banned_msg,omitempty"`
		Description string                 `json:"description,omitempty"`
	}

	ChannelDeleteResult struct {
		Deleted bool   `json:"deleted"`
		Channel string `json:"channel"`
	}
)