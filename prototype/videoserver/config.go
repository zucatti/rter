// rtER Project - SRL, McGill University, 2013
//
// Author: echa@cim.mcgill.ca

package main

import (
	"encoding/json"
	"io/ioutil"
	"flag"
	"log"
)

// Command Line Options
var configfile = flag.String("config", "config.json", "server config file")

// Server Configuration
type ServerConfig struct {
	// server
	Server struct {
		Addr string `json:"addr"`
		Port uint64 `json:"port"`
		Production_mode bool `json:"production_mode"`
		Secure_mode bool `json:"secure_mode"`
		Cert_file string `json:"cert_file"`
		Key_file string `json:"key_file"`
	}
	// limits
	Limits struct {
		Max_memory_mbytes uint64 `json:"max_memory_mbytes"`
		Max_ingest_sessions uint64 `json:"max_ingest_sessions"`
		Max_ingest_bandwidth_kbit uint64 `json:"max_ingest_bandwidth_kbit"`
		Rate_limit_enable bool `json:"rate_limit_enable"`
		Rate_limit_ingest_window uint64 `json:"rate_limit_ingest_window"`
		Rate_limit_ingest_sessions_per_source uint64 `json:"rate_limit_ingest_sessions_per_source"`
		Rate_limit_ingest_bytes_per_source uint64 `json:"rate_limit_ingest_bytes_per_source"`
	}
	Auth struct {
	// auth
		Signkey string `json:"signkey"`
	}
	// ingest
	Ingest struct {
		Enable_avc_ingest bool `json:"avc"`
		Enable_ts_ingest bool `json:"ts"`
		Enable_chunk_ingest bool `json:"chunk"`
	}
	// transcode
	Transcode struct {
		Hls struct {
			Enabled bool `json:"enabled"`
			Segment_length uint64 `json:"segment_length"`
			Path string `json:"path"`
		}
		Dash struct {
			Enabled bool `json:"enabled"`
			Segment_length uint64 `json:"segment_length"`
			Path string `json:"path"`
		}
		Mp4 struct {
			Enabled bool `json:"enabled"`
			Path string `json:"path"`
		}
		Ogg struct {
			Enabled bool `json:"enabled"`
			Path string `json:"path"`
		}
		Webm struct {
			Enabled bool `json:"enabled"`
			Path string `json:"path"`
		}
		Thumb struct {
			Enabled bool `json:"enabled"`
			Path string `json:"path"`
		}
		Poster struct {
			Enabled bool `json:"enabled"`
			Path string `json:"path"`
		}
	}
}

func ParseConfig(c *ServerConfig) {

	// parse command line
	flag.Parse()

	// set default values
	c.Server.Addr = "127.0.0.1"
	c.Server.Port = 8080
	c.Server.Production_mode = false
	c.Server.Secure_mode = false
	c.Server.Cert_file = ""
	c.Server.Key_file = ""
	c.Limits.Max_memory_mbytes = 128
	c.Limits.Max_ingest_sessions = 10
	c.Limits.Max_ingest_bandwidth_kbit = 10000
	c.Limits.Rate_limit_enable = false
	c.Limits.Rate_limit_ingest_window = 15
	c.Limits.Rate_limit_ingest_sessions_per_source = 100
	c.Limits.Rate_limit_ingest_bytes_per_source = 134217728
	c.Auth.Signkey = "none"
	c.Ingest.Enable_avc_ingest = true
	c.Ingest.Enable_ts_ingest = true
	c.Ingest.Enable_chunk_ingest = false
	c.Transcode.Hls.Enabled = false
	c.Transcode.Hls.Segment_length = 2
	c.Transcode.Hls.Path = "./data/hls"
	c.Transcode.Dash.Enabled = false
	c.Transcode.Dash.Segment_length = 2
	c.Transcode.Dash.Path = "./data/dash"
	c.Transcode.Mp4.Enabled = true
	c.Transcode.Mp4.Path = "./data"
	c.Transcode.Ogg.Enabled = false
	c.Transcode.Ogg.Path = "./data"
	c.Transcode.Webm.Enabled = false
	c.Transcode.Webm.Path = "./data"
	c.Transcode.Thumb.Enabled = false
	c.Transcode.Thumb.Path = "./data"
	c.Transcode.Poster.Enabled = false
	c.Transcode.Poster.Path = "./data"

	// read config
    jsonconfig, err := ioutil.ReadFile(*configfile)
	if err != nil {
    	log.Fatalf("Error reading config file: %s\n", err)
    }

    // unpack config from JSON into Go struct (sets only the defined values)
	err = json.Unmarshal(jsonconfig, &c)
	if err != nil {
    	log.Fatalf("Error parsing config file: %s\n", err)
    }

    log.Printf("ServerConfig: %+v\n", c)

    if c.Server.Production_mode && !c.Server.Secure_mode {
    	log.Printf("Warning: HTTPS is strongly recommended for production mode!")
    }
}