package config
//reads and writes json

import (
	"encoding/json"
	"os"
	"path/filepath"
)
const jsonFilePath = ".gatorconfig.json"

type Config struct {
	Db_url 				string `json:"db_url"`
	Current_user_name 	string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	//set struct username field
	cfg.Current_user_name = userName
	//write to file
	return Write(*cfg)
}

func Read() (Config, error) {
	//creates full path
	fullPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	//read json data from cfg file and puts into 'file' var
	file, err := os.ReadFile(fullPath)
	if err != nil {
		return Config{}, err
	}

	//unmarshall json data into newStruct
	newStruct := Config{}
	if err := json.Unmarshal(file, &newStruct); err != nil {
		return Config{}, err
	}
	//return new struct
	return newStruct, nil
}

func getConfigFilePath() (string, error) {
	//get home directory
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	//joins paths for fullPath
	fullPath := filepath.Join(home, jsonFilePath)
	return fullPath, nil
}

func Write(cfg Config) error  {
	//grab full path
	fullPath, err := getConfigFilePath()
	if err != nil {
		return err
	}

	//creates or truncates (reduces size to min existing size or nothing) file at end of path
	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}

	defer file.Close()

	//returns encoder that encodes and writes to the give file
	encoder := json.NewEncoder((file))
	//encode->write cfg to file
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}