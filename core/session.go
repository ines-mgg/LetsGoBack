// Package core provides the implementation of a lightweight HTTP router
// with support for dynamic routes, middleware, and route grouping.
// This file, session.go, implements session management functionality,
// including session creation, retrieval, and clearing, using cookies and an in-memory store.
package core

import (
	"net/http"
	"sync"
	"time"
)

var (
	sessions   = make(map[string]string)
	mu         sync.RWMutex
	cookieName = "auth_token"
)

// func generateSessionID() string {
// 	b := make([]byte, 32)
// 	rand.Read(b)
// 	return base64.URLEncoding.EncodeToString(b)
// }

func GetUserID(req *http.Request) (string, bool) {
	cookie, err := req.Cookie(cookieName)
	if err != nil {
		return "", false
	}

	mu.RLock()
	userID, ok := sessions[cookie.Value]
	mu.RUnlock()

	return userID, ok
}

func SetSession(w http.ResponseWriter, token string) {
	// sessionID := generateSessionID()

	// mu.Lock()
	// sessions[sessionID] = userID
	// mu.Unlock()

	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w, cookie)
}

func ClearSession(w http.ResponseWriter) {
	cookie := &http.Cookie{
		Name:     cookieName,
		Value:    "",
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
	}
	http.SetCookie(w, cookie)
}
