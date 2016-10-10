package api

import (
	"context"
	"net/http"
)

// Index endpoint
func (a *API) Index(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	sendJSON(w, 200, map[string]string{
		"version":     a.version,
		"name":        "Netlify Commerce",
		"description": "Netlify Commerce is a flexible Ecommerce API for JAMStack sites",
	})
}
