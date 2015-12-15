package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"jarvis/log"
	"net/http"
	"time"
)

type DiningMenu struct {
	Location string       `json:"Location"`
	Date     string       `json:"Date"`
	Meals    []DiningMeal `json:"Meals"`
}

type DiningMeal struct {
	Name   string `json:"Name"`
	Order  int    `json:"Order"`
	Type   string `json:"Type"`
	Status string `json:"Status"`
	Hours  struct {
		StartTime string `json:"StartTime"`
		EndTime   string `json:"EndTime"`
	} `json:"Hours"`
	Stations []DiningStation `json:"Stations"`
}

type DiningStation struct {
	Name  string       `json:"Name"`
	Items []DiningItem `json:"Items"`
}

type DiningItem struct {
	Id           string           `json:"ID"`
	Name         string           `json:"Name"`
	IsVegetarian bool             `json:"IsVegetarian"`
	Allergens    []DiningAllergen `json:"Allergens"`
}

type DiningAllergen struct {
	Name  string `json:"Name"`
	Value bool   `json:"Value"`
}

type Purdue struct{}

func (p Purdue) GetDiningMenu(location string, date time.Time) (*DiningMenu, error) {
	url := fmt.Sprintf("http://api.hfs.purdue.edu/menus/v2/locations/%v/%v-%v-%v", location, date.Year(), date.Month(), date.Day())
	res, err := http.Get(url)
	if log.Error(err) {
		return nil, errors.New("It seems like I'm having problems contacting the Dining Courts API.")
	}
	resB, err := ioutil.ReadAll(res.Body)
	if log.Error(err) {
		return nil, errors.New("It seems like I'm having problems contacting the Dining Courts API.")
	}
	menu := &DiningMenu{}
	err = json.Unmarshal(resB, menu)
	if log.Error(err) {
		return nil, errors.New("It seems like I'm having problems contacting the Dining Courts API.")
	}
	return menu, nil
}
