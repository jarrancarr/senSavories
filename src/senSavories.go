package main

import (
	//"crypto/md5"
	//"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	//"os"
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
	//website
	senSavories = website.CreateSite("senSavories")
	senSavories.AddMenu("nav").
		AddItem(&html.HTMLMenuItem{"/test", "Test", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/edit", "Edit", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/secure", "Secure", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/home", "Home", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/message", "message", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/login", "login", html.HTMLElement{}}).
		Add("nav nav-pills nav-stacked", "", "")

	acs := website.CreateAccountService()
	senSavories.AddService("account", acs)
	senSavories.AddService("message", service.CreateMessageService())

	// template subpages
	senSavories.AddPage("head", "head", "")
	senSavories.AddPage("nav", "nav", "")

	// pages
	main := senSavories.AddPage("senSavories", "main", "/")
	main.AddTable("cart", []string{"A", "B", "C", "D"}, []string{"1", "2", "3", "4"}).AddClass("table")
	senSavories.AddPage("Home", "home", "/home")
	senSavories.AddPage("senSavories-edit", "edit", "/edit")
	senSavories.AddPage("senSavories", "test", "/test")
	senSavories.AddPage("message", "message", "/message")
	login := senSavories.AddPage("login", "login", "/login")
	login.AddPostHandler("login", acs.LoginPostHandler)

	//	Shelf = []ecommerse.Category{
	//		ecommerse.Category{"Oils", "Olive Oils", "oils.png"},
	//		ecommerse.Category{"Vinegars", "Aged Balsamic Vinegars", "vinegars.png"},
	//		ecommerse.Category{"Spices", "African Spices", "spices.png"},
	//		ecommerse.Category{"Teas", "Quality East African Teas", "teas.png"},
	//		ecommerse.Category{"Coffee", "Coffee from Africa", "coffee.png"},
	//	}
}
