package main

import (
	//"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	//"strconv"
	"html/template"
	"strings"
	//"time"

	"github.com/jarrancarr/sensavories/src/html"
	//"github.com/jarrancarr/sensavories/src/sensavories"
)

// var home, test *Page
var resourceDir = "../../"

func main() {
	//	senSavories := createSite("senSavories")
	//	home, _ := loadPage(senSavories, "SenSavories", "main")
	//	sensavories.Init()
	//	home.AddTable("cart", []string{"A", "B"}, []string{"1", "2", "3", "4"})
	//	home.GetTable("cart").AddClass("table")
	//	home.AddMenu("nav", []html.HTMLLink{html.Link("Menu1", "/test"), html.Link("Menu2", "/test"), html.Link("Menu3", "/test")})

	testSite := createSite("testSite")
	testBody, _ := loadPage(testSite, "testNavData", "body", "")
	testSite.AddPage("home", "/", testBody)
	testBody.AddTable("data", []string{"A", "B"}, []string{"1", "2", "3", "4"})
	testBody.GetTable("data").AddClass("table")
	testNav, _ := loadPage(testSite, "testNavData", "testNav", "")
	testNav.AddMenu("nav", []html.HTMLLink{html.Link("edit", "/edit"), html.Link("TestMenu2", "/"), html.Link("TestMenu3", "/")})
	test, _ := loadPage(testSite, "SenSavories", "test", "/test")
	test.AddPage("nav", testNav)
	test.AddPage("body", testBody)
	loadPage(testSite, "SenSavories", "edit", "/edit")
	loadPage(testSite, "SenSavories", "edit", "/edit/addProduct")

	http.HandleFunc("/js/", serveResource)
	http.HandleFunc("/css/", serveResource)
	http.HandleFunc("/img/", serveResource)
	http.HandleFunc("/upload", testSite.upload)
	http.ListenAndServe(":8080", nil)
}

func serveResource(w http.ResponseWriter, r *http.Request) {
	path := resourceDir + "public" + r.URL.Path
	if strings.HasSuffix(r.URL.Path, "js") {
		w.Header().Add("Content-Type", "application/javascript")
	} else if strings.HasSuffix(r.URL.Path, "css") {
		w.Header().Add("Content-Type", "text/css")
	} else if strings.HasSuffix(r.URL.Path, "png") {
		w.Header().Add("Content-Type", "image/svg+xml")
	} else if strings.HasSuffix(r.URL.Path, "svg") {
		w.Header().Add("Content-Type", "image/svg+xml")
	}

	data, err := ioutil.ReadFile(path)

	if err == nil {
		w.Write(data)
	} else {
		w.WriteHeader(404)
		w.Write([]byte("404, My Friend - " + http.StatusText(404)))
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		//		crutime := time.Now().Unix()
		//		h := md5.New()
		//		io.WriteString(h, strconv.FormatInt(crutime, 10))
		//		token := fmt.Sprintf("%x", h.Sum(nil))

		//		t, _ := template.ParseFiles("upload.gtpl")
		//		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		f, err := os.OpenFile("../../temp/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		t, _ := template.New("foo").Parse(`File Uploaded`)
		t.Execute(w, "")
	}
}
