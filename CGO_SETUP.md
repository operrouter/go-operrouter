# CGO 配置说明

## 问题背景

`ffi_client.go` 使用 CGO 调用 C 库，这会导致 gopls（Go 语言服务器）显示错误：

```
go list failed to return CompiledGoFiles. This may indicate failure to perform cgo processing
```

## 解决方案

### 1. 已添加的配置

项目已包含以下配置文件：

- **`.vscode/settings.json`** - VS Code 的 gopls 配置
- **`.golangci.yml`** - Linter 配置，忽略 CGO 错误
- **`operrouter/ffi_client.go`** - 添加了 `//go:build cgo` 标签

### 2. IDE 设置

#### VS Code

配置已自动应用，如果仍有问题：

1. 重启 Go 语言服务器：
   - `Ctrl+Shift+P` (Windows/Linux) 或 `Cmd+Shift+P` (Mac)
   - 输入 "Go: Restart Language Server"

2. 重启 VS Code

#### GoLand / IntelliJ IDEA

在 Settings → Go → Build Tags and Vendoring 中添加：
```
CGO_ENABLED=1
```

### 3. 命令行编译

确保启用 CGO：

```bash
# 构建
CGO_ENABLED=1 go build ./...

# 测试
CGO_ENABLED=1 go test ./...

# 运行示例
CGO_ENABLED=1 go run examples/datasource_ffi/main.go
```

### 4. 环境变量设置（可选）

在 shell 配置文件中添加：

```bash
# ~/.bashrc 或 ~/.zshrc
export CGO_ENABLED=1
```

## 为什么会出现这个错误？

这是 gopls 的已知限制（[Issue #38990](https://github.com/golang/go/issues/38990)）：

1. gopls 不执行完整的 CGO 编译流程
2. 无法解析 `import "C"` 和 C 类型定义
3. 显示大量 `undefined: C.*` 错误

**这些错误是误报** - 代码在命令行可以正常编译和运行。

## 验证配置

运行以下命令验证配置是否正确：

```bash
# 应该显示 "1"
go env CGO_ENABLED

# 应该成功编译
CGO_ENABLED=1 go build ./operrouter

# 验证所有包
CGO_ENABLED=1 go build ./...
```

## 常见问题

### Q: 为什么 IDE 仍然显示错误？

A: 这是正常的，gopls 无法完全理解 CGO 代码。只要命令行编译成功即可。

### Q: 如何完全隐藏这些错误？

A: 无法完全隐藏，但可以：
1. 忽略 `ffi_client.go` 中的红色波浪线
2. 关注其他非 CGO 文件的错误
3. 使用命令行进行最终验证

### Q: FFI 后端是否必需？

A: 不是。如果不需要 FFI 功能，可以只使用 HTTP 或 gRPC 后端：

```go
// 使用 HTTP 后端（不需要 CGO）
client := operrouter.NewHTTP("http://localhost:8080")

// 使用 gRPC 后端（不需要 CGO）
client, _ := operrouter.NewGRPC("localhost:50051")
```

## 相关链接

- [Go Issue #38990](https://github.com/golang/go/issues/38990)
- [CGO Documentation](https://golang.org/cmd/cgo/)
- [gopls Configuration](https://github.com/golang/tools/blob/master/gopls/doc/settings.md)
