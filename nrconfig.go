package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

// NRconfig structure
// type NRconfig struct {
// 	ENV                     string        `json:"ENV"`
// 	SCENARIO                string        `json:"SCENARIO"`
// 	CONFIG                  string        `json:"CONFIG"`
// 	LAYOUTTYPE              int           `json:"LAYOUTTYPE"`
// 	FcGHz                   float64       `json:"FcGHz"`
// 	NumTRxP                 int           `json:"NumTRxP"`
// 	AntennaScheme           string        `json:"AntennaScheme"`
// 	BSAntennaConfig         []int         `json:"BSAntennaConfig"`
// 	UEAntennaConfig         []interface{} `json:"UEAntennaConfig"`
// 	SCSKHz                  int           `json:"SCSKHz"`
// 	FrameStructure          string        `json:"FrameStructure"`
// 	ChannelModel            string        `json:"ChannelModel"`
// 	ReliabiltyAntennaScheme string        `json:"ReliabiltyAntennaScheme"`
// 	MobilityClass           int           `json:"MobilityClass"`
// 	fname                   string
// }

type NRconfig struct {
	ENV                     string        `json:"ENV"`
	SCENARIO                string        `json:"SCENARIO"`
	CONFIG                  string        `json:"CONFIG"`
	LAYOUTTYPE              int           `json:"LAYOUTTYPE"`
	FcGHz                   float64       `json:"FcGHz"`
	NumTRxP                 int           `json:"NumTRxP"`
	AntennaScheme           string        `json:"AntennaScheme"`
	BSAntennaConfig         []int         `json:"BSAntennaConfig"`
	UEAntennaConfig         []interface{} `json:"UEAntennaConfig"`
	SCSKHz                  int           `json:"SCSKHz"`
	FrameStructure          string        `json:"FrameStructure"`
	ChannelModel            string        `json:"ChannelModel"`
	ReliabiltyAntennaScheme string        `json:"ReliabiltyAntennaScheme"`
	MobilityClass           int           `json:"MobilityClass"`
	GainDb                  float64       `json:"GainDb"`
	VBeamWidth              float64       `json:"VBeamWidth"`
	HBeamWidth              float64       `json:"HBeamWidth"`
	SLAV                    float64       `json:"SLAV"`
	ESpacingVFactor         float64       `json:"EspacingVfactor"`
	ESpacingHFactor         float64       `json:"EspacingHfactor"`
	ElectricalTilt          []float64     `json:"ElectricalTilt"`
	Escan                   []float64     `json:"Escan"`
	MechanicalTilt          float64       `json:"MechanicalTilt"`
	fname                   string
	Vtilt                   float64
}

// Save ...
func (i *NRconfig) Save() {

	SwitchOutput()
	vlib.SaveStructure(i, i.fname, true)
	SwitchBack()
	// os.Chdir(CurrDIR)
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
