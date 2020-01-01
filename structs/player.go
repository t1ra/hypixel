package structs

/*
 * Copyright 2020 t1ra

 * Permission to use, copy, modify, and/or distribute this software for any purpose with or without
 * fee is hereby granted, provided that the above copyright notice and this permission notice appear
 * in all copies.

 * THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES WITH REGARD TO THIS
 * SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE
 * AUTHOR BE LIABLE FOR ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
 * WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN ACTION OF CONTRACT,
 * NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF OR IN CONNECTION WITH THE USE OR PERFORMANCE
 * OF THIS SOFTWARE.
 */

/*
 *                           Player
 */

// Player contains information about a player.
type Player struct {
	ID                  string               `json:"_id"`
	UUID                string               `json:"uuid"`
	DisplayName         string               `json:"displayname"`
	PlayerName          string               `json:"playername"`
	FirstLogin          int64                `json:"firstLogin"`
	LastLogin           int64                `json:"lastLogin"`
	Aliases             []string             `json:"knownAliases"`
	Achievements        map[string]int       `json:"achievements"`
	AchievementsOneTime map[int]string       `json:"activementsOneTime"`
	Version             string               `json:"mcVersionRp"`
	VanityTokens        int                  `json:"vanityTokens"`
	Karma               int                  `json:"karma"`
	Stats               PlayerStats          `json:"stats"`
	Settings            PlayerSettings       `json:"settings"`
	PetConsumables      PlayerPetConsumables `json:"petConsumables"`
	HousingMeta         PlayerHousingMeta    `json:"housingMeta"`
	VanityMeta          PlayerVanityMeta     `json:"vanityMeta"`
	Quests              PlayerQuests         `json:"quests"`
}

// PlayerSettings stores lobby settings for a player.
type PlayerSettings struct {
	SpecNightVision  bool `json:"spec_night_vision"`
	SpecSpeed        int  `json:"spec_speed"`
	playerVisibility bool `json:"playerVisibility"`
}

// PlayerPetConsumables stores pet consumables for a player.
type PlayerPetConsumables struct {
	Carrot       int `json:"CARROT_ITEM"`
	Feather      int `json:"FEATHER"`
	Hay          int `json:"HAY_BLOCK"`
	RottenFlesh  int `json:"ROTTEN_FLESH"`
	Slimeball    int `json:"SLIME_BALL"`
	Cake         int `json:"CAKE"`
	Steak        int `json:"COOKED_BEEF"`
	Water        int `json:"WATER_BUCKET"`
	WoodenSword  int `json:"WOOD_SWORD"`
	Stick        int `json:"STICK"`
	MilkBucket   int `json:"MILK_BUCKET"`
	GoldenRecord int `json:"GOLD_RECORD"`
	Leash        int `json:"LEASH"`
	Lava         int `json:"LAVA_BUCKET"`
	MagmaCream   int `json:"MAGMA_CREAM"`
	Cookie       int `json:"COOKIE"`
	BakedPotato  int `json:"BAKED_POTATO"`
	Bread        int `json:"BREAD"`
	Pork         int `json:"PORK"`
	MushroomSoup int `json:"MUSHROOM_SOUP"`
	PumpkinPie   int `json:"PUMPKIN_PIE"`
	Bone         int `json:"BONE"`
	Rose         int `json:"RED_ROSE"`
	Apple        int `json:"APPLE"`
	Melon        int `json:"MELON"`
	RawFish      int `json:"RAW_FISH"`
	Wheat        int `json:"WHEAT"`
}

// PlayerHousingMeta stores housing information for a player.
type PlayerHousingMeta struct {
	AllowedBlocks []string `json:"allowedBlocks"`
	Packages      []string `json:"packages"`
	TutorialStage string   `json:"tutorialStep"`
}

// PlayerVanityMeta stores vanity information for a player.
type PlayerVanityMeta struct {
	Packages []string `json:"packages"`
}

// PlayerQuests stores player quest information.
type PlayerQuests struct {
	SkywarsSoloKills struct {
		Completions []struct {
			Time int `json:"time"`
		} `json:"completions"`
	} `json:"skywars_solo_kills"`
	SkywarsSoloWin struct {
		Completions []struct {
			Time int `json:"time"`
		} `json:"completions"`
	} `json:"skywars_solo_win"`
	SkywarsTeamWin struct {
		Completions []struct {
			Time int `json:"time"`
		} `json:"completions"`
	} `json:"skywars_team_win"`
	SkywarsTeamKills struct {
		Completions []struct {
			Time int `json:"time"`
		} `json:"completions"`
	} `json:"skywars_team_kills"`
	SkywarsWeeklyKills struct{} `json:"skywars_weekly_kills"`
	QuakeDailyPlay     struct {
		Active struct {
			Objectives struct {
				QuakeDailyPlay int `json:"quake_daily_play"`
			} `json:"objectives"`
			Started int64 `json:"started"`
		} `json:"active"`
	} `json:"quake_daily_play"`
	QuakeDailyKill struct {
		Active struct {
			Objectives struct {
				QuakeDailyKill int `json:"quake_daily_kill"`
			} `json:"objectives"`
			Started int64 `json:"started"`
		} `json:"active"`
	} `json:"quake_daily_kill"`
	QuakeDailyWin struct {
		Active struct {
			Objectives struct {
			} `json:"objectives"`
			Started int64 `json:"started"`
		} `json:"active"`
	} `json:"quake_daily_win"`
	QuakeWeeklyPlay struct {
		Active struct {
			Objectives struct {
				QuakeWeeklyPlay int `json:"quake_weekly_play"`
			} `json:"objectives"`
			Started int64 `json:"started"`
		} `json:"active"`
	} `json:"quake_weekly_play"`
	SkywarsDailyTokens struct {
		Completions []struct {
			Time int64 `json:"time"`
		} `json:"completions"`
	} `json:"skywars_daily_tokens"`
	PrototypePitDailyKills struct {
		Active struct {
			Objectives struct {
				Kill int `json:"kill"`
			} `json:"objectives"`
			Started int64 `json:"started"`
		} `json:"active"`
	} `json:"prototype_pit_daily_kills"`
	FriendRequests []string `json:"friendRequestsUuid"`
	Outfit         struct {
		Helmet     string `json:"HELMET"`
		Chestplate string `json:"CHESTPLATE"`
		Leggings   string `json:"LEGGINGS"`
		Boots      string `json:"BOOTS"`
	} `json:"outfit"`
	ParkourCompletions struct {
		Skywars []struct {
			TimeStart int64 `json:"timeStart"`
			TimeTaken int   `json:"timeTook"`
		} `json:"Skywars"`
		SkywarsAug2017 []struct {
			TimeStart int64 `json:"timeStart"`
			TimeTaken int   `json:"timeTook"`
		} `json:"SkywarsAug2017"`
		Bedwars []struct {
			TimeStart int64 `json:"timeStart"`
			TimeTaken int   `json:"timeTook"`
		} `json:"Bedwars"`
		Duels []struct {
			TimeStart int64 `json:"timeStart"`
			TimeTaken int   `json:"timeTook"`
		} `json:"Duels"`
	} `json:"parkourCompletions"`
	Voting struct {
		Total           int   `json:"total"`
		TotalMcsorg     int   `json:"total_mcsorg"`
		SecondaryMcsorg int   `json:"secondary_mcsorg"`
		LastMcsorg      int64 `json:"last_mcsorg"`
		VotesToday      int   `json:"votesToday"`
	} `json:"voting"`
	AchievementTracking []interface{} `json:"achievementTracking"`
	LastLogout          int64         `json:"lastLogout"`
	AchievementPoints   int           `json:"achievementPoints"`
	AchievementSync     struct {
		QuakeTiered int `json:"quake_tiered"`
	} `json:"achievementSync"`
	Challenges struct {
		AllTime struct {
			SurvivalGamesResistanceChallenge int `json:"SURVIVAL_GAMES__resistance_challenge"`
			QuakeCraftComboChallenge         int `json:"QUAKECRAFT__combo_challenge"`
			SkywarsfeedingTheVoidChallenge   int `json:"SKYWARS__feeding_the_void_challenge"`
			ArcadeZombiesChallenge           int `json:"ARCADE__zombies_challenge"`
			BedwarsSOffensive                int `json:"BEDWARS__offensive"`
			BedwarsSupport                   int `json:"BEDWARS__support"`
			BuildBattleTop3Challenge         int `json:"BUILD_BATTLE__top_3_challenge"`
			DUelsTeamsChallenge              int `json:"DUELS__teams_challenge"`
			BedwarsDefensive                 int `json:"BEDWARS__defensive"`
			SkywarsEndermanChallenge         int `json:"SKYWARS__enderman_challenge"`
			DuelsFeedTheVoidChallenge        int `json:"DUELS__feed_the_void_challenge"`
		} `json:"all_time"`
		DayF struct {
			BedwarsSupport    int `json:"BEDWARS__support"`
			BedwarsSOffensive int `json:"BEDWARS__offensive"`
		} `json:"day_f"`
		DayG struct {
			BedwarsSOffensive int `json:"BEDWARS__offensive"`
			BedwarsSupport    int `json:"BEDWARS__support"`
		} `json:"day_g"`
		DayH struct {
			BedwarsSOffensive int `json:"BEDWARS__offensive"`
			BedwarsSupport    int `json:"BEDWARS__support"`
		} `json:"day_h"`
	} `json:"challenges"`
	AchievementRewardsNew struct {
		ForPoints600 int64 `json:"for_points_600"`
		ForPoints300 int64 `json:"for_points_300"`
		ForPoints200 int64 `json:"for_points_200"`
		ForPoints400 int64 `json:"for_points_400"`
	} `json:"achievementRewardsNew"`
	LevelingReward0        bool `json:"levelingReward_0"`
	LevelingReward6        bool `json:"levelingReward_6"`
	LevelingReward1        bool `json:"levelingReward_1"`
	LevelingReward2        bool `json:"levelingReward_2"`
	LevelingReward3        bool `json:"levelingReward_3"`
	LevelingReward5        bool `json:"levelingReward_5"`
	LevelingReward4        bool `json:"levelingReward_4"`
	ParkourCheckpointBests struct {
		SkywarsAug2017 struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
		} `json:"SkywarsAug2017"`
		Bedwars struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
		} `json:"Bedwars"`
		Duels struct {
			Num0 int `json:"0"`
			Num1 int `json:"1"`
			Num2 int `json:"2"`
			Num3 int `json:"3"`
		} `json:"Duels"`
		Prototype struct {
			Num0 int `json:"0"`
		} `json:"Prototype"`
	} `json:"parkourCheckpointBests"`
	QuickjoinTimestamp      int64  `json:"quickjoin_timestamp"`
	QuickjoinUses           int    `json:"quickjoin_uses"`
	LevelingReward7         bool   `json:"levelingReward_7"`
	Channel                 string `json:"channel"`
	LastAdsenseGenerateTime int64  `json:"lastAdsenseGenerateTime"`
	LastClaimedReward       int64  `json:"lastClaimedReward"`
	RewardHighScore         int    `json:"rewardHighScore"`
	RewardScore             int    `json:"rewardScore"`
	RewardStreak            int    `json:"rewardStreak"`
	TotalDailyRewards       int    `json:"totalDailyRewards"`
	TotalRewards            int    `json:"totalRewards"`
	Monthlycrates           struct {
		Eight2019 struct {
			REGULAR bool `json:"REGULAR"`
		} `json:"8-2019"`
	} `json:"monthlycrates"`
	Xmas2019MAINLOBBY13 bool   `json:"xmas2019_MAIN_LOBBY_13"`
	Xmas2019MAINLOBBY31 bool   `json:"xmas2019_MAIN_LOBBY_31"`
	Xmas2019MAINLOBBY30 bool   `json:"xmas2019_MAIN_LOBBY_30"`
	MostRecentGameType  string `json:"mostRecentGameType"`
}

// PlayerStats holds a collection of structs containing information about a player.
type PlayerStats struct {
	HungerGames   HungerGamesStats   `json:"HungerGames"`
	Paintball     PaintballStats     `json:"Paintball"`
	TNTGames      TNTGamesStats      `json:"TNTGames"`
	MegaWalls     MegaWallsStats     `json:"Walls3"`
	Skywars       SkywarsStats       `json:"Skywars"`
	Battleground  BattlegroundStats  `json:"Battleground"`
	TrueCombat    TrueCombatStats    `json:"TrueCombat"`
	UHC           UHCStats           `json:"UHC"`
	Walls         WallsStats         `json:"Walls"`
	Quake         QuakeStats         `json:"Quake"`
	Arcade        ArcadeStats        `json:"Arcade"`
	Arena         ArenaStats         `json:"Arena"`
	SuperSmash    SuperSmashStats    `json:"SuperSmashStats"`
	GingerBread   GingerBreadStats   `json:"GingerBread"`
	VampireZ      VampireZStats      `json:"VampireZ"`
	Bedwars       BedwarsStats       `json:"Bedwars"`
	Legacy        LegacyStats        `json:"Legacy"`
	Duels         DuelsStats         `json:"Duels"`
	Skyblock      SkyblockStats      `json:"Skyblock"`
	BuildBattle   BuildBattleStats   `json:"BuildBattle"`
	MurderMystery MurderMysteryStats `json:"MurderMystery"`
	Pit           PitStats           `json:"Pit"`
	MCGO          MCGOStats          `json:"MCGO"`
}

// HungerGamesStats is hunger games statistics about a player.
type HungerGamesStats struct {
	// General
	TimePlayed       int      `json:"time_played"`
	Kills            int      `json:"kills"`
	Deaths           int      `json:"deaths"`
	Coins            int      `json:"coins"`
	Wins             int      `json:"wins"`
	Archer           int      `json:"archer"`
	Armorer          int      `json:"armorer"`
	Knight           int      `json:"knight"`
	Scout            int      `json:"scout"`
	WeeklyKillsA     int      `json:"weekly_kills_a"`
	WeeklyKillsB     int      `json:"weekly_kills_b"`
	MonthlyKillsA    int      `json:"monthly_kills_a"`
	MonthlyKillsB    int      `json:"monthly_kills_b"`
	Autoarmor        bool     `json:"autoarmor"`
	Packages         []string `json:"packages"`
	PotionsDrunk     int      `json:"potions_drunk"`
	ArrowsFired      int      `json:"arrows_fired"`
	Damage           int      `json:"damage"`
	ChestsOpened     int      `json:"chests_opened"`
	DamageTaken      int      `json:"damage_taken"`
	GamesPlayed      int      `json:"games_played"`
	PotionsThrown    int      `json:"potions_thown"`
	MobsSpawned      int      `json:"mobs_spawned"`
	ArrowsHit        int      `json:"arrows_hit"`
	WinsTeamsNormal  int      `json:"wins_teams_normal"`
	WinsBackup       int      `json:"wins_backup"`
	WinsSoloNormal   int      `json:"wins_solo_normal"`
	KillsSoloNormal  int      `json:"kills_solo_normal"`
	KillsTeamsNormal int      `json:"kills_teams_normal"`
	// Random
	TimePlayedRandom   int `json:"time_played_random"`
	GamesPlayedRandom  int `json:"games_played_random"`
	DamageRandom       int `json:"damage_random"`
	DamageTakenRandom  int `json:"damage_taken_random"`
	PotionsDrunkRandom int `json:"potions_drunk_random"`
	ArrowsFiredRandom  int `json:"arrows_fired_random"`
	ChestsOpenedRandom int `json:"chests_opened_random"`
	// Scout
	TimePlayedScout    int `json:"time_played_scout"`
	GamesPlayedScout   int `json:"games_played_scout"`
	DamageScout        int `json:"damage_scout"`
	DamageTakenScout   int `json:"damage_taken_scout"`
	PotionsDrunkScout  int `json:"potions_drunk_scout"`
	PotionsThrownScout int `json:"potions_thrown_scout"`
	ArrowsFiredScout   int `json:"arrows_fired_scout"`
	ChestsOpenedScout  int `json:"chests_opened_scout"`
	ExperienceScout    int `json:"exp_scout"`
	KillsScout         int `json:"kills_scout"`
	// Fisherman
	TimePlayedFisherman   int `json:"time_played_fisherman"`
	GamesPlayedFisherman  int `json:"games_played_fisherman"`
	ChestsOpenedFisherman int `json:"chests_opened_fisherman"`
	DamageFisherman       int `json:"damge_fisherman"`
	DamageTakenFisherman  int `json:"damage_taken_fisherman"`
	PotionsDrunkFisherman int `json:"potions_drunk_fisherman"`
	ArrowsFiredFisherman  int `json:"arrows_fired_fisherman"`
	ArrowsHitFisherman    int `json:"arrows_hit_fisherman"`
	// Knight
	TimePlayedKnight   int `json:"time_played_knight"`
	ChestsOpenedKnight int `json:"chests_opened_scout"`
	DamageKnight       int `json:"damage_knight"`
	MobsSpawnedKnight  int `json:"mobs_spawned_knight"`
	DamageTakenKnight  int `json:"damage_taken_knight"`
	PotionsDrunkKnight int `json:"potions_drunk_knight"`
	GamesPlayedKnight  int `json:"games_played_knight"`
	KillsKnight        int `json:"kills_knight"`
	ArrowsFiredKnight  int `json:"arrows_fired_knight"`
	ExperienceKnight   int `json:"experience_knight"`
	// Archer
	TimePlayedArcher   int `json:"time_played_archer"`
	KillsArcher        int `json:"kills_archer"`
	DamageArcher       int `json:"damage_archer"`
	DamageTakenArcher  int `json:"damage_taken_archer"`
	ArrowsFiredArcher  int `json:"arrows_fired_archer"`
	ChestsOpenedArcher int `json:"chests_opened_archer"`
	GamesPlayedArcher  int `json:"games_played_archer"`
	PotionsDrunkArcher int `json:"potions_drunk_archer"`
}

