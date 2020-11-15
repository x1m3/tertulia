package ulid

import (
	ulidx "github.com/oklog/ulid"
	"math/rand"
	"time"
)

type ID struct {
	ulidx.ULID
}

func New() ID {
	return ID{ULID: ulidx.MustNew(ulidx.Now(), rand.New(rand.NewSource(time.Now().UnixNano())))}
}
