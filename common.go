package fineract

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/meson10/highbrow"
)

var (
	authenticationKey AuthenticationKey
)

type AuthenticationKey struct {
	Data string `json:"base64EncodedAuthenticationKey"`
}

func BasicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func (client *Client) MakeRequest(reqType, url string, payload interface{}, response interface{}) error {
	b, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		rawMessage := json.RawMessage([]byte(err.Error()))
		return &FineractError{ErrCodeSerialization, &rawMessage}
	}

	req, err := http.NewRequest(reqType, url, bytes.NewBuffer(b))
	if err != nil {
		log.Println(err)
		rawMessage := json.RawMessage([]byte(err.Error()))
		return &FineractError{ErrBadRequest, &rawMessage}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("fineract-platform-tenantid", "default")
	req.Header.Set("Authorization", "Basic "+BasicAuth(client.UserName, client.Password))

	var resp *http.Response
	errTry := highbrow.Try(5, func() error {
		resp, err = client.Option.Transport.Do(req)
		return err
	})
	if errTry != nil {
		rawMessage := json.RawMessage([]byte(errTry.Error()))
		return &FineractError{ErrCodeResponseTimeout, &rawMessage}
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	//log.Println(string(body))
	//log.Println(resp.StatusCode)
	//log.Println("------------------")
	if resp.StatusCode != 200 {
		rawMessage := json.RawMessage(body)
		return &FineractError{GetFineractStatusCode(resp.StatusCode), &rawMessage}
	}

	if err = json.Unmarshal(body, &response); err != nil {
		log.Println(err)
		rawMessage := json.RawMessage([]byte(err.Error()))
		return &FineractError{ErrCodeSerialization, &rawMessage}
	}

	return nil
}
