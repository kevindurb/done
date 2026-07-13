package layouts

import (
	g "maragu.dev/gomponents"
	c "maragu.dev/gomponents/components"
	h "maragu.dev/gomponents/html"
)

func Layout(children ...g.Node) g.Node {
	return c.HTML5(c.HTML5Props{
		Title:    "Done",
		Language: "en",
		Body: []g.Node{
			h.Main(
				children...,
			),
		},
	})
}
