package dbops

import (
	"database/sql"
	"log"
)

func AddUserCredential(loginName string, pwd string) error{
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, password) VALUES (?, ?)")
	defer stmtIns.Close()
	if err != nil{
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}

func GetUserCredential(loginName string) (string, error){
	stmtOut, err := dbConn.Prepare("SELECT password FROM users WHERE login_name = ?")
	defer stmtOut.Close()  // 在堆栈退出的时候才会调用 稍微会损耗一些性能
	if err != nil{
		log.Println(err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows{
		return "", err
	}
	//stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users where login_name=? AND password=?")
	defer stmtDel.Close()
	if err != nil{
		log.Println(err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	return nil
}


