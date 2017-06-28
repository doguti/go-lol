package lol

import (
	"context"
	"fmt"
)

// LeagueService represents a service to interact with League API.
type LeagueService service

// LeagueList represent a LeagueListDTO.
type LeagueList struct {
	Tier    *string     `json:"tier,omitempty"`
	Queue   *string     `json:"queue,omitempty"`
	Name    *string     `json:"name,omitempty"`
	Entries []LeagueItem     `json:"pages,omitempty"`
}

func (t LeagueList) String() string {
	return Stringify(t)
}

// LeagueItem represent a LeagueItemDTO.
type LeagueItem struct {
	Rank             *string     `json:"rank,omitempty"`
	HotStreak        *bool     `json:"hotStreak,omitempty"`
	MiniSeries       *MiniSeries     `json:"miniSeries,omitempty"`
	Wins             *int     `json:"wins,omitempty"`
	Veteran          *bool     `json:"veteran,omitempty"`
	Losses           *int     `json:"losses,omitempty"`
	PlayerOrTeamID   *string     `json:"playerOrTeamId,omitempty"`
	PlayerOrTeamName *string     `json:"playerOrTeamName,omitempty"`
	Inactive         *bool     `json:"inactive,omitempty"`
	FreshBlood       *bool     `json:"freshBlood,omitempty"`
	LeaguePoints     *int     `json:"leaguePoints,omitempty"`
}

func (t LeagueItem) String() string {
	return Stringify(t)
}

// MiniSeries represent a MiniSeriesDTO.
type MiniSeries struct {
	Wins     *int     `json:"wins,omitempty"`
	Losses   *int     `json:"losses,omitempty"`
	Target   *int     `json:"target,omitempty"`
	Progress *string     `json:"progress,omitempty"`
}

func (t MiniSeries) String() string {
	return Stringify(t)
}

// LeaguePosition represent a LeaguePositionDTO.
type LeaguePosition struct {
	Rank             *string     `json:"rank,omitempty"`
	QueueType        *string     `json:"queueType,omitempty"`
	HotStreak        *bool     `json:"hotStreak,omitempty"`
	MiniSeries       *MiniSeries     `json:"miniSeries,omitempty"`
	Wins             *int     `json:"wins,omitempty"`
	Veteran          *bool     `json:"veteran,omitempty"`
	Losses           *int     `json:"losses,omitempty"`
	PlayerOrTeamID   *string     `json:"playerOrTeamId,omitempty"`
	LeagueName       *string     `json:"leagueName,omitempty"`
	PlayerOrTeamName *string     `json:"playerOrTeamName,omitempty"`
	Inactive         *bool     `json:"inactive,omitempty"`
	FreshBlood       *bool     `json:"freshBlood,omitempty"`
	Tier             *string     `json:"tier,omitempty"`
	LeaguePoints     *int     `json:"leaguePoints,omitempty"`
}

func (t LeaguePosition) String() string {
	return Stringify(t)
}

// GetChallengerLeaguesByQueue fetches the challenger league for a given queue.
func (s *LeagueService) GetChallengerLeaguesByQueue(ctx context.Context, queueString string) (*LeagueList, *Response, error) {
	c := fmt.Sprintf("%v/challengerleagues/by-queue/%v", s.client.LeagueURL, queueString)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(LeagueList)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetLeaguesBySummoner fetches a collection of leagues in all queues for a given summoner ID
func (s *LeagueService) GetLeaguesBySummoner(ctx context.Context, id int) ([]*LeagueList, *Response, error) {
	c := fmt.Sprintf("%v/leagues/by-summoner/%v", s.client.LeagueURL, id)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	var leagues []*LeagueList
	resp, err := s.client.Do(ctx, req, &leagues)
	if err != nil {
		return nil, resp, err
	}

	return leagues, resp, nil
}

// GetMasterLeaguesByQueue fetches the master league for a given queue
func (s *LeagueService) GetMasterLeaguesByQueue(ctx context.Context, queue string) (*LeagueList, *Response, error) {
	c := fmt.Sprintf("%v/masterleagues/by-queue/%v", s.client.LeagueURL, queue)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(LeagueList)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetPositionsBySummoner fetches a collection of league positions in all queues for a given summoner ID
func (s *LeagueService) GetPositionsBySummoner(ctx context.Context, id int) ([]*LeaguePosition, *Response, error) {
	c := fmt.Sprintf("%v/positions/by-summoner/%v", s.client.LeagueURL, id)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	var leagues []*LeaguePosition
	resp, err := s.client.Do(ctx, req, &leagues)
	if err != nil {
		return nil, resp, err
	}

	return leagues, resp, nil
}