// PaintballStats is paintball statistics about a player.
type PaintballStats struct {
	Coins       int      `json:"coins"`
	Deaths      int      `json:"deaths"`
	Kills       int      `json:"kills"`
	ShotsFired  int      `json:"shots_fired"`
	Packages    []string `json:"packages"`
	Killstreaks int      `json:"killstreaks"`
}

// TNTGamesStats is tnt games statistics about a player.
type TNTGamesStats struct {
	Coins                       int      `json:"coins"`
	DeathsBowspleef             int      `json:"deaths_bowspleef"`
	DeathsCapture               int      `json:"deaths_capture"`
	KillsCapture                int      `json:"kills_capture"`
	TagsBowspleef               int      `json:"tags_bowspleef"`
	WinsCapture                 int      `json:"wins_capture"`
	TNTRunRecord                int      `json:"record_tntrun"`
	Wins                        int      `json:"wins"`
	Packages                    []string `json:"packages"`
	NewIcewizardRegen           int      `json:"new_icewizard_regen"`
	NewBloodwizardExplode       int      `json:"new_bloodwizard_explode"`
	NewTntagSpeedy              int      `json:"new_tntag_speedy"`
	NewFirewizardRegen          int      `json:"new_firewizard_regen"`
	NewKineticwizardRegen       int      `json:"new_kineticwizard_regen"`
	NewSpleefTripleshot         int      `json:"new_spleef_tripleshot"`
	NewIcewizardExplode         int      `json:"new_icewizard_explode"`
	NewFirewizardExplode        int      `json:"new_firewizard_explode"`
	NewKineticwizardExplode     int      `json:"new_kineticwizard_explode"`
	NewPvprunDoubleJumps        int      `json:"new_pvprun_double_jumps"`
	NewWitherwizardExplode      int      `json:"new_witherwizard_explode"`
	NewTntrunDoubleJumps        int      `json:"new_tntrun_double_jumps"`
	NewSpleefRepulsor           int      `json:"new_spleef_repulsor"`
	NewSpleefDoubleJumps        int      `json:"new_spleef_double_jumps"`
	NewWitherwizardRegen        int      `json:"new_witherwizard_regen"`
	NewBloodwizardRegen         int      `json:"new_bloodwizard_regen"`
	RunPotionsSplashedOnPlayers int      `json:"run_potions_splashed_on_players"`
	DeathsTntrun                int      `json:"deaths_tntrun"`
	Winstreak                   int      `json:"winstreak"`
	Flags                       struct {
		ShowTipHolograms        bool `json:"show_tip_holograms"`
		ShowTntrunActionbarInfo bool `json:"show_tntrun_actionbar_info"`
	} `json:"flags"`
}

// MegaWallsStats is mega walls tatistics about a player.
type MegaWallsStats struct {
	ChosenClass        string   `json:"chosen_class"`
	Coins              int      `json:"coins"`
	Deaths             int      `json:"deaths"`
	DeathsHerobrine    int      `json:"deaths_Herobrine"`
	DeathsSkeleton     int      `json:"deaths_Skeleton"`
	DeathsZombie       int      `json:"deaths_Zombie"`
	FinalDeaths        int      `json:"finalDeaths"`
	FinalKills         int      `json:"finalKills"`
	FinalKillsSkeleton int      `json:"finalKills_Skeleton"`
	FinalKillsZombie   int      `json:"finalKills_Zombie"`
	HerobrineG         int      `json:"herobrine_g"`
	Kills              int      `json:"kills"`
	KillsHerobrine     int      `json:"kills_Herobrine"`
	KillsSkeleton      int      `json:"kills_Skeleton"`
	KillsZombie        int      `json:"kills_Zombie"`
	Losses             int      `json:"losses"`
	LossesHerobrine    int      `json:"losses_Herobrine"`
	LossesSkeleton     int      `json:"losses_Skeleton"`
	LossesZombie       int      `json:"losses_Zombie"`
	Wins               int      `json:"wins"`
	WinsSkeleton       int      `json:"wins_Skeleton"`
	Packages           []string `json:"packages"`
	Classes            struct {
		Skeleton struct {
			SkillLevelDChecked5 bool `json:"skill_level_dChecked5"`
			SkillLevelD         int  `json:"skill_level_d"`
			Unlocked            bool `json:"unlocked"`
			Checked4            bool `json:"checked4"`
		} `json:"skeleton"`
		Herobrine struct {
			SkillLevelD         int  `json:"skill_level_d"`
			SkillLevelDChecked5 bool `json:"skill_level_dChecked5"`
			Checked4            bool `json:"checked4"`
			Unlocked            bool `json:"unlocked"`
		} `json:"herobrine"`
		Enderman struct {
			SkillLevelD         int  `json:"skill_level_d"`
			SkillLevelDChecked5 bool `json:"skill_level_dChecked5"`
			Unlocked            bool `json:"unlocked"`
			Checked4            bool `json:"checked4"`
		} `json:"enderman"`
		Zombie struct {
			SkillLevelDChecked5 bool `json:"skill_level_dChecked5"`
			SkillLevelD         int  `json:"skill_level_d"`
			Unlocked            bool `json:"unlocked"`
			Checked4            bool `json:"checked4"`
		} `json:"zombie"`
	} `json:"classes"`
	SkeletonFinalKills                    int `json:"skeleton_final_kills"`
	SkeletonWins                          int `json:"skeleton_wins"`
	HerobrineDeaths                       int `json:"herobrine_deaths"`
	SkeletonKills                         int `json:"skeleton_kills"`
	ZombieDeaths                          int `json:"zombie_deaths"`
	ZombieTotalFinalKillsStandard         int `json:"zombie_total_final_kills_standard"`
	SkeletonLosses                        int `json:"skeleton_losses"`
	ZombieLosses                          int `json:"zombie_losses"`
	HerobrineLosses                       int `json:"herobrine_losses"`
	SkeletonTotalFinalKills               int `json:"skeleton_total_final_kills"`
	SkeletonWinsStandard                  int `json:"skeleton_wins_standard"`
	ZombieKills                           int `json:"zombie_kills"`
	ZombieTotalFinalKills                 int `json:"zombie_total_final_kills"`
	FinalKillsStandard                    int `json:"final_kills_standard"`
	TotalFinalKills                       int `json:"total_final_kills"`
	HerobrineKills                        int `json:"herobrine_kills"`
	ZombieFinalKills                      int `json:"zombie_final_kills"`
	SkeletonFinalKillsStandard            int `json:"skeleton_final_kills_standard"`
	SkeletonDeaths                        int `json:"skeleton_deaths"`
	SkeletonTotalFinalKillsStandard       int `json:"skeleton_total_final_kills_standard"`
	TotalFinalKillsStandard               int `json:"total_final_kills_standard"`
	ZombieFinalKillsStandard              int `json:"zombie_final_kills_standard"`
	ArrowsFired                           int `json:"arrows_fired"`
	ArrowsFiredStandard                   int `json:"arrows_fired_standard"`
	ArrowsHit                             int `json:"arrows_hit"`
	ArrowsHitStandard                     int `json:"arrows_hit_standard"`
	Assists                               int `json:"assists"`
	AssistsStandard                       int `json:"assists_standard"`
	BlocksBroken                          int `json:"blocks_broken"`
	BlocksBrokenStandard                  int `json:"blocks_broken_standard"`
	BlocksPlaced                          int `json:"blocks_placed"`
	BlocksPlacedPreparation               int `json:"blocks_placed_preparation"`
	BlocksPlacedPreparationStandard       int `json:"blocks_placed_preparation_standard"`
	BlocksPlacedStandard                  int `json:"blocks_placed_standard"`
	DeathsStandard                        int `json:"deaths_standard"`
	FoodEaten                             int `json:"food_eaten"`
	FoodEatenStandard                     int `json:"food_eaten_standard"`
	GamesPlayed                           int `json:"games_played"`
	GamesPlayedStandard                   int `json:"games_played_standard"`
	GoldenApplesEaten                     int `json:"golden_apples_eaten"`
	GoldenApplesEatenStandard             int `json:"golden_apples_eaten_standard"`
	HunterArrowsFired                     int `json:"hunter_arrows_fired"`
	HunterArrowsFiredStandard             int `json:"hunter_arrows_fired_standard"`
	HunterArrowsHit                       int `json:"hunter_arrows_hit"`
	HunterArrowsHitStandard               int `json:"hunter_arrows_hit_standard"`
	HunterAssists                         int `json:"hunter_assists"`
	HunterAssistsStandard                 int `json:"hunter_assists_standard"`
	HunterBlocksBroken                    int `json:"hunter_blocks_broken"`
	HunterBlocksBrokenStandard            int `json:"hunter_blocks_broken_standard"`
	HunterBlocksPlaced                    int `json:"hunter_blocks_placed"`
	HunterBlocksPlacedPreparation         int `json:"hunter_blocks_placed_preparation"`
	HunterBlocksPlacedPreparationStandard int `json:"hunter_blocks_placed_preparation_standard"`
	HunterBlocksPlacedStandard            int `json:"hunter_blocks_placed_standard"`
	HunterDeaths                          int `json:"hunter_deaths"`
	HunterDeathsStandard                  int `json:"hunter_deaths_standard"`
	HunterFoodEaten                       int `json:"hunter_food_eaten"`
	HunterFoodEatenStandard               int `json:"hunter_food_eaten_standard"`
	HunterGamesPlayed                     int `json:"hunter_games_played"`
	HunterGamesPlayedStandard             int `json:"hunter_games_played_standard"`
	HunterGoldenApplesEaten               int `json:"hunter_golden_apples_eaten"`
	HunterGoldenApplesEatenStandard       int `json:"hunter_golden_apples_eaten_standard"`
	HunterIronOreBroken                   int `json:"hunter_iron_ore_broken"`
	HunterIronOreBrokenStandard           int `json:"hunter_iron_ore_broken_standard"`
	HunterLosses                          int `json:"hunter_losses"`
	HunterLossesStandard                  int `json:"hunter_losses_standard"`
	HunterMetersFallen                    int `json:"hunter_meters_fallen"`
	HunterMetersFallenStandard            int `json:"hunter_meters_fallen_standard"`
	HunterMetersWalked                    int `json:"hunter_meters_walked"`
	HunterMetersWalkedSpeed               int `json:"hunter_meters_walked_speed"`
	HunterMetersWalkedSpeedStandard       int `json:"hunter_meters_walked_speed_standard"`
	HunterMetersWalkedStandard            int `json:"hunter_meters_walked_standard"`
	HunterTimePlayed                      int `json:"hunter_time_played"`
	HunterTimePlayedStandard              int `json:"hunter_time_played_standard"`
	HunterTotalDeaths                     int `json:"hunter_total_deaths"`
	HunterTotalDeathsStandard             int `json:"hunter_total_deaths_standard"`
	HunterTotalKills                      int `json:"hunter_total_kills"`
	HunterTotalKillsStandard              int `json:"hunter_total_kills_standard"`
	HunterTreasuresFound                  int `json:"hunter_treasures_found"`
	HunterTreasuresFoundStandard          int `json:"hunter_treasures_found_standard"`
	HunterWoodChopped                     int `json:"hunter_wood_chopped"`
	HunterWoodChoppedStandard             int `json:"hunter_wood_chopped_standard"`
	IronOreBroken                         int `json:"iron_ore_broken"`
	IronOreBrokenStandard                 int `json:"iron_ore_broken_standard"`
	LossesStandard                        int `json:"losses_standard"`
	MetersFallen                          int `json:"meters_fallen"`
	MetersFallenStandard                  int `json:"meters_fallen_standard"`
	MetersWalked                          int `json:"meters_walked"`
	MetersWalkedSpeed                     int `json:"meters_walked_speed"`
	MetersWalkedSpeedStandard             int `json:"meters_walked_speed_standard"`
	MetersWalkedStandard                  int `json:"meters_walked_standard"`
	TimePlayed                            int `json:"time_played"`
	TimePlayedStandard                    int `json:"time_played_standard"`
	TotalDeaths                           int `json:"total_deaths"`
	TotalDeathsStandard                   int `json:"total_deaths_standard"`
	TotalKills                            int `json:"total_kills"`
	TotalKillsStandard                    int `json:"total_kills_standard"`
	TreasuresFound                        int `json:"treasures_found"`
	TreasuresFoundStandard                int `json:"treasures_found_standard"`
	WoodChopped                           int `json:"wood_chopped"`
	WoodChoppedStandard                   int `json:"wood_chopped_standard"`
}

