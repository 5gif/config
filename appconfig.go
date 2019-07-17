package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	log "github.com/Sirupsen/logrus"

	"math"

	"github.com/spf13/viper"
	"github.com/wiless/vlib"
)

//AppConfig  Struct for the app parameteres
type AppConfig struct {
	CarriersGHz      float64
	ISD              float64
	TxPowerDbm       float64
	Out2IndoorLossDb float64
	UENoiseFigureDb  float64
	BSNoiseFigureDb  float64
	INDOORRatio      float64
	INCARRatio       float64
	INCARLossdB      float64
	// ActiveCells     int
	NCells        int
	ActiveBSCells int // The number of cells where the BS are enabled (for link)
	ActiveUECells int // The number of cells where the UEs are dropped ..
	AntennaVTilt  float64
	Extended      bool
	ForceAllLOS   bool
	ShadowLoss    bool
	BandwidthMHz  float64
	BSHeight      float64
	UEHeight      float64
	LogInfo       bool
	NumUEperCell  int
	UEcells       []int
	BScells       []int
}

// NewAppConfig - contains nested structures of other config structures
type NewAppConfig struct {
	ITUconfig
	NRconfig
	SIMconfig
}

// C1 Global Configuration variable
var C1 AppConfig

// C2 Global Configuration variable
var C2 NewAppConfig

// SetDefaults loads the default values for the simulation
func (C1 *AppConfig) SetDefaults() {
	C1.CarriersGHz = .7 // 700MHz
	C1.INDOORRatio = 0
	C1.INCARRatio = 0
	C1.INCARLossdB = 0
	C1.Out2IndoorLossDb = 0
	C1.NCells = 19
	C1.ActiveBSCells = -1 // Default all the cells are active
	C1.ActiveUECells = -1 // UEs are dropped in all the cells
	C1.Extended = false
	C1.ForceAllLOS = false
	C1.BandwidthMHz = 10
	C1.UENoiseFigureDb = 7
	C1.BSNoiseFigureDb = 5
	C1.ShadowLoss = true
	C1.LogInfo = false
	C1.NumUEperCell = 30
	C1.UEcells = []int{0, 10}
	C1.BScells = []int{0, 1, 2}
	// C.TrueCells = -1   // Default to all the cells
	// Do for others too
}

// ReadAppConfig reads all the configuration for the app
func ReadAppConfig(configname string, indir string) (AppConfig, float64, float64, float64, []float64) {
	C1.SetDefaults()
	log.Print("Reading APP config ")
	viper.AddConfigPath(indir)
	viper.SetConfigName(configname)

	err := viper.ReadInConfig()
	if err != nil {
		log.Print("ReadInConfig ", err)
	}

	fmt.Printf("\n INPUT CONFIGURATION %#v", C1)
	err = viper.Unmarshal(&C1)
	if err != nil {
		log.Print("Error unmarshalling ", err)
	}

	if C1.ActiveBSCells == -1 {
		if len(C1.BScells) > 0 {
			C1.ActiveBSCells = len(C1.BScells)
		} else {
			C1.ActiveBSCells = C1.NCells
			C1.BScells = vlib.NewSegmentI(0, C1.ActiveBSCells)
		}
	} else {
		C1.BScells = vlib.NewSegmentI(0, C1.ActiveBSCells)
	}

	if C1.ActiveUECells == -1 {
		if len(C1.UEcells) > 0 {
			C1.ActiveUECells = len(C1.UEcells)
		} else {
			C1.ActiveUECells = C1.NCells
			C1.UEcells = vlib.NewSegmentI(0, C1.ActiveUECells)
		}
	} else {
		C1.UEcells = vlib.NewSegmentI(0, C1.ActiveUECells)
		fmt.Println(C1.UEcells)
	}

	// Set all the default values
	// {
	// 	viper.SetDefault("TxPowerDbm", TxPowerDbm)
	// 	viper.SetDefault("ISD", ISD)
	// 	viper.SetDefault("INDOORRatio", C.INDOORRatio)
	// 	viper.SetDefault("INCARRatio", C.INCARRatio)
	// 	viper.SetDefault("INCARLossdB", C.INCARLossdB)
	// 	viper.SetDefault("Out2IndoorLossDb", C.Out2IndoorLossDb)
	// 	viper.SetDefault("UENoiseFigureDb", C.UENoiseFigureDb)
	// 	viper.SetDefault("BSNoiseFigureDb", C.BSNoiseFigureDb)
	// 	viper.SetDefault("ActiveUECells", C.ActiveUECells)
	// 	viper.SetDefault("ActiveBSCells", C.ActiveBSCells)
	// 	viper.SetDefault("ForceAllLOS", C.ForceAllLOS)
	// 	CellRadius = ISD / math.Sqrt(3.0)
	// 	log.Println("AppConfig : ", C)
	// }

	// Load from the external configuration files
	ISD := viper.GetFloat64("ISD")
	TxPowerDbm := viper.GetFloat64("TxpowerDBm")
	CellRadius := ISD / math.Sqrt(3.0)
	CarriersGHz := []float64{C1.CarriersGHz}
	SaveAppConfig()
	return C1, ISD, TxPowerDbm, CellRadius, CarriersGHz
}

