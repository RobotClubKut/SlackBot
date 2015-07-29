package conf

import (
	"encoding/json"
	"io/ioutil"

	"github.com/RobotClubKut/SlackBot/lib/log"
)

//設定とかの制御関係

// Configure is configures
type Configure struct {
	DB DataBase `json:"db_configure"`
}

// DataBase setting
type DataBase struct {
	ServerName string `json:"server_name"`
	Port       int    `json:"port"`
	Pool       string
	DBName     string `json:"db_name"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}

// NewCoufigure is init configure
func NewCoufigure() *Configure {
	return &Configure{
		DB: *newDataBase("localhost", 8091, "default", "slack_bot", "", ""),
	}
}

func newDataBase(serverName string, port int, pool string, dbName string, userName string, password string) *DataBase {
	return &DataBase{
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
