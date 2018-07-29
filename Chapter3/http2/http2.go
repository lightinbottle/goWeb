package main

//http/2
//多个http handler的串联
//字符匹配

import (
	logg "log"
	"net/http"
	"runtime"
	"reflect"
	"fmt"
	//"golang.org/x/net/http2"
	"golang.org/x/net/http2"
)

func hello(w http.ResponseWriter,r *http.Request)  {

	fmt.Println("hello")
	//w.Write([]byte("hello"))
}

func index(w http.ResponseWriter,r *http.Request)  {

	fmt.Println("index")
	w.Write([]byte("index"))
}

func helloworld(w http.ResponseWriter,r *http.Request)  {

	fmt.Println("helloworld")
	w.Write([]byte("helloword"))
}

// reflect.ValueOf() returns a new Value initialized to the concrete value
// Pointer returns v's value as a uintptr. It panics if v's Kind is not Chan, Func, Map, Ptr, Slice, or UnsafePointer.
// FuncForPC returns a *Func describing the function that contains the given program counter address, or else nil.


func log(h http.HandlerFunc)http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		logg.Println("http handler: ",runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name())
		h(w,r)
	}
}

func main(){
	mux :=http.NewServeMux()
	mux.HandleFunc("/",log(index))
	mux.HandleFunc("/hello/",log(hello))       //  /hello/和/hello的区别就是是否是精确匹配
	mux.HandleFunc("/hello/world/",log(helloworld)) // 如果url为: http://localhost:1235/hello/world 会返回301 Moved Permanently


	server :=http.Server{
		Addr:":1235",
		Handler:mux,
	}

	// func ConfigureServer(s *http.Server, conf *Server)
	// ConfigureServer adds HTTP/2 support to a net/http Server.
	http2.ConfigureServer(&server,&http2.Server{})

	err :=server.ListenAndServeTLS("/home/vode/goProject/cert.pem","/home/vode/goProject/key.pem")
	//err :=server.ListenAndServe()
	if err!=nil{
		logg.Fatalln(err)
	}
}
