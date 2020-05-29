package main

import (
	"github.com/apache/pulsar/pulsar-function-go/pf"

	pulsarFlogoTrigger "github.com/skothari-tibco/pulsar/function"
)

func main() {
	pf.Start(pulsarFlogoTrigger.Invoke)
}
