package session

import (
	"encoding/base64"
	"github.com/gorilla/sessions"
	"log"
	"web2/config"
)

var Store *sessions.CookieStore

func Init() {
	// Decode keys from base64
	hashKey, err := base64.StdEncoding.DecodeString(string(config.HashKey))
	if err != nil {
		log.Fatal("Failed to decode SESSION_HASH_KEY:", err)
	}
	blockKey, err := base64.StdEncoding.DecodeString(string(config.BlockKey))
	if err != nil {
		log.Fatal("Failed to decode SESSION_BLOCK_KEY:", err)
	}

	Store = sessions.NewCookieStore(hashKey, blockKey)
	Store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   config.SessionLife,
		HttpOnly: true,
		// Secure: true, // enable if using HTTPS
	}
}
