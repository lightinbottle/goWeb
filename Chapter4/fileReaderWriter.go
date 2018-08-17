package main

import (
	//"io/ioutil"
	"log"
	"fmt"
	//"path/filepath"
	"os"
	"bufio"
	"io"
	"path/filepath"
	"io/ioutil"
)

func main(){
	//测试 文件的读写方式

	var err error
	//ioutil


	//filename , _ := exec.LookPath(os.Args[0])  //返回的是编译之后的可执行文件的路径
	filename :="/home/vode/goProject/goWeb/Chapter4/ioutil.txt"

	fmt.Println(filename)
	path,_ :=filepath.Abs(filename)        //绝对路径
	fmt.Println(path)
	rst := filepath.Dir(path)              //文件所在的目录
	fmt.Println(rst)

	fmt.Println(filepath.Dir(rst))         //上一级目录
	fmt.Println(filepath.Base(rst))        //returns the last element of path.



	// WriteFile writes data to a file named by filename.
	// If the file does not exist, WriteFile creates it with permissions perm;
	// otherwise WriteFile truncates it before writing.
	err=ioutil.WriteFile(filename,[]byte("ioutill\nvode\ntest"),0777)
	if err!=nil{
		log.Fatalln(err)
	}
	ioutil.ReadFile("")

	// ReadFile reads the file named by filename and returns the contents.
	// A successful call returns err == nil, not err == EOF. Because ReadFile
	// reads the whole file, it does not treat an EOF from Read as an error
	// to be reported.
	var readfile []byte
	readfile,err =ioutil.ReadFile(filename)     //读文件的所有的内容
	fmt.Println(string(readfile))


	//通过File结构对文件进行读写
	filename ="/home/vode/goProject/goWeb/Chapter4/file.txt"
	file,_ :=os.OpenFile(filename, os.O_RDWR|os.O_TRUNC|os.O_CREATE,0777)
	defer file.Close()
	// Write writes len(b) bytes to the File.
	// It returns the number of bytes written and an error, if any.
	// Write returns a non-nil error when n != len(b).
	fmt.Println(file.Write([]byte("file\nvode\ntest")))   //返回写入的字节数


	file,_ =os.OpenFile(filename, os.O_RDWR|os.O_CREATE,0777)    //注意要再开一个File结构进行读
	defer file.Close()

	readfile =make([]byte,6,10)
	fmt.Println(len(readfile))
	fmt.Println(file.Read(readfile))   //读最多len(readfile)的字节数
	fmt.Println(string(readfile))


	//file按行读取
	filename  ="/home/vode/goProject/goWeb/Chapter4/file.txt"
	file,_ =os.Open(filename)
	buffReader :=bufio.NewReader(file)

	for{
		line,_,err :=buffReader.ReadLine()
		if err!=nil{
			if err==io.EOF{
				break
			}else{
				log.Fatalln(err)
			}
		}
		fmt.Println(string(line))

	}


}
