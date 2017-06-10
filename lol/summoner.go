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

func (s *SummonerService) Get(ctx context.Context, param string, method string) (*Summoner, *Response, error) {
	var c string

	switch method{
	case "Account":
		if param == "" {
			return nil, nil ,errors.New("Need to set a Account ID")
		}
		c = fmt.Sprintf("%v/by-account/%v?api_key=%s", s.client.SummonerURL, param, s.client.keyLol)
	case "ID":
		if param == "" {
			return nil, nil ,errors.New("Need to set a Summoner ID")
		}
		c = fmt.Sprintf("%v/%v?api_key=%s", s.client.SummonerURL,param, s.client.keyLol)
	default:
		if param == "" {
			return nil, nil ,errors.New("Need to set a Summoner Name")
		}
		c = fmt.Sprintf("%v/by-name/%v?api_key=%s", s.client.SummonerURL,param, s.client.keyLol)
	}

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
