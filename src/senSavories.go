package main

import (
	"net/http"

	"github.com/jarrancarr/website"
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
	senSavories.AddMenu("nav").
		AddItem("Test", "/test").
		AddItem("Home", "/home").
		AddItem("Login", "/login").
		Add("nav nav-pills nav-stacked", "", "")
	
	// services
	acs := website.CreateAccountService()
	senSavories.AddService("account", acs)
	ecs := ecommerse.CreateService(acs)
	senSavories.AddService("ecommerse", ecs)
	senSavories.AddSiteProcessor("secure", acs.CheckSecure)

	ecs.AddCategories("Oils", "Olive Oils", "oils.png")
	ecs.AddCategories("Vinegars", "Aged Balsamic Vinegars", "vinegars.png")
	ecs.AddCategories("Spices", "African Spices", "spices.png")
	ecs.AddCategories("Teas", "Quality East African Teas", "teas.png")
	ecs.AddCategories("Coffee", "Coffee from Africa", "coffee.png")

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
}