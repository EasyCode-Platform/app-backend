package storage

import (
	"testing"

	"github.com/EasyCode-Platform/app-backend/src/driver/mongodb"
	"github.com/EasyCode-Platform/app-backend/src/driver/postgres"
	"github.com/EasyCode-Platform/app-backend/src/model"
	"github.com/EasyCode-Platform/app-backend/src/utils/config"
	"github.com/EasyCode-Platform/app-backend/src/utils/logger"
)

func Test_Storage(t *testing.T) {
	globalConfig := config.GetInstance()
	logger := logger.NewSugardLogger()

	// init validator

	// init driver
	postgresDriver, err := postgres.NewPostgresConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		t.Errorf("Error in startup, postgres init failed.")
	}
	mongodbDriver, err := mongodb.NewMongodbConnectionByGlobalConfig(globalConfig, logger)
	if err != nil {
		t.Errorf("Error in startup, mongodb init failed")
	}
	storage := NewStorage(postgresDriver, mongodbDriver, logger)

	ComponentId, err := storage.ComponentStorage.CreateNewComponent(`{
  "squadName": "Super hero squad",
  "homeTown": "Metro City",
  "formed": 2016,
  "secretBase": "Super tower",
  "active": true,
  "members": [
    {
      "name": "Molecule Man",
      "age": 29,
      "secretIdentity": "Dan Jukes",
      "powers": ["Radiation resistance", "Turning tiny", "Radiation blast"]
    },
    {
      "name": "Madame Uppercut",
      "age": 39,
      "secretIdentity": "Jane Wilson",
      "powers": [
        "Million tonne punch",
        "Damage resistance",
        "Superhuman reflexes"
      ]
    },
    {
      "name": "Eternal Flame",
      "age": 1000000,
      "secretIdentity": "Unknown",
      "powers": [
        "Immortality",
        "Heat Immunity",
        "Inferno",
        "Teleportation",
        "Interdimensional travel"
      ]
    }
  ]
}`)
	if err != nil {
		t.Errorf("Failed to create component")
	}
	resultJson, err := storage.ComponentStorage.RetrieveComponent(ComponentId)
	if err != nil {
		t.Errorf("Failed to retrieve component")
	}
	t.Logf("retrieve json %s", resultJson)

	app := model.NewApp("test app", 123, 123, ComponentId)
	appId, err := storage.AppStorage.CreateApp(app)
	if err != nil {
		t.Errorf("Failed to create a new app")
	}
	t.Logf("new app id is %d", appId)
	apps, err := storage.AppStorage.RetrieveAllApp(123)
	if err != nil {
		t.Errorf("Failed to retrieve app from teamId : %d", 123)
	}
	t.Logf("Ans of retriving all app is %+v", apps)

	appans, err := storage.AppStorage.RetrieveAppByUid(appId)
	if err != nil {
		t.Errorf("Failed to retrieve app from Uid : %d", appId)
	}
	t.Logf("Ans of retriving app by id is %+v", appans)
}
