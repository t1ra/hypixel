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
	"fmt"
	. "hypixel/structs"
	"net/url"

	"github.com/google/uuid"
)

var (
	// Endpoints
	banStats = APIBaseURL + "watchdogstats"
	// resources/*
	resourcesAchievements        = APIBaseURL + "resources/achievements"
	resourcesChallenges          = APIBaseURL + "resources/challenges"
	resourcesQuests              = APIBaseURL + "resources/quests"
	resourcesGuildsAchievements  = APIBaseURL + "resources/guilds/achievements"
	resourcesGuildsPermissions   = APIBaseURL + "resources/guilds/permissions"
	resourcesSkyblockCollections = APIBaseURL + "resources/skyblock/collections"
	resourcesSkyblockSkills      = APIBaseURL + "resources/skyblock/skills"
	playerCount                  = APIBaseURL + "playerCount"
	player                       = APIBaseURL + "player"
	leaderboards                 = APIBaseURL + "leaderboards"
	key                          = APIBaseURL + "key"
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

	stats, err := session.APIRequest(banStats)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(stats, &result)

	return result, nil
}

// Achievements returns all available achievements for all gamemodes.
func (session *Hypixel) Achievements() (Achievements, error) {
	var result Achievements

	achievements, err := session.APIRequest(resourcesAchievements)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(achievements, &result)
	return result, nil
}

// Challenges returns all the available challenges for all gamemodes.
func (session *Hypixel) Challenges() (Challenges, error) {
	var result Challenges

	challenges, err := session.APIRequest(resourcesChallenges)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(challenges, &result)

	return result, nil
}

// Quests returns all the available quests for all gamemodes.
func (session *Hypixel) Quests() (Quests, error) {
	var result Quests

	quests, err := session.APIRequest(resourcesQuests)
	if err != nil {
		return result, err
	}

	fmt.Println(string(quests))

	err = json.Unmarshal(quests, &result)

	return result, nil
}

// GuildAchievements returns all the available guild achievements.
func (session *Hypixel) GuildAchievements() (GuildAchievements, error) {
	var result GuildAchievements

	achievements, err := session.APIRequest(resourcesGuildsAchievements)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(achievements, &result)

	return result, nil
}

// GuildPermissions returns all the available guild permissions.
func (session *Hypixel) GuildPermissions() (GuildPermissions, error) {
	var result GuildPermissions

	permissions, err := session.APIRequest(resourcesGuildsPermissions)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(permissions, &result)

	return result, nil
}

// SkyblockCollections returns all the available skyblock collections.
func (session *Hypixel) SkyblockCollections() (SkyblockCollections, error) {
	var result SkyblockCollections

	collections, err := session.APIRequest(resourcesSkyblockCollections)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(collections, &result)

	return result, nil
}

// SkyblockSkills returns all the available skyblock skills.
func (session *Hypixel) SkyblockSkills() (SkyblockSkills, error) {
	var result SkyblockSkills

	skills, err := session.APIRequest(resourcesSkyblockSkills)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(skills, &result)

	return result, nil
}

// PlayerCount returns the current amount of players online.
func (session *Hypixel) PlayerCount() (int, error) {
	var result PlayerCount

	playerCount, err := session.APIRequest(playerCount)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(playerCount, &result)
	if err != nil {
		return 0, err
	}

	return result.PlayerCount, nil
}

// Player returns information about a player.
func (session *Hypixel) Player(uuid string) (Player, error) {
	type PlayerContainer struct {
		Player Player
	}

	var result PlayerContainer

	player, err := session.APIRequest(player, "uuid", uuid)
	if err != nil {
		return result.Player, err
	}

	err = json.Unmarshal(player, &result)

	return result.Player, err
}

// Leaderboards returns information about the current Hypixel leaderboards.
func (session *Hypixel) Leaderboards() (Leaderboards, error) {
	var result Leaderboards

	leaderboards, err := session.APIRequest(leaderboards)
	if err != nil {
		return leaderboards, err
	}

	err = json.Unmarshal(leaderboards, &result)

	return result, err
}

// Key returns information about the current active key.
func (session *Hypixel) Key() (Key, error) {
	var result Key

	keyData, err := session.APIRequest(key)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(keyData, &result)

	return result, err
}

// Guild returns information about a guild.
func (session *Hypixel) Guild(guild string) (Guild, error) {
	var result Guild

	var guildData interface{}
	_, err := uuid.Parse(guild)
	if err == nil { // It's a player UUID
		guildData, err = session.APIRequest(guild, "player", guild)
	} else { // It's a guild name
		guildData, err = session.APIRequest(guild, "name", url.QueryEscape(guild))
	}

	err = json.Unmarshal(guildData, &result)

	return result, err
}

// Stats returns statistics about a guild.
func (session *Hypixel) Stats() (Stats, error) {
	var result Stats

	statsData := session.APIRequest(stats)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(statsData, &result)

	return result, err
}

// Friends returns frend information about a player.
func (session *Hypixel) Friends() (Friends, error) {
	var result Friends

	friendData := session.APIRequest(friends)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(friendData, &result)

	return result, err
}

// FindGuild returns the ID of a requested guild.
func (session *Hypixel) FindGuild(guild string) (string, error) {
	var result string

	findGuild := struct {
		Guild string `json:"guild"`
	}

	_, err := uuid.Parse(guild)
	if err == nil { // It's a guild UUID
		findGuild, err = session.APIRequest(findguild, "byUuid", guild)
	} else { // It's a guild name
		findGuild, err = session.APIRequest(findguild, "byName", url.QueryEscape(guild))
	}

	err = json.Unmarshal(findGuild, &result)

	if err != nil {
		return "", err
	}

	return findGuild.Guild, err
}

// Boosters returns a list of active boosters.
func (session *Hypixel) Boosters() (Boosters, error) {
	var result Boosters

	boosterData := session.APIRequest(boosters)
	if err != nil {
		return result, err
	}

	err = json.Unmarshal(boosterData, &result)

	return result, err
}