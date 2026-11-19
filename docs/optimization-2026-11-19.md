# 2026-11-19 扫描：问题与优化步骤

只读扫描本仓库。总览见 [r2day/docs](https://github.com/r2day/docs/blob/main/architecture/2026-11-19-api-optimization.md)。

只有 `Base64ToFile`。当前 18 个 API 的 `go.mod` 没有直接 require 本模块。本次没有发现密钥。

## 问题

### P2 测试把文件写在模块目录

`base64tofile_test.go` 使用 `os.CreateTemp("./", "test.txt")`。已跟踪的 `test.txt633158242` 是这次写法留下的 5 字节文件。

`Base64ToFile` 把整段解码进内存，没有长度上限。调用方如果接受用户上传的 base64，要在调用前限制大小。没有逗号时 `IndexByte` 为 -1，切片从 0 开始，不会越界。

## 步骤

1. 删除 `test.txt633158242`。
2. 测试改用 `t.TempDir()`。不要在模块目录创建文件。
3. 调用方限制输入长度后再解码。
