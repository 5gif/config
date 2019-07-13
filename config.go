package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
)

// InDIR This is a comment
var InDIR string

// OutDIR This is a comment
var OutDIR string

// CurrDIR This is a comment
var CurrDIR string

type cfgSettings struct {
	Indir    string `json:"indir"`
	Outdir   string `json:"outdir"`
	Itufname string `json:"itufname"`
	Nrfname  string `json:"nrfname"`
	Simfname string `json:"simfname"`
}

func init() {
	InDIR = "."
	OutDIR = "./results/default"
	CurrDIR = "."
}

var c cfgSettings

// ReadCfgSettings reads all the configuration for the app
func (c *cfgSettings) ReadCfgSettings(fname string) {
	pwd, _ := os.Getwd()
	viper.SetConfigFile(pwd + "/" + fname)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig: ", err)
	}
	err = viper.Unmarshal(c)
	if err != nil {
		log.Print("Error unmarshalling in viper: ", err)
	} else {
		PrintStructsPretty(c)
	}
}

// SwitchInput ...
func SwitchInput(indir string) {
	GOPATH := os.Getenv("GOPATH")
	CurrDIR, _ = os.Getwd()
	InDIR = GOPATH + "/src/github.com/5gif/" + indir
	log.Println("Input Directory: ", InDIR)
	os.Chdir(InDIR)
}

// SwitchBack ...
func SwitchBack() {
	os.Chdir(CurrDIR)
}

// PrintStructsPretty ...
func PrintStructsPretty(c interface{}) {
	fmt.Println(reflect.TypeOf(c), c)
	b, err := json.MarshalIndent(c, "", "    ")
	if err == nil {
		fmt.Println(reflect.TypeOf(c), " struct:")
		fmt.Println(string(b))
	}
}

// Setup ...
// func Setup(fname string) ITUconfig {
func Setup(fname string) (ITUconfig, NRconfig, SIMconfig) {
	c.ReadCfgSettings(fname)
	InDIR = c.Indir
	SwitchInput(InDIR)

	log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")

	var ITUcfg ITUconfig
	ITUcfg = ReadITUConfig(c.Itufname, c.Indir)

	var NRcfg NRconfig
	NRcfg = ReadNRConfig(c.Nrfname, c.Indir)

	var SIMcfg SIMconfig
	SIMcfg = ReadSIMConfig(c.Simfname, c.Indir)

	SwitchBack()
	return ITUcfg, NRcfg, SIMcfg
	// return ITUcfg
}
