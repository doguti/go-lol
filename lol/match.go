package lol

import (
	"context"
	"fmt"
)

// MatchService represents a service to interact with Match API.
type MatchService service

// Match represent a MatchDto.
type Match struct {
	SeasonID              *int                  `json:"seasonId,omitempty"`
	QueueID               *int                  `json:"queueId,omitempty"`
	GameID                *int                  `json:"gameId,omitempty"`
	ParticipantIdentities []ParticipantIdentity `json:"participantIdentities,omitempty"`
	GameVersion           *string               `json:"gameVersion,omitempty"`
	PlatformID            *string               `json:"platformId,omitempty"`
	GameMode              *string               `json:"gameMode,omitempty"`
	MapID                 *int                  `json:"mapId,omitempty"`
	GameType              *string               `json:"gameType,omitempty"`
	Teams                 []TeamStats           `json:"teams,omitempty"`
	Participants          []ParticipantDto         `json:"participants,omitempty"`
	GameDuration          *int                  `json:"gameDuration,omitempty"`
	GameCreation          *int                  `json:"gameCreation,omitempty"`
}

func (t Match) String() string {
	return Stringify(t)
}

// ParticipantIdentity represent a ParticipantIdentityDto.
type ParticipantIdentity struct {
	Player        *Player `json:"player,omitempty"`
	ParticipantID *int    `json:"participantId,omitempty"`
}

func (t ParticipantIdentity) String() string {
	return Stringify(t)
}

// Player represent a PlayerDto.
type Player struct {
	CurrentPlatformID *string `json:"currentPlatformId,omitempty"`
	SummonerName      *string `json:"summonerName,omitempty"`
	MatchHistoryURI   *string `json:"matchHistoryUri,omitempty"`
	PlatformID        *string `json:"platformId,omitempty"`
	CurrentAccountID  *int    `json:"currentAccountId,omitempty"`
	ProfileIcon       *int    `json:"profileIcon,omitempty"`
	SummonerID        *int    `json:"summonerId,omitempty"`
	AccountID         *int    `json:"accountId,omitempty"`
}

func (t Player) String() string {
	return Stringify(t)
}

// TeamStats represent a TeamStatsDto.
type TeamStats struct {
	FirstDragon          *bool      `json:"firstDragon,omitempty"`
	FirstInhibitor       *bool      `json:"firstInhibitor,omitempty"`
	Bans                 []TeamBans `json:"bans,omitempty"`
	BaronKills           *int       `json:"baronKills,omitempty"`
	FirstRiftHerald      *bool      `json:"firstRiftHerald,omitempty"`
	FirstBaron           *bool      `json:"firstBaron,omitempty"`
	RiftHeraldKills      *int       `json:"riftHeraldKills,omitempty"`
	FirstBlood           *bool      `json:"firstBlood,omitempty"`
	TeamID               *int       `json:"teamId,omitempty"`
	FirstTower           *bool      `json:"firstTower,omitempty"`
	VilemawKills         *int       `json:"vilemawKills,omitempty"`
	InhibitorKills       *int       `json:"inhibitorKills,omitempty"`
	TowerKills           *int       `json:"towerKills,omitempty"`
	DominionVictoryScore *int       `json:"dominionVictoryScore,omitempty"`
	Win                  *string    `json:"win,omitempty"`
	DragonKills          *int       `json:"dragonKills,omitempty"`
}

func (t TeamStats) String() string {
	return Stringify(t)
}

// TeamBans represent a TeamBansDto.
type TeamBans struct {
	PickTurn   *int `json:"pickTurn,omitempty"`
	ChampionID *int `json:"championId,omitempty"`
}

func (t TeamBans) String() string {
	return Stringify(t)
}

// ParticipantDto represent a ParticipantDto.
type ParticipantDto struct {
	Stats                     *ParticipantStats    `json:"stats,omitempty"`
	ParticipantID             *int                 `json:"participantId,omitempty"`
	Runes                     []Rune               `json:"runes,omitempty"`
	Timeline                  *ParticipantTimeline `json:"timeline,omitempty"`
	TeamID                    *int                 `json:"teamId,omitempty"`
	Spell2ID                  *int                 `json:"spell2Id,omitempty"`
	Masteries                 []Mastery            `json:"masteries,omitempty"`
	HighestAchievedSeasonTier *string              `json:"highestAchievedSeasonTier,omitempty"`
	Spell1ID                  *int                 `json:"spell1Id,omitempty"`
	ChampionID                *int                 `json:"championId,omitempty"`
}

func (t ParticipantDto) String() string {
	return Stringify(t)
}

