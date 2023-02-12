package tournament

import (
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	name string
	mp   int
	w    int
	d    int
	l    int
}

func (t *Team) won() {
	t.mp += 1
	t.w += 1
}

func (t *Team) draw() {
	t.mp += 1
	t.d += 1
}

func (t *Team) lost() {
	t.mp += 1
	t.l += 1
}

func (t *Team) points() int {
	return t.w*3 + t.d
}

func getTeam(teamName string, teamsMap map[string]*Team) *Team {
	team, exists := teamsMap[teamName]
	if !exists {
		team = &Team{}
		team.name = teamName
		teamsMap[teamName] = team
	}

	return team
}

func formatLine(col1, col2, col3, col4, col5, col6 interface{}) string {
	return fmt.Sprintf("%-30v | %2v | %2v | %2v | %2v | %2v\n", col1, col2, col3, col4, col5, col6)
}

func Tally(reader io.Reader, writer io.Writer) error {
	buf, _ := io.ReadAll(reader)
	lines := strings.Split(string(buf), "\n")

	teamsMap := make(map[string]*Team)
	for _, line := range lines {
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.Count(line, ";") != 2 {
			return errors.New("error parsing input; not enough semicolons")
		}
		lineChunks := strings.Split(line, ";")
		homeTeamName, awayTeamName, result := lineChunks[0], lineChunks[1], lineChunks[2]

		homeTeam := getTeam(homeTeamName, teamsMap)
		awayTeam := getTeam(awayTeamName, teamsMap)

		switch result {
		case "win":
			homeTeam.won()
			awayTeam.lost()
		case "draw":
			homeTeam.draw()
			awayTeam.draw()
		case "loss":
			homeTeam.lost()
			awayTeam.won()
		default:
			return errors.New(fmt.Sprintf("Game result is not valid: %s", result))
		}
	}

	var teams []*Team

	for _, team := range teamsMap {
		teams = append(teams, team)
	}

	sort.Slice(teams, func(i, j int) bool {
		if teams[i].points() == teams[j].points() {
			return teams[i].name < teams[j].name
		}

		return teams[i].points() > teams[j].points()
	})

	_, err := io.WriteString(writer, formatLine("Team", "MP", "W", "D", "L", "P"))
	if err != nil {
		return err
	}
	for _, team := range teams {
		_, err = io.WriteString(writer, formatLine(team.name, team.mp, team.w, team.d, team.l, team.points()))
		if err != nil {
			return err
		}
	}

	return nil
}
