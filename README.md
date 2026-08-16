# go-bytes-dup

字节克隆仍共享

internal/bufx/dup.go 的 Dup：返回 s[:len:len] 仍共享底层，未 copy