// SkywarsStats is skywars stats about a player.
type SkywarsStats struct {
	Packages                                      []string `json:"packages"`
	UsedSoulWell                                  bool     `json:"usedSoulWell"`
	SoulWell                                      int      `json:"soul_well"`
	MegaMiningExpertise                           int      `json:"mega_mining_expertise"`
	ActiveKitTEAM                                 string   `json:"activeKit_TEAM"`
	ActiveKitMEGA                                 string   `json:"activeKit_MEGA"`
	WinStreak                                     int      `json:"win_streak"`
	LossesKitMiningTeamDefault                    int      `json:"losses_kit_mining_team_default"`
	BlocksBroken                                  int      `json:"blocks_broken"`
	LossesTeamNormal                              int      `json:"losses_team_normal"`
	Losses                                        int      `json:"losses"`
	SurvivedPlayersKitMiningTeamDefault           int      `json:"survived_players_kit_mining_team_default"`
	LossesTeam                                    int      `json:"losses_team"`
	SurvivedPlayersTeam                           int      `json:"survived_players_team"`
	Quits                                         int      `json:"quits"`
	DeathsTeamNormal                              int      `json:"deaths_team_normal"`
	SurvivedPlayers                               int      `json:"survived_players"`
	DeathsTeam                                    int      `json:"deaths_team"`
	DeathsKitMiningTeamDefault                    int      `json:"deaths_kit_mining_team_default"`
	Deaths                                        int      `json:"deaths"`
	LossesMegaNormal                              int      `json:"losses_mega_normal"`
	Coins                                         int      `json:"coins"`
	DeathsMega                                    int      `json:"deaths_mega"`
	SurvivedPlayersKitMegaMegaDefault             int      `json:"survived_players_kit_mega_mega_default"`
	DeathsKitMegaMegaDefault                      int      `json:"deaths_kit_mega_mega_default"`
	DeathsMegaNormal                              int      `json:"deaths_mega_normal"`
	Assists                                       int      `json:"assists"`
	LossesKitMegaMegaDefault                      int      `json:"losses_kit_mega_mega_default"`
	SurvivedPlayersMega                           int      `json:"survived_players_mega"`
	LossesMega                                    int      `json:"losses_mega"`
	AssistsMega                                   int      `json:"assists_mega"`
	AssistsKitMegaMegaDefault                     int      `json:"assists_kit_mega_mega_default"`
	BlocksPlaced                                  int      `json:"blocks_placed"`
	Kills                                         int      `json:"kills"`
	WinsKitMegaMegaDefault                        int      `json:"wins_kit_mega_mega_default"`
	KillsKitMegaMegaDefault                       int      `json:"kills_kit_mega_mega_default"`
	KillsMega                                     int      `json:"kills_mega"`
	Wins                                          int      `json:"wins"`
	WinsMegaNormal                                int      `json:"wins_mega_normal"`
	SoulsGathered                                 int      `json:"souls_gathered"`
	WinsMega                                      int      `json:"wins_mega"`
	Souls                                         int      `json:"souls"`
	ArrowsShot                                    int      `json:"arrows_shot"`
	KillsMegaNormal                               int      `json:"kills_mega_normal"`
	EggThrown                                     int      `json:"egg_thrown"`
	ActiveKitSOLO                                 string   `json:"activeKit_SOLO"`
	SurvivedPlayersSolo                           int      `json:"survived_players_solo"`
	LossesSolo                                    int      `json:"losses_solo"`
	DeathsSoloNormal                              int      `json:"deaths_solo_normal"`
	SurvivedPlayersKitBasicSoloDefault            int      `json:"survived_players_kit_basic_solo_default"`
	LossesKitBasicSoloDefault                     int      `json:"losses_kit_basic_solo_default"`
	DeathsKitBasicSoloDefault                     int      `json:"deaths_kit_basic_solo_default"`
	LossesSoloNormal                              int      `json:"losses_solo_normal"`
	DeathsSolo                                    int      `json:"deaths_solo"`
	LossesSoloInsane                              int      `json:"losses_solo_insane"`
	DeathsSoloInsane                              int      `json:"deaths_solo_insane"`
	ItemsEnchanted                                int      `json:"items_enchanted"`
	ArrowsHit                                     int      `json:"arrows_hit"`
	DeathsTeamInsane                              int      `json:"deaths_team_insane"`
	LossesTeamInsane                              int      `json:"losses_team_insane"`
	KillsTeamInsane                               int      `json:"kills_team_insane"`
	KillsTeam                                     int      `json:"kills_team"`
	KillsKitMiningTeamDefault                     int      `json:"kills_kit_mining_team_default"`
	GamesTeam                                     int      `json:"games_team"`
	Games                                         int      `json:"games"`
	GamesKitMiningTeamDefault                     int      `json:"games_kit_mining_team_default"`
	WinsKitMiningTeamDefault                      int      `json:"wins_kit_mining_team_default"`
	WinsTeamInsane                                int      `json:"wins_team_insane"`
	WinsTeam                                      int      `json:"wins_team"`
	GamesSolo                                     int      `json:"games_solo"`
	GamesKitBasicSoloDefault                      int      `json:"games_kit_basic_solo_default"`
	AssistsKitMiningTeamDefault                   int      `json:"assists_kit_mining_team_default"`
	AssistsTeam                                   int      `json:"assists_team"`
	AssistsKitBasicSoloDefault                    int      `json:"assists_kit_basic_solo_default"`
	AssistsSolo                                   int      `json:"assists_solo"`
	RefillChestDestroy                            int      `json:"refill_chest_destroy"`
	KillsTeamNormal                               int      `json:"kills_team_normal"`
	SoloResistanceBoost                           int      `json:"solo_resistance_boost"`
	KillsSoloNormal                               int      `json:"kills_solo_normal"`
	KillsKitBasicSoloDefault                      int      `json:"kills_kit_basic_solo_default"`
	KillsSolo                                     int      `json:"kills_solo"`
	KillsSoloInsane                               int      `json:"kills_solo_insane"`
	TeamEnderMastery                              int      `json:"team_ender_mastery"`
	WinsTeamNormal                                int      `json:"wins_team_normal"`
	EnderpearlsThrown                             int      `json:"enderpearls_thrown"`
	PaidSouls                                     int      `json:"paid_souls"`
	SoulWellRares                                 int      `json:"soul_well_rares"`
	SurvivedPlayersKitAttackingTeamScout          int      `json:"survived_players_kit_attacking_team_scout"`
	LossesKitAttackingTeamScout                   int      `json:"losses_kit_attacking_team_scout"`
	DeathsKitAttackingTeamScout                   int      `json:"deaths_kit_attacking_team_scout"`
	KillsKitAttackingTeamScout                    int      `json:"kills_kit_attacking_team_scout"`
	AssistsKitAttackingTeamScout                  int      `json:"assists_kit_attacking_team_scout"`
	GamesKitAttackingTeamScout                    int      `json:"games_kit_attacking_team_scout"`
	LossesKitBasicSoloEcologist                   int      `json:"losses_kit_basic_solo_ecologist"`
	DeathsKitBasicSoloEcologist                   int      `json:"deaths_kit_basic_solo_ecologist"`
	SurvivedPlayersKitBasicSoloEcologist          int      `json:"survived_players_kit_basic_solo_ecologist"`
	GamesKitBasicSoloEcologist                    int      `json:"games_kit_basic_solo_ecologist"`
	WinsKitAttackingTeamScout                     int      `json:"wins_kit_attacking_team_scout"`
	KillsKitBasicSoloEcologist                    int      `json:"kills_kit_basic_solo_ecologist"`
	ActiveCage                                    string   `json:"activeCage"`
	AssistsKitBasicSoloEcologist                  int      `json:"assists_kit_basic_solo_ecologist"`
	WinsKitBasicSoloEcologist                     int      `json:"wins_kit_basic_solo_ecologist"`
	WinsSoloInsane                                int      `json:"wins_solo_insane"`
	WinsSolo                                      int      `json:"wins_solo"`
	GamesMega                                     int      `json:"games_mega"`
	GamesKitMegaMegaDefault                       int      `json:"games_kit_mega_mega_default"`
	TeamInstantSmelting                           int      `json:"team_instant_smelting"`
	MegaRusher                                    int      `json:"mega_rusher"`
	MegaEnderMastery                              int      `json:"mega_ender_mastery"`
	TeamResistanceBoost                           int      `json:"team_resistance_boost"`
	LossesKitAdvancedSoloEnchanter                int      `json:"losses_kit_advanced_solo_enchanter"`
	DeathsKitAdvancedSoloEnchanter                int      `json:"deaths_kit_advanced_solo_enchanter"`
	SurvivedPlayersKitAdvancedSoloEnchanter       int      `json:"survived_players_kit_advanced_solo_enchanter"`
	WinsSoloNormal                                int      `json:"wins_solo_normal"`
	VotesShire                                    int      `json:"votes_Shire"`
	VotesChronos                                  int      `json:"votes_Chronos"`
	DeathsKitBasicSoloArmorsmith                  int      `json:"deaths_kit_basic_solo_armorsmith"`
	LossesKitBasicSoloArmorsmith                  int      `json:"losses_kit_basic_solo_armorsmith"`
	SurvivedPlayersKitBasicSoloArmorsmith         int      `json:"survived_players_kit_basic_solo_armorsmith"`
	KillsKitBasicSoloArmorsmith                   int      `json:"kills_kit_basic_solo_armorsmith"`
	GamesKitBasicSoloArmorsmith                   int      `json:"games_kit_basic_solo_armorsmith"`
	WinsKitBasicSoloArmorsmith                    int      `json:"wins_kit_basic_solo_armorsmith"`
	VotesToadstool                                int      `json:"votes_Toadstool"`
	VotesTowers                                   int      `json:"votes_Towers"`
	DeathsKitSupportingTeamHealer                 int      `json:"deaths_kit_supporting_team_healer"`
	LossesKitSupportingTeamHealer                 int      `json:"losses_kit_supporting_team_healer"`
	SurvivedPlayersKitSupportingTeamHealer        int      `json:"survived_players_kit_supporting_team_healer"`
	KillsKitSupportingTeamHealer                  int      `json:"kills_kit_supporting_team_healer"`
	VotesSteampunk                                int      `json:"votes_Steampunk"`
	WinsKitAdvancedSoloEnchanter                  int      `json:"wins_kit_advanced_solo_enchanter"`
	GamesKitAdvancedSoloEnchanter                 int      `json:"games_kit_advanced_solo_enchanter"`
	KillsKitAdvancedSoloEnchanter                 int      `json:"kills_kit_advanced_solo_enchanter"`
	VotesTribal                                   int      `json:"votes_Tribal"`
	SoloEnderMastery                              int      `json:"solo_ender_mastery"`
	SoulWellLegendaries                           int      `json:"soul_well_legendaries"`
	SoloInstantSmelting                           int      `json:"solo_instant_smelting"`
	AssistsKitBasicSoloArmorsmith                 int      `json:"assists_kit_basic_solo_armorsmith"`
	XezbethLuck                                   int      `json:"xezbeth_luck"`
	SurvivedPlayersKitAttackingTeamKnight         int      `json:"survived_players_kit_attacking_team_knight"`
	DeathsKitAttackingTeamKnight                  int      `json:"deaths_kit_attacking_team_knight"`
	LossesKitAttackingTeamKnight                  int      `json:"losses_kit_attacking_team_knight"`
	AssistsKitAttackingTeamKnight                 int      `json:"assists_kit_attacking_team_knight"`
	KillsKitAttackingTeamKnight                   int      `json:"kills_kit_attacking_team_knight"`
	WinsKitAttackingTeamKnight                    int      `json:"wins_kit_attacking_team_knight"`
	GamesKitAttackingTeamKnight                   int      `json:"games_kit_attacking_team_knight"`
	SurvivedPlayersKitDefendingTeamGuardian       int      `json:"survived_players_kit_defending_team_guardian"`
	LossesKitDefendingTeamGuardian                int      `json:"losses_kit_defending_team_guardian"`
	DeathsKitDefendingTeamGuardian                int      `json:"deaths_kit_defending_team_guardian"`
	KillsKitDefendingTeamGuardian                 int      `json:"kills_kit_defending_team_guardian"`
	GamesKitDefendingTeamGuardian                 int      `json:"games_kit_defending_team_guardian"`
	VotesCongo                                    int      `json:"votes_Congo"`
	VotesDragonice                                int      `json:"votes_Dragonice"`
	VotesLongIsland                               int      `json:"votes_LongIsland"`
	VotesAtuin                                    int      `json:"votes_Atuin"`
	TeamMiningExpertise                           int      `json:"team_mining_expertise"`
	VotesOverfall                                 int      `json:"votes_Overfall"`
	SoloMiningExpertise                           int      `json:"solo_mining_expertise"`
	DeathsKitSupportingTeamEcologist              int      `json:"deaths_kit_supporting_team_ecologist"`
	KillsKitSupportingTeamEcologist               int      `json:"kills_kit_supporting_team_ecologist"`
	AssistsKitSupportingTeamEcologist             int      `json:"assists_kit_supporting_team_ecologist"`
	SurvivedPlayersKitSupportingTeamEcologist     int      `json:"survived_players_kit_supporting_team_ecologist"`
	LossesKitSupportingTeamEcologist              int      `json:"losses_kit_supporting_team_ecologist"`
	AssistsKitAdvancedSoloEnchanter               int      `json:"assists_kit_advanced_solo_enchanter"`
	SoloSpeedBoost                                int      `json:"solo_speed_boost"`
	GamesKitSupportingTeamHealer                  int      `json:"games_kit_supporting_team_healer"`
	AssistsKitSupportingTeamHealer                int      `json:"assists_kit_supporting_team_healer"`
	KillsMonthlyB                                 int      `json:"kills_monthly_b"`
	KillsWeeklyB                                  int      `json:"kills_weekly_b"`
	KillsWeeklyA                                  int      `json:"kills_weekly_a"`
	KillsMonthlyA                                 int      `json:"kills_monthly_a"`
	Killstreak                                    int      `json:"killstreak"`
	HighestKillstreak                             int      `json:"highestKillstreak"`
	SoloNourishment                               int      `json:"solo_nourishment"`
	KillstreakSolo                                int      `json:"killstreak_solo"`
	WinstreakSolo                                 int      `json:"winstreak_solo"`
	Winstreak                                     int      `json:"winstreak"`
	WinstreakKitBasicSoloEcologist                int      `json:"winstreak_kit_basic_solo_ecologist"`
	KillstreakKitBasicSoloEcologist               int      `json:"killstreak_kit_basic_solo_ecologist"`
	HighestWinstreak                              int      `json:"highestWinstreak"`
	KillstreakKitBasicSoloArmorsmith              int      `json:"killstreak_kit_basic_solo_armorsmith"`
	WinstreakKitBasicSoloArmorsmith               int      `json:"winstreak_kit_basic_solo_armorsmith"`
	MegaNourishment                               int      `json:"mega_nourishment"`
	SurvivedPlayersKitDefendingTeamBaseballPlayer int      `json:"survived_players_kit_defending_team_baseball-player"`
	LossesKitDefendingTeamBaseballPlayer          int      `json:"losses_kit_defending_team_baseball-player"`
	DeathsKitDefendingTeamBaseballPlayer          int      `json:"deaths_kit_defending_team_baseball-player"`
	GamesKitDefendingTeamBaseballPlayer           int      `json:"games_kit_defending_team_baseball-player"`
	SurvivedPlayersKitBasicSoloRookie             int      `json:"survived_players_kit_basic_solo_rookie"`
	GamesKitBasicSoloRookie                       int      `json:"games_kit_basic_solo_rookie"`
	KillsKitBasicSoloRookie                       int      `json:"kills_kit_basic_solo_rookie"`
	LossesKitBasicSoloRookie                      int      `json:"losses_kit_basic_solo_rookie"`
	DeathsKitBasicSoloRookie                      int      `json:"deaths_kit_basic_solo_rookie"`
	KillsKitDefendingTeamBaseballPlayer           int      `json:"kills_kit_defending_team_baseball-player"`
	WinstreakKitAttackingTeamScout                int      `json:"winstreak_kit_attacking_team_scout"`
	WinstreakTeam                                 int      `json:"winstreak_team"`
	LossesKitDefendingTeamArmorer                 int      `json:"losses_kit_defending_team_armorer"`
	AssistsKitDefendingTeamArmorer                int      `json:"assists_kit_defending_team_armorer"`
	SurvivedPlayersKitDefendingTeamArmorer        int      `json:"survived_players_kit_defending_team_armorer"`
	DeathsKitDefendingTeamArmorer                 int      `json:"deaths_kit_defending_team_armorer"`
	MostKillsGameKitBasicSoloRookie               int      `json:"most_kills_game_kit_basic_solo_rookie"`
	VoidKills                                     int      `json:"void_kills"`
	VoidKillsKitBasicSoloRookie                   int      `json:"void_kills_kit_basic_solo_rookie"`
	ChestsOpened                                  int      `json:"chests_opened"`
	TimePlayedKitBasicSoloRookie                  int      `json:"time_played_kit_basic_solo_rookie"`
	MostKillsGameSolo                             int      `json:"most_kills_game_solo"`
	ChestsOpenedSolo                              int      `json:"chests_opened_solo"`
	TimePlayedSolo                                int      `json:"time_played_solo"`
	ChestsOpenedKitBasicSoloRookie                int      `json:"chests_opened_kit_basic_solo_rookie"`
	TimePlayed                                    int      `json:"time_played"`
	MostKillsGame                                 int      `json:"most_kills_game"`
	VoidKillsSolo                                 int      `json:"void_kills_solo"`
	ChestsOpenedKitBasicSoloArmorsmith            int      `json:"chests_opened_kit_basic_solo_armorsmith"`
	MostKillsGameKitBasicSoloArmorsmith           int      `json:"most_kills_game_kit_basic_solo_armorsmith"`
	VoidKillsKitBasicSoloArmorsmith               int      `json:"void_kills_kit_basic_solo_armorsmith"`
	TimePlayedKitBasicSoloArmorsmith              int      `json:"time_played_kit_basic_solo_armorsmith"`
	FreeLootChestNpc                              int64    `json:"freeLootChestNpc"`
	SkywarsChests                                 int      `json:"skywars_chests"`
	SkyWarsOpenedChests                           int      `json:"SkyWars_openedChests"`
	SkyWarsOpenedRares                            int      `json:"SkyWars_openedRares"`
	SkywarsChestHistory                           []string `json:"skywars_chest_history"`
	SkyWarsOpenedCommons                          int      `json:"SkyWars_openedCommons"`
	ActiveBalloon                                 string   `json:"active_balloon"`
	ActiveKitTEAMS                                string   `json:"activeKit_TEAMS"`
	GamesPlayedSkywars                            int      `json:"games_played_skywars"`
	LastMode                                      string   `json:"lastMode"`
	TimePlayedKitDefendingTeamBaseballPlayer      int      `json:"time_played_kit_defending_team_baseball-player"`
	ChestsOpenedTeam                              int      `json:"chests_opened_team"`
	ChestsOpenedKitDefendingTeamBaseballPlayer    int      `json:"chests_opened_kit_defending_team_baseball-player"`
	TimePlayedTeam                                int      `json:"time_played_team"`
	AssistsKitDefendingTeamBaseballPlayer         int      `json:"assists_kit_defending_team_baseball-player"`
	ActiveKitTEAMSRandom                          bool     `json:"activeKit_TEAMS_random"`
	FastestWin                                    int      `json:"fastest_win"`
	FastestWinTeam                                int      `json:"fastest_win_team"`
	FastestWinKitAttackingTeamScout               int      `json:"fastest_win_kit_attacking_team_scout"`
	ChestsOpenedKitAttackingTeamScout             int      `json:"chests_opened_kit_attacking_team_scout"`
	TimePlayedKitAttackingTeamScout               int      `json:"time_played_kit_attacking_team_scout"`
	ActiveKitSOLORandom                           bool     `json:"activeKit_SOLO_random"`
	LongestBowKillSolo                            int      `json:"longest_bow_kill_solo"`
	LongestBowShotSolo                            int      `json:"longest_bow_shot_solo"`
	LongestBowShotKitBasicSoloArmorsmith          int      `json:"longest_bow_shot_kit_basic_solo_armorsmith"`
	LongestBowShot                                int      `json:"longest_bow_shot"`
	FastestWinSolo                                int      `json:"fastest_win_solo"`
	LongestBowKillKitBasicSoloArmorsmith          int      `json:"longest_bow_kill_kit_basic_solo_armorsmith"`
	LongestBowKill                                int      `json:"longest_bow_kill"`
	FastestWinKitBasicSoloArmorsmith              int      `json:"fastest_win_kit_basic_solo_armorsmith"`
	MeleeKillsSolo                                int      `json:"melee_kills_solo"`
	ArrowsHitKitBasicSoloArmorsmith               int      `json:"arrows_hit_kit_basic_solo_armorsmith"`
	ArrowsShotKitBasicSoloArmorsmith              int      `json:"arrows_shot_kit_basic_solo_armorsmith"`
	ArrowsHitSolo                                 int      `json:"arrows_hit_solo"`
	MeleeKills                                    int      `json:"melee_kills"`
	MeleeKillsKitBasicSoloArmorsmith              int      `json:"melee_kills_kit_basic_solo_armorsmith"`
	ArrowsShotSolo                                int      `json:"arrows_shot_solo"`
	SurvivedPlayersKitSupportingTeamArmorsmith    int      `json:"survived_players_kit_supporting_team_armorsmith"`
	ArrowsShotTeam                                int      `json:"arrows_shot_team"`
	ArrowsShotKitSupportingTeamArmorsmith         int      `json:"arrows_shot_kit_supporting_team_armorsmith"`
	LossesKitSupportingTeamArmorsmith             int      `json:"losses_kit_supporting_team_armorsmith"`
	ChestsOpenedKitSupportingTeamArmorsmith       int      `json:"chests_opened_kit_supporting_team_armorsmith"`
	DeathsKitSupportingTeamArmorsmith             int      `json:"deaths_kit_supporting_team_armorsmith"`
	TimePlayedKitSupportingTeamArmorsmith         int      `json:"time_played_kit_supporting_team_armorsmith"`
	MostKillsGameTeam                             int      `json:"most_kills_game_team"`
	LongestBowShotTeam                            int      `json:"longest_bow_shot_team"`
	MostKillsGameKitSupportingTeamArmorsmith      int      `json:"most_kills_game_kit_supporting_team_armorsmith"`
	LongestBowKillKitSupportingTeamArmorsmith     int      `json:"longest_bow_kill_kit_supporting_team_armorsmith"`
	LongestBowKillTeam                            int      `json:"longest_bow_kill_team"`
	LongestBowShotKitSupportingTeamArmorsmith     int      `json:"longest_bow_shot_kit_supporting_team_armorsmith"`
	MeleeKillsTeam                                int      `json:"melee_kills_team"`
	KillsKitSupportingTeamArmorsmith              int      `json:"kills_kit_supporting_team_armorsmith"`
	ArrowsHitTeam                                 int      `json:"arrows_hit_team"`
	ArrowsHitKitSupportingTeamArmorsmith          int      `json:"arrows_hit_kit_supporting_team_armorsmith"`
	MeleeKillsKitSupportingTeamArmorsmith         int      `json:"melee_kills_kit_supporting_team_armorsmith"`
	GamesKitSupportingTeamArmorsmith              int      `json:"games_kit_supporting_team_armorsmith"`
	LongestBowShotKitAttackingTeamScout           int      `json:"longest_bow_shot_kit_attacking_team_scout"`
	MostKillsGameKitAttackingTeamScout            int      `json:"most_kills_game_kit_attacking_team_scout"`
	ArrowsShotKitAttackingTeamScout               int      `json:"arrows_shot_kit_attacking_team_scout"`
	VoidKillsKitAttackingTeamScout                int      `json:"void_kills_kit_attacking_team_scout"`
	ArrowsHitKitAttackingTeamScout                int      `json:"arrows_hit_kit_attacking_team_scout"`
	VoidKillsTeam                                 int      `json:"void_kills_team"`
	ChestsOpenedKitSupportingTeamHealer           int      `json:"chests_opened_kit_supporting_team_healer"`
	TimePlayedKitSupportingTeamHealer             int      `json:"time_played_kit_supporting_team_healer"`
	MostKillsGameKitSupportingTeamHealer          int      `json:"most_kills_game_kit_supporting_team_healer"`
	LongestBowKillKitSupportingTeamHealer         int      `json:"longest_bow_kill_kit_supporting_team_healer"`
	MeleeKillsKitSupportingTeamHealer             int      `json:"melee_kills_kit_supporting_team_healer"`
	ArrowsShotKitSupportingTeamHealer             int      `json:"arrows_shot_kit_supporting_team_healer"`
	LongestBowShotKitSupportingTeamHealer         int      `json:"longest_bow_shot_kit_supporting_team_healer"`
	ArrowsHitKitSupportingTeamHealer              int      `json:"arrows_hit_kit_supporting_team_healer"`
	FastestWinKitSupportingTeamHealer             int      `json:"fastest_win_kit_supporting_team_healer"`
	VoidKillsKitSupportingTeamHealer              int      `json:"void_kills_kit_supporting_team_healer"`
	KillstreakKitSupportingTeamHealer             int      `json:"killstreak_kit_supporting_team_healer"`
	KillstreakTeam                                int      `json:"killstreak_team"`
	WinstreakKitSupportingTeamHealer              int      `json:"winstreak_kit_supporting_team_healer"`
	WinsKitSupportingTeamHealer                   int      `json:"wins_kit_supporting_team_healer"`
	LongestBowKillKitAttackingTeamScout           int      `json:"longest_bow_kill_kit_attacking_team_scout"`
	BowKillsTeam                                  int      `json:"bow_kills_team"`
	BowKills                                      int      `json:"bow_kills"`
	BowKillsKitAttackingTeamScout                 int      `json:"bow_kills_kit_attacking_team_scout"`
	LongestBowKillKitAttackingTeamKnight          int      `json:"longest_bow_kill_kit_attacking_team_knight"`
	LongestBowShotKitAttackingTeamKnight          int      `json:"longest_bow_shot_kit_attacking_team_knight"`
	MostKillsGameKitAttackingTeamKnight           int      `json:"most_kills_game_kit_attacking_team_knight"`
	TimePlayedKitAttackingTeamKnight              int      `json:"time_played_kit_attacking_team_knight"`
	ArrowsHitKitAttackingTeamKnight               int      `json:"arrows_hit_kit_attacking_team_knight"`
	MeleeKillsKitAttackingTeamKnight              int      `json:"melee_kills_kit_attacking_team_knight"`
	ArrowsShotKitAttackingTeamKnight              int      `json:"arrows_shot_kit_attacking_team_knight"`
	ChestsOpenedKitAttackingTeamKnight            int      `json:"chests_opened_kit_attacking_team_knight"`
	VoidKillsKitAttackingTeamKnight               int      `json:"void_kills_kit_attacking_team_knight"`
	TeamSpeedBoost                                int      `json:"team_speed_boost"`
	SkywarsExperience                             int      `json:"skywars_experience"`
	TeamNourishment                               int      `json:"team_nourishment"`
	FastestWinKitBasicSoloScout                   int      `json:"fastest_win_kit_basic_solo_scout"`
	MostKillsGameKitBasicSoloScout                int      `json:"most_kills_game_kit_basic_solo_scout"`
	LongestBowKillKitBasicSoloScout               int      `json:"longest_bow_kill_kit_basic_solo_scout"`
	KillsKitBasicSoloScout                        int      `json:"kills_kit_basic_solo_scout"`
	SurvivedPlayersKitBasicSoloScout              int      `json:"survived_players_kit_basic_solo_scout"`
	MeleeKillsKitBasicSoloScout                   int      `json:"melee_kills_kit_basic_solo_scout"`
	TimePlayedKitBasicSoloScout                   int      `json:"time_played_kit_basic_solo_scout"`
	WinsKitBasicSoloScout                         int      `json:"wins_kit_basic_solo_scout"`
	WinstreakKitBasicSoloScout                    int      `json:"winstreak_kit_basic_solo_scout"`
	KillstreakKitBasicSoloScout                   int      `json:"killstreak_kit_basic_solo_scout"`
	GamesKitBasicSoloScout                        int      `json:"games_kit_basic_solo_scout"`
	VoidKillsKitBasicSoloScout                    int      `json:"void_kills_kit_basic_solo_scout"`
	ChestsOpenedKitBasicSoloScout                 int      `json:"chests_opened_kit_basic_solo_scout"`
	LevelFormatted                                string   `json:"levelFormatted"`
	QuickjoinUsesAtuin                            int      `json:"quickjoin_uses_Atuin"`
	QuickjoinUsesTotal                            int      `json:"quickjoin_uses_total"`
	DeathsKitBasicSoloScout                       int      `json:"deaths_kit_basic_solo_scout"`
	LossesKitBasicSoloScout                       int      `json:"losses_kit_basic_solo_scout"`
	QuickjoinUsesRandom                           int      `json:"quickjoin_uses_random"`
	ChestsOpenedKitDefendingTeamArmorer           int      `json:"chests_opened_kit_defending_team_armorer"`
	GamesKitDefendingTeamArmorer                  int      `json:"games_kit_defending_team_armorer"`
	KillsKitDefendingTeamArmorer                  int      `json:"kills_kit_defending_team_armorer"`
	MostKillsGameKitDefendingTeamArmorer          int      `json:"most_kills_game_kit_defending_team_armorer"`
	TimePlayedKitDefendingTeamArmorer             int      `json:"time_played_kit_defending_team_armorer"`
	VoidKillsKitDefendingTeamArmorer              int      `json:"void_kills_kit_defending_team_armorer"`
	ArrowsShotKitDefendingTeamBaseballPlayer      int      `json:"arrows_shot_kit_defending_team_baseball-player"`
	ChestsOpenedKitBasicSoloBaseballPlayer        int      `json:"chests_opened_kit_basic_solo_baseball-player"`
	DeathsKitBasicSoloBaseballPlayer              int      `json:"deaths_kit_basic_solo_baseball-player"`
	KillsKitBasicSoloBaseballPlayer               int      `json:"kills_kit_basic_solo_baseball-player"`
	LossesKitBasicSoloBaseballPlayer              int      `json:"losses_kit_basic_solo_baseball-player"`
	MostKillsGameKitBasicSoloBaseballPlayer       int      `json:"most_kills_game_kit_basic_solo_baseball-player"`
	SurvivedPlayersKitBasicSoloBaseballPlayer     int      `json:"survived_players_kit_basic_solo_baseball-player"`
	TimePlayedKitBasicSoloBaseballPlayer          int      `json:"time_played_kit_basic_solo_baseball-player"`
	VoidKillsKitBasicSoloBaseballPlayer           int      `json:"void_kills_kit_basic_solo_baseball-player"`
	ArrowsShotKitBasicSoloBaseballPlayer          int      `json:"arrows_shot_kit_basic_solo_baseball-player"`
	GamesKitBasicSoloBaseballPlayer               int      `json:"games_kit_basic_solo_baseball-player"`
	ArrowsHitKitBasicSoloScout                    int      `json:"arrows_hit_kit_basic_solo_scout"`
	ArrowsShotKitBasicSoloScout                   int      `json:"arrows_shot_kit_basic_solo_scout"`
	LongestBowShotKitBasicSoloScout               int      `json:"longest_bow_shot_kit_basic_solo_scout"`
	ChestsOpenedKitBasicSoloHealer                int      `json:"chests_opened_kit_basic_solo_healer"`
	DeathsKitBasicSoloHealer                      int      `json:"deaths_kit_basic_solo_healer"`
	LossesKitBasicSoloHealer                      int      `json:"losses_kit_basic_solo_healer"`
	TimePlayedKitBasicSoloHealer                  int      `json:"time_played_kit_basic_solo_healer"`
	ArrowsHitKitDefendingTeamBaseballPlayer       int      `json:"arrows_hit_kit_defending_team_baseball-player"`
	LongestBowShotKitDefendingTeamBaseballPlayer  int      `json:"longest_bow_shot_kit_defending_team_baseball-player"`
	MostKillsGameKitDefendingTeamBaseballPlayer   int      `json:"most_kills_game_kit_defending_team_baseball-player"`
	VoidKillsKitDefendingTeamBaseballPlayer       int      `json:"void_kills_kit_defending_team_baseball-player"`
	ArrowsHitKitDefendingTeamArmorer              int      `json:"arrows_hit_kit_defending_team_armorer"`
	ArrowsShotKitDefendingTeamArmorer             int      `json:"arrows_shot_kit_defending_team_armorer"`
	BowKillsKitDefendingTeamArmorer               int      `json:"bow_kills_kit_defending_team_armorer"`
	LongestBowKillKitDefendingTeamArmorer         int      `json:"longest_bow_kill_kit_defending_team_armorer"`
	LongestBowShotKitDefendingTeamArmorer         int      `json:"longest_bow_shot_kit_defending_team_armorer"`
	MeleeKillsKitDefendingTeamArmorer             int      `json:"melee_kills_kit_defending_team_armorer"`
	MeleeKillsKitAttackingTeamScout               int      `json:"melee_kills_kit_attacking_team_scout"`
	Heads                                         int      `json:"heads"`
	HeadsKitAttackingTeamScout                    int      `json:"heads_kit_attacking_team_scout"`
	HeadsSalty                                    int      `json:"heads_salty"`
	HeadsSaltyKitAttackingTeamScout               int      `json:"heads_salty_kit_attacking_team_scout"`
	HeadsSaltySolo                                int      `json:"heads_salty_solo"`
	HeadsSolo                                     int      `json:"heads_solo"`
	HeadCollection                                struct {
		Recent []struct {
			UUID      string `json:"uuid"`
			Timestamp int64  `json:"timestamp"`
			Mode      string `json:"mode"`
			Sacrifice string `json:"sacrifice"`
		} `json:"recent"`
		Prestigious []struct {
			UUID      string `json:"uuid"`
			Timestamp int64  `json:"timestamp"`
			Mode      string `json:"mode"`
			Sacrifice string `json:"sacrifice"`
		} `json:"prestigious"`
	} `json:"head_collection"`
	SurvivedPlayersKitBasicSoloHealer        int    `json:"survived_players_kit_basic_solo_healer"`
	GamesKitBasicSoloHealer                  int    `json:"games_kit_basic_solo_healer"`
	KillsKitBasicSoloHealer                  int    `json:"kills_kit_basic_solo_healer"`
	LongestBowKillKitBasicSoloHealer         int    `json:"longest_bow_kill_kit_basic_solo_healer"`
	MeleeKillsKitBasicSoloHealer             int    `json:"melee_kills_kit_basic_solo_healer"`
	MostKillsGameKitBasicSoloHealer          int    `json:"most_kills_game_kit_basic_solo_healer"`
	VoidKillsKitBasicSoloHealer              int    `json:"void_kills_kit_basic_solo_healer"`
	AssistsKitBasicSoloScout                 int    `json:"assists_kit_basic_solo_scout"`
	ChestsOpenedKitAdvancedSoloKnight        int    `json:"chests_opened_kit_advanced_solo_knight"`
	DeathsKitAdvancedSoloKnight              int    `json:"deaths_kit_advanced_solo_knight"`
	GamesKitAdvancedSoloKnight               int    `json:"games_kit_advanced_solo_knight"`
	KillsKitAdvancedSoloKnight               int    `json:"kills_kit_advanced_solo_knight"`
	LongestBowKillKitAdvancedSoloKnight      int    `json:"longest_bow_kill_kit_advanced_solo_knight"`
	LossesKitAdvancedSoloKnight              int    `json:"losses_kit_advanced_solo_knight"`
	MeleeKillsKitAdvancedSoloKnight          int    `json:"melee_kills_kit_advanced_solo_knight"`
	MostKillsGameKitAdvancedSoloKnight       int    `json:"most_kills_game_kit_advanced_solo_knight"`
	SurvivedPlayersKitAdvancedSoloKnight     int    `json:"survived_players_kit_advanced_solo_knight"`
	TimePlayedKitAdvancedSoloKnight          int    `json:"time_played_kit_advanced_solo_knight"`
	AssistsKitAdvancedSoloArmorer            int    `json:"assists_kit_advanced_solo_armorer"`
	ChestsOpenedKitAdvancedSoloArmorer       int    `json:"chests_opened_kit_advanced_solo_armorer"`
	DeathsKitAdvancedSoloArmorer             int    `json:"deaths_kit_advanced_solo_armorer"`
	KillsKitAdvancedSoloArmorer              int    `json:"kills_kit_advanced_solo_armorer"`
	LongestBowKillKitAdvancedSoloArmorer     int    `json:"longest_bow_kill_kit_advanced_solo_armorer"`
	LossesKitAdvancedSoloArmorer             int    `json:"losses_kit_advanced_solo_armorer"`
	MeleeKillsKitAdvancedSoloArmorer         int    `json:"melee_kills_kit_advanced_solo_armorer"`
	MostKillsGameKitAdvancedSoloArmorer      int    `json:"most_kills_game_kit_advanced_solo_armorer"`
	SurvivedPlayersKitAdvancedSoloArmorer    int    `json:"survived_players_kit_advanced_solo_armorer"`
	TimePlayedKitAdvancedSoloArmorer         int    `json:"time_played_kit_advanced_solo_armorer"`
	ChestsOpenedKitDefendingTeamFarmer       int    `json:"chests_opened_kit_defending_team_farmer"`
	DeathsKitDefendingTeamFarmer             int    `json:"deaths_kit_defending_team_farmer"`
	KillsKitDefendingTeamFarmer              int    `json:"kills_kit_defending_team_farmer"`
	LongestBowKillKitDefendingTeamFarmer     int    `json:"longest_bow_kill_kit_defending_team_farmer"`
	LossesKitDefendingTeamFarmer             int    `json:"losses_kit_defending_team_farmer"`
	MeleeKillsKitDefendingTeamFarmer         int    `json:"melee_kills_kit_defending_team_farmer"`
	MostKillsGameKitDefendingTeamFarmer      int    `json:"most_kills_game_kit_defending_team_farmer"`
	SurvivedPlayersKitDefendingTeamFarmer    int    `json:"survived_players_kit_defending_team_farmer"`
	TimePlayedKitDefendingTeamFarmer         int    `json:"time_played_kit_defending_team_farmer"`
	SoloKnowledge                            int    `json:"solo_knowledge"`
	SkyWarsOpenedLegendaries                 int    `json:"SkyWars_openedLegendaries"`
	CosmeticTokens                           int    `json:"cosmetic_tokens"`
	ActiveDeathcry                           string `json:"active_deathcry"`
	ActiveProjectiletrail                    string `json:"active_projectiletrail"`
	KillstreakKitAttackingTeamScout          int    `json:"killstreak_kit_attacking_team_scout"`
	MegaInstantSmelting                      int    `json:"mega_instant_smelting"`
	HeadsEww                                 int    `json:"heads_eww"`
	HeadsEwwKitAttackingTeamScout            int    `json:"heads_eww_kit_attacking_team_scout"`
	HeadsEwwTeam                             int    `json:"heads_eww_team"`
	HeadsSaltyTeam                           int    `json:"heads_salty_team"`
	HeadsTeam                                int    `json:"heads_team"`
	SoloBridger                              int    `json:"solo_bridger"`
	TeamJuggernaut                           int    `json:"team_juggernaut"`
	HeadsDivine                              int    `json:"heads_divine"`
	HeadsDivineKitAttackingTeamScout         int    `json:"heads_divine_kit_attacking_team_scout"`
	HeadsDivineSolo                          int    `json:"heads_divine_solo"`
	HeadsMeh                                 int    `json:"heads_meh"`
	HeadsMehKitAttackingTeamScout            int    `json:"heads_meh_kit_attacking_team_scout"`
	HeadsMehSolo                             int    `json:"heads_meh_solo"`
	ChestsOpenedKitSupportingTeamEcologist   int    `json:"chests_opened_kit_supporting_team_ecologist"`
	TimePlayedKitSupportingTeamEcologist     int    `json:"time_played_kit_supporting_team_ecologist"`
	GamesKitSupportingTeamEcologist          int    `json:"games_kit_supporting_team_ecologist"`
	LongestBowKillKitSupportingTeamEcologist int    `json:"longest_bow_kill_kit_supporting_team_ecologist"`
	MeleeKillsKitSupportingTeamEcologist     int    `json:"melee_kills_kit_supporting_team_ecologist"`
	MostKillsGameKitSupportingTeamEcologist  int    `json:"most_kills_game_kit_supporting_team_ecologist"`
	FastestWinKitSupportingTeamEcologist     int    `json:"fastest_win_kit_supporting_team_ecologist"`
	KillstreakKitSupportingTeamEcologist     int    `json:"killstreak_kit_supporting_team_ecologist"`
	VoidKillsKitSupportingTeamEcologist      int    `json:"void_kills_kit_supporting_team_ecologist"`
	WinsKitSupportingTeamEcologist           int    `json:"wins_kit_supporting_team_ecologist"`
	ArrowsHitKitSupportingTeamEcologist      int    `json:"arrows_hit_kit_supporting_team_ecologist"`
	ArrowsShotKitSupportingTeamEcologist     int    `json:"arrows_shot_kit_supporting_team_ecologist"`
	LongestBowShotKitSupportingTeamEcologist int    `json:"longest_bow_shot_kit_supporting_team_ecologist"`
	ExtraWheels                              int    `json:"extra_wheels"`
	ChestsOpenedKitDefendingTeamGuardian     int    `json:"chests_opened_kit_defending_team_guardian"`
	MostKillsGameKitDefendingTeamGuardian    int    `json:"most_kills_game_kit_defending_team_guardian"`
	TimePlayedKitDefendingTeamGuardian       int    `json:"time_played_kit_defending_team_guardian"`
	VoidKillsKitDefendingTeamGuardian        int    `json:"void_kills_kit_defending_team_guardian"`
	HeadsDivineKitSupportingTeamHealer       int    `json:"heads_divine_kit_supporting_team_healer"`
	HeadsKitSupportingTeamHealer             int    `json:"heads_kit_supporting_team_healer"`
	HeadsYucky                               int    `json:"heads_yucky"`
	HeadsYuckyKitSupportingTeamHealer        int    `json:"heads_yucky_kit_supporting_team_healer"`
	HeadsYuckySolo                           int    `json:"heads_yucky_solo"`
	MegaBridger                              int    `json:"mega_bridger"`
	TeamSavior                               int    `json:"team_savior"`
	ChestsOpenedKitAttackingTeamEnderman     int    `json:"chests_opened_kit_attacking_team_enderman"`
	DeathsKitAttackingTeamEnderman           int    `json:"deaths_kit_attacking_team_enderman"`
	KillsKitAttackingTeamEnderman            int    `json:"kills_kit_attacking_team_enderman"`
	LossesKitAttackingTeamEnderman           int    `json:"losses_kit_attacking_team_enderman"`
	MostKillsGameKitAttackingTeamEnderman    int    `json:"most_kills_game_kit_attacking_team_enderman"`
	SurvivedPlayersKitAttackingTeamEnderman  int    `json:"survived_players_kit_attacking_team_enderman"`
	TimePlayedKitAttackingTeamEnderman       int    `json:"time_played_kit_attacking_team_enderman"`
	VoidKillsKitAttackingTeamEnderman        int    `json:"void_kills_kit_attacking_team_enderman"`
	LongestBowKillKitAttackingTeamEnderman   int    `json:"longest_bow_kill_kit_attacking_team_enderman"`
	MeleeKillsKitAttackingTeamEnderman       int    `json:"melee_kills_kit_attacking_team_enderman"`
	ArrowsHitKitAttackingTeamEnderman        int    `json:"arrows_hit_kit_attacking_team_enderman"`
	ArrowsShotKitAttackingTeamEnderman       int    `json:"arrows_shot_kit_attacking_team_enderman"`
	FastestWinKitAttackingTeamEnderman       int    `json:"fastest_win_kit_attacking_team_enderman"`
	GamesKitAttackingTeamEnderman            int    `json:"games_kit_attacking_team_enderman"`
	KillstreakKitAttackingTeamEnderman       int    `json:"killstreak_kit_attacking_team_enderman"`
	LongestBowShotKitAttackingTeamEnderman   int    `json:"longest_bow_shot_kit_attacking_team_enderman"`
	WinsKitAttackingTeamEnderman             int    `json:"wins_kit_attacking_team_enderman"`
	AssistsKitAttackingTeamEnderman          int    `json:"assists_kit_attacking_team_enderman"`
	ArrowsHitKitDefendingTeamFarmer          int    `json:"arrows_hit_kit_defending_team_farmer"`
	ArrowsShotKitDefendingTeamFarmer         int    `json:"arrows_shot_kit_defending_team_farmer"`
	GamesKitDefendingTeamFarmer              int    `json:"games_kit_defending_team_farmer"`
	LongestBowShotKitDefendingTeamFarmer     int    `json:"longest_bow_shot_kit_defending_team_farmer"`
	VoidKillsKitDefendingTeamFarmer          int    `json:"void_kills_kit_defending_team_farmer"`
	FreeEventKeySkywarsChristmasBoxes2019    bool   `json:"free_event_key_skywars_christmas_boxes_2019"`
	SkywarsChristmasBoxes                    int    `json:"skywars_christmas_boxes"`
	FastestWinKitDefendingTeamFarmer         int    `json:"fastest_win_kit_defending_team_farmer"`
	KillstreakKitDefendingTeamFarmer         int    `json:"killstreak_kit_defending_team_farmer"`
	WinsKitDefendingTeamFarmer               int    `json:"wins_kit_defending_team_farmer"`
	AssistsKitDefendingTeamFarmer            int    `json:"assists_kit_defending_team_farmer"`
	InGamePresentsCap20197                   int    `json:"inGamePresentsCap_2019_7"`
	ArrowsHitKitDefendingTeamDisco           int    `json:"arrows_hit_kit_defending_team_disco"`
	ArrowsShotKitDefendingTeamDisco          int    `json:"arrows_shot_kit_defending_team_disco"`
	ChestsOpenedKitDefendingTeamDisco        int    `json:"chests_opened_kit_defending_team_disco"`
	DeathsKitDefendingTeamDisco              int    `json:"deaths_kit_defending_team_disco"`
	GamesKitDefendingTeamDisco               int    `json:"games_kit_defending_team_disco"`
	LongestBowShotKitDefendingTeamDisco      int    `json:"longest_bow_shot_kit_defending_team_disco"`
	LossesKitDefendingTeamDisco              int    `json:"losses_kit_defending_team_disco"`
	SurvivedPlayersKitDefendingTeamDisco     int    `json:"survived_players_kit_defending_team_disco"`
	TimePlayedKitDefendingTeamDisco          int    `json:"time_played_kit_defending_team_disco"`
	KillsKitDefendingTeamDisco               int    `json:"kills_kit_defending_team_disco"`
	LongestBowKillKitDefendingTeamDisco      int    `json:"longest_bow_kill_kit_defending_team_disco"`
	MeleeKillsKitDefendingTeamDisco          int    `json:"melee_kills_kit_defending_team_disco"`
	MostKillsGameKitDefendingTeamDisco       int    `json:"most_kills_game_kit_defending_team_disco"`
	HeadsEwwKitSupportingTeamHealer          int    `json:"heads_eww_kit_supporting_team_healer"`
	HeadsEwwSolo                             int    `json:"heads_eww_solo"`
	HeadsSaltyKitSupportingTeamHealer        int    `json:"heads_salty_kit_supporting_team_healer"`
	HeadsKitSupportingTeamEcologist          int    `json:"heads_kit_supporting_team_ecologist"`
	HeadsYuckyKitSupportingTeamEcologist     int    `json:"heads_yucky_kit_supporting_team_ecologist"`
	InGamePresentsCap201911                  int    `json:"inGamePresentsCap_2019_11"`
	InGamePresentsCap201916                  int    `json:"inGamePresentsCap_2019_16"`
}

