package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

type ConcurrentWork interface {
	work()
}

func main() {
	app := &cli.App{
		Name: "Concurrency pattern benchmark.",
		Commands: []*cli.Command{
			{
				Name:  "pool",
				Usage: "Pool pattern.",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:  "workers",
						Value: 4,
						Usage: "Number of workers.",
					},
					&cli.IntFlag{
						Name:  "tasks",
						Value: 1000,
						Usage: "Number of tasks.",
					},
				},
				Action: func(cCtx *cli.Context) error {
					c := PoolWork{cCtx.Int("workers"), cCtx.Int("tasks")}
					s := fmt.Sprintf("Running pool pattern with %d workers and %d tasks.", c.w, c.t)
					fmt.Println(s)

					c.work()
					return nil
				},
			},
			{
				Name:  "mutex",
				Usage: "Mutex pattern.",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "mutex",
						Value: false,
						Usage: "With/out mutex.",
					},
				},
				Action: func(cCtx *cli.Context) error {
					c := MutexWork{cCtx.Bool("mutex")}
					s := fmt.Sprintf("Running mutex pattern with mutex set to %t.", c.m)
					fmt.Println(s)

					c.work()
					return nil
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
