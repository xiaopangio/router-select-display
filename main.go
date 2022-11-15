package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"log"
	"os"
)

type Config struct {
	app        *fyne.App
	mainWindow fyne.Window
	InfoLog    *log.Logger
	ErrorLog   *log.Logger
}

func main() {
	//create our app
	myApp := Config{}
	//create an app
	a := app.New()
	myApp.app = &a
	//create logger
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	//create a window
	window := a.NewWindow("router-select")
	myApp.mainWindow = window
	//makeUI for window
	myApp.makeUI()
	//resize a window
	myApp.mainWindow.Resize(fyne.NewSize(770, 410))
	myApp.mainWindow.SetMaster()
	myApp.mainWindow.SetFixedSize(true)
	//show and run window
	myApp.mainWindow.ShowAndRun()
}
