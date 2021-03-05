package shibshib

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"

	authv1 "github.com/harmony-development/shibshib/gen/auth/v1"
	chatv1 "github.com/harmony-development/shibshib/gen/chat/v1"
	types "github.com/harmony-development/shibshib/gen/harmonytypes/v1"
	profilev1 "github.com/harmony-development/shibshib/gen/profile/v1"
)

type Client struct {
	ChatKit *chatv1.ChatServiceClient
	AuthKit *authv1.AuthServiceClient
	ProfKit *profilev1.ProfileServiceClient

	ErrorHandler func(error)

	UserID uint64

	incomingEvents <-chan *chatv1.Event
	outgoingEvents chan<- *chatv1.StreamEventsRequest

	subscribedGuilds []uint64

	events       chan *types.Message
	homeserver   string
	sessionToken string

	streaming bool

	mtx *sync.Mutex
}

var ErrEndOfStream = errors.New("end of stream")

func NewClient(homeserver, token string, userid uint64) (ret *Client, err error) {
	ret = &Client{}
	ret.events = make(chan *types.Message)

	ret.mtx = new(sync.Mutex)

	ret.homeserver = homeserver
	ret.ChatKit = chatv1.NewChatServiceClient(homeserver)
	ret.ChatKit.Header.Add("Authorization", token)
	ret.AuthKit = authv1.NewAuthServiceClient(homeserver)
	ret.AuthKit.Header.Add("Authorization", token)

	ret.sessionToken = token
	ret.UserID = userid

	ret.ErrorHandler = func(e error) {
		panic(e)
	}

	ret.StreamEvents()

	return
}

func (c *Client) StreamEvents() (err error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if c.streaming {
		return
	}

	c.outgoingEvents, c.incomingEvents, err = c.ChatKit.StreamEvents()
	if err != nil {
		err = fmt.Errorf("StreamEvents: failed to open stream: %w", err)
		return
	}

	c.streaming = true

	go func() {
		for ev := range c.incomingEvents {
			msg, ok := ev.Event.(*chatv1.Event_SentMessage)
			if !ok {
				continue
			}

			c.events <- msg.SentMessage.Message
		}

		c.mtx.Lock()
		defer c.mtx.Unlock()

		c.streaming = false
		c.ErrorHandler(ErrEndOfStream)
	}()

	return nil
}

func (c *Client) SubscribeToGuild(community uint64) {
	for _, g := range c.subscribedGuilds {
		if g == community {
			return
		}
	}
	c.outgoingEvents <- &chatv1.StreamEventsRequest{
		Request: &chatv1.StreamEventsRequest_SubscribeToGuild_{
			SubscribeToGuild: &chatv1.StreamEventsRequest_SubscribeToGuild{
				GuildId: community,
			},
		},
	}
	c.subscribedGuilds = append(c.subscribedGuilds, community)
}

func (c *Client) SubscribedGuilds() []uint64 {
	return c.subscribedGuilds
}

func (c *Client) SendMessage(msg *chatv1.SendMessageRequest) (*chatv1.SendMessageResponse, error) {
	return c.ChatKit.SendMessage(msg)
}

func (c *Client) transformHMCURL(hmc string) string {
	if !strings.HasPrefix(hmc, "hmc://") {
		return fmt.Sprintf("https://%s/_harmony/media/download/%s", c.homeserver, hmc)
	}

	trimmed := strings.TrimPrefix(hmc, "hmc://")
	split := strings.Split(trimmed, "/")
	if len(split) != 2 {
		return fmt.Sprintf("malformed URL: %s", hmc)
	}

	return fmt.Sprintf("https://%s/_harmony/media/download/%s", split[0], split[1])
}

func (c *Client) UsernameFor(m *types.Message) string {
	if m.Overrides != nil {
		return m.Overrides.Name
	}

	resp, err := c.ChatKit.GetUser(&chatv1.GetUserRequest{
		UserId: m.AuthorId,
	})
	if err != nil {
		return strconv.FormatUint(m.AuthorId, 10)
	}

	return resp.UserName
}

func (c *Client) AvatarFor(m *types.Message) string {
	if m.Overrides != nil {
		return m.Overrides.Avatar
	}

	resp, err := c.ChatKit.GetUser(&chatv1.GetUserRequest{
		UserId: m.AuthorId,
	})
	if err != nil {
		return ""
	}

	return c.transformHMCURL(resp.UserAvatar)
}
