package application

import (
	"github.com/antikytheraton/portScanner/pkg/config"
	"github.com/antikytheraton/portScanner/pkg/db"
)

// Application holds an instance of the app with db and config
type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

// Get initializes and returns a new app with config and db session
func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}
	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}
