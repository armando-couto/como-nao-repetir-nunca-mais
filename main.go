package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"log"
	"os/exec"
)

var message = ""

const SEPARETOR1 = "=========================================================================================================================\n"
const SEPARETOR2 = "\n=========================================================================================================================\n"

func main() {
	err := ui.Main(func() {
		buttonPostgreSQLStart := ui.NewButton("PostgreSQL - Start")
		buttonPostgreSQLStop := ui.NewButton("PostgreSQL - Stop")
		buttonRedisStart := ui.NewButton("Redis - Start")
		buttonMongoDbStart := ui.NewButton("MongoDB - Start")
		buttonApacheStart := ui.NewButton("Apache - Start")
		buttonApacheStop := ui.NewButton("Apache - Stop")
		buttonMySQLStart := ui.NewButton("MySQL - Start")
		buttonMySQLStop := ui.NewButton("MySQL - Stop")

		password := ui.NewPasswordEntry()

		logger := ui.NewMultilineEntry()
		logger.SetReadOnly(true) // Só leitura

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Digite sua senha de sudo:"), false)
		box.Append(password, false)
		box.Append(buttonPostgreSQLStart, false)
		box.Append(buttonPostgreSQLStop, false)
		box.Append(buttonRedisStart, false)
		box.Append(buttonMongoDbStart, false)
		box.Append(buttonApacheStart, false)
		box.Append(buttonApacheStop, false)
		box.Append(buttonMySQLStart, false)
		box.Append(buttonMySQLStop, false)

		box.Append(logger, false)

		window := ui.NewWindow("Gerenciador de Aplicativos", 1024, 768, false)
		window.SetChild(box)

		buttonPostgreSQLStart.OnClicked(func(*ui.Button) {
			err := exec.Command("pg_ctl", "-D", "/usr/local/var/postgres", "start").Run()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Start no PostgreSQL!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Start no PostgreSQL!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep postgres"), SEPARETOR1)
			logger.SetText(message)
		})
		buttonPostgreSQLStop.OnClicked(func(*ui.Button) {
			err := exec.Command("pg_ctl", "-D", "/usr/local/var/postgres", "stop").Run()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Stop no PostgreSQL!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Stop no PostgreSQL!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep postgres"), SEPARETOR1)
			logger.SetText(message)
		})

		buttonRedisStart.OnClicked(func(*ui.Button) {
			err := exec.Command("bash", "-c", "redis-server /usr/local/etc/redis.conf &").Start()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Start no Redis!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Start no Redis!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep redis-server"), SEPARETOR1)
			logger.SetText(message)
		})

		buttonMongoDbStart.OnClicked(func(*ui.Button) {
			err := exec.Command("bash", "-c", "mongod").Start()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Start no MongoDB!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Start no MongoDB!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep mongod"), SEPARETOR1)
			logger.SetText(message)
		})

		buttonApacheStart.OnClicked(func(*ui.Button) {
			cmd := exec.Command("sudo", "true")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}

			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo -S /usr/sbin/apachectl start")).Output()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Start no Apache!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Start no Apache!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep apachectl"), SEPARETOR1)
			logger.SetText(message)
		})
		buttonApacheStop.OnClicked(func(*ui.Button) {
			cmd := exec.Command("sudo", "true")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}

			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo -S /usr/sbin/apachectl stop")).Output()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Stop no Apache!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Stop no Apache!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep apachectl"), SEPARETOR1)
			logger.SetText(message)
		})

		buttonMySQLStart.OnClicked(func(*ui.Button) {
			cmd := exec.Command("sudo", "true")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}

			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo launchctl load -w /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist")).Output()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Start no MySQL!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Start no MySQL!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep mysql"), SEPARETOR1)
			logger.SetText(message)
		})
		buttonMySQLStop.OnClicked(func(*ui.Button) {
			cmd := exec.Command("sudo", "true")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}

			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo launchctl unload -w /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist")).Output()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Stop no MySQL!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Stop no MySQL!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep mysql"), SEPARETOR1)
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

func psAux(command string) string {
	retorno, err := exec.Command("bash", "-c", command).Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(retorno)
}
