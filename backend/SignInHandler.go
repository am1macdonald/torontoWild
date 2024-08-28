package main

import (
	"bytes"
	"net/http"

	"github.com/am1macdonald/torontoWild/internal/magiclink"
)

func extractEmail(r *http.Request) (string, error) {
	err := r.ParseForm()
	if err != nil {
		return "", err
	}

	email := r.Form.Get("email")
	if email == "" {
		return "", err
	}
	return email, nil
}

func (cfg *apiConfig) HandleSignIn(w http.ResponseWriter, r *http.Request) {
	// send a magic link to the provided email address
	email, err := extractEmail(r)
	if err != nil {
		errorResponse(err, 400, w)
		return
	}

	link, err := magiclink.CreateMagicLink(email, *cfg.valKey, r.Context())
	if err != nil {
		errorResponse(err, 500, w)
		return
	}

	var tpl bytes.Buffer

	err = cfg.templates.Render(&tpl, "magiclink", struct{ MagicLink string }{MagicLink: link}, r.Context())
	if err != nil {
		errorResponse(err, 500, w)
		return
	}

	err = cfg.mailer.Send(email,
		"torontoWild MagicLink",
		"", tpl.String())
	if err != nil {
		errorResponse(err, 500, w)
		return
	}

	w.WriteHeader(200)
}

func (cfg *apiConfig) HandleSignOut(w http.ResponseWriter, r *http.Request) {}
