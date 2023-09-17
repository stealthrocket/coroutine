//go:build durable

package compiler

import (
	"testing"

	"github.com/stealthrocket/coroutine"
)

func assertSerializable[R, S any](t *testing.T, g coroutine.Coroutine[R, S]) coroutine.Coroutine[R, S] {
	c := g.Context()
	b, err := c.MarshalAppend(nil)
	if err != nil {
		t.Fatal(err)
	}
	var reconstructed coroutine.Context[R, S]
	if n, err := reconstructed.Unmarshal(b); err != nil {
		t.Fatal(err)
	} else if n != len(b) {
		t.Fatal("invalid number of bytes read when reconstructing context")
	}
	e := c.Entrypoint
	*c = reconstructed
	// TODO: the context reconstruction needs to capture the
	// coroutine entry point.
	//
	// https://www.notion.so/stealthrocket/Durable-Coroutines-1487e78403804b5f871cf37275a55cc8?pvs=4#395d316dc79e432ca58dd59df9f561f0
	c.Entrypoint = e

	return g
}