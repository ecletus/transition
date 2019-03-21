package transition

import (
	"github.com/ecletus/db"
)

type Plugin struct {
	db.DisDBNames
}

func (p *Plugin) OnRegister() {
	p.DBOnMigrateGorm(func(e *db.GormDBEvent) error {
		return e.DB.AutoMigrate(&StateChangeLog{}).Error
	})
}
