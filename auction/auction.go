package auction

// NewBidder creates a new bidder with the specified parameters.
func NewBidder(name string, initialBid, maxBid, bidIncrement float64) *Bidder {
	return &Bidder{
		Name:         name,
		InitialBid:   initialBid,
		MaxBid:       maxBid,
		BidIncrement: bidIncrement,
		CurrentBid:   initialBid,
	}
}

// DetermineWinner determines the winner of the auction based on the bidding rules.
func DetermineWinner(bidders []*Bidder) *Bidder {
	for {
		// Track if any bidder can still increase their bid.
		anyBidUpdated := false

		// Increment each bidder's bid if possible.
		for _, bidder := range bidders {
			if bidder.CurrentBid+bidder.BidIncrement <= bidder.MaxBid {
				bidder.CurrentBid += bidder.BidIncrement
				anyBidUpdated = true
			}
		}

		// Determine the bidder with the lowest possible winning bid.
		lowestWinningBid := findLowestWinningBid(bidders)

		// Check if there is a clear winner.
		winnerCandidates := findWinnerCandidates(bidders, lowestWinningBid)

		// If only one bidder has the lowest winning bid, they are the winner.
		if len(winnerCandidates) == 1 {
			return winnerCandidates[0]
		}

		// If no further increments are possible, and we have a tie, the first bidder wins.
		if !anyBidUpdated {
			return winnerCandidates[0]
		}
	}
}

func findLowestWinningBid(bidders []*Bidder) float64 {
	lowestWinningBid := bidders[0].CurrentBid
	for _, bidder := range bidders {
		if bidder.CurrentBid < lowestWinningBid {
			lowestWinningBid = bidder.CurrentBid
		}
	}
	return lowestWinningBid
}

func findWinnerCandidates(bidders []*Bidder, lowestWinningBid float64) []*Bidder {
	var candidates []*Bidder
	for _, bidder := range bidders {
		if bidder.CurrentBid == lowestWinningBid {
			candidates = append(candidates, bidder)
		}
	}
	return candidates
}
