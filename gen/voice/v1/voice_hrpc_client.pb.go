// Code generated by protoc-gen-go-hrpc. DO NOT EDIT.

package voicev1

import (
	context "context"
	errors "errors"
	websocket "github.com/gorilla/websocket"
	proto "google.golang.org/protobuf/proto"
	http "net/http"
	url "net/url"
)

type VoiceServiceClient interface {
	// Endpoint to stream messages between client and server.
	//
	// - One StreamMessage stream corresponds to being in one voice channel.
	// - It's recommended that users should not be able to be in more than one voice channel,
	// but this limitation is left up to the server implementation.
	StreamMessage(context.Context, chan *StreamMessageRequest) (chan *StreamMessageResponse, error)
}

type HTTPVoiceServiceClient struct {
	Client         http.Client
	BaseURL        string
	WebsocketProto string
	WebsocketHost  string
	Header         http.Header
}

func (client *HTTPVoiceServiceClient) StreamMessage(req chan *StreamMessageRequest) (chan *StreamMessageResponse, error) {
	u := url.URL{Scheme: client.WebsocketProto, Host: client.WebsocketHost, Path: "/protocol.voice.v1.VoiceService/StreamMessage"}
	inC := req
	outC := make(chan *StreamMessageResponse)
	c, _, err := websocket.DefaultDialer.Dial(u.String(), client.Header)
	if err != nil {
		return nil, err
	}
	go func() {
		defer c.Close()
		msgs := make(chan []byte)
		go func() {
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					close(msgs)
					break
				}
				msgs <- message
			}
		}()
		for {
			select {
			case msg, ok := <-msgs:
				if !ok {
					close(inC)
					close(outC)
					return
				}
				thing := &StreamMessageResponse{}
				err := proto.Unmarshal(msg[1:], thing)
				if err != nil {
					close(inC)
					close(outC)
					return
				}
				outC <- thing
			case send, ok := <-inC:
				if !ok {
					close(outC)
					return
				}
				data, err := proto.Marshal(send)
				if err != nil {
					return
				}
				err = c.WriteMessage(websocket.BinaryMessage, data)
				if err != nil {
					return
				}
			}
		}
	}()
	return outC, nil
}

type HTTPTestVoiceServiceClient struct {
	Client interface {
		Test(*http.Request, ...int) (*http.Response, error)
	}
}

func (client *HTTPTestVoiceServiceClient) StreamMessage(req chan *StreamMessageRequest) (chan *StreamMessageResponse, error) {
	return nil, errors.New("unimplemented")
}
