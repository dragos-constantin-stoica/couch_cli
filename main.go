package main

import (
	"context"

	"github.com/gdamore/tcell/v2"
	kivik "github.com/go-kivik/kivik/v4"
	_ "github.com/go-kivik/kivik/v4/couchdb"
	"github.com/rivo/tview"
)

type CouchDBURL struct {
	fullURL string
}

var pages = tview.NewPages()
var app = tview.NewApplication()
var flex = tview.NewFlex()
var form = tview.NewForm()
var dbList = tview.NewList().ShowSecondaryText(false)
var docList = tview.NewList().ShowSecondaryText(true)
var docDetails = tview.NewTextView()

var text = tview.NewTextView().
	SetTextColor(tcell.ColorGreen).
	SetText("(o) to connect to CouchDB :: (q) to quit")

func main() {

	flex.SetDirection(tview.FlexRow).
		AddItem(tview.NewFlex().AddItem(dbList, 0, 1, true).AddItem(docList, 0, 4, false), 0, 6, false).
		AddItem(text, 0, 1, false)

	flex.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		if event.Rune() == 113 {
			app.Stop()
		} else if event.Rune() == 'o' {
			form.Clear(true)
			addOpenDBForm()
			pages.SwitchToPage("Open")
		}
		return event
	})

	pages.AddPage("Menu", flex, true, true)
	pages.AddPage("Open", form, true, false)

	if err := app.SetRoot(pages, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}

func addOpenDBForm() *tview.Form {
	openurl := CouchDBURL{}

	form.AddInputField("URL", "", 60, nil, func(dburl string) {
		openurl.fullURL = dburl
	})

	form.AddButton("Save", func() {
		client, err := kivik.New("couch", openurl.fullURL)
		if err != nil {
			panic(err)
		}
		dbs, err := client.AllDBs(context.TODO(), nil)
		if err != nil {
			panic(err)
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
