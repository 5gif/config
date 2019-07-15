package config

import (
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

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

// Save ...
func (i *NRconfig) Save() {
	CurrDIR, _ := os.Getwd()
	t := time.Now()
	year, month, day := t.Date()
	root, _ := os.Getwd()
	OutDIR := root + "/" + OutDIR + "/" + strconv.Itoa(day) + "_" + strconv.Itoa(int(month)) + "_" + strconv.Itoa(year)
	// fmt.Println(OutDIR)
	_, err := os.Stat(OutDIR)
	if err != nil {
		os.MkdirAll(OutDIR, 0700)
	} //else {
	// 	fmt.Println(x)
	// }
	os.Chdir(OutDIR)
	log.Println("Saving NR config OUTPUT DIR: ", OutDIR)
	vlib.SaveStructure(i, i.fname, true)
	os.Chdir(CurrDIR)
}

func (i *NRconfig) Read(f string) error {
	//i.SetDefaults()
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
		log.Print("Error unmarshalling: ", err)
	}
	return err
}

// ReadNRConfig reads all the configuration for the app
func ReadNRConfig(configname string) (NRconfig, error) {
	var cfg NRconfig
	// fmt.Println(InDIR)
	err := cfg.Read(configname)
	return cfg, err
}
