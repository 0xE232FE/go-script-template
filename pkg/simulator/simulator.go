// Code generated DO NOT EDIT. (@generated)

package simulator

type Price struct {
	Metal int64
	Crystal int64
	Deuterium int64
}
func (Price) String() string { panic("not implemented") }
func (Price) Total() int64 { panic("not implemented") }
type SimulatorResult struct {
	TacticalRetreat string
	Simulations int64
	AttackerWin int64
	DefenderWin int64
	Draw int64
	Rounds int64
	AttackerLosses Price
	AttackerProfit Price
	DefenderLosses Price
	DefenderProfit Price
	Debris Price
	PossiblePlunder Price
	CapturedPlunder Price
	PctCaptured int64
	Recycler int64
	LCNeeded int64
	SCNeeded int64
	PFNeeded int64
	Moonchance int64
	Logs string
	FlightTime int64
	Fuel int64
	AttackerSmallCargoRemaining int64
	AttackerLargeCargoRemaining int64
	AttackerLightFighterRemaining int64
	AttackerHeavyFighterRemaining int64
	AttackerCruiserRemaining int64
	AttackerBattleshipRemaining int64
	AttackerColonyShipRemaining int64
	AttackerRecyclerRemaining int64
	AttackerEspionageProbeRemaining int64
	AttackerBomberRemaining int64
	AttackerSolarSatelliteRemaining int64
	AttackerDestroyerRemaining int64
	AttackerDeathstarRemaining int64
	AttackerBattlecruiserRemaining int64
	AttackerCrawlerRemaining int64
	AttackerReaperRemaining int64
	AttackerPathfinderRemaining int64
	DefenderSmallCargoRemaining int64
	DefenderLargeCargoRemaining int64
	DefenderLightFighterRemaining int64
	DefenderHeavyFighterRemaining int64
	DefenderCruiserRemaining int64
	DefenderBattleshipRemaining int64
	DefenderColonyShipRemaining int64
	DefenderRecyclerRemaining int64
	DefenderEspionageProbeRemaining int64
	DefenderBomberRemaining int64
	DefenderSolarSatelliteRemaining int64
	DefenderDestroyerRemaining int64
	DefenderDeathstarRemaining int64
	DefenderBattlecruiserRemaining int64
	DefenderCrawlerRemaining int64
	DefenderReaperRemaining int64
	DefenderPathfinderRemaining int64
	DefenderRocketLauncherRemaining int64
	DefenderLightLaserRemaining int64
	DefenderHeavyLaserRemaining int64
	DefenderGaussCannonRemaining int64
	DefenderIonCannonRemaining int64
	DefenderPlasmaTurretRemaining int64
	DefenderSmallShieldDomeRemaining int64
	DefenderLargeShieldDomeRemaining int64
}
func (SimulatorResult) String() string { panic("not implemented") }
