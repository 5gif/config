package config

import (
	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

type Antenna struct {
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
}

type NRconfig struct {
	AntennaScheme           string  `json:"AntennaScheme"`
	BS                      Antenna `json:"BS"`
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
	UE                      Antenna `json:"UE"`
	fname                   string  `json:"fname"`
	Vtilt                   float64

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

func (i *NRconfig) DefaultNRconfig() {

	i.ENV = "RMa"
	i.SCENARIO = "eMBB"
	i.CONFIG = "A"
	i.LAYOUTTYPE = 0
	i.FcGHz = 0.7
	i.NumTRxP = 57
	i.AntennaScheme = "32x4 MU-MIMO, Reciprocity based, 4T SRS"
	i.BS.AntennaConfig = []int{8, 4, 2, 1, 1, 1, 2}
	i.BS.SLAV = 30
	i.BS.HBeamWidth = 65
	i.BS.VBeamWidth = 65
	i.BS.GainDb = 8
	i.BS.ElectricalTilt = []float64{99.2}
	i.BS.Escan = []float64{0.0}
	i.BS.Omni = false
	i.BS.PanelAz = []float64{0.0}
	i.BS.PanelEl = []float64{90.0}
	i.BS.EspacingHfactor = 0.5
	i.BS.EspacingVfactor = 0.8
	i.BS.MechanicalTilt = 90
	i.BS.Polarization = []float64{45, -45}

	i.UE.SLAV = 25
	i.UE.AntennaConfig = []int{8, 4, 2, 1, 1, 1, 2}
	i.UE.HBeamWidth = 90
	i.UE.VBeamWidth = 90
	i.UE.GainDb = 0
	i.UE.ElectricalTilt = []float64{}
	i.UE.Escan = []float64{}
	i.UE.Omni = true
	i.UE.PanelAz = []float64{}
	i.UE.PanelEl = []float64{}
	i.UE.EspacingHfactor = 0.5
	i.UE.EspacingVfactor = 0.8
	i.UE.MechanicalTilt = 90
	i.UE.Polarization = []float64{0, 90}
	i.SCSKHz = 30
	i.FrameStructure = "DDDSU"
	i.ChannelModel = "A"
	i.ReliabiltyAntennaScheme = ""
	i.MobilityClass = 0
	i.fname = "3GPP_RMa_configA"

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
