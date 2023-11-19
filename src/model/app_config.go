package model

import (
	"encoding/json"
)

const APP_CONFIG_FIELD_PUBLIC = "public"
const APP_CONFIG_FIELD_WATER_MARK = "waterMark"
const APP_CONFIG_FIELD_DESCRIPTION = "description"

type AppConfig struct {
	Public                 bool   `json:"public"` // switch for public app (which can view by anonymous user)
	WaterMark              bool   `json:"waterMark"`
	Description            string `json:"description"`
	PublishedToMarketplace bool   `json:"publishedToMarketplace"`
	Cover                  string `json:"cover"`
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		Public:      false,
		WaterMark:   true,
		Description: "",
	}
}

func (appConfig *AppConfig) ExportToJSONString() string {
	r, _ := json.Marshal(appConfig)
	return string(r)
}

func (appConfig *AppConfig) IsPublic() bool {
	return appConfig.Public
}

func (appConfig *AppConfig) EnableWaterMark() {
	appConfig.WaterMark = true
}

func (appConfig *AppConfig) DisableWaterMark() {
	appConfig.WaterMark = false
}

func (appConfig *AppConfig) SetublishedToMarketplace() {
	appConfig.PublishedToMarketplace = true
}

func (appConfig *AppConfig) SetNotPublishedToMarketplace() {
	appConfig.PublishedToMarketplace = false
}

func (appConfig *AppConfig) SetCover(cover string) {
	appConfig.Cover = cover
}
func NewAppConfigByDefault() *AppConfig {
	return &AppConfig{
		Public:      false,
		WaterMark:   true,
		Description: "",
	}
}
