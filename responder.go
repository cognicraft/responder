package responder

import (
	"github.com/cognicraft/hyper"
	"github.com/cognicraft/icu"
)

func New(translate icu.TranslatorFunc) *Responder {
	return &Responder{
		translate: translate,
	}
}

type Responder struct {
	translate icu.TranslatorFunc
}

func (r *Responder) Translate(key string, ps ...icu.Parameter) string {
	return r.translate(key, ps...)
}

func (r *Responder) Created(msg string, location string) hyper.Item {
	res := hyper.Item{
		Label: msg,
		Type:  "response",
		Properties: hyper.Properties{
			{
				Label: r.translate("response:message:label"),
				Name:  "message",
				Value: msg,
			},
		},
		Links: hyper.Links{
			{
				Label: r.translate("details:label"),
				Rel:   hyper.RelDetails,
				Href:  location,
			},
		},
	}
	return res
}

func (r *Responder) Changed(msg string) hyper.Item {
	res := hyper.Item{
		Label: msg,
		Type:  "response",
		Properties: hyper.Properties{
			{
				Label: r.translate("response:message:label"),
				Name:  "message",
				Value: msg,
			},
		},
	}
	return res
}

func (r *Responder) Deleted(msg string) hyper.Item {
	res := hyper.Item{
		Label: msg,
		Type:  "response",
		Properties: hyper.Properties{
			{
				Label: r.translate("response:message:label"),
				Name:  "message",
				Value: msg,
			},
		},
	}
	return res
}

func (r *Responder) Error(msg string, errs ...error) hyper.Item {
	res := hyper.Item{
		Label: msg,
		Type:  "response",
		Properties: hyper.Properties{
			{
				Label: r.translate("response:message:label"),
				Name:  "message",
				Value: msg,
			},
		},
	}
	for _, err := range errs {
		e := hyper.Error{Message: err.Error()}
		if errC, ok := err.(errorCoder); ok {
			e.Code = errC.Code()
		}
		res.Errors = append(res.Errors, e)
	}
	return res
}

type errorCoder interface {
	Code() string
}
