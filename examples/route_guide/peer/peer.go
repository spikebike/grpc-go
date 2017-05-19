
import (
		  "fmt"

		  "golang.org/x/net/context"
)

type key int

const peerIDKey key=0

func NewContext(ctx context.Context, peerID int) context.Context {
	return context.WithValue(ctx, peerIDKey, peerID)
}

