package controller

import (
	"Zabbix-Webassembly/backend/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func ConfigReader() model.Config {

	var config model.Config

	jsonFile, err := os.Open("../config.json")
	if err != nil {
		fmt.Println(err)
		panic("não foi possivel ler o arquivo")
	}
	defer jsonFile.Close()

	configValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
		panic("não foi possivel fazer parse do arquivo")
	}

	json.Unmarshal(configValue, &config)

	return config

}
