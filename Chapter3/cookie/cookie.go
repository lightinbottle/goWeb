package main

import (
	"net/http"
	"fmt"
	"time"
)

func setCookie(w http.ResponseWriter,r *http.Request){
	_,err :=r.Cookie("cookietest")
	if err !=http.ErrNoCookie{
		w.Write([]byte("cookie has equippded"))
		return
	}
	cookie :=http.Cookie{
		Name:"cookietest",
		Value:"test http cookie",

		}
	http.SetCookie(w,&cookie)
	return
}

func blinkResponse(w http.ResponseWriter,r *http.Request){
	_,err :=r.Cookie("cookietest")
	if err==http.ErrNoCookie{
		w.Write([]byte("no cookie found"))
		return
	}
	cookie :=http.Cookie{
		Name:"cookietest",
		Value:"test http cookie",
		//MaxAge:-1,                    //cookie 生命周期时间 单位是second
		Expires:time.Now().Add(time.Second*10),   //直到 Expires 指定的时间失效

	}
	http.SetCookie(w,&cookie)
	w.Write([]byte("blink response, cookie outdate"))
	return
}

func hello(w http.ResponseWriter,r *http.Request){
	_,err :=r.Cookie("cookietest")
	if err==http.ErrNoCookie{
		w.Write([]byte("no cookie found"))
		return
	}
	w.Write([]byte("hello"))
	return
}

func main(){
	mux :=http.NewServeMux()
	mux.HandleFunc("/setcookie",setCookie)
	mux.HandleFunc("/hello",hello)
	mux.HandleFunc("/blink",blinkResponse)

	server :=http.Server{
		Handler:mux,
		Addr:":1234",

	}

	err :=server.ListenAndServe()
	if err !=nil{
		fmt.Println(err)
	}

}