// BattlegroundStats is Battleground stats about a player.
type BattlegroundStats struct {
	WarriorSpec               string   `json:"warrior_spec"`
	ShamanSpec                string   `json:"shaman_spec"`
	MageSpec                  string   `json:"mage_spec"`
	SelectedMount             string   `json:"selected_mount"`
	PaladinSpec               string   `json:"paladin_spec"`
	Packages                  []string `json:"packages"`
	ChosenClass               string   `json:"chosen_class"`
	PlayStreak                int      `json:"play_streak"`
	AvengerPlays              int      `json:"avenger_plays"`
	DamagePreventedPaladin    int      `json:"damage_prevented_paladin"`
	Deaths                    int      `json:"deaths"`
	DamagePrevented           int      `json:"damage_prevented"`
	Damage                    int      `json:"damage"`
	HealAvenger               int      `json:"heal_avenger"`
	DamageAvenger             int      `json:"damage_avenger"`
	PaladinPlays              int      `json:"paladin_plays"`
	Assists                   int      `json:"assists"`
	Heal                      int      `json:"heal"`
	DamagePreventedAvenger    int      `json:"damage_prevented_avenger"`
	HealPaladin               int      `json:"heal_paladin"`
	Coins                     int      `json:"coins"`
	DamagePaladin             int      `json:"damage_paladin"`
	DamageTaken               int      `json:"damage_taken"`
	DamagePyromancer          int      `json:"damage_pyromancer"`
	HealMage                  int      `json:"heal_mage"`
	DamageMage                int      `json:"damage_mage"`
	HealPyromancer            int      `json:"heal_pyromancer"`
	FlagConquerTeam           int      `json:"flag_conquer_team"`
	PyromancerPlays           int      `json:"pyromancer_plays"`
	DamagePreventedPyromancer int      `json:"damage_prevented_pyromancer"`
	MagePlays                 int      `json:"mage_plays"`
	Kills                     int      `json:"kills"`
	DamagePreventedMage       int      `json:"damage_prevented_mage"`
}

