package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/RichardKnop/machinery/v1"
	"github.com/RichardKnop/machinery/v1/config"
	"github.com/RichardKnop/machinery/v1/tasks"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	cfg := &config.Config{
		Broker:        "redis://localhost:6379",
		ResultBackend: "redis://localhost:6379",
	}
	server, err := machinery.NewServer(cfg)
	if err != nil {
		log.Fatal(err)
	}
	// tasks := map[string]interface{}{
	//     "task": task,
	// }
	// server.RegisterTasks(tasks)
	server.RegisterTask("task", task)

	worker := server.NewWorker("worker1", 10)

	errCh := make(chan error)
	worker.LaunchAsync(errCh)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	sig := tasks.Signature{
		Name: "task",
		Args: []tasks.Arg{
			{
				Type:  "string",
				Value: "hello please",
			},
		},
	}
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	asyncResult, err := server.SendTaskWithContext(ctx, &sig)
	if err != nil {
		log.Fatal(err)
	}
	spew.Dump("state before ", asyncResult.GetState())
	// rs, err := asyncResult.Get(10 * time.Millisecond)
	rs, err := asyncResult.GetWithTimeout(1000*time.Millisecond, 10*time.Millisecond)
	spew.Dump("state after ", asyncResult.GetState())

	for _, r := range rs {
		spew.Dump("results", r)
	}
	spew.Dump("results: ", tasks.HumanReadableResults(rs))

	worker.Quit()
	err = <-errCh
	if err != nil {
		log.Fatal(err)
	}
}

func task(ctx context.Context, msg string) (string, error) {
	fmt.Println("received ", msg)

	return "returning from task", nil
}
