package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"todo_list/global"
	"todo_list/model"
)

type TodoDAO struct {
	TodoDB *sql.DB
}

func NewTodoDAO() *TodoDAO {
	return &TodoDAO{
		TodoDB: global.GetDb(),
	}
}

func (t *TodoDAO) CreateNewList(info *model.CreatNewList) error {
	_, err := global.Db.Exec("insert into list(UserID,Info,Type,CreateAt,UpdateAt,Title)value(?,?,?,?,?,?)", info.UserID, info.Info, info.Type, info.CreatedAt, info.UpdatedAt, info.Title)
	if err != nil {
		fmt.Println("Error Occurred When Database INSERT INTO")
		return err
	}
	return nil
}

func (t *TodoDAO) GetList(listID int) (*model.CreatNewList, error) {
	var list model.CreatNewList

	query := "SELECT ListID, UserID, Info, Type, CreateAt, UpdateAt, Title FROM list WHERE ListID = ?"

	err := global.Db.QueryRow(query, listID).Scan(&list.ListID, &list.UserID, &list.Info, &list.Type, &list.CreatedAt, &list.UpdatedAt, &list.Title)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No Such List")
			return nil, errors.New("No Such List")
		}
		return nil, err
	}

	return &list, nil
}

func (t *TodoDAO) ShowAllList(userID int) (*[]model.CreatNewList, error) {
	// SQL查询语句，查询所有列
	query := `SELECT ListID, UserID, Info, Type, CreateAt, UpdateAt, Title 
              FROM list 
              WHERE UserID = ?`

	// 执行查询
	rows, err := global.Db.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %v", err)
	}
	defer rows.Close() // 确保关闭rows释放连接

	// 创建用户切片
	var lists []model.CreatNewList

	// 遍历结果集
	for rows.Next() {
		var list model.CreatNewList

		// 扫描每一列到结构体字段
		err := rows.Scan(
			&list.ListID,
			&list.UserID,
			&list.Info,
			&list.Type,
			&list.CreatedAt,
			&list.UpdatedAt,
			&list.Title,
		)
		if err != nil {
			return nil, fmt.Errorf("扫描用户数据失败: %v", err)
		}

		// 将用户添加到切片
		lists = append(lists, list)
	}

	// 检查遍历过程中的错误
	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("遍历结果集时发生错误: %v", err)
	}

	return &lists, nil
}

func (t *TodoDAO) UpdateList(listID int, updatelist *model.UpdateList) error {
	// 更新语句
	query := "UPDATE list SET Info = ?, Type = ? , UpdateAt = ? , Title = ? WHERE ListID = ?"

	// 执行更新
	_, err := global.Db.Exec(query, updatelist.Info, updatelist.Type, updatelist.UpdatedAt, updatelist.Title, listID)
	if err != nil {
		return fmt.Errorf("更新失败: %v", err)
	}

	return nil
}

func (t *TodoDAO) DeleteList(listID int) error {
	query := "DELETE FROM list WHERE ListID = ?"

	_, err := global.Db.Exec(query, listID)
	if err != nil {
		return err
	}
	return nil
}
