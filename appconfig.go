package config

import (
	cirgen "github.com/5gif/channel/Cirgen/config"
	"github.com/wiless/cellular/antenna"
)

// //SIMconfig  Struct for the app parameteres
// type SIMconfig struct {
// 	NCells        int
// 	ActiveBSCells int // The number of cells where the BS are enabled (for link)
// 	ActiveUECells int // The number of cells where the UEs are dropped ..
// 	// AntennaVTilt  float64
// 	Extended    bool
// 	ForceAllLOS bool
// 	ShadowLoss  bool
// 	// BandwidthMHz  float64
// 	// BSHeight      float64
// 	// UEHeight      float64
// 	LogInfo bool
// 	// NumUEperCell  int
// 	UEcells    []int
// 	BScells    []int
// 	CellRadius float64
// }

// AppConfigs - contains nested structures of other config structures
type AppConfigs struct {
	AppSetting
	ITUcfg     ITUconfig
	NRcfg      NRconfig
	SIMcfg     SIMconfig
	AAScfg     antenna.SettingAAS
	Channelcfg cirgen.TestEnvironment
}

// DefaultApp Global variable of all Configurations used in the Application
var DefaultApp AppConfigs

// // SetDefaults loads the default values for the simulation
// func (s *SIMconfig) SetDefaults() {

// 	s.ActiveBSCells = -1 // Default all the cells are active
// 	s.ActiveUECells = -1 // UEs are dropped in all the cells
// 	s.Extended = false
// 	s.ForceAllLOS = false

// 	s.ShadowLoss = true
// 	s.LogInfo = false
// 	s.UEcells = []int{0, 10}
// 	s.BScells = []int{0, 1, 2}

// }
