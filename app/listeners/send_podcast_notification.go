package listeners

import (
	"github.com/goravel/framework/contracts/event"
)

type SendPodcastNotification struct {
}

func (receiver *SendPodcastNotification) Signature() string {
	return "send_podcast_notification"
}

func (receiver *SendPodcastNotification) Queue(args ...any) event.Queue {
	return event.Queue{
		Enable:     true,
		Connection: "",
		Queue:      "",
	}
}

func (receiver *SendPodcastNotification) Handle(args ...any) error {
	name := args[0]
	print(name)
	return nil
}
