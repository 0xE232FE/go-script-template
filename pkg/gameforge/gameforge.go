package gameforge

type Server struct {
	Language string
	Number int64
	AccountGroup string
	Name string
	PlayerCount int64
	PlayersOnline int64
	Opened string
	StartDate string
	EndDate *string
	ServerClosed int64
	Prefered int64
	SignupClosed int64
	MultiLanguage int64
	AvailableOn []string
	Settings struct { AKS int64; FleetSpeedWar int64; FleetSpeedHolding int64; FleetSpeedPeaceful int64; WreckField int64; ServerLabel string; EconomySpeed interface {}; PlanetFields int64; UniverseSize int64; ServerCategory string; EspionageProbeRaids int64; PremiumValidationGift int64; DebrisFieldFactorShips int64; ResearchDurationDivisor float64; DebrisFieldFactorDefence int64 }
}
func (Server) ProbeRaidsEnabled() bool { panic("not implemented") }
