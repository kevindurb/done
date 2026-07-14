package html

import (
	"fmt"
	g "maragu.dev/gomponents"
	h "maragu.dev/gomponents/html"
)

func Actionf(format string, a ...any) g.Node {
	return h.Action(fmt.Sprintf(format, a...))
}
