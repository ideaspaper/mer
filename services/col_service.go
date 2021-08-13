package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/ideaspaper/mer/entities"
)

type IColService interface {
	Search(string, string) ([]entities.Col, []string, error)
}

type colService struct {
	url string
}

func NewColService(url string) IColService {
	return &colService{
		url: url,
	}
}

func (cs *colService) Search(keyword, apiKey string) ([]entities.Col, []string, error) {
	resp, err := http.Get(fmt.Sprintf("%s/%s?key=%s", cs.url, keyword, apiKey))
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	result := []entities.Col{}
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		resutDidYouMean := []string{}
		err = json.Unmarshal([]byte(body), &resutDidYouMean)
		if err != nil {
			return nil, nil, errors.New(string(body))
		}
		return nil, resutDidYouMean, nil
	}

	return result, nil, nil
}
