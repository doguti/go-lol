package lol


import (
	"context"
	"fmt"
	"errors"
)

// SummonerService represents a service to interact with Summoner API.
type SummonerService service

// Summoner represent a SummonerDTO.
type Summoner struct {
	ProfileIconID *int     `json:"profileIconId,omitempty"` //ID of the summoner icon associated with the summoner.
	Name          *string     `json:"name,omitempty"`       //Summoner name.
	SummonerLevel *int     `json:"summonerLevel,omitempty"` //Summoner level associated with the summoner.
	RevisionDate  *int     `json:"revisionDate,omitempty"`  //Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: profile icon change, playing the tutorial or advanced tutorial, finishing a game, summoner name change
	ID            *int     `json:"id,omitempty"`            //Summoner ID.
	AccountID     *int     `json:"accountId,omitempty"`     //Account ID.
}

func (t Summoner) String() string {
	return Stringify(t)
}

// GetByName fetches a summoner by summoner name
func (s *SummonerService) GetByName(ctx context.Context, name string) (*Summoner, *Response, error) {
	if name == "" {
		return nil, nil ,errors.New("Need to set a Summoner Name")
	}

	c := fmt.Sprintf("%v/by-name/%v", s.client.SummonerURL,name)
	return getSummoner(ctx, s, c)
}

// GetByID fetches a summoner by summoner ID
func (s *SummonerService) GetByID(ctx context.Context, id int) (*Summoner, *Response, error) {
	c := fmt.Sprintf("%v/%v", s.client.SummonerURL, id)
	return getSummoner(ctx, s, c)
}

// GetByAccountID fetches a summoner by account ID
func (s *SummonerService) GetByAccountID(ctx context.Context, id int) (*Summoner, *Response, error) {
	c := fmt.Sprintf("%v/by-account/%v", s.client.SummonerURL, id)
	return getSummoner(ctx, s, c)
}

// getSummoner fetches a summoner by string
func getSummoner(ctx context.Context, s *SummonerService, c string) (*Summoner, *Response, error){
	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(Summoner)
	resp, err := s.client.Do(ctx, req, uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}

// GetImgURL fetches an URL for profileIcon
func (s *SummonerService) GetImgURL(profileIcon int) string{
	return fmt.Sprintf("%v/%v.png", s.client.ProfileIconURL.String(), profileIcon)
}
