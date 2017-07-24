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
func NewConnection(cfg *Config) (string, error) {

	// Only use 3 of the config variables.
	// It’s hard to tell which members are used w/o pulling apart all the functions that Config is passed in.
	newConnection := cfg.Host + cfg.Port + cfg.DBName

	return newConnection, nil
}

func main() {
	// create entire config struct
	c := Config{}

	// pass ENTIRE config struct into new connection
	dbconnection, err := NewConnection(&c)
	if err != nil {
		// error this bad boy
	}

	fmt.Println(dbconnection)
}
