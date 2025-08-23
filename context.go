package stdx

import (
	"context"
	"time"
)

type Context struct {
	context.Context
	cancel context.CancelFunc
}

func NewContext(parent context.Context) *Context {
	ctx, cancel := context.WithCancel(parent)
	return &Context{
		Context: ctx,
		cancel:  cancel,
	}
}

func NewContextWithTimeout(parent context.Context, timeout time.Duration, cause ...error) *Context {
	var err error
	if len(cause) > 0 {
		err = cause[0]
	}
	ctx, cancel := context.WithTimeoutCause(parent, timeout, err)
	return &Context{
		Context: ctx,
		cancel:  cancel,
	}
}

func NewContextWithDeadline(parent context.Context, deadline time.Time, cause ...error) *Context {
	var err error
	if len(cause) > 0 {
		err = cause[0]
	}
	ctx, cancel := context.WithDeadlineCause(parent, deadline, err)
	return &Context{
		Context: ctx,
		cancel:  cancel,
	}
}

func (c *Context) Wait() {
	<-c.Done()
}

func (c *Context) Cancel() {
	c.cancel()
}

func (c *Context) Cancelled() bool {
	select {
	case <-c.Done():
		return true
	default:
		return false
	}
}

func (c *Context) Cause() error {
	return context.Cause(c.Context)
}
