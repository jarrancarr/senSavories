package main

import (
	//"crypto/md5"
	"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	"os"
	//"strconv"
	//"html/template"
	//"strings"
	//"time"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
	"github.com/jarrancarr/website/html"
	"github.com/jarrancarr/website/service"
)

var Shelf []ecommerse.Category
var senSavories *website.Site

func main() {
	website.ResourceDir = ".."
	setup()

	http.HandleFunc("/js/", website.ServeResource)
	http.HandleFunc("/css/", website.ServeResource)
	http.HandleFunc("/img/", website.ServeResource)
	http.ListenAndServe(":8090", nil)
}

func setup() {
	senSavories = website.CreateSite("senSavories")
	senSavories.AddMenu("nav").
		AddItem(&html.HTMLMenuItem{"/test", "Test", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/edit", "Edit", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/secure", "Secure", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/home", "Home", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/message", "message", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/login", "login", html.HTMLElement{}}).
		Add("nav nav-pills nav-stacked", "", "")

	senSavories.AddService("message", service.CreateMessageService())
	secure := service.CreateAccountService()
	senSavories.AddService("account", secure)
	head := addPage(senSavories, "", "head", "/head")
	senSavories.AddPage("head", head)
	nav := addPage(senSavories, "nav", "nav", "")
	senSavories.AddPage("nav", nav)

	main := addPage(senSavories, "senSavories", "main", "/")
	main.AddTable("cart", []string{"A", "B", "C", "D"}, []string{"1", "2", "3", "4"}).AddClass("table")
	addPage(senSavories, "Home", "home", "/home")
	addPage(senSavories, "senSavories-edit", "edit", "/edit")
	addPage(senSavories, "senSavories", "test", "/test")
	addPage(senSavories, "Home", "home", "/secure").AddInitProcessor(secure.ValidateSession)
	addPage(senSavories, "message", "message", "/message")
	addPage(senSavories, "", "login", "/login").AddPostHandler("login",
		func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte(senSavories.Service["account"].Execute(
				"Login",
				r.FormValue("UserName"),
				r.FormValue("Password"))))
		})

	Shelf = []ecommerse.Category{
		ecommerse.Category{"Oils", "Olive Oils", "oils.png"},
		ecommerse.Category{"Vinegars", "Aged Balsamic Vinegars", "vinegars.png"},
		ecommerse.Category{"Spices", "African Spices", "spices.png"},
		ecommerse.Category{"Teas", "Quality East African Teas", "teas.png"},
		ecommerse.Category{"Coffee", "Coffee from Africa", "coffee.png"},
	}
}

func addPage(site *website.Site, name, template, url string) *website.Page {
	page, err := website.LoadPage(site, name, template, url)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return page
}
