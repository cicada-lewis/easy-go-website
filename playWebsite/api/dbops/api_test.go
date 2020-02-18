package dbops

import (
	"github.com/Catelemmon/easy-go-website/playWebsite/api/defs"
	"strconv"
	"testing"
	"time"
)

var tempvid string

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


func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.VideoId
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil{
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUser)
	t.Run("AddComments",testAddComments)
	t.Run("ListComments", testListComments)
}


func testAddComments(t *testing.T) {
	vid := "55555"
	aid := 1
	content := "This video is nice"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error at First AddComments: %v", err)
	}

	content = "I wanna know the next of the video"
	err = AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error at Second AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "55555"
	from := 1514766000
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()  / 1000000000, 10))
	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error at ListComments: %v", err)
	}
	for i, ele := range res {

		t.Logf("comment %d: %v \n", i, ele)
		//fmt.Printf("comment %d: %v \n", i, ele)
	}
}


func TestSessions(t *testing.T) {
	t.Run("AddSessions", testInsertSession)
	t.Run("RetrieveSessions", testRetrieveSession)
	t.Run("RetrieveAllSessions", testRetrieveAllSessions)
	t.Run("DeleteSession", testDeleteSession)
}

func testDeleteSession(t *testing.T) {
	sid := "12345"
	err := DeleteSession(sid)
	if err != nil {
		t.Logf("Error at DeleteSession: %v", err)
	}
	return
}



func testRetrieveAllSessions(t *testing.T) {
	sss, err := RetrieveAllSessions()
	if err != nil {
		t.Errorf("Error at RetrieveAllSessions: %v", err)
	}
	sss.Range(func(k, v interface{}) bool {
		ss, ok := v.(*defs.SimpleSession)
		if !ok {
			return false
		}
		t.Logf("Sid: %s, Username: %s, TTL: %d", k, ss.Username, ss.TTL)
		return true
	})
}

func testRetrieveSession(t *testing.T) {
	ss, err := RetrieveSession("12345")
	if err != nil {
		t.Errorf("Error at RetrieveSession: %v", err)
	}
	t.Logf("tll: %d, username: %s", ss.TTL, ss.Username)
}

func testInsertSession(t *testing.T) {
	err := InsertSession("12345", 233333, "cicada")
	if err != nil {
		t.Errorf("Error at First InsertSession: %v", err)
	}

	err = InsertSession("54321", 233333, "naruto")
	if err != nil {
		t.Errorf("Error at Second InsertSession: %v", err)
	}
}
