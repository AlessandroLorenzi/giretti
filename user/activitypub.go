package user

func (u *User) ToActivityPub() map[string]interface{} {
	return map[string]interface{}{
		"@context":          []string{"https://www.w3.org/ns/activitystreams", "https://w3id.org/security/v1"},
		"id":                "https://giretti.com/@" + u.Username,
		"type":              "Person",
		"preferredUsername": u.Username,
		"inbox":             "https://giretti.com/@" + u.Username + "/inbox",
		"outbox":            "https://giretti.com/@" + u.Username + "/outbox",
		"publicKey": map[string]interface{}{
			"id":           "https://giretti.com/@" + u.Username + "#main-key",
			"owner":        "https://giretti.com/@" + u.Username,
			"publicKeyPem": u.PublicKeyPem,
		},
	}
}
