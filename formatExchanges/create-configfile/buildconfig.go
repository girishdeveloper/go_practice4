package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
)

type Server struct {
	Host string `json:"host"`
	Port uint32 `json:"port"`
}

type Database struct {
	Host     string `json:"host"`
	Port     uint32 `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}

type Logging struct {
	Filename   string `json:"filename"`
	Extension  string `json:"extension"`
	Path       string `json:"path"`
	Permission uint32 `json:"permission"`
	Level      string `json:"level"`
}

type Configuration struct {
	Namespace string   `json:"namespace"`
	Server    Server   `json:"server"`
	Database  Database `json:"database"`
	Logging   Logging  `json:"logging"`
}

func main() {
	var namespace string
	var configfile string
	var configuration Configuration
	var serverHost string
	var serverPort uint32
	var databaseHost string
	var databasePort uint32
	var databaseUser string
	var databasePassword string
	var databaseName string
	var logFilename string
	var logExtension string
	var logPath string
	var logPermission string
	var logLevel string

	fmt.Println("Build a config file")
	namespace, found := os.LookupEnv("NAMESPACE")
	if !found {
	NAMESPACE:
		fmt.Print("Enter namespace: ")
		fmt.Scanf("%s", &namespace)
		if len(namespace) == 0 {
			goto NAMESPACE
		}
	}
CONFIGNAME:
	fmt.Print("Enter config file name:")
	fmt.Scanf("%s", &configfile)
	if len(configfile) == 0 {
		goto CONFIGNAME
	}
	fileHandler, err := os.OpenFile("./output/"+configfile+".json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0766)
	if err != nil {
		panic(err)
	}
	existingConf, err := io.ReadAll(fileHandler)
	if err != nil {
		panic(err)
	}
	/*err = fileHandler.Truncate(0)
	if err != nil {
		panic(err)
	}*/
	err = json.Unmarshal(existingConf, &configuration)
	fmt.Print("Enter server host:")
	fmt.Scanf("%s", &serverHost)
	fmt.Print("Enter server port:")
	fmt.Scanf("%d", &serverPort)

	fmt.Print("Enter database host:")
	fmt.Scanf("%s", &databaseHost)
	fmt.Print("Enter database port:")
	fmt.Scanf("%d", &databasePort)
	fmt.Print("Enter database user:")
	fmt.Scanf("%s", &databaseUser)
	fmt.Print("Enter database password:")
	fmt.Scanf("%s", &databasePassword)
	fmt.Print("Enter database name:")
	fmt.Scanf("%s", &databaseName)

	fmt.Print("Enter log file name (without extension): " + namespace + "-")
	fmt.Scanf("%s", &logFilename)
	fmt.Print("Enter log file extension:")
	fmt.Scanf("%s", &logExtension)
	fmt.Print("Enter log file path (directory path):")
	fmt.Scanf("%s", &logPath)
	fmt.Print("Enter log file permission:")
	fmt.Scanf("%s", &logPermission)
	permValue, err := strconv.ParseInt(logPermission, 10, 32)
	fmt.Print("Enter log level:")
	fmt.Scanf("%s", &logLevel)

	//new configuration values
	configuration.Namespace = namespace
	configuration.Server = Server{
		Host: serverHost,
		Port: serverPort,
	}
	configuration.Database = Database{
		Host:     databaseHost,
		Port:     databasePort,
		User:     databaseUser,
		Password: databasePassword,
		Dbname:   databaseName,
	}
	configuration.Logging = Logging{
		Filename:   configuration.Namespace + "-" + logFilename,
		Extension:  logExtension,
		Path:       logPath,
		Permission: uint32(permValue),
		Level:      logLevel,
	}
	bI, err := json.MarshalIndent(configuration, "", " ")
	n, err := io.WriteString(fileHandler, string(bI))
	if err != nil {
		panic(err)
	}
	fmt.Println("output of WriteString", n)
	fileHandler.Close()
	fmt.Println(configuration)
	fmt.Println("Config file is built")
}
