package lol

import (
	"context"
	"fmt"
)

type MasteriesService service

// Masteries represent a MasteryPagesDto.
type Masteries struct {
	SummonerId *int     `json:"summonerId,omitempty"`     // Summoner name.
	Pages      []MasteryPage     `json:"pages,omitempty"` //Summoner name.
}

func (t Masteries) String() string {
	return Stringify(t)
}

// MasteryPage represent a MasteryPageDto.
type MasteryPage struct {
	Current   *bool     `json:"current,omitempty"`   // Summoner name.
	Masteries []Mastery     `json:"masteries,omitempty"` // Summoner name.
	Name      *string     `json:"name,omitempty"`      // Summoner name.
	Id        *int     `json:"id,omitempty"`        // Summoner name.
}

func (t MasteryPage) String() string {
	return Stringify(t)
}

// Mastery represent a MasteryDto.
type Mastery struct {
	Id   *int     `json:"id,omitempty"`   // Summoner name.
	Rank *int     `json:"rank,omitempty"` //Summoner name.
}

func (t Mastery) String() string {
	return Stringify(t)
}

func (s *MasteriesService) GetBySummonerId(ctx context.Context, id int) (*Masteries, *Response, error) {
	c := fmt.Sprintf("%v/by-summoner/%v", s.client.MasteriesURL, id)
	return getMasteries(s, ctx, c)
}

func getMasteries(s *MasteriesService, ctx context.Context, c string) (*Masteries, *Response, error){
	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Masteries)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}
