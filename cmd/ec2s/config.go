package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
	"time"

	ntx "github.com/npillmayer/nestext"
)

type config struct {
	players int
	host    string
	sysop   string

	launchDate   time.Time
	maintTime    time.Time
	lastMaintRun time.Time

	MaintPeriod int
	IP          string
	Port        int
}

func (c config) NumPlayers() int {
	return c.players
}

func (c config) LaunchDate() string {
	return c.launchDate.Format("2006-01-02")
}

func (c config) MaintTime() string {
	return c.maintTime.Format("15:04")
}

func (c config) Encode() (bytes.Buffer, error) {

	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	if err := enc.Encode(c); err != nil {
		return buff, err
	}

	return buff, nil

}

func (c config) LastMaintRun() string {
	return c.lastMaintRun.Format("15:04")
}

func (cfg *config) Setup(path string) error {

	nasty := errors.New("\nYour " + configFile + " file is not in the proper format! " +
		"Please read game docs for an example.\n")

	filePath := path + configFile
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return errors.New("Error reading config.ns file! Recheck your path.")
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
		nasty2 := "Error parsing config.ns file: "
		switch key.Interface().(string) {
		case "players":
			s := strings.TrimSpace(v.Interface().(string))
			m, err := strconv.Atoi(s)
			if err != nil {
				return errors.New(nasty2 + "players")
			}
			cfg.players = m
		case "host":
			cfg.host = strings.TrimSpace(v.Interface().(string))
		case "sysop":
			cfg.sysop = strings.TrimSpace(v.Interface().(string))
		case "launchDate":
			date := strings.TrimSpace(v.Interface().(string))
			t, err := time.Parse("2006-01-02", date)
			if err != nil {
				return errors.New(nasty2 + "launchDate")
			}
			cfg.launchDate = t
		case "ip":
			cfg.IP = strings.TrimSpace(v.Interface().(string))
		case "port":
			s := strings.TrimSpace(v.Interface().(string))
			m, err := strconv.Atoi(s)
			if err != nil {
				return errors.New(nasty2 + "port")
			}
			cfg.Port = m
		case "maintTime":
			ts := strings.TrimSpace(v.Interface().(string))
			t, err := time.Parse("15:04", ts)
			if err != nil {
				return errors.New(nasty2 + "maintTime")
			}
			cfg.maintTime = t
		case "maintPeriod":
			s := strings.TrimSpace(v.Interface().(string))
			m, err := strconv.Atoi(s)
			if err != nil {
				return errors.New(nasty2 + "maintPeriod")
			}
			cfg.MaintPeriod = m
		default:
			return nasty
		}
	}

	return nil

}
