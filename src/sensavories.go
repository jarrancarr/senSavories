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
		AddItem(&html.HTMLMenuItem{"/login", "login", html.HTMLElement{}}).
		Add("nav nav-pills nav-stacked", "", "")

	nav, navPageErr := website.LoadPage(senSavories, "nav", "nav", "")
	if navPageErr != nil {
		fmt.Println("Error creating web page: " + navPageErr.Error())
		os.Exit(1)
	}
	senSavories.AddPage("nav", nav)

	main, mainPageErr := website.LoadPage(senSavories, "SenSavories", "main", "/")
	if mainPageErr != nil {
		fmt.Println("Error creating web page: " + mainPageErr.Error())
		os.Exit(1)
	}

	website.LoadPage(senSavories, "Home", "home", "/home")
	//	if homePageErr != nil {
	//		fmt.Println("Error creating web page: " + homePageErr.Error())
	//		os.Exit(1)
	//	}

	website.LoadPage(senSavories, "SenSavories", "edit", "/edit")

	secure, securePageErr := website.LoadPage(senSavories, "Home", "home", "/secure")
	if securePageErr != nil {
		fmt.Println("Error creating web page: " + securePageErr.Error())
		os.Exit(1)
	}
	secure.SetRequireLogin()

	main.AddTable("cart", []string{"A", "B", "C", "D"}, []string{"1", "2", "3", "4"}).AddClass("table")

	Shelf = []ecommerse.Category{
		ecommerse.Category{"Oils", "Olive Oils", "oils.png"},
		ecommerse.Category{"Vinegars", "Aged Balsamic Vinegars", "vinegars.png"},
		ecommerse.Category{"Spices", "African Spices", "spices.png"},
		ecommerse.Category{"Teas", "Quality East African Teas", "teas.png"},
		ecommerse.Category{"Coffee", "Coffee from Africa", "coffee.png"},
	}
}
