package auth

// CookieShape is the shape of the auth cookie.
type CookieShape struct {
	UserID        string `json:"user_id"`
	SessionID     string `json:"session_id"`
	SessionSecret string `json:"token_key"`
}
