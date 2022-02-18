package start

import "adapter/src/presentation"

func StartApplication() <-chan error {
	return presentation.SubscribeConsumers()
}
