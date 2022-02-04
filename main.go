package main

import (
	"database/sql"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
var err error

func init() {

	// Connect to database
	db, err = sql.Open("sqlite3", "./ninierp.db")
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(1)
	defer db.Close()
	log.Println("db connection successful")

}

func main() {

	addClient("Tareq", "087393", "tareq@gmail.com", "Dhaka")
	// os.Exit(1)

	App := app.New()
	Window := App.NewWindow("Add client")

	Window.Resize(fyne.NewSize(500, 400))

	title := canvas.NewText("Add new client", color.White)
	title.TextSize = 20
	title.Alignment = fyne.TextAlignCenter

	name := widget.NewEntry()
	name.PlaceHolder = "Enter Customer's name"
	mobile := widget.NewEntry()
	mobile.PlaceHolder = "Enter Customer's Mobile No."
	email := widget.NewEntry()
	email.PlaceHolder = "Enter Customer's email address"
	address := widget.NewEntry()
	address.PlaceHolder = "Enter Customer's address"

	row1 := widget.NewFormItem("Name", name)
	row2 := widget.NewFormItem("Mobile", mobile)
	row3 := widget.NewFormItem("Email", email)
	row4 := widget.NewFormItem("Address", address)

	wform := widget.NewForm(row1, row2, row3, row4)
	wform.SubmitText = "Save"
	wform.OnSubmit = func() {
		name := name.Text
		phone := mobile.Text
		email_id := email.Text
		address := address.Text
		myData := fmt.Sprintf(`%s %s %s %s`, name, phone, email_id, address)
		dialog.NewInformation("Success!", myData, Window).Show()

	}
	wform.OnCancel = func() {
		Window.Close()
	}

	Window.SetContent(container.NewVBox(
		title,
		wform,
	))
	Window.ShowAndRun()

}
