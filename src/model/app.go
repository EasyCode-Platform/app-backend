package model

import (
	"time"

	"github.com/gofrs/uuid"
)

type App struct {
	ID          int       `json:"id" 				gorm:"column:id;type:bigserial;primary_key;unique"`
	UID         uuid.UUID `json:"uid"   		    gorm:"column:uid;type:uuid;not null"`
	TeamID      int       `json:"teamID" 		    gorm:"column:team_id"`
	Name        string    `json:"name" 				gorm:"column:name;type:text"`
	ComponentId string    `json:"name" 				gorm:"column:coponent_id;type:text;notnull"`
	Config      string    `json:"config" 	        gorm:"column:config;type:jsonb"`
	CreatedAt   time.Time `json:"createdAt" 		gorm:"column:created_at;type:timestamp"`
	UpdatedAt   time.Time `json:"updatedAt" 		gorm:"column:updated_at;type:timestamp"`
}

func NewApp(appName string, teamID int, modifyUserID int, ComponentId string) *App {
	return &App{
		TeamID:      teamID,
		Name:        appName,
		ComponentId: ComponentId,
		Config:      NewAppConfig().ExportToJSONString(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}
}
