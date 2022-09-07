package core

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"

	"github.com/greenm01/ec2game/internal/global"
	
	ntx "github.com/npillmayer/nestext"
)

const CONFIGFILE = global.CONFIGFILE

type Config struct {

	GameYear int
	
	NumPlayers int
	Host    string
	Sysop   string

	LaunchDate   time.Time
	MaintTime    time.Time
	LastMaintRun time.Time

	MaintPeriod int
	IP          string
	Port        string
	
}

// LDate return the game launch date as text
func (c *Config) LDate() string {
	return c.LaunchDate.Format("2006-01-02")
}

// MTime return the maintenance time as text
func (c *Config) MTime() string {
	return c.MaintTime.Format("15:04")
}

// LMRun returns last maintenance run time as text
func (c *Config) LMRun() string {
	return c.LastMaintRun.Format("15:04")
}

func (c *Config) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(c); err != nil {
		return buff, err
	}

	return buff, nil

}

func (cfg *Config) Load(path string) error {

	nasty := errors.New("\nYour " + CONFIGFILE + " file is not in the proper format! " +
		"Please read game docs for an example.\n")

	filePath := path + CONFIGFILE
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New("Error reading " + CONFIGFILE + " file! Recheck your path.")
	}

	nestedText, err := ntx.Parse(strings.NewReader(string(f)))
	if err != nil {
		return nasty
	}

	nt := reflect.ValueOf(nestedText)
	if nt.Kind() != reflect.Map {
		return nasty
	}

	/* https://riptutorial.com/go/example/29810/time-parsing */

	for _, key := range nt.MapKeys() {
		v := nt.MapIndex(key)
		nasty2 := "Error parsing " + CONFIGFILE + " file: "
		switch key.Interface().(string) {
		case "players":
			s := strings.TrimSpace(v.Interface().(string))
			m, err := strconv.Atoi(s)
			if err != nil {
				return errors.New(nasty2 + "players")
			}
			cfg.NumPlayers = m
		case "host":
			cfg.Host = strings.TrimSpace(v.Interface().(string))
		case "sysop":
			cfg.Sysop = strings.TrimSpace(v.Interface().(string))
		case "launchDate":
			date := strings.TrimSpace(v.Interface().(string))
			t, err := time.Parse("2006-01-02", date)
			if err != nil {
				return errors.New(nasty2 + "launchDate")
			}
			cfg.LaunchDate = t
		case "ip":
			cfg.IP = strings.TrimSpace(v.Interface().(string))
		case "port":
			cfg.Port = strings.TrimSpace(v.Interface().(string))
		case "maintTime":
			ts := strings.TrimSpace(v.Interface().(string))
			t, err := time.Parse("15:04", ts)
			if err != nil {
				return errors.New(nasty2 + "maintTime")
			}
			cfg.MaintTime = t
		case "maintPeriod":
			s := strings.TrimSpace(v.Interface().(string))
			m, err := strconv.Atoi(s)
			if err != nil {
				return errors.New(nasty2 + "maintPeriod")
			}
			cfg.MaintPeriod = m
		default:
			return errors.New(nasty2 + "unknown value")
		}
	}

	return nil

}
