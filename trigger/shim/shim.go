package main

import (
	"github.com/apache/pulsar/pulsar-function-go/pf"

	pulsarFlogoTrigger "github.com/skothari-tibco/messaging-contrib/pulsar/trigger"
)

func main() {
	pf.Start(pulsarFlogoTrigger.Invoke)
}
