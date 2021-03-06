package play

import (
	"net/http"
	"time"
)

type Context struct {
	RequestTime time.Time
	ActionName  string
	Render      string
	Input       *Input
	Output      Output

	HttpRequest  *http.Request
	HttpResponse http.ResponseWriter
	Session      Session

	doneFlag string
}

func (ctx *Context) Done(doneFlag string) error {
	ctx.doneFlag = doneFlag
	return nil
}

func NewContextWithInput(input *Input) *Context {
	ctx := &Context{Input: input, Output: &playKvOutput{}, RequestTime: time.Now()}
	return ctx
}
