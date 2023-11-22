package main

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/gdamore/tcell/v2"
	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	"github.com/rivo/tview"
)

type CouchDBURL struct {
	fullURL string // http(s)://user:password@server:port
	protocol string //http, https
	user string
	password string
	server string
	port string
	DBname string
}

var pages = tview.NewPages()
var app = tview.NewApplication()
var flex = tview.NewFlex()
var form = tview.NewForm()
var dbList = tview.NewList().ShowSecondaryText(false)
var docList = tview.NewList().ShowSecondaryText(false)
var docDetails = tview.NewTextArea()
var docflex = tview.NewFlex()
var msgBox = tview.NewModal()

var client kivik.Client
var clientURL = CouchDBURL{}

var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(o) to connect to CouchDB :: (q) to quit")

func main() {

	dbList.SetBorder(true).SetTitle("Databases")
	docList.SetBorder(true).SetTitle("Documents")
	docDetails.SetTitle("Doc Details").SetBorder(true)

	docflex.SetDirection(tview.FlexRow).
	AddItem(docList, 0, 1, false).
	AddItem(docDetails, 0, 2, false)

	msgBox.AddButtons([]string{"OK"}).
		    		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		    			if buttonLabel == "OK" {
		    				pages.SwitchToPage("Menu")
		    			}
		    		})

	dbList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
		clientURL.DBname = name
		//messageBox(name)				
		//pages.SwitchToPage("Message")

		client, err := kivik.New("couch", clientURL.fullURL)
		if err != nil {
			messageBox(err.Error())
			pages.SendToFront("Messages")
			pages.ShowPage("Message")
			return
		}

		crtDB := client.DB(clientURL.DBname)
		
		if crtDB.Err() != nil {
			messageBox(crtDB.Err().Error())
			pages.SendToFront("Messages")
			pages.ShowPage("Message")
		}
		
		allDocs := crtDB.AllDocs(context.TODO(), nil)
		defer allDocs.Close()

		    docList.Clear()
		    docDetails.SetText("", false)
		    idx := 0
			for allDocs.Next() {				
				//add doc_id to docList
				var docid, _ = allDocs.ID()
				docList.AddItem(docid, " ", rune(49+idx) , nil)
				idx++
			}
			
			if allDocs.Err() != nil {
				messageBox(allDocs.Err().Error())
			pages.SendToFront("Messages")
			pages.ShowPage("Message")
			}		
	})

	docList.SetSelectedFunc(func(index int, name string, second_name string, shortcut rune) {
	
		client, err := kivik.New("couch", clientURL.fullURL)
		if err != nil {
			messageBox(err.Error())
			pages.SendToFront("Messages")
			pages.ShowPage("Message")
			return
		}        
		crtDB := client.DB(clientURL.DBname)
		//get the document with the _id:name
		var doc interface{}
	    err = crtDB.Get(context.TODO(), name).ScanDoc(&doc)
	    if err != nil {
			messageBox(err.Error())
			pages.SendToFront("Messages")
			pages.ShowPage("Message")
		}
		docString, err := json.MarshalIndent(doc,"","\t")
		docDetails.SetText(strings.ReplaceAll(string(docString),"\\n", "\n"), true)
		
	})
		
	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().AddItem(dbList, 0, 1, false).AddItem(docflex, 0, 4, false), 0, 1, false).
		AddItem(text, 1, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 'q' {
			app.Stop()
		} else if event.Rune() == 'o' {
			form.Clear(true)
			addOpenDBForm()
			pages.SendToFront("Open")
			pages.ShowPage("Open")
		}
		return event
	})

	// Returns a new primitive which puts the provided primitive in the center and
	// sets its size to the given width and height.
	modal := func(p tview.Primitive, width, height int) tview.Primitive {
		return tview.NewFlex().
			AddItem(nil, 0, 1, false).
			AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
				AddItem(nil, 0, 1, false).
				AddItem(p, height, 1, true).
				AddItem(nil, 0, 1, false), width, 1, true).
			AddItem(nil, 0, 1, false)
	}

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Open", modal(form, 70, 7), true, false)
	pages.AddPage("Message", msgBox, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addOpenDBForm() *tview.Form {

	form.AddInputField("URL", "", 60, nil, func(dburl string) {
		clientURL.fullURL = dburl
	})

	form.AddButton("Connect", func() {
		client, err := kivik.New("couch", clientURL.fullURL)
		if err != nil {
			messageBox(err.Error())
			pages.HidePage("Open")
			pages.SendToFront("Messages")
			pages.ShowPage("Message")
			return
		   	//panic(err)
		}
		dbs, err := client.AllDBs(context.TODO(), nil)
		if err != nil {
			messageBox(err.Error())
			pages.HidePage("Open")
			pages.SendToFront("Messages")
			pages.ShowPage("Message")
			return
			//panic(err)
		}
		dbList.Clear()
		for index, db := range dbs {
			dbList.AddItem(db, " ", rune(49+index), nil)
		}
		pages.SwitchToPage("Menu")
	})

	form.SetBorder(true).SetTitle("Connect to CouchDB server").SetTitleAlign(tview.AlignCenter)

	return form
}

func messageBox(msg string) *tview.Modal{
  msgBox.SetText(msg)
  return msgBox
}
