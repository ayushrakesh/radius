// ------------------------------------------------------------
// Copyright (c) Microsoft Corporation.
// Licensed under the MIT License.
// ------------------------------------------------------------

package hosting

import (
	"context"
	"fmt"
	"sync"
)

// noCopy may be embedded into structs which must not be copied
// after the first use.
//
// See https://golang.org/issues/8005#issuecomment-190753527
// for details.
type noCopy struct{} //nolint:golint,unused

// Lock is a no-op used by -copylocks checker from `go vet`.
func (*noCopy) Lock()   {} //nolint:golint,unused
func (*noCopy) Unlock() {} //nolint:golint,unused

type AsyncValue struct {
	noCopy noCopy //nolint

	Cond  *sync.Cond
	Value interface{}
	Err   error
}

func NewAsyncValue() *AsyncValue {
	return &AsyncValue{Cond: &sync.Cond{L: &sync.Mutex{}}}
}

func (a *AsyncValue) Get(ctx context.Context) (interface{}, error) {
	type result struct {
		Value interface{}
		Err   error
	}

	initialized := make(chan result, 1)
	go func() {
		a.Cond.L.Lock()

		for {
			if a.Value != nil || a.Err != nil {
				break
			}

			// Not ready to proceed, wait to be woken up
			a.Cond.Wait()
		}

		a.Cond.L.Unlock()
		initialized <- result{Value: a.Value, Err: a.Err}
		close(initialized)
	}()

	select {
	case <-ctx.Done():
		return nil, fmt.Errorf("failed to retrieve value: %w", ctx.Err())

	case result := <-initialized:
		return result.Value, result.Err
	}
}

func (a *AsyncValue) Put(value interface{}) {
	a.Cond.L.Lock()
	a.Value = value
	a.Cond.L.Unlock()
	a.Cond.Broadcast()
}

func (a *AsyncValue) PutErr(err error) {
	a.Cond.L.Lock()
	a.Err = err
	a.Cond.L.Unlock()
	a.Cond.Broadcast()
}