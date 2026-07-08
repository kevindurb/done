package form

import (
	"net/http"

	"github.com/go-playground/form/v4"
	"github.com/go-playground/validator/v10"
)

type Parser struct {
	d *form.Decoder
	v *validator.Validate
}

func New() *Parser {
	return &Parser{
		form.NewDecoder(),
		validator.New(),
	}
}

func (p *Parser) Parse(dst any, r *http.Request) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	if err := p.d.Decode(dst, r.PostForm); err != nil {
		return err
	}

	if err := p.v.StructCtx(r.Context(), dst); err != nil {
		return err
	}

	return nil
}
