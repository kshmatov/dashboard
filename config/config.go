package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kshmatov/dashboard/types"
)

// Configuration stores configuration data for application
type Configuration struct {
	Server   types.Server
	Postgres types.Database
	LogFile  string
}

// Init loads configuration data from file <fname>
// Parser expects json in file
func Init(fname string) (*Configuration, error) {
	jconf, err := ioutil.ReadFile(fname)
	if err != nil {
		return nil, err
	}
	c := Configuration{}
	err = json.Unmarshal(jconf, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
