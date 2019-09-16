package channel

import (
	"fmt"
	"log"
	jpaypp "github.com/jpaypp/jpaypp-go/jpaypp"
)

type Client struct {
	B   jpaypp.Backend
	Key string
}

func getC() Client {
	return Client{jpaypp.GetBackend(jpaypp.APIBackend), jpaypp.Key}
}

func New(appId, subAppId string, params *jpaypp.ChannelParams) (*jpaypp.Channel, error) {
	return getC().New(appId, subAppId, params)
}

func (c Client) New(appId, subAppId string, params *jpaypp.ChannelParams) (*jpaypp.Channel, error) {
	paramsString, errs := jpaypp.JsonEncode(params)
	if errs != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("ChannelParams Marshall Errors is : %q/n", errs)
		}
		return nil, errs
	}
	if jpaypp.LogLevel > 2 {
		log.Printf("params of create user is :\n %v\n ", string(paramsString))
	}

	channel := &jpaypp.Channel{}
	err := c.B.Call("POST", fmt.Sprintf("/apps/%s/sub_apps/%s/channels", appId, subAppId), c.Key, nil, paramsString, channel)
	return channel, err
}

func Get(appId, subAppId, channel string) (*jpaypp.Channel, error) {
	return getC().Get(appId, subAppId, channel)
}

func (c Client) Get(appId, subAppId, channelName string) (*jpaypp.Channel, error) {
	channel := &jpaypp.Channel{}

	err := c.B.Call("GET", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, nil, channel)
	return channel, err
}

func Delete(appId, subAppId, channelName string) (*jpaypp.ChannelDeleteResult, error) {
	return getC().Delete(appId, subAppId, channelName)
}

func (c Client) Delete(appId, subAppId, channelName string) (*jpaypp.ChannelDeleteResult, error) {
	result := &jpaypp.ChannelDeleteResult{}

	err := c.B.Call("DELETE", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, nil, result)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return result, err
}

func Update(appId, subAppId, channelName string, params jpaypp.ChannelUpdateParams) (*jpaypp.Channel, error) {
	return getC().Update(appId, subAppId, channelName, params)
}

func (c Client) Update(appId, subAppId, channelName string, params jpaypp.ChannelUpdateParams) (*jpaypp.Channel, error) {
	paramsString, _ := jpaypp.JsonEncode(params)
	if jpaypp.LogLevel > 2 {
		log.Printf("params of update Channel  to jpaypp is :\n %v\n ", string(paramsString))
	}

	channel := &jpaypp.Channel{}

	err := c.B.Call("PUT", fmt.Sprintf("/apps/%s/sub_apps/%s/channels/%s", appId, subAppId, channelName), c.Key, nil, paramsString, channel)
	if err != nil {
		if jpaypp.LogLevel > 0 {
			log.Printf("%v\n", err)
		}
		return nil, err
	}
	return channel, err
}
