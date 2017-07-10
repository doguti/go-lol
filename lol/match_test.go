package lol

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"strings"
	"testing"
)

func TestMatch_marshall(t *testing.T) {
	testJSONMarshal(t, &Match{}, "{}")

	player := &Player{
		CurrentPlatformID: String("World"),
		SummonerName:      String("World"),
		MatchHistoryURI:   String("World"),
		PlatformID:        String("World"),
		CurrentAccountID:  Int(3),
		ProfileIcon:       Int(3),
		SummonerID:        Int(3),
		AccountID:         Int(3),
	}

	participantIdentities := []ParticipantIdentity{
		{
			Player:        player,
			ParticipantID: Int(3),
		},
	}

	teamBanses := []TeamBans{
		{
			PickTurn:   Int(1),
			ChampionID: Int(3),
		},
	}

	teamStatses := []TeamStats{
		{
			FirstDragon:          Bool(true),
			FirstInhibitor:       Bool(true),
			Bans:                 teamBanses,
			BaronKills:           Int(1),
			FirstRiftHerald:      Bool(true),
			FirstBaron:           Bool(true),
			RiftHeraldKills:      Int(1),
			FirstBlood:           Bool(true),
			TeamID:               Int(1),
			FirstTower:           Bool(true),
			VilemawKills:         Int(1),
			InhibitorKills:       Int(1),
			TowerKills:           Int(1),
			DominionVictoryScore: Int(1),
			Win:                  String("Winner"),
			DragonKills:          Int(1),
		},
	}

	participantStats := &ParticipantStats{
		PhysicalDamageDealt:            Int(1),
		NeutralMinionsKilledTeamJungle: Int(1),
		MagicDamageDealt:               Int(1),
		TotalPlayerScore:               Int(1),
		Deaths:                         Int(1),
		Win:                            Bool(true),
		NeutralMinionsKilledEnemyJungle: Int(1),
		AltarsCaptured:                  Int(1),
		LargestCriticalStrike:           Int(1),
		TotalDamageDealt:                Int(1),
		MagicDamageDealtToChampions:     Int(1),
		VisionWardsBoughtInGame:         Int(1),
		DamageDealtToObjectives:         Int(1),
		LargestKillingSpree:             Int(1),
		Item1:                           Int(1),
		QuadraKills:                     Int(1),
		TeamObjective:                   Int(1),
		TotalTimeCrowdControlDealt:      Int(1),
		LongestTimeSpentLiving:          Int(1),
		WardsKilled:                     Int(1),
		FirstTowerAssist:                Bool(true),
		FirstTowerKill:                  Bool(true),
		Item2:                           Int(1),
		Item3:                           Int(1),
		Item0:                           Int(1),
		FirstBloodAssist:                Bool(true),
		VisionScore:                     Int(1),
		WardsPlaced:                     Int(1),
		Item4:                           Int(1),
		Item5:                           Int(1),
		Item6:                           Int(1),
		TurretKills:                     Int(1),
		TripleKills:                     Int(1),
		DamageSelfMitigated:             Int(1),
		ChampLevel:                      Int(1),
		NodeNeutralizeAssist:            Int(1),
		FirstInhibitorKill:              Bool(true),
		GoldEarned:                      Int(1),
		MagicalDamageTaken:              Int(1),
		Kills:                           Int(1),
		DoubleKills:                     Int(1),
		NodeCaptureAssist:               Int(1),
		TrueDamageTaken:                 Int(1),
		NodeNeutralize:                  Int(1),
		FirstInhibitorAssist:            Bool(true),
		Assists:                         Int(1),
		UnrealKills:                     Int(1),
		NeutralMinionsKilled:            Int(1),
		ObjectivePlayerScore:            Int(1),
		CombatPlayerScore:               Int(1),
		DamageDealtToTurrets:            Int(1),
		AltarsNeutralized:               Int(1),
		PhysicalDamageDealtToChampions:  Int(1),
		GoldSpent:                       Int(1),
		TrueDamageDealt:                 Int(1),
		TrueDamageDealtToChampions:      Int(1),
		ParticipantID:                   Int(1),
		PentaKills:                      Int(1),
		TotalHeal:                       Int(1),
		TotalMinionsKilled:              Int(1),
		FirstBloodKill:                  Bool(true),
		NodeCapture:                     Int(1),
		LargestMultiKill:                Int(1),
		SightWardsBoughtInGame:          Int(1),
		TotalDamageDealtToChampions:     Int(1),
		TotalUnitsHealed:                Int(1),
		InhibitorKills:                  Int(1),
		TotalScoreRank:                  Int(1),
		TotalDamageTaken:                Int(1),
		KillingSprees:                   Int(1),
		TimeCCingOthers:                 Int(1),
		PhysicalDamageTaken:             Int(1),
	}

	mapStrFloat := &map[string]float64{
		"value1": float64(1.1),
		"value2": float64(1.2),
	}

	participantTimeline := &ParticipantTimeline{
		Lane:               String("World"),
		ParticipantID:      Int(1),
		CsDiffPerMinDeltas: mapStrFloat,
		GoldPerMinDeltas:   mapStrFloat,
		XpDiffPerMinDeltas: mapStrFloat,
		CreepsPerMinDeltas: mapStrFloat,
		XpPerMinDeltas:     mapStrFloat,
		Role:               String("Wizard"),
		DamageTakenDiffPerMinDeltas: mapStrFloat,
		DamageTakenPerMinDeltas:     mapStrFloat,
	}

	runes := []Rune{
		{
			RuneID: Int(12),
			Rank:   Int(321123),
		},
		{
			RuneID: Int(13),
			Rank:   Int(111222),
		},
	}

	masteries := []Mastery{
		{
			ID:   Int(12),
			Rank: Int(321123),
		},
		{
			ID:   Int(13),
			Rank: Int(111222),
		},
	}

	participants := []ParticipantDto{
		{
			Stats:                     participantStats,
			ParticipantID:             Int(1),
			Runes:                     runes,
			Timeline:                  participantTimeline,
			TeamID:                    Int(1),
			Spell2ID:                  Int(1),
			Masteries:                 masteries,
			HighestAchievedSeasonTier: String("season"),
			Spell1ID:                  Int(1),
			ChampionID:                Int(1),
		},
	}

	match := &Match{
		SeasonID:              Int(1),
		QueueID:               Int(1),
		GameID:                Int(1),
		ParticipantIdentities: participantIdentities,
		GameVersion:           String("1.0"),
		PlatformID:            String("1234"),
		GameMode:              String("World"),
		MapID:                 Int(2),
		GameType:              String("World"),
		Teams:                 teamStatses,
		Participants:          participants,
		GameDuration:          Int(1),
		GameCreation:          Int(1),
	}

	want := `{
		"seasonId":1,
		"queueId":1,
		"gameId":1,
		"participantIdentities":[
			{
				"player":{
					"currentPlatformId":"World",
					"summonerName":"World",
					"matchHistoryUri":"World",
					"platformId":"World",
					"currentAccountId":3,
					"profileIcon":3,
					"summonerId":3,
					"accountId":3
				},
				"participantId":3
			}
		],
		"gameVersion":"1.0",
		"platformId":"1234",
		"gameMode":"World",
		"mapId":2,
		"gameType":"World",
		"teams":[
			{
				"firstDragon":true,
				"firstInhibitor":true,
				"bans":[
					{
						"pickTurn":1,
						"championId":3
					}
				],
				"baronKills":1,
				"firstRiftHerald":true,
				"firstBaron":true,
				"riftHeraldKills":1,
				"firstBlood":true,
				"teamId":1,
				"firstTower":true,
				"vilemawKills":1,
				"inhibitorKills":1,
				"towerKills":1,
				"dominionVictoryScore":1,
				"win":"Winner",
				"dragonKills":1
			}
		],
		"participants":[
			{
				"stats":{
					"physicalDamageDealt":1,
					"neutralMinionsKilledTeamJungle":1,
					"magicDamageDealt":1,
					"totalPlayerScore":1,
					"deaths":1,
					"win":true,
					"neutralMinionsKilledEnemyJungle":1,
					"altarsCaptured":1,
					"largestCriticalStrike":1,
					"totalDamageDealt":1,
					"magicDamageDealtToChampions":1,
					"visionWardsBoughtInGame":1,
					"damageDealtToObjectives":1,
					"largestKillingSpree":1,
					"item1":1,
					"quadraKills":1,
					"teamObjective":1,
					"totalTimeCrowdControlDealt":1,
					"longestTimeSpentLiving":1,
					"wardsKilled":1,
					"firstTowerAssist":true,
					"firstTowerKill":true,
					"item2":1,
					"item3":1,
					"item0":1,
					"firstBloodAssist":true,
					"visionScore":1,
					"wardsPlaced":1,
					"item4":1,
					"item5":1,
					"item6":1,
					"turretKills":1,
					"tripleKills":1,
					"damageSelfMitigated":1,
					"champLevel":1,
					"nodeNeutralizeAssist":1,
					"firstInhibitorKill":true,
					"goldEarned":1,
					"magicalDamageTaken":1,
					"kills":1,
					"doubleKills":1,
					"nodeCaptureAssist":1,
					"trueDamageTaken":1,
					"nodeNeutralize":1,
					"firstInhibitorAssist":true,
					"assists":1,
					"unrealKills":1,
					"neutralMinionsKilled":1,
					"objectivePlayerScore":1,
					"combatPlayerScore":1,
					"damageDealtToTurrets":1,
					"altarsNeutralized":1,
					"physicalDamageDealtToChampions":1,
					"goldSpent":1,
					"trueDamageDealt":1,
					"trueDamageDealtToChampions":1,
					"participantId":1,
					"pentaKills":1,
					"totalHeal":1,
					"totalMinionsKilled":1,
					"firstBloodKill":true,
					"nodeCapture":1,
					"largestMultiKill":1,
					"sightWardsBoughtInGame":1,
					"totalDamageDealtToChampions":1,
					"totalUnitsHealed":1,
					"inhibitorKills":1,
					"totalScoreRank":1,
					"totalDamageTaken":1,
					"killingSprees":1,
					"timeCCingOthers":1,
					"physicalDamageTaken":1
				},
				"participantId":1,
				"runes":[
					{
						"runeId":12,
						"rank":321123
					},
					{
						"runeId":13,
						"rank":111222
					}
				],
				"timeline":{
					"lane":"World",
					"participantId":1,
					"csDiffPerMinDeltas":{
						"value1":1.1,
						"value2":1.2
					},
					"goldPerMinDeltas":{
						"value1":1.1,
						"value2":1.2
					},
					"xpDiffPerMinDeltas":{
						"value1":1.1,
						"value2":1.2
					},
					"creepsPerMinDeltas":{
						"value1":1.1,
						"value2":1.2
					},
					"xpPerMinDeltas":{
						"value1":1.1,
						"value2":1.2
					},
					"role":"Wizard",
					"damageTakenDiffPerMinDeltas":{
						"value1":1.1,
						"value2":1.2
					},
					"damageTakenPerMinDeltas":{
						"value1":1.1,
						"value2":1.2
					}
				},
				"teamId":1,
				"spell2Id":1,
				"masteries":[
					{
						"id":12,
						"rank":321123
					},
					{
						"id":13,
						"rank":111222
					}
				],
				"highestAchievedSeasonTier":"season",
				"spell1Id":1,
				"championId":1
			}
		],
		"gameDuration":1,
		"gameCreation":1
	}`

	testJSONMarshal(t, match, want)
}

