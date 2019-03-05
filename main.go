package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"log"
	"os/exec"
)

var message = ""

const SEPARETOR1 = "===========================================================================================================================\n"
const SEPARETOR2 = "\n===========================================================================================================================\n"

func main() {
	err := ui.Main(func() {
		buttonPostgreSQLStart := ui.NewButton("PostgreSQL - Start")
		buttonPostgreSQLStop := ui.NewButton("PostgreSQL - Stop")

		logger := ui.NewLabel("")

		box := ui.NewVerticalBox()
		box.Append(buttonPostgreSQLStart, false)
		box.Append(buttonPostgreSQLStop, false)

		box.Append(logger, false)

		window := ui.NewWindow("Gerenciador de Aplicativos", 1024, 768, false)
		window.SetChild(box)

		buttonPostgreSQLStart.OnClicked(func(*ui.Button) {
			cmd := exec.Command("pg_ctl", "-D", "/usr/local/var/postgres", "start")
			err := cmd.Run()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Start no PostgreSQL!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint(psAux(), SEPARETOR1)
			logger.SetText(message)
		})
		buttonPostgreSQLStop.OnClicked(func(*ui.Button) {
			cmd := exec.Command("pg_ctl", "-D", "/usr/local/var/postgres", "stop")
			err := cmd.Run()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Stop no PostgreSQL!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint(psAux(), SEPARETOR1)
			logger.SetText(message)
		})

		window.OnClosing(func(*ui.Window) bool {
			ui.Quit()
			return true
		})
		window.Show()
	})
	if err != nil {
		panic(err)
	}
}

func psAux() string {
	retorno, err := exec.Command("bash", "-c", "ps aux | grep postgres").Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(retorno)
}
