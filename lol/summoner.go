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

// Get fetches a user. Passing the empty string will fetch the authenticated
// user.
//
// GitHub API docs: https://developer.github.com/v3/users/#get-a-single-user
func (s *SummonerService) Get(ctx context.Context, summoner string) (*Summoner, *Response, error) {
	var c string
	if summoner != "" {
		c = fmt.Sprintf("summoner/v3/summoners/by-name/%v?api_key=%s", summoner, s.client.keyLol)
	} else {
		return nil, nil ,errors.New("Need to set a champion name")
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