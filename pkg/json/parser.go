package json

import (
	"encoding/json"
	"log"
	"regexp"
	"strings"
)

type Token struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
}

func UnmarshallStringToTokenStruct(tokenString string, token *Token)  (*Token, error) {


	log.Printf("before %v", tokenString)
	//err := json.Unmarshal([]byte(tokenString),&token)
	//if err != nil {
	//	log.Fatalf("%v",err)
	//	return nil, err
	//}
	rgx := regexp.MustCompile(`(\\n)|\\`)
	result := rgx.ReplaceAllString(strings.TrimPrefix(strings.TrimSuffix(tokenString,"\""),"\""),"")
	log.Printf("after %v", result)
	err := json.Unmarshal([]byte(result),&token)
	if err != nil {
		log.Fatalf("%v",err)
		return nil, err
	}
	log.Printf("%v", result)
	return token, nil
}
