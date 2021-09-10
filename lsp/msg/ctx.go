package msg

import (
	"github.com/Mrs4s/MiraiGo/message"
	"github.com/sirupsen/logrus"
)

type IMsgCtx interface {
	TextReply(text string) interface{}
	Reply(m *MSG) interface{}
	Send(m *MSG) interface{}
	NoPermissionReply() interface{}
	GetLog() *logrus.Entry
	GetTarget() Target
	GetSender() *message.Sender
}