// ParticipantStats represent a ParticipantStatsDto.
type ParticipantStats struct {
	PhysicalDamageDealt             *int  `json:"physicalDamageDealt,omitempty"`
	NeutralMinionsKilledTeamJungle  *int  `json:"neutralMinionsKilledTeamJungle,omitempty"`
	MagicDamageDealt                *int  `json:"magicDamageDealt,omitempty"`
	TotalPlayerScore                *int  `json:"totalPlayerScore,omitempty"`
	Deaths                          *int  `json:"deaths,omitempty"`
	Win                             *bool `json:"win,omitempty"`
	NeutralMinionsKilledEnemyJungle *int  `json:"neutralMinionsKilledEnemyJungle,omitempty"`
	AltarsCaptured                  *int  `json:"altarsCaptured,omitempty"`
	LargestCriticalStrike           *int  `json:"largestCriticalStrike,omitempty"`
	TotalDamageDealt                *int  `json:"totalDamageDealt,omitempty"`
	MagicDamageDealtToChampions     *int  `json:"magicDamageDealtToChampions,omitempty"`
	VisionWardsBoughtInGame         *int  `json:"visionWardsBoughtInGame,omitempty"`
	DamageDealtToObjectives         *int  `json:"damageDealtToObjectives,omitempty"`
	LargestKillingSpree             *int  `json:"largestKillingSpree,omitempty"`
	Item1                           *int  `json:"item1,omitempty"`
	QuadraKills                     *int  `json:"quadraKills,omitempty"`
	TeamObjective                   *int  `json:"teamObjective,omitempty"`
	TotalTimeCrowdControlDealt      *int  `json:"totalTimeCrowdControlDealt,omitempty"`
	LongestTimeSpentLiving          *int  `json:"longestTimeSpentLiving,omitempty"`
	WardsKilled                     *int  `json:"wardsKilled,omitempty"`
	FirstTowerAssist                *bool `json:"firstTowerAssist,omitempty"`
	FirstTowerKill                  *bool `json:"firstTowerKill,omitempty"`
	Item2                           *int  `json:"item2,omitempty"`
	Item3                           *int  `json:"item3,omitempty"`
	Item0                           *int  `json:"item0,omitempty"`
	FirstBloodAssist                *bool `json:"firstBloodAssist,omitempty"`
	VisionScore                     *int  `json:"visionScore,omitempty"`
	WardsPlaced                     *int  `json:"wardsPlaced,omitempty"`
	Item4                           *int  `json:"item4,omitempty"`
	Item5                           *int  `json:"item5,omitempty"`
	Item6                           *int  `json:"item6,omitempty"`
	TurretKills                     *int  `json:"turretKills,omitempty"`
	TripleKills                     *int  `json:"tripleKills,omitempty"`
	DamageSelfMitigated             *int  `json:"damageSelfMitigated,omitempty"`
	ChampLevel                      *int  `json:"champLevel,omitempty"`
	NodeNeutralizeAssist            *int  `json:"nodeNeutralizeAssist,omitempty"`
	FirstInhibitorKill              *bool `json:"firstInhibitorKill,omitempty"`
	GoldEarned                      *int  `json:"goldEarned,omitempty"`
	MagicalDamageTaken              *int  `json:"magicalDamageTaken,omitempty"`
	Kills                           *int  `json:"kills,omitempty"`
	DoubleKills                     *int  `json:"doubleKills,omitempty"`
	NodeCaptureAssist               *int  `json:"nodeCaptureAssist,omitempty"`
	TrueDamageTaken                 *int  `json:"trueDamageTaken,omitempty"`
	NodeNeutralize                  *int  `json:"nodeNeutralize,omitempty"`
	FirstInhibitorAssist            *bool `json:"firstInhibitorAssist,omitempty"`
	Assists                         *int  `json:"assists,omitempty"`
	UnrealKills                     *int  `json:"unrealKills,omitempty"`
	NeutralMinionsKilled            *int  `json:"neutralMinionsKilled,omitempty"`
	ObjectivePlayerScore            *int  `json:"objectivePlayerScore,omitempty"`
	CombatPlayerScore               *int  `json:"combatPlayerScore,omitempty"`
	DamageDealtToTurrets            *int  `json:"damageDealtToTurrets,omitempty"`
	AltarsNeutralized               *int  `json:"altarsNeutralized,omitempty"`
	PhysicalDamageDealtToChampions  *int  `json:"physicalDamageDealtToChampions,omitempty"`
	GoldSpent                       *int  `json:"goldSpent,omitempty"`
	TrueDamageDealt                 *int  `json:"trueDamageDealt,omitempty"`
	TrueDamageDealtToChampions      *int  `json:"trueDamageDealtToChampions,omitempty"`
	ParticipantID                   *int  `json:"participantId,omitempty"`
	PentaKills                      *int  `json:"pentaKills,omitempty"`
	TotalHeal                       *int  `json:"totalHeal,omitempty"`
	TotalMinionsKilled              *int  `json:"totalMinionsKilled,omitempty"`
	FirstBloodKill                  *bool `json:"firstBloodKill,omitempty"`
	NodeCapture                     *int  `json:"nodeCapture,omitempty"`
	LargestMultiKill                *int  `json:"largestMultiKill,omitempty"`
	SightWardsBoughtInGame          *int  `json:"sightWardsBoughtInGame,omitempty"`
	TotalDamageDealtToChampions     *int  `json:"totalDamageDealtToChampions,omitempty"`
	TotalUnitsHealed                *int  `json:"totalUnitsHealed,omitempty"`
	InhibitorKills                  *int  `json:"inhibitorKills,omitempty"`
	TotalScoreRank                  *int  `json:"totalScoreRank,omitempty"`
	TotalDamageTaken                *int  `json:"totalDamageTaken,omitempty"`
	KillingSprees                   *int  `json:"killingSprees,omitempty"`
	TimeCCingOthers                 *int  `json:"timeCCingOthers,omitempty"`
	PhysicalDamageTaken             *int  `json:"physicalDamageTaken,omitempty"`
}

