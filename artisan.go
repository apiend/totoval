package main

import (
	"log"
	"os"

	"github.com/urfave/cli"

	"github.com/totoval/framework/cache"
	"github.com/totoval/framework/cmd"
	"github.com/totoval/framework/cmd/migration"
	"github.com/totoval/framework/database"
	"github.com/totoval/framework/helpers/m"
	"github.com/totoval/framework/queue"
	"totoval/app/console/commands"

	"totoval/app/events"
	"totoval/app/jobs"
	"totoval/app/listeners"
	"totoval/config"
	"totoval/resources/lang"
)

func init() {
	config.Initialize()
	cache.Initialize()
	database.Initialize()
	m.Initialize()
	lang.Initialize() // an translation must contains resources/lang/xx.json file (then a resources/lang/validation_translator/xx.go)
	queue.Initialize()
	jobs.Initialize()
	events.Initialize()
	listeners.Initialize()

	migration.Initialize()
	commands.Initialize()
}

func main() {
	app := cli.NewApp()
	app.Name = "artisan"
	app.Usage = "Let's work like an artisan"
	app.Version = "0.4.6"

	app.Commands = cmd.List()

	app.Action = func(c *cli.Context) error {
		cmd.Println(cmd.CODE_INFO, "COMMANDS:")
		for _, cate := range app.Categories() {
			categoryName := cate.Name
			if categoryName == "" {
				categoryName = "kernel"
			}
			cmd.Println(cmd.CODE_WARNING, "    "+categoryName+":")

			for _, cmds := range cate.Commands {
				cmd.Println(cmd.CODE_SUCCESS, "        "+cmds.Name+"    "+cmd.Sprintf(cmd.CODE_WARNING, "%s", cmds.Usage))
			}
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
