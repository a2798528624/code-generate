package main

import quickstart "github.com/ACking-you/quickstart_project"

func autoQuickStart() {
	config := quickstart.DefaultConfig("go-code-generate/quick-start-demo/example", "root", "a1131657944", "127.0.0.1", 3306, "sqlc-test").
		//打印出生成结果
		EnableDebug(true).
		//改变基本路径（默认为项目根目录）
		BasePath("./example")

	err := quickstart.Run(config)
	if err != nil {
		panic(err)
	}
}
func main() {
	autoQuickStart()
}
