# Auction Package

## Purpose

The `auction` package provides functionality to simulate a bidding process in an auction. It allows for the creation of bidders and determines the winner based on the lowest possible winning bid according to specific rules.

## Features

- **Bidder Creation**: Create bidders with a name, initial bid, maximum bid, and bid increment.
- **Winner Determination**: Determine the winner of the auction based on the lowest bid that satisfies the auction rules.
- **Tiebreaker Logic**: Handles ties by prioritizing bids that were placed first.

## Usage

To use this package, follow the example below:

```go
package main

import (
	"fmt"
	"github.com/leandrorichardardan/auction"
)

func main() {
	// Create bidders
	bidder1 := auction.NewBidder("Sasha", 50.00, 80.00, 3.00)
	bidder2 := auction.NewBidder("John", 60.00, 82.00, 2.00)
	bidder3 := auction.NewBidder("Pat", 55.00, 85.00, 5.00)

	// Determine the winner
	winner := auction.DetermineWinner([]*auction.Bidder{bidder1, bidder2, bidder3})

	// Print the winner
	fmt.Printf("The winner is %s with a bid of $%.2f\n", winner.Name, winner.CurrentBid)
}
```

### Auction Rules

1. **Initial Bid**: The first and lowest bid the bidder is willing to offer.
2. **Max Bid**: The maximum amount the bidder is willing to pay.
3. **Bid Increment**: The amount by which the current bid should be incremented if the bidder is losing, but it should not exceed the Max Bid.
4. **Winning Bid**: The algorithm determines the winner based on the lowest possible bid that will win the auction.
5. **Ties**: In the event of a tie, the bidder who placed their bid first will be considered the winner.

## Running Tests

To ensure that the package works as expected, you can run the included unit tests:

**Run the Tests**: Execute the following command to run all tests:

   ```bash
   make test
   ```

**Tidy Up Dependencies**: If you need to tidy up your Go module dependencies, use:

   ```bash
   make tidy
   ```