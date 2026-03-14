package service

import (
	"database/sql"
	"errors"
	"todo_list/dao"
	"todo_list/initjwt"
	"todo_list/model"
)

type UserService struct {
	UserDAO *dao.UserDAO
}

func NewUserService() *UserService {
	return &UserService{
		UserDAO: dao.NewUserDAO(),
	}
}

func (u *UserService) compairpassword(password, confirmpassword string) bool {
	if password == confirmpassword {
		return true
	}
	return false
}

func (u *UserService) checkrepeatuser(username string) (error, int) {
	_, err := u.UserDAO.GetUserByName(username)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, 0
		}
		return err, 1
	}
	return nil, 1
}

func (u *UserService) UserRegister(user *model.Userregistermodel) (error, *model.User) {
	k := u.compairpassword(user.Password, user.Confirmpassword)
	if k != true {
		return errors.New("The two passwords do not match."), nil
	}
	//
	err, exists := u.checkrepeatuser(user.Username)
	if err != nil {
		return err, nil
	} else if exists == 1 {
		return errors.New("Already Have This User"), nil
	}
	//
	err = u.UserDAO.UserRegister(user)
	if err != nil {
		return err, nil
	}
	//TODO：调用GetUserByName查找数据库是否存在该用户，以检验插入是否成功
	userinfo := &model.User{
		UserName: user.Username,
		Password: user.Password,
		NickName: user.Nickname,
	}

	return nil, userinfo
}

func (u *UserService) UserLogIn(user *model.UserLogIn) (error, *model.User) {
	userinfo, err := u.UserDAO.GetUserByName(user.UserName)
	if err != nil {
		return err, nil
	}
	if !u.compairpassword(user.PassWord, userinfo.Password) {
		return errors.New("Wrong Password"), nil
	}
	token, err := initjwt.CreateNewToken(userinfo)
	if err != nil {
		return err, nil
	}
	err = initjwt.StoreTokenInRedis(userinfo.UserID, token)
	if err != nil {
		return err, nil
	}
	return nil, userinfo
}

func (u *UserService) ChangePassword(userid int, userCP model.ModelForCP) error {
	user, err := u.UserDAO.GetUserByID(userid)
	if err != nil {
		return err
	}
	if user.Password != userCP.OldPassword {
		return errors.New("OldPassword Not Correct")
	}
	err = u.UserDAO.ChangePassword(userid, userCP)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ChangeUserInfo(userid int, nuserinfo model.ModelForCI) error {
	return u.UserDAO.ChangeUserInfo(userid, nuserinfo)
}
