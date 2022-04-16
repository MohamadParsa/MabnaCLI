package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"./postgres"
	"github.com/urfave/cli"
)

var app = cli.NewApp()

func main() {
	configApp()
	app.Run(os.Args)
}

func configApp() {
	app := cli.NewApp()
	app.Name = "Insert Random Date"
	app.Usage = "to insert random date into trade table."
	app.Version = "1.0.0"
	configAppAction()
}

func configAppAction() {
	app.Action = func(c *cli.Context) error {

		if c.Args().Len() != 1 {
			err := "Error: You should set a parameter as cout of random rows.( for example type: main 10 )"
			println(err)
			return errors.New(err)
		}

		count, err := strconv.Atoi(c.Args().First())
		if err != nil {
			err := "Error: You should set a number as cout of random rows.( for example type: main 10 )"
			println(err)
			return errors.New(err)
		}

		channel := runInsertCommand(count)
		waitForChannel(channel, count)

		return nil
	}
}

func runInsertCommand(count int) chan string {
	channel := make(chan string, count)

	var i int
	for i = 1; i <= count; i++ {
		go func(i int) {
			pg, err := postgres.InitializeDatabase("postgresql://postgres:Abc1234@localhost:5432/postgres?sslmode=disable")
			if err != nil {
				log.Println("error in connecting to the database")
				close(channel)
			} else {
				err := pg.InsertRandomDataIntoTrade()
				result := "- row number " + strconv.Itoa(i) + " was inserted"
				if err != nil {
					result = "- row number " + strconv.Itoa(i) + " has an error"
				}
				channel <- result
			}
		}(i)
	}
	return channel
}
func waitForChannel(channel chan string, count int) {
	for {
		select {
		case message, status := <-channel:
			if status {
				log.Println(message)
				count--
				if count == 0 {
					close(channel)
					return
				}
			} else {
				fmt.Println("---------------------------------")
				return
			}

		}
	}
}
