package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/sejal-varu/buffalo_test/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
