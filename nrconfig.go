package config

import (
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

func init() {
	InDIR = "."
	OutDIR = "./results"
}

// NRconfig structure
type NRconfig struct {
	ENV                     string  `json:"ENV"`
	SCENARIO                string  `json:"SCENARIO"`
	CONFIG                  string  `json:"CONFIG"`
	LAYOUTTYPE              int     `json:"LAYOUTTYPE"`
	ISD                     int     `json:"ISD"`
	FcGHz                   float64 `json:"FcGHz"`
	BandwidthMHz            float64 `json:"BandwidthMHz"`
	SCSKHz                  int     `json:"SCSKHz"`
	Duplexity               string  `json:"Duplexity"`
	AntennaScheme           string  `json:"AntennaScheme"`
	ReliabiltyAntennaScheme string  `json:"ReliabiltyAntennaScheme"`
	BSAntennaConfig         []int   `json:"BSAntennaConfig"`
	UEAntennaConfig         string  `json:"UEAntennaConfig"`
	FrameStructure          string  `json:"FrameStructure"`
	ChannelModel            string  `json:"ChannelModel"`
	NumTRxP                 string  `json:"NumTRxP"`
	MobilityClass           int     `json:"MobilityClass"`
	fname                   string
}

func (i *NRconfig) Save() {
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

func (i *NRconfig) Read(f string) {
	//i.SetDefaults()
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

// ReadNRConfig reads all the configuration for the app
func ReadNRConfig(configname string, indir string) NRconfig {
	var cfg NRconfig
	InDIR = indir
	cfg.Read(configname)
	return cfg
}
