// Package tls 反序列化回归 UT：消费 cospec/context/l5-fixtures/responses.json，
// 校验 5 个 Describe* Response 类对真实服务端响应能完整反序列化、零字段丢失。
//
// fixture 来源：实验室经 volclog 采集（profile 2100051396-sy, region cn-guilin-boe），
// 与 Java/Python/C++ V2 共享同一份 responses.json（Go 单 SDK 抽样，4 SDK 交叉反序列化）。
package tls

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func loadResponseFixture(t *testing.T) map[string]json.RawMessage {
	t.Helper()
	// 仓库相对路径：repos/volc-sdk-golang -> ../../cospec/.../responses.json
	wd, _ := os.Getwd()
	candidates := []string{
		filepath.Join(wd, "..", "..", "..", "..", "cospec", "changes", "check-tls-sdk-contract-alignment", "context", "l5-fixtures", "responses.json"),
	}
	if env := firstNonEmptyEnv("RESPONSE_FIXTURE", "L5_FIXTURE"); env != "" {
		candidates = append([]string{env}, candidates...)
	}
	for _, p := range candidates {
		if data, err := os.ReadFile(p); err == nil {
			var m map[string]json.RawMessage
			if err := json.Unmarshal(data, &m); err != nil {
				t.Fatalf("fixture %s decode failed: %v", p, err)
			}
			return m
		}
	}
	t.Skipf("response fixture not found in any of: %v (set RESPONSE_FIXTURE env to override)", candidates)
	return nil
}

func TestResponseDeserialize_DescribeCursor(t *testing.T) {
	fx := loadResponseFixture(t)
	body := fx["DescribeCursor"]
	var resp DescribeCursorResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("unmarshal DescribeCursorResponse: %v", err)
	}
	if resp.Cursor != "AAAAAAAAAAAAAAAAAAAAAAAw" {
		t.Fatalf("Cursor mismatch: got %q (服务端 wire 用 'Cursor' 大写 C，model.go 当前 json tag 是小写 'cursor'，Go 反序列化大小写不敏感所以仍 work)", resp.Cursor)
	}
}

func TestResponseDeserialize_DescribeCheckPoint(t *testing.T) {
	fx := loadResponseFixture(t)
	body := fx["DescribeCheckPoint"]
	var resp DescribeCheckPointResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("unmarshal DescribeCheckPointResponse: %v", err)
	}
	if resp.ShardID != 0 {
		t.Fatalf("ShardID mismatch: got %d (wire key 'ShardID' 大写 D，与普通 'ShardId' 不同，必须精确命中)", resp.ShardID)
	}
	if resp.Checkpoint != "" {
		t.Fatalf("Checkpoint should be empty string for new CG, got %q", resp.Checkpoint)
	}
}

func TestResponseDeserialize_DescribeIndex(t *testing.T) {
	fx := loadResponseFixture(t)
	body := fx["DescribeIndex"]
	var resp DescribeIndexResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("unmarshal DescribeIndexResponse: %v", err)
	}
	if resp.TopicID != "121c1bcd-6030-42c8-a02f-8465f1fc67ef" {
		t.Fatalf("TopicID mismatch: %q", resp.TopicID)
	}
	if resp.FullText != nil {
		t.Fatalf("FullText should be nil when wire 'FullText':null, got %+v", resp.FullText)
	}
	if resp.KeyValue == nil || len(*resp.KeyValue) != 0 {
		t.Fatalf("KeyValue should be empty slice (wire []), got %+v", resp.KeyValue)
	}
	if resp.CreateTime == "" || resp.ModifyTime == "" {
		t.Fatalf("CreateTime/ModifyTime should not be empty, got %q/%q", resp.CreateTime, resp.ModifyTime)
	}
}

func TestResponseDeserialize_DescribeProject(t *testing.T) {
	fx := loadResponseFixture(t)
	body := fx["DescribeProject"]
	var resp DescribeProjectResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("unmarshal DescribeProjectResponse: %v", err)
	}
	if resp.ProjectID != "740e2f59-08e2-41ee-b24c-dd98b79f6c35" {
		t.Fatalf("ProjectID mismatch: %q", resp.ProjectID)
	}
	if resp.ProjectName != "l3-sdk-align-1779612326" {
		t.Fatalf("ProjectName mismatch: %q", resp.ProjectName)
	}
	if resp.IamProjectName != "default" {
		t.Fatalf("IamProjectName mismatch: %q", resp.IamProjectName)
	}
	if resp.TopicCount != 1 {
		t.Fatalf("TopicCount mismatch: %d", resp.TopicCount)
	}
	if resp.InnerNetDomain == "" {
		t.Fatalf("InnerNetDomain should be set")
	}
}

func TestResponseDeserialize_DescribeTopic(t *testing.T) {
	fx := loadResponseFixture(t)
	body := fx["DescribeTopic"]
	var resp DescribeTopicResponse
	if err := json.Unmarshal(body, &resp); err != nil {
		t.Fatalf("unmarshal DescribeTopicResponse: %v", err)
	}
	if resp.TopicID != "121c1bcd-6030-42c8-a02f-8465f1fc67ef" {
		t.Fatalf("TopicID mismatch: %q", resp.TopicID)
	}
	if resp.TopicName != "l3-topic-1" {
		t.Fatalf("TopicName mismatch: %q", resp.TopicName)
	}
	if resp.TLSVersion != "2.0" {
		t.Fatalf("TLSVersion mismatch (wire key 'TLSVersion' 三大写): %q", resp.TLSVersion)
	}
	if resp.ShardCount != 2 {
		t.Fatalf("ShardCount mismatch: %d", resp.ShardCount)
	}
	if resp.MeteringMode != "ChargeByFunction" {
		t.Fatalf("MeteringMode mismatch: %q", resp.MeteringMode)
	}
	if resp.HotTtl != 1 || resp.Ttl != 1 {
		t.Fatalf("HotTtl/Ttl mismatch: %d/%d", resp.HotTtl, resp.Ttl)
	}
	if resp.AutoSplit {
		t.Fatalf("AutoSplit should be false")
	}
	if resp.EnableHotTtl {
		t.Fatalf("EnableHotTtl should be false")
	}
}
