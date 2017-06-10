package lol


import (
	"context"
	"fmt"
	"errors"
)

type SummonerService service

type Summoner struct {
	ProfileIconId	*int	 `json:"profileIconId,omitempty"` 	//ID of the summoner icon associated with the summoner.
	Name		    *string	 `json:"name,omitempty"` 		    //Summoner name.
	SummonerLevel	*int	 `json:"summonerLevel,omitempty"` 	//Summoner level associated with the summoner.
	RevisionDate	*int	 `json:"revisionDate,omitempty"`	//Date summoner was last modified specified as epoch milliseconds. The following events will update this timestamp: profile icon change, playing the tutorial or advanced tutorial, finishing a game, summoner name change
	ID		        *int	 `json:"id,omitempty"` 			    //Summoner ID.
	AccountID	    *int	 `json:"accountId,omitempty"` 		//Account ID.
}

// Get fetches a summoner. Passing the empty string will fetch the authenticated
// summoner.
// Methods
//  Name
//  Account
//  Id
// By Default will be Name



func (s *SummonerService) GetByName(ctx context.Context, name string) (interface{}, *Response, error) {
	if name == "" {

		return nil, nil ,errors.New("Need to set a Summoner Name")

	}else{

		c := fmt.Sprintf("%v/by-name/%v", s.client.SummonerURL,name)

		return getSummoner(s, ctx, c, new(Summoner))

	}
}

func (s *SummonerService) GetByID(ctx context.Context, id int) (interface{}, *Response, error) {

		c := fmt.Sprintf("%v/%v", s.client.SummonerURL, id)
		return getSummoner(s, ctx, c, new(Summoner))

}

func (s *SummonerService) GetByAccountID(ctx context.Context, id int) (interface{}, *Response, error) {

	c := fmt.Sprintf("%v/by-account/%v", s.client.SummonerURL, id)
	return getSummoner(s, ctx, c, new(Summoner))

}

func getSummoner(s *SummonerService, ctx context.Context, c string, inter interface{}) (interface{}, *Response, error){
	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	resp, err := s.client.Do(ctx, req, inter)
	if err != nil {
		return nil, resp, err
	}

	return inter, resp, nil
}

func (s *SummonerService) GetImgURL(profileIcon int) string{

	return fmt.Sprintf("%v/%v.png", s.client.ProfileIconURL.String(), profileIcon)

}
