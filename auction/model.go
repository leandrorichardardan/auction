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
