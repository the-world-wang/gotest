package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//func Open(name string) (file *File, err error) {
	// return OpenFile(name, O_RDONLY, 0)
	// }
	// Open方法其实也是利用OpenFile实现，只读模式
	file, _ := os.Open("test.txt")
	br := bufio.NewReader(file)
	for {
		line, readErr := br.ReadString('\n')
		if readErr == io.EOF {
			break
		}
		fmt.Println(line)
	}
	defer file.Close()

	//2.利用直接读取字符串到一个byte切片中
	readByte, erro := ioutil.ReadFile("test.txt")
	if erro != nil {
		fmt.Println(erro)
	}
	fmt.Println(string(readByte))

	//3.带缓冲的读取,Read reads data into p.
	// buf := make([]byte, 1024)
	// br.Read(buf)

	// fmt.Println(string(buf))

	// 4.写文件,以只读形式打开
	writeFile, outputError := os.OpenFile("test.txt", os.O_WRONLY, 0666)
	if outputError != nil {
		fmt.Println(outputError)
	}
	defer writeFile.Close()

	writer := bufio.NewWriter(writeFile)
	writer.WriteString("我才不想")
	writer.Flush()

	// 5.Copy
	// 利用io.Copy()完成Fie的拷贝
}
