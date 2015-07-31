package mysql

import (
	"database/sql"

	"github.com/RobotClubKut/SlackBot/lib/conf"
	"github.com/RobotClubKut/SlackBot/lib/log"
	"github.com/RobotClubKut/SlackBot/lib/nosub"
	_ "github.com/go-sql-driver/mysql"
)

//InsertNoSubData is insert Nosub data
func InsertNoSubData(data []nosub.Data) {
	data = func() []nosub.Data {
		s := data
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			s[i], s[j] = s[j], s[i]
		}
		return s
	}()
	configure := conf.ReadConfigure()
	server := configure.MysqlConf.UserName + ":" + configure.MysqlConf.Password + "@/" + configure.MysqlConf.DBName
	db, err := sql.Open("mysql", server)
	log.Terminate(err)
	defer db.Close()
	sqlStr := "INSERT INTO nosub_new(Title, URL, ImageURL, Time)VALUES"
	values := []interface{}{}

	for _, row := range data {
		sqlStr += "(?, ?, ?, ?), "
		values = append(values, row.Title, row.URL, row.ImageURL, row.Time)
	}

	//fmt.Println(values)
	sqlStr = sqlStr[0 : len(sqlStr)-2]
	stmt, _ := db.Prepare(sqlStr)
	_, err = stmt.Exec(values...)
	log.Terminate(err)
}

//GetAnimeMostNewAnime DBにある一番IDが大きいやつを取得
func GetAnimeMostNewAnime() string {
	configure := conf.ReadConfigure()
	server := configure.MysqlConf.UserName + ":" + configure.MysqlConf.Password + "@/" + configure.MysqlConf.DBName
	db, err := sql.Open("mysql", server)
	log.Terminate(err)
	defer db.Close()

	sqlStr := "SELECT Title FROM nosub_new ORDER BY ID DESC LIMIT 1"

	rows, err := db.Query(sqlStr)
	log.Terminate(err)
	return func() string {
		for rows.Next() {
			var title string
			if err := rows.Scan(&title); err != nil {
				log.WriteErrorLog(err)
				return ""
			}
			return title
		}
		return ""
	}()
}

func deleteRedundancy() {
	configure := conf.ReadConfigure()
	server := configure.MysqlConf.UserName + ":" + configure.MysqlConf.Password + "@/" + configure.MysqlConf.DBName
	db, err := sql.Open("mysql", server)
	log.Terminate(err)
	defer db.Close()

	sqlStr := "DELETE FROM nosub_deny_word WHERE ID in ( SELECT ID FROM (SELECT ID FROM nosub_deny_word GROUP BY Word HAVING COUNT(*) >= 2) AS x )"
	stmt, err := db.Prepare(sqlStr)
	log.Terminate(err)
	stmt.Exec()
	log.Terminate(err)
}

//InsertDenyWord is insert words
func InsertDenyWord(w []string) {
	configure := conf.ReadConfigure()
	server := configure.MysqlConf.UserName + ":" + configure.MysqlConf.Password + "@/" + configure.MysqlConf.DBName
	db, err := sql.Open("mysql", server)
	log.Terminate(err)
	defer db.Close()

	sqlStr := "INSERT INTO nosub_deny_word(Word)VALUES"
	values := []interface{}{}

	for _, row := range w {
		sqlStr += "(?), "
		values = append(values, row)
	}

	sqlStr = sqlStr[0 : len(sqlStr)-2]
	stmt, _ := db.Prepare(sqlStr)
	_, err = stmt.Exec(values...)
	deleteRedundancy()
	log.Terminate(err)
}
