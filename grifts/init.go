package grifts

import (
	"mta_hosting_optimiser/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
