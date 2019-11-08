package voting

import (
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"testing"
)

func TestVoting(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		
		votesCount := 5
		votes := make([]string, votesCount)
		votes[0]="1"
		votes[1]="2"
		votes[2]="3"
		votes[3]="4"
		votes[4]="5"


		x := make([]string, votesCount)
		y := make([]*big.Int, votesCount)
		V := make([]*big.Int, votesCount)
		Y := make([]*big.Int, votesCount)
		RegVote := make([]*big.Int, votesCount)

		g, p := get_g_p("512")

		// Voter random values
		x[0] = strconv.Itoa(rand.Intn(1e12))
		x[1] = strconv.Itoa(rand.Intn(1e12))
		x[2] = strconv.Itoa(rand.Intn(1e12))
		x[3] = strconv.Itoa(rand.Intn(1e12))
		x[4] = strconv.Itoa(rand.Intn(1e12))

		// Voting keys (broadcast by each voter)
		y[0] = makeG(g, x[0], p)
		y[1] = makeG(g, x[1], p)
		y[2] = makeG(g, x[2], p)
		y[3] = makeG(g, x[3], p)
		y[4] = makeG(g, x[4], p)

		// Voter values
		V[0] = makeG(g, votes[0], p)
		V[1] = makeG(g, votes[1], p)
		V[2] = makeG(g, votes[2], p)
		V[3] = makeG(g, votes[3], p)
		V[4] = makeG(g, votes[4], p)

		// Each voter calculates Y[i]
		Y[0] = calcit(0, votesCount, g, p, y)
		Y[1] = calcit(1, votesCount, g, p, y)
		Y[2] = calcit(2, votesCount, g, p, y)
		Y[3] = calcit(3, votesCount, g, p, y)
		Y[4] = calcit(4, votesCount, g, p, y)

		RegVote[0] = getVote(Y[0], x[0], p, V[0])
		RegVote[1] = getVote(Y[1], x[1], p, V[1])
		RegVote[2] = getVote(Y[2], x[2], p, V[2])
		RegVote[3] = getVote(Y[3], x[3], p, V[3])
		RegVote[4] = getVote(Y[4], x[4], p, V[4])

		res0 := mult(votesCount, RegVote, p)

		fmt.Printf("Vote1: %s, Vote2: %s, Vote3: %s, Vote4: %s, Vote5: %s\n", votes[0], votes[1], votes[2], votes[3], votes[4])
		fmt.Printf("\n\nResult: %s", res0)

		for i := 1; i <= 10000; i++ {
			m, _ := new(big.Int).SetString(strconv.Itoa(i), 10)
			gm := new(big.Int).Exp(g, m, p)
			if gm.Cmp(res0) == 0 {
				fmt.Printf("\n\nTotal votes: %d", i)
				break
			}
		}

		fmt.Printf("\n\nVoter 1 (Vote Registration): %s", y[0])
		fmt.Printf("\n\nVoter 2 (Vote Registration): %s", y[1])
		fmt.Printf("\n\nVoter 3 (Vote Registration): %s", y[2])
		fmt.Printf("\n\nVoter 4 (Vote Registration): %s", y[3])
		fmt.Printf("\n\nVoter 5 (Vote Registration): %s", y[4])
		
		fmt.Printf("\n\nVoter 1 (Vote Value): %s", RegVote[0])
		fmt.Printf("\n\nVoter 2 (Vote Value): %s", RegVote[1])
		fmt.Printf("\n\nVoter 3 (Vote Value): %s", RegVote[2])
		fmt.Printf("\n\nVoter 4 (Vote Value): %s", RegVote[3])
		fmt.Printf("\n\nVoter 5 (Vote Value): %s", RegVote[4])
	})
}
