package lol

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestShardStatus_String(t *testing.T) {
	testJSONMarshal(t, &ShardStatus{}, "{}")

	translations := []Translation{
		{
			Locale:     String("es_ES"),
			Content:    String("word"),
			UpdatedAt: String("20170710"),
		},
	}

	messages := []Message{
		{
			Severity:     String("low"),
			Author:       String("root"),
			CreatedAt:   String("20170710"),
			Translations: translations,
			UpdatedAt:   String("20170710"),
			Content:      String("word"),
			ID:           String("123456"),
		},
	}

	incidents := []Incident{
		{
			Active:     Bool(true),
			CreatedAt: String("20170710"),
			ID:         Int(123456),
			Updates:    messages,
		},
	}

	services := []Service{
		{
			Status:    String("online"),
			Incidents: incidents,
			Name:      String("Client"),
			Slug:      String("client"),
		},
	}

	locales := []string{"en_GB", "es_ES"}

	shardStatus := &ShardStatus{
		Name:       String("EU West"),
		RegionTag: String("eu"),
		Hostname:   String("prod.euw1.lol.riotgames.com"),
		Services:   services,
		Slug:       String("euw"),
		Locales:    locales,
	}

	want := `{
		"name":"EU West",
		"region_tag":"eu",
		"hostname":"prod.euw1.lol.riotgames.com",
		"services":[
			{
				"status":"online",
				"incidents":[
					{
						"active":true,
						"created_at":"20170710",
						"id":123456,
						"updates":[
							{
								"severity":"low",
								"author":"root",
								"created_at":"20170710",
								"translations":[
									{
										"locale":"es_ES",
										"content":"word",
										"updated_at":"20170710"
									}
								],
								"updated_at":"20170710",
								"content":"word",
								"id":"123456"
							}
						]
					}
				],
				"name":"Client",
				"slug":"client"
			}
		],
		"slug":"euw",
		"locales":[
			"en_GB",
			"es_ES"
		]
	}`

	testJSONMarshal(t, shardStatus, want)
}

func TestLolStatusService_GetShardStatus(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.LolStatusURL+"/shard-data", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"name":"EU West"}`)
	})

	lolStatus, _, err := client.LolStatus.GetShardStatus(context.Background())
	if err != nil {
		t.Errorf("LolStatus.GetShardStatus returned error: %v", err)
	}

	want := &ShardStatus{Name: String("EU West")}
	if !reflect.DeepEqual(lolStatus, want) {
		t.Errorf("LolStatus.GetShardStatus returned %+v, want %+v", lolStatus, want)
	}
}
