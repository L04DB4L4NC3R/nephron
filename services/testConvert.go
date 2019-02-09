package services

import (
	"code.sajari.com/docconv"
	"github.com/angadsharma1016/nephron/model"
)

func ConvertToText(path string, c chan model.StringReturn) {

	res, err := docconv.ConvertPath(path)
	if err != nil {
		c <- model.StringReturn{"error while converting", err}
		return
	}
	c <- model.StringReturn{res.Body, nil}

}
