package database

/*funtion yang diakses secara otomatis ketika package diakses, dibuat dengan nama init*/

var connection string

func init() {
	connection = "MySql"
}

func GetDataBase() string {
	return connection
}
