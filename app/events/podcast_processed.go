package events

import "github.com/goravel/framework/contracts/event"

type PodcastProcessed struct {
}

func (receiver *PodcastProcessed) Handle(args []event.Arg) ([]event.Arg, error) {
	return args, nil
}
