package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/yeboahd24/recaptcha-demo/internal/config"
	"github.com/yeboahd24/recaptcha-demo/internal/recaptcha"
)

func RegisterRoutes(mux *http.ServeMux, cfg *config.Config) {
	// Parse templates once during initialization
	tmpl, err := template.ParseFiles(
		filepath.Join("static", "index.html"),
		filepath.Join("static", "success.html"),
	)
	if err != nil {
		panic(fmt.Sprintf("Failed to parse templates: %v", err))
	}

	mux.HandleFunc("/", serveForm(cfg, tmpl))
	mux.HandleFunc("/verify", verifyRecaptcha(cfg, tmpl))
}

func serveForm(cfg *config.Config, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		// Render index.html with SiteKey and optional error message
		data := struct {
			SiteKey string
			Error   string
		}{
			SiteKey: cfg.RecaptchaSiteKey,
			Error:   r.URL.Query().Get("error"), // Get error from query param if present
		}
		if err := tmpl.ExecuteTemplate(w, "index.html", data); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}

func verifyRecaptcha(cfg *config.Config, tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Redirect(w, r, "/?error=Failed+to+parse+form", http.StatusSeeOther)
			return
		}

		token := r.FormValue("g-recaptcha-response")
		if token == "" {
			http.Redirect(w, r, "/?error=reCAPTCHA+token+missing", http.StatusSeeOther)
			return
		}

		result, err := recaptcha.Verify(token, cfg.RecaptchaSecret)
		if err != nil {
			http.Redirect(w, r, fmt.Sprintf("/?error=reCAPTCHA+verification+failed:+%v", err), http.StatusSeeOther)
			return
		}

		if !result.Success || result.Score < 0.5 || result.Action != "submit" {
			errMsg := fmt.Sprintf("reCAPTCHA+verification+failed:+Score=%.2f,+Action=%s", result.Score, result.Action)
			http.Redirect(w, r, "/?error="+errMsg, http.StatusSeeOther)
			return
		}

		// Render success.html on successful verification
		if err := tmpl.ExecuteTemplate(w, "success.html", nil); err != nil {
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
		}
	}
}
