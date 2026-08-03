//go:build replay
// +build replay

package replay

import (
	"fmt"
	"os"
	"sort"
	"strings"
	"testing"
)

// TestReplay 端到端全量回放：逐条用真实流量 fixture 的 out 喂 SDK 反序列化链路。
//
// 无 fixtures 时（CI/新克隆，testdata/replay 不存在且未设 REPLAY_FIXTURES_DIR）跳过 PASS。
// 默认仅记录字段覆盖差异不 fail；设环境变量 REPLAY_STRICT=1 时缺失字段视为失败。
// 设 REPLAY_LIMIT_PER_URI=N 可限制每 uri 回放条数（0=全量），用于快速验证。
func TestReplay(t *testing.T) {
	fixtures := loadReplayFixtures("")
	if len(fixtures) == 0 {
		t.Skip("无回放 fixtures（testdata/replay 不存在）。生成方式见 tools/replaygen/README.md")
	}

	limitPerURI := envInt("REPLAY_LIMIT_PER_URI", 0)
	strict := os.Getenv("REPLAY_STRICT") == "1"

	server := newDynamicReplayServer()
	defer server.Close()
	client, err := newReplayClient(server.URL)
	if err != nil {
		t.Fatalf("构造回放 client 失败: %v", err)
	}

	// 按 uri 排序保证输出稳定
	uris := make([]string, 0, len(fixtures))
	for u := range fixtures {
		uris = append(uris, u)
	}
	sort.Strings(uris)

	var totalReplayed, totalMissing int
	for _, uri := range uris {
		uri := uri
		entry, ok := replayMap[uri]
		if !ok {
			entry = newRawReplayEntry(uri)
		}
		rawEntry := newRawReplayEntry(uri)
		t.Run(uri, func(t *testing.T) {
			fs := fixtures[uri]
			// 每 uri 回放的缺失字段并集（去重），最终汇总
			missingSet := map[string]struct{}{}
			var buildFail, callFail, deserialFail, fallbackCount int
			n := 0
			for i := range fs {
				if limitPerURI > 0 && n >= limitPerURI {
					break
				}
				f := &fs[i]
				// 切换 mock 响应为当前这条 fixture 的 out
				server.setOut(uri, f.Out)

				activeEntry := entry
				usedRaw := entry.family == "raw"
				apiReq, err := activeEntry.build(f.In)
				if err != nil {
					buildFail++
					activeEntry = rawEntry
					usedRaw = true
					apiReq, err = activeEntry.build(f.In)
					if err != nil {
						continue
					}
					fallbackCount++
				}
				resp, err := activeEntry.call(client, apiReq)
				if err != nil {
					// 反序列化错误是重要发现（类型不匹配），记录但不阻断
					callFail++
					if isDeserialErr(err) {
						deserialFail++
						apiReq, rawErr := rawEntry.build(f.In)
						if rawErr == nil {
							var rawResp interface{}
							rawResp, rawErr = rawEntry.call(client, apiReq)
							if rawErr == nil {
								resp = rawResp
								activeEntry = rawEntry
								usedRaw = true
								fallbackCount++
							}
						}
					}
					if !usedRaw {
						continue
					}
				}
				if resp == nil {
					continue
				}
				_, replyData := activeEntry.reply(resp)

				// 字段覆盖差异（成功响应才对比）
				if f.Category == "success" && entry.family != "raw" {
					for _, m := range fieldCoverage(f.Out["data"], replyData) {
						missingSet[m] = struct{}{}
					}
				}
				n++
			}
			totalReplayed += n

			missing := sortedKeys(missingSet)
			totalMissing += len(missing)
			t.Logf("[%s] %s: 回放 %d 条（共 %d）build失败=%d call失败=%d(其中反序列化=%d) 原始降级=%d 缺失字段 %d 个",
				entry.family, uri, n, len(fs), buildFail, callFail, deserialFail, fallbackCount, len(missing))
			if len(missing) > 0 {
				t.Logf("  缺失: %s", strings.Join(missing, ", "))
				if strict {
					t.Errorf("strict 模式：缺失字段 %v", missing)
				}
			}
		})
	}

	t.Logf("=== 回放汇总：覆盖 %d 个接口，回放 %d 条，缺失字段总数 %d（strict=%v, limit=%d）===",
		len(uris), totalReplayed, totalMissing, strict, limitPerURI)
}

// TestReplaySkipWhenNoFixtures 验证无 fixtures 时的跳过逻辑（CI 场景）。
func TestReplaySkipWhenNoFixtures(t *testing.T) {
	fixtures := loadReplayFixtures("/nonexistent/replay/path")
	if fixtures != nil {
		t.Fatalf("期望 nil，got %v", fixtures)
	}
}

// sortedKeys 返回 map key 的排序切片。
func sortedKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

// isDeserialErr 判断是否为 JSON 反序列化错误（类型不匹配等重要发现）。
func isDeserialErr(err error) bool {
	return err != nil && containsAny(err.Error(),
		"cannot unmarshal", "json:", "into Go struct field")
}

func containsAny(s string, subs ...string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}

func envInt(key string, def int) int {
	v := os.Getenv(key)
	if v == "" {
		return def
	}
	var n int
	fmt.Sscanf(v, "%d", &n)
	return n
}
