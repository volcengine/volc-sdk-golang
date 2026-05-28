package tls

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

// Signer dump.
//
// 读取固定 fixture，调用 base.GetSignRequest 复现签名流程（绕过 now()），
// 将关键签名结果写入 l4-snapshots/sign-go.txt 用于跨 SDK 比对。

type signFixture struct {
	AK          string            `json:"ak"`
	SK          string            `json:"sk"`
	Region      string            `json:"region"`
	Service     string            `json:"service"`
	Method      string            `json:"method"`
	Host        string            `json:"host"`
	Path        string            `json:"path"`
	Query       map[string]string `json:"query"`
	Body        string            `json:"body"`
	ContentType string            `json:"content_type"`
	XDate       string            `json:"x_date"`
}

func TestSignDump(t *testing.T) {
	fixturePath := firstNonEmptyEnv("SIGN_FIXTURE_PATH", "L4_FIXTURE_PATH")
	if fixturePath == "" {
		var ok bool
		fixturePath, ok = findSignFixturePath()
		if !ok {
			t.Skip("sign fixture not found; set SIGN_FIXTURE_PATH to run dump")
		}
	}

	raw, err := os.ReadFile(fixturePath)
	if err != nil {
		if os.IsNotExist(err) {
			t.Skipf("sign fixture not found: %s", fixturePath)
		}
		t.Fatalf("read fixture: %v", err)
	}
	var fx signFixture
	if err := json.Unmarshal(raw, &fx); err != nil {
		t.Fatalf("parse fixture: %v", err)
	}

	fixedTime, err := time.Parse("20060102T150405Z", fx.XDate)
	if err != nil {
		t.Fatalf("parse x_date: %v", err)
	}

	header := http.Header{}
	header.Set("Content-Type", fx.ContentType)

	queryList := url.Values{}
	for k, v := range fx.Query {
		queryList.Set(k, v)
	}

	rp := base.RequestParam{
		IsSignUrl: false,
		Body:      []byte(fx.Body),
		Host:      fx.Host,
		Path:      fx.Path,
		Method:    fx.Method,
		Date:      fixedTime,
		QueryList: queryList,
		Headers:   header,
	}
	creds := base.Credentials{
		AccessKeyID:     fx.AK,
		SecretAccessKey: fx.SK,
		Region:          fx.Region,
		Service:         fx.Service,
	}

	sr := base.GetSignRequest(rp, creds)

	out := "Authorization: " + sr.Authorization + "\n" +
		"ContentType: " + sr.ContentType + "\n" +
		"Host: " + sr.Host + "\n" +
		"XContentSha256: " + sr.XContentSha256 + "\n" +
		"XDate: " + sr.XDate + "\n"

	outDir := firstNonEmptyEnv("SIGN_OUT_DIR", "L4_OUT_DIR")
	if outDir == "" {
		outDir = t.TempDir()
	} else if err := os.MkdirAll(outDir, 0755); err != nil {
		t.Fatalf("create SIGN_OUT_DIR: %v", err)
	}
	outPath := filepath.Join(outDir, "sign-go.txt")
	if err := os.WriteFile(outPath, []byte(out), 0644); err != nil {
		t.Fatalf("write output: %v", err)
	}

	if sr.Authorization == "" {
		t.Fatalf("Authorization is empty")
	}
	t.Logf("Authorization: %s", sr.Authorization)
	t.Logf("sign dump written to %s", outPath)
}

func firstNonEmptyEnv(keys ...string) string {
	for _, key := range keys {
		if key == "" {
			continue
		}
		if value := os.Getenv(key); value != "" {
			return value
		}
	}
	return ""
}

func findSignFixturePath() (string, bool) {
	const rel = "cospec/changes/check-tls-sdk-contract-alignment/context/l4-snapshots/fixture.json"

	dir, err := os.Getwd()
	if err != nil {
		return "", false
	}
	for {
		candidate := filepath.Join(dir, filepath.FromSlash(rel))
		if _, err := os.Stat(candidate); err == nil {
			return candidate, true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", false
		}
		dir = parent
	}
}
