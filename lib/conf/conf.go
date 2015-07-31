package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/RobotClubKut/SlackBot/lib/log"
)

//設定とかの制御関係

// Configure is configures
type Configure struct {
	CB        CouchBase `json:"couchbase_configure"`
	MysqlConf Mysql     `json:"mysql_configure"`
	SlackConf Slack     `json:"slack_configure"`
}

// CouchBase setting
type CouchBase struct {
	ServerName string `json:"server_name"`
	Port       int    `json:"port"`
	Pool       string
	DBName     string `json:"db_name"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}

//Mysql setting
type Mysql struct {
	ServerName string `json:"server_name"`
	Port       string `json:"port"`
	DBName     string `json:"db_name"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}

//Slack setting
type Slack struct {
	Token    string `json:"token"`
	UserName string `json:"user_name"`
}

// NewCoufigure is init configure
func NewCoufigure() *Configure {
	return &Configure{
		CB:        *newCouchBase("localhost", 8091, "", "", "", ""),
		MysqlConf: *newMysql("localhost", "3306", "", "", ""),
		SlackConf: *newSlack("", "slackbot"),
	}
}

func newSlack(token string, userName string) *Slack {
	return &Slack{Token: token, UserName: userName}
}

func newMysql(
	serverName string,
	port string,
	dbName string,
	userName string,
	password string,
) *Mysql {
	return &Mysql{
		ServerName: serverName,
		Port:       port,
		DBName:     dbName,
		UserName:   userName,
		Password:   password,
	}
}

func newCouchBase(serverName string, port int, pool string, dbName string, userName string, password string) *CouchBase {
	return &CouchBase{
		ServerName: serverName,
		Port:       port,
		Pool:       pool,
		DBName:     dbName,
		UserName:   userName,
		Password:   password,
	}
}

//ReadConfigure is read all configure
func ReadConfigure() *Configure {
	confBin, err := ioutil.ReadFile("../conf/bot.json")
	log.TerminateAndWriteMessage(err, "can not read configure file.")

	var ret Configure
	json.Unmarshal(confBin, &ret)
	return &ret
}
