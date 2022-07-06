package config

import (
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

// // SetDir sets the directory
// func SetDir(in, out string) {
// 	CurrDIR, _ = os.Getwd()
// 	InDIR = CurrDIR + "/" + in
// 	OutDIR = CurrDIR + "/" + out
// 	os.MkdirAll(OutDIR, os.ModePerm)
// }

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
	UETxDbm             float64 `json:"UETxDbm"`
	BuildingTypeLoss    float64 `json:"BuildingTypeLoss"`
	ISD                 float64 `json:"ISD"`
	NumBSelements       int     `json:"NumBSelements"`
	NumUEelements       int     `json:"NumUEelements"`
	INDOORRatio         float64 `json:"INDOORRatio"`
	IndoorSpeed         float64 `json:"IndoorSpeed"`
	Outdoorspeed        float64 `json:"Outdoorspeed"`
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

// // Save config
// func (i *ITUconfig) UnmarshalJSON(d []byte) error {
// 	log.Println("Custom Unmarshaller ", string(d))
// 	// err := json.Decoder(d, i)
// 	r := bytes.NewBuffer(d)
// 	dec := json.NewDecoder(r)
// 	err := dec.Decode(i)
// 	if err == nil && i.fname == "" {
// 		i.SetFname("ITUcfg.json")
// 	}
// 	return err
// }

// Save config
func (i *ITUconfig) Save() {

	//	SwitchOutput()
	// vlib.SaveStructure(s, "OutputSetting.json", true)
	fname := filepath.Base(i.fname)
	vlib.SaveStructure(i, fname, true)
	//	SwitchBack()

	// os.Chdir(CurrDIR)
}

func (i *ITUconfig) SetFname(f string) {
	i.fname = f

}
func (i ITUconfig) FileName() string {
	return i.fname

}
func (i *ITUconfig) Read(f string) error {
	i.SetDefaults()
	i.fname = f
	// viper.AddConfigPath(InDIR) // ssk look for indir
	// viper.SetConfigName(f)
	viper.SetConfigFile(f)
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
	// filepath.Base(configname)
	err := cfg.Read(configname)
	return cfg, err
}
