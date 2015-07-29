package conf

//設定とかの制御関係

// Configure is configures
type Configure struct {
	DB DataBase `json:"db_configure"`
}

// DataBase setting
type DataBase struct {
	ServerName string `json:"server_name"`
	Port       int    `json:"port"`
	UserName   string `json:"user_name"`
	Password   string `json:"password"`
}

// NewCoufigure is init configure
func NewCoufigure() *Configure {
	return &Configure{
		DB: *newDataBase("localhost", 5984, "", ""),
	}
}

func newDataBase(serverName string, port int, userName string, password string) *DataBase {
	return &DataBase{
		ServerName: serverName,
		Port:       port,
		UserName:   userName,
		Password:   password,
	}
}

func readConf() {

}
