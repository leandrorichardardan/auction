package auction

// Bidder represents a participant in the auction.
type Bidder struct {
	// Name of the bidder.
	Name string
	// InitialBid is the first and lowest bid the bidder is willing to offer.
	InitialBid float64
	// MaxBid is the maximum amount the bidder is willing to pay.
	MaxBid float64
	// BidIncrement is the amount by which the current bid should be incremented
	// if the bidder is losing, but it should not exceed the MaxBid.
	BidIncrement float64
	// CurrentBid is the current bid amount for the bidder, updated during the
	// auction process.
	CurrentBid float64
}

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
