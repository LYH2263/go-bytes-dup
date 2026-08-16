# 字节克隆仍共享

internal/bufx/dup.go 的 Dup：返回 s[:len:len] 仍共享底层，未 copy

```bash
go build ./...
go test ./... -count=1
go vet ./...
```
