package config

import (
	"os"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

// SetDir sets the directory
func SetDir(in, out string) {
	CurrDIR, _ = os.Getwd()
	InDIR = CurrDIR + "/" + in
	OutDIR = CurrDIR + "/" + out
	os.MkdirAll(OutDIR, os.ModePerm)
}

// ITUconfig ....
type ITUconfig struct {
	ENV                 string  `json:"ENV"`
	SCENARIO            string  `json:"SCENARIO"`
	CONFIG              string  `json:"CONFIG"`
	LAYOUTTYPE          int     `json:"LAYOUTTYPE"`
	BSantennaType       string  `json:"BSantennaType"`
	UEantennaType       string  `json:"UEantennaType"`
	CarriersGHz         float64 `json:"CarriersGHz "`
	Duplexity           string  `json:"Duplexity"`
	BSHeight            float64 `json:"BSHeight"`
	UEHeightout         float64 `json:"UEHeightout"`
	UEHeightin          float64 `json:"UEHeightin"`
	TxPowerDbm          float64 `json:"TxPowerDbm"`
	UETxDbm             int     `json:"UETxDbm"`
	BuildingTypeLoss    float64 `json:"BuildingTypeLoss"`
	ISD                 float64 `json:"ISD"`
	NumBSelements       int     `json:"NumBSelements"`
	NumUEelements       int     `json:"NumUEelements"`
	INDOORRatio         float64 `json:"INDOORRatio"`
	IndoorSpeed         int     `json:"IndoorSpeed"`
	Outdoorspeed        int     `json:"Outdoorspeed"`
	BSNoiseFigureDb     float64 `json:"BSNoiseFigureDb"`
	UENoiseFigureDb     float64 `json:"UENoiseFigureDb"`
	BSAntennaEleGainDbi int     `json:"BSAntennaEleGainDbi"`
	UEAntennaEleGainDbi int     `json:"UEAntennaEleGainDbi"`
	N0                  int     `json:"N0"`
	TrafficModel        int     `json:"TrafficModel"`
	NumUEperCell        int     `json:"NumUEperCell"`
	INCARRatio          float64 `json:"INCARRatio"`
	INCARLossdB         float64 `json:"INCARLossdB"`
	Out2IndoorLossDb    float64 `json:"Out2IndoorLossDb"`
	fname               string
	NCells              int     `json:"NCells"`
	BandwidthMHz        float64 `json:"BandwidthMHz"`
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

// Save config
func (i *ITUconfig) Save() {

	vlib.SaveStructure(i, i.fname, true)
	os.Chdir(CurrDIR)
}

func (i *ITUconfig) Read(f string) error {
	i.SetDefaults()
	i.fname = f
	viper.AddConfigPath(InDIR)
	// viper.SetConfigName(f)
	viper.SetConfigFile(InDIR + "/" + f)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig Error: ", err)
	}
	err = viper.Unmarshal(i)
	if err != nil {
		log.Print("Error unmarshalling ", err)
	}
	return err
}

// ReadITUConfig reads all the configuration for the app
func ReadITUConfig(configname string) (ITUconfig, error) {
	var cfg ITUconfig
	// fmt.Println(InDIR)
	err := cfg.Read(configname)
	return cfg, err
}
