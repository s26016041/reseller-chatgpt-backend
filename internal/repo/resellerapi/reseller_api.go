package resellerapi

import (
	"net/http"
)

type Repo struct {
	client *http.Client
}

func NewRepo() *Repo {
	return &Repo{
		client: &http.Client{},
	}
}
