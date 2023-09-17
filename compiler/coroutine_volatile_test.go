//go:build !durable

package compiler

import (
	"testing"

	"github.com/stealthrocket/coroutine"
)

func assertSerializable[R, S any](t *testing.T, g coroutine.Coroutine[R, S]) coroutine.Coroutine[R, S] {
	// Not serializable unless the "durable" build tag is present.
	return g
}