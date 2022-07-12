package daos

import (
	"errors"
	_ "log"

	"github.com/macduyhai/loadBalancingGrpcApi/models"

	"github.com/jinzhu/gorm"
)

type UserDao interface {
	Login(username, pass string) (*models.User, error)
	Create(user models.User) (*models.User, error)
	Edit(user models.User, id int64) (*models.User, error)
	Delete(user models.User, id int64) error
}

type userDaoImpl struct {
	db *gorm.DB
}

func NewUserDao(db *gorm.DB) UserDao {
	return &userDaoImpl{db: db}
}

func (dao *userDaoImpl) Login(username, pass string) (*models.User, error) {
	var user models.User
	if err := dao.db.Where("username = ? AND password = ?", username, pass).Find(&user).Error; err != nil {
		return nil, err
	}
	if user.Active == 0 {
		err := errors.New("Inactive account")
		return nil, err
	}
	return &user, nil

}
func (dao *userDaoImpl) Create(user models.User) (*models.User, error) {

	if err := dao.db.Create(&user).Error; err != nil {
		return nil, err

	}
	return &user, nil
}
func (dao *userDaoImpl) Edit(edituser models.User, id int64) (*models.User, error) {
	var user models.User

	if err := dao.db.Where("ID=?", id).Find(&user).Error; err != nil {
		return nil, err
	}
	if user.Username != edituser.Username {
		err := errors.New("Username do not match")
		return nil, err
	}

	if edituser.Fullname != "" {
		user.Fullname = edituser.Fullname
	}
	user.Salary = edituser.Salary
	user.UpdateTime = edituser.UpdateTime

	if err := dao.db.Save(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
func (dao *userDaoImpl) Delete(userdel models.User, id int64) error {
	var user models.User

	if err := dao.db.Where("ID=?", id).Find(&user).Error; err != nil {
		return err
	}
	if user.Username != userdel.Username {
		err := errors.New("Username do not match")
		return err
	}
	if err := dao.db.Delete(&user).Error; err != nil {
		return err
	}
	return nil
}
