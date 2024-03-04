package model

type HttpClient interface {
	Get(url string) (string, error)
}
