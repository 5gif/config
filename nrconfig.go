package config

import (
	"path/filepath"

	"github.com/5gif/aas"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

type NRconfig struct {
	AntennaScheme           string      `json:"AntennaScheme"`
	BS                      aas.Antenna `json:"BS"`
	CONFIG                  string      `json:"CONFIG"`
	ChannelModel            string      `json:"ChannelModel"`
	ENV                     string      `json:"ENV"`
	FcGHz                   float64     `json:"FcGHz"`
	FrameStructure          string      `json:"FrameStructure"`
	LAYOUTTYPE              int         `json:"LAYOUTTYPE"`
	MobilityClass           int         `json:"MobilityClass"`
	NumTRxP                 int         `json:"NumTRxP"`
	ReliabiltyAntennaScheme string      `json:"ReliabiltyAntennaScheme"`
	SCENARIO                string      `json:"SCENARIO"`
	SCSKHz                  int         `json:"SCSKHz"`
	UE                      aas.Antenna `json:"UE"`
	fname                   string      `json:"fname"`
	Vtilt                   float64
}

func (i *NRconfig) SetFname(f string) {
	i.fname = f
}

func (i *NRconfig) FileName() string {
	return i.fname
}

func (i *NRconfig) DefaultNRconfig() {

	i.ENV = "RMa"
	i.SCENARIO = "eMBB"
	i.CONFIG = "A"
	i.LAYOUTTYPE = 0
	i.FcGHz = 0.7
	i.NumTRxP = 57
	i.AntennaScheme = "32x4 MU-MIMO, Reciprocity based, 4T SRS"
	i.BS.AntennaConfig = []int{8, 8, 1, 1, 1, 2, 1}
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
	i.UE.AntennaConfig = []int{8, 4, 1, 1, 1, 1, 1}
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

	// SwitchOutput()
	fname := filepath.Base(i.fname)
	vlib.SaveStructure(i, fname, true)
	// SwitchBack()
	// os.Chdir(CurrDIR)
}

func (i *NRconfig) Read(f string) error {
	//i.SetDefaults()
	i.fname = f
	//viper.AddConfigPath(InDIR)
	// viper.SetConfigName(f)
	viper.SetConfigFile(f)
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
