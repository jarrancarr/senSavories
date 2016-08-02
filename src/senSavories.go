package main

import (
	//"crypto/md5"
	"fmt"
	//"io"
	//"io/ioutil"
	"net/http"
	//"os"
	//"strconv"
	"html/template"
	//"strings"
	//"time"

	"github.com/jarrancarr/website"
	"github.com/jarrancarr/website/ecommerse"
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
		AddItem("Test", "/test").
		AddItem("Edit", "/edit").
		AddItem("Secure", "/secure").
		AddItem("Home", "/home").
		AddItem("Message", "/message").
		AddItem("Login", "/login").
		Add("nav nav-pills nav-stacked", "", "")

	// services
	acs := website.CreateAccountService()
	senSavories.AddService("account", acs)
	mgs := service.CreateMessageService(acs)
	senSavories.AddService("message", mgs)
	senSavories.AddSiteProcessor("secure", acs.CheckSecure)

	// template subpages
	senSavories.AddPage("", "head", "")
	senSavories.AddPage("", "banner", "")
	senSavories.AddPage("nav", "nav", "")

	// pages
	main := senSavories.AddPage("senSavories", "main", "/")
	main.AddBypassSiteProcessor("secure")
	main.AddTable("cart", []string{"A", "B", "C", "D"}, []string{"1", "2", "3", "4"}).AddClass("table").AddId("cart")
	main.AddTable("ppppxxxx", []string{"X", "Y", "Z"}, []string{"91", "82", "73", "64", "55", "46"}).AddClass("table").AddId("ppppxxxx")
	senSavories.AddPage("Home", "home", "/home")
	senSavories.AddPage("senSavories-edit", "edit", "/edit")
	test := senSavories.AddPage("test", "test", "/test")
	test.AddAJAXHandler("test123", mgs.TestAJAX)
	senSavories.AddPage("", "secure", "/secure")
	senSavories.AddPage("message", "message", "/message")
	login := senSavories.AddPage("login", "login", "/login")
	login.AddPostHandler("login", acs.LoginPostHandler)
	login.AddBypassSiteProcessor("secure")
	
	chess := senSavories.AddPage("chess", "chess", "/chess")
	
	scaleX := 45
	scaleY := 35
	offX := 120
	offY := 0
	spaces := 4
	for y := 0; y<spaces; y++ {
		for x := 0; x<spaces+1+y; x++ {
			if x>0 {				
				chess.Data["svg"] = append(chess.Data["svg"], 
					template.HTML(
						fmt.Sprintf(
							"<polygon points='%d,%d %d,%d %d,%d' style='fill:#842;stroke:purple;stroke-width:1' />",
								offX+(2*x-y)*scaleX, offY+2*scaleY*y, 
								offX+(2*x-y+1)*scaleX, offY+scaleY*2+2*scaleY*y, 
								offX+(2*x-y+2)*scaleX, offY+2*scaleY*y))) 
			}
			chess.Data["svg"] = append(chess.Data["svg"], 
				template.HTML(
					fmt.Sprintf(
						"<polygon points='%d,%d %d,%d %d,%d' style='fill:#482;stroke:purple;stroke-width:1' />",
							offX+(2*x-y+1)*scaleX, offY+scaleY*2+2*scaleY*y, 
							offX+(2*x-y+2)*scaleX, offY+2*scaleY*y, 
							offX+(2*x-y+3)*scaleX, offY+scaleY*2+2*scaleY*y)))
			
		}
	}
	for y := spaces; y<spaces*2; y++ {
		for x := 0; x<spaces*3-y; x++ {
			if x>0 {		
				chess.Data["svg"] = append(chess.Data["svg"], 
					template.HTML(
						fmt.Sprintf(
							"<polygon points='%d,%d %d,%d %d,%d' style='fill:#482;stroke:purple;stroke-width:1' />",
								offX+(2*x+y+3-spaces*2)*scaleX, offY+scaleY*2+2*scaleY*y, 
								offX+(2*x+y+2-spaces*2)*scaleX, offY+2*scaleY*y, 
								offX+(2*x+y+1-spaces*2)*scaleX, offY+scaleY*2+2*scaleY*y)))	
			}	
			chess.Data["svg"] = append(chess.Data["svg"], 
				template.HTML(
					fmt.Sprintf(
						"<polygon points='%d,%d %d,%d %d,%d' style='fill:#842;stroke:purple;stroke-width:1' />",
							offX+(2*x+y+2-spaces*2)*scaleX, offY+2*scaleY*y, 
							offX+(2*x+y+3-spaces*2)*scaleX, offY+scaleY*2+2*scaleY*y, 
							offX+(2*x+y+4-spaces*2)*scaleX, offY+2*scaleY*y))) 
			
		}
	}
	chess.AddAJAXHandler("test123", mgs.TestAJAX)
	chess.AddBypassSiteProcessor("secure")
	

	//	Shelf = []ecommerse.Category{
	//		ecommerse.Category{"Oils", "Olive Oils", "oils.png"},
	//		ecommerse.Category{"Vinegars", "Aged Balsamic Vinegars", "vinegars.png"},
	//		ecommerse.Category{"Spices", "African Spices", "spices.png"},
	//		ecommerse.Category{"Teas", "Quality East African Teas", "teas.png"},
	//		ecommerse.Category{"Coffee", "Coffee from Africa", "coffee.png"},
	//	}
}
