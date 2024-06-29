package api

import (
	"gvb_server/api/advert_api"
	"gvb_server/api/images_api"
	"gvb_server/api/menu_api"
	"gvb_server/api/setting_api"
	"gvb_server/api/user_api"
)

type ApiGroup struct {
	SettingsApi setting_api.SettingsApi
	ImagesApi   images_api.ImagesApi
	AdvertApi   advert_api.AdvertApi
	MenuAPi     menu_api.MenuAPi
	LoginApi    user_api.UserApi
}

var ApiGroupApp = new(ApiGroup)
