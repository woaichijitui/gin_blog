package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/images_api"
	"gvb_server/api/setting_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
}

var ApiGroupApp = new(ApiGroup)
