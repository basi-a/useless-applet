package main

import (
	"log"
	"os/exec"

	"github.com/basi-a/useless-applet/config"
	"github.com/basi-a/useless-applet/fileserver"
	"github.com/basi-a/useless-applet/icon"
	"github.com/getlantern/systray"
	"github.com/skratchdot/open-golang/open"
)

func main()  {
	systray.Run(onReady, onExit)	
}

func onExit()  {
	log.Println("app exited")
}

func onReady() {
	var fileServercmd *exec.Cmd

	systray.SetTemplateIcon(icon.Data, icon.Data)
	mChecked := systray.AddMenuItemCheckbox("File Server", "Check Me", false)
	systray.AddSeparator()
	searchEngineUrl := systray.AddMenuItem("Search Something", "search")
	translatorUrl := systray.AddMenuItem("Translator", "Go to translator")
	mUrl := systray.AddMenuItem("My Blog", "Go to my blog")
	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	// We can manipulate the systray in other goroutines
	go func() {
		for {
			select {
			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					mChecked.Uncheck()
					mChecked.SetTitle("File Server")
					//stop the file server
					fileserver.FileServerStop(fileServercmd)
				} else {
					mChecked.Check()
					mChecked.SetTitle("File Server")
					//start file server
					fileServercmd = fileserver.FileServerStart()
				}
			case <-mQuit.ClickedCh:
				systray.Quit()
				log.Println("Quit now...")
				return
			case <-searchEngineUrl.ClickedCh:
				open.Run(config.SearchEngine)
			case <-translatorUrl.ClickedCh:
				open.Run(config.TranslatorUrl)
			case <-mUrl.ClickedCh:
				open.Run(config.MyBlog)
			}
		}
	}()
}
