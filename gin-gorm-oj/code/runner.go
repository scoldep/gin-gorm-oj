package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	// go run code-user/main.go
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, stderr bytes.Buffer
	cmd.Stderr = &stderr // 标准err
	cmd.Stdout = &out    // Stdout-标准输出值
	stdinPipe, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	io.WriteString(stdinPipe, "23 11\n") // 通过stdinPipe这个管道，往里面写数据
	// 根据测试的输入案例进行运行，拿到输出结果和标准的输出结果进行比对（如果一样，说明用户的答案正确）
	if err := cmd.Run(); err != nil {
		log.Fatal(err, stderr.String())
	}
	fmt.Println(out.String()) // 以字符串的形式打印输出

	fmt.Println(out.String() == "34\n") // 进行比对，结果为true，测试通过
}
