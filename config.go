package config

import (
	log "github.com/sirupsen/logrus"

	"os"

	"github.com/spf13/viper"
)

// // InDIR This is a comment
// var InDIR string

// // OutDIR This is a comment
// var OutDIR string

// // CurrDIR This is a comment
// var CurrDIR string

// AppSetting setting to read and write config files
type AppSetting struct {
	INdir          string `json:"inputdir"`
	OUTdir         string `json:"outputdir"`
	SFdir          string `json:"shadowdir"`
	ITUfname       string `json:"itu"`
	NRfname        string `json:"nr"`
	SIMfname       string `json:"sim"`
	CHANNELfname   string `json:"channelcfg"`
	ASSOCRxTxfname string `json:"assocRxTx"`
	UELOCfname     string `json:"uelocation"`
	BSLOCfname     string `json:"bslocation"`
	LINKfname      string `json:"linkproperties"`
}

func init() {
	// InDIR = "."
	// OutDIR = "."
	// CurrDIR, _ = os.Getwd()
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
	// OutDIR = app.OUTdir
	// InDIR = app.INdir

}

// Setup single big quick function
func Setup(inputcfg string) (ITUconfig, NRconfig, SIMconfig, error) {
	var app App
	// var app AppSetting
	app.AppSetting.FromJSON(inputcfg)
	err := app.LoadCfgs()
	// Appcfg, err := app.LoadApp()
	return app.ITUcfg, app.NRcfg, app.SIMcfg, err
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
