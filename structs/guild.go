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
 *                           Guild
 */

type Guild struct {
	ID        string `json:"_id"`
	Name      string `json:"name"`
	NameLower string `json:"name_lower"`
	Coins     int    `json:"coins"`
	CoinsEver int    `json:"coinsEver"`
	Created   int64  `json:"created"`
	Members   []struct {
		UUID               string      `json:"uuid"`
		Rank               string      `json:"rank"`
		Joined             int64       `json:"joined"`
		QuestParticipation int         `json:"questParticipation,omitempty"`
		ExpHistory         interface{} `json:"expHistory"`
	} `json:"members"`
	Ranks []struct {
		Name     string `json:"name"`
		Default  bool   `json:"default,omitempty"`
		Priority int    `json:"priority"`
		Created  int64  `json:"created"`
	} `json:"ranks"`
	Achievements struct {
		OnlinePlayers   int `json:"ONLINE_PLAYERS"`
		Winners         int `json:"WINNERS"`
		ExperienceKings int `json:"EXPERIENCE_KINGS"`
	} `json:"achievements"`
	Exp                int    `json:"exp"`
	Tag                string `json:"tag"`
	GuildExpByGameType struct {
		Arcade        int `json:"ARCADE"`
		UHC           int `json:"UHC"`
		Skywars       int `json:"SKYWARS"`
		Prototype     int `json:"PROTOTYPE"`
		Bedwars       int `json:"BEDWARS"`
		MurderMystery int `json:"MURDER_MYSTERY"`
		BuildBattle   int `json:"BUILD_BATTLE"`
		Duels         int `json:"DUELS"`
	} `json:"guildExpByGameType"`
}
