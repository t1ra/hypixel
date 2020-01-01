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
 *                           BanStats
 */

// BanStats stores data returned from Hypixel.BanStats()
// https://github.com/HypixelDev/PublicAPI/blob/master/Documentation/methods/watchdogstats.md
type BanStats struct {
	// Watchdog
	WatchdogTotal      int `json:"watchdog_total"`
	WatchdogLastDay    int `json:"watchdog_rollingDaily"`
	WatchdogLastMinute int `json:"watchdog_lastMinute"`
	// Staff
	StaffTotal   int `json:"staff_total"`
	StaffLastDay int `json:"staff_rollingDaily"`
}

/*
 *                           Achievements
 */

// Achievements stores data for the achievements subresource in the resources endpoint.
// This endpoint's JSON response is complicated so it requires a bunch of structs.
// https://github.com/HypixelDev/PublicAPI/blob/master/Documentation/methods/resources.md
type Achievements struct {
	LastUpdated int64                       `json:"lastUpdated"`
	Modes       map[string]AchievementsMode `json:"achievements"`
}

// AchievementsMode structs are created for every game mode, it contains sub-structs (seperate
// ones for both one time and tiered achievements).
type AchievementsMode struct {
	TotalPoints       int `json:"total_points"`
	TotalLegacyPoints int `json:"total_legacy_points"`
	// Achievements
	OneTime map[string]OneTimeAchievement `json:"one_time"`
	Tiered  map[string]TieredAchievement  `json:"tiered"`
}

// OneTimeAchievement is created for one time achievements (i.e. achievements with one tier)
type OneTimeAchievement struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Points      int    `json:"points"`
}

// TieredAchievement is created for achievements that have multiple tiers, which are stored in a
// seperate Tier struct.
type TieredAchievement struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Tiers       []Tier `json:"tiers"`
}

// Tier is a tier of a TieredAchievement. It contains the tier level (usually 0 - 5), the amount
// of points the tier awards, and amount, which I bet means something.
type Tier struct {
	Tier   int `json:"tier"`
	Points int `json:"points"`
	Amount int `json:"amount"`
}

/*
 *                           Challenges
 */

// Challenges stores data for the challenges subresource in the resources endpoint.
// This endpoint's JSON response is complicated so it requires a bunch of structs.
// https://github.com/HypixelDev/PublicAPI/blob/master/Documentation/methods/resources.md
type Challenges struct {
	LastUpdated int64                          `json:"lastUpdated"`
	Modes       map[string]ChallengesChallenge `json:"challenges"`
}

// ChallengesChallenge stores data about a challenge. Rewards requires its own struct because
// for some reason, the API developers decided that even though no challenge has more than
// one reward, they should be in an array of objects. Every challenge reward is at index 0.
// I like to think this is here for future-proofing, and its easy to just demarshal it like this,
// so we'll keep it.
type ChallengesChallenge struct {
	ID      string                  `json:"id"`
	Name    string                  `json:"name"`
	Rewards map[int]ChallengeReward `json:"rewards"`
}

// ChallengeReward stores a challenge reward. It contains the type, which defines how the gained
// experience (presumably stored in `amount`) affects your total experience, e.g.
// "MultipliedExperienceReward".
type ChallengeReward struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

/*
 *                           Quests
 */

// Quests stores data for the quests subresource in the resources endpoint.
// This endpoint's JSON response is complicated so it requires a bunch of structs.
// https://github.com/HypixelDev/PublicAPI/blob/master/Documentation/methods/resources.md
type Quests struct {
	LastUpdated int64                    `json:"lastUpdated"`
	Modes       map[string][]QuestsQuest `json:"quests"`
}

// QuestsQuest stores data about a quest. It suffers from the same rewards issue as Challenges.
type QuestsQuest struct {
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Rewards      []QuestReward      `json:"rewards"`
	Requirements []QuestRequirement `json:"requirements"`
	Description  string             `json:"description"`
}

// QuestReward stores a quest reward.
type QuestReward struct {
	Type   string `json:"type"`
	Amount int    `json:"amount"`
}

// QuestRequirement simply contains the type of requirement for a quest.
type QuestRequirement struct {
	Type string `json:"type"`
}

/*
 *                           Guild Achievements
 */

// GuildAchievements contains all achievements obtainable in guilds.
type GuildAchievements struct {
	LastUpdated int64 `json:"lastUpdated"`
	// There is a "OneTime" field in the JSON response, but its empty.
	Tiered map[string]TieredGuildAchievement
}

