package lol

import (
	"context"
	"fmt"
)

// SpectatorService represents a service to interact with Spectator API.
type SpectatorService service

// CurrentGameInfo represent a CurrentGameInfo.
type CurrentGameInfo struct {
	GameID            *int                     `json:"gameId,omitempty"`            // The ID of the game
	GameStartTime     *int                     `json:"gameStartTime,omitempty"`     // The game start time represented in epoch milliseconds
	PlatformID        *string                  `json:"platformId,omitempty"`        // The ID of the platform on which the game is being played
	GameMode          *string                  `json:"gameMode,omitempty"`          // The game mode
	MapID             *int                     `json:"mapId,omitempty"`             // The ID of the map
	GameType          *string                  `json:"gameType,omitempty"`          // The game type
	BannedChampions   []BannedChampion         `json:"bannedChampions,omitempty"`   // Banned champion information
	Observers         *Observer                `json:"observers,omitempty"`         // The observer information
	Participants      []CurrentGameParticipant `json:"participants,omitempty"`      // The participant information
	GameLength        *int                     `json:"gameLength,omitempty"`        // The amount of time in seconds that has passed since the game started
	GameQueueConfigID *int                     `json:"gameQueueConfigId,omitempty"` // The queue type (queue types are documented on the Game Constants page)
}

func (t CurrentGameInfo) String() string {
	return Stringify(t)
}

// BannedChampion represent a BannedChampion.
type BannedChampion struct {
	PickTurn   *int `json:"pickTurn,omitempty"`   // The turn during which the champion was banned
	ChampionID *int `json:"championId,omitempty"` // The ID of the banned champion
	TeamID     *int `json:"teamId,omitempty"`     // The ID of the team that banned the champion
}

func (t BannedChampion) String() string {
	return Stringify(t)
}

// Observer represent a Observer.
type Observer struct {
	EncryptionKey *string `json:"encryptionKey,omitempty"` // 	Key used to decrypt the spectator grid game data for playback
}

func (t Observer) String() string {
	return Stringify(t)
}

// CurrentGameParticipant represent a CurrentGameParticipant.
type CurrentGameParticipant struct {
	ProfileIconID *int    `json:"profileIconId,omitempty"` // The ID of the profile icon used by this participant
	ChampionID    *int    `json:"championId,omitempty"`    // The ID of the champion played by this participant
	SummonerName  *string `json:"summonerName,omitempty"`  // 	The summoner name of this participant
	Runes         []Rune  `json:"runes,omitempty"`         // 	The runes used by this participant
	Bot           *bool   `json:"bot,omitempty"`           // Flag indicating whether or not this participant is a bot
	TeamID        *int    `json:"teamId,omitempty"`        // The team ID of this participant, indicating the participant's team
	Spell2ID      *int    `json:"spell2Id,omitempty"`      // The ID of the second summoner spell used by this participant
	Masteries     *int    `json:"masteries,omitempty"`     // The masteries used by this participant
	Spell1ID      *int    `json:"spell1Id,omitempty"`      // The ID of the first summoner spell used by this participant
	SummonerID    *int    `json:"summonerId,omitempty"`    // The summoner ID of this participant
}

func (t CurrentGameParticipant) String() string {
	return Stringify(t)
}

// FeaturedGames represent a FeaturedGames.
type FeaturedGames struct {
	ClientRefreshInterval *int               `json:"clientRefreshInterval,omitempty"` // The suggested interval to wait before requesting FeaturedGames again
	GameList              []FeaturedGameInfo `json:"gameList,omitempty"`              // The list of featured games
}

func (t FeaturedGames) String() string {
	return Stringify(t)
}

// FeaturedGameInfo represent a FeaturedGameInfo.
type FeaturedGameInfo struct {
	GameID            *int             `json:"gameId,omitempty"`            // The ID of the game
	GameStartTime     *int             `json:"gameStartTime,omitempty"`     // The game start time represented in epoch milliseconds
	PlatformID        *string          `json:"platformId,omitempty"`        // The ID of the platform on which the game is being played
	GameMode          *string          `json:"gameMode,omitempty"`          // The game mode
	MapID             *int             `json:"mapId,omitempty"`             // 	The ID of the map
	GameType          *string          `json:"gameType,omitempty"`          // The game type
	BannedChampions   []BannedChampion `json:"bannedChampions,omitempty"`   // Banned champion information
	Observers         *Observer        `json:"observers,omitempty"`         // The observer information
	Participants      []Participant    `json:"participants,omitempty"`      // The participant information
	GameLength        *int             `json:"gameLength,omitempty"`        // The amount of time in seconds that has passed since the game started
	GameQueueConfigID *int             `json:"gameQueueConfigId,omitempty"` // The queue type (queue types are documented on the Game Constants page)
}

func (t FeaturedGameInfo) String() string {
	return Stringify(t)
}

// Participant represent a Participant.
type Participant struct {
	ProfileIconID *int    `json:"profileIconId,omitempty"` // The ID of the profile icon used by this participant
	ChampionID    *int    `json:"championId,omitempty"`    // The ID of the champion played by this participant
	SummonerName  *string `json:"summonerName,omitempty"`  // 	The summoner name of this participant
	Bot           *bool   `json:"bot,omitempty"`           // Flag indicating whether or not this participant is a bot
	Spell2ID      *int    `json:"spell2Id,omitempty"`      // The ID of the second summoner spell used by this participant
	TeamID        *int    `json:"teamId,omitempty"`        // The team ID of this participant, indicating the participant's tea
	Spell1ID      *int    `json:"spell1Id,omitempty"`      // The ID of the first summoner spell used by this participant
}

func (t Participant) String() string {
	return Stringify(t)
}

// GetActiveGamesBySummonerID fetches a current game information for the given summoner ID
func (s *SpectatorService) GetActiveGamesBySummonerID(ctx context.Context, ID int) (*CurrentGameInfo, *Response, error) {
	c := fmt.Sprintf("%v/active-games/by-summoner/%v", s.client.SpectatorURL, ID)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(CurrentGameInfo)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetFeaturedGames fetches a list of featured games
func (s *SpectatorService) GetFeaturedGames(ctx context.Context) (*FeaturedGames, *Response, error) {
	c := fmt.Sprintf("%v/featured-games", s.client.SpectatorURL)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(FeaturedGames)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}
