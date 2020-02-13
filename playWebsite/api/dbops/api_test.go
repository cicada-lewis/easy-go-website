package dbops

import (
	"testing"
)


func clearTables() {
	dbConn.Exec("TRUNCATE users")
	dbConn.Exec("TRUNCATE video_info")
	dbConn.Exec("TRUNCATE comments")
	dbConn.Exec("TRUNCATE sessions")
}

func TestUserWorkflow(t *testing.T) {

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
		t.Errorf("Error of DelUser", err)
	}
}

func TestMain(m *testing.M) {

}