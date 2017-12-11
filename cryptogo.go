//fork from https://github.com/isfonzar/CryptoGo
package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/konvict/filecrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func main() {
	// If not enough args, return help text
	if len(os.Args) < 2 {
		printHelp()
		os.Exit(0)
	}

	function := os.Args[1]

	switch function {
	case "h":
		printHelp()
	case "e":
		encryptHandle()
	case "d":
		decryptHandle()
	default:
		fmt.Println("加密或解密一个文件。")
		os.Exit(1)
	}

}

func printHelp() {
	fmt.Println("CryptoGo")
	fmt.Println("MIT Open Source License")
	fmt.Println("原作者isfonzar，https://github.com/isfonzar/")
	fmt.Println("一个简单的问加密软件。")
	fmt.Println("")
	fmt.Println("用法:")
	fmt.Println("")
	fmt.Println("\tCryptoGo e 文件")
	fmt.Println("")
	fmt.Println("命令:")
	fmt.Println("")
	fmt.Println("\t e\t\t使用密码加密一个文件")
	fmt.Println("\t d\t\t使用密码解密一个文件")
	fmt.Println("\t h\t\t显示帮助信息")
	fmt.Println("")
}

func encryptHandle() {

	if len(os.Args) < 3 {
		println("缺少必要参数。需指定要加密的文件。")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("文件未找到！")
	}

	password := getPassword()

	fmt.Println("加密中...")
	filecrypt.Encrypt(file, password)
	fmt.Println("\n文件已成功加密")

}

func getPassword() []byte {
	fmt.Print("请输入加密密码: ")
	password, _ := terminal.ReadPassword(0)
	fmt.Print("\n请在输入一次: ")
	password2, _ := terminal.ReadPassword(0)
	if !validatePassword(password, password2) {
		fmt.Print("\n密码不匹配，请在输入一次。\n")
		return getPassword()
	}
	return password
}

func decryptHandle() {

	if len(os.Args) < 3 {
		println("缺少必要参数。需指定要加密的文件。")
		os.Exit(0)
	}

	file := os.Args[2]

	if !validateFile(file) {
		panic("文件未找到")
	}

	fmt.Print("请输入密码: ")
	password, _ := terminal.ReadPassword(0)

	fmt.Println("\n解密中...")
	filecrypt.Decrypt(file, password)
	fmt.Println("\n文件已成功解密。")

}

func validatePassword(password1 []byte, password2 []byte) bool {
	if !bytes.Equal(password1, password2) {
		return false
	}

	return true
}

func validateFile(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}

	return true
}
