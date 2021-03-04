package main

import (
	"fmt"
)

//Team is an struct
type Team struct {
	TeamName     string
	TableOfPlays []int
}

//Play is an structure
type Play struct {
	Team1  string
	Team2  string
	Result string
}

//Teams is an slice
var Teams []Team

//Plays is an slice
var Plays []Play

func main() {

	addTeam("Devastating Donkeys")
	addTeam("Allegoric Alaskians")
	addTeam("Blithering Badgers")
	addTeam("Courageous Californians")
	//addTeam("Allegoric Alaskians")

	starToPlay(1, 2, 1)
	addPlay(Teams[1].TeamName, Teams[2].TeamName, 1)
	starToPlay(0, 3, 2)
	addPlay(Teams[0].TeamName, Teams[3].TeamName, 3)
	starToPlay(0, 1, 1)
	addPlay(Teams[0].TeamName, Teams[1].TeamName, 1)
	starToPlay(2, 3, 1)
	addPlay(Teams[3].TeamName, Teams[2].TeamName, 2)
	starToPlay(0, 2, 1)
	addPlay(Teams[2].TeamName, Teams[0].TeamName, 2)
	starToPlay(1, 3, 1)
	addPlay(Teams[1].TeamName, Teams[3].TeamName, 1)

	screenPrintTeams(Teams)
	screenPrintPlays(Plays)

}

/*
Team                           | MP |  W |  D |  L |  P
Devastating Donkeys            |  3 |  2 |  1 |  0 |  7
Allegoric Alaskans             |  3 |  2 |  0 |  1 |  6
Blithering Badgers             |  3 |  1 |  0 |  2 |  3
Courageous Californians        |  3 |  0 |  1 |  2 |  1

Allegoric Alaskans;Blithering Badgers;win
Devastating Donkeys;Courageous Californians;draw
Devastating Donkeys;Allegoric Alaskans;win
Courageous Californians;Blithering Badgers;loss
Blithering Badgers;Devastating Donkeys;loss
Allegoric Alaskans;Courageous Californians;win
*/
//addTeam is to add a Team into the Slice
func addTeam(theTeamName string) {
	if !existTeam(theTeamName) {
		Teams = append(Teams, Team{
			TeamName:     theTeamName,
			TableOfPlays: []int{0, 0, 0, 0, 0},
		})
	} else {
		fmt.Println("Team already exist")
	}
}

//verify is the Team exist or not
func existTeam(name string) bool {
	exist := false
	for i := 0; i < len(Teams); i++ {
		if Teams[i].TeamName == name {
			exist = true
		}
	}
	return exist
}

//To add the play and the result of the play
func addPlay(n1 string, n2 string, rslt int) {
	switch rslt {
	case 1:
		Plays = append(Plays, Play{
			Team1:  n1,
			Team2:  n2,
			Result: "win",
		})
		break

	case 2:
		Plays = append(Plays, Play{
			Team1:  n1,
			Team2:  n2,
			Result: "loss",
		})
		break

	case 3:
		Plays = append(Plays, Play{
			Team1:  n1,
			Team2:  n2,
			Result: "draw",
		})
		break
	}
}

//To add the results of each Team
func starToPlay(posT1 int, posT2 int, result int) {
	mp1 := Teams[posT1].TableOfPlays[0]
	w1 := Teams[posT1].TableOfPlays[1]
	d1 := Teams[posT1].TableOfPlays[2]
	l1 := Teams[posT1].TableOfPlays[3]
	p1 := Teams[posT1].TableOfPlays[4]

	mp2 := Teams[posT2].TableOfPlays[0]
	w2 := Teams[posT2].TableOfPlays[1]
	d2 := Teams[posT2].TableOfPlays[2]
	l2 := Teams[posT2].TableOfPlays[3]
	p2 := Teams[posT2].TableOfPlays[4]
	switch result {
	case 1:

		Teams[posT1].TableOfPlays[0] = mp1 + 1
		Teams[posT1].TableOfPlays[1] = w1 + 1
		Teams[posT1].TableOfPlays[2] = d1
		Teams[posT1].TableOfPlays[3] = l1
		Teams[posT1].TableOfPlays[4] = p1 + 3

		Teams[posT2].TableOfPlays[0] = mp2 + 1
		Teams[posT2].TableOfPlays[1] = w2
		Teams[posT2].TableOfPlays[2] = d2
		Teams[posT2].TableOfPlays[3] = l2 + 1
		Teams[posT2].TableOfPlays[4] = p2

		break
	case 2:
		Teams[posT1].TableOfPlays[0] = mp1 + 1
		Teams[posT1].TableOfPlays[1] = w1
		Teams[posT1].TableOfPlays[2] = d1 + 1
		Teams[posT1].TableOfPlays[3] = l1
		Teams[posT1].TableOfPlays[4] = p1 + 1

		Teams[posT2].TableOfPlays[0] = mp2 + 1
		Teams[posT2].TableOfPlays[1] = w2
		Teams[posT2].TableOfPlays[2] = d2 + 1
		Teams[posT2].TableOfPlays[3] = l2
		Teams[posT2].TableOfPlays[4] = p2 + 1
		break
	}
}

//To show Teams slice
func screenPrintTeams(rs []Team) {
	println()
	fmt.Printf("%-30s | MP |  W |  D |  L |  P\n", "Team")
	for ind := range rs {
		fmt.Printf("%-30s | %2d | %2d | %2d | %2d | %2d\n", rs[ind].TeamName, rs[ind].TableOfPlays[0],
			rs[ind].TableOfPlays[1], rs[ind].TableOfPlays[2], rs[ind].TableOfPlays[3], rs[ind].TableOfPlays[4])
	}
}

//to show Plays slice
func screenPrintPlays(p []Play) {
	println()
	for ind := range p {
		fmt.Printf("%s , %s , %s \n", p[ind].Team1, p[ind].Team2, p[ind].Result)
	}
}
