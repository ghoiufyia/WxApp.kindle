package resource


import (
	"github.com/ghoiufyia/kindle/app/web/kindle/internal/models"
)

func (r *Resource) First() (usermail *models.UserEmail,err error) {
	usermail = &models.UserEmail{}
	r.originDB.First(usermail)
	return
}