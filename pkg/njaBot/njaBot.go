// Code generated DO NOT EDIT. (@generated)

package njaBot

import (
	"context"
	"time"
	"nja/pkg/database"
	"nja/pkg/ogame"
	"nja/pkg/simulator"
	"nja/pkg/tgbotapi"
	"nja/pkg/wrapper"
)
type ExpeditionsConfigs struct {
	Strategy int64
	WaveDelayMin int64
	WaveDelayMax int64
	DelayMin int64
	DelayMax int64
	MinimumPathfinders int64
	PathfindersHome int64
	LargeCargoHome int64
	SmallCargoHome int64
	EarlyPathfinders int64
	EarlyPFSecsMin int64
	EarlyPFSecsMax int64
	MinSpeedEarlyPathfinders ogame.Speed
	RecycleDebris bool
	Sheets []Sheet
}
type FarmAttack struct {
	Session *FarmSession
	CelestialID ogame.CelestialID
	Ships ogame.ShipsInfos
	Where ogame.Coordinate
	Speed ogame.Speed
}
type FarmSession struct {
	FarmSession database.FarmSession
	Cancel context.CancelFunc
}
type IVMBot interface {
	DeleteScript(string) error
	GetID() database.BotID
	GetLang() string
	GetPlayerID() int64
	GetPlayerName() string
	GetServerID() database.ServerID
	GetServerNumber() int64
	GetUniverse() string
	IsPausedScript(string) bool
	IsScriptRunning(string) bool
	PauseScript(string) error
	Publish(interface {}) bool
	ResumeScript(string) error
	SendMessage(int64, string) error
	SetScriptRunAtStart(string, bool) error
	Start() error
	StartScript(string) error
	Stop()
	StopScript(string) error
}
type InlineKeyboardMarkup struct {
	InlineKeyboardMarkup tgbotapi.InlineKeyboardMarkup
}
type NJAFleetBuilder struct {
	FleetBuilder *wrapper.FleetBuilder
}
func (NJAFleetBuilder) AddShips(param1 ogame.ID, param2 int64) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) FlightTime() (int64, int64) { panic("not implemented") }
func (NJAFleetBuilder) OnError(param1 func(error)) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) OnSuccess(param1 func(ogame.Fleet)) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SendNow() (ogame.Fleet, error) { panic("not implemented") }
func (NJAFleetBuilder) SetAllCrystal() *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetAllDeuterium() *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetAllMetal() *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetAllResources() *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetAllShips() *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetCrystal(param1 int64) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetDestination(param1 wrapper.IntoCoordinate) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetDeuterium(param1 int64) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetDuration(param1 int64) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetMetal(param1 int64) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetMinimumDeuterium(param1 int64) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetMission(param1 ogame.MissionID) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetOrigin(param1 wrapper.IntoCelestial) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetRecallIn(param1 int64) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetResources(param1 ogame.Resources) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetShips(param1 ogame.ShipsInfos) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetSpeed(param1 ogame.Speed) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetTx(param1 wrapper.Prioritizable) *wrapper.FleetBuilder { panic("not implemented") }
func (NJAFleetBuilder) SetUnionID(param1 int64) *wrapper.FleetBuilder { panic("not implemented") }
type PhalanxReport struct {
	ID int64
	SessionID int64
	Data PhalanxReportData
	CreatedAt time.Time
}
type PhalanxReportData []ogame.PhalanxFleet
func (PhalanxReportData) Value() (driver.Value, error) { panic("not implemented") }
type PlayerData struct {
	ID int64
	Name string
	Timestamp int64
	PointsTotal int64
	PointsEconomy int64
	PointsResearch int64
	PointsMilitary int64
	PointsMilitaryBuilt int64
	PointsMilitaryDestroyed int64
	PointsMilitaryLost int64
	PointsHonor int64
	PositionTotal int64
	PositionEconomy int64
	PositionResearch int64
	PositionMilitary int64
	PositionMilitaryBuilt int64
	PositionMilitaryDestroyed int64
	PositionMilitaryLost int64
	PositionHonor int64
	MilitaryShips int64
	Celestials []struct { ID ogame.CelestialID; Name string; Coordinate ogame.Coordinate }
	Alliance *struct { ID int64; Name string; Tag string }
}
type PlayerDataShort struct {
	ID int64
	Name string
	IsAdmin bool
	IsInactive bool
	IsLongInactive bool
	Vacation bool
	AllianceID *int64
}
type Sheet struct {
	UUID string
	Name string
	Active bool
	IgnoreMinimumDeuterium bool
	SlotsAllowed int64
	Origin ogame.CelestialID
	SystemFrom int64
	SystemTo int64
	ExpeditionSpeed ogame.Speed
	Duration int64
	SmallCargo int64
	LargeCargo int64
	LightFighter int64
	HeavyFighter int64
	Cruiser int64
	Battleship int64
	ColonyShip int64
	Recycler int64
	EspionageProbe int64
	Bomber int64
	Destroyer int64
	Deathstar int64
	Battlecruiser int64
	Reaper int64
	Pathfinder int64
	SmallCargoAuto bool
	LargeCargoAuto bool
	LightFighterAuto bool
	HeavyFighterAuto bool
	CruiserAuto bool
	BattleshipAuto bool
	ColonyShipAuto bool
	RecyclerAuto bool
	EspionageProbeAuto bool
	BomberAuto bool
	DestroyerAuto bool
	DeathstarAuto bool
	BattlecruiserAuto bool
	ReaperAuto bool
	PathfinderAuto bool
	AutoFleet bool
}
func (Sheet) AutoSplit(param1 ogame.ID) bool { panic("not implemented") }
func (Sheet) ByID(param1 ogame.ID) int64 { panic("not implemented") }
func (Sheet) CountShips() int64 { panic("not implemented") }
func (Sheet) ShouldAutoSplit() bool { panic("not implemented") }
func (Sheet) UnlimitedSlotsAllowed() bool { panic("not implemented") }
type SimulatorResult struct {
	SimulatorResult simulator.SimulatorResult
	AttackerLosses ogame.Resources
	AttackerProfit ogame.Resources
	DefenderLosses ogame.Resources
	DefenderProfit ogame.Resources
	Debris ogame.Resources
	PossiblePlunder ogame.Resources
	CapturedPlunder ogame.Resources
}
func (SimulatorResult) String() string { panic("not implemented") }
