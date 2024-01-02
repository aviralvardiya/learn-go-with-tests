package poker

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
	// "time"
)

// type BlindAlerter interface {
// 	ScheduleAlertAt(duration time.Duration, amount int)
// }

type CLI struct {
	in   *bufio.Scanner
	out  io.Writer
	game Game
}

func NewCLI(in io.Reader, out io.Writer, game Game) *CLI {
	return &CLI{
		in:   bufio.NewScanner(in),
		out:  out,
		game: game,
	}
}

const PlayerPrompt = "Please enter the number of players: "

func (cli *CLI) PlayPoker() {
	fmt.Fprint(cli.out, PlayerPrompt)

	numberOfPlayersInput := cli.readLine()
	numberOfPlayers, err := strconv.Atoi(strings.Trim(numberOfPlayersInput, "\n"))

	if err!=nil{
		fmt.Print("Please enter a numeric input\n")
		return
	}
	cli.game.Start(numberOfPlayers, cli.out)

	winnerInput := cli.readLine()
	winner,err := extractWinner(winnerInput)

	if err != nil {
		fmt.Fprint(cli.out, err)
		return
	}

	cli.game.Finish(winner)
}

func extractWinner(userInput string) (string, error) {
	if !strings.Contains(userInput, " wins") {
		return "", errors.New(BadWinnerInputMsg)
	}
	return strings.Replace(userInput, " wins", "", 1), nil
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}

// func (cli *CLI) scheduleBlindAlerts(numberOfPlayers int) {
// 	blindIncrement := time.Duration(5+numberOfPlayers) * time.Minute

// 	blinds := []int{100, 200, 300, 400, 500, 600, 800, 1000, 2000, 4000, 8000}
// 	blindTime := 0 * time.Second
// 	for _, blind := range blinds {
// 		cli.alerter.ScheduleAlertAt(blindTime, blind)
// 		blindTime = blindTime + blindIncrement
// 	}
// }

type Game interface {
	Start(numberOfPlayers int,alertsDestination io.Writer)
	Finish(winner string)
}

const BadPlayerInputErrMsg = "Bad value received for number of players, please try again with a number"

const BadWinnerInputMsg = "invalid winner input, expect format of 'PlayerName wins'"