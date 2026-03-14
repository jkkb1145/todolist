package dao

import (
	"database/sql"
	"fmt"
	"todo_list/global"
	"todo_list/model"
)

type UserDAO struct {
	UserDB *sql.DB
}

func NewUserDAO() *UserDAO {
	return &UserDAO{
		UserDB: global.GetDb(),
	}
}

func (u *UserDAO) GetUserByName(username string) (*model.User, error) {
	user := &model.User{}
	// QueryRow 查询单行数据
	err := global.Db.QueryRow("SELECT UserID,UserName,Password,NickName FROM user WHERE UserName = ?", username).Scan(
		&user.UserID, &user.UserName, &user.Password, &user.NickName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No Such User")
			return nil, err
		}
		fmt.Printf("Error Occurred When Database SELECT. \n %v", err)
		return nil, err
	}
	return user, nil
}

func (u *UserDAO) GetUserByID(userid int) (*model.User, error) {
	user := &model.User{}
	// QueryRow 查询单行数据
	err := global.Db.QueryRow("SELECT UserID,UserName,Password,NickName FROM user WHERE UserID = ?", userid).Scan(
		&user.UserID, &user.UserName, &user.Password, &user.NickName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("No Such User")
			return nil, err
		}
		fmt.Printf("Error Occurred When Database SELECT. \n %v", err)
		return nil, err
	}
	return user, nil
}

func (u *UserDAO) UserRegister(user *model.Userregistermodel) error {
	query := "INSERT INTO user(UserName,Password,NickName) VALUE (?,?,?)"
	_, err := global.Db.Exec(query, user.Username, user.Password, user.Nickname)
	if err != nil {
		fmt.Printf("Error Occurred When Database INSERT INTO. \n %v", err)
		return err
	}
	return nil
}

func (u *UserDAO) ChangePassword(userid int, userCP model.ModelForCP) error {
	query := "UPDATE user SET Password = ? WHERE UserID = ?"
	_, err := global.Db.Exec(query, userCP.NewPassword, userid)
	if err != nil {
		fmt.Printf("Error Occurred When Database UPDATE. \n %v", err)
		return err
	}
	return nil
}

func (u *UserDAO) ChangeUserInfo(userid int, nuserinfo model.ModelForCI) error {
	query := "UPDATE user SET NickName = ? WHERE UserID = ?"
	_, err := global.Db.Exec(query, nuserinfo.NickName, userid)
	if err != nil {
		fmt.Printf("Error Occurred When Database UPDATE. \n %v", err)
		return err
	}
	return nil
}
