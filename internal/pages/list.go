package pages

import (
	"gostorefront/pkg/router"
	"gostorefront/pkg/view"
)

type ListViewData struct {
	Title string
}

func List() (router.Response, error) {
	data := ListViewData{Title: "Product list"}
	body, err := view.Load("list", data)
	if err != nil {
		return router.EmptyResponse(), nil
	}

	return router.Response{Body: body, StatusCode: 200}, nil
}
