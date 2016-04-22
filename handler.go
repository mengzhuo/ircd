package ircd

import (
	"errors"
	"fmt"
	"github.com/mengzhuo/irc"
)

var noHandlerFunc = errors.New("no such handler")

func NewHandler() *Handler {
	mapper := make(map[string]HandlerFunc)
	return &Handler{mapper: mapper}
}

type Handler struct {
	mapper map[string]HandlerFunc
}

type HandlerFunc func(client *Client, msg *irc.Msg) error

func (r *Handler) Add(cmd string, f HandlerFunc) {
	r.mapper[cmd] = f
}

func (r *Handler) Handle(client *Client, msg *irc.Msg) (err error) {

	f, existed := r.mapper[string(msg.Cmd())] // compiler won't allocate string
	if !existed {
		return errors.New(fmt.Sprintf("No %s handler", msg.Cmd()))
	}
	return f(client, msg)
}

func PingHandler(c *Client, msg *irc.Msg) (err error) {

	msg.SetCmd([]byte(irc.PONG))
	if len(msg.Params()) >= 1 {
		msg.SetTrailing(msg.Params()[0])
	}
	msg.SetParams(c.Server.Name)
	_, err = c.Encoder.Encode(msg)
	msgPool.Put(msg)
	return
}
