package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

type NRconfig struct {
	AntennaScheme string `json:"AntennaScheme"`
	BS            struct {
		AntennaConfig   []int     `json:"AntennaConfig"`
		ElectricalTilt  []float64 `json:"ElectricalTilt"`
		Escan           []float64 `json:"Escan"`
		GainDb          float64   `json:"GainDb"`
		HBeamWidth      float64   `json:"HBeamWidth"`
		Omni            bool      `json:"Omni"`
		SLAV            float64   `json:"SLAV"`
		VBeamWidth      float64   `json:"VBeamWidth"`
		EspacingHfactor float64   `json:"EspacingHfactor"`
		EspacingVfactor float64   `json:"EspacingVfactor"`
		MechanicalTilt  float64   `json:"MechanicalTilt"`
		Polarization    []float64 `json:"Polarization"`
		PanelAz         []float64 `json:"PanelAz"`
		PanelEl         []float64 `json:"PanelEl"`
	} `json:"BS"`
	CONFIG                  string  `json:"CONFIG"`
	ChannelModel            string  `json:"ChannelModel"`
	ENV                     string  `json:"ENV"`
	FcGHz                   float64 `json:"FcGHz"`
	FrameStructure          string  `json:"FrameStructure"`
	LAYOUTTYPE              int     `json:"LAYOUTTYPE"`
	MobilityClass           int     `json:"MobilityClass"`
	NumTRxP                 int     `json:"NumTRxP"`
	ReliabiltyAntennaScheme string  `json:"ReliabiltyAntennaScheme"`
	SCENARIO                string  `json:"SCENARIO"`
	SCSKHz                  int     `json:"SCSKHz"`
	UE                      struct {
		AntennaConfig   []int     `json:"AntennaConfig"`
		ElectricalTilt  []float64 `json:"ElectricalTilt"`
		Escan           []float64 `json:"Escan"`
		GainDb          float64   `json:"GainDb"`
		HBeamWidth      float64   `json:"HBeamWidth"`
		Omni            bool      `json:"Omni"`
		SLAV            float64   `json:"SLAV"`
		VBeamWidth      float64   `json:"VBeamWidth"`
		EspacingHfactor float64   `json:"EspacingHfactor"`
		EspacingVfactor float64   `json:"EspacingVfactor"`
		MechanicalTilt  float64   `json:"MechanicalTilt"`
		Polarization    []float64 `json:"Polarization"`
		PanelAz         []float64 `json:"PanelAz"`
		PanelEl         []float64 `json:"PanelEl"`
	} `json:"UE"`
	fname string `json:"fname"`
	Vtilt float64

	//        float64ENV                     string        `json:"ENV"`
	// SCENARIO                string        `json:"SCENARIO"`
	// CONFIG                  string        `json:"CONFIG"`
	// LAYOUTTYPE              int           `json:"LAYOUTTYPE"`
	// FcGHz                   float64       `json:"FcGHz"`
	// NumTRxP                 int           `json:"NumTRxP"`
	// AntennaScheme           string        `json:"AntennaScheme"`
	// BSAntennaConfig         []int         `json:"BSAntennaConfig"`
	// UEAntennaConfig         []int          `json:"UEAntennaConfig"`
	// SCSKHz                  int           `json:"SCSKHz"`
	// FrameStructure          string        `json:"FrameStructure"`
	// ChannelModel            string        `json:"ChannelModel"`
	// ReliabiltyAntennaScheme string        `json:"ReliabiltyAntennaScheme"`
	// MobilityClass           int           `json:"MobilityClass"`
	// GainDb                  float64       `json:"GainDb"`
	// VBeamWidth              float64       `json:"VBeamWidth"`
	// HBeamWidth              float64       `json:"HBeamWidth"`
	// SLAV                    float64       `json:"SLAV"`
	// UEGainDb                  float64       `json:"UEGainDb"`
	// UEVBeamWidth              float64       `json:"UEVBeamWidth"`
	// UEHBeamWidth              float64       `json:"UEHBeamWidth"`
	// UESLAV                    float64       `json:"UESLAV"`
	// ESpacingVFactor         float64       `json:"EspacingVfactor"`
	// ESpacingHFactor         float64       `json:"EspacingHfactor"`
	// ElectricalTilt          []float64     `json:"ElectricalTilt"`
	// Escan                   []float64     `json:"Escan"`
	// UEElectricalTilt        []float64     `json:"UEElectricalTilt"`
	// UEEscan                 []float64     `json:"UEEscan"`
	// MechanicalTilt          float64       `json:"MechanicalTilt"`
	// Omni                    bool           `json:"Omni"`
	// fname                   string
	// Vtilt                   float64
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
