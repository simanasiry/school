package main

import (
	"github.com/simanasiry/school/app"
	"github.com/spf13/viper"
	"os"
)

func main() {
	medinaTask := app.New(config())

	medinaTask.InitCore()
	medinaTask.InitModules()

	medinaTask.Start()
}

func config() *viper.Viper {
	v := viper.New()
	v.SetConfigType("yaml")
	f, err := os.Open("./config.yaml")
	if err != nil {
		panic("cannot read config: " + err.Error())
	}
	err = v.ReadConfig(f)
	if err != nil {
		panic("cannot read config" + err.Error())
	}

	return v
}
