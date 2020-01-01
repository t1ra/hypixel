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
 *                           Stats
 */

type Stats struct {
	GameStats struct {
		BUILDBATTLE struct {
			Modes struct {
				BUILDBATTLECHRISTMASNEWTEAMS int `json:"BUILD_BATTLE_CHRISTMAS_NEW_TEAMS"`
				BUILDBATTLESOLONORMALLATEST  int `json:"BUILD_BATTLE_SOLO_NORMAL_LATEST"`
				BUILDBATTLEGUESSTHEBUILD     int `json:"BUILD_BATTLE_GUESS_THE_BUILD"`
				BUILDBATTLETEAMSNORMAL       int `json:"BUILD_BATTLE_TEAMS_NORMAL"`
				BUILDBATTLESOLONORMAL        int `json:"BUILD_BATTLE_SOLO_NORMAL"`
				BUILDBATTLESOLOPRO           int `json:"BUILD_BATTLE_SOLO_PRO"`
				BUILDBATTLECHRISTMASNEWSOLO  int `json:"BUILD_BATTLE_CHRISTMAS_NEW_SOLO"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"BUILD_BATTLE"`
		UHC struct {
			Modes struct {
				TEAMS int `json:"TEAMS"`
				SOLO  int `json:"SOLO"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"UHC"`
		LEGACY struct {
			Modes struct {
				WALLS       int `json:"WALLS"`
				QUAKECRAFT  int `json:"QUAKECRAFT"`
				PAINTBALL   int `json:"PAINTBALL"`
				VAMPIREZ    int `json:"VAMPIREZ"`
				ARENA       int `json:"ARENA"`
				GINGERBREAD int `json:"GINGERBREAD"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"LEGACY"`
		SKYBLOCK struct {
			Modes struct {
				Combat3   int `json:"combat_3"`
				Hub       int `json:"hub"`
				Farming1  int `json:"farming_1"`
				Farming2  int `json:"farming_2"`
				Foraging1 int `json:"foraging_1"`
				Dynamic   int `json:"dynamic"`
				Combat2   int `json:"combat_2"`
				Mining1   int `json:"mining_1"`
				Combat1   int `json:"combat_1"`
				Mining2   int `json:"mining_2"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"SKYBLOCK"`
		HOUSING struct {
			Players int `json:"players"`
		} `json:"HOUSING"`
		MCGO struct {
			Modes struct {
				Normal      int `json:"normal"`
				NormalParty int `json:"normal_party"`
				Deathmatch  int `json:"deathmatch"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"MCGO"`
		BATTLEGROUND struct {
			Modes struct {
				CtfMini        int `json:"ctf_mini"`
				TeamDeathmatch int `json:"team_deathmatch"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"BATTLEGROUND"`
		SURVIVALGAMES struct {
			Modes struct {
				SoloNormal  int `json:"solo_normal"`
				TeamsNormal int `json:"teams_normal"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"SURVIVAL_GAMES"`
		MURDERMYSTERY struct {
			Modes struct {
				MURDERDOUBLEUP  int `json:"MURDER_DOUBLE_UP"`
				MURDERINFECTION int `json:"MURDER_INFECTION"`
				MURDERASSASSINS int `json:"MURDER_ASSASSINS"`
				MURDERCLASSIC   int `json:"MURDER_CLASSIC"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"MURDER_MYSTERY"`
		ARCADE struct {
			Modes struct {
				PARTY                  int `json:"PARTY"`
				HOLEINTHEWALL          int `json:"HOLE_IN_THE_WALL"`
				DEFENDER               int `json:"DEFENDER"`
				MINIWALLS              int `json:"MINI_WALLS"`
				ZOMBIESBADBLOOD        int `json:"ZOMBIES_BAD_BLOOD"`
				SANTASIMULATOR         int `json:"SANTA_SIMULATOR"`
				HIDEANDSEEKPARTYPOOPER int `json:"HIDE_AND_SEEK_PARTY_POOPER"`
				DAYONE                 int `json:"DAYONE"`
				DRAWTHEIRTHING         int `json:"DRAW_THEIR_THING"`
				ZOMBIESALIENARCADIUM   int `json:"ZOMBIES_ALIEN_ARCADIUM"`
				ONEINTHEQUIVER         int `json:"ONEINTHEQUIVER"`
				SOCCER                 int `json:"SOCCER"`
				ENDER                  int `json:"ENDER"`
				THROWOUT               int `json:"THROW_OUT"`
				GRINCH                 int `json:"GRINCH"`
				STARWARS               int `json:"STARWARS"`
				SANTASAYS              int `json:"SANTA_SAYS"`
				DRAGONWARS2            int `json:"DRAGONWARS2"`
				ZOMBIESDEADEND         int `json:"ZOMBIES_DEAD_END"`
				FARMHUNT               int `json:"FARM_HUNT"`
				HIDEANDSEEKPROPHUNT    int `json:"HIDE_AND_SEEK_PROP_HUNT"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"ARCADE"`
		TOURNAMENTLOBBY struct {
			Players int `json:"players"`
		} `json:"TOURNAMENT_LOBBY"`
		TNTGAMES struct {
			Modes struct {
				PVPRUN    int `json:"PVPRUN"`
				TNTAG     int `json:"TNTAG"`
				TNTRUN    int `json:"TNTRUN"`
				CAPTURE   int `json:"CAPTURE"`
				BOWSPLEEF int `json:"BOWSPLEEF"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"TNTGAMES"`
		SKYWARS struct {
			Modes struct {
				RankedNormal         int `json:"ranked_normal"`
				SoloInsaneTntMadness int `json:"solo_insane_tnt_madness"`
				SoloInsaneLucky      int `json:"solo_insane_lucky"`
				SoloInsaneSlime      int `json:"solo_insane_slime"`
				TeamsInsaneSlime     int `json:"teams_insane_slime"`
				SoloInsaneRush       int `json:"solo_insane_rush"`
				TeamsInsane          int `json:"teams_insane"`
				SoloNormal           int `json:"solo_normal"`
				SoloInsane           int `json:"solo_insane"`
				TeamsNormal          int `json:"teams_normal"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"SKYWARS"`
		PROTOTYPE struct {
			Modes struct {
				TOWERWARSSOLO      int `json:"TOWERWARS_SOLO"`
				PVPCTW             int `json:"PVP_CTW"`
				TOWERWARSTEAMOFTWO int `json:"TOWERWARS_TEAM_OF_TWO"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"PROTOTYPE"`
		WALLS3 struct {
			Modes struct {
				Standard int `json:"standard"`
				FaceOff  int `json:"face_off"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"WALLS3"`
		BEDWARS struct {
			Modes struct {
				BEDWARSTWOFOUR       int `json:"BEDWARS_TWO_FOUR"`
				BEDWARSFOURTHREE     int `json:"BEDWARS_FOUR_THREE"`
				BEDWARSEIGHTONE      int `json:"BEDWARS_EIGHT_ONE"`
				BEDWARSCASTLE        int `json:"BEDWARS_CASTLE"`
				BEDWARSEIGHTTWORUSH  int `json:"BEDWARS_EIGHT_TWO_RUSH"`
				BEDWARSEIGHTTWOARMED int `json:"BEDWARS_EIGHT_TWO_ARMED"`
				BEDWARSEIGHTTWO      int `json:"BEDWARS_EIGHT_TWO"`
				BEDWARSFOURFOUR      int `json:"BEDWARS_FOUR_FOUR"`
				BEDWARSEIGHTTWOLUCKY int `json:"BEDWARS_EIGHT_TWO_LUCKY"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"BEDWARS"`
		LIMBO struct {
			Players int `json:"players"`
		} `json:"LIMBO"`
		QUEUE struct {
			Players int `json:"players"`
		} `json:"QUEUE"`
		MAINLOBBY struct {
			Players int `json:"players"`
		} `json:"MAIN_LOBBY"`
		SUPERSMASH struct {
			Modes struct {
				OneV1Normal int `json:"1v1_normal"`
				SoloNormal  int `json:"solo_normal"`
				TwoV2Normal int `json:"2v2_normal"`
				TeamsNormal int `json:"teams_normal"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"SUPER_SMASH"`
		TRUECOMBAT struct {
			Modes struct {
				Solo      int `json:"solo"`
				Team      int `json:"team"`
				SoloChaos int `json:"solo_chaos"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"TRUE_COMBAT"`
		PIT struct {
			Modes struct {
				PIT int `json:"PIT"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"PIT"`
		SPEEDUHC struct {
			Modes struct {
				TeamNormal int `json:"team_normal"`
				SoloNormal int `json:"solo_normal"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"SPEED_UHC"`
		DUELS struct {
			Modes struct {
				DUELSBOWSPLEEFDUEL int `json:"DUELS_BOWSPLEEF_DUEL"`
				DUELSBRIDGE2V2V2V2 int `json:"DUELS_BRIDGE_2V2V2V2"`
				DUELSBOWDUEL       int `json:"DUELS_BOW_DUEL"`
				DUELSMWDUEL        int `json:"DUELS_MW_DUEL"`
				DUELSUHCFOUR       int `json:"DUELS_UHC_FOUR"`
				DUELSUHCMEETUP     int `json:"DUELS_UHC_MEETUP"`
				DUELSBRIDGEDOUBLES int `json:"DUELS_BRIDGE_DOUBLES"`
				DUELSSWDOUBLES     int `json:"DUELS_SW_DOUBLES"`
				DUELSUHCDOUBLES    int `json:"DUELS_UHC_DOUBLES"`
				DUELSBRIDGEFOUR    int `json:"DUELS_BRIDGE_FOUR"`
				DUELSBRIDGE3V3V3V3 int `json:"DUELS_BRIDGE_3V3V3V3"`
				DUELSSUMODUEL      int `json:"DUELS_SUMO_DUEL"`
				DUELSUHCDUEL       int `json:"DUELS_UHC_DUEL"`
				DUELSOPDOUBLES     int `json:"DUELS_OP_DOUBLES"`
				DUELSOPDUEL        int `json:"DUELS_OP_DUEL"`
				DUELSMWDOUBLES     int `json:"DUELS_MW_DOUBLES"`
				DUELSBLITZDUEL     int `json:"DUELS_BLITZ_DUEL"`
				DUELSCLASSICDUEL   int `json:"DUELS_CLASSIC_DUEL"`
				DUELSPOTIONDUEL    int `json:"DUELS_POTION_DUEL"`
				DUELSCOMBODUEL     int `json:"DUELS_COMBO_DUEL"`
				DUELSBRIDGEDUEL    int `json:"DUELS_BRIDGE_DUEL"`
				DUELSSWDUEL        int `json:"DUELS_SW_DUEL"`
			} `json:"modes"`
			Players int `json:"players"`
		} `json:"DUELS"`
	} `json:"games"`
	PlayerCount int `json:"playerCount"`
}
