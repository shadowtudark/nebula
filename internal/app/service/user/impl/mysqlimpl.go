package impl

import (
	"github.com/douyu/jupiter/pkg/store/gorm"
	"meta/internal/app/model/db"
	"meta/internal/app/service/user"
)
type mysqlImpl struct {
	gh *gorm.DB
}
// NewMysqlImpl construct an instance of mysqlImpl
func NewMysqlImpl(gh *gorm.DB) user.Repository {
	return &mysqlImpl{
		gh: gh,
	}
}
func (m *mysqlImpl)Get(id int) (user db.User,err error){
	user = db.User{}
	err = m.gh.Where("id = ?", id).Find(&user).Error
	return
}