// TrueCombatStats is true combat stats about a player.
type TrueCombatStats struct {
	Packages                         []string `json:"packages"`
	WinStreak                        int      `json:"win_streak"`
	Games                            int      `json:"games"`
	CrazywallsLossesSolo             int      `json:"crazywalls_losses_solo"`
	ItemsEnchanted                   int      `json:"items_enchanted"`
	Deaths                           int      `json:"deaths"`
	Coins                            int      `json:"coins"`
	CrazywallsDeathsSolo             int      `json:"crazywalls_deaths_solo"`
	Losses                           int      `json:"losses"`
	CrazywallsGamesSolo              int      `json:"crazywalls_games_solo"`
	SurvivedPlayers                  int      `json:"survived_players"`
	Kills                            int      `json:"kills"`
	CrazywallsKillsMonthlyBSolo      int      `json:"crazywalls_kills_monthly_b_solo"`
	CrazywallsKillsSolo              int      `json:"crazywalls_kills_solo"`
	CrazywallsKillsWeeklyASolo       int      `json:"crazywalls_kills_weekly_a_solo"`
	CrazywallsLossesSoloChaos        int      `json:"crazywalls_losses_solo_chaos"`
	CrazywallsGamesSoloChaos         int      `json:"crazywalls_games_solo_chaos"`
	ArrowsShot                       int      `json:"arrows_shot"`
	CrazywallsDeathsSoloChaos        int      `json:"crazywalls_deaths_solo_chaos"`
	CrazywallsKillsMonthlyBSoloChaos int      `json:"crazywalls_kills_monthly_b_solo_chaos"`
	CrazywallsKillsSoloChaos         int      `json:"crazywalls_kills_solo_chaos"`
	CrazywallsKillsWeeklyASoloChaos  int      `json:"crazywalls_kills_weekly_a_solo_chaos"`
	CrazywallsDeathsTeam             int      `json:"crazywalls_deaths_team"`
	CrazywallsLossesTeam             int      `json:"crazywalls_losses_team"`
	CrazywallsGamesTeam              int      `json:"crazywalls_games_team"`
	CrazywallsDeathsTeamChaos        int      `json:"crazywalls_deaths_team_chaos"`
	CrazywallsGamesTeamChaos         int      `json:"crazywalls_games_team_chaos"`
	CrazywallsLossesTeamChaos        int      `json:"crazywalls_losses_team_chaos"`
	CrazywallsKillsWeeklyBSolo       int      `json:"crazywalls_kills_weekly_b_solo"`
	SkullsGathered                   int      `json:"skulls_gathered"`
	CrazywallsKillsMonthlyASolo      int      `json:"crazywalls_kills_monthly_a_solo"`
	GoldenSkulls                     int      `json:"golden_skulls"`
	GiantZombie                      int      `json:"giant_zombie"`
	ActiveKitTeam                    string   `json:"activeKit_Team"`
	CrazywallsKillsWeeklyBTeam       int      `json:"crazywalls_kills_weekly_b_team"`
	CrazywallsKillsTeam              int      `json:"crazywalls_kills_team"`
	CrazywallsKillsMonthlyATeam      int      `json:"crazywalls_kills_monthly_a_team"`
	CrazywallsWinsTeamChaos          int      `json:"crazywalls_wins_team_chaos"`
	Wins                             int      `json:"wins"`
}

