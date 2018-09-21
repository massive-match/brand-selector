package internal

import (
	"encoding/json"
	"net/http"
)

type Request struct {
	Brand  string
	Header Header
}

type Header struct {
	AccessToken string
	Client      string
	BusinessId  string
	CompanyId   string
	UserType    string
}

type Response struct {
	Page []Page `json:"page"`
}

type Page struct {
	Code string `json:"code"`
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (request *Request) GetBrand() (Response, error) {
	var response Response
	url := "https://api-falabella-svl.buffetcloud.io/bff/api/v1/brands?filter[where][name]=" + request.Brand + "&locale=es"
	client := &http.Client{}
	if req, err := http.NewRequest("GET", url, nil); err != nil {
		return response, err
	} else {
		req.Header.Set("access-token", request.Header.AccessToken)
		req.Header.Set("client", request.Header.Client)
		req.Header.Set("locale", "es")
		req.Header.Set("X-Business-ID", request.Header.BusinessId)
		req.Header.Set("X-Company-Id", request.Header.CompanyId)
		req.Header.Set("X-UserType", request.Header.UserType)
		if resp, err := client.Do(req); err != nil {
			return response, err
		} else {
			if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
				return response, err
			} else {
				return response, nil
			}
		}
	}
}
