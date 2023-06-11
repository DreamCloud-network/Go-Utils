package blockbucketutils

import "github.com/GreenMan-Network/Go-Utils/pkg/queueutils/fifoutils"

type BlockBucket struct {
	fifo *fifoutils.Fifo // The underlying FIFO queue
}
