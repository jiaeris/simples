package cookie

import (
	"net/http"
	"time"
)

func open() {
	http.Handle("/1", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	}))
	http.HandleFunc("/2", func(w http.ResponseWriter, r *http.Request) {

	})
	http.Handle("/3", &CC{})
	http.HandleFunc("/4", login)
}

type CC struct {

}

func (c *CC) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	cookie, err := r.Cookie("test")
	if err != nil {
		panic(err)
	}
	if cookie.Value == "" {
		//TODO
		return
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	name := r.Form.Get("name")
	pass := r.Form.Get("pass")
	//TODO

	cookie := http.Cookie{
		Name:"test",
		Value:name + pass,
		Expires:time.Now().Add(time.Hour * 30),
		//可访问域名，即主目录
		//Domain:"www.heiheihei.com",
		//在domain下的范围目录，如www.heiheihei.com/test表示只有test目录可以访问cookie，"/"所有目录可以访问
		//Path:"/test",
		//是否仅通过安全的https,值为0或1
		//Secure:1,
	}
	http.SetCookie(w, cookie)
}
