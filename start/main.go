package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"

	"app"
)

func main() {
	// Create the client object just once per process
	c, err := client.NewClient(client.Options{
		HostPort: app.HostPort,
	})
	if err != nil {
		panic(fmt.Sprintf("unable to create Temporal client; err: %v", err))
	}
	defer c.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "OK!")
	})
	http.HandleFunc("/trigger-workflow", func(w http.ResponseWriter, r *http.Request) {
		options := client.StartWorkflowOptions{
			ID:        "transfer-money-workflow",
			TaskQueue: app.TransferMoneyTaskQueue,
		}
		transferDetails := app.TransferDetails{
			Amount:      54.99,
			FromAccount: "001-001",
			ToAccount:   "002-002",
			ReferenceID: uuid.New().String(),
		}
		we, err := c.ExecuteWorkflow(context.Background(), options, app.TransferMoney, transferDetails)
		if err != nil {
			log.Fatalln("error starting TransferMoney workflow", err)
		}
		_, err = fmt.Fprintf(w,
			"Transfer of $%f from account %s to account %s is processing. ReferenceID: %s\n\nWorkflowID: %s RunID: %s",
			transferDetails.Amount,
			transferDetails.FromAccount,
			transferDetails.ToAccount,
			transferDetails.ReferenceID,
			we.GetID(),
			we.GetRunID(),
		)
	})

	if err := http.ListenAndServe(":10000", nil); err != nil {
		panic(err)
	}
}
