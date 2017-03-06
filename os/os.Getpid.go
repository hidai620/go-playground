package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("プロセスID",os.Getpid())
	fmt.Println("親プロセスID",os.Getppid())
	fmt.Println("ユーザーID",os.Getuid())
	fmt.Println("グループID",os.Getgid())
	fmt.Println("実行グループID", os.Getegid())
}