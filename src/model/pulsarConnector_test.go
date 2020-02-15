package model

import (
	"testing"

	"github.com/apache/pulsar/pulsar-client-go/pulsar"
)

func TestPulsar_Connect(t *testing.T) {
	type fields struct {
		Status          int
		Client          pulsar.Client
		TwitterProducer pulsar.Producer
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pulsar{
				Status:          tt.fields.Status,
				Client:          tt.fields.Client,
				TwitterProducer: tt.fields.TwitterProducer,
			}
			p.Connect()
		})
	}
}
