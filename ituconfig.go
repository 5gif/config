package config

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

// SetDir sets the directory
func SetDir(in, out string) {
	InDIR = in
	OutDIR = out
}

// ITUconfig ....
type ITUconfig struct {
	ENV                 string  `json:"ENV"`
	SCENARIO            string  `json:"SCENARIO"`
	CONFIG              string  `json:"CONFIG"`
	LAYOUTTYPE          int     `json:"LAYOUTTYPE"`
	BSantennaType       string  `json:"BSantennaType"`
	UEantennaType       string  `json:"UEantennaType"`
	CarriersGHz         float64 `json:"FcGHz"`
	Duplexity           string  `json:"Duplexity"`
	BSHeight            int     `json:"BSHeight"`
	UEHeightout         float64 `json:"UEHeightout"`
	UEHeightin          float64 `json:"UEHeightin"`
	TxPowerDbm          int     `json:"BSTxPowerDbm"`
	BandwidthMHz        int     `json:"BandwidthMHz"`
	UETxDbm             int     `json:"UETxDbm"`
	BuildingTypeLoss    float64 `json:"BuildingTypeLoss"`
	ISD                 int     `json:"ISD"`
	NumBSelements       int     `json:"NumBSelements"`
	NumUEelements       int     `json:"NumUEelements"`
	INDOORRatio         float64 `json:"INDOORRatio"`
	IndoorSpeed         int     `json:"IndoorSpeed"`
	Outdoorspeed        int     `json:"Outdoorspeed"`
	BSNoiseFigureDb     int     `json:"BSNoiseFigureDb"`
	UENoiseFigureDb     int     `json:"UENoiseFigureDb"`
	BSAntennaEleGainDbi int     `json:"BSAntennaEleGainDbi"`
	UEAntennaEleGainDbi int     `json:"UEAntennaEleGainDbi"`
	N0                  int     `json:"N0"`
	TrafficModel        int     `json:"TrafficModel"`
	NumUEperCell        int     `json:"NumUEperCell"`
	fname               string
	INCARRatio          int
	INCARLossdB         int
	Out2IndoorLossDb    int
	NCells              int
}

//SetDefaults loads the default values for the simulation
func (i *ITUconfig) SetDefaults() {

	i.CarriersGHz = .7 // 700MHz
	i.INDOORRatio = 0
	i.INCARRatio = 0
	i.INCARLossdB = 0
	i.Out2IndoorLossDb = 0
	i.NCells = 19
	i.BandwidthMHz = 10
	i.UENoiseFigureDb = 7
	i.BSNoiseFigureDb = 5
	i.NumUEperCell = 30

}

//Save config
func (i *ITUconfig) Save() {
	//Switch Input
	pwd, _ := os.Getwd()
	currentdir := pwd
	rel, _ := filepath.Rel(currentdir, OutDIR)
	_ = rel
	os.Mkdir(OutDIR, 0700)
	os.Chdir(OutDIR)
	log.Println("Switching to OUTPUT DIR ", OutDIR)
	vlib.SaveStructure(i, i.fname, true)
	//SwitchBack()
	os.Chdir(currentdir)
}

func (i *ITUconfig) Read(f string) {
	i.SetDefaults()
	i.fname = f
	viper.AddConfigPath(InDIR)
	// viper.SetConfigName(f)
	viper.SetConfigFile(InDIR + "/" + f)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig ", err)
	}
	err = viper.Unmarshal(i)
	if err != nil {
		log.Print("Error unmarshalling ", err)
	}

}

// ReadITUConfig reads all the configuration for the app
func ReadITUConfig(configname string, indir string) ITUconfig {
	var cfg ITUconfig
	fmt.Println(InDIR)
	cfg.Read(configname)
	return cfg
}
