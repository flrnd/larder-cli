package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

// Token String
type Token struct {
	TokenString string `json:"token"`
}

// Config type to store configuration file info
type Config struct {
	fileName  string
	directory string
	path      string
}

// Parse the configuration token string
func (t *Token) Parse(handle []byte) (string, error) {
	return t.TokenString, json.Unmarshal(handle, &t)
}

// Marshal yaml ...
func (t *Token) Marshal(c Config) ([]byte, error) {
	return json.Marshal(&Token{TokenString: string(t.TokenString)})
}

func usrHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Println("User home error: ", err)
	}
	return join(usr.HomeDir, "/")
}

func firstTimeRun(c Config) error {
	// Since we are going to store our api token,
	// for security reasons we don't want give read access
	// to other users in our system.
	if err := os.MkdirAll(c.directory, 0700); err != nil {
		return err
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("larder-cli needs to create a config file.")
	fmt.Println("Please introduce your API token: ")
	t := Token{}

	b, _ := reader.ReadString('\n')

	// remove '/n' char at the end of the string
	t.TokenString = b[:len(b)-1]
	data, err := t.Marshal(c)
	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(c.path, data, 0700); err != nil {
		return err
	}

	return nil
}
