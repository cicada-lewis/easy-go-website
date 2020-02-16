package dbops

import (
	"database/sql"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/defs"
	"github.com/Catelemmon/easy-go-website/playWebsite/api/utils"
	"log"
	"time"
)

//User

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


// videos

func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid , err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	ct := time.Now()
	ctime := ct.Format("Jan 02 2006, 15:04:05") // M D y, HH:MM:SS
	stmtIns, err := dbConn.Prepare("INSERT INTO video_info (video_id, author_id, name, display_ctime, create_time) VALUES (?, ?, ?, ?, ?)")
	defer stmtIns.Close()
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime, ct)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{
		VideoId: vid,
		AuthorId: aid,
		Name: name,
		DisplayCtime: ctime,
	}

	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE video_id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows{
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{VideoId: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE video_id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	commentId, err := utils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare(
		"INSERT INTO comments (comment_id, video_id, author_id, content, create_time) VALUES (?, ?, ?, ?, ?)")
	defer stmtIns.Close()
	if err != nil {
		return nil
	}

	_, err = stmtIns.Exec(commentId, vid, aid, content, time.Now())
	if err != nil {
		return err
	}
	return nil

}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.comment_id, users.login_name, comments.content from comments
											INNER JOIN users ON comments.author_id = users.user_id 
											WHERE comments.video_id = ? 
											AND comments.create_time > FROM_UNIXTIME(?) AND comments.create_time <= FROM_UNIXTIME(?) `)
	if err != nil {
		return nil, err
	}
	defer stmtOut.Close()
	var res []*defs.Comment
	rows, err := stmtOut.Query(vid, from, to)
	if err != nil{
		return nil, err
	}
	for rows.Next() {
		var commentId, name, content string
		if err := rows.Scan(&commentId, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{
			CommentId: commentId,
			VideoId: vid,
			Author: name,
			Content: content,
		}
		res = append(res, c)
	}
	return res, nil
}