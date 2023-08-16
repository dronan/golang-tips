package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	XType int = iota
	OtherType
)

type XResultType struct {
	// some fields
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type OtherResultType struct {
	// some fields
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type ApiJob struct {
	ctx   context.Context
	url   string
	gType int
}

func (api *ApiJob) genericFunction() (interface{}, []string, error) {
	var err error
	var r interface{}

	token := `12345`

	if api.gType == XType {
		r, err = get[[]XResultType](api.ctx, api.url)
	}

	if api.gType == OtherType {
		r, err = get[[]OtherResultType](api.ctx, api.url, token)
	}

	if err != nil {
		return nil, nil, err
	}

	return r, nil, nil
}

func get[T any](ctx context.Context, url string, bearerToken ...string) (*T, error) {
	client := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	if len(bearerToken) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", bearerToken[0]))
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code '%d' is different from '%d'", res.StatusCode, http.StatusOK)
	}

	var result T
	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &result, nil

}
