# wire in action

使用 `wire`

问题：原生的 `wire` 无法递归生成

解决：可使用 `go install github.com/ranxx/wire/cmd/wire@dev`，然后在你的项目敲如下命令 `wire gen --recursion`
> github.com/ranxx/wire 是在 github.com/google/wire 的 `wire gen` 命令基础上新增 `--recursion` flag 完成的递归逻辑，请放心使用

示例：
```bash
➜  wire-in-action: go install github.com/ranxx/wire/cmd/wire@dev

➜  wire-in-action: wire gen --recursion

wire: github.com/ranxx/wire-in-action/foobarbaz: wrote /Users/axing/Code/ranxx/wire-in-action/foobarbaz/wire_gen.go
wire: github.com/ranxx/wire-in-action/foobarbaz2: wrote /Users/axing/Code/ranxx/wire-in-action/foobarbaz2/wire_gen.go
```