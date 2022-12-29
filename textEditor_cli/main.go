package main

import (
	"bufio"
	"strings"
	//"io/ioutil"
	"os"

	//"io"
	"fmt"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var fileName string

	fmt.Print("Enter path + file name [txt files only]: ")
	fmt.Scan(&fileName)

	file, err := os.Create(fileName)
	checkPanic(err)

	fmt.Println("\t Type content..... \t\t\t\n'Note add a ^ and press enter to stop writting'")
	para, errRe := reader.ReadString('^')
	checkPanic(errRe)

	fmt.Println(para)
	para = strings.ReplaceAll(para, "^", " ")
	res, err2 := file.WriteString(para)
	checkPanic(err2)
	println(res)

	fmt.Println("Written Successfully")
	println(file)
	defer file.Close()

}

func checkPanic(e error) {
	if e != nil {
		panic(e)
	}
}
