package tls

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestIndexRequestsSerializeEnablePhraseIndex(t *testing.T) {
	enabled := true

	createBody, err := json.Marshal(&CreateIndexRequest{
		TopicID:           "topic-id",
		EnablePhraseIndex: &enabled,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(createBody), `"EnablePhraseIndex":true`) {
		t.Fatalf("CreateIndexRequest missing EnablePhraseIndex: %s", string(createBody))
	}

	modifyBody, err := json.Marshal(&ModifyIndexRequest{
		TopicID:           "topic-id",
		EnablePhraseIndex: &enabled,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(modifyBody), `"EnablePhraseIndex":true`) {
		t.Fatalf("ModifyIndexRequest missing EnablePhraseIndex: %s", string(modifyBody))
	}

	var response DescribeIndexResponse
	if err := json.Unmarshal([]byte(`{"TopicId":"topic-id","EnablePhraseIndex":true}`), &response); err != nil {
		t.Fatal(err)
	}
	if !response.EnablePhraseIndex {
		t.Fatal("DescribeIndexResponse did not parse EnablePhraseIndex")
	}
}

func TestProcessorRequestsSerializeServiceContractFields(t *testing.T) {
	body, err := json.Marshal(&CreateProcessorRequest{
		ProjectID:        "project-id",
		ProcessorName:    "processor-name",
		Description:      "desc",
		DSLContent:       "f_set(\"k\", \"v\")",
		ProcessorType:    ProcessorTypeIngester,
		ProcessorDSLType: ProcessorDSLTypeDSL,
		ProcessorStatus:  ProcessorStatusEnabled,
		FailStrategy:     ProcessorFailStrategyKeepRaw,
		TimeoutMs:        5000,
		MaxQps:           10,
	})
	if err != nil {
		t.Fatal(err)
	}

	got := string(body)
	for _, want := range []string{
		`"ProjectId":"project-id"`,
		`"ProcessorName":"processor-name"`,
		`"ProcessorType":"ingester"`,
		`"ProcessorDSLType":"dsl"`,
		`"ProcessorStatus":"enabled"`,
		`"FailStrategy":"keep_raw"`,
		`"TimeoutMs":5000`,
		`"MaxQps":10`,
	} {
		if !strings.Contains(got, want) {
			t.Fatalf("CreateProcessorRequest missing %s in %s", want, got)
		}
	}
}
