package config

import (
	"errors"

	log "github.com/Sirupsen/logrus"

	"os"

	cirConfig "github.com/5gif/channel/Cirgen/config"
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

// AppSetting setting to read and write config files
type AppSetting struct {
	INdir        string `json:"inputdir"`
	OUTdir       string `json:"outputdir"`
	ITUfname     string `json:"itu"`
	NRfname      string `json:"nr"`
	SIMfname     string `json:"sim"`
	CHANNELfname string `json: "channelcfg"`
}

func init() {
	InDIR = "."
	OutDIR = "."
	CurrDIR, _ = os.Getwd()
}

// ReadCfgSettings reads all the configuration for the app
func (app *AppSetting) FromJSON(fname string) {
	pwd, _ := os.Getwd()
	viper.SetConfigFile(pwd + "/" + fname)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig: ", err)
	}
	err = viper.Unmarshal(app)
	if err != nil {
		log.Print("Error unmarshalling in viper: ", err)
	} else {
		PrintStructsPretty(app)
	}
	OutDIR = app.OUTdir
	InDIR = app.INdir

}

// Setup single big quick function
func Setup(inputcfg string) (ITUconfig, NRconfig, SIMconfig, antenna.SettingAAS, error) {
	var app AppSetting
	app.FromJSON(inputcfg)
	Appcfg, err := app.LoadApp()
	return Appcfg.ITUcfg, Appcfg.NRcfg, Appcfg.SIMcfg, Appcfg.AAScfg, err
}

// ReadCfgSettings reads all the configuration for the app
func (app *AppSetting) read(fname string) {
	pwd, _ := os.Getwd()
	viper.SetConfigFile(pwd + "/" + fname)
	// viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig: ", err)
	}
	err = viper.Unmarshal(app)
	if err != nil {
		log.Print("Error unmarshalling in viper: ", err)
	} else {
		PrintStructsPretty(app)
	}
}

// LoadApp loads all the config in the App
func (app *AppSetting) LoadApp() (*AppConfigs, error) {
	//app.read(fname)
	DefaultApp.AppSetting = *app
	OutDIR = app.OUTdir
	InDIR = app.INdir

	log.Info("Loading App Configurations", OutDIR)
	SetDir(app.INdir, app.OUTdir)
	app.INdir = InDIR
	app.OUTdir = OutDIR
	SwitchInput()

	var err1, err2, err3, err error
	var ITUcfg ITUconfig
	var NRcfg NRconfig
	var SIMcfg SIMconfig
	var Channelcfg cirConfig.TestEnvironment

	log.Info("Loading ITU Config")
	ITUcfg, err1 = ReadITUConfig(app.ITUfname)
	log.Info("Loading ITU Config ..done")

	log.Info("Loading NR Config")
	NRcfg, err2 = ReadNRConfig(app.NRfname)
	log.Info("Loading NR Config ..done")

	log.Info("Loading SIM Config")
	SIMcfg, err3 = ReadSIMConfig(app.SIMfname)
	SIMcfg.SetSIMconfig(ITUcfg, NRcfg)
	log.Info("Loading SIM Config ..done")

	log.Info("Loading CHANNEL Config")
	Channelcfg, err3 = cirConfig.ReadChannelConfig(app.CHANNELfname)
	log.Info("Loading Channel Config ..done")

	var AAS antenna.SettingAAS
	if _, err := os.Stat("sector.json"); os.IsNotExist(err) {
		// path/to/whatever does not exist
		log.Fatal("Unable to find sector.json")
	} else {
		vlib.LoadStructure("sector.json", &AAS)
	}
	/// ----- DONE loading all the ITU, 3GPP ,  SimConfig and related Antenna config
	SwitchBack()

	if err1 != nil || err2 != nil || err3 != nil {
		log.Println(err1, err2, err3)
		err = errors.New("Setup Error: Files not read correctly. Using default values")
	} else {
		err = nil
	}

	// (ITUconfig, NRconfig, SIMconfig, antenna.SettingAAS, error)
	// GenerateSIMconfig - Generates configuration for the app using ITU, NR and SIM config files
	// func SetAppConfigs(itucfg ITUconfig, nrcfg NRconfig, simcfg SIMconfig) {
	DefaultApp.ITUcfg = ITUcfg
	DefaultApp.NRcfg = NRcfg
	DefaultApp.SIMcfg = SIMcfg
	DefaultApp.AAScfg = AAS
	DefaultApp.Channelcfg = Channelcfg

	return &DefaultApp, err
	// return ITUcfg
}