func (t ParticipantStats) String() string {
	return Stringify(t)
}

// Rune represent a RuneDto.
type Rune struct {
	RuneID *int `json:"runeId,omitempty"` // The ID of the rune
	Rank   *int `json:"rank,omitempty"`
	Count   *int `json:"count,omitempty"` // The count of this rune used by the participant
}

func (t Rune) String() string {
	return Stringify(t)
}

// ParticipantTimeline represent a ParticipantTimelineDto.
type ParticipantTimeline struct {
	Lane                        *string             `json:"lane,omitempty"`
	ParticipantID               *int                `json:"participantId,omitempty"`
	CsDiffPerMinDeltas          *map[string]float64 `json:"csDiffPerMinDeltas,omitempty"` // Map[string, double]
	GoldPerMinDeltas            *map[string]float64 `json:"goldPerMinDeltas,omitempty"`   // Map[string, double]
	XpDiffPerMinDeltas          *map[string]float64 `json:"xpDiffPerMinDeltas,omitempty"` // Map[string, double]
	CreepsPerMinDeltas          *map[string]float64 `json:"creepsPerMinDeltas,omitempty"` // Map[string, double]
	XpPerMinDeltas              *map[string]float64 `json:"xpPerMinDeltas,omitempty"`     // Map[string, double]
	Role                        *string             `json:"role,omitempty"`
	DamageTakenDiffPerMinDeltas *map[string]float64 `json:"damageTakenDiffPerMinDeltas,omitempty"` // Map[string, double]
	DamageTakenPerMinDeltas     *map[string]float64 `json:"damageTakenPerMinDeltas,omitempty"`     // Map[string, double]
}

func (t ParticipantTimeline) String() string {
	return Stringify(t)
}

// MatchList represent a MatchListDto.
type MatchList struct {
	Matches    []MatchReference `json:"matches,omitempty"`
	TotalGames *int             `json:"totalGames,omitempty"`
	StartIndex *int             `json:"startIndex,omitempty"`
	EndIndex   *int             `json:"endIndex,omitempty"`
}

func (t MatchList) String() string {
	return Stringify(t)
}

// MatchReference represent a MatchReferenceDto.
type MatchReference struct {
	Lane       *string `json:"lane,omitempty"`
	GameID     *int    `json:"gameId,omitempty"`
	Champion   *int    `json:"champion,omitempty"`
	PlatformID *string `json:"platformId,omitempty"`
	Season     *int    `json:"season,omitempty"`
	Queue      *int    `json:"queue,omitempty"`
	Role       *string `json:"role,omitempty"`
	Timestamp  *int    `json:"timestamp,omitempty"`
}

func (t MatchReference) String() string {
	return Stringify(t)
}

// MatchTimeline represent a MatchTimelineDto.
type MatchTimeline struct {
	Frames        []MatchFrame `json:"frames,omitempty"`
	FrameInterval *int         `json:"frameInterval,omitempty"`
}

func (t MatchTimeline) String() string {
	return Stringify(t)
}

// MatchFrame represent a MatchFrameDto.
type MatchFrame struct {
	Timestamp         *string                        `json:"timestamp,omitempty"`
	ParticipantFrames *map[int]MatchParticipantFrame `json:"participantFrames,omitempty"` // Map[int, MatchParticipantFrame]
	Events            []MatchEvent                   `json:"events,omitempty"`
}

func (t MatchFrame) String() string {
	return Stringify(t)
}

