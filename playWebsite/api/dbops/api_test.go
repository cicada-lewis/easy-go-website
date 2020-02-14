package dbops

import (
	"testing"
)

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func clearTables() {
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE video_info")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestUserWorkflow(t *testing.T) {
	 t.Run("ADD", testAddUser)
	 t.Run("GET", testGetUser)
	 t.Run("DEL", testDelUser)
	 t.Run("REGET", testRegetUser)
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("cicada", "123")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUser(t *testing.T) {
	 res, err := GetUserCredential("cicada")
	 if err != nil{
	 	t.Errorf("Error of GetUser %v", err)
	 }
	 t.Logf("the Result is %s", res)
}

func testDelUser(t *testing.T) {
	err := DeleteUser("cicada", "123")
	if err != nil{
		t.Errorf("Error of DelUser %v", err)
	}
}

func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("cicada")
	if err != nil {
		t.Errorf("error of RegetUser %v", err)
	}

	if pwd != "" {
		t.Error("Delete has failed!")
	}
}

