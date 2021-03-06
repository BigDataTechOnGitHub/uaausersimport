package token

import "encoding/json"
import "fmt"
import "github.com/pivotalservices/gtils/http"

import "github.com/pivotalservices/uaausersimport/functions"
import "io/ioutil"
import . "net/http"

type Token struct {
	AccessToken string `json:"access_token"`
}

var NewGateway = func() http.HttpGateway {
	return http.NewHttpGateway()
}

var GetToken functions.TokenFunc = func(info *functions.Info) (token string, err error) {
	fmt.Println("Getting token.............")
	entity := http.HttpRequestEntity{
		Url:      fmt.Sprintf("%s/oauth/token?grant_type=client_credentials", info.Uaaurl),
		Username: info.Clientid,
		Password: info.Secret,
	}
	httpGateway := NewGateway()
	request := httpGateway.Post(entity, nil)
	response, err := request()
	if err != nil {
		return
	}
	return parse(response)
}

func parse(response *Response) (tokenString string, err error) {
	body := response.Body
	defer body.Close()
	token := &Token{}
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &token)
	if err != nil {
		return
	}
	tokenString = token.AccessToken
	return
}