func TestMatchList_marshall(t *testing.T) {
	testJSONMarshal(t, &MatchList{}, "{}")

	matchReferences := []MatchReference{
		{
			Lane:       String("World"),
			GameID:     Int(3),
			Champion:   Int(3),
			PlatformID: String("123456"),
			Season:     Int(3),
			Queue:      Int(3),
			Role:       String("World"),
			Timestamp:  Int(3),
		},
	}

	matchList := &MatchList{
		Matches:    matchReferences,
		TotalGames: Int(3),
		StartIndex: Int(3),
		EndIndex:   Int(3),
	}

	want := `{
		"matches":[
			{
				"lane":"World",
				"gameId":3,
				"champion":3,
				"platformId":"123456",
				"season":3,
				"queue":3,
				"role":"World",
				"timestamp":3
			}
		],
		"totalGames":3,
		"startIndex":3,
		"endIndex":3
	}`

	testJSONMarshal(t, matchList, want)
}

func TestMatchTimeline_marshall(t *testing.T) {
	testJSONMarshal(t, &MatchTimeline{}, "{}")

	matchPosition := &MatchPosition{
		X: Int(2),
		Y: Int(2),
	}

	listInt := []int{1, 2}

	matchEvents := []MatchEvent{
		{
			EventType:               String("World"),
			TowerType:               String("World"),
			TeamID:                  Int(3),
			AscendedType:            String("World"),
			KillerID:                Int(3),
			LevelUpType:             String("World"),
			PointCaptured:           String("World"),
			AssistingParticipantIds: listInt,
			WardType:                String("World"),
			MonsterType:             String("World"),
			Type:                    String("World"),
			SkillSlot:               Int(3),
			VictimID:                Int(3),
			Timestamp:               Int(3),
			AfterID:                 Int(3),
			MonsterSubType:          String("World"),
			LaneType:                String("World"),
			ItemID:                  Int(3),
			ParticipantID:           Int(3),
			BuildingType:            String("World"),
			CreatorID:               Int(3),
			Position:                matchPosition,
			BeforeID:                Int(3),
		},
	}

	mapIntMatchParticipantFrame := &map[int]MatchParticipantFrame{
		2:  {},
		11: {},
	}

	matchFrames := []MatchFrame{
		{
			Timestamp:         String("123456789"),
			ParticipantFrames: mapIntMatchParticipantFrame,
			Events:            matchEvents,
		},
	}

	matchTimeline := &MatchTimeline{
		Frames:        matchFrames,
		FrameInterval: Int(3),
	}

	want := `{
		"frames":[
			{
				"timestamp":"123456789",
				"participantFrames":{
					"11":{

					},
					"2":{

					}
				},
				"events":[
					{
						"eventType":"World",
						"towerType":"World",
						"teamId":3,
						"ascendedType":"World",
						"killerId":3,
						"levelUpType":"World",
						"pointCaptured":"World",
						"assistingParticipantIds":[
							1,
							2
						],
						"wardType":"World",
						"monsterType":"World",
						"type":"World",
						"skillSlot":3,
						"victimId":3,
						"timestamp":3,
						"afterId":3,
						"monsterSubType":"World",
						"laneType":"World",
						"itemId":3,
						"participantId":3,
						"buildingType":"World",
						"creatorId":3,
						"position":{
							"x":2,
							"y":2
						},
						"beforeId":3
					}
				]
			}
		],
		"frameInterval":3
	}`

	testJSONMarshal(t, matchTimeline, want)
}