// UHCStats is UHC stats about a player.
type UHCStats struct {
	Coins int `json:"coins"`
}

// WallsStats is Walls stats about a player.
type WallsStats struct {
	Losses int `json:"losses"`
}

// QuakeStats is Quake stats about a player.
type QuakeStats struct {
	Packages                        []string `json:"packages"`
	KillsTeams                      int      `json:"kills_teams"`
	Coins                           int      `json:"coins"`
	DeathsTeams                     int      `json:"deaths_teams"`
	WinsTeams                       int      `json:"wins_teams"`
	CompassSelected                 bool     `json:"compass_selected"`
	AlternativeGunCooldownIndicator bool     `json:"alternative_gun_cooldown_indicator"`
	EnableSound                     bool     `json:"enable_sound"`
	InstantRespawn                  bool     `json:"instantRespawn"`
	HighestKillstreak               int      `json:"highest_killstreak"`
	ShowDashCooldown                bool     `json:"showDashCooldown"`
	Kills                           int      `json:"kills"`
	Headshots                       int      `json:"headshots"`
	DistanceTravelled               int      `json:"distance_travelled"`
	ShotsFired                      int      `json:"shots_fired"`
	KillsSinceUpdateFeb2017         int      `json:"kills_since_update_feb_2017"`
	Deaths                          int      `json:"deaths"`
	Barrel                          string   `json:"barrel"`
	MessageOthersKillsDeaths        bool     `json:"messageOthers' Kills/deaths"`
	MessageYourKills                bool     `json:"messageYour Kills"`
	MessageCoinMessages             bool     `json:"messageCoin Messages"`
	MessageKillstreaks              bool     `json:"messageKillstreaks"`
	MessageYourDeaths               bool     `json:"messageYour Deaths"`
	MessagePowerupCollections       bool     `json:"messagePowerup Collections"`
	MessageMultiKills               bool     `json:"messageMulti-kills"`
	Killstreaks                     int      `json:"killstreaks"`
	DashCooldown                    string   `json:"dash_cooldown"`
	KillsDmTeams                    int      `json:"kills_dm_teams"`
	KillsTimeattack                 int      `json:"kills_timeattack"`
	KillsDm                         int      `json:"kills_dm"`
	ShowKillPrefix                  bool     `json:"showKillPrefix"`
}

// ArcadeStats is arcade stats about a player.
type ArcadeStats struct {
	Coins                                   int    `json:"coins"`
	RoundsSantaSays                         int    `json:"rounds_santa_says"`
	MonthlyCoinsB                           int    `json:"monthly_coins_b"`
	WeeklyCoinsA                            int    `json:"weekly_coins_a"`
	MiniwallsActiveKit                      string `json:"miniwalls_activeKit"`
	ArrowsHitMiniWalls                      int    `json:"arrows_hit_mini_walls"`
	MonthlyCoinsA                           int    `json:"monthly_coins_a"`
	KillsMiniWalls                          int    `json:"kills_mini_walls"`
	DeathsMiniWalls                         int    `json:"deaths_mini_walls"`
	ArrowsShotMiniWalls                     int    `json:"arrows_shot_mini_walls"`
	BasicZombieKillsZombies                 int    `json:"basic_zombie_kills_zombies"`
	BestRoundZombies                        int    `json:"best_round_zombies"`
	BestRoundZombiesDeadend                 int    `json:"best_round_zombies_deadend"`
	BestRoundZombiesDeadendNormal           int    `json:"best_round_zombies_deadend_normal"`
	BulletsHitZombies                       int    `json:"bullets_hit_zombies"`
	BulletsShotZombies                      int    `json:"bullets_shot_zombies"`
	DeathsZombies                           int    `json:"deaths_zombies"`
	DeathsZombiesDeadend                    int    `json:"deaths_zombies_deadend"`
	DeathsZombiesDeadendNormal              int    `json:"deaths_zombies_deadend_normal"`
	DoorsOpenedZombies                      int    `json:"doors_opened_zombies"`
	DoorsOpenedZombiesDeadend               int    `json:"doors_opened_zombies_deadend"`
	DoorsOpenedZombiesDeadendNormal         int    `json:"doors_opened_zombies_deadend_normal"`
	EmpoweredZombieKillsZombies             int    `json:"empowered_zombie_kills_zombies"`
	FastestTime10Zombies                    int    `json:"fastest_time_10_zombies"`
	FastestTime10ZombiesDeadendNormal       int    `json:"fastest_time_10_zombies_deadend_normal"`
	FireZombieKillsZombies                  int    `json:"fire_zombie_kills_zombies"`
	HeadshotsZombies                        int    `json:"headshots_zombies"`
	MagmaCubeZombieKillsZombies             int    `json:"magma_cube_zombie_kills_zombies"`
	MagmaZombieKillsZombies                 int    `json:"magma_zombie_kills_zombies"`
	PigZombieZombieKillsZombies             int    `json:"pig_zombie_zombie_kills_zombies"`
	PlayersRevivedZombies                   int    `json:"players_revived_zombies"`
	PlayersRevivedZombiesDeadend            int    `json:"players_revived_zombies_deadend"`
	PlayersRevivedZombiesDeadendNormal      int    `json:"players_revived_zombies_deadend_normal"`
	TimesKnockedDownZombies                 int    `json:"times_knocked_down_zombies"`
	TimesKnockedDownZombiesDeadend          int    `json:"times_knocked_down_zombies_deadend"`
	TimesKnockedDownZombiesDeadendNormal    int    `json:"times_knocked_down_zombies_deadend_normal"`
	TntBabyZombieKillsZombies               int    `json:"tnt_baby_zombie_kills_zombies"`
	TotalRoundsSurvivedZombies              int    `json:"total_rounds_survived_zombies"`
	TotalRoundsSurvivedZombiesDeadend       int    `json:"total_rounds_survived_zombies_deadend"`
	TotalRoundsSurvivedZombiesDeadendNormal int    `json:"total_rounds_survived_zombies_deadend_normal"`
	WindowsRepairedZombies                  int    `json:"windows_repaired_zombies"`
	WindowsRepairedZombiesDeadend           int    `json:"windows_repaired_zombies_deadend"`
	WindowsRepairedZombiesDeadendNormal     int    `json:"windows_repaired_zombies_deadend_normal"`
	WolfZombieKillsZombies                  int    `json:"wolf_zombie_kills_zombies"`
	ZombieKillsZombies                      int    `json:"zombie_kills_zombies"`
	ZombieKillsZombiesDeadend               int    `json:"zombie_kills_zombies_deadend"`
	ZombieKillsZombiesDeadendNormal         int    `json:"zombie_kills_zombies_deadend_normal"`
}

// ArenaStats is Arena stats about a player.
type ArenaStats struct {
	Rating        float64 `json:"rating"`
	WinStreaks4V4 int     `json:"win_streaks_4v4"`
	Coins         int     `json:"coins"`
	Wins4V4       int     `json:"wins_4v4"`
	Healed4V4     int     `json:"healed_4v4"`
	Kills4V4      int     `json:"kills_4v4"`
	Damage4V4     int     `json:"damage_4v4"`
	Deaths4V4     int     `json:"deaths_4v4"`
	Games4V4      int     `json:"games_4v4"`
	Damage1V1     int     `json:"damage_1v1"`
	Games1V1      int     `json:"games_1v1"`
	WinStreaks1V1 int     `json:"win_streaks_1v1"`
	Wins1V1       int     `json:"wins_1v1"`
	Healed1V1     int     `json:"healed_1v1"`
	Deaths2V2     int     `json:"deaths_2v2"`
	WinStreaks2V2 int     `json:"win_streaks_2v2"`
	Damage2V2     int     `json:"damage_2v2"`
	Wins2V2       int     `json:"wins_2v2"`
	Healed2V2     int     `json:"healed_2v2"`
	Games2V2      int     `json:"games_2v2"`
	Losses2V2     int     `json:"losses_2v2"`
	Wins          int     `json:"wins"`
}

// SuperSmashStats is Super Smash stats about a player.
type SuperSmashStats struct {
	LastLevelTHEBULK int `json:"lastLevel_THE_BULK"`
	SmashLevel       int `json:"smashLevel"`
}

// GingerBreadStats is Gingerbread stats about a player
type GingerBreadStats struct {
	FrameActive   string   `json:"frame_active"`
	PantsActive   string   `json:"pants_active"`
	ShoesActive   string   `json:"shoes_active"`
	HelmetActive  string   `json:"helmet_active"`
	BoosterActive string   `json:"booster_active"`
	EngineActive  string   `json:"engine_active"`
	JacketActive  string   `json:"jacket_active"`
	Packages      []string `json:"packages"`
	SkinActive    string   `json:"skin_active"`
	Coins         int      `json:"coins"`
}

// VampireZStats is VampireZ stats about a player.
type VampireZStats struct {
	UpdatedStats bool `json:"updated_stats"`
}

