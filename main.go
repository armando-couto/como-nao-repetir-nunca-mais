package main

import (
	"fmt"
	"github.com/andlabs/ui"
	"log"
	"os/exec"
)

var message = ""
var logger = ui.NewMultilineEntry()

const SEPARETOR1 = "=========================================================================================================================\n"
const SEPARETOR2 = "\n=========================================================================================================================\n"

func main() {
	err := ui.Main(func() {
		buttonPostgreSQLStart := ui.NewButton("PostgreSQL - Start")
		buttonPostgreSQLStop := ui.NewButton("PostgreSQL - Stop")
		buttonRedisStart := ui.NewButton("Redis - Start")
		buttonRedisStop := ui.NewButton("Redis - Stop")
		buttonMongoDbStart := ui.NewButton("MongoDB - Start")
		buttonMongoDbStop := ui.NewButton("MongoDB - Stop")
		buttonApacheStart := ui.NewButton("Apache - Start")
		buttonApacheStop := ui.NewButton("Apache - Stop")
		buttonMySQLStart := ui.NewButton("MySQL - Start")
		buttonMySQLStop := ui.NewButton("MySQL - Stop")

		password := ui.NewPasswordEntry()

		logger.SetReadOnly(true) // Só leitura

		box := ui.NewVerticalBox()
		box.Append(ui.NewLabel("Digite sua senha de sudo:"), false)
		box.Append(password, false)
		box.Append(buttonPostgreSQLStart, false)
		box.Append(buttonPostgreSQLStop, false)
		box.Append(buttonRedisStart, false)
		box.Append(buttonRedisStop, false)
		box.Append(buttonMongoDbStart, false)
		box.Append(buttonMongoDbStop, false)
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

		buttonRedisStop.OnClicked(func(*ui.Button) {
			pid, err := exec.Command("bash", "-c", "pgrep redis-server").Output()
			err = exec.Command("bash", "-c", "kill -9 ", string(pid)).Start()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Stop no Redis!", SEPARETOR2)
				logger.SetText(message)
			}

			message += fmt.Sprint("Stop no Redis!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep redis-server"), SEPARETOR1)
			logger.SetText(message)
		})

		buttonMongoDbStart.OnClicked(func(*ui.Button) {
			cmd := exec.Command("sudo", "true")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo -S mongod &")).Output()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Start no MongoDB!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Start no MongoDB!", SEPARETOR2)
			logger.SetText(message)
			message += fmt.Sprint(psAux("ps aux | grep mongod"), SEPARETOR1)
			logger.SetText(message)
		})

		buttonMongoDbStop.OnClicked(func(*ui.Button) {
			cmd := exec.Command("sudo", "true")
			err := cmd.Start()
			if err != nil {
				log.Fatal(err)
			}
			pid, err := exec.Command("bash", "-c", "pgrep mongod").Output()
			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo -S kill -9 ", string(pid))).Output()
			if err != nil {
				message += fmt.Sprint("Erro ao dar Stop no MongoDB!", SEPARETOR2)
				logger.SetText(message)
			}
			message += fmt.Sprint("Stop no MongoDB!", SEPARETOR2)
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

			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo -S launchctl load -w /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist")).Output()
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

			_, err = exec.Command("bash", "-c", fmt.Sprint("echo '", password.Text(), "' | sudo -S launchctl unload -w /Library/LaunchDaemons/com.oracle.oss.mysql.mysqld.plist")).Output()
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
