package httpservice

import (
	"io"
	"net/http"
)

type Response struct {
	StatusCode int
	Header     map[string][]string
	Body       io.ReadCloser
}

//go:generate mockgen --build_flags=--mod=mod -destination=../mocks/httpService.go -package=mocks sk8.town/backside/httpservice HttpService
type HttpService interface {
	Get(url string) (resp *Response, err error)
}

type DefaultHttpService struct{}

func (s DefaultHttpService) Get(url string) (resp *Response, err error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	return &Response{
		StatusCode: response.StatusCode,
		Header:     response.Header,
		Body:       response.Body,
	}, nil
}

func NewHttpService() DefaultHttpService {
	return DefaultHttpService{}
}
