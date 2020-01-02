package hypixel

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

import (
	"encoding/json"
	"net/url"

	. "github.com/t1ra/hypixel/structs"

	"github.com/google/uuid"
)

var (
	// Endpoints
	banStats     = APIBaseURL + "watchdogstats"
	playerCount  = APIBaseURL + "playerCount"
	player       = APIBaseURL + "player"
	leaderboards = APIBaseURL + "leaderboards"
	key          = APIBaseURL + "key"
	gameStats    = APIBaseURL + "gameCounts"
	findGuild    = APIBaseURL + "findGuild"
	guild        = APIBaseURL + "guild"
	friends      = APIBaseURL + "friends"
	boosters     = APIBaseURL + "boosters"
	// resources/* endpoints
	resourcesAchievements        = APIBaseURL + "resources/achievements"
	resourcesChallenges          = APIBaseURL + "resources/challenges"
	resourcesQuests              = APIBaseURL + "resources/quests"
	resourcesGuildsAchievements  = APIBaseURL + "resources/guilds/achievements"
	resourcesGuildsPermissions   = APIBaseURL + "resources/guilds/permissions"
	resourcesSkyblockCollections = APIBaseURL + "resources/skyblock/collections"
	resourcesSkyblockSkills      = APIBaseURL + "resources/skyblock/skills"
)

/*
 * Each endpoint has its own struct that holds the values that the associated function will return.
 * The struct should have the same name as the function, and should not marshall the `success`
 * field. Structs should have the method's documentation linked, the function should not.
 *
 * These get pretty boilerplate; most of the work is done in `structs`.
 */

// BanStats returns statistics about Watchdog and staff bans.
func (session *Hypixel) BanStats() (BanStats, error) {
	var result BanStats

	stats, err := session.apiRequest(banStats)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(stats, &result)

	return result, err
}

// Achievements returns all available achievements for all gamemodes.
func (session *Hypixel) Achievements() (Achievements, error) {
	var result Achievements

	achievements, err := session.apiRequest(resourcesAchievements)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(achievements, &result)

	return result, err
}

// Challenges returns all the available challenges for all gamemodes.
func (session *Hypixel) Challenges() (Challenges, error) {
	var result Challenges

	challenges, err := session.apiRequest(resourcesChallenges)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(challenges, &result)

	return result, err
}

// Quests returns all the available quests for all gamemodes.
func (session *Hypixel) Quests() (Quests, error) {
	var result Quests

	quests, err := session.apiRequest(resourcesQuests)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(quests, &result)

	return result, err
}

// GuildAchievements returns all the available guild achievements.
func (session *Hypixel) GuildAchievements() (GuildAchievements, error) {
	var result GuildAchievements

	achievements, err := session.apiRequest(resourcesGuildsAchievements)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(achievements, &result)

	return result, err
}

// GuildPermissions returns all the available guild permissions.
func (session *Hypixel) GuildPermissions() ([]GuildPermission, error) {
	var guildPermissionsContainer struct {
		LastUpdated int64 `json:"lastUpdated"`
		Permissions []struct {
			Permission GuildPermission `json:"en_us"`
		} `json:"permissions"`
	}

	var guildPermissions []GuildPermission

	permissions, err := session.apiRequest(resourcesGuildsPermissions)
	if err != nil {
		return guildPermissions, err
	}

	err = json.Unmarshal(permissions, &guildPermissionsContainer)

	for _, permission := range guildPermissionsContainer.Permissions {
		guildPermissions = append(guildPermissions, permission.Permission)
	}

	return guildPermissions, err
}

// SkyblockCollections returns all the available skyblock collections.
func (session *Hypixel) SkyblockCollections() (map[string]SkyblockCollection, error) {
	// skyblockCollections is a container for all of the collections in Skyblock.
	// Instead of returning a needless container, SkyblockCollections() returns
	// skyblockCollections.Collections directly. The other data here isn't really
	// necessary.
	var skyblockCollections struct {
		LastUpdated int                           `json:"lastUpdated"`
		Version     string                        `json:"version"`
		Collections map[string]SkyblockCollection `json:"collections"`
	}

	collections, err := session.apiRequest(resourcesSkyblockCollections)
	if err != nil {
		return skyblockCollections.Collections, err
	}

	err = json.Unmarshal(collections, &skyblockCollections)

	return skyblockCollections.Collections, err
}

