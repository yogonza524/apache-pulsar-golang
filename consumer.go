package main

import (
	"yogonza524/pulsar-client/src/model"
)

func main() {
	p := model.Pulsar{}

	p.Connect()
	p.Consume()
}