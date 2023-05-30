package blockbucket

import "github.com/GreenMan-Network/Go-Utils/pkg/queue/fifo"

type BlockBucket struct {
	fifo *fifo.Fifo // The underlying FIFO queue
}
