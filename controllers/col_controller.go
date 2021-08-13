package controllers

import (
	"github.com/ideaspaper/mer/services"
	"github.com/ideaspaper/mer/views"
	"github.com/spf13/viper"
)

type IColController interface {
	Search(string)
}

type colController struct {
	service services.IColService
	view    views.IColView
}

func NewColController(service services.IColService, view views.IColView) IColController {
	return &colController{
		service: service,
		view:    view,
	}
}

func (cc *colController) Search(keyword string) {
	result, didYouMean, err := cc.service.Search(keyword, viper.GetString("API_KEY"))
	if err != nil {
		cc.view.DisplayError(err)
	} else if len(didYouMean) > 0 {
		cc.view.DidYouMean(didYouMean)
	} else if len(result) > 0 {
		cc.view.Search(result)
	} else {
		cc.view.NoResult()
	}
}
