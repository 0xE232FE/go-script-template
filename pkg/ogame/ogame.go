// Code generated DO NOT EDIT. (@generated)

package ogame

import "time"

type ACSValues struct {
	Galaxy int64
	System int64
	Position int64
	CelestialType CelestialType
	Name string
	ACSValues string
	Union int64
}
type ActiveItem struct {
	ID int64
	Ref string
	Name string
	TimeRemaining int64
	TotalDuration int64
	ImgSmall string
}
type AllianceClass int64
func (AllianceClass) IsResearcher() bool { panic("not implemented") }
func (AllianceClass) IsTrader() bool { panic("not implemented") }
func (AllianceClass) IsWarrior() bool { panic("not implemented") }
func (AllianceClass) String() string { panic("not implemented") }
type AllianceInfos struct {
	ID int64
	Name string
	Tag string
	Rank int64
	Member int64
}
type AttackEvent struct {
	ID int64
	MissionType MissionID
	Origin Coordinate
	Destination Coordinate
	DestinationName string
	ArrivalTime time.Time
	ArriveIn int64
	AttackerName string
	AttackerID int64
	UnionID int64
	Missiles int64
	Ships *ShipsInfos
}
func (AttackEvent) String() string { panic("not implemented") }
type Auction struct {
	Ref string
	HasFinished bool
	Endtime int64
	NumBids int64
	CurrentBid int64
	AlreadyBid int64
	MinimumBid int64
	DeficitBid int64
	HighestBidder string
	HighestBidderUserID int64
	CurrentItem string
	CurrentItemLong string
	Inventory int64
	Token string
	ResourceMultiplier struct { Metal float64; Crystal float64; Deuterium float64; Honor int64 }
	Resources map[string]interface {}
}
func (Auction) String() string { panic("not implemented") }
type BuildAccelerators interface {
	GetNaniteFactory() int64
	GetResearchLab() int64
	GetRoboticsFactory() int64
	GetShipyard() int64
}
type BuildingAccelerators interface {
	GetNaniteFactory() int64
	GetRoboticsFactory() int64
}
type CelestialID int64
type CelestialType int64
func (CelestialType) Int() int64 { panic("not implemented") }
func (CelestialType) Int64() int64 { panic("not implemented") }
func (CelestialType) String() string { panic("not implemented") }
type CharacterClass int64
func (CharacterClass) IsCollector() bool { panic("not implemented") }
func (CharacterClass) IsDiscoverer() bool { panic("not implemented") }
func (CharacterClass) IsGeneral() bool { panic("not implemented") }
type ChatMsg struct {
	SenderID int64
	SenderName string
	AssociationID int64
	Text string
	ID int64
	Date int64
}
func (ChatMsg) String() string { panic("not implemented") }
type CombatReportSummary struct {
	ID int64
	APIKey string
	Origin *Coordinate
	Destination Coordinate
	AttackerName string
	DefenderName string
	Loot int64
	Metal int64
	Crystal int64
	Deuterium int64
	DebrisField int64
	CreatedAt time.Time
}
type Coordinate struct {
	Galaxy int64
	System int64
	Position int64
	Type CelestialType
}
func (Coordinate) Debris() Coordinate { panic("not implemented") }
func (Coordinate) Equal(param1 Coordinate) bool { panic("not implemented") }
func (Coordinate) IsDebris() bool { panic("not implemented") }
func (Coordinate) IsMoon() bool { panic("not implemented") }
func (Coordinate) IsPlanet() bool { panic("not implemented") }
func (Coordinate) Moon() Coordinate { panic("not implemented") }
func (Coordinate) Planet() Coordinate { panic("not implemented") }
func (Coordinate) String() string { panic("not implemented") }
type CostTimeBonus struct {
	Cost float64
	Duration float64
}
type CostTimeBonuses map[ID]CostTimeBonus
type DMCost struct {
	Cost int64
	CanBuy bool
	Complete bool
	OGameID ID
	Nbr int64
	BuyAndActivateToken string
	Token string
}
func (DMCost) String() string { panic("not implemented") }
type DMCosts struct {
	Buildings DMCost
	Research DMCost
	Shipyard DMCost
}
func (DMCosts) String() string { panic("not implemented") }
type DMType string
func (DMType) IsValid() bool { panic("not implemented") }
type DefenseAccelerators interface {
	GetNaniteFactory() int64
	GetShipyard() int64
}
type DefensesInfos struct {
	RocketLauncher int64
	LightLaser int64
	HeavyLaser int64
	GaussCannon int64
	IonCannon int64
	PlasmaTurret int64
	SmallShieldDome int64
	LargeShieldDome int64
	AntiBallisticMissiles int64
	InterplanetaryMissiles int64
}
func (DefensesInfos) AttackableValue() int64 { panic("not implemented") }
func (DefensesInfos) ByID(param1 ID) int64 { panic("not implemented") }
func (DefensesInfos) CountShipDefenses() int64 { panic("not implemented") }
func (DefensesInfos) HasMissilesDefense() bool { panic("not implemented") }
func (DefensesInfos) HasShipDefense() bool { panic("not implemented") }
func (DefensesInfos) String() string { panic("not implemented") }
type EmpireCelestial struct {
	Name string
	Diameter int64
	Img string
	ID CelestialID
	Type CelestialType
	Fields Fields
	Temperature Temperature
	Coordinate Coordinate
	Resources Resources
	Supplies ResourcesBuildings
	LfBuildings LfBuildings
	LfResearches LfResearches
	Facilities Facilities
	Defenses DefensesInfos
	Researches Researches
	Ships ShipsInfos
}
type EspionageReport struct {
	Resources Resources
	ID int64
	Username string
	CharacterClass CharacterClass
	AllianceClass AllianceClass
	LastActivity int64
	CounterEspionage int64
	APIKey string
	HasFleetInformation bool
	HasDefensesInformation bool
	HasBuildingsInformation bool
	HasResearchesInformation bool
	HonorableTarget bool
	IsBandit bool
	IsStarlord bool
	IsInactive bool
	IsLongInactive bool
	MetalMine *int64
	CrystalMine *int64
	DeuteriumSynthesizer *int64
	SolarPlant *int64
	FusionReactor *int64
	SolarSatellite *int64
	MetalStorage *int64
	CrystalStorage *int64
	DeuteriumTank *int64
	RoboticsFactory *int64
	Shipyard *int64
	ResearchLab *int64
	AllianceDepot *int64
	MissileSilo *int64
	NaniteFactory *int64
	Terraformer *int64
	SpaceDock *int64
	LunarBase *int64
	SensorPhalanx *int64
	JumpGate *int64
	EnergyTechnology *int64
	LaserTechnology *int64
	IonTechnology *int64
	HyperspaceTechnology *int64
	PlasmaTechnology *int64
	CombustionDrive *int64
	ImpulseDrive *int64
	HyperspaceDrive *int64
	EspionageTechnology *int64
	ComputerTechnology *int64
	Astrophysics *int64
	IntergalacticResearchNetwork *int64
	GravitonTechnology *int64
	WeaponsTechnology *int64
	ShieldingTechnology *int64
	ArmourTechnology *int64
	RocketLauncher *int64
	LightLaser *int64
	HeavyLaser *int64
	GaussCannon *int64
	IonCannon *int64
	PlasmaTurret *int64
	SmallShieldDome *int64
	LargeShieldDome *int64
	AntiBallisticMissiles *int64
	InterplanetaryMissiles *int64
	LightFighter *int64
	HeavyFighter *int64
	Cruiser *int64
	Battleship *int64
	Battlecruiser *int64
	Bomber *int64
	Destroyer *int64
	Deathstar *int64
	SmallCargo *int64
	LargeCargo *int64
	ColonyShip *int64
	Recycler *int64
	EspionageProbe *int64
	Crawler *int64
	Reaper *int64
	Pathfinder *int64
	Coordinate Coordinate
	Type EspionageReportType
	Date time.Time
}
func (EspionageReport) Add(param1 Resources) Resources { panic("not implemented") }
func (EspionageReport) CanAfford(param1 Resources) bool { panic("not implemented") }
func (EspionageReport) CanAfford2(param1 Resources, param2 bool) bool { panic("not implemented") }
func (EspionageReport) DefensesInfos() *DefensesInfos { panic("not implemented") }
func (EspionageReport) Div(param1 Resources) int64 { panic("not implemented") }
func (EspionageReport) Facilities() *Facilities { panic("not implemented") }
func (EspionageReport) FitsIn(param1 Ship, param2 Researches, param3 LfBonuses, param4 CharacterClass, param5 float64, param6 bool) int64 { panic("not implemented") }
func (EspionageReport) Gte(param1 Resources) bool { panic("not implemented") }
func (EspionageReport) Gte2(param1 Resources, param2 bool) bool { panic("not implemented") }
func (EspionageReport) IsDefenceless() bool { panic("not implemented") }
func (EspionageReport) Loot(param1 CharacterClass) Resources { panic("not implemented") }
func (EspionageReport) Lte(param1 Resources) bool { panic("not implemented") }
func (EspionageReport) Mul(param1 int64) Resources { panic("not implemented") }
func (EspionageReport) PlunderRatio(param1 CharacterClass) float64 { panic("not implemented") }
func (EspionageReport) Researches() *Researches { panic("not implemented") }
func (EspionageReport) ResourcesBuildings() *ResourcesBuildings { panic("not implemented") }
func (EspionageReport) ShipsInfos() *ShipsInfos { panic("not implemented") }
func (EspionageReport) String() string { panic("not implemented") }
func (EspionageReport) Sub(param1 Resources) Resources { panic("not implemented") }
func (EspionageReport) SubPercent(param1 float64) Resources { panic("not implemented") }
func (EspionageReport) Total() int64 { panic("not implemented") }
func (EspionageReport) Value() int64 { panic("not implemented") }
type EspionageReportSummary struct {
	ID int64
	Type EspionageReportType
	From string
	Target Coordinate
	LootPercentage float64
}
type EspionageReportType int
type ExpeditionMessage struct {
	ID int64
	Coordinate Coordinate
	Content string
	Resources Resources
	Ships ShipsInfos
	CreatedAt time.Time
}
type Facilities struct {
	RoboticsFactory int64
	Shipyard int64
	ResearchLab int64
	AllianceDepot int64
	MissileSilo int64
	NaniteFactory int64
	Terraformer int64
	SpaceDock int64
	LunarBase int64
	SensorPhalanx int64
	JumpGate int64
}
func (Facilities) ByID(param1 ID) int64 { panic("not implemented") }
func (Facilities) GetAllianceDepot() int64 { panic("not implemented") }
func (Facilities) GetJumpGate() int64 { panic("not implemented") }
func (Facilities) GetLunarBase() int64 { panic("not implemented") }
func (Facilities) GetMissileSilo() int64 { panic("not implemented") }
func (Facilities) GetNaniteFactory() int64 { panic("not implemented") }
func (Facilities) GetResearchLab() int64 { panic("not implemented") }
func (Facilities) GetRoboticsFactory() int64 { panic("not implemented") }
func (Facilities) GetSensorPhalanx() int64 { panic("not implemented") }
func (Facilities) GetShipyard() int64 { panic("not implemented") }
func (Facilities) GetSpaceDock() int64 { panic("not implemented") }
func (Facilities) GetTerraformer() int64 { panic("not implemented") }
func (Facilities) Lazy() LazyFacilities { panic("not implemented") }
func (Facilities) String() string { panic("not implemented") }
type Fields struct {
	Built int64
	Total int64
}
func (Fields) HasFieldAvailable() bool { panic("not implemented") }
type Fleet struct {
	Mission MissionID
	ReturnFlight bool
	InDeepSpace bool
	ID FleetID
	Resources Resources
	Origin Coordinate
	Destination Coordinate
	Ships ShipsInfos
	StartTime time.Time
	ArrivalTime time.Time
	BackTime time.Time
	ArriveIn int64
	BackIn int64
	UnionID int64
	TargetPlanetID int64
}
type FleetDispatchInfos struct {
	Resources Resources
	Ships ShipsInfos
	Slots Slots
	ACSValues []ACSValues
}
type FleetID int64
func (FleetID) IsSet() bool { panic("not implemented") }
func (FleetID) String() string { panic("not implemented") }
type Highscore struct {
	NbPage int64
	CurrPage int64
	Category int64
	Type int64
	Players []HighscorePlayer
}
func (Highscore) String() string { panic("not implemented") }
type HighscorePlayer struct {
	Position int64
	ID int64
	Name string
	Score int64
	AllianceID int64
	HonourPoints int64
	Homeworld Coordinate
	Ships int64
}
func (HighscorePlayer) String() string { panic("not implemented") }
type ID int64
func (ID) Int() int64 { panic("not implemented") }
func (ID) Int64() int64 { panic("not implemented") }
func (ID) IsBuilding() bool { panic("not implemented") }
func (ID) IsCombatShip() bool { panic("not implemented") }
func (ID) IsDefense() bool { panic("not implemented") }
func (ID) IsFacility() bool { panic("not implemented") }
func (ID) IsFlyableShip() bool { panic("not implemented") }
func (ID) IsLfBuilding() bool { panic("not implemented") }
func (ID) IsLfTech() bool { panic("not implemented") }
func (ID) IsResourceBuilding() bool { panic("not implemented") }
func (ID) IsSet() bool { panic("not implemented") }
func (ID) IsShip() bool { panic("not implemented") }
func (ID) IsTech() bool { panic("not implemented") }
func (ID) IsValid() bool { panic("not implemented") }
func (ID) IsValidIPMTarget() bool { panic("not implemented") }
func (ID) String() string { panic("not implemented") }
type IFacilities interface {
	ByID(ID) int64
	GetAllianceDepot() int64
	GetJumpGate() int64
	GetLunarBase() int64
	GetMissileSilo() int64
	GetNaniteFactory() int64
	GetResearchLab() int64
	GetRoboticsFactory() int64
	GetSensorPhalanx() int64
	GetShipyard() int64
	GetSpaceDock() int64
	GetTerraformer() int64
}
type ILfBuildings interface {
	ByID(ID) int64
	GetAcademyOfSciences() int64
	GetAdvancedRecyclingPlant() int64
	GetAntimatterCondenser() int64
	GetAntimatterConvector() int64
	GetAssemblyLine() int64
	GetAutomatisedAssemblyCentre() int64
	GetBioModifier() int64
	GetBiosphereFarm() int64
	GetBiotechLab() int64
	GetChipMassProduction() int64
	GetChrysalisAccelerator() int64
	GetCloningLaboratory() int64
	GetCrystalFarm() int64
	GetCrystalRefinery() int64
	GetDeuteriumSynthesiser() int64
	GetDisruptionChamber() int64
	GetFoodSilo() int64
	GetForumOfTranscendence() int64
	GetFusionCellFactory() int64
	GetFusionPoweredProduction() int64
	GetHallsOfRealisation() int64
	GetHighEnergySmelting() int64
	GetHighPerformanceSynthesiser() int64
	GetHighPerformanceTransformer() int64
	GetMagmaForge() int64
	GetMeditationEnclave() int64
	GetMegalith() int64
	GetMetropolis() int64
	GetMicrochipAssemblyLine() int64
	GetMineralResearchCentre() int64
	GetNanoRepairBots() int64
	GetNeuroCalibrationCentre() int64
	GetOriktorium() int64
	GetPlanetaryShield() int64
	GetProductionAssemblyHall() int64
	GetPsionicModulator() int64
	GetQuantumComputerCentre() int64
	GetResearchCentre() int64
	GetResidentialSector() int64
	GetRoboticsResearchCentre() int64
	GetRuneForge() int64
	GetRuneTechnologium() int64
	GetSanctuary() int64
	GetShipManufacturingHall() int64
	GetSkyscraper() int64
	GetSupraRefractor() int64
	GetUpdateNetwork() int64
	GetVortexChamber() int64
}
type ILfResearches interface {
	ByID(ID) *int64
	GetAcousticScanning() *int64
	GetArtificialSwarmIntelligence() *int64
	GetAutomatedTransportLines() *int64
	GetBattlecruiserMkII() *int64
	GetBomberMkII() *int64
	GetCargoHoldExpansionCivilianShips() *int64
	GetCatalyserTechnology() *int64
	GetCruiserMkII() *int64
	GetDepotAI() *int64
	GetDepthSounding() *int64
	GetDestroyerMkII() *int64
	GetDiamondEnergyTransmitter() *int64
	GetEfficiencyModule() *int64
	GetEfficientSwarmIntelligence() *int64
	GetEnhancedProductionTechnologies() *int64
	GetEnhancedSensorTechnology() *int64
	GetExperimentalRecyclingTechnology() *int64
	GetExperimentalWeaponsTechnology() *int64
	GetFusionDrives() *int64
	GetGeneralOverhaulBattlecruiser() *int64
	GetGeneralOverhaulBattleship() *int64
	GetGeneralOverhaulBomber() *int64
	GetGeneralOverhaulCruiser() *int64
	GetGeneralOverhaulDestroyer() *int64
	GetGeneralOverhaulLightFighter() *int64
	GetGeothermalPowerPlants() *int64
	GetGravitationSensors() *int64
	GetHardenedDiamondDrillHeads() *int64
	GetHeatRecovery() *int64
	GetHighEnergyPumpSystems() *int64
	GetHighPerformanceExtractors() *int64
	GetHighPerformanceTerraformer() *int64
	GetHighTemperatureSuperconductors() *int64
	GetImprovedDroneAI() *int64
	GetImprovedLabTechnology() *int64
	GetImprovedStellarator() *int64
	GetIntergalacticEnvoys() *int64
	GetInterplanetaryAnalysisNetwork() *int64
	GetIonCrystalEnhancementHeavyFighter() *int64
	GetIonCrystalModules() *int64
	GetKaeleshDiscovererEnhancement() *int64
	GetLightFighterMkII() *int64
	GetLowTemperatureDrives() *int64
	GetMagmaPoweredProduction() *int64
	GetMagmaPoweredPumpSystems() *int64
	GetMechanGeneralEnhancement() *int64
	GetNeuroInterface() *int64
	GetNeuromodalCompressor() *int64
	GetObsidianShieldReinforcement() *int64
	GetOptimisedSiloConstructionMethod() *int64
	GetOrbitalDen() *int64
	GetOverclockingBattleship() *int64
	GetOverclockingHeavyFighter() *int64
	GetOverclockingLargeCargo() *int64
	GetPlasmaDrive() *int64
	GetPlasmaTerraformer() *int64
	GetPsionicNetwork() *int64
	GetPsionicShieldMatrix() *int64
	GetPsychoharmoniser() *int64
	GetResearchAI() *int64
	GetRobotAssistants() *int64
	GetRocktalCollectorEnhancement() *int64
	GetRuneShields() *int64
	GetSeismicMiningTechnology() *int64
	GetSixthSense() *int64
	GetSlingshotAutopilot() *int64
	GetStealthFieldGenerator() *int64
	GetSulphideProcess() *int64
	GetSupercomputer() *int64
	GetTelekineticDrive() *int64
	GetTelekineticTractorBeam() *int64
	GetVolcanicBatteries() *int64
}
type IResearches interface {
	ByID(ID) int64
	GetArmourTechnology() int64
	GetAstrophysics() int64
	GetCombustionDrive() int64
	GetComputerTechnology() int64
	GetEnergyTechnology() int64
	GetEspionageTechnology() int64
	GetGravitonTechnology() int64
	GetHyperspaceDrive() int64
	GetHyperspaceTechnology() int64
	GetImpulseDrive() int64
	GetIntergalacticResearchNetwork() int64
	GetIonTechnology() int64
	GetLaserTechnology() int64
	GetPlasmaTechnology() int64
	GetShieldingTechnology() int64
	GetWeaponsTechnology() int64
}
type IResourcesBuildings interface {
	ByID(ID) int64
	GetCrystalMine() int64
	GetCrystalStorage() int64
	GetDeuteriumSynthesizer() int64
	GetDeuteriumTank() int64
	GetFusionReactor() int64
	GetMetalMine() int64
	GetMetalStorage() int64
	GetSolarPlant() int64
	GetSolarSatellite() int64
}
type Item struct {
	Ref string
	Name string
	Image string
	ImageLarge string
	Title string
	Rarity string
	Amount int64
	AmountFree int64
	AmountBought int64
}
type LazyFacilities func() Facilities
func (LazyFacilities) ByID(param1 ID) int64 { panic("not implemented") }
func (LazyFacilities) GetAllianceDepot() int64 { panic("not implemented") }
func (LazyFacilities) GetJumpGate() int64 { panic("not implemented") }
func (LazyFacilities) GetLunarBase() int64 { panic("not implemented") }
func (LazyFacilities) GetMissileSilo() int64 { panic("not implemented") }
func (LazyFacilities) GetNaniteFactory() int64 { panic("not implemented") }
func (LazyFacilities) GetResearchLab() int64 { panic("not implemented") }
func (LazyFacilities) GetRoboticsFactory() int64 { panic("not implemented") }
func (LazyFacilities) GetSensorPhalanx() int64 { panic("not implemented") }
func (LazyFacilities) GetShipyard() int64 { panic("not implemented") }
func (LazyFacilities) GetSpaceDock() int64 { panic("not implemented") }
func (LazyFacilities) GetTerraformer() int64 { panic("not implemented") }
type LazyLfBuildings func() LfBuildings
func (LazyLfBuildings) ByID(param1 ID) int64 { panic("not implemented") }
func (LazyLfBuildings) GetAcademyOfSciences() int64 { panic("not implemented") }
func (LazyLfBuildings) GetAdvancedRecyclingPlant() int64 { panic("not implemented") }
func (LazyLfBuildings) GetAntimatterCondenser() int64 { panic("not implemented") }
func (LazyLfBuildings) GetAntimatterConvector() int64 { panic("not implemented") }
func (LazyLfBuildings) GetAssemblyLine() int64 { panic("not implemented") }
func (LazyLfBuildings) GetAutomatisedAssemblyCentre() int64 { panic("not implemented") }
func (LazyLfBuildings) GetBioModifier() int64 { panic("not implemented") }
func (LazyLfBuildings) GetBiosphereFarm() int64 { panic("not implemented") }
func (LazyLfBuildings) GetBiotechLab() int64 { panic("not implemented") }
func (LazyLfBuildings) GetChipMassProduction() int64 { panic("not implemented") }
func (LazyLfBuildings) GetChrysalisAccelerator() int64 { panic("not implemented") }
func (LazyLfBuildings) GetCloningLaboratory() int64 { panic("not implemented") }
func (LazyLfBuildings) GetCrystalFarm() int64 { panic("not implemented") }
func (LazyLfBuildings) GetCrystalRefinery() int64 { panic("not implemented") }
func (LazyLfBuildings) GetDeuteriumSynthesiser() int64 { panic("not implemented") }
func (LazyLfBuildings) GetDisruptionChamber() int64 { panic("not implemented") }
func (LazyLfBuildings) GetFoodSilo() int64 { panic("not implemented") }
func (LazyLfBuildings) GetForumOfTranscendence() int64 { panic("not implemented") }
func (LazyLfBuildings) GetFusionCellFactory() int64 { panic("not implemented") }
func (LazyLfBuildings) GetFusionPoweredProduction() int64 { panic("not implemented") }
func (LazyLfBuildings) GetHallsOfRealisation() int64 { panic("not implemented") }
func (LazyLfBuildings) GetHighEnergySmelting() int64 { panic("not implemented") }
func (LazyLfBuildings) GetHighPerformanceSynthesiser() int64 { panic("not implemented") }
func (LazyLfBuildings) GetHighPerformanceTransformer() int64 { panic("not implemented") }
func (LazyLfBuildings) GetMagmaForge() int64 { panic("not implemented") }
func (LazyLfBuildings) GetMeditationEnclave() int64 { panic("not implemented") }
func (LazyLfBuildings) GetMegalith() int64 { panic("not implemented") }
func (LazyLfBuildings) GetMetropolis() int64 { panic("not implemented") }
func (LazyLfBuildings) GetMicrochipAssemblyLine() int64 { panic("not implemented") }
func (LazyLfBuildings) GetMineralResearchCentre() int64 { panic("not implemented") }
func (LazyLfBuildings) GetNanoRepairBots() int64 { panic("not implemented") }
func (LazyLfBuildings) GetNeuroCalibrationCentre() int64 { panic("not implemented") }
func (LazyLfBuildings) GetOriktorium() int64 { panic("not implemented") }
func (LazyLfBuildings) GetPlanetaryShield() int64 { panic("not implemented") }
func (LazyLfBuildings) GetProductionAssemblyHall() int64 { panic("not implemented") }
func (LazyLfBuildings) GetPsionicModulator() int64 { panic("not implemented") }
func (LazyLfBuildings) GetQuantumComputerCentre() int64 { panic("not implemented") }
func (LazyLfBuildings) GetResearchCentre() int64 { panic("not implemented") }
func (LazyLfBuildings) GetResidentialSector() int64 { panic("not implemented") }
func (LazyLfBuildings) GetRoboticsResearchCentre() int64 { panic("not implemented") }
func (LazyLfBuildings) GetRuneForge() int64 { panic("not implemented") }
func (LazyLfBuildings) GetRuneTechnologium() int64 { panic("not implemented") }
func (LazyLfBuildings) GetSanctuary() int64 { panic("not implemented") }
func (LazyLfBuildings) GetShipManufacturingHall() int64 { panic("not implemented") }
func (LazyLfBuildings) GetSkyscraper() int64 { panic("not implemented") }
func (LazyLfBuildings) GetSupraRefractor() int64 { panic("not implemented") }
func (LazyLfBuildings) GetUpdateNetwork() int64 { panic("not implemented") }
func (LazyLfBuildings) GetVortexChamber() int64 { panic("not implemented") }
type LazyLfResearches func() LfResearches
func (LazyLfResearches) ByID(param1 ID) *int64 { panic("not implemented") }
func (LazyLfResearches) GetAcousticScanning() *int64 { panic("not implemented") }
func (LazyLfResearches) GetArtificialSwarmIntelligence() *int64 { panic("not implemented") }
func (LazyLfResearches) GetAutomatedTransportLines() *int64 { panic("not implemented") }
func (LazyLfResearches) GetBattlecruiserMkII() *int64 { panic("not implemented") }
func (LazyLfResearches) GetBomberMkII() *int64 { panic("not implemented") }
func (LazyLfResearches) GetCargoHoldExpansionCivilianShips() *int64 { panic("not implemented") }
func (LazyLfResearches) GetCatalyserTechnology() *int64 { panic("not implemented") }
func (LazyLfResearches) GetCruiserMkII() *int64 { panic("not implemented") }
func (LazyLfResearches) GetDepotAI() *int64 { panic("not implemented") }
func (LazyLfResearches) GetDepthSounding() *int64 { panic("not implemented") }
func (LazyLfResearches) GetDestroyerMkII() *int64 { panic("not implemented") }
func (LazyLfResearches) GetDiamondEnergyTransmitter() *int64 { panic("not implemented") }
func (LazyLfResearches) GetEfficiencyModule() *int64 { panic("not implemented") }
func (LazyLfResearches) GetEfficientSwarmIntelligence() *int64 { panic("not implemented") }
func (LazyLfResearches) GetEnhancedProductionTechnologies() *int64 { panic("not implemented") }
func (LazyLfResearches) GetEnhancedSensorTechnology() *int64 { panic("not implemented") }
func (LazyLfResearches) GetExperimentalRecyclingTechnology() *int64 { panic("not implemented") }
func (LazyLfResearches) GetExperimentalWeaponsTechnology() *int64 { panic("not implemented") }
func (LazyLfResearches) GetFusionDrives() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGeneralOverhaulBattlecruiser() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGeneralOverhaulBattleship() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGeneralOverhaulBomber() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGeneralOverhaulCruiser() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGeneralOverhaulDestroyer() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGeneralOverhaulLightFighter() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGeothermalPowerPlants() *int64 { panic("not implemented") }
func (LazyLfResearches) GetGravitationSensors() *int64 { panic("not implemented") }
func (LazyLfResearches) GetHardenedDiamondDrillHeads() *int64 { panic("not implemented") }
func (LazyLfResearches) GetHeatRecovery() *int64 { panic("not implemented") }
func (LazyLfResearches) GetHighEnergyPumpSystems() *int64 { panic("not implemented") }
func (LazyLfResearches) GetHighPerformanceExtractors() *int64 { panic("not implemented") }
func (LazyLfResearches) GetHighPerformanceTerraformer() *int64 { panic("not implemented") }
func (LazyLfResearches) GetHighTemperatureSuperconductors() *int64 { panic("not implemented") }
func (LazyLfResearches) GetImprovedDroneAI() *int64 { panic("not implemented") }
func (LazyLfResearches) GetImprovedLabTechnology() *int64 { panic("not implemented") }
func (LazyLfResearches) GetImprovedStellarator() *int64 { panic("not implemented") }
func (LazyLfResearches) GetIntergalacticEnvoys() *int64 { panic("not implemented") }
func (LazyLfResearches) GetInterplanetaryAnalysisNetwork() *int64 { panic("not implemented") }
func (LazyLfResearches) GetIonCrystalEnhancementHeavyFighter() *int64 { panic("not implemented") }
func (LazyLfResearches) GetIonCrystalModules() *int64 { panic("not implemented") }
func (LazyLfResearches) GetKaeleshDiscovererEnhancement() *int64 { panic("not implemented") }
func (LazyLfResearches) GetLightFighterMkII() *int64 { panic("not implemented") }
func (LazyLfResearches) GetLowTemperatureDrives() *int64 { panic("not implemented") }
func (LazyLfResearches) GetMagmaPoweredProduction() *int64 { panic("not implemented") }
func (LazyLfResearches) GetMagmaPoweredPumpSystems() *int64 { panic("not implemented") }
func (LazyLfResearches) GetMechanGeneralEnhancement() *int64 { panic("not implemented") }
func (LazyLfResearches) GetNeuroInterface() *int64 { panic("not implemented") }
func (LazyLfResearches) GetNeuromodalCompressor() *int64 { panic("not implemented") }
func (LazyLfResearches) GetObsidianShieldReinforcement() *int64 { panic("not implemented") }
func (LazyLfResearches) GetOptimisedSiloConstructionMethod() *int64 { panic("not implemented") }
func (LazyLfResearches) GetOrbitalDen() *int64 { panic("not implemented") }
func (LazyLfResearches) GetOverclockingBattleship() *int64 { panic("not implemented") }
func (LazyLfResearches) GetOverclockingHeavyFighter() *int64 { panic("not implemented") }
func (LazyLfResearches) GetOverclockingLargeCargo() *int64 { panic("not implemented") }
func (LazyLfResearches) GetPlasmaDrive() *int64 { panic("not implemented") }
func (LazyLfResearches) GetPlasmaTerraformer() *int64 { panic("not implemented") }
func (LazyLfResearches) GetPsionicNetwork() *int64 { panic("not implemented") }
func (LazyLfResearches) GetPsionicShieldMatrix() *int64 { panic("not implemented") }
func (LazyLfResearches) GetPsychoharmoniser() *int64 { panic("not implemented") }
func (LazyLfResearches) GetResearchAI() *int64 { panic("not implemented") }
func (LazyLfResearches) GetRobotAssistants() *int64 { panic("not implemented") }
func (LazyLfResearches) GetRocktalCollectorEnhancement() *int64 { panic("not implemented") }
func (LazyLfResearches) GetRuneShields() *int64 { panic("not implemented") }
func (LazyLfResearches) GetSeismicMiningTechnology() *int64 { panic("not implemented") }
func (LazyLfResearches) GetSixthSense() *int64 { panic("not implemented") }
func (LazyLfResearches) GetSlingshotAutopilot() *int64 { panic("not implemented") }
func (LazyLfResearches) GetStealthFieldGenerator() *int64 { panic("not implemented") }
func (LazyLfResearches) GetSulphideProcess() *int64 { panic("not implemented") }
func (LazyLfResearches) GetSupercomputer() *int64 { panic("not implemented") }
func (LazyLfResearches) GetTelekineticDrive() *int64 { panic("not implemented") }
func (LazyLfResearches) GetTelekineticTractorBeam() *int64 { panic("not implemented") }
func (LazyLfResearches) GetVolcanicBatteries() *int64 { panic("not implemented") }
type LazyResearches func() Researches
func (LazyResearches) ByID(param1 ID) int64 { panic("not implemented") }
func (LazyResearches) GetArmourTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetAstrophysics() int64 { panic("not implemented") }
func (LazyResearches) GetCombustionDrive() int64 { panic("not implemented") }
func (LazyResearches) GetComputerTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetEnergyTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetEspionageTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetGravitonTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetHyperspaceDrive() int64 { panic("not implemented") }
func (LazyResearches) GetHyperspaceTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetImpulseDrive() int64 { panic("not implemented") }
func (LazyResearches) GetIntergalacticResearchNetwork() int64 { panic("not implemented") }
func (LazyResearches) GetIonTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetLaserTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetPlasmaTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetShieldingTechnology() int64 { panic("not implemented") }
func (LazyResearches) GetWeaponsTechnology() int64 { panic("not implemented") }
type LazyResourcesBuildings func() ResourcesBuildings
func (LazyResourcesBuildings) ByID(param1 ID) int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetCrystalMine() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetCrystalStorage() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetDeuteriumSynthesizer() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetDeuteriumTank() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetFusionReactor() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetMetalMine() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetMetalStorage() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetSolarPlant() int64 { panic("not implemented") }
func (LazyResourcesBuildings) GetSolarSatellite() int64 { panic("not implemented") }
type LfBonuses struct {
	LfResourceBonuses LfResourceBonuses
	LfShipBonuses LfShipBonuses
	CostTimeBonuses CostTimeBonuses
	PlanetLfResearchCostTimeBonus CostTimeBonus
}
type LfBuildings struct {
	LifeformType LifeformType
	ResidentialSector int64
	BiosphereFarm int64
	ResearchCentre int64
	AcademyOfSciences int64
	NeuroCalibrationCentre int64
	HighEnergySmelting int64
	FoodSilo int64
	FusionPoweredProduction int64
	Skyscraper int64
	BiotechLab int64
	Metropolis int64
	PlanetaryShield int64
	MeditationEnclave int64
	CrystalFarm int64
	RuneTechnologium int64
	RuneForge int64
	Oriktorium int64
	MagmaForge int64
	DisruptionChamber int64
	Megalith int64
	CrystalRefinery int64
	DeuteriumSynthesiser int64
	MineralResearchCentre int64
	AdvancedRecyclingPlant int64
	AssemblyLine int64
	FusionCellFactory int64
	RoboticsResearchCentre int64
	UpdateNetwork int64
	QuantumComputerCentre int64
	AutomatisedAssemblyCentre int64
	HighPerformanceTransformer int64
	MicrochipAssemblyLine int64
	ProductionAssemblyHall int64
	HighPerformanceSynthesiser int64
	ChipMassProduction int64
	NanoRepairBots int64
	Sanctuary int64
	AntimatterCondenser int64
	VortexChamber int64
	HallsOfRealisation int64
	ForumOfTranscendence int64
	AntimatterConvector int64
	CloningLaboratory int64
	ChrysalisAccelerator int64
	BioModifier int64
	PsionicModulator int64
	ShipManufacturingHall int64
	SupraRefractor int64
}
func (LfBuildings) ByID(param1 ID) int64 { panic("not implemented") }
func (LfBuildings) GetAcademyOfSciences() int64 { panic("not implemented") }
func (LfBuildings) GetAdvancedRecyclingPlant() int64 { panic("not implemented") }
func (LfBuildings) GetAntimatterCondenser() int64 { panic("not implemented") }
func (LfBuildings) GetAntimatterConvector() int64 { panic("not implemented") }
func (LfBuildings) GetAssemblyLine() int64 { panic("not implemented") }
func (LfBuildings) GetAutomatisedAssemblyCentre() int64 { panic("not implemented") }
func (LfBuildings) GetBioModifier() int64 { panic("not implemented") }
func (LfBuildings) GetBiosphereFarm() int64 { panic("not implemented") }
func (LfBuildings) GetBiotechLab() int64 { panic("not implemented") }
func (LfBuildings) GetChipMassProduction() int64 { panic("not implemented") }
func (LfBuildings) GetChrysalisAccelerator() int64 { panic("not implemented") }
func (LfBuildings) GetCloningLaboratory() int64 { panic("not implemented") }
func (LfBuildings) GetCrystalFarm() int64 { panic("not implemented") }
func (LfBuildings) GetCrystalRefinery() int64 { panic("not implemented") }
func (LfBuildings) GetDeuteriumSynthesiser() int64 { panic("not implemented") }
func (LfBuildings) GetDisruptionChamber() int64 { panic("not implemented") }
func (LfBuildings) GetFoodSilo() int64 { panic("not implemented") }
func (LfBuildings) GetForumOfTranscendence() int64 { panic("not implemented") }
func (LfBuildings) GetFusionCellFactory() int64 { panic("not implemented") }
func (LfBuildings) GetFusionPoweredProduction() int64 { panic("not implemented") }
func (LfBuildings) GetHallsOfRealisation() int64 { panic("not implemented") }
func (LfBuildings) GetHighEnergySmelting() int64 { panic("not implemented") }
func (LfBuildings) GetHighPerformanceSynthesiser() int64 { panic("not implemented") }
func (LfBuildings) GetHighPerformanceTransformer() int64 { panic("not implemented") }
func (LfBuildings) GetMagmaForge() int64 { panic("not implemented") }
func (LfBuildings) GetMeditationEnclave() int64 { panic("not implemented") }
func (LfBuildings) GetMegalith() int64 { panic("not implemented") }
func (LfBuildings) GetMetropolis() int64 { panic("not implemented") }
func (LfBuildings) GetMicrochipAssemblyLine() int64 { panic("not implemented") }
func (LfBuildings) GetMineralResearchCentre() int64 { panic("not implemented") }
func (LfBuildings) GetNanoRepairBots() int64 { panic("not implemented") }
func (LfBuildings) GetNeuroCalibrationCentre() int64 { panic("not implemented") }
func (LfBuildings) GetOriktorium() int64 { panic("not implemented") }
func (LfBuildings) GetPlanetaryShield() int64 { panic("not implemented") }
func (LfBuildings) GetProductionAssemblyHall() int64 { panic("not implemented") }
func (LfBuildings) GetPsionicModulator() int64 { panic("not implemented") }
func (LfBuildings) GetQuantumComputerCentre() int64 { panic("not implemented") }
func (LfBuildings) GetResearchCentre() int64 { panic("not implemented") }
func (LfBuildings) GetResidentialSector() int64 { panic("not implemented") }
func (LfBuildings) GetRoboticsResearchCentre() int64 { panic("not implemented") }
func (LfBuildings) GetRuneForge() int64 { panic("not implemented") }
func (LfBuildings) GetRuneTechnologium() int64 { panic("not implemented") }
func (LfBuildings) GetSanctuary() int64 { panic("not implemented") }
func (LfBuildings) GetShipManufacturingHall() int64 { panic("not implemented") }
func (LfBuildings) GetSkyscraper() int64 { panic("not implemented") }
func (LfBuildings) GetSupraRefractor() int64 { panic("not implemented") }
func (LfBuildings) GetUpdateNetwork() int64 { panic("not implemented") }
func (LfBuildings) GetVortexChamber() int64 { panic("not implemented") }
func (LfBuildings) Lazy() LazyLfBuildings { panic("not implemented") }
type LfResearchDetails struct {
	LfResearches LfResearches
	Slots [18]LfSlot
	ArtefactsCollected int64
	ArtefactsLimit int64
}
func (LfResearchDetails) ByID(param1 ID) *int64 { panic("not implemented") }
func (LfResearchDetails) GetAcousticScanning() *int64 { panic("not implemented") }
func (LfResearchDetails) GetArtificialSwarmIntelligence() *int64 { panic("not implemented") }
func (LfResearchDetails) GetAutomatedTransportLines() *int64 { panic("not implemented") }
func (LfResearchDetails) GetBattlecruiserMkII() *int64 { panic("not implemented") }
func (LfResearchDetails) GetBomberMkII() *int64 { panic("not implemented") }
func (LfResearchDetails) GetCargoHoldExpansionCivilianShips() *int64 { panic("not implemented") }
func (LfResearchDetails) GetCatalyserTechnology() *int64 { panic("not implemented") }
func (LfResearchDetails) GetCruiserMkII() *int64 { panic("not implemented") }
func (LfResearchDetails) GetDepotAI() *int64 { panic("not implemented") }
func (LfResearchDetails) GetDepthSounding() *int64 { panic("not implemented") }
func (LfResearchDetails) GetDestroyerMkII() *int64 { panic("not implemented") }
func (LfResearchDetails) GetDiamondEnergyTransmitter() *int64 { panic("not implemented") }
func (LfResearchDetails) GetEfficiencyModule() *int64 { panic("not implemented") }
func (LfResearchDetails) GetEfficientSwarmIntelligence() *int64 { panic("not implemented") }
func (LfResearchDetails) GetEnhancedProductionTechnologies() *int64 { panic("not implemented") }
func (LfResearchDetails) GetEnhancedSensorTechnology() *int64 { panic("not implemented") }
func (LfResearchDetails) GetExperimentalRecyclingTechnology() *int64 { panic("not implemented") }
func (LfResearchDetails) GetExperimentalWeaponsTechnology() *int64 { panic("not implemented") }
func (LfResearchDetails) GetFusionDrives() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGeneralOverhaulBattlecruiser() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGeneralOverhaulBattleship() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGeneralOverhaulBomber() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGeneralOverhaulCruiser() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGeneralOverhaulDestroyer() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGeneralOverhaulLightFighter() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGeothermalPowerPlants() *int64 { panic("not implemented") }
func (LfResearchDetails) GetGravitationSensors() *int64 { panic("not implemented") }
func (LfResearchDetails) GetHardenedDiamondDrillHeads() *int64 { panic("not implemented") }
func (LfResearchDetails) GetHeatRecovery() *int64 { panic("not implemented") }
func (LfResearchDetails) GetHighEnergyPumpSystems() *int64 { panic("not implemented") }
func (LfResearchDetails) GetHighPerformanceExtractors() *int64 { panic("not implemented") }
func (LfResearchDetails) GetHighPerformanceTerraformer() *int64 { panic("not implemented") }
func (LfResearchDetails) GetHighTemperatureSuperconductors() *int64 { panic("not implemented") }
func (LfResearchDetails) GetImprovedDroneAI() *int64 { panic("not implemented") }
func (LfResearchDetails) GetImprovedLabTechnology() *int64 { panic("not implemented") }
func (LfResearchDetails) GetImprovedStellarator() *int64 { panic("not implemented") }
func (LfResearchDetails) GetIntergalacticEnvoys() *int64 { panic("not implemented") }
func (LfResearchDetails) GetInterplanetaryAnalysisNetwork() *int64 { panic("not implemented") }
func (LfResearchDetails) GetIonCrystalEnhancementHeavyFighter() *int64 { panic("not implemented") }
func (LfResearchDetails) GetIonCrystalModules() *int64 { panic("not implemented") }
func (LfResearchDetails) GetKaeleshDiscovererEnhancement() *int64 { panic("not implemented") }
func (LfResearchDetails) GetLightFighterMkII() *int64 { panic("not implemented") }
func (LfResearchDetails) GetLowTemperatureDrives() *int64 { panic("not implemented") }
func (LfResearchDetails) GetMagmaPoweredProduction() *int64 { panic("not implemented") }
func (LfResearchDetails) GetMagmaPoweredPumpSystems() *int64 { panic("not implemented") }
func (LfResearchDetails) GetMechanGeneralEnhancement() *int64 { panic("not implemented") }
func (LfResearchDetails) GetNeuroInterface() *int64 { panic("not implemented") }
func (LfResearchDetails) GetNeuromodalCompressor() *int64 { panic("not implemented") }
func (LfResearchDetails) GetObsidianShieldReinforcement() *int64 { panic("not implemented") }
func (LfResearchDetails) GetOptimisedSiloConstructionMethod() *int64 { panic("not implemented") }
func (LfResearchDetails) GetOrbitalDen() *int64 { panic("not implemented") }
func (LfResearchDetails) GetOverclockingBattleship() *int64 { panic("not implemented") }
func (LfResearchDetails) GetOverclockingHeavyFighter() *int64 { panic("not implemented") }
func (LfResearchDetails) GetOverclockingLargeCargo() *int64 { panic("not implemented") }
func (LfResearchDetails) GetPlasmaDrive() *int64 { panic("not implemented") }
func (LfResearchDetails) GetPlasmaTerraformer() *int64 { panic("not implemented") }
func (LfResearchDetails) GetPsionicNetwork() *int64 { panic("not implemented") }
func (LfResearchDetails) GetPsionicShieldMatrix() *int64 { panic("not implemented") }
func (LfResearchDetails) GetPsychoharmoniser() *int64 { panic("not implemented") }
func (LfResearchDetails) GetResearchAI() *int64 { panic("not implemented") }
func (LfResearchDetails) GetRobotAssistants() *int64 { panic("not implemented") }
func (LfResearchDetails) GetRocktalCollectorEnhancement() *int64 { panic("not implemented") }
func (LfResearchDetails) GetRuneShields() *int64 { panic("not implemented") }
func (LfResearchDetails) GetSeismicMiningTechnology() *int64 { panic("not implemented") }
func (LfResearchDetails) GetSixthSense() *int64 { panic("not implemented") }
func (LfResearchDetails) GetSlingshotAutopilot() *int64 { panic("not implemented") }
func (LfResearchDetails) GetStealthFieldGenerator() *int64 { panic("not implemented") }
func (LfResearchDetails) GetSulphideProcess() *int64 { panic("not implemented") }
func (LfResearchDetails) GetSupercomputer() *int64 { panic("not implemented") }
func (LfResearchDetails) GetTelekineticDrive() *int64 { panic("not implemented") }
func (LfResearchDetails) GetTelekineticTractorBeam() *int64 { panic("not implemented") }
func (LfResearchDetails) GetVolcanicBatteries() *int64 { panic("not implemented") }
func (LfResearchDetails) Lazy() LazyLfResearches { panic("not implemented") }
type LfResearches struct {
	IntergalacticEnvoys *int64
	HighPerformanceExtractors *int64
	FusionDrives *int64
	StealthFieldGenerator *int64
	OrbitalDen *int64
	ResearchAI *int64
	HighPerformanceTerraformer *int64
	EnhancedProductionTechnologies *int64
	LightFighterMkII *int64
	CruiserMkII *int64
	ImprovedLabTechnology *int64
	PlasmaTerraformer *int64
	LowTemperatureDrives *int64
	BomberMkII *int64
	DestroyerMkII *int64
	BattlecruiserMkII *int64
	RobotAssistants *int64
	Supercomputer *int64
	VolcanicBatteries *int64
	AcousticScanning *int64
	HighEnergyPumpSystems *int64
	CargoHoldExpansionCivilianShips *int64
	MagmaPoweredProduction *int64
	GeothermalPowerPlants *int64
	DepthSounding *int64
	IonCrystalEnhancementHeavyFighter *int64
	ImprovedStellarator *int64
	HardenedDiamondDrillHeads *int64
	SeismicMiningTechnology *int64
	MagmaPoweredPumpSystems *int64
	IonCrystalModules *int64
	OptimisedSiloConstructionMethod *int64
	DiamondEnergyTransmitter *int64
	ObsidianShieldReinforcement *int64
	RuneShields *int64
	RocktalCollectorEnhancement *int64
	CatalyserTechnology *int64
	PlasmaDrive *int64
	EfficiencyModule *int64
	DepotAI *int64
	GeneralOverhaulLightFighter *int64
	AutomatedTransportLines *int64
	ImprovedDroneAI *int64
	ExperimentalRecyclingTechnology *int64
	GeneralOverhaulCruiser *int64
	SlingshotAutopilot *int64
	HighTemperatureSuperconductors *int64
	GeneralOverhaulBattleship *int64
	ArtificialSwarmIntelligence *int64
	GeneralOverhaulBattlecruiser *int64
	GeneralOverhaulBomber *int64
	GeneralOverhaulDestroyer *int64
	ExperimentalWeaponsTechnology *int64
	MechanGeneralEnhancement *int64
	HeatRecovery *int64
	SulphideProcess *int64
	PsionicNetwork *int64
	TelekineticTractorBeam *int64
	EnhancedSensorTechnology *int64
	NeuromodalCompressor *int64
	NeuroInterface *int64
	InterplanetaryAnalysisNetwork *int64
	OverclockingHeavyFighter *int64
	TelekineticDrive *int64
	SixthSense *int64
	Psychoharmoniser *int64
	EfficientSwarmIntelligence *int64
	OverclockingLargeCargo *int64
	GravitationSensors *int64
	OverclockingBattleship *int64
	PsionicShieldMatrix *int64
	KaeleshDiscovererEnhancement *int64
}
func (LfResearches) ByID(param1 ID) *int64 { panic("not implemented") }
func (LfResearches) GetAcousticScanning() *int64 { panic("not implemented") }
func (LfResearches) GetArtificialSwarmIntelligence() *int64 { panic("not implemented") }
func (LfResearches) GetAutomatedTransportLines() *int64 { panic("not implemented") }
func (LfResearches) GetBattlecruiserMkII() *int64 { panic("not implemented") }
func (LfResearches) GetBomberMkII() *int64 { panic("not implemented") }
func (LfResearches) GetCargoHoldExpansionCivilianShips() *int64 { panic("not implemented") }
func (LfResearches) GetCatalyserTechnology() *int64 { panic("not implemented") }
func (LfResearches) GetCruiserMkII() *int64 { panic("not implemented") }
func (LfResearches) GetDepotAI() *int64 { panic("not implemented") }
func (LfResearches) GetDepthSounding() *int64 { panic("not implemented") }
func (LfResearches) GetDestroyerMkII() *int64 { panic("not implemented") }
func (LfResearches) GetDiamondEnergyTransmitter() *int64 { panic("not implemented") }
func (LfResearches) GetEfficiencyModule() *int64 { panic("not implemented") }
func (LfResearches) GetEfficientSwarmIntelligence() *int64 { panic("not implemented") }
func (LfResearches) GetEnhancedProductionTechnologies() *int64 { panic("not implemented") }
func (LfResearches) GetEnhancedSensorTechnology() *int64 { panic("not implemented") }
func (LfResearches) GetExperimentalRecyclingTechnology() *int64 { panic("not implemented") }
func (LfResearches) GetExperimentalWeaponsTechnology() *int64 { panic("not implemented") }
func (LfResearches) GetFusionDrives() *int64 { panic("not implemented") }
func (LfResearches) GetGeneralOverhaulBattlecruiser() *int64 { panic("not implemented") }
func (LfResearches) GetGeneralOverhaulBattleship() *int64 { panic("not implemented") }
func (LfResearches) GetGeneralOverhaulBomber() *int64 { panic("not implemented") }
func (LfResearches) GetGeneralOverhaulCruiser() *int64 { panic("not implemented") }
func (LfResearches) GetGeneralOverhaulDestroyer() *int64 { panic("not implemented") }
func (LfResearches) GetGeneralOverhaulLightFighter() *int64 { panic("not implemented") }
func (LfResearches) GetGeothermalPowerPlants() *int64 { panic("not implemented") }
func (LfResearches) GetGravitationSensors() *int64 { panic("not implemented") }
func (LfResearches) GetHardenedDiamondDrillHeads() *int64 { panic("not implemented") }
func (LfResearches) GetHeatRecovery() *int64 { panic("not implemented") }
func (LfResearches) GetHighEnergyPumpSystems() *int64 { panic("not implemented") }
func (LfResearches) GetHighPerformanceExtractors() *int64 { panic("not implemented") }
func (LfResearches) GetHighPerformanceTerraformer() *int64 { panic("not implemented") }
func (LfResearches) GetHighTemperatureSuperconductors() *int64 { panic("not implemented") }
func (LfResearches) GetImprovedDroneAI() *int64 { panic("not implemented") }
func (LfResearches) GetImprovedLabTechnology() *int64 { panic("not implemented") }
func (LfResearches) GetImprovedStellarator() *int64 { panic("not implemented") }
func (LfResearches) GetIntergalacticEnvoys() *int64 { panic("not implemented") }
func (LfResearches) GetInterplanetaryAnalysisNetwork() *int64 { panic("not implemented") }
func (LfResearches) GetIonCrystalEnhancementHeavyFighter() *int64 { panic("not implemented") }
func (LfResearches) GetIonCrystalModules() *int64 { panic("not implemented") }
func (LfResearches) GetKaeleshDiscovererEnhancement() *int64 { panic("not implemented") }
func (LfResearches) GetLightFighterMkII() *int64 { panic("not implemented") }
func (LfResearches) GetLowTemperatureDrives() *int64 { panic("not implemented") }
func (LfResearches) GetMagmaPoweredProduction() *int64 { panic("not implemented") }
func (LfResearches) GetMagmaPoweredPumpSystems() *int64 { panic("not implemented") }
func (LfResearches) GetMechanGeneralEnhancement() *int64 { panic("not implemented") }
func (LfResearches) GetNeuroInterface() *int64 { panic("not implemented") }
func (LfResearches) GetNeuromodalCompressor() *int64 { panic("not implemented") }
func (LfResearches) GetObsidianShieldReinforcement() *int64 { panic("not implemented") }
func (LfResearches) GetOptimisedSiloConstructionMethod() *int64 { panic("not implemented") }
func (LfResearches) GetOrbitalDen() *int64 { panic("not implemented") }
func (LfResearches) GetOverclockingBattleship() *int64 { panic("not implemented") }
func (LfResearches) GetOverclockingHeavyFighter() *int64 { panic("not implemented") }
func (LfResearches) GetOverclockingLargeCargo() *int64 { panic("not implemented") }
func (LfResearches) GetPlasmaDrive() *int64 { panic("not implemented") }
func (LfResearches) GetPlasmaTerraformer() *int64 { panic("not implemented") }
func (LfResearches) GetPsionicNetwork() *int64 { panic("not implemented") }
func (LfResearches) GetPsionicShieldMatrix() *int64 { panic("not implemented") }
func (LfResearches) GetPsychoharmoniser() *int64 { panic("not implemented") }
func (LfResearches) GetResearchAI() *int64 { panic("not implemented") }
func (LfResearches) GetRobotAssistants() *int64 { panic("not implemented") }
func (LfResearches) GetRocktalCollectorEnhancement() *int64 { panic("not implemented") }
func (LfResearches) GetRuneShields() *int64 { panic("not implemented") }
func (LfResearches) GetSeismicMiningTechnology() *int64 { panic("not implemented") }
func (LfResearches) GetSixthSense() *int64 { panic("not implemented") }
func (LfResearches) GetSlingshotAutopilot() *int64 { panic("not implemented") }
func (LfResearches) GetStealthFieldGenerator() *int64 { panic("not implemented") }
func (LfResearches) GetSulphideProcess() *int64 { panic("not implemented") }
func (LfResearches) GetSupercomputer() *int64 { panic("not implemented") }
func (LfResearches) GetTelekineticDrive() *int64 { panic("not implemented") }
func (LfResearches) GetTelekineticTractorBeam() *int64 { panic("not implemented") }
func (LfResearches) GetVolcanicBatteries() *int64 { panic("not implemented") }
func (LfResearches) Lazy() LazyLfResearches { panic("not implemented") }
type LfResourceBonuses struct {
}
type LfShipBonus struct {
	ID ID
	StructuralIntegrity float64
	ShieldPower float64
	WeaponPower float64
	Speed float64
	CargoCapacity float64
	FuelConsumption float64
}
type LfShipBonuses map[ID]LfShipBonus
type LfSlot struct {
	TechID ID
	Level int64
	Allowed bool
	Locked bool
}
type LifeformType int64
type MarketplaceMessage struct {
	ID int64
	Type int64
	CreatedAt time.Time
	Token string
	MarketTransactionID int64
}
type MessagesTabID int64
type MissionID int
func (MissionID) String() string { panic("not implemented") }
type Moon struct {
	ID MoonID
	Img string
	Name string
	Diameter int64
	Coordinate Coordinate
	Fields Fields
}
func (Moon) GetCoordinate() Coordinate { panic("not implemented") }
func (Moon) GetDiameter() int64 { panic("not implemented") }
func (Moon) GetFields() Fields { panic("not implemented") }
func (Moon) GetID() CelestialID { panic("not implemented") }
func (Moon) GetImg() string { panic("not implemented") }
func (Moon) GetName() string { panic("not implemented") }
func (Moon) GetType() CelestialType { panic("not implemented") }
type MoonID int64
func (MoonID) Celestial() CelestialID { panic("not implemented") }
type MoonInfos struct {
	ID int64
	Name string
	Diameter int64
	Activity int64
}
type PhalanxFleet struct {
	Fleet Fleet
	BaseSpeed int64
}
type Planet struct {
	Img string
	ID PlanetID
	Name string
	Diameter int64
	Coordinate Coordinate
	Fields Fields
	Temperature Temperature
	Moon *Moon
}
func (Planet) GetCoordinate() Coordinate { panic("not implemented") }
func (Planet) GetDiameter() int64 { panic("not implemented") }
func (Planet) GetFields() Fields { panic("not implemented") }
func (Planet) GetID() CelestialID { panic("not implemented") }
func (Planet) GetImg() string { panic("not implemented") }
func (Planet) GetMoon() *Moon { panic("not implemented") }
func (Planet) GetName() string { panic("not implemented") }
func (Planet) GetTemperature() Temperature { panic("not implemented") }
func (Planet) GetType() CelestialType { panic("not implemented") }
type PlanetID int64
func (PlanetID) Celestial() CelestialID { panic("not implemented") }
func (PlanetID) String() string { panic("not implemented") }
type PlanetInfos struct {
	ID int64
	Activity int64
	Name string
	Img string
	Coordinate Coordinate
	Administrator bool
	Destroyed bool
	Inactive bool
	Vacation bool
	StrongPlayer bool
	Newbie bool
	HonorableTarget bool
	Banned bool
	Debris struct { Metal int64; Crystal int64; Deuterium int64; RecyclersNeeded int64 }
	Moon *MoonInfos
	Player struct { ID int64; Name string; Rank int64; IsBandit bool; IsStarlord bool }
	Alliance *AllianceInfos
}
type Preferences struct {
	SpioAnz int64
	SpySystemAutomaticQuantity int64
	SpySystemTargetPlanetTypes int64
	SpySystemTargetPlayerTypes int64
	SpySystemIgnoreSpiedInLastXMinutes int64
	DisableChatBar bool
	DisableOutlawWarning bool
	MobileVersion bool
	ShowOldDropDowns bool
	ActivateAutofocus bool
	Language string
	EventsShow int64
	SortSetting int64
	SortOrder int64
	ShowDetailOverlay bool
	AnimatedSliders bool
	AnimatedOverview bool
	PopupsNotices bool
	PopopsCombatreport bool
	SpioReportPictures bool
	MsgResultsPerPage int64
	AuctioneerNotifications bool
	EconomyNotifications bool
	ShowActivityMinutes bool
	PreserveSystemOnPlanetChange bool
	DiscoveryWarningEnabled bool
	UrlaubsModus bool
	Notifications struct { BuildList bool; FriendlyFleetActivities bool; HostileFleetActivities bool; ForeignEspionage bool; AllianceBroadcasts bool; AllianceMessages bool; Auctions bool; Account bool }
}
type Quantifiable struct {
	ID ID
	Nbr int64
}
type Researches struct {
	EnergyTechnology int64
	LaserTechnology int64
	IonTechnology int64
	HyperspaceTechnology int64
	PlasmaTechnology int64
	CombustionDrive int64
	ImpulseDrive int64
	HyperspaceDrive int64
	EspionageTechnology int64
	ComputerTechnology int64
	Astrophysics int64
	IntergalacticResearchNetwork int64
	GravitonTechnology int64
	WeaponsTechnology int64
	ShieldingTechnology int64
	ArmourTechnology int64
}
func (Researches) ByID(param1 ID) int64 { panic("not implemented") }
func (Researches) GetArmourTechnology() int64 { panic("not implemented") }
func (Researches) GetAstrophysics() int64 { panic("not implemented") }
func (Researches) GetCombustionDrive() int64 { panic("not implemented") }
func (Researches) GetComputerTechnology() int64 { panic("not implemented") }
func (Researches) GetEnergyTechnology() int64 { panic("not implemented") }
func (Researches) GetEspionageTechnology() int64 { panic("not implemented") }
func (Researches) GetGravitonTechnology() int64 { panic("not implemented") }
func (Researches) GetHyperspaceDrive() int64 { panic("not implemented") }
func (Researches) GetHyperspaceTechnology() int64 { panic("not implemented") }
func (Researches) GetImpulseDrive() int64 { panic("not implemented") }
func (Researches) GetIntergalacticResearchNetwork() int64 { panic("not implemented") }
func (Researches) GetIonTechnology() int64 { panic("not implemented") }
func (Researches) GetLaserTechnology() int64 { panic("not implemented") }
func (Researches) GetPlasmaTechnology() int64 { panic("not implemented") }
func (Researches) GetShieldingTechnology() int64 { panic("not implemented") }
func (Researches) GetWeaponsTechnology() int64 { panic("not implemented") }
func (Researches) Lazy() LazyResearches { panic("not implemented") }
func (Researches) String() string { panic("not implemented") }
func (Researches) ToPtr() *Researches { panic("not implemented") }
type ResourceSettings struct {
	MetalMine int64
	CrystalMine int64
	DeuteriumSynthesizer int64
	SolarPlant int64
	FusionReactor int64
	SolarSatellite int64
	Crawler int64
}
func (ResourceSettings) String() string { panic("not implemented") }
type Resources struct {
	Metal int64
	Crystal int64
	Deuterium int64
	Energy int64
	Darkmatter int64
	Population int64
	Food int64
}
func (Resources) Add(param1 Resources) Resources { panic("not implemented") }
func (Resources) CanAfford(param1 Resources) bool { panic("not implemented") }
func (Resources) CanAfford2(param1 Resources, param2 bool) bool { panic("not implemented") }
func (Resources) Div(param1 Resources) int64 { panic("not implemented") }
func (Resources) FitsIn(param1 Ship, param2 Researches, param3 LfBonuses, param4 CharacterClass, param5 float64, param6 bool) int64 { panic("not implemented") }
func (Resources) Gte(param1 Resources) bool { panic("not implemented") }
func (Resources) Gte2(param1 Resources, param2 bool) bool { panic("not implemented") }
func (Resources) Lte(param1 Resources) bool { panic("not implemented") }
func (Resources) Mul(param1 int64) Resources { panic("not implemented") }
func (Resources) String() string { panic("not implemented") }
func (Resources) Sub(param1 Resources) Resources { panic("not implemented") }
func (Resources) SubPercent(param1 float64) Resources { panic("not implemented") }
func (Resources) Total() int64 { panic("not implemented") }
func (Resources) Value() int64 { panic("not implemented") }
type ResourcesBuildings struct {
	MetalMine int64
	CrystalMine int64
	DeuteriumSynthesizer int64
	SolarPlant int64
	FusionReactor int64
	SolarSatellite int64
	MetalStorage int64
	CrystalStorage int64
	DeuteriumTank int64
}
func (ResourcesBuildings) ByID(param1 ID) int64 { panic("not implemented") }
func (ResourcesBuildings) GetCrystalMine() int64 { panic("not implemented") }
func (ResourcesBuildings) GetCrystalStorage() int64 { panic("not implemented") }
func (ResourcesBuildings) GetDeuteriumSynthesizer() int64 { panic("not implemented") }
func (ResourcesBuildings) GetDeuteriumTank() int64 { panic("not implemented") }
func (ResourcesBuildings) GetFusionReactor() int64 { panic("not implemented") }
func (ResourcesBuildings) GetMetalMine() int64 { panic("not implemented") }
func (ResourcesBuildings) GetMetalStorage() int64 { panic("not implemented") }
func (ResourcesBuildings) GetSolarPlant() int64 { panic("not implemented") }
func (ResourcesBuildings) GetSolarSatellite() int64 { panic("not implemented") }
func (ResourcesBuildings) Lazy() LazyResourcesBuildings { panic("not implemented") }
func (ResourcesBuildings) String() string { panic("not implemented") }
type ResourcesDetails struct {
	Metal struct { Available int64; StorageCapacity int64; CurrentProduction int64 }
	Crystal struct { Available int64; StorageCapacity int64; CurrentProduction int64 }
	Deuterium struct { Available int64; StorageCapacity int64; CurrentProduction int64 }
	Food struct { Available int64; StorageCapacity int64; Overproduction int64; ConsumedIn int64; TimeTillFoodRunsOut int64 }
	Population struct { Available int64; T2Lifeforms int64; T3Lifeforms int64; LivingSpace int64; Satisfied int64; Hungry float64; GrowthRate float64; BunkerSpace int64 }
	Energy struct { Available int64; CurrentProduction int64; Consumption int64 }
	Darkmatter struct { Available int64; Purchased int64; Found int64 }
}
func (ResourcesDetails) Available() Resources { panic("not implemented") }
type Ship interface {
	ConstructionTime(int64, int64, BuildAccelerators, LfBonuses, CharacterClass, bool) time.Duration
	DefenderConstructionTime(int64, int64, DefenseAccelerators, LfBonuses) time.Duration
	GetCargoCapacity(IResearches, LfBonuses, CharacterClass, float64, bool) int64
	GetFuelConsumption(IResearches, LfBonuses, CharacterClass, float64) int64
	GetID() ID
	GetName() string
	GetPrice(int64, LfBonuses) Resources
	GetRapidfireAgainst() map[ID]int64
	GetRapidfireFrom() map[ID]int64
	GetRequirements() map[ID]int64
	GetShieldPower(IResearches) int64
	GetSpeed(IResearches, LfBonuses, CharacterClass, AllianceClass) int64
	GetStructuralIntegrity(IResearches) int64
	GetWeaponPower(IResearches) int64
	IsAvailable(CelestialType, IResourcesBuildings, ILfBuildings, ILfResearches, IFacilities, IResearches, int64, CharacterClass) bool
}
type ShipsInfos struct {
	LightFighter int64
	HeavyFighter int64
	Cruiser int64
	Battleship int64
	Battlecruiser int64
	Bomber int64
	Destroyer int64
	Deathstar int64
	SmallCargo int64
	LargeCargo int64
	ColonyShip int64
	Recycler int64
	EspionageProbe int64
	SolarSatellite int64
	Crawler int64
	Reaper int64
	Pathfinder int64
}
func (ShipsInfos) ByID(param1 ID) int64 { panic("not implemented") }
func (ShipsInfos) Cargo(param1 IResearches, param2 LfBonuses, param3 CharacterClass, param4 float64, param5 bool) int64 { panic("not implemented") }
func (ShipsInfos) CountShips() int64 { panic("not implemented") }
func (ShipsInfos) Each(param1 func(ID, int64))  { panic("not implemented") }
func (ShipsInfos) EachFlyable(param1 func(ID, int64))  { panic("not implemented") }
func (ShipsInfos) Equal(param1 ShipsInfos) bool { panic("not implemented") }
func (ShipsInfos) FleetCost(param1 LfBonuses) Resources { panic("not implemented") }
func (ShipsInfos) FleetValue(param1 LfBonuses) int64 { panic("not implemented") }
func (ShipsInfos) FromQuantifiables(param1 []Quantifiable) ShipsInfos { panic("not implemented") }
func (ShipsInfos) Has(param1 ShipsInfos) bool { panic("not implemented") }
func (ShipsInfos) HasFlyableShips() bool { panic("not implemented") }
func (ShipsInfos) HasShips() bool { panic("not implemented") }
func (ShipsInfos) IsEmpty() bool { panic("not implemented") }
func (ShipsInfos) Speed(param1 IResearches, param2 LfBonuses, param3 CharacterClass, param4 AllianceClass) int64 { panic("not implemented") }
func (ShipsInfos) String() string { panic("not implemented") }
func (ShipsInfos) ToPtr() *ShipsInfos { panic("not implemented") }
func (ShipsInfos) ToQuantifiables() []Quantifiable { panic("not implemented") }
type Slots struct {
	InUse int64
	Total int64
	ExpInUse int64
	ExpTotal int64
}
func (Slots) IsAllSlotsInUse(param1 MissionID) bool { panic("not implemented") }
type Speed float64
func (Speed) Float64() float64 { panic("not implemented") }
func (Speed) Int() int64 { panic("not implemented") }
func (Speed) Int64() int64 { panic("not implemented") }
func (Speed) String() string { panic("not implemented") }
type SystemInfos struct {
	ExpeditionDebris struct { Metal int64; Crystal int64; Deuterium int64; PathfindersNeeded int64 }
	Events struct { Darkmatter int64; HasAsteroid bool }
	OverlayToken string
}
func (SystemInfos) Each(param1 func(*PlanetInfos))  { panic("not implemented") }
func (SystemInfos) Galaxy() int64 { panic("not implemented") }
func (SystemInfos) MarshalJSON() ([]uint8, error) { panic("not implemented") }
func (SystemInfos) Position(param1 int64) *PlanetInfos { panic("not implemented") }
func (SystemInfos) System() int64 { panic("not implemented") }
type TechAccelerators interface {
	GetResearchLab() int64
}
type TechnologyDetails struct {
	TechnologyID ID
	ProductionDuration time.Duration
	Price Resources
	Level int64
	TearDownEnabled bool
}
type Temperature struct {
	Min int64
	Max int64
}
func (Temperature) Mean() int64 { panic("not implemented") }
type UserInfos struct {
	PlayerID int64
	PlayerName string
	Points int64
	Rank int64
	Total int64
	HonourPoints int64
}
