package main

import (
    "reflect"
	"time"	
	"strings"
	"strconv"
	"io/ioutil"
	"errors"
	 
	ntx "github.com/npillmayer/nestext"	
)

type configData struct {
	Players int				
	Host string 
	Sysop string
	LaunchDate time.Time 
	MaintPeriod int
	MaintTime time.Time
	IP string
	Port int
}

func loadConfig(path string) (configData, error) {
	
	config := configData{}
	
	nasty := errors.New("\nGame config.ns file not in the proper format! " +
		                "Please read game docs for an example.\n")
	
	filePath := path+configFile	
	f, err := ioutil.ReadFile(filePath)
	if err != nil {
		return config, nasty
	}
	
	nestedText, err := ntx.Parse(strings.NewReader(string(f)))
	if err != nil {
	    return config, nasty
	}
	
	nt := reflect.ValueOf(nestedText)
	if nt.Kind() != reflect.Map {
		return config, nasty
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
					return config, errors.New(nasty2 + "players")
				}
				config.Players = m
			case "host":
				config.Host = strings.TrimSpace(v.Interface().(string))
			case "sysop":
				config.Sysop = strings.TrimSpace(v.Interface().(string))
			case "launchDate":
				date := strings.TrimSpace(v.Interface().(string))
				t,err := time.Parse("2006-01-02", date)
				if err != nil {
					return config, errors.New(nasty2 + "launchDate")
				}
				config.LaunchDate = t
			case "ip":
				config.IP = strings.TrimSpace(v.Interface().(string))
			case "port":
				s := strings.TrimSpace(v.Interface().(string))
				m, err := strconv.Atoi(s)
				if err != nil {
					return config, errors.New(nasty2 + "port")
				}
				config.Port = m
			case "maintTime":
				ts := strings.TrimSpace(v.Interface().(string))
				t,err := time.Parse("15:04", ts)
				if err != nil {
					return config, errors.New(nasty2 + "maintTime")
				}
				config.MaintTime = t				
			case "maintPeriod":
				s := strings.TrimSpace(v.Interface().(string))
				m, err := strconv.Atoi(s)
				if err != nil {
					return config, errors.New(nasty2 + "maintPeriod")
				}
				config.MaintPeriod = m
			default:
				return config, nasty	
		}
	}
	
	return config, nil
	
}