// SetAppConfig reads all the configuration for the app
func SetAppConfig(itucfg ITUconfig, nrcfg NRconfig, simcfg SIMconfig) (AppConfig, float64, []float64) {
	C1.SetDefaults()
	log.Print("Setting App configurations")

	C1.CarriersGHz = itucfg.CarriersGHz
	C1.ISD = itucfg.ISD
	C1.TxPowerDbm = itucfg.TxPowerDbm
	C1.INDOORRatio = itucfg.INDOORRatio
	C1.BSNoiseFigureDb = itucfg.BSNoiseFigureDb
	C1.UENoiseFigureDb = itucfg.UENoiseFigureDb
	C1.NumUEperCell = itucfg.NumUEperCell
	C1.INCARRatio = itucfg.INCARRatio
	C1.INCARLossdB = itucfg.INCARLossdB
	C1.Out2IndoorLossDb = itucfg.Out2IndoorLossDb
	C1.NCells = itucfg.NCells

	C1.BandwidthMHz = nrcfg.BandwidthMHz

	C1.ActiveBSCells = simcfg.ActiveBSCells
	C1.ActiveUECells = simcfg.ActiveUECells
	C1.Extended = simcfg.Extended
	C1.ForceAllLOS = simcfg.ForceAllLOS
	C1.ShadowLoss = simcfg.ShadowLoss
	C1.LogInfo = simcfg.LogInfo
	C1.UEcells = simcfg.UEcells
	C1.BScells = simcfg.BScells

	// C1.AntennaVTilt=
	// C1.BSHeight=
	// C1.UEHeight=

	if C1.ActiveBSCells == -1 {
		if len(C1.BScells) > 0 {
			C1.ActiveBSCells = len(C1.BScells)
		} else {
			C1.ActiveBSCells = C1.NCells
			C1.BScells = vlib.NewSegmentI(0, C1.ActiveBSCells)
		}
	} else {
		C1.BScells = vlib.NewSegmentI(0, C1.ActiveBSCells)
	}

	if C1.ActiveUECells == -1 {
		if len(C1.UEcells) > 0 {
			C1.ActiveUECells = len(C1.UEcells)
		} else {
			C1.ActiveUECells = C1.NCells
			C1.UEcells = vlib.NewSegmentI(0, C1.ActiveUECells)
		}
	} else {
		C1.UEcells = vlib.NewSegmentI(0, C1.ActiveUECells)
		fmt.Println(C1.UEcells)
	}

	// Load from the external configuration files
	// ISD := viper.GetFloat64("ISD")
	// TxPowerDbm := viper.GetFloat64("TxpowerDBm")
	CellRadius := C1.ISD / math.Sqrt(3.0)
	CarriersGHz := []float64{C1.CarriersGHz}
	SaveAppConfig()
	return C1, CellRadius, CarriersGHz
}

// GenerateAppConfig - Generates configuration for the app using ITU, NR and SIM config files
func GenerateAppConfig(itucfg ITUconfig, nrcfg NRconfig, simcfg SIMconfig) {
	// C1.SetDefaults()
	log.Print("Generating APP config ")
	fmt.Println("Hello world!!")
	C2.ITUconfig = itucfg
	C2.NRconfig = nrcfg
	C2.SIMconfig = simcfg
	PrintStructsPretty(C2)
}

// SaveAppConfig ....
func SaveAppConfig() {
	// PrintStructsPretty(C1)
	t := time.Now()
	year, month, day := t.Date()
	root, _ := os.Getwd()
	opdir := root + "/" + OutDIR + "/" + strconv.Itoa(day) + "_" + strconv.Itoa(int(month)) + "_" + strconv.Itoa(year)
	// fmt.Println(opdir)
	_, err := os.Stat(opdir)
	if err != nil {
		os.MkdirAll(opdir, 0700)
	}
	//SwitchOutput()
	// pwd, _ := os.Getwd()
	// currentdir := pwd
	os.Chdir(opdir)
	log.Println("Saving APP config to OUTPUT DIR: ", opdir)
	vlib.SaveStructure(C1, "OutputSetting.json", true)
	//SwitchBack()
	os.Chdir(CurrDIR)

}
