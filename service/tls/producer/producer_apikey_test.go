package producer

import (
	"testing"

	. "github.com/volcengine/volc-sdk-golang/service/tls"
)

func TestNewProducerPropagatesAPIKeyToTLSClient(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://example.com"
	cfg.Region = "cn-beijing"
	cfg.APIKey = "api-key"

	p := newProducer(cfg)
	client, ok := p.cli.(*LsClient)
	if !ok {
		t.Fatalf("expected *tls.LsClient, got %T", p.cli)
	}
	if client.APIKey != "api-key" {
		t.Fatalf("expected producer client APIKey to be propagated")
	}
}

func TestProducerSetAPIKeyUpdatesTLSClient(t *testing.T) {
	cfg := GetDefaultProducerConfig()
	cfg.Endpoint = "http://example.com"
	cfg.Region = "cn-beijing"
	cfg.APIKey = "old-api-key"

	var api Producer = newProducer(cfg)
	api.SetAPIKey("new-api-key")
	p := api.(*producer)
	client, ok := p.cli.(*LsClient)
	if !ok {
		t.Fatalf("expected *tls.LsClient, got %T", p.cli)
	}
	if client.APIKey != "new-api-key" {
		t.Fatalf("expected producer client APIKey to be updated")
	}
}
