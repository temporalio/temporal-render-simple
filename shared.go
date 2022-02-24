package app

import "os"

const TransferMoneyTaskQueue = "TRANSFER_MONEY_TASK_QUEUE"

var HostPort = os.Getenv("TEMPORAL_CLUSTER_HOST") + ":7233"

type TransferDetails struct {
	Amount      float32
	FromAccount string
	ToAccount   string
	ReferenceID string
}
