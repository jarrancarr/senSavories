package main

import (
	//"fmt"
	"net/http"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/html"
	"github.com/jarrancarr/website/ecommerse"
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
	senSavories.Html.Tag("nav", html.NewTag3("ul", "", "nav nav-pills nav-stacked", "", "").
		AppendChild(html.NewTag("li").AppendChild(html.NewTag("a href=/test Test"))).
		AppendChild(html.NewTag("li").AppendChild(html.NewTag("a href=/home Home"))).
		AppendChild(html.NewTag("li").AppendChild(html.NewTag("a href=/login Login"))))
	
	// services
	acs := website.CreateAccountService()
	senSavories.AddService("account", acs)
	ecs := ecommerse.CreateService(acs)
	senSavories.AddService("ecommerse", ecs)
	senSavories.AddSiteProcessor("secure", acs.CheckSecure)

	ecs.AddCategory("Oils", "Olive Oils", "oils.png")
	ecs.AddCategory("Vinegars", "Aged Balsamic Vinegars", "vinegars.png")
	ecs.AddCategory("Spices", "African Spices", "spices.png")
	ecs.AddCategory("Teas", "Quality East African Teas", "teas.png")
	ecs.AddCategory("Coffee", "Coffee from Africa", "coffee.png")
	ecs.AddProduct("Oils", "Orange Oil", "Virgin press olive oil with orange oil essence.", "OrangeOliveOil.bpg", 1790, 72)
	ecs.AddProduct("Oils", "Sage Oil", "Virgin press olive oil with infused sage.", "SageOliveOil.bpg", 1590, 41)
	ecs.AddProduct("Oils", "Fescheu Oil", "First press virgin olive oil from the Fescheu orchard.", "FescheuOliveOil.bpg", 1490, 95)
	ecs.AddProduct("Coffee", "Kenyan Coffee", "Premier Karibou Coffee.", "KaribouCoffee.bpg", 1590, 33)

	// template subpages
	senSavories.AddPage("", "head", "")
	senSavories.AddPage("", "banner", "")
	senSavories.AddPage("nav", "nav", "")

	// pages
	main := senSavories.AddPage("senSavories", "main", "/")
	main.AddBypassSiteProcessor("secure")
	login := senSavories.AddPage("login", "login", "/login")
	login.AddBypassSiteProcessor("secure")
	login.AddPostHandler("login", acs.LoginPostHandler)
	test := senSavories.AddPage("test", "test", "/test")
	test.AddAJAXHandler("categories", ecs.GetCategories)
	test.AddAJAXHandler("products", ecs.GetProducts)
	test.AddBypassSiteProcessor("secure")
}