// MatchEvent represent a MatchEventDto.
type MatchEvent struct {
	EventType               *string        `json:"eventType,omitempty"`
	TowerType               *string        `json:"towerType,omitempty"`
	TeamID                  *int           `json:"teamId,omitempty"`
	AscendedType            *string        `json:"ascendedType,omitempty"`
	KillerID                *int           `json:"killerId,omitempty"`
	LevelUpType             *string        `json:"levelUpType,omitempty"`
	PointCaptured           *string        `json:"pointCaptured,omitempty"`
	AssistingParticipantIds []int          `json:"assistingParticipantIds,omitempty"`
	WardType                *string        `json:"wardType,omitempty"`
	MonsterType             *string        `json:"monsterType,omitempty"`
	Type                    *string        `json:"type,omitempty"`
	SkillSlot               *int           `json:"skillSlot,omitempty"`
	VictimID                *int           `json:"victimId,omitempty"`
	Timestamp               *int           `json:"timestamp,omitempty"`
	AfterID                 *int           `json:"afterId,omitempty"`
	MonsterSubType          *string        `json:"monsterSubType,omitempty"`
	LaneType                *string        `json:"laneType,omitempty"`
	ItemID                  *int           `json:"itemId,omitempty"`
	ParticipantID           *int           `json:"participantId,omitempty"`
	BuildingType            *string        `json:"buildingType,omitempty"`
	CreatorID               *int           `json:"creatorId,omitempty"`
	Position                *MatchPosition `json:"position,omitempty"`
	BeforeID                *int           `json:"beforeId,omitempty"`
}

func (t MatchEvent) String() string {
	return Stringify(t)
}

// MatchParticipantFrame represent a MatchParticipantFrameDto.
type MatchParticipantFrame struct {
	TotalGold           *string        `json:"totalGold,omitempty"`
	TeamScore           *int           `json:"teamScore,omitempty"`
	ParticipantID       *int           `json:"participantId,omitempty"`
	Level               *string        `json:"level,omitempty"`
	CurrentGold         *int           `json:"currentGold,omitempty"`
	MinionsKilled       *int           `json:"minionsKilled,omitempty"`
	DominionScore       *string        `json:"dominionScore,omitempty"`
	Position            *MatchPosition `json:"position,omitempty"`
	Xp                  *int           `json:"xp,omitempty"`
	JungleMinionsKilled *int           `json:"jungleMinionsKilled,omitempty"`
}

func (t MatchParticipantFrame) String() string {
	return Stringify(t)
}

// MatchPosition represent a MatchPositionDto.
type MatchPosition struct {
	X *int `json:"x,omitempty"`
	Y *int `json:"y,omitempty"`
}

func (t MatchPosition) String() string {
	return Stringify(t)
}

// GetMatchesByMatchID fetches match by match ID
func (s *MatchService) GetMatchesByMatchID(ctx context.Context, matchID int) (*Match, *Response, error) {
	c := fmt.Sprintf("%v/matches/%v", s.client.MatchURL, matchID)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Match)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetMatchListByAccountID fetches matchlist for ranked games played on given account ID and platform ID and filtered using given filter parameters, if any
func (s *MatchService) GetMatchListByAccountID(ctx context.Context, accountID int) (*MatchList, *Response, error) {
	c := fmt.Sprintf("%v/matchlists/by-account/%v", s.client.MatchURL, accountID)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(MatchList)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetMatchListRecentByAccountID fetches matchlist for last 20 matches played on given account ID and platform ID.
func (s *MatchService) GetMatchListRecentByAccountID(ctx context.Context, accountID int) (*MatchList, *Response, error) {
	c := fmt.Sprintf("%v/matchlists/by-account/%v/recent", s.client.MatchURL, accountID)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(MatchList)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetTimelineByMatchID fetches match timeline by match ID.
func (s *MatchService) GetTimelineByMatchID(ctx context.Context, matchID int) (*MatchTimeline, *Response, error) {
	c := fmt.Sprintf("%v/timelines/by-match/%v", s.client.MatchURL, matchID)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(MatchTimeline)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetMatchIdsByTournamentCode fetches a collection of match IDs by tournament code
func (s *MatchService) GetMatchIdsByTournamentCode(ctx context.Context, tournamentCode string) ([]int, *Response, error) {
	c := fmt.Sprintf("%v/matches/by-tournament-code/%v/ids", s.client.MatchURL, tournamentCode)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	var uResp []int
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetMatchesByTournamentCode fetches match by match ID and tournament code
func (s *MatchService) GetMatchesByTournamentCode(ctx context.Context, tournamentCode string) (*Match, *Response, error) {
	c := fmt.Sprintf("%v/matches/by-tournament-code/%v", s.client.MatchURL, tournamentCode)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Match)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}
