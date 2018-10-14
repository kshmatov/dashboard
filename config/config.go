package config

import (
	"encoding/json"
	"io/ioutil"

	"github.com/kshmatov/dashboard/types"
)

type Configuration struct {
	Server types.Server
	Postgres types.Database
	LogFile string
}

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