package db

import (
	"database/sql"
	mydb "filestore/db/mysql"
	"fmt"
)

// UserSignup 通过用户名及密码完成user表的注册
func UserSignup(username string, passwd string) bool {
	stmt, err := mydb.DbConn().Prepare("insert ignore into tbl_user(`user_name`,`user_pwd`)value (?,?)")
	if err != nil {
		fmt.Println("Failed to insert,err:" + err.Error())
		return false
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("stmt 对象关闭失败")
		}
	}(stmt)

	ret, err := stmt.Exec(username, passwd)
	if err != nil {
		fmt.Println("Failed to Exec,err:" + err.Error())
		return false
	}
	if rowsAffected, err := ret.RowsAffected(); nil == err && rowsAffected > 0 {
		return true
	}
	return false
}

// UserSignIn 判断密码是否一致
func UserSignIn(username string, encPwd string) bool {
	stmt, err := mydb.DbConn().Prepare("SELECT * from tbl_user where user_name=? limit 1")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}

	rows, err := stmt.Query(username)
	if err != nil {
		fmt.Println(err.Error())
		return false
	} else if rows == nil {
		fmt.Println("username not found:" + username)
		return false
	}

	pRows := mydb.ParseRows(rows)
	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encPwd {
		return true
	}

	return false

}

// UpdateToken 刷新用户登录的Token
func UpdateToken(username string, token string) bool {
	stmt, err := mydb.DbConn().Prepare("replace into tbl_user_token(`user_name`,`user_token`)values (?,?)")
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(username, token)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

type User struct {
	Username string
	Email    string
	Phone    string
	SignupAt string
}

func GetUserInfo(username string) (User, error) {
	user := User{}

	stmt, err := mydb.DbConn().Prepare("SELECT user_name, signup_at FROM tbl_user WHERE user_name=? LIMIT 1")
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
	if err != nil {
		fmt.Println(err.Error())
		return user, err
	}

	return user, nil
}
