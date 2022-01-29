package redis

import "encoding/json"

var token = "{\n  \"access_token\": \"eyJhbGciOiJSUzI1NiIsImtpZCI6InB1YmxpYzo4MzQ3Zjc3OS1hMDIxLTRlMzEtYTQ4ZC1iNWU1NjdjMzg2ZmMiLCJ0eXAiOiJKV1QifQ.eyJhdWQiOltdLCJjbGllbnRfaWQiOiJhcHBJRDpOTURQVFJJQUxfa2ltX3Nlb2tfbnVhbmNlX2NvbV8yMDIxMDQyN1QwNDQzNTE1MjY3NDg6Z2VvOnVzOmNsaWVudE5hbWU6ZXZlbnRfbG9nX2NvbGxlY3RvciIsImV4cCI6MTYzMTYxMDI0NywiZXh0Ijp7fSwiaWF0IjoxNjMxNjA5MzQ3LCJpc3MiOiJodHRwczovL2F1dGguY3J0Lm51YW5jZS5jb20vIiwianRpIjoiMjM1NzM4ZWYtYmI4NS00OTJiLTg0ZTktM2M2NTk2ODUyNWVmIiwibmJmIjoxNjMxNjA5MzQ3LCJzY3AiOlsibG9nLndyaXRlIl0sInN1YiI6ImFwcElEOk5NRFBUUklBTF9raW1fc2Vva19udWFuY2VfY29tXzIwMjEwNDI3VDA0NDM1MTUyNjc0ODpnZW86dXM6Y2xpZW50TmFtZTpldmVudF9sb2dfY29sbGVjdG9yIn0.g3PTUacr9-v_Bhg2gqsjGJIP--zSe1fSyGadqkYurYmY9VwqEqcORtP3hgN0iuatufsLp586Yjcucl8D2agZ9aHOo_h0za0JcZQPSlgPgKWzkEbgb3eUzUxqingUb5CsbSnUtMS5NL3gm-fa9JAQQNtP7FjUXjr8jQaHL0biQwBMGQhczGyNTWsBuKkzECHHlMRT76nut0bpEz3M7f2UiQgAMH4YDYgshofIX6QZeDpNckfKAinuip88mV4Ga_f6jH23pUZFpCzD-O0UMeb0y7SmDMAE7XnFUHWy83xc33Ojv9AUT1qOpjxarUtMcPuwHB4H3tawO0inn2ZC9Ymh7CUVwWBK8e5iGesR0Yf49F_JuxEjhPvsPC5rER6s409z-kTLvYlSsTvj0FJk97Jw7-Bapez2P8n-UdBba3Csf6oi5Oymv1keJGm3GVpOrQshfJOr89GieOmrFXpgOBtjnYwYBML05ecCFA8r8v6aqkWTG8yqh_9z1lTIptBRyJ9bYos-kkaPkpsgVRkZgoOpxTyhT3rRsK9xLIPcFca_4opJt4ZBJ2gatOY1KeuC6HszUCSdul4tPjP04mCpNZiismLAlYcE2krnxPwmRsliakP6E1SmwVUilPM6ghtkGdkWIh4P-4dSVh08l1nmBkEmTfpip0misI9jJVVBUVP8ZkA\",\n  \"expires_in\": 899,\n  \"scope\": \"log.write\",\n  \"token_type\": \"bearer\"\n}"

type Token struct {
	AccessToken string `json:"access_token,omitempty"`
	ExpriresIn  int    `json:"expires_in,omitempty"`
	Scope       string `json:"scope,omitempty"`
	TokenType   string `json:"token_type,omitempty"`
}

func getToken() string {

	return token

}

func decodeToken(t *Token) error {
	err := json.Unmarshal([]byte(getToken()), &t)
	if err != nil {
		return err
	}
	return nil
}
