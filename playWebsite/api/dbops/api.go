package dbops

import (
	"log"
)

func AddUserCredential(loginName string, pwd string) error{
	stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name, password) VALUES (?, ?)")
	if err != nil{
		return err
	}
	stmtIns.Exec(loginName, pwd)
	stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error){
	stmtOut, err := dbConn.Prepare("SELECT password FROM users WHERE login_name = ?")
	//defer stmtOut.Close()  // 在堆栈退出的时候才会调用 稍微会损耗一些性能
	if err != nil{
		log.Println(err)
		return "", err
	}
	var pwd string
	stmtOut.QueryRow(loginName).Scan(&pwd)
	stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM users where login_name=? AND password=?")
	if err != nil{
		log.Println(err)
		return err
	}
	stmtDel.Exec(loginName, pwd)
	stmtDel.Close()
	return nil
}


