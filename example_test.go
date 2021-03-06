package hypixel

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var (
	apiKey   string
	instance *Hypixel
)

// Tests
func TestMain(m *testing.M) {
	keyBytes, err := ioutil.ReadFile("api_key")
	apiKey = string(keyBytes)
	if err != nil {
		log.Println("Did you add your API key? See" +
			" https://github.com/t1ra/hypixel/README.md#Testing.\nError:")
		log.Fatalln(err)
	}
	instance, err = New(apiKey)
	if err != nil {
		log.Println("Failed to create Hypixel instance. Are you connected to the internet?")
		log.Fatalln(err)
	}
	os.Exit(m.Run())
}

// Examples
func ExampleNew() {
	instance, err := New("00000000-0000-0000-0000-000000000000")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(instance.getKey())
}

func ExampleHypixel_Achievements() {
	Achievements, err := instance.Achievements()
	if err != nil {
		log.Fatalln("Couldn't get achievements:", err)
	}

	fmt.Println(Achievements.LastUpdated)
}

func ExampleHypixel_BanStats() {
	BanStats, err := instance.BanStats()
	if err != nil {
		log.Fatalln("Couldn't get Ban stats:", err)
	}
	fmt.Println(BanStats.WatchdogLastMinute)
}

func ExampleHypixel_Boosters() {
	_, err := instance.Boosters()
	if err != nil {
		log.Fatalln("Couldn't get boosters:", err)
	}
}

func ExampleHypixel_Challenges() {
	_, err := instance.Challenges()
	if err != nil {
		log.Fatalln("Couldn't get boosters:", err)
	}

}

func ExampleHypixel_FindGuild() {
	// Get a guild ID by its name
	GuildID, err := instance.FindGuild("Sex Havers")
	if err != nil {
		log.Fatalln("Couldn't get guild ID from name:", err)
	}
	fmt.Println(GuildID)

	// Get a guild ID by the UUID of a member
	GuildID, err = instance.FindGuild("69a7da3c-5e94-4bbe-883e-61568c928d05")
	if err != nil {
		log.Fatalln("Couldn't get guild ID from UUID:", err)
	}
	fmt.Println(GuildID)

	// Output: 5de1d9c68ea8c96c82070699
	// 5de1d9c68ea8c96c82070699
}

func ExampleHypixel_Friends() {
	Friends, err := instance.Friends("d4acada6-bc84-4dd3-8406-0bb77e207a7a")
	// Me! :)
	if err != nil {
		log.Fatalln("Couldn't get friends:", err)
	}

	fmt.Println(len(Friends.Records)) // :(
	// Output: 1
}

func ExampleHypixel_Guild() {
	Guild, err := instance.Guild("69a7da3c-5e94-4bbe-883e-61568c928d05")
	if err != nil {
		log.Fatalln("Couldn't get guild information from player UUID:", err)
	}

	fmt.Println(Guild.Name)

	Guild, err = instance.Guild("Sex Havers")
	if err != nil {
		log.Fatalln("Couldn't get guild information from name:", err)
	}

	fmt.Println(Guild.ID)

	// Output: Sex Havers
	// 5de1d9c68ea8c96c82070699
}

func ExampleHypixel_GuildAchievements() {
	GuildAchievements, err := instance.GuildAchievements()
	if err != nil {
		log.Fatalln("Couldn't get guild achievements", err)
	}

	fmt.Println(GuildAchievements.Tiered["PRESTIGE"].Description)

	// Output: Reach Guild level %s
}

func ExampleHypixel_GuildPermissions() {
	GuildPermissions, err := instance.GuildPermissions()
	if err != nil {
		log.Fatalln("Couldn't get guild permissions", err)
	}

	fmt.Println(GuildPermissions[0].Name)
	// The "en_us" access is because, apparently, theres the ability
	// to have localised permissions. Even though all of this data is static.
	// Who knows.

	// Output: Modify Guild Name
}

func ExampleHypixel_Key() {
	Key, err := instance.Key()
	if err != nil {
		log.Fatalln("Couldn't get default key information", err)
	}

	fmt.Println(Key.Owner)

	// Output: d4acada6bc844dd384060bb77e207a7a
}

func ExampleHypixel_Leaderboards() {
	Leaderboards, err := instance.Leaderboards()
	if err != nil {
		log.Fatalln("Couldn't get leaderboards", err)
	}

	fmt.Printf("%v %v",
		Leaderboards.Skywars[1].Prefix, Leaderboards.Skywars[1].Title)

	// Output: Overall Wins
}

func ExampleHypixel_Player() {
	Player, err := instance.Player("d4acada6-bc84-4dd3-8406-0bb77e207a7a")
	if err != nil {
		log.Fatalln("Couldn't get player information", err)
	}

	fmt.Println(Player.PlayerName)

	// Output: tiraa
}

func ExampleHypixel_PlayerCount() {
	PlayerCount, err := instance.PlayerCount()
	if err != nil {
		log.Fatalln("Couldn't get player count", err)
	}

	fmt.Println(PlayerCount > 0)
	// This per-se can fail and still be truthful, but we'll just have to assume
	// that Hypixel still exists for the mean time.

	// Output: true
}

func ExampleHypixel_Quests() {
	Quests, err := instance.Quests()
	if err != nil {
		log.Fatalln("Couldn't get quests", err)
	}

	fmt.Println(Quests.Modes["hungergames"][0].Name)

	// Output: Daily Quest: Game of the Day
}

func ExampleHypixel_SkyblockCollections() {
	SkyblockCollections, err := instance.SkyblockCollections()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(SkyblockCollections["FARMING"].Items["WHEAT"].Name)
	fmt.Println(SkyblockCollections["FARMING"].Items["WHEAT"].MaxTier)

	// Output: Wheat
	// 9
}

func ExampleHypixel_SkyblockSkills() {
	SkyblockSkills, err := instance.SkyblockSkills()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(SkyblockSkills.Collections["FARMING"].Name)
	fmt.Println(SkyblockSkills.Collections["FARMING"].Description)

	// Output: Farming
	// Harvest crops and shear sheep to earn Farming XP!
}

func ExampleHypixel_GameStats() {
	Stats, err := instance.GameStats()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(Stats.GameStats.SKYWARS.Players > 0)
	// This too can theoretically return false and still be correct, but
	// practically, there should be at least one person playing it at any time.

	// Output: true
}