func TestMatchService_GetMatchesByMatchID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.MatchURL+"/matches/23231", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"seasonID":2}`)
	})

	match, _, err := client.Match.GetMatchesByMatchID(context.Background(), 23231)
	if err != nil {
		t.Errorf("Match.GetByMatchID returned error: %v", err)
	}

	want := &Match{SeasonID: Int(2)}
	if !reflect.DeepEqual(match, want) {
		t.Errorf("Match.GetByMatchID returned %+v, want %+v", match, want)
	}
}

func TestMatchService_GetMatchesByTournamentCode(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.MatchURL+"/matches/by-tournament-code/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"seasonID":2}`)
	})

	match, _, err := client.Match.GetMatchesByTournamentCode(context.Background(), "1234")
	if err != nil {
		t.Errorf("Match.GetMatchesByTournamentCode returned error: %v", err)
	}

	want := &Match{SeasonID: Int(2)}
	if !reflect.DeepEqual(match, want) {
		t.Errorf("Match.GetMatchesByTournamentCode returned %+v, want %+v", match, want)
	}
}

func TestMatchService_Get_invalidMatchesByTournamentCode(t *testing.T) {
	_, _, err := client.Match.GetMatchesByTournamentCode(context.Background(), "%")
	testURLParseError(t, err)
}

