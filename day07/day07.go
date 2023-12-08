package day07

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/HSLU-Student/AoC-2023/util"
)

type Day07 struct{}

var CARDVALUE = map[rune]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

var CARDVALUEJOKER = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
}

type Game struct {
	wintype int
	cards   []int
	bid     int
}

func (d Day07) Part1(input string) util.Solution {
	starttime := time.Now()
	games := []Game{}
	for _, game := range util.SplitContentLine(input) {
		unrankedgame, cardmap, mostcommoncard := parseGame(game, CARDVALUE)
		games = append(games, rankGame(unrankedgame, cardmap[mostcommoncard], cardmap))
	}

	//sort
	sortGames(games)

	//calc result
	total := 0
	for rank, game := range games {
		total += (rank + 1) * game.bid
	}

	return util.NewSolution(total, 1, time.Since(starttime))
}

func (d Day07) Part2(input string) util.Solution {
	starttime := time.Now()
	games := []Game{}
	for _, game := range util.SplitContentLine(input) {
		unrankedgame, cardmap, mostcommoncard := parseGameFiltered(game, CARDVALUEJOKER, CARDVALUEJOKER['J'])

		//manipulate cardmap if jokers exist
		jokers, exist := cardmap[CARDVALUEJOKER['J']]
		if exist {
			//add jokers to most common card - if most common card = -1 there are only jokers
			if mostcommoncard != -1 {
				cardmap[mostcommoncard] += jokers
			} else {
				cardmap[CARDVALUEJOKER['A']] += jokers
			}

			//drop joker key value
			delete(cardmap, CARDVALUEJOKER['J'])
		}

		games = append(games, rankGame(unrankedgame, cardmap[mostcommoncard], cardmap))
	}

	//sort
	sortGames(games)

	//calc res
	total := 0
	for rank, game := range games {
		total += (rank + 1) * game.bid
	}

	return util.NewSolution(total, 2, time.Since(starttime))
}

func parseGame(input string, cardvalues map[rune]int) (Game, map[int]int, int) {
	return parseGameFiltered(input, cardvalues, -1)
}

func parseGameFiltered(input string, cardvalues map[rune]int, jokerfilter int) (Game, map[int]int, int) {

	splitstr := strings.Split(input, " ")

	//bid
	bid, _ := strconv.Atoi(splitstr[1])

	//find game type
	cardmap := map[int]int{}
	cards := []int{}

	maxsamecard := 0
	mostcommoncard := -1
	for _, char := range splitstr[0] {
		value := cardvalues[char]
		cardmap[value] = cardmap[value] + 1

		if cardmap[value] > maxsamecard && value != jokerfilter {
			maxsamecard = cardmap[value]
			mostcommoncard = value
		}

		cards = append(cards, value)
	}
	return Game{-1, cards, bid}, cardmap, mostcommoncard
}

func rankGame(unrankedGame Game, maxsamecard int, cardmap map[int]int) Game {
	// 5 of a kind
	if len(cardmap) == 1 {
		return Game{6, unrankedGame.cards, unrankedGame.bid}
	}

	// 4 of a kind | full house
	if len(cardmap) == 2 {
		if maxsamecard == 4 {
			return Game{5, unrankedGame.cards, unrankedGame.bid}
		} else {
			return Game{4, unrankedGame.cards, unrankedGame.bid}
		}
	}

	//3 of a kind | two pair
	if len(cardmap) == 3 {
		if maxsamecard == 3 {
			return Game{3, unrankedGame.cards, unrankedGame.bid}
		} else {
			return Game{2, unrankedGame.cards, unrankedGame.bid}
		}
	}

	//1 pair
	if maxsamecard == 2 {
		return Game{1, unrankedGame.cards, unrankedGame.bid}
	}

	//high card
	return Game{0, unrankedGame.cards, unrankedGame.bid}
}

func sortGames(games []Game) {
	sort.Slice(games, func(i, j int) bool {
		return games[i].wintype < games[j].wintype
	})

	sort.SliceStable(games, func(i, j int) bool {
		if games[i].wintype == games[j].wintype {
			for idx, val := range games[i].cards {
				if val == games[j].cards[idx] {
					continue
				}
				return val < games[j].cards[idx]
			}
		}
		return false
	})

}
