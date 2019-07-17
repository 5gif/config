package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"reflect"

	"github.com/spf13/viper"
	"github.com/wiless/cellular/antenna"
	"github.com/wiless/vlib"
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
	OutDIR = "./results"
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
func Setup(fname string) (ITUconfig, NRconfig, SIMconfig, antenna.SettingAAS, error) {
	c.ReadCfgSettings(fname)
	SetDir(c.Indir, c.Outdir)
	SwitchInput(InDIR)

	log.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
	var err1, err2, err3, err error
	var ITUcfg ITUconfig
	ITUcfg, err1 = ReadITUConfig(c.Itufname)

	var NRcfg NRconfig
	NRcfg, err2 = ReadNRConfig(c.Nrfname)

	var SIMcfg SIMconfig
	SIMcfg, err3 = ReadSIMConfig(c.Simfname)

	var AAS antenna.SettingAAS
	// SIMcfg, err3 = ReadSIMConfig(c.Simfname)
	vlib.LoadStructure("sector.json", &AAS)

	SwitchBack()

	if err1 != nil || err2 != nil || err3 != nil {
		err = errors.New("Setup Error: Files not read correctly. Using default values")
	} else {
		err = nil
	}
	return ITUcfg, NRcfg, SIMcfg, AAS, err
	// return ITUcfg
}
