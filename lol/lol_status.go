package lol

import (
	"context"
	"fmt"
)

// LolStatusService represents a service to interact with LolStatus API.
type LolStatusService service

// ShardStatus represent a ShardStatus.
type ShardStatus struct {
	Name      *string   `json:"name,omitempty"`
	RegionTag *string   `json:"region_tag,omitempty"`
	Hostname  *string   `json:"hostname,omitempty"`
	Services  []Service `json:"services,omitempty"`
	Slug      *string   `json:"slug,omitempty"`
	Locales   []string  `json:"locales,omitempty"`
}

func (t ShardStatus) String() string {
	return Stringify(t)
}

// Service represent a Service.
type Service struct {
	Status    *string    `json:"status,omitempty"`
	Incidents []Incident `json:"incidents,omitempty"`
	Name      *string    `json:"name,omitempty"`
	Slug      *string    `json:"slug,omitempty"`
}

func (t Service) String() string {
	return Stringify(t)
}

// Incident represent a Incident.
type Incident struct {
	Active    *bool     `json:"active,omitempty"`
	CreatedAt *string   `json:"created_at,omitempty"`
	ID        *int      `json:"id,omitempty"`
	Updates   []Message `json:"updates,omitempty"`
}

func (t Incident) String() string {
	return Stringify(t)
}

// Message represent a Message.
type Message struct {
	Severity     *string       `json:"severity,omitempty"`
	Author       *string       `json:"author,omitempty"`
	CreatedAt    *string       `json:"created_at,omitempty"`
	Translations []Translation `json:"translations,omitempty"`
	UpdatedAt    *string       `json:"updated_at,omitempty"`
	Content      *string       `json:"content,omitempty"`
	ID           *string       `json:"id,omitempty"`
}

func (t Message) String() string {
	return Stringify(t)
}

// Translation represent a Translation.
type Translation struct {
	Locale    *string `json:"locale,omitempty"`
	Content   *string `json:"content,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty"`
}

func (t Translation) String() string {
	return Stringify(t)
}

// GetShardStatus fetches the League of Legends status for the given shard.
func (s *LolStatusService) GetShardStatus(ctx context.Context) (*ShardStatus, *Response, error) {
	c := fmt.Sprintf("%v/shard-data", s.client.LolStatusURL)

	req, err := s.client.NewRequest("GET", c, nil)
	if err != nil {
		return nil, nil, err
	}

	uResp := new(ShardStatus)
	resp, err := s.client.Do(ctx, req, &uResp)
	if err != nil {
		return nil, resp, err
	}

	return uResp, resp, nil
}