// BedwarsStats is Bedwars stats about a player.
type BedwarsStats struct {
	Experience                                int      `json:"Experience"`
	FirstJoin7                                bool     `json:"first_join_7"`
	BedwarsBoxes                              int      `json:"bedwars_boxes"`
	GamesPlayedBedwars1                       int      `json:"games_played_bedwars_1"`
	Favourites2                               string   `json:"favourites_2"`
	EightTwoWinstreak                         int      `json:"eight_two_winstreak"`
	FavoriteSlots                             string   `json:"favorite_slots"`
	FinalDeathsBedwars                        int      `json:"final_deaths_bedwars"`
	GoldResourcesCollectedBedwars             int      `json:"gold_resources_collected_bedwars"`
	BedsLostBedwars                           int      `json:"beds_lost_bedwars"`
	EightTwoFinalDeathsBedwars                int      `json:"eight_two_final_deaths_bedwars"`
	EightTwoItemsPurchasedBedwars             int      `json:"eight_two__items_purchased_bedwars"`
	EightTwoEntityAttackFinalDeathsBedwars    int      `json:"eight_two_entity_attack_final_deaths_bedwars"`
	EntityAttackFinalDeathsBedwars            int      `json:"entity_attack_final_deaths_bedwars"`
	EightTwoGamesPlayedBedwars                int      `json:"eight_two_games_played_bedwars"`
	ResourcesCollectedBedwars                 int      `json:"resources_collected_bedwars"`
	EightTwoResourcesCollectedBedwars         int      `json:"eight_two_resources_collected_bedwars"`
	EightTwoBedsLostBedwars                   int      `json:"eight_two_beds_lost_bedwars"`
	Coins                                     int      `json:"coins"`
	EightTwoGoldResourcesCollectedBedwars     int      `json:"eight_two_gold_resources_collected_bedwars"`
	LossesBedwars                             int      `json:"losses_bedwars"`
	GamesPlayedBedwars                        int      `json:"games_played_bedwars"`
	IronResourcesCollectedBedwars             int      `json:"iron_resources_collected_bedwars"`
	EightTwoIronResourcesCollectedBedwars     int      `json:"eight_two_iron_resources_collected_bedwars"`
	EightTwoLossesBedwars                     int      `json:"eight_two_losses_bedwars"`
	ItemsPurchasedBedwars                     int      `json:"_items_purchased_bedwars"`
	EightOneWinstreak                         int      `json:"eight_one_winstreak"`
	DiamondResourcesCollectedBedwars          int      `json:"diamond_resources_collected_bedwars"`
	EightTwoDiamondResourcesCollectedBedwars  int      `json:"eight_two_diamond_resources_collected_bedwars"`
	EightTwoKillsBedwars                      int      `json:"eight_two_kills_bedwars"`
	EightTwoEntityAttackKillsBedwars          int      `json:"eight_two_entity_attack_kills_bedwars"`
	KillsBedwars                              int      `json:"kills_bedwars"`
	EntityAttackKillsBedwars                  int      `json:"entity_attack_kills_bedwars"`
	DeathsBedwars                             int      `json:"deaths_bedwars"`
	EightTwoDeathsBedwars                     int      `json:"eight_two_deaths_bedwars"`
	EightTwoEntityAttackDeathsBedwars         int      `json:"eight_two_entity_attack_deaths_bedwars"`
	EntityAttackDeathsBedwars                 int      `json:"entity_attack_deaths_bedwars"`
	EightOneBedsLostBedwars                   int      `json:"eight_one_beds_lost_bedwars"`
	EightOneFinalDeathsBedwars                int      `json:"eight_one_final_deaths_bedwars"`
	EightOneDeathsBedwars                     int      `json:"eight_one_deaths_bedwars"`
	EightOneGamesPlayedBedwars                int      `json:"eight_one_games_played_bedwars"`
	EightOneEntityAttackDeathsBedwars         int      `json:"eight_one_entity_attack_deaths_bedwars"`
	EightOneEntityAttackFinalDeathsBedwars    int      `json:"eight_one_entity_attack_final_deaths_bedwars"`
	EightOneLossesBedwars                     int      `json:"eight_one_losses_bedwars"`
	EightOneIronResourcesCollectedBedwars     int      `json:"eight_one_iron_resources_collected_bedwars"`
	EightOneItemsPurchasedBedwars             int      `json:"eight_one__items_purchased_bedwars"`
	EightOneResourcesCollectedBedwars         int      `json:"eight_one_resources_collected_bedwars"`
	EightOneGoldResourcesCollectedBedwars     int      `json:"eight_one_gold_resources_collected_bedwars"`
	QuickjoinUsesRandom                       int      `json:"quickjoin_uses_random"`
	QuickjoinUsesTotal                        int      `json:"quickjoin_uses_total"`
	BedsBrokenBedwars                         int      `json:"beds_broken_bedwars"`
	EightOneBedsBrokenBedwars                 int      `json:"eight_one_beds_broken_bedwars"`
	EightOneDiamondResourcesCollectedBedwars  int      `json:"eight_one_diamond_resources_collected_bedwars"`
	EightOneEntityAttackFinalKillsBedwars     int      `json:"eight_one_entity_attack_final_kills_bedwars"`
	EightOneEntityAttackKillsBedwars          int      `json:"eight_one_entity_attack_kills_bedwars"`
	EightOneFinalKillsBedwars                 int      `json:"eight_one_final_kills_bedwars"`
	EightOneKillsBedwars                      int      `json:"eight_one_kills_bedwars"`
	EntityAttackFinalKillsBedwars             int      `json:"entity_attack_final_kills_bedwars"`
	FinalKillsBedwars                         int      `json:"final_kills_bedwars"`
	EightOneVoidKillsBedwars                  int      `json:"eight_one_void_kills_bedwars"`
	VoidKillsBedwars                          int      `json:"void_kills_bedwars"`
	EightTwoBedsBrokenBedwars                 int      `json:"eight_two_beds_broken_bedwars"`
	EightTwoFinalKillsBedwars                 int      `json:"eight_two_final_kills_bedwars"`
	EightTwoProjectileDeathsBedwars           int      `json:"eight_two_projectile_deaths_bedwars"`
	EightTwoVoidDeathsBedwars                 int      `json:"eight_two_void_deaths_bedwars"`
	EightTwoVoidFinalDeathsBedwars            int      `json:"eight_two_void_final_deaths_bedwars"`
	EightTwoVoidFinalKillsBedwars             int      `json:"eight_two_void_final_kills_bedwars"`
	EightTwoVoidKillsBedwars                  int      `json:"eight_two_void_kills_bedwars"`
	ProjectileDeathsBedwars                   int      `json:"projectile_deaths_bedwars"`
	VoidDeathsBedwars                         int      `json:"void_deaths_bedwars"`
	VoidFinalDeathsBedwars                    int      `json:"void_final_deaths_bedwars"`
	VoidFinalKillsBedwars                     int      `json:"void_final_kills_bedwars"`
	EightTwoEmeraldResourcesCollectedBedwars  int      `json:"eight_two_emerald_resources_collected_bedwars"`
	EightTwoEntityAttackFinalKillsBedwars     int      `json:"eight_two_entity_attack_final_kills_bedwars"`
	EightTwoEntityExplosionDeathsBedwars      int      `json:"eight_two_entity_explosion_deaths_bedwars"`
	EightTwoFallDeathsBedwars                 int      `json:"eight_two_fall_deaths_bedwars"`
	EightTwoPermanentItemsPurchasedBedwars    int      `json:"eight_two_permanent _items_purchased_bedwars"`
	EightTwoWinsBedwars                       int      `json:"eight_two_wins_bedwars"`
	EmeraldResourcesCollectedBedwars          int      `json:"emerald_resources_collected_bedwars"`
	EntityExplosionDeathsBedwars              int      `json:"entity_explosion_deaths_bedwars"`
	FallDeathsBedwars                         int      `json:"fall_deaths_bedwars"`
	PermanentItemsPurchasedBedwars            int      `json:"permanent _items_purchased_bedwars"`
	WinsBedwars                               int      `json:"wins_bedwars"`
	Winstreak                                 int      `json:"winstreak"`
	EightTwoFallKillsBedwars                  int      `json:"eight_two_fall_kills_bedwars"`
	FallKillsBedwars                          int      `json:"fall_kills_bedwars"`
	EightTwoFallFinalDeathsBedwars            int      `json:"eight_two_fall_final_deaths_bedwars"`
	FallFinalDeathsBedwars                    int      `json:"fall_final_deaths_bedwars"`
	EightOneEntityExplosionDeathsBedwars      int      `json:"eight_one_entity_explosion_deaths_bedwars"`
	EightOnePermanentItemsPurchasedBedwars    int      `json:"eight_one_permanent _items_purchased_bedwars"`
	EightOneVoidFinalKillsBedwars             int      `json:"eight_one_void_final_kills_bedwars"`
	EightOneEmeraldResourcesCollectedBedwars  int      `json:"eight_one_emerald_resources_collected_bedwars"`
	EightOneFallKillsBedwars                  int      `json:"eight_one_fall_kills_bedwars"`
	EightOneVoidDeathsBedwars                 int      `json:"eight_one_void_deaths_bedwars"`
	EightOneWinsBedwars                       int      `json:"eight_one_wins_bedwars"`
	Packages                                  []string `json:"packages"`
	BedwarsOpenedChests                       int      `json:"Bedwars_openedChests"`
	ChestHistoryNew                           []string `json:"chest_history_new"`
	BedwarsOpenedCommons                      int      `json:"Bedwars_openedCommons"`
	BedwarsOpenedRares                        int      `json:"Bedwars_openedRares"`
	BedwarsOpenedEpics                        int      `json:"Bedwars_openedEpics"`
	ActiveBedDestroy                          string   `json:"activeBedDestroy"`
	EightTwoEntityExplosionKillsBedwars       int      `json:"eight_two_entity_explosion_kills_bedwars"`
	EntityExplosionKillsBedwars               int      `json:"entity_explosion_kills_bedwars"`
	ActiveProjectileTrail                     string   `json:"activeProjectileTrail"`
	ActiveNPCSkin                             string   `json:"activeNPCSkin"`
	EightOneVoidFinalDeathsBedwars            int      `json:"eight_one_void_final_deaths_bedwars"`
	ActiveVictoryDance                        string   `json:"activeVictoryDance"`
	ActiveKillEffect                          string   `json:"activeKillEffect"`
	ActiveIslandTopper                        string   `json:"activeIslandTopper"`
	ActiveDeathCry                            string   `json:"activeDeathCry"`
	ActiveKillMessages                        string   `json:"activeKillMessages"`
	ActiveGlyph                               string   `json:"activeGlyph"`
	QuickjoinUsesSpeedway                     int      `json:"quickjoin_uses_Speedway"`
	EightTwoFallFinalKillsBedwars             int      `json:"eight_two_fall_final_kills_bedwars"`
	FallFinalKillsBedwars                     int      `json:"fall_final_kills_bedwars"`
	ShopSort                                  string   `json:"shop_sort"`
	BedwarsOpenedLegendaries                  int      `json:"Bedwars_openedLegendaries"`
	FourFourWinstreak                         int      `json:"four_four_winstreak"`
	FourFourBedsLostBedwars                   int      `json:"four_four_beds_lost_bedwars"`
	FourFourDeathsBedwars                     int      `json:"four_four_deaths_bedwars"`
	FourFourDiamondResourcesCollectedBedwars  int      `json:"four_four_diamond_resources_collected_bedwars"`
	FourFourEntityAttackKillsBedwars          int      `json:"four_four_entity_attack_kills_bedwars"`
	FourFourFallDeathsBedwars                 int      `json:"four_four_fall_deaths_bedwars"`
	FourFourFinalDeathsBedwars                int      `json:"four_four_final_deaths_bedwars"`
	FourFourGamesPlayedBedwars                int      `json:"four_four_games_played_bedwars"`
	FourFourGoldResourcesCollectedBedwars     int      `json:"four_four_gold_resources_collected_bedwars"`
	FourFourIronResourcesCollectedBedwars     int      `json:"four_four_iron_resources_collected_bedwars"`
	FourFourItemsPurchasedBedwars             int      `json:"four_four_items_purchased_bedwars"`
	FourFourKillsBedwars                      int      `json:"four_four_kills_bedwars"`
	FourFourLossesBedwars                     int      `json:"four_four_losses_bedwars"`
	FourFourPermanentItemsPurchasedBedwars    int      `json:"four_four_permanent _items_purchased_bedwars"`
	FourFourResourcesCollectedBedwars         int      `json:"four_four_resources_collected_bedwars"`
	FourFourVoidFinalDeathsBedwars            int      `json:"four_four_void_final_deaths_bedwars"`
	FourFourVoidKillsBedwars                  int      `json:"four_four_void_kills_bedwars"`
	EightOneFireTickKillsBedwars              int      `json:"eight_one_fire_tick_kills_bedwars"`
	FireTickKillsBedwars                      int      `json:"fire_tick_kills_bedwars"`
	EightOneFallFinalDeathsBedwars            int      `json:"eight_one_fall_final_deaths_bedwars"`
	ShopSortEnableOwnedFirst                  bool     `json:"shop_sort_enable_owned_first"`
	ActiveSprays                              string   `json:"activeSprays"`
	EightTwoProjectileKillsBedwars            int      `json:"eight_two_projectile_kills_bedwars"`
	ProjectileKillsBedwars                    int      `json:"projectile_kills_bedwars"`
	EightTwoEntityExplosionFinalDeathsBedwars int      `json:"eight_two_entity_explosion_final_deaths_bedwars"`
	EntityExplosionFinalDeathsBedwars         int      `json:"entity_explosion_final_deaths_bedwars"`
	EightOneFallDeathsBedwars                 int      `json:"eight_one_fall_deaths_bedwars"`
	BedwarsHalloweenBoxes                     int      `json:"bedwars_halloween_boxes"`
	FreeEventKeyBedwarsHalloweenBoxes2019     bool     `json:"free_event_key_bedwars_halloween_boxes_2019"`
	QuickjoinUsesPernicious                   int      `json:"quickjoin_uses_Pernicious"`
	EightTwoEntityExplosionFinalKillsBedwars  int      `json:"eight_two_entity_explosion_final_kills_bedwars"`
	EntityExplosionFinalKillsBedwars          int      `json:"entity_explosion_final_kills_bedwars"`
	BedwarsChristmasBoxes                     int      `json:"bedwars_christmas_boxes"`
	FreeEventKeyBedwarsChristmasBoxes2019     bool     `json:"free_event_key_bedwars_christmas_boxes_2019"`
	EightTwoFireTickFinalDeathsBedwars        int      `json:"eight_two_fire_tick_final_deaths_bedwars"`
	FireTickFinalDeathsBedwars                int      `json:"fire_tick_final_deaths_bedwars"`
	FourThreeWinstreak                        int      `json:"four_three_winstreak"`
	FourThreeBedsBrokenBedwars                int      `json:"four_three_beds_broken_bedwars"`
	FourThreeBedsLostBedwars                  int      `json:"four_three_beds_lost_bedwars"`
	FourThreeDeathsBedwars                    int      `json:"four_three_deaths_bedwars"`
	FourThreeDiamondResourcesCollectedBedwars int      `json:"four_three_diamond_resources_collected_bedwars"`
	FourThreeEntityAttackDeathsBedwars        int      `json:"four_three_entity_attack_deaths_bedwars"`
	FourThreeEntityAttackFinalDeathsBedwars   int      `json:"four_three_entity_attack_final_deaths_bedwars"`
	FourThreeEntityAttackFinalKillsBedwars    int      `json:"four_three_entity_attack_final_kills_bedwars"`
	FourThreeFinalDeathsBedwars               int      `json:"four_three_final_deaths_bedwars"`
	FourThreeFinalKillsBedwars                int      `json:"four_three_final_kills_bedwars"`
	FourThreeGamesPlayedBedwars               int      `json:"four_three_games_played_bedwars"`
	FourThreeGoldResourcesCollectedBedwars    int      `json:"four_three_gold_resources_collected_bedwars"`
	FourThreeIronResourcesCollectedBedwars    int      `json:"four_three_iron_resources_collected_bedwars"`
	FourThreeItemsPurchasedBedwars            int      `json:"four_three_items_purchased_bedwars"`
	FourThreeLossesBedwars                    int      `json:"four_three_losses_bedwars"`
	FourThreePermanentItemsPurchasedBedwars   int      `json:"four_three_permanent _items_purchased_bedwars"`
	FourThreeResourcesCollectedBedwars        int      `json:"four_three_resources_collected_bedwars"`
	FourThreeVoidDeathsBedwars                int      `json:"four_three_void_deaths_bedwars"`
	FourThreeEmeraldResourcesCollectedBedwars int      `json:"four_three_emerald_resources_collected_bedwars"`
	FourThreeEntityAttackKillsBedwars         int      `json:"four_three_entity_attack_kills_bedwars"`
	FourThreeFallDeathsBedwars                int      `json:"four_three_fall_deaths_bedwars"`
	FourThreeKillsBedwars                     int      `json:"four_three_kills_bedwars"`
	FourThreeVoidKillsBedwars                 int      `json:"four_three_void_kills_bedwars"`
	FourThreeVoidFinalDeathsBedwars           int      `json:"four_three_void_final_deaths_bedwars"`
	EightTwoFireFinalDeathsBedwars            int      `json:"eight_two_fire_final_deaths_bedwars"`
	FireFinalDeathsBedwars                    int      `json:"fire_final_deaths_bedwars"`
	EightTwoFireTickDeathsBedwars             int      `json:"eight_two_fire_tick_deaths_bedwars"`
	FireTickDeathsBedwars                     int      `json:"fire_tick_deaths_bedwars"`
}

// LegacyStats is legacy stats about a player.
type LegacyStats struct {
	NextTokensSeconds int `json:"next_tokens_seconds"`
	QuakecraftTokens  int `json:"quakecraft_tokens"`
	TotalTokens       int `json:"total_tokens"`
	Tokens            int `json:"tokens"`
}

