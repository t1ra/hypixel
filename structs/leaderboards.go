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
 *                           Leaderboards
 */

// Leaderboards holds the current Hypixel leaderboards.
type Leaderboards struct {
	TNTGames []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"TNTGAMES"`
	Walls []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"WALLS"`
	QuakeCraft []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"QUAKECRAFT"`
	Skywars []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"SKYWARS"`
	VampireZ []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"VAMPIREZ"`
	Prototype []interface{} `json:"PROTOTYPE"`
	MegaWalls []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"WALLS3"`
	Skyclash []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"SKYCLASH"`
	Bedwars []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"BEDWARS"`
	BuildBattle []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"BUILD_BATTLE"`
	UHC []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"UHC"`
	Paintball []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"PAINTBALL"`
	SuperSmash []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"SUPER_SMASH"`
	TrueCombat []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"TRUE_COMBAT"`
	MCGO []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"MCGO"`
	Battleground []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"BATTLEGROUND"`
	SurvivalGames []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"SURVIVAL_GAMES"`
	MurderMystery []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"MURDER_MYSTERY"`
	Arcade []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"ARCADE"`
	Arena []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"ARENA"`
	SpeedUHC []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"SPEED_UHC"`
	Duels []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"DUELS"`
	Gingerbread []struct {
		Path     string   `json:"path"`
		Prefix   string   `json:"prefix"`
		Count    int      `json:"count"`
		Location string   `json:"location"`
		Leaders  []string `json:"leaders"`
		Title    string   `json:"title"`
	} `json:"GINGERBREAD"`
}