// SkyblockSkills returns all the available skyblock skills.
func (session *Hypixel) SkyblockSkills() (SkyblockSkills, error) {
	var result SkyblockSkills

	skills, err := session.apiRequest(resourcesSkyblockSkills)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(skills, &result)

	return result, err
}

// PlayerCount returns the current amount of players online.
func (session *Hypixel) PlayerCount() (int, error) {
	var result PlayerCount

	playerCount, err := session.apiRequest(playerCount)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(playerCount, &result)
	if err != nil {
		return 0, err
	}

	return result.PlayerCount, err
}

// Player returns information about a player.
func (session *Hypixel) Player(uuid string) (Player, error) {
	type PlayerContainer struct {
		Player Player
	}

	var result PlayerContainer

	player, err := session.apiRequest(player, "uuid", uuid)
	if err != nil {
		return result.Player, err
	}

	err = json.Unmarshal(player, &result)

	return result.Player, err
}

// Leaderboards returns information about the current Hypixel leaderboards.
func (session *Hypixel) Leaderboards() (Leaderboards, error) {
	var leaderboardsContainer struct {
		Leaderboards Leaderboards `json:"leaderboards"`
	}

	leaderboardData, err := session.apiRequest(leaderboards)
	if err != nil {
		return leaderboardsContainer.Leaderboards, err
	}

	err = json.Unmarshal(leaderboardData, &leaderboardsContainer)

	return leaderboardsContainer.Leaderboards, err
}

// Key returns information about the current active key.
// If no arguments are provided, it looks up the key that's set in the
// Hypixel struct. If an argument is provided, it looks up that key instead.
func (session *Hypixel) Key(customKey ...string) (Key, error) {
	var keyDataContainer struct {
		Key Key `json:"record"`
	}

	var keyData []byte
	var err error

	if len(customKey) == 0 {
		keyData, err = session.apiRequest(key)
	} else {
		keyData, err = session.apiRequestCustomKey(key, customKey[0])
	}

	if err != nil {
		return keyDataContainer.Key, err
	}

	err = json.Unmarshal(keyData, &keyDataContainer)

	return keyDataContainer.Key, err
}

// FindGuild returns the ID of a requested guild.
func (session *Hypixel) FindGuild(guild string) (string, error) {
	var findGuildContainer struct {
		Guild string `json:"guild"`
	}

	var findGuildBytes []byte

	_, err := uuid.Parse(guild)
	if err == nil { // It's a guild UUID
		findGuildBytes, err = session.apiRequest(findGuild, "byUuid", guild)
	} else { // It's a guild name
		findGuildBytes, err = session.apiRequest(findGuild, "byName", url.QueryEscape(guild))
	}

	if err != nil {
		return "", err
	}

	err = json.Unmarshal(findGuildBytes, &findGuildContainer)

	return findGuildContainer.Guild, err
}

// Guild returns information about a guild.
func (session *Hypixel) Guild(userGuild string) (Guild, error) {
	var result struct {
		Guild Guild `json:"guild"`
	}

	var guildData []byte
	_, err := uuid.Parse(userGuild)

	if err == nil { // It's a player UUID
		guildData, err = session.apiRequest(guild, "player", userGuild)
	} else { // It's a guild name
		guildData, err = session.apiRequest(guild, "name", url.QueryEscape(userGuild))
	}

	if err != nil {
		return result.Guild, err
	}

	err = json.Unmarshal(guildData, &result)

	return result.Guild, err
}

// GameStats returns the total players online, total players for every minigame,
// and total players for every mode of that minigame
// e.g. (Solo/Doubles/Mega Skywars).
func (session *Hypixel) GameStats() (Stats, error) {
	var result Stats

	stats, err := session.apiRequest(gameStats)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(stats, &result)

	return result, err
}

// Friends returns frend information about a player.
func (session *Hypixel) Friends(uuid string) (Friends, error) {
	var result Friends

	friendData, err := session.apiRequest(friends, "uuid", uuid)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(friendData, &result)

	return result, err
}

// Boosters returns a list of active boosters.
func (session *Hypixel) Boosters() (Boosters, error) {
	var result Boosters

	boosterData, err := session.apiRequest(boosters)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(boosterData, &result)

	return result, err
}
