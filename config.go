package config

import (
	"os"
	"path/filepath"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

var InDIR string
var OutDIR string

func init() {
	InDIR = "."
	OutDIR = "./results"
}

// NRconfig is crap...
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
}

// ITUconfig is crap...
type ITUconfig struct {
	ENV                 string  `json:"ENV"`
	SCENARIO            string  `json:"SCENARIO"`
	CONFIG              string  `json:"CONFIG"`
	LAYOUTTYPE          int     `json:"LAYOUTTYPE"`
	BSantennaType       string  `json:"BSantennaType"`
	UEantennaType       string  `json:"UEantennaType"`
	FcGHz               float64 `json:"FcGHz"`
	Duplexity           string  `json:"Duplexity"`
	BSHeight            int     `json:"BSHeight"`
	UEHeightout         float64 `json:"UEHeightout"`
	UEHeightin          float64 `json:"UEHeightin"`
	BSTxPowerDbm        int     `json:"BSTxPowerDbm"`
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
}

// SetDefaults loads the default values for the simulation
func (i *ITUconfig) SetDefaults() {
	// TODO - Set default values
	// C1.CarriersGHz = .7 // 700MHz
	// C1.INDOORRatio = 0
	// C1.INCARRatio = 0
	// C1.INCARLossdB = 0
	// C1.Out2IndoorLossDb = 0
	// C1.NCells = 19
	// C1.ActiveBSCells = -1 // Default all the cells are active
	// C1.ActiveUECells = -1 // UEs are dropped in all the cells
	// C1.Extended = false
	// C1.ForceAllLOS = false
	// C1.BandwidthMHz = 10
	// C1.UENoiseFigureDb = 7
	// C1.BSNoiseFigureDb = 5
	// C1.ShadowLoss = true
	// C1.LogInfo = false
	// C1.NumUEperCell = 30
	// C1.UEcells = []int{0, 10}
	// C1.BScells = []int{0, 1, 2}
	// C.TrueCells = -1   // Default to all the cells
	// Do for others too
}

func (i *ITUconfig) Save() {
	// log.Printf("ITUconfig : %#v ", i)
	//SwitchOutput()
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
	InDIR = indir
	// pwd, _ := os.Getwd()
	cfg.Read(configname)
	return cfg
}
