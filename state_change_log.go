package transition

import (
	"fmt"
	"strings"

	"github.com/moisespsena-go/aorm"
	"github.com/aghape/admin"
	"github.com/aghape/audited"
	"github.com/aghape/aghape/resource"
	"github.com/aghape/roles"
)

// StateChangeLog a model that used to keep state change logs
type StateChangeLog struct {
	aorm.Model
	ReferTable string
	ReferID    string
	From       string
	To         string
	Note       string `sql:"size:1024"`
	audited.AuditedModel
}

// GenerateReferenceKey generate reference key used for change log
func GenerateReferenceKey(model interface{}, db *aorm.DB) string {
	var (
		scope         = db.NewScope(model)
		primaryValues []string
	)

	for _, field := range scope.PrimaryFields() {
		primaryValues = append(primaryValues, fmt.Sprint(field.Field.Interface()))
	}

	return strings.Join(primaryValues, "::")
}

// GetStateChangeLogs get state change logs
func GetStateChangeLogs(model interface{}, db *aorm.DB) []StateChangeLog {
	var (
		changelogs []StateChangeLog
		scope      = db.NewScope(model)
	)

	db.Where("refer_table = ? AND refer_id = ?", scope.TableName(), GenerateReferenceKey(model, db)).Find(&changelogs)

	return changelogs
}

// ConfigureQorResource used to configure transition for qor admin
func (stageChangeLog *StateChangeLog) ConfigureQorResource(res resource.Resourcer) {
	if res, ok := res.(*admin.Resource); ok {
		if res.Permission == nil {
			res.Permission = roles.Deny(roles.Update, roles.Anyone).Deny(roles.Create, roles.Anyone)
		} else {
			res.Permission = res.Permission.Deny(roles.Update, roles.Anyone).Deny(roles.Create, roles.Anyone)
		}
	}
}
