package nja

import "time"
import "context"
import "nja/pkg/gameforge"
import "nja/pkg/ogame"
import "nja/pkg/wrapper"
import "nja/pkg/njaBot"
import "nja/pkg/tgbotapi"
import "nja/pkg/database"
import "nja/pkg/cron"
import "nja/pkg/limiter"

var ACADEMYOFSCIENCES ogame.ID
var ACOUSTICSCANNING ogame.ID
var ADVANCEDRECYCLINGPLANT ogame.ID
var ALLIANCEDEPOT ogame.ID
var ANTIBALLISTICMISSILES ogame.ID
var ANTIMATTERCONDENSER ogame.ID
var ANTIMATTERCONVECTOR ogame.ID
var ARMOURTECHNOLOGY ogame.ID
var ARTIFICIALSWARMINTELLIGENCE ogame.ID
var ASSEMBLYLINE ogame.ID
var ASTROPHYSICS ogame.ID
var ATTACK int
var AUTOMATEDTRANSPORTLINES ogame.ID
var AUTOMATISEDASSEMBLYCENTRE ogame.ID
func Abandon(wrapper.IntoPlanet) error { panic("not implemented") }
func ActivateItem(string, ogame.CelestialID) error { panic("not implemented") }
func AddItemToQueue(ogame.CelestialID, ogame.ID, int64) error { panic("not implemented") }
func AllBots() []njaBot.IVMBot { panic("not implemented") }
func AnswerCallbackQuery(string, string) error { panic("not implemented") }
func Atoi(string) int { panic("not implemented") }
var BATTLECRUISER ogame.ID
var BATTLECRUISERMKII ogame.ID
var BATTLESHIP ogame.ID
var BIOMODIFIER ogame.ID
var BIOSPHEREFARM ogame.ID
var BIOTECHLAB ogame.ID
var BOMBER ogame.ID
var BOMBERMKII ogame.ID
func Base64(string) string { panic("not implemented") }
func Base64Decode(string) (string, error) { panic("not implemented") }
func Begin() wrapper.Prioritizable { panic("not implemented") }
var BotID database.BotID
func BotsCount() int64 { panic("not implemented") }
func Build(wrapper.IntoCelestial, ogame.ID, int64) error { panic("not implemented") }
func BuildBuilding(wrapper.IntoCelestial, ogame.ID) error { panic("not implemented") }
func BuildCancelable(wrapper.IntoCelestial, ogame.ID) error { panic("not implemented") }
func BuildDefense(wrapper.IntoCelestial, ogame.ID, int64) error { panic("not implemented") }
func BuildProduction(wrapper.IntoCelestial, ogame.ID, int64) error { panic("not implemented") }
func BuildShips(wrapper.IntoCelestial, ogame.ID, int64) error { panic("not implemented") }
func BuildTechnology(wrapper.IntoCelestial, ogame.ID) error { panic("not implemented") }
var BuildingsArr []ogame.ID
func BuyOfferOfTheDay() error { panic("not implemented") }
func BuyResetTree(wrapper.IntoPlanet, int64) error { panic("not implemented") }
func Bytes2Str([]uint8) string { panic("not implemented") }
var CARGOHOLDEXPANSIONCIVILIANSHIPS ogame.ID
var CATALYSERTECHNOLOGY ogame.ID
var CHIPMASSPRODUCTION ogame.ID
var CHRYSALISACCELERATOR ogame.ID
var CLONINGLABORATORY ogame.ID
var COLLECTOR ogame.CharacterClass
var COLONIZE int
var COLONYSHIP ogame.ID
var COMBUSTIONDRIVE ogame.ID
var COMPUTERTECHNOLOGY ogame.ID
var CRAWLER ogame.ID
var CRUISER ogame.ID
var CRUISERMKII ogame.ID
var CRYSTALFARM ogame.ID
var CRYSTALMINE ogame.ID
var CRYSTALREFINERY ogame.ID
var CRYSTALSTORAGE ogame.ID
func CalcCargo(int64) (int64, int64) { panic("not implemented") }
func CalcFastCargo(int64, int64, int64) (int64, int64, int64) { panic("not implemented") }
func CalcFastCargoPF(int64, int64, int64, int64) (int64, int64, int64, int64) { panic("not implemented") }
func CalcPreferredCargo(ogame.ID, int64, int64, int64, int64, int64, bool) (int64, int64, int64, int64, int64) { panic("not implemented") }
func CancelBuilding(wrapper.IntoCelestial) error { panic("not implemented") }
func CancelFleet(ogame.FleetID) error { panic("not implemented") }
func CancelResearch(wrapper.IntoCelestial) error { panic("not implemented") }
func Cargo(ogame.ShipsInfos) int64 { panic("not implemented") }
func ClearAllConstructionQueues() error { panic("not implemented") }
func ClearOut() { panic("not implemented") }
func Clock() (int, int, int) { panic("not implemented") }
func ConstructionTime(ogame.ID, int64, ogame.Facilities) time.Duration { panic("not implemented") }
func ConstructionsBeingBuilt(wrapper.IntoCelestial) (ogame.ID, int64, ogame.ID, int64) { panic("not implemented") }
func ConstructionsBeingBuiltLf(wrapper.IntoCelestial) (ogame.ID, int64, ogame.ID, int64, ogame.ID, int64, ogame.ID, int64) { panic("not implemented") }
func ConvertIntoCoordinate(interface {}) (ogame.Coordinate, error) { panic("not implemented") }
func CoordinatesAvailableForDiscoveryFleet(wrapper.IntoCelestial, int64, int64) ([]ogame.Coordinate, error) { panic("not implemented") }
func CreateFleet(ogame.FleetID, string) { panic("not implemented") }
func CreateUnion(ogame.Fleet, []string) (int64, error) { panic("not implemented") }
func CronExec(string, func()) (cron.EntryID, error) { panic("not implemented") }
func DBGetDefenses(wrapper.IntoCelestial) (ogame.DefensesInfos, error) { panic("not implemented") }
func DBGetFacilities(wrapper.IntoCelestial) (ogame.Facilities, error) { panic("not implemented") }
func DBGetResearches() ogame.Researches { panic("not implemented") }
func DBGetResourceBuildings(wrapper.IntoCelestial) (ogame.ResourcesBuildings, error) { panic("not implemented") }
func DBGetResourceSettings(wrapper.IntoCelestial) (ogame.ResourceSettings, error) { panic("not implemented") }
func DBGetShips(wrapper.IntoCelestial) (ogame.ShipsInfos, error) { panic("not implemented") }
var DEATHSTAR ogame.ID
var DEBRIS_TYPE ogame.CelestialType
var DEPOTAI ogame.ID
var DEPTHSOUNDING ogame.ID
var DESTROY int
var DESTROYER ogame.ID
var DESTROYERMKII ogame.ID
var DEUTERIUMSYNTHESISER ogame.ID
var DEUTERIUMSYNTHESIZER ogame.ID
var DEUTERIUMTANK ogame.ID
var DIAMONDENERGYTRANSMITTER ogame.ID
func DISCORD_WEBHOOK() string { panic("not implemented") }
var DISCOVERER ogame.CharacterClass
var DISRUPTIONCHAMBER ogame.ID
func Date() (int, time.Month, int) { panic("not implemented") }
var DefencesArr []ogame.ID
func DeleteAllMessagesFromTab(int64) error { panic("not implemented") }
func DeleteMessage(int64) error { panic("not implemented") }
func DestroyRockets(ogame.PlanetID, int64, int64) error { panic("not implemented") }
func DisableNJA() { panic("not implemented") }
func Distance(ogame.Coordinate, ogame.Coordinate) int64 { panic("not implemented") }
func DoAuction(map[ogame.CelestialID]ogame.Resources) error { panic("not implemented") }
func Dotify(int64) string { panic("not implemented") }
func DurationBetweenTimeStrings(string, string) time.Duration { panic("not implemented") }
var EFFICIENCYMODULE ogame.ID
var EFFICIENTSWARMINTELLIGENCE ogame.ID
var EIGHTY_FIVE_PERCENT float64
var EIGHTY_PERCENT float64
var ENERGYTECHNOLOGY ogame.ID
var ENHANCEDPRODUCTIONTECHNOLOGIES ogame.ID
var ENHANCEDSENSORTECHNOLOGY ogame.ID
var ESPIONAGEPROBE ogame.ID
var ESPIONAGETECHNOLOGY ogame.ID
var EXPEDITION int
var EXPERIMENTALRECYCLINGTECHNOLOGY ogame.ID
var EXPERIMENTALWEAPONSTECHNOLOGY ogame.ID
func EnableNJA() { panic("not implemented") }
func ExecAt(interface {}, func()) context.CancelFunc { panic("not implemented") }
func ExecAtCh(interface {}, func()) <-chan struct {} { panic("not implemented") }
func ExecAtSync(interface {}, func()) { panic("not implemented") }
func ExecIn(int64, func()) context.CancelFunc { panic("not implemented") }
func ExecInCh(int64, func()) <-chan struct {} { panic("not implemented") }
func ExecInSync(int64, func()) { panic("not implemented") }
var FIFTEEN_PERCENT float64
var FIFTY_FIVE_PERCENT float64
var FIFTY_PERCENT float64
var FILE string
var FIVE_PERCENT float64
var FOODSILO ogame.ID
var FORTY_FIVE_PERCENT float64
var FORTY_PERCENT float64
var FORUMOFTRANSCENDENCE ogame.ID
var FOURTY_PERCENT float64
var FUSIONCELLFACTORY ogame.ID
var FUSIONDRIVES ogame.ID
var FUSIONPOWEREDPRODUCTION ogame.ID
var FUSIONREACTOR ogame.ID
func FindAbandonedPlanetWithMinimumTravelTime(ogame.Coordinate, ogame.ShipsInfos, time.Duration, bool) (ogame.Coordinate, time.Duration, int64, error) { panic("not implemented") }
func FindDebrisFieldWithMinimumTravelTime(ogame.Coordinate, ogame.ShipsInfos, time.Duration) (ogame.Coordinate, time.Duration, int64, error) { panic("not implemented") }
func FindEmptyPlanetWithMinimumTravelTime(ogame.Coordinate, ogame.ShipsInfos, time.Duration) (ogame.Coordinate, time.Duration, int64, error) { panic("not implemented") }
func FindInactivePlanetWithMinimumTravelTime(ogame.Coordinate, ogame.ShipsInfos, time.Duration, bool) (ogame.Coordinate, time.Duration, int64, error) { panic("not implemented") }
func FlightTime(wrapper.IntoCoordinate, wrapper.IntoCoordinate, ogame.Speed, ogame.ShipsInfos, ogame.MissionID) (int64, int64) { panic("not implemented") }
func FreeResetTree(wrapper.IntoPlanet, int64) error { panic("not implemented") }
var GALAXIES int64
var GAUSSCANNON ogame.ID
var GENERAL ogame.CharacterClass
var GENERALOVERHAULBATTLECRUISER ogame.ID
var GENERALOVERHAULBATTLESHIP ogame.ID
var GENERALOVERHAULBOMBER ogame.ID
var GENERALOVERHAULCRUISER ogame.ID
var GENERALOVERHAULDESTROYER ogame.ID
var GENERALOVERHAULLIGHTFIGHTER ogame.ID
var GEOTHERMALPOWERPLANTS ogame.ID
var GRAVITATIONSENSORS ogame.ID
var GRAVITONTECHNOLOGY ogame.ID
var GROUPEDATTACK int
func GalaxyInfos(int64, int64, ...wrapper.Option) (ogame.SystemInfos, error) { panic("not implemented") }
func GalaxyInfosUsing(int64, int64, ogame.CelestialID) (ogame.SystemInfos, error) { panic("not implemented") }
func GetActiveItems(wrapper.IntoCelestial) ([]ogame.ActiveItem, error) { panic("not implemented") }
func GetAllResources() (map[ogame.CelestialID]ogame.Resources, error) { panic("not implemented") }
func GetAttacks() []ogame.AttackEvent { panic("not implemented") }
func GetAttacksUsing(ogame.CelestialID) []ogame.AttackEvent { panic("not implemented") }
func GetAuction() (ogame.Auction, error) { panic("not implemented") }
func GetBotByID(database.BotID) njaBot.IVMBot { panic("not implemented") }
func GetCachedCelestial(wrapper.IntoCelestial) wrapper.Celestial { panic("not implemented") }
func GetCachedCelestials() []wrapper.Celestial { panic("not implemented") }
func GetCachedLfBonuses() (ogame.LfBonuses, error) { panic("not implemented") }
func GetCachedMoons() []wrapper.Moon { panic("not implemented") }
func GetCachedPlanets() []wrapper.Planet { panic("not implemented") }
func GetCachedPlayer() ogame.UserInfos { panic("not implemented") }
func GetCachedResearch() ogame.Researches { panic("not implemented") }
func GetCelestial(wrapper.IntoCelestial) (wrapper.Celestial, error) { panic("not implemented") }
func GetCelestials() ([]wrapper.Celestial, error) { panic("not implemented") }
func GetCombatReportSummaryFor(ogame.Coordinate) (ogame.CombatReportSummary, error) { panic("not implemented") }
func GetDMCosts(ogame.CelestialID) (ogame.DMCosts, error) { panic("not implemented") }
func GetDefense(wrapper.IntoCelestial) (ogame.DefensesInfos, error) { panic("not implemented") }
func GetDepartureTime(interface {}, wrapper.IntoCoordinate, wrapper.IntoCoordinate, ogame.ShipsInfos, ogame.Speed, ogame.MissionID) time.Time { panic("not implemented") }
func GetEspionageReport(int64) (ogame.EspionageReport, error) { panic("not implemented") }
func GetEspionageReportFor(ogame.Coordinate) (ogame.EspionageReport, error) { panic("not implemented") }
func GetEspionageReportMessages() ([]ogame.EspionageReportSummary, error) { panic("not implemented") }
func GetExpeditionsConfigs() *njaBot.ExpeditionsConfigs { panic("not implemented") }
func GetFacilities(wrapper.IntoCelestial) (ogame.Facilities, error) { panic("not implemented") }
func GetFleetSlotsReserved() int64 { panic("not implemented") }
func GetFleets(...wrapper.Option) ([]ogame.Fleet, ogame.Slots) { panic("not implemented") }
func GetHighscore(int64, int64, int64) (ogame.Highscore, error) { panic("not implemented") }
func GetHomeWorld() wrapper.Celestial { panic("not implemented") }
func GetItems(wrapper.IntoCelestial) ([]ogame.Item, error) { panic("not implemented") }
func GetLanguage() string { panic("not implemented") }
func GetLfBonuses() (ogame.LfBonuses, error) { panic("not implemented") }
func GetLfBuildings(wrapper.IntoCelestial) (ogame.LfBuildings, error) { panic("not implemented") }
func GetLfResearch(wrapper.IntoCelestial) (ogame.LfResearches, error) { panic("not implemented") }
func GetLfResearchDetails(wrapper.IntoCelestial) (ogame.LfResearchDetails, error) { panic("not implemented") }
func GetMoon(wrapper.IntoMoon) (wrapper.Moon, error) { panic("not implemented") }
func GetMoons() []wrapper.Moon { panic("not implemented") }
func GetNextDatetimeAt(int64, int64, int64) (time.Time, error) { panic("not implemented") }
func GetPlanet(wrapper.IntoPlanet) (wrapper.Planet, error) { panic("not implemented") }
func GetPlanetInfo(wrapper.IntoCoordinate, ...wrapper.Option) (ogame.PlanetInfos, error) { panic("not implemented") }
func GetPlanetInfoUsing(interface {}, ogame.CelestialID) (ogame.PlanetInfos, error) { panic("not implemented") }
func GetPlanetLifeformType(wrapper.IntoPlanet) ogame.LifeformType { panic("not implemented") }
func GetPlanets() []wrapper.Planet { panic("not implemented") }
func GetPlayerCoordinates(int64) []ogame.Coordinate { panic("not implemented") }
func GetPreferences() *ogame.Preferences { panic("not implemented") }
func GetPrice(ogame.ID, int64) ogame.Resources { panic("not implemented") }
func GetPriceLfResearch(wrapper.IntoCelestial, ogame.ID, int64) ogame.Resources { panic("not implemented") }
func GetProduction(wrapper.IntoCelestial) ([]ogame.Quantifiable, int64, error) { panic("not implemented") }
func GetPublicIP() (string, error) { panic("not implemented") }
func GetRequirements(ogame.ID) map[ogame.ID]int64 { panic("not implemented") }
func GetResearch() ogame.Researches { panic("not implemented") }
func GetResourceSettings(wrapper.IntoPlanet) (ogame.ResourceSettings, error) { panic("not implemented") }
func GetResources(wrapper.IntoCelestial) (ogame.Resources, error) { panic("not implemented") }
func GetResourcesBuildings(wrapper.IntoCelestial) (ogame.ResourcesBuildings, error) { panic("not implemented") }
func GetResourcesDetails(wrapper.IntoCelestial) (ogame.ResourcesDetails, error) { panic("not implemented") }
func GetResourcesProductions(wrapper.IntoPlanet) (ogame.Resources, error) { panic("not implemented") }
func GetResourcesProductionsLight(ogame.ResourcesBuildings, ogame.Researches, ogame.ResourceSettings, ogame.Temperature) ogame.Resources { panic("not implemented") }
func GetRunningScripts() []string { panic("not implemented") }
func GetScripts() []string { panic("not implemented") }
func GetServer() gameforge.Server { panic("not implemented") }
func GetSession() string { panic("not implemented") }
func GetShips(wrapper.IntoCelestial) (ogame.ShipsInfos, error) { panic("not implemented") }
func GetSlot(string) <-chan struct {} { panic("not implemented") }
func GetSlots() ogame.Slots { panic("not implemented") }
func GetSortedCelestials(ogame.Coordinate) []wrapper.Celestial { panic("not implemented") }
func GetSortedMoons(ogame.Coordinate) []wrapper.Moon { panic("not implemented") }
func GetSortedPlanets(ogame.Coordinate) []wrapper.Planet { panic("not implemented") }
func GetSupplies(wrapper.IntoCelestial) (ogame.ResourcesBuildings, error) { panic("not implemented") }
func GetSystemsInRange(int64, int64) []int64 { panic("not implemented") }
func GetSystemsInRangeAsc(int64, int64) []int64 { panic("not implemented") }
func GetSystemsInRangeDesc(int64, int64) []int64 { panic("not implemented") }
func GetTechs(wrapper.IntoCelestial) (ogame.ResourcesBuildings, ogame.Facilities, ogame.ShipsInfos, ogame.DefensesInfos, ogame.Researches, error) { panic("not implemented") }
func GetTechs2(wrapper.IntoCelestial) (ogame.ResourcesBuildings, ogame.Facilities, ogame.ShipsInfos, ogame.DefensesInfos, ogame.Researches, ogame.LfBuildings, ogame.LfResearches, error) { panic("not implemented") }
func GetTimestamp() int64 { panic("not implemented") }
func GetUniverseName() string { panic("not implemented") }
func GetUserInfos() ogame.UserInfos { panic("not implemented") }
func GetUsername() string { panic("not implemented") }
func Go(func()) { panic("not implemented") }
var HALLSOFREALISATION ogame.ID
var HARDENEDDIAMONDDRILLHEADS ogame.ID
var HEATRECOVERY ogame.ID
var HEAVYFIGHTER ogame.ID
var HEAVYLASER ogame.ID
var HIGHENERGYPUMPSYSTEMS ogame.ID
var HIGHENERGYSMELTING ogame.ID
var HIGHPERFORMANCEEXTRACTORS ogame.ID
var HIGHPERFORMANCESYNTHESISER ogame.ID
var HIGHPERFORMANCETERRAFORMER ogame.ID
var HIGHPERFORMANCETRANSFORMER ogame.ID
var HIGHTEMPERATURESUPERCONDUCTORS ogame.ID
var HUMANS ogame.LifeformType
var HUNDRED_PERCENT float64
var HYPERSPACEDRIVE ogame.ID
var HYPERSPACETECHNOLOGY ogame.ID
func ID2Str(ogame.ID) string { panic("not implemented") }
var IMPROVEDDRONEAI ogame.ID
var IMPROVEDLABTECHNOLOGY ogame.ID
var IMPROVEDSTELLARATOR ogame.ID
var IMPULSEDRIVE ogame.ID
var INTERGALACTICENVOYS ogame.ID
var INTERGALACTICRESEARCHNETWORK ogame.ID
var INTERPLANETARYANALYSISNETWORK ogame.ID
var INTERPLANETARYMISSILES ogame.ID
var IONCANNON ogame.ID
var IONCRYSTALENHANCEMENTHEAVYFIGHTER ogame.ID
var IONCRYSTALMODULES ogame.ID
var IONTECHNOLOGY ogame.ID
var IS_CLOUD bool
var IS_SELF_HOST bool
func IntervalExec(int64, func()) context.CancelFunc { panic("not implemented") }
func IsActive() bool { panic("not implemented") }
func IsAvailable(ogame.ID, ogame.CelestialID, ogame.ResourcesBuildings, ogame.LfBuildings, ogame.LfResearches, ogame.Facilities, ogame.Researches, int64) bool { panic("not implemented") }
func IsNJAEnabled() bool { panic("not implemented") }
func IsPausedScript(string) bool { panic("not implemented") }
func IsPioneers() bool { panic("not implemented") }
func IsScriptRunning(string) bool { panic("not implemented") }
func IsStarted() bool { panic("not implemented") }
func IsUnderAttack() bool { panic("not implemented") }
var JUMPGATE ogame.ID
func JumpGate(ogame.MoonID, ogame.MoonID, ogame.ShipsInfos) error { panic("not implemented") }
func JumpGate2(ogame.MoonID, ogame.MoonID, ogame.ShipsInfos) (bool, int64, error) { panic("not implemented") }
func JumpGateDestinations(ogame.MoonID) ([]ogame.MoonID, int64, error) { panic("not implemented") }
var KAELESH ogame.LifeformType
var KAELESHDISCOVERERENHANCEMENT ogame.ID
var LARGECARGO ogame.ID
var LARGESHIELDDOME ogame.ID
var LASERTECHNOLOGY ogame.ID
var LIGHTFIGHTER ogame.ID
var LIGHTFIGHTERMKII ogame.ID
var LIGHTLASER ogame.ID
var LOWTEMPERATUREDRIVES ogame.ID
var LUNARBASE ogame.ID
var LfBuildingsHumansArr []ogame.ID
var LfBuildingsKaeleshArr []ogame.ID
var LfBuildingsMechasArr []ogame.ID
var LfBuildingsRocktalArr []ogame.ID
func LogDebug(...interface {}) { panic("not implemented") }
func LogDebugf(string, ...interface {}) { panic("not implemented") }
func LogError(...interface {}) { panic("not implemented") }
func LogErrorf(string, ...interface {}) { panic("not implemented") }
func LogInfo(...interface {}) { panic("not implemented") }
func LogInfof(string, ...interface {}) { panic("not implemented") }
func LogWarn(...interface {}) { panic("not implemented") }
func LogWarnf(string, ...interface {}) { panic("not implemented") }
func Login() error { panic("not implemented") }
func Logout() { panic("not implemented") }
func LowestSpeed(interface {}, wrapper.IntoCoordinate, wrapper.IntoCoordinate, ogame.ShipsInfos) (ogame.Speed, int64, time.Time, error) { panic("not implemented") }
var MAGMAFORGE ogame.ID
var MAGMAPOWEREDPRODUCTION ogame.ID
var MAGMAPOWEREDPUMPSYSTEMS ogame.ID
var MECHANGENERALENHANCEMENT ogame.ID
var MECHAS ogame.LifeformType
var MEDITATIONENCLAVE ogame.ID
var MEGALITH ogame.ID
var METALMINE ogame.ID
var METALSTORAGE ogame.ID
var METROPOLIS ogame.ID
var MICROCHIPASSEMBLYLINE ogame.ID
var MINERALRESEARCHCENTRE ogame.ID
var MISSILEATTACK int
var MISSILESILO ogame.ID
var MOON_TYPE ogame.CelestialType
func MillisecondsBetweenTimeStrings(string, string) int64 { panic("not implemented") }
var MoonBuildingsArr []ogame.ID
var NANITEFACTORY ogame.ID
var NANOREPAIRBOTS ogame.ID
var NEUROCALIBRATIONCENTRE ogame.ID
var NEUROINTERFACE ogame.ID
var NEUROMODALCOMPRESSOR ogame.ID
var NINETY_FIVE_PERCENT float64
var NINETY_PERCENT float64
var NONE_LF_TYPE ogame.LifeformType
var NO_ALLIANCE_CLASS ogame.AllianceClass
var NO_CLASS ogame.CharacterClass
func NewCoordinate(int64, int64, int64, ogame.CelestialType) ogame.Coordinate { panic("not implemented") }
func NewFleet() *njaBot.NJAFleetBuilder { panic("not implemented") }
func NewInlineKeyboardButtonData(string, string) tgbotapi.InlineKeyboardButton { panic("not implemented") }
func NewInlineKeyboardMarkup(...[]tgbotapi.InlineKeyboardButton) *njaBot.InlineKeyboardMarkup { panic("not implemented") }
func NewInlineKeyboardRow(...tgbotapi.InlineKeyboardButton) []tgbotapi.InlineKeyboardButton { panic("not implemented") }
func NewKeyboardButton(string) tgbotapi.KeyboardButton { panic("not implemented") }
func NewKeyboardButtonRow(...tgbotapi.KeyboardButton) []tgbotapi.KeyboardButton { panic("not implemented") }
func NewQuantifiable(ogame.ID, int64) ogame.Quantifiable { panic("not implemented") }
func NewReplyKeyboard(...[]tgbotapi.KeyboardButton) tgbotapi.ReplyKeyboardMarkup { panic("not implemented") }
func NewResourceSettings(int64, int64, int64, int64, int64, int64, int64) ogame.ResourceSettings { panic("not implemented") }
func NewResources(int64, int64, int64) ogame.Resources { panic("not implemented") }
func NewShipsInfos() *ogame.ShipsInfos { panic("not implemented") }
func NewSimpleLimiter(time.Duration) *limiter.Limiter { panic("not implemented") }
func NewTemperature(int64, int64) ogame.Temperature { panic("not implemented") }
func Notify(string, string) { panic("not implemented") }
func NowInTimeRange(string, string) bool { panic("not implemented") }
func NowTimeString() string { panic("not implemented") }
var OBSIDIANSHIELDREINFORCEMENT ogame.ID
var OGAME_SERVER gameforge.Server
var OPTIMISEDSILOCONSTRUCTIONMETHOD ogame.ID
var ORBITALDEN ogame.ID
var ORIKTORIUM ogame.ID
var OVERCLOCKINGBATTLESHIP ogame.ID
var OVERCLOCKINGHEAVYFIGHTER ogame.ID
var OVERCLOCKINGLARGECARGO ogame.ID
func OnFleetDispatch() chan ogame.Fleet { panic("not implemented") }
func OwnBots() []njaBot.IVMBot { panic("not implemented") }
var PARK int
var PARKINTHATALLY int
var PATHFINDER ogame.ID
var PLANETARYSHIELD ogame.ID
var PLANET_TYPE ogame.CelestialType
var PLASMADRIVE ogame.ID
var PLASMATECHNOLOGY ogame.ID
var PLASMATERRAFORMER ogame.ID
var PLASMATURRET ogame.ID
var PRODUCTIONASSEMBLYHALL ogame.ID
var PSIONICMODULATOR ogame.ID
var PSIONICNETWORK ogame.ID
var PSIONICSHIELDMATRIX ogame.ID
var PSYCHOHARMONISER ogame.ID
func ParseCoord(string) (ogame.Coordinate, error) { panic("not implemented") }
func ParseNextDatetimeAt(string) (time.Time, error) { panic("not implemented") }
func PauseScript(string) error { panic("not implemented") }
func Phalanx(wrapper.IntoMoon, wrapper.IntoCoordinate) ([]ogame.Fleet, error) { panic("not implemented") }
var PlanetBuildingsArr []ogame.ID
func PlaySound(int64, float64) error { panic("not implemented") }
func PlayerDataByID(int64) (njaBot.PlayerData, error) { panic("not implemented") }
func PlayerDataByName(string) (njaBot.PlayerData, error) { panic("not implemented") }
func PlayersData() ([]njaBot.PlayerDataShort, error) { panic("not implemented") }
func Print(...interface {}) { panic("not implemented") }
func Printf(string, ...interface {}) { panic("not implemented") }
var QUANTUMCOMPUTERCENTRE ogame.ID
var REAPER ogame.ID
var RECYCLEDEBRISFIELD int
var RECYCLER ogame.ID
var RESEARCHAI ogame.ID
var RESEARCHCENTRE ogame.ID
var RESEARCHER ogame.AllianceClass
var RESEARCHLAB ogame.ID
var RESIDENTIALSECTOR ogame.ID
var ROBOTASSISTANTS ogame.ID
var ROBOTICSFACTORY ogame.ID
var ROBOTICSRESEARCHCENTRE ogame.ID
var ROCKETLAUNCHER ogame.ID
var ROCKTAL ogame.LifeformType
var ROCKTALCOLLECTORENHANCEMENT ogame.ID
var RUNEFORGE ogame.ID
var RUNESHIELDS ogame.ID
var RUNETECHNOLOGIUM ogame.ID
func Random(int64, int64) int64 { panic("not implemented") }
func RangeCronExec(string, string, string, func()) (cron.EntryID, error) { panic("not implemented") }
func RepatriateCelestial(wrapper.IntoCelestial) (ogame.Resources, error) { panic("not implemented") }
func RepatriateNow() { panic("not implemented") }
func RepatriateSetAllDestinations(wrapper.IntoCelestial) error { panic("not implemented") }
func ResumeScript(string) error { panic("not implemented") }
var SANCTUARY ogame.ID
var SEABEDDEUTERIUMDEN ogame.ID
var SEARCHFORLIFEFORMS int
var SEISMICMININGTECHNOLOGY ogame.ID
var SENSORPHALANX ogame.ID
var SEVENTY_FIVE_PERCENT float64
var SEVENTY_PERCENT float64
var SHIELDEDMETALDEN ogame.ID
var SHIELDINGTECHNOLOGY ogame.ID
var SHIPMANUFACTURINGHALL ogame.ID
var SHIPYARD ogame.ID
var SIXTHSENSE ogame.ID
var SIXTY_FIVE_PERCENT float64
var SIXTY_PERCENT float64
var SKYSCRAPER ogame.ID
var SLINGSHOTAUTOPILOT ogame.ID
var SMALLCARGO ogame.ID
var SMALLSHIELDDOME ogame.ID
var SOLARPLANT ogame.ID
var SOLARSATELLITE ogame.ID
var SPACEDOCK ogame.ID
var SPY int
var STEALTHFIELDGENERATOR ogame.ID
var SULPHIDEPROCESS ogame.ID
var SUPERCOMPUTER ogame.ID
var SUPRAREFRACTOR ogame.ID
var SYSTEMS int64
func SelectLfResearchArtifacts(wrapper.IntoPlanet, int64, ogame.ID) error { panic("not implemented") }
func SelectLfResearchRandom(wrapper.IntoPlanet, int64) error { panic("not implemented") }
func SelectLfResearchSelect(wrapper.IntoPlanet, int64) error { panic("not implemented") }
func SendDiscord(string, string) error { panic("not implemented") }
func SendDiscoveryFleet(wrapper.IntoCelestial, wrapper.IntoCoordinate) error { panic("not implemented") }
func SendDiscoveryFleet2(wrapper.IntoCelestial, wrapper.IntoCoordinate) (ogame.Fleet, error) { panic("not implemented") }
func SendFleet(ogame.CelestialID, ogame.ShipsInfos, ogame.Speed, ogame.Coordinate, ogame.MissionID, ogame.Resources, int64, int64) (ogame.Fleet, error) { panic("not implemented") }
func SendIPM(ogame.PlanetID, ogame.Coordinate, int64, ogame.ID) (int64, error) { panic("not implemented") }
func SendMessage(int64, string) error { panic("not implemented") }
func SendMessageAlliance(int64, string) error { panic("not implemented") }
func SendSavedFleet(int64) (ogame.Fleet, error) { panic("not implemented") }
func SendSlack(string, string) error { panic("not implemented") }
func SendTelegram(int64, string) error { panic("not implemented") }
func ServerTime() time.Time { panic("not implemented") }
func ServerURL() string { panic("not implemented") }
func ServerVersion() string { panic("not implemented") }
func SetExpeditionsConfigs(*njaBot.ExpeditionsConfigs) { panic("not implemented") }
func SetFleetSlotsReserved(int64) { panic("not implemented") }
func SetHomeWorld(ogame.CelestialID) { panic("not implemented") }
func SetPreferences(*ogame.Preferences) error { panic("not implemented") }
func SetResourceSettings(wrapper.IntoPlanet, ogame.ResourceSettings) error { panic("not implemented") }
func ShipCargo(ogame.ID) int64 { panic("not implemented") }
func ShipFuel(ogame.ID) int64 { panic("not implemented") }
func ShipSpeed(ogame.ID) int64 { panic("not implemented") }
var ShipsArr []ogame.ID
func ShortDur(interface {}) string { panic("not implemented") }
func ShortShipsInfos(ogame.ShipsInfos) string { panic("not implemented") }
func Shuffle([]interface {}) { panic("not implemented") }
func Simulator(map[string]map[string]interface {}) (njaBot.SimulatorResult, error) { panic("not implemented") }
func Sleep(int64) { panic("not implemented") }
func SleepDur(time.Duration) { panic("not implemented") }
func SleepMin(int64) { panic("not implemented") }
func SleepMs(int64) { panic("not implemented") }
func SleepRandHour(int64, int64) { panic("not implemented") }
func SleepRandMin(int64, int64) { panic("not implemented") }
func SleepRandMs(int64, int64) { panic("not implemented") }
func SleepRandSec(int64, int64) { panic("not implemented") }
func SleepSec(int64) { panic("not implemented") }
func SleepUntil(string) { panic("not implemented") }
func SlotAvailable(string) bool { panic("not implemented") }
func SolarSatelliteProduction(ogame.Temperature, int64) int64 { panic("not implemented") }
func Speed(ogame.ShipsInfos) int64 { panic("not implemented") }
func StartRemoteScript(string, int64, string, string) error { panic("not implemented") }
func StartScript(string) error { panic("not implemented") }
func StopRemoteScript(string, int64, string, string) error { panic("not implemented") }
func StopScript(string) error { panic("not implemented") }
var TELEGRAM_CHAT_ID int64
var TELEKINETICDRIVE ogame.ID
var TELEKINETICTRACTORBEAM ogame.ID
var TEN_PERCENT float64
var TERRAFORMER ogame.ID
var THIRTY_FIVE_PERCENT float64
var THIRTY_PERCENT float64
var TRADER ogame.AllianceClass
var TRANSPORT int
var TWENTY_FIVE_PERCENT float64
var TWENTY_PERCENT float64
func TearDown(wrapper.IntoCelestial, ogame.ID) error { panic("not implemented") }
var TechnologiesArr []ogame.ID
func TechnologyDetails(wrapper.IntoCelestial, ogame.ID) (ogame.TechnologyDetails, error) { panic("not implemented") }
func Tx(func(wrapper.Prioritizable) error) error { panic("not implemented") }
var UNDERGROUNDCRYSTALDEN ogame.ID
var UPDATENETWORK ogame.ID
func UnsafePhalanx(wrapper.IntoMoon, wrapper.IntoCoordinate) ([]ogame.Fleet, error) { panic("not implemented") }
func UseDM(string, ogame.CelestialID) error { panic("not implemented") }
var VERSION string
var VOLCANICBATTERIES ogame.ID
var VORTEXCHAMBER ogame.ID
var WARRIOR ogame.AllianceClass
var WEAPONSTECHNOLOGY ogame.ID
func Weekday() int64 { panic("not implemented") }
var __FILE__ string
var __IS_CLOUD__ bool
var __IS_SELF_HOST__ bool
var __VERSION__ string
func print(...interface {}) { panic("not implemented") }
func println(...interface {}) { panic("not implemented") }