// TieredGuildAchievement contains a guild achievement with more than one level, or tier.
type TieredGuildAchievement struct {
	Name        string                        `json:"name"`
	Description string                        `json:"description"`
	Tiers       map[int]GuildAchievementTiers `json:"tiers"`
}

// GuildAchievementTiers stores a tier of a TieredGuildAchievement.
type GuildAchievementTiers struct {
	Tier   int `json:"tier"`
	Amount int `json:"amount"`
}

/*
 *                           Guild Permissions
 */

// GuildPermissions contains all of the permissions a guild member can hold.
type GuildPermissions struct {
	LastUpdated int64                   `json:"lastUpdated"`
	Permissions map[int]GuildPermission `json:"permissions"`
}

// GuildPermission actually just contains an array of objects. The objects' names are the locale
// of the permission's text.
type GuildPermission struct {
	Locales map[string]GuildPermissionLocalised
}

// GuildPermissionLocalised contains guild permission information.
type GuildPermissionLocalised struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Items       map[string]string `json:"item"`
}

/*
 *                           Skyblock Collections
 */

// SkyblockCollections contains all of the collections in Skyblock.
type SkyblockCollections struct {
	LastUpdated int                           `json:"lastUpdated"`
	Version     string                        `json:"version"`
	Collections map[string]SkyblockCollection `json:"collections"`
}

// SkyblockCollection stores a Skyblock collection's name and the item its tied to.
type SkyblockCollection struct {
	Name  string                            `json:"name"`
	Items map[string]SkyblockCollectionItem `json:"items"`
}

// SkyblockCollectionItem defines information about a SkyblockCollection's item
type SkyblockCollectionItem struct {
	Name    string                   `json:"name"`
	MaxTier int                      `json:"maxTiers"`
	Tiers   []SkyblockCollectionTier `json:"tiers"`
}

// SkyblockCollectionTier defines information about the tier of a SkyblockCollectionItem.
type SkyblockCollectionTier struct {
	Tier           int            `json:"tier"`
	AmountRequired int            `json:"amountRequired"`
	Unlocks        map[int]string `json:"unlock"`
}

/*
 *                           Skyblock Skills
 */

// SkyblockSkills contains all of the available Skyblock skills.
type SkyblockSkills struct {
	LastUpdated int64                              `json:"lastUpdated"`
	Version     string                             `json:"version"`
	Collections map[string]SkyblockSkillCollection `json:"collections"`
}

// SkyblockSkillCollection contains collections of Skyblock skills (grouped into farming, etc.)
type SkyblockSkillCollection struct {
	Name        string                          `json:"name"`
	Description string                          `json:"description"`
	MaxLevel    int                             `json:"maxLevel"`
	Levels      []SkyblockSkillCollectionLevels `json:"levels"`
}

// SkyblockSkillCollectionLevels contains information about what certain skills unlock at
// certain levels.
type SkyblockSkillCollectionLevels struct {
	Level              int      `json:"level"`
	ExperienceRequired int      `json:"totalExpRequired"`
	Unlocks            []string `json:"unlocks"`
}

/*
 *                           Player Count
 */

// PlayerCount just contains.. the player count.
type PlayerCount struct {
	PlayerCount int `json:"playerCount"`
}

/*
 *                           Key
 */

type Key struct {
	Record struct {
		Owner        string `json:"ownerUuid"`
		Key          string `json:"key"`
		TotalQueries int    `json:"totalQueries"`
	}
}

/*
 *                           Friends
 */

type Friends struct {
	Records []struct {
		ID           string `json:"_id"`
		UUIDSender   string `json:"uuidSender"`
		UUIDReceiver string `json:"uuidReceiver"`
		Started      int64  `json:"started"`
	} `json:"records"`
}

/*
 *                           Boosters
 */

type Boosters struct {
	Boosters []struct {
		ID             string `json:"_id"`
		PurchaserUUID  string `json:"purchaserUuid"`
		Amount         int    `json:"amount"`
		OriginalLength int    `json:"originalLength"`
		Length         int    `json:"length"`
		GameType       int    `json:"gameType"`
		Stacked        bool   `json:"stacked,omitempty"`
		DateActivated  int64  `json:"dateActivated"`
	} `json:"boosters"`
	BoosterState struct {
		Decrementing bool `json:"decrementing"`
	} `json:"boosterState"`
}
