package config

import (
	"fmt"
	"math"

	log "github.com/Sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

// SIMconfig ...
type SIMconfig struct {
	SCENARIO        string  `json:"SCENARIO"`
	SimulationTime  int     `json:"SimulationTime"`
	FrameType       int     `json:"FrameType"`
	SlotType        int     `json:"SlotType"`
	NumUEperSlot    int     `json:"NumUEperSlot"`
	NumInterference int     `json:"NumInterference"`
	Fname           string  `json:"fname"`
	ActiveBSCells   int     `json:"ActiveBSCells"`
	ActiveUECells   int     `json:"ActiveUECells"`
	Extended        bool    `json:"Extended"`
	ForceAllLOS     bool    `json:"ForceAllLOS"`
	ShadowLoss      bool    `json:"ShadowLoss"`
	LogInfo         bool    `json:"LogInfo"`
	UEcells         []int   `json:"UEcells"`
	BScells         []int   `json:"BScells"`
	TrueCells       int     `json:"TrueCells"`
	SchedulerType   int     `json:"SchedulerType"`
	IsDownLink      bool    `json:"IsDownLink"`
	AntennaVTilt    float64 `json:"AntennaVTilt"`
	CellRadius      float64 `json:"-"`
	NCells          int     `json:"-"`
	ISD             float64 `json:",omitempty"`
	WrapAround      bool    `json:"WrapAround"`
}

//SetDefaults loads the default values for the simulation
func (i *SIMconfig) SetDefaults() {

	i.ActiveBSCells = -1 // Default all the cells are active
	i.ActiveUECells = -1 // UEs are dropped in all the cells
	i.Extended = false
	i.ForceAllLOS = false
	i.ShadowLoss = false
	i.LogInfo = false
	// i.UEcells = []int{0, 1, 2}
	// i.BScells = []int{0, 1, 2}
	i.TrueCells = -1 // Default to all the cells
	i.IsDownLink = true
	i.WrapAround = false
	i.NumInterference = 0

}

// Save ...
func (s *SIMconfig) Save() {
	SwitchOutput()
	vlib.SaveStructure(s, s.Fname+".json", true)
	SwitchBack()

}

func (s *SIMconfig) Read(f string) error {
	s.SetDefaults()
	s.Fname = f
	viper.AddConfigPath(InDIR)
	// viper.SetConfigName(f)
	viper.SetConfigFile(InDIR + "/" + f)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig Error: ", err)
		return err
	}
	err = viper.Unmarshal(s)
	if err != nil {
		log.Print("Error unmarshalling ", err)
		return err
	}
	return nil
}

// ReadSIMConfig reads all the configuration for the app
func ReadSIMConfig(configname string) (SIMconfig, error) {
	var cfg SIMconfig
	// fmt.Println(InDIR)
	err := cfg.Read(configname)
	return cfg, err
}

// SetSIMconfig reads all the configuration for the app
func (s *SIMconfig) SetSIMconfig(itucfg ITUconfig, nrcfg NRconfig) {

	//s.SetDefaults()
	log.Print("Configuring Simulator")
	s.NCells = itucfg.NCells
	//	s.Extended = simcfg.Extended  set based on RMa ??

	if s.ActiveBSCells == -1 {
		if len(s.BScells) > 0 {
			s.ActiveBSCells = len(s.BScells)
		} else {
			s.ActiveBSCells = itucfg.NCells
			s.BScells = vlib.NewSegmentI(0, s.ActiveBSCells)
		}
	} else {
		s.BScells = vlib.NewSegmentI(0, s.ActiveBSCells)
	}

	if s.ActiveUECells == -1 {
		if len(s.UEcells) > 0 {
			s.ActiveUECells = len(s.UEcells)
		} else {
			s.ActiveUECells = itucfg.NCells
			s.UEcells = vlib.NewSegmentI(0, s.ActiveUECells)
		}
	} else {
		s.UEcells = vlib.NewSegmentI(0, s.ActiveUECells)
		fmt.Println(s.UEcells)
	}

	// Load from the external configuration files
	// ISD := viper.GetFloat64("ISD")
	// TxPowerDbm := viper.GetFloat64("TxpowerDBm")
	s.CellRadius = itucfg.ISD / math.Sqrt(3.0)
	s.ISD = itucfg.ISD
	log.Infof("SIMconfig Initialized : %#v", s)
	// return C1, CellRadius, CarriersGHz
}

// ReadSIMconfig reads all the configuration for the app //( float64, float64, float64, []float64)
func (s *SIMconfig) ReadSIMconfig(configname string, indir string) {
	s.SetDefaults()
	log.Print("Reading APP config ")
	viper.AddConfigPath(indir)
	viper.SetConfigName(configname)

	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig ", err)
	}

	fmt.Printf("\n INPUT CONFIGURATION %#v", s)
	err = viper.Unmarshal(&s)
	if err != nil {
		log.Print("Error unmarshalling ", err)
	}

	s.SaveSIMconfig()

}

//SaveSIMconfig ....
func (s *SIMconfig) SaveSIMconfig() {
	SwitchOutput()
	vlib.SaveStructure(s, s.Fname, true)
	SwitchBack()
}
