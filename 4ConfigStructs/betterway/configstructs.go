package configstructs

import "fmt"

// Config - All config things
type Config struct {
	MaxLineLength int
	BackBuff      int
	BatchSize     int
	NumOutlets    int
	InputFormat   int
	MaxAttempts   int
	Prival        string
	Procid        string
	Hostname      string
	Appname       string
	Msgid         string
	Host          string
	Port          string
	DBName        string
	// ...
}

// NewConnection -
func NewConnection(host string, port string, dbname string) (string, error) {

	newConnection := host + port + dbname

	return newConnection, nil
}

func main() {
	// create entire config struct
	c := Config{}

	// pass JUST the config members needed to create the new connection.
	dbconnection, err := NewConnection(c.Host, c.Port, c.DBName)
	if err != nil {
		// error this bad boy
	}

	fmt.Println(dbconnection)
}