// DuelsStats is duels stats about a player.
type DuelsStats struct {
	TntGamesRookieTitlePrestige     int      `json:"tnt_games_rookie_title_prestige"`
	SumoRookieTitlePrestige         int      `json:"sumo_rookie_title_prestige"`
	SkywarsRookieTitlePrestige      int      `json:"skywars_rookie_title_prestige"`
	ClassicRookieTitlePrestige      int      `json:"classic_rookie_title_prestige"`
	BridgeRookieTitlePrestige       int      `json:"bridge_rookie_title_prestige"`
	BlitzRookieTitlePrestige        int      `json:"blitz_rookie_title_prestige"`
	NoDebuffRookieTitlePrestige     int      `json:"no_debuff_rookie_title_prestige"`
	ComboRookieTitlePrestige        int      `json:"combo_rookie_title_prestige"`
	AllModesRookieTitlePrestige     int      `json:"all_modes_rookie_title_prestige"`
	MegaWallsRookieTitlePrestige    int      `json:"mega_walls_rookie_title_prestige"`
	OpRookieTitlePrestige           int      `json:"op_rookie_title_prestige"`
	BowRookieTitlePrestige          int      `json:"bow_rookie_title_prestige"`
	UhcRookieTitlePrestige          int      `json:"uhc_rookie_title_prestige"`
	DuelsRecentlyPlayed2            string   `json:"duels_recently_played2"`
	Selected2New                    string   `json:"selected_2_new"`
	Selected1New                    string   `json:"selected_1_new"`
	ShowLbOption                    string   `json:"show_lb_option"`
	GamesPlayedDuels                int      `json:"games_played_duels"`
	ChatEnabled                     string   `json:"chat_enabled"`
	MapsWonOn                       []string `json:"maps_won_on"`
	MatchHistory2                   []string `json:"matchHistory2"`
	CurrentWinstreak                int      `json:"current_winstreak"`
	CurrentOpWinstreak              int      `json:"current_op_winstreak"`
	CurrentWinstreakModeOpDoubles   int      `json:"current_winstreak_mode_op_doubles"`
	Coins                           int      `json:"coins"`
	DamageDealt                     int      `json:"damage_dealt"`
	Deaths                          int      `json:"deaths"`
	HealthRegenerated               int      `json:"health_regenerated"`
	Kills                           int      `json:"kills"`
	Losses                          int      `json:"losses"`
	MeleeHits                       int      `json:"melee_hits"`
	MeleeSwings                     int      `json:"melee_swings"`
	OpDoublesDamageDealt            int      `json:"op_doubles_damage_dealt"`
	OpDoublesDeaths                 int      `json:"op_doubles_deaths"`
	OpDoublesHealthRegenerated      int      `json:"op_doubles_health_regenerated"`
	OpDoublesKills                  int      `json:"op_doubles_kills"`
	OpDoublesLosses                 int      `json:"op_doubles_losses"`
	OpDoublesMeleeHits              int      `json:"op_doubles_melee_hits"`
	OpDoublesMeleeSwings            int      `json:"op_doubles_melee_swings"`
	OpDoublesRoundsPlayed           int      `json:"op_doubles_rounds_played"`
	RoundsPlayed                    int      `json:"rounds_played"`
	CurrentWinstreakModeOpDuel      int      `json:"current_winstreak_mode_op_duel"`
	OpDuelDamageDealt               int      `json:"op_duel_damage_dealt"`
	OpDuelDeaths                    int      `json:"op_duel_deaths"`
	OpDuelHealthRegenerated         int      `json:"op_duel_health_regenerated"`
	OpDuelLosses                    int      `json:"op_duel_losses"`
	OpDuelMeleeHits                 int      `json:"op_duel_melee_hits"`
	OpDuelMeleeSwings               int      `json:"op_duel_melee_swings"`
	OpDuelRoundsPlayed              int      `json:"op_duel_rounds_played"`
	BestWinstreakModeOpDuel         int      `json:"best_winstreak_mode_op_duel"`
	BestOpWinstreak                 int      `json:"best_op_winstreak"`
	BestOverallWinstreak            int      `json:"best_overall_winstreak"`
	OpDuelKills                     int      `json:"op_duel_kills"`
	OpDuelWins                      int      `json:"op_duel_wins"`
	Wins                            int      `json:"wins"`
	DuelsWinstreakBestOpDuel        int      `json:"duels_winstreak_best_op_duel"`
	CurrentClassicWinstreak         int      `json:"current_classic_winstreak"`
	CurrentWinstreakModeClassicDuel int      `json:"current_winstreak_mode_classic_duel"`
	BowShots                        int      `json:"bow_shots"`
	ClassicDuelBowShots             int      `json:"classic_duel_bow_shots"`
	ClassicDuelDamageDealt          int      `json:"classic_duel_damage_dealt"`
	ClassicDuelDeaths               int      `json:"classic_duel_deaths"`
	ClassicDuelHealthRegenerated    int      `json:"classic_duel_health_regenerated"`
	ClassicDuelLosses               int      `json:"classic_duel_losses"`
	ClassicDuelMeleeHits            int      `json:"classic_duel_melee_hits"`
	ClassicDuelMeleeSwings          int      `json:"classic_duel_melee_swings"`
	ClassicDuelRoundsPlayed         int      `json:"classic_duel_rounds_played"`
	BestWinstreakModeClassicDuel    int      `json:"best_winstreak_mode_classic_duel"`
	BestClassicWinstreak            int      `json:"best_classic_winstreak"`
	ClassicDuelKills                int      `json:"classic_duel_kills"`
	ClassicDuelWins                 int      `json:"classic_duel_wins"`
	DuelsWinstreakBestClassicDuel   int      `json:"duels_winstreak_best_classic_duel"`
	BowHits                         int      `json:"bow_hits"`
	ClassicDuelBowHits              int      `json:"classic_duel_bow_hits"`
	BestWinstreakModeOpDoubles      int      `json:"best_winstreak_mode_op_doubles"`
	OpDoublesWins                   int      `json:"op_doubles_wins"`
	DuelsChests                     int      `json:"duels_chests"`
	SwDuelsKitNew3                  string   `json:"sw_duels_kit_new3"`
	CurrentSkywarsWinstreak         int      `json:"current_skywars_winstreak"`
	CurrentWinstreakModeSwDuel      int      `json:"current_winstreak_mode_sw_duel"`
	BlocksPlaced                    int      `json:"blocks_placed"`
	SwDuelBlocksPlaced              int      `json:"sw_duel_blocks_placed"`
	SwDuelDeaths                    int      `json:"sw_duel_deaths"`
	SwDuelLosses                    int      `json:"sw_duel_losses"`
	SwDuelRoundsPlayed              int      `json:"sw_duel_rounds_played"`
	BestSkywarsWinstreak            int      `json:"best_skywars_winstreak"`
	BestWinstreakModeSwDuel         int      `json:"best_winstreak_mode_sw_duel"`
	KitWins                         int      `json:"kit_wins"`
	PaladinKitWins                  int      `json:"paladin_kit_wins"`
	SwDuelKitWins                   int      `json:"sw_duel_kit_wins"`
	SwDuelPaladinKitWins            int      `json:"sw_duel_paladin_kit_wins"`
	SwDuelWins                      int      `json:"sw_duel_wins"`
	DuelsWinstreakBestSwDuel        int      `json:"duels_winstreak_best_sw_duel"`
	SwDuelMeleeSwings               int      `json:"sw_duel_melee_swings"`
	SwDuelDamageDealt               int      `json:"sw_duel_damage_dealt"`
	SwDuelHealthRegenerated         int      `json:"sw_duel_health_regenerated"`
	SwDuelMeleeHits                 int      `json:"sw_duel_melee_hits"`
	SumoDuelMeleeHits               int      `json:"sumo_duel_melee_hits"`
	SumoDuelMeleeSwings             int      `json:"sumo_duel_melee_swings"`
	SumoDuelRoundsPlayed            int      `json:"sumo_duel_rounds_played"`
	DuelsWinstreakBestOpDoubles     int      `json:"duels_winstreak_best_op_doubles"`
	LeaderboardPageWinStreak        int      `json:"leaderboardPage_win_streak"`
	BestUhcWinstreak                int      `json:"best_uhc_winstreak"`
	BestWinstreakModeUhcDoubles     int      `json:"best_winstreak_mode_uhc_doubles"`
	CurrentWinstreakModeUhcDoubles  int      `json:"current_winstreak_mode_uhc_doubles"`
	CurrentUhcWinstreak             int      `json:"current_uhc_winstreak"`
	GoldenApplesEaten               int      `json:"golden_apples_eaten"`
	UhcDoublesBlocksPlaced          int      `json:"uhc_doubles_blocks_placed"`
	UhcDoublesBowShots              int      `json:"uhc_doubles_bow_shots"`
	UhcDoublesDamageDealt           int      `json:"uhc_doubles_damage_dealt"`
	UhcDoublesGoldenApplesEaten     int      `json:"uhc_doubles_golden_apples_eaten"`
	UhcDoublesHealthRegenerated     int      `json:"uhc_doubles_health_regenerated"`
	UhcDoublesKills                 int      `json:"uhc_doubles_kills"`
	UhcDoublesMeleeHits             int      `json:"uhc_doubles_melee_hits"`
	UhcDoublesMeleeSwings           int      `json:"uhc_doubles_melee_swings"`
	UhcDoublesRoundsPlayed          int      `json:"uhc_doubles_rounds_played"`
	UhcDoublesWins                  int      `json:"uhc_doubles_wins"`
	DuelsWinstreakBestUhcDoubles    int      `json:"duels_winstreak_best_uhc_doubles"`
	UhcDoublesBowHits               int      `json:"uhc_doubles_bow_hits"`
	UhcDoublesDeaths                int      `json:"uhc_doubles_deaths"`
	UhcDoublesLosses                int      `json:"uhc_doubles_losses"`
	UhcMeetupBowShots               int      `json:"uhc_meetup_bow_shots"`
	UhcMeetupDamageDealt            int      `json:"uhc_meetup_damage_dealt"`
	UhcMeetupHealthRegenerated      int      `json:"uhc_meetup_health_regenerated"`
	UhcMeetupMeleeHits              int      `json:"uhc_meetup_melee_hits"`
	UhcMeetupMeleeSwings            int      `json:"uhc_meetup_melee_swings"`
	UhcMeetupRoundsPlayed           int      `json:"uhc_meetup_rounds_played"`
	UhcMeetupBowHits                int      `json:"uhc_meetup_bow_hits"`
	CurrentWinstreakModeUhcDuel     int      `json:"current_winstreak_mode_uhc_duel"`
	UhcDuelBlocksPlaced             int      `json:"uhc_duel_blocks_placed"`
	UhcDuelDamageDealt              int      `json:"uhc_duel_damage_dealt"`
	UhcDuelDeaths                   int      `json:"uhc_duel_deaths"`
	UhcDuelGoldenApplesEaten        int      `json:"uhc_duel_golden_apples_eaten"`
	UhcDuelHealthRegenerated        int      `json:"uhc_duel_health_regenerated"`
	UhcDuelLosses                   int      `json:"uhc_duel_losses"`
	UhcDuelMeleeHits                int      `json:"uhc_duel_melee_hits"`
	UhcDuelMeleeSwings              int      `json:"uhc_duel_melee_swings"`
	UhcDuelRoundsPlayed             int      `json:"uhc_duel_rounds_played"`
	CurrentWinstreakModeUhcMeetup   int      `json:"current_winstreak_mode_uhc_meetup"`
	UhcMeetupDeaths                 int      `json:"uhc_meetup_deaths"`
	UhcMeetupLosses                 int      `json:"uhc_meetup_losses"`
	UhcDuelBowHits                  int      `json:"uhc_duel_bow_hits"`
	UhcDuelBowShots                 int      `json:"uhc_duel_bow_shots"`
	HealPotsUsed                    int      `json:"heal_pots_used"`
	PotionDuelDamageDealt           int      `json:"potion_duel_damage_dealt"`
	PotionDuelHealPotsUsed          int      `json:"potion_duel_heal_pots_used"`
	PotionDuelHealthRegenerated     int      `json:"potion_duel_health_regenerated"`
	PotionDuelMeleeHits             int      `json:"potion_duel_melee_hits"`
	PotionDuelMeleeSwings           int      `json:"potion_duel_melee_swings"`
	PotionDuelRoundsPlayed          int      `json:"potion_duel_rounds_played"`
	ComboDuelGoldenApplesEaten      int      `json:"combo_duel_golden_apples_eaten"`
	ComboDuelHealthRegenerated      int      `json:"combo_duel_health_regenerated"`
	ComboDuelMeleeHits              int      `json:"combo_duel_melee_hits"`
	ComboDuelMeleeSwings            int      `json:"combo_duel_melee_swings"`
	ComboDuelRoundsPlayed           int      `json:"combo_duel_rounds_played"`
	SwDuelBowHits                   int      `json:"sw_duel_bow_hits"`
	SwDuelBowShots                  int      `json:"sw_duel_bow_shots"`
	EndermanKitWins                 int      `json:"enderman_kit_wins"`
	SwDuelEndermanKitWins           int      `json:"sw_duel_enderman_kit_wins"`
	CurrentWinstreakModeSwDoubles   int      `json:"current_winstreak_mode_sw_doubles"`
	SwDoublesDeaths                 int      `json:"sw_doubles_deaths"`
	SwDoublesHealthRegenerated      int      `json:"sw_doubles_health_regenerated"`
	SwDoublesLosses                 int      `json:"sw_doubles_losses"`
	SwDoublesMeleeSwings            int      `json:"sw_doubles_melee_swings"`
	SwDoublesRoundsPlayed           int      `json:"sw_doubles_rounds_played"`
}

// SkyblockStats is Skyblock stats about a player.
type SkyblockStats struct {
	Profiles struct {
	} `json:"profiles"`
}

// BuildBattleStats is build battle stats about a player.
type BuildBattleStats struct {
	Coins          int `json:"coins"`
	GamesPlayed    int `json:"games_played"`
	MonthlyCoinsB  int `json:"monthly_coins_b"`
	Score          int `json:"score"`
	SoloMostPoints int `json:"solo_most_points"`
	TotalVotes     int `json:"total_votes"`
	WeeklyCoinsB   int `json:"weekly_coins_b"`
	Wins           int `json:"wins"`
	WinsSoloNormal int `json:"wins_solo_normal"`
}

// MurderMysteryStats is Murder Mystery stats about a player.
type MurderMysteryStats struct {
	MmChests int `json:"mm_chests"`
}

// PitStats is The Pit stats about a player.
type PitStats struct {
	PitStatsPtl struct {
		ArrowsFired           int `json:"arrows_fired"`
		Assists               int `json:"assists"`
		BowDamageReceived     int `json:"bow_damage_received"`
		CashEarned            int `json:"cash_earned"`
		DamageDealt           int `json:"damage_dealt"`
		DamageReceived        int `json:"damage_received"`
		GappleEaten           int `json:"gapple_eaten"`
		Joins                 int `json:"joins"`
		JumpedIntoPit         int `json:"jumped_into_pit"`
		Kills                 int `json:"kills"`
		LeftClicks            int `json:"left_clicks"`
		MaxStreak             int `json:"max_streak"`
		MeleeDamageDealt      int `json:"melee_damage_dealt"`
		MeleeDamageReceived   int `json:"melee_damage_received"`
		SwordHits             int `json:"sword_hits"`
		ArrowHits             int `json:"arrow_hits"`
		BowDamageDealt        int `json:"bow_damage_dealt"`
		Deaths                int `json:"deaths"`
		LaunchedByLaunchers   int `json:"launched_by_launchers"`
		PlaytimeMinutes       int `json:"playtime_minutes"`
		ChatMessages          int `json:"chat_messages"`
		DiamondItemsPurchased int `json:"diamond_items_purchased"`
		GheadEaten            int `json:"ghead_eaten"`
		ContractsStarted      int `json:"contracts_started"`
	} `json:"pit_stats_ptl"`
	Profile struct {
		OutgoingOffers         []interface{} `json:"outgoing_offers"`
		LastMidfightDisconnect int64         `json:"last_midfight_disconnect"`
		ContractChoices        interface{}   `json:"contract_choices"`
		LeaderboardStats       struct {
			PitTdmRedKills2019Winter int `json:"Pit_tdm_red_kills_2019_winter"`
			PitCakeEaten2019Winter   int `json:"Pit_cake_eaten_2019_winter"`
			PitKotlTime2019Winter    int `json:"Pit_kotl_time_2019_winter"`
			PitKotlGold2019Winter    int `json:"Pit_kotl_gold_2019_winter"`
		} `json:"leaderboard_stats"`
		LastSave  int64 `json:"last_save"`
		KingQuest struct {
			Kills int `json:"kills"`
		} `json:"king_quest"`
		InvArmor struct {
			Type int   `json:"type"`
			Data []int `json:"data"`
		} `json:"inv_armor"`
		SelectedPerk1   string        `json:"selected_perk_1"`
		SelectedPerk0   string        `json:"selected_perk_0"`
		LastContract    int64         `json:"last_contract"`
		LoginMessages   []interface{} `json:"login_messages"`
		HotbarFavorites []int         `json:"hotbar_favorites"`
		RecentKills     []struct {
			Victim    string `json:"victim"`
			Timestamp int64  `json:"timestamp"`
		} `json:"recent_kills"`
		Xp          int `json:"xp"`
		InvContents struct {
			Type int   `json:"type"`
			Data []int `json:"data"`
		} `json:"inv_contents"`
		ZeroPointThreeGoldTransfer bool `json:"zero_point_three_gold_transfer"`
		DeathRecaps                struct {
			Type int   `json:"type"`
			Data []int `json:"data"`
		} `json:"death_recaps"`
		EndedContracts []struct {
			Difficulty   string `json:"difficulty"`
			GoldReward   int    `json:"gold_reward"`
			Requirements struct {
				Ingots int `json:"ingots"`
			} `json:"requirements"`
			Progress struct {
				Ingots int `json:"ingots"`
			} `json:"progress"`
			ChunkOfVilesReward int    `json:"chunk_of_viles_reward"`
			CompletionDate     int    `json:"completion_date"`
			RemainingTicks     int    `json:"remaining_ticks"`
			Key                string `json:"key"`
		} `json:"ended_contracts"`
		Bounties []interface{} `json:"bounties"`
		Unlocks  []struct {
			Tier        int    `json:"tier"`
			AcquireDate int64  `json:"acquireDate"`
			Key         string `json:"key"`
		} `json:"unlocks"`
		Cash                float64 `json:"cash"`
		CashDuringPrestige0 float64 `json:"cash_during_prestige_0"`
	} `json:"profile"`
}

// MCGOStats is MCGO stats about a player.
type MCGOStats struct {
	LastTourneyAd int64 `json:"lastTourneyAd"`
}
