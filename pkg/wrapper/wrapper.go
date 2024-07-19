// Code generated DO NOT EDIT. (@generated)

package wrapper

import (
	"net/http"
	"net/url"
	"sync"
	"time"
	"nja/pkg/ogame"
)
type Celestial interface {
	ActivateItem(string) error
	Build(ogame.ID, int64) error
	BuildBuilding(ogame.ID) error
	BuildDefense(ogame.ID, int64) error
	BuildTechnology(ogame.ID) error
	CancelBuilding() error
	CancelLfBuilding() error
	CancelResearch() error
	ConstructionsBeingBuilt() (ogame.ID, int64, ogame.ID, int64, ogame.ID, int64, ogame.ID, int64)
	EnsureFleet(ogame.ShipsInfos, ogame.Speed, ogame.Coordinate, ogame.MissionID, ogame.Resources, int64, int64) (ogame.Fleet, error)
	GetCoordinate() ogame.Coordinate
	GetDefense(...Option) (ogame.DefensesInfos, error)
	GetDiameter() int64
	GetFacilities(...Option) (ogame.Facilities, error)
	GetFields() ogame.Fields
	GetID() ogame.CelestialID
	GetImg() string
	GetItems() ([]ogame.Item, error)
	GetLfBuildings(...Option) (ogame.LfBuildings, error)
	GetLfResearch(...Option) (ogame.LfResearches, error)
	GetName() string
	GetProduction() ([]ogame.Quantifiable, int64, error)
	GetResources() (ogame.Resources, error)
	GetResourcesBuildings(...Option) (ogame.ResourcesBuildings, error)
	GetResourcesDetails() (ogame.ResourcesDetails, error)
	GetShips(...Option) (ogame.ShipsInfos, error)
	GetTechs() (ogame.ResourcesBuildings, ogame.Facilities, ogame.ShipsInfos, ogame.DefensesInfos, ogame.Researches, ogame.LfBuildings, ogame.LfResearches, error)
	GetType() ogame.CelestialType
	SendFleet(ogame.ShipsInfos, ogame.Speed, ogame.Coordinate, ogame.MissionID, ogame.Resources, int64, int64) (ogame.Fleet, error)
	TearDown(ogame.ID) error
}
type FleetBuilder struct {
}
type IntoCelestial interface {
	
}
type IntoCoordinate interface {
	
}
type IntoMoon interface {
	
}
type IntoPlanet interface {
	
}
type Moon struct {
	Moon ogame.Moon
}
func (Moon) ActivateItem(param1 string) error { panic("not implemented") }
func (Moon) Build(param1 ogame.ID, param2 int64) error { panic("not implemented") }
func (Moon) BuildBuilding(param1 ogame.ID) error { panic("not implemented") }
func (Moon) BuildDefense(param1 ogame.ID, param2 int64) error { panic("not implemented") }
func (Moon) BuildTechnology(param1 ogame.ID) error { panic("not implemented") }
func (Moon) CancelBuilding() error { panic("not implemented") }
func (Moon) CancelLfBuilding() error { panic("not implemented") }
func (Moon) CancelResearch() error { panic("not implemented") }
func (Moon) ConstructionsBeingBuilt() (ogame.ID, int64, ogame.ID, int64, ogame.ID, int64, ogame.ID, int64) { panic("not implemented") }
func (Moon) EnsureFleet(param1 ogame.ShipsInfos, param2 ogame.Speed, param3 ogame.Coordinate, param4 ogame.MissionID, param5 ogame.Resources, param6 int64, param7 int64) (ogame.Fleet, error) { panic("not implemented") }
func (Moon) GetCoordinate() ogame.Coordinate { panic("not implemented") }
func (Moon) GetDefense(param1 ...Option) (ogame.DefensesInfos, error) { panic("not implemented") }
func (Moon) GetDiameter() int64 { panic("not implemented") }
func (Moon) GetFacilities(param1 ...Option) (ogame.Facilities, error) { panic("not implemented") }
func (Moon) GetFields() ogame.Fields { panic("not implemented") }
func (Moon) GetID() ogame.CelestialID { panic("not implemented") }
func (Moon) GetImg() string { panic("not implemented") }
func (Moon) GetItems() ([]ogame.Item, error) { panic("not implemented") }
func (Moon) GetLfBuildings(param1 ...Option) (ogame.LfBuildings, error) { panic("not implemented") }
func (Moon) GetLfResearch(param1 ...Option) (ogame.LfResearches, error) { panic("not implemented") }
func (Moon) GetName() string { panic("not implemented") }
func (Moon) GetProduction() ([]ogame.Quantifiable, int64, error) { panic("not implemented") }
func (Moon) GetResources() (ogame.Resources, error) { panic("not implemented") }
func (Moon) GetResourcesBuildings(param1 ...Option) (ogame.ResourcesBuildings, error) { panic("not implemented") }
func (Moon) GetResourcesDetails() (ogame.ResourcesDetails, error) { panic("not implemented") }
func (Moon) GetShips(param1 ...Option) (ogame.ShipsInfos, error) { panic("not implemented") }
func (Moon) GetTechs() (ogame.ResourcesBuildings, ogame.Facilities, ogame.ShipsInfos, ogame.DefensesInfos, ogame.Researches, ogame.LfBuildings, ogame.LfResearches, error) { panic("not implemented") }
func (Moon) GetType() ogame.CelestialType { panic("not implemented") }
func (Moon) Phalanx(param1 ogame.Coordinate) ([]ogame.PhalanxFleet, error) { panic("not implemented") }
func (Moon) SendFleet(param1 ogame.ShipsInfos, param2 ogame.Speed, param3 ogame.Coordinate, param4 ogame.MissionID, param5 ogame.Resources, param6 int64, param7 int64) (ogame.Fleet, error) { panic("not implemented") }
func (Moon) TearDown(param1 ogame.ID) error { panic("not implemented") }
type OGame struct {
	Mutex sync.Mutex
	Player ogame.UserInfos
	CachedPreferences ogame.Preferences
	Universe string
	Username string
}
type Option func(*Options)
type Options struct {
	DebugGalaxy bool
	SkipInterceptor bool
	SkipRetry bool
	SkipCacheFullPage bool
	ChangePlanet ogame.CelestialID
	Delay time.Duration
}
type Planet struct {
	Planet ogame.Planet
	Moon *Moon
}
func (Planet) ActivateItem(param1 string) error { panic("not implemented") }
func (Planet) Build(param1 ogame.ID, param2 int64) error { panic("not implemented") }
func (Planet) BuildBuilding(param1 ogame.ID) error { panic("not implemented") }
func (Planet) BuildCancelable(param1 ogame.ID) error { panic("not implemented") }
func (Planet) BuildDefense(param1 ogame.ID, param2 int64) error { panic("not implemented") }
func (Planet) BuildShips(param1 ogame.ID, param2 int64) error { panic("not implemented") }
func (Planet) BuildTechnology(param1 ogame.ID) error { panic("not implemented") }
func (Planet) CancelBuilding() error { panic("not implemented") }
func (Planet) CancelLfBuilding() error { panic("not implemented") }
func (Planet) CancelResearch() error { panic("not implemented") }
func (Planet) ConstructionsBeingBuilt() (ogame.ID, int64, ogame.ID, int64, ogame.ID, int64, ogame.ID, int64) { panic("not implemented") }
func (Planet) EnsureFleet(param1 ogame.ShipsInfos, param2 ogame.Speed, param3 ogame.Coordinate, param4 ogame.MissionID, param5 ogame.Resources, param6 int64, param7 int64) (ogame.Fleet, error) { panic("not implemented") }
func (Planet) FlightTime(param1 ogame.Coordinate, param2 ogame.Speed, param3 ogame.ShipsInfos, param4 ogame.MissionID) (int64, int64) { panic("not implemented") }
func (Planet) GetCoordinate() ogame.Coordinate { panic("not implemented") }
func (Planet) GetDefense(param1 ...Option) (ogame.DefensesInfos, error) { panic("not implemented") }
func (Planet) GetDiameter() int64 { panic("not implemented") }
func (Planet) GetFacilities(param1 ...Option) (ogame.Facilities, error) { panic("not implemented") }
func (Planet) GetFields() ogame.Fields { panic("not implemented") }
func (Planet) GetID() ogame.CelestialID { panic("not implemented") }
func (Planet) GetImg() string { panic("not implemented") }
func (Planet) GetItems() ([]ogame.Item, error) { panic("not implemented") }
func (Planet) GetLfBuildings(param1 ...Option) (ogame.LfBuildings, error) { panic("not implemented") }
func (Planet) GetLfResearch(param1 ...Option) (ogame.LfResearches, error) { panic("not implemented") }
func (Planet) GetMoon() *Moon { panic("not implemented") }
func (Planet) GetName() string { panic("not implemented") }
func (Planet) GetProduction() ([]ogame.Quantifiable, int64, error) { panic("not implemented") }
func (Planet) GetResourceSettings(param1 ...Option) (ogame.ResourceSettings, error) { panic("not implemented") }
func (Planet) GetResources() (ogame.Resources, error) { panic("not implemented") }
func (Planet) GetResourcesBuildings(param1 ...Option) (ogame.ResourcesBuildings, error) { panic("not implemented") }
func (Planet) GetResourcesDetails() (ogame.ResourcesDetails, error) { panic("not implemented") }
func (Planet) GetResourcesProductions() (ogame.Resources, error) { panic("not implemented") }
func (Planet) GetShips(param1 ...Option) (ogame.ShipsInfos, error) { panic("not implemented") }
func (Planet) GetTechs() (ogame.ResourcesBuildings, ogame.Facilities, ogame.ShipsInfos, ogame.DefensesInfos, ogame.Researches, ogame.LfBuildings, ogame.LfResearches, error) { panic("not implemented") }
func (Planet) GetTemperature() ogame.Temperature { panic("not implemented") }
func (Planet) GetType() ogame.CelestialType { panic("not implemented") }
func (Planet) SendFleet(param1 ogame.ShipsInfos, param2 ogame.Speed, param3 ogame.Coordinate, param4 ogame.MissionID, param5 ogame.Resources, param6 int64, param7 int64) (ogame.Fleet, error) { panic("not implemented") }
func (Planet) SendIPM(param1 ogame.PlanetID, param2 ogame.Coordinate, param3 int64, param4 ogame.ID) (int64, error) { panic("not implemented") }
func (Planet) String() string { panic("not implemented") }
func (Planet) TearDown(param1 ogame.ID) error { panic("not implemented") }
type Prioritizable interface {
	Abandon(IntoPlanet) error
	ActivateItem(string, ogame.CelestialID) error
	Begin() Prioritizable
	BeginNamed(string) Prioritizable
	Build(ogame.CelestialID, ogame.ID, int64) error
	BuildBuilding(ogame.CelestialID, ogame.ID) error
	BuildCancelable(ogame.CelestialID, ogame.ID) error
	BuildDefense(ogame.CelestialID, ogame.ID, int64) error
	BuildProduction(ogame.CelestialID, ogame.ID, int64) error
	BuildShips(ogame.CelestialID, ogame.ID, int64) error
	BuildTechnology(ogame.CelestialID, ogame.ID) error
	BuyMarketplace(int64, ogame.CelestialID) error
	BuyOfferOfTheDay() error
	BuyResetTree(ogame.PlanetID, int64) error
	CancelBuilding(ogame.CelestialID) error
	CancelFleet(ogame.FleetID) error
	CancelLfBuilding(ogame.CelestialID) error
	CancelResearch(ogame.CelestialID) error
	CollectAllMarketplaceMessages() error
	CollectMarketplaceMessage(ogame.MarketplaceMessage) error
	ConstructionsBeingBuilt(ogame.CelestialID) (ogame.ID, int64, ogame.ID, int64, ogame.ID, int64, ogame.ID, int64)
	CreateUnion(ogame.Fleet, []string) (int64, error)
	DeleteAllMessagesFromTab(ogame.MessagesTabID) error
	DeleteMessage(int64) error
	DestroyRockets(ogame.PlanetID, int64, int64) error
	DoAuction(map[ogame.CelestialID]ogame.Resources) error
	Done()
	EnsureFleet(ogame.CelestialID, ogame.ShipsInfos, ogame.Speed, ogame.Coordinate, ogame.MissionID, ogame.Resources, int64, int64) (ogame.Fleet, error)
	FlightTime(ogame.Coordinate, ogame.Coordinate, ogame.Speed, ogame.ShipsInfos, ogame.MissionID) (int64, int64)
	FreeResetTree(ogame.PlanetID, int64) error
	GalaxyInfos(int64, int64, ...Option) (ogame.SystemInfos, error)
	GetActiveItems(ogame.CelestialID) ([]ogame.ActiveItem, error)
	GetAllResources() (map[ogame.CelestialID]ogame.Resources, error)
	GetAttacks(...Option) ([]ogame.AttackEvent, error)
	GetAuction() (ogame.Auction, error)
	GetAvailableDiscoveries(...Option) int64
	GetCachedAllianceClass() (ogame.AllianceClass, error)
	GetCachedLfBonuses() (ogame.LfBonuses, error)
	GetCachedResearch() ogame.Researches
	GetCelestial(IntoCelestial) (Celestial, error)
	GetCelestials() ([]Celestial, error)
	GetCombatReportSummaryFor(ogame.Coordinate) (ogame.CombatReportSummary, error)
	GetDMCosts(ogame.CelestialID) (ogame.DMCosts, error)
	GetDefense(ogame.CelestialID, ...Option) (ogame.DefensesInfos, error)
	GetEmpire(ogame.CelestialType) ([]ogame.EmpireCelestial, error)
	GetEmpireJSON(ogame.CelestialType) (interface {}, error)
	GetEspionageReport(int64) (ogame.EspionageReport, error)
	GetEspionageReportFor(ogame.Coordinate) (ogame.EspionageReport, error)
	GetEspionageReportMessages(int64) ([]ogame.EspionageReportSummary, error)
	GetExpeditionMessageAt(time.Time) (ogame.ExpeditionMessage, error)
	GetExpeditionMessages(int64) ([]ogame.ExpeditionMessage, error)
	GetFacilities(ogame.CelestialID, ...Option) (ogame.Facilities, error)
	GetFleetDispatch(ogame.CelestialID, ...Option) (ogame.FleetDispatchInfos, error)
	GetFleets(...Option) ([]ogame.Fleet, ogame.Slots)
	GetFleetsFromEventList() []ogame.Fleet
	GetItems(ogame.CelestialID) ([]ogame.Item, error)
	GetLfBonuses() (ogame.LfBonuses, error)
	GetLfBuildings(ogame.CelestialID, ...Option) (ogame.LfBuildings, error)
	GetLfResearch(ogame.CelestialID, ...Option) (ogame.LfResearches, error)
	GetLfResearchDetails(ogame.CelestialID, ...Option) (ogame.LfResearchDetails, error)
	GetMoon(IntoMoon) (Moon, error)
	GetMoons() ([]Moon, error)
	GetPageContent(url.Values) ([]uint8, error)
	GetPlanet(IntoPlanet) (Planet, error)
	GetPlanets() ([]Planet, error)
	GetPositionsAvailableForDiscoveryFleet(int64, int64, ...Option) ([]ogame.Coordinate, error)
	GetProduction(ogame.CelestialID) ([]ogame.Quantifiable, int64, error)
	GetResearch() (ogame.Researches, error)
	GetResourceSettings(ogame.PlanetID, ...Option) (ogame.ResourceSettings, error)
	GetResources(ogame.CelestialID) (ogame.Resources, error)
	GetResourcesBuildings(ogame.CelestialID, ...Option) (ogame.ResourcesBuildings, error)
	GetResourcesDetails(ogame.CelestialID) (ogame.ResourcesDetails, error)
	GetResourcesProductions(ogame.PlanetID) (ogame.Resources, error)
	GetResourcesProductionsLight(ogame.ResourcesBuildings, ogame.Researches, ogame.ResourceSettings, ogame.Temperature) ogame.Resources
	GetShips(ogame.CelestialID, ...Option) (ogame.ShipsInfos, error)
	GetSlots() (ogame.Slots, error)
	GetTechs(ogame.CelestialID) (ogame.ResourcesBuildings, ogame.Facilities, ogame.ShipsInfos, ogame.DefensesInfos, ogame.Researches, ogame.LfBuildings, ogame.LfResearches, error)
	GetUserInfos() (ogame.UserInfos, error)
	HeadersForPage(string) (http.Header, error)
	Highscore(int64, int64, int64) (ogame.Highscore, error)
	IsUnderAttack(...Option) (bool, error)
	JumpGate(ogame.MoonID, ogame.MoonID, ogame.ShipsInfos) (bool, int64, error)
	JumpGateDestinations(ogame.MoonID) ([]ogame.MoonID, int64, error)
	Login() error
	LoginWithBearerToken(string) (bool, error)
	LoginWithExistingCookies() (bool, error)
	Logout()
	OfferBuyMarketplace(interface {}, int64, int64, int64, int64, ogame.CelestialID) error
	OfferSellMarketplace(interface {}, int64, int64, int64, int64, ogame.CelestialID) error
	Phalanx(ogame.MoonID, ogame.Coordinate) ([]ogame.PhalanxFleet, error)
	PostPageContent(url.Values, url.Values) ([]uint8, error)
	RecruitOfficer(int64, int64) error
	SelectLfResearchArtifacts(ogame.PlanetID, int64, ogame.ID) error
	SelectLfResearchRandom(ogame.PlanetID, int64) error
	SelectLfResearchSelect(ogame.PlanetID, int64) error
	SendDiscoveryFleet(ogame.CelestialID, ogame.Coordinate, ...Option) error
	SendDiscoveryFleet2(ogame.CelestialID, ogame.Coordinate, ...Option) (ogame.Fleet, error)
	SendFleet(ogame.CelestialID, ogame.ShipsInfos, ogame.Speed, ogame.Coordinate, ogame.MissionID, ogame.Resources, int64, int64) (ogame.Fleet, error)
	SendIPM(ogame.PlanetID, ogame.Coordinate, int64, ogame.ID) (int64, error)
	SendMessage(int64, string) error
	SendMessageAlliance(int64, string) error
	ServerTime() (time.Time, error)
	SetInitiator(string) Prioritizable
	SetPreferences(ogame.Preferences) error
	SetPreferencesLang(string) error
	SetResourceSettings(ogame.PlanetID, ogame.ResourceSettings) error
	SetVacationMode() error
	TearDown(ogame.CelestialID, ogame.ID) error
	TechnologyDetails(ogame.CelestialID, ogame.ID) (ogame.TechnologyDetails, error)
	Tx(func(Prioritizable) error) error
	TxNamed(string, func(Prioritizable) error) error
	UnsafePhalanx(ogame.MoonID, ogame.Coordinate) ([]ogame.PhalanxFleet, error)
	UseDM(ogame.DMType, ogame.CelestialID) error
}
