package grifts

import (
	"maratona_full_cycle_3_desafio_2/actions"

	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
