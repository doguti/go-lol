package lol


import (
	"context"
	"fmt"
	"errors"
)

type SummonerService service

// Summoner represent a SummonerDTO.
type Summoner struct {
	ProfileIconId *int     `json:"profileIconId,omitempty"` //ID of the summoner icon associated with the summoner.
	Name          *string     `json:"name,omitempty"`       //Summoner name.
	SummonerLevel *int     `json:"summonerLevel,omitempty"` //Summoner level associated with the summoner.
	RevisionDate  *int     `json:"revisionDate,omitempty"`  //Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: profile icon change, playing the tutorial or advanced tutorial, finishing a game, summoner name change
	ID            *int     `json:"id,omitempty"`            //Summoner ID.
	AccountID     *int     `json:"accountId,omitempty"`     //Account ID.
}

func (t Summoner) String() string {
	return Stringify(t)
}

// Get fetches a summoner. Passing the empty string will fetch the authenticated
// summoner.
// Methods
//  Name
//  Account
//  Id
// By Default will be Name
func (s *SummonerService) GetByName(ctx context.Context, name string) (*Summoner, *Response, error) {
	if name == "" {
		return nil, nil ,errors.New("Need to set a Summoner Name")

	}else{
		c := fmt.Sprintf("%v/by-name/%v", s.client.SummonerURL,name)
		return getSummoner(s, ctx, c)
	}
}

func (s *SummonerService) GetByID(ctx context.Context, id int) (*Summoner, *Response, error) {
		c := fmt.Sprintf("%v/%v", s.client.SummonerURL, id)
		return getSummoner(s, ctx, c)
}

func (s *SummonerService) GetByAccountID(ctx context.Context, id int) (*Summoner, *Response, error) {
	c := fmt.Sprintf("%v/by-account/%v", s.client.SummonerURL, id)
	return getSummoner(s, ctx, c)
}

func getSummoner(s *SummonerService, ctx context.Context, c string) (*Summoner, *Response, error){
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

func (s *SummonerService) GetImgURL(profileIcon int) string{
	return fmt.Sprintf("%v/%v.png", s.client.ProfileIconURL.String(), profileIcon)
}