func TestMatchService_GetMatchIdsByTournamentCode(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.MatchURL+"/matches/by-tournament-code/1234/ids", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `[1]`)
	})

	match, _, err := client.Match.GetMatchIdsByTournamentCode(context.Background(), "1234")

	errStr := fmt.Sprintf("%v", err)
	if err != nil && !strings.Contains(errStr, "[]int") {
		t.Errorf("Match.GetMatchIdsByTournamentCode did not return type: %v", "[]int")
		t.Errorf("Match.GetMatchIdsByTournamentCode returned error: %v", err)
		t.Errorf("Match.GetMatchIdsByTournamentCode returned: %v", match)
	}
}

func TestMatchService_Get_invalidMatchIdsByTournamentCode(t *testing.T) {
	_, _, err := client.Match.GetMatchIdsByTournamentCode(context.Background(), "%")
	testURLParseError(t, err)
}

func TestMatchService_GetMatchListByAccountID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.MatchURL+"/matchlists/by-account/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"totalGames":2}`)
	})

	match, _, err := client.Match.GetMatchListByAccountID(context.Background(), 1234)
	if err != nil {
		t.Errorf("Match.GetMatchListByAccountID returned error: %v", err)
	}

	want := &MatchList{TotalGames: Int(2)}
	if !reflect.DeepEqual(match, want) {
		t.Errorf("Match.GetMatchListByAccountID returned %+v, want %+v", match, want)
	}
}

func TestMatchService_GetMatchListRecentByAccountID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.MatchURL+"/matchlists/by-account/1234/recent", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"totalGames":2}`)
	})

	match, _, err := client.Match.GetMatchListRecentByAccountID(context.Background(), 1234)
	if err != nil {
		t.Errorf("Match.GetMatchListRecentByAccountID returned error: %v", err)
	}

	want := &MatchList{TotalGames: Int(2)}
	if !reflect.DeepEqual(match, want) {
		t.Errorf("Match.GetMatchListRecentByAccountID returned %+v, want %+v", match, want)
	}
}

func TestMatchService_GetTimelineByMatchID(t *testing.T) {
	setup()
	defer teardown()

	mux.HandleFunc("/"+client.MatchURL+"/timelines/by-match/1234", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		fmt.Fprint(w, `{"frameInterval":2}`)
	})

	match, _, err := client.Match.GetTimelineByMatchID(context.Background(), 1234)
	if err != nil {
		t.Errorf("Match.GetTimelineByMatchID returned error: %v", err)
	}

	want := &MatchTimeline{FrameInterval: Int(2)}
	if !reflect.DeepEqual(match, want) {
		t.Errorf("Match.GetTimelineByMatchID returned %+v, want %+v", match, want)
	}
}
