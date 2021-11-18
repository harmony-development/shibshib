// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package syncv1

import (
	bytes "bytes"
	context "context"
	proto "google.golang.org/protobuf/proto"
	ioutil "io/ioutil"
	http "net/http"
	httptest "net/http/httptest"
)

type PostboxServiceClient interface {
	// Endpoint to pull events.
	Pull(context.Context, *PullRequest) (*PullResponse, error)
	// Endpoint to push events.
	Push(context.Context, *PushRequest) (*PushResponse, error)
	// Endpoint to notify a server of a server ID change. It is called by the server
	// that had it's server ID changed for all servers it has federated with.
	NotifyNewId(context.Context, *NotifyNewIdRequest) (*NotifyNewIdResponse, error)
}

type HTTPPostboxServiceClient struct {
	Client         http.Client
	BaseURL        string
	WebsocketProto string
	WebsocketHost  string
	Header         http.Header
}

func (client *HTTPPostboxServiceClient) Pull(req *PullRequest) (*PullResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	hreq, err := http.NewRequest("POST", client.BaseURL+"/protocol.sync.v1.PostboxService.Pull/", reader)
	if err != nil {
		return nil, err
	}
	for k, v := range client.Header {
		hreq.Header[k] = v
	}
	hreq.Header.Add("content-typ", "application/hrpc")
	resp, err := client.Client.Do(hreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &PullResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPPostboxServiceClient) Push(req *PushRequest) (*PushResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	hreq, err := http.NewRequest("POST", client.BaseURL+"/protocol.sync.v1.PostboxService.Push/", reader)
	if err != nil {
		return nil, err
	}
	for k, v := range client.Header {
		hreq.Header[k] = v
	}
	hreq.Header.Add("content-typ", "application/hrpc")
	resp, err := client.Client.Do(hreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &PushResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPPostboxServiceClient) NotifyNewId(req *NotifyNewIdRequest) (*NotifyNewIdResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	hreq, err := http.NewRequest("POST", client.BaseURL+"/protocol.sync.v1.PostboxService.NotifyNewId/", reader)
	if err != nil {
		return nil, err
	}
	for k, v := range client.Header {
		hreq.Header[k] = v
	}
	hreq.Header.Add("content-typ", "application/hrpc")
	resp, err := client.Client.Do(hreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &NotifyNewIdResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}

type HTTPTestPostboxServiceClient struct {
	Client interface {
		Test(*http.Request, ...int) (*http.Response, error)
	}
}

func (client *HTTPTestPostboxServiceClient) Pull(req *PullRequest) (*PullResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.sync.v1.PostboxService.Pull/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &PullResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestPostboxServiceClient) Push(req *PushRequest) (*PushResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.sync.v1.PostboxService.Push/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &PushResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
func (client *HTTPTestPostboxServiceClient) NotifyNewId(req *NotifyNewIdRequest) (*NotifyNewIdResponse, error) {
	data, marshalErr := proto.Marshal(req)
	if marshalErr != nil {
		return nil, marshalErr
	}
	reader := bytes.NewReader(data)
	testreq := httptest.NewRequest("POST", "/protocol.sync.v1.PostboxService.NotifyNewId/", reader)
	resp, err := client.Client.Test(testreq)
	if err != nil {
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	ret := &NotifyNewIdResponse{}
	unmarshalErr := proto.Unmarshal(body, ret)
	if unmarshalErr != nil {
		return nil, unmarshalErr
	}
	return ret, nil
}
