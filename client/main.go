package main

import (
	"log"
	"time"

	"github.com/pkg/errors"
	"github.com/spf13/viper"

	"github.com/7574-sistemas-distribuidos/docker-compose-init/client/common"
)


type Data struct {
	SERVER_PORT, SERVER_LISTEN_BACKLOG, CLI_LOOP_LAPSE, CLI_LOOP_PERIOD,CLI_ID string
}

// InitConfig Function that uses viper library to parse env variables. If
// some of the variables cannot be parsed, an error is returned
func InitConfig() (*viper.Viper, error) {
	v := viper.New()

	
	v.SetConfigName("test") // name of config file (without extension)
	v.AddConfigPath("./data1")   // path to look for the config file in

	if err := v.ReadInConfig(); err != nil {
		log.Print("No config file\n")
	        // Config file not found, try env
	    if _, ok := err.(viper.ConfigFileNotFoundError); ok {
	
			v.AutomaticEnv()
			v.SetEnvPrefix("cli")

			// Add env variables supported
			v.BindEnv("id")

			v.BindEnv("loop", "period")
			v.BindEnv("loop", "lapse")

			v.BindEnv("server", "address")

	    } else {
	    	log.Fatalf("Fatal error config file: %s \n", err)
	    }
	}


	// Parse time.Duration variables and return an error
	// if those variables cannot be parsed
	if _, err := time.ParseDuration(v.GetString("loop_lapse")); err != nil {
		return nil, errors.Wrapf(err, "Could not parse CLI_LOOP_LAPSE env var as time.Duration.")
	}

	if _, err := time.ParseDuration(v.GetString("loop_period")); err != nil {
		return nil, errors.Wrapf(err, "Could not parse CLI_LOOP_PERIOD env var as time.Duration.")
	}

	log.Print("my serve add: %s",v.GetString("server_address"))

	return v, nil
	// Configure viper to read env variables with the CLI_ prefix
	
}

func main() {
	v, err := InitConfig()
	if err != nil {
		log.Fatalf("%s", err)
	}

	clientConfig := common.ClientConfig{
		ServerAddress: v.GetString("server_address"),
		ID:            v.GetString("id"),
		LoopLapse:     v.GetDuration("loop_lapse"),
		LoopPeriod:    v.GetDuration("loop_period"),
	}

	client := common.NewClient(clientConfig)
	client.StartClientLoop()
}
