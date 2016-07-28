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
	senSavories = website.CreateSite("senSavories", "localhost:8090")
	senSavories.AddMenu("nav").
		AddItem(&html.HTMLMenuItem{"/test", "Test", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/edit", "Edit", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/secure", "Secure", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/home", "Home", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/message", "Message", html.HTMLElement{}}).
		AddItem(&html.HTMLMenuItem{"/login", "Login", html.HTMLElement{}}).
		Add("nav nav-pills nav-stacked", "", "")

	acs := website.CreateAccountService()
	senSavories.AddService("account", acs)
	mgs := service.CreateMessageService(acs)
	senSavories.AddService("message", mgs)

	// template subpages
	senSavories.AddPage("", "head", "")
	senSavories.AddPage("", "banner", "")
	senSavories.AddPage("nav", "nav", "")

	// pages
	main := senSavories.AddPage("senSavories", "main", "/")
	main.AddTable("cart", []string{"A", "B", "C", "D"}, []string{"1", "2", "3", "4"}).AddClass("table")
	senSavories.AddPage("Home", "home", "/home")
	senSavories.AddPage("senSavories-edit", "edit", "/edit")
	test := senSavories.AddPage("test", "test", "/test")
	test.AddAJAXHandler("test123", mgs.TestAJAX)
	senSavories.AddPage("", "secure", "/secure")
	senSavories.AddPage("message", "message", "/message")
	login := senSavories.AddPage("login", "login", "/login")
	login.AddPostHandler("login", acs.LoginPostHandler)
	login.AddBypassSiteProcessor("secure")
	
	senSavories.AddSiteProcessor("secure", acs.CheckSecure)

	//	Shelf = []ecommerse.Category{
	//		ecommerse.Category{"Oils", "Olive Oils", "oils.png"},
	//		ecommerse.Category{"Vinegars", "Aged Balsamic Vinegars", "vinegars.png"},
	//		ecommerse.Category{"Spices", "African Spices", "spices.png"},
	//		ecommerse.Category{"Teas", "Quality East African Teas", "teas.png"},
	//		ecommerse.Category{"Coffee", "Coffee from Africa", "coffee.png"},
	//	}
}
