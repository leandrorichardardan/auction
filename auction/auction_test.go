package auction

import "testing"

func TestDetermineWinner(t *testing.T) {
	tests := []struct {
		name     string
		bidders  []*Bidder
		expected string
	}{
		{
			name: "Auction 1",
			bidders: []*Bidder{
				NewBidder("Sasha", 50.00, 80.00, 3.00),
				NewBidder("John", 60.00, 82.00, 2.00),
				NewBidder("Pat", 55.00, 85.00, 5.00),
			},
			expected: "Sasha",
		},
		{
			name: "Auction 2",
			bidders: []*Bidder{
				NewBidder("Riley", 700.00, 725.00, 2.00),
				NewBidder("Morgan", 599.00, 725.00, 15.00),
				NewBidder("Charlie", 625.00, 725.00, 8.00),
			},
			expected: "Morgan",
		},
		{
			name: "Auction 3",
			bidders: []*Bidder{
				NewBidder("Alex", 2500.00, 3000.00, 500.00),
				NewBidder("Jesse", 2800.00, 3100.00, 201.00),
				NewBidder("Drew", 2501.00, 3200.00, 247.00),
			},
			expected: "Drew",
		},
		{
			name: "Auction 4",
			bidders: []*Bidder{
				NewBidder("Alex", 2500.00, 3000.00, 500.00),
				NewBidder("Jesse", 2800.00, 3100.00, 201.00),
				NewBidder("Drew", 2501.00, 3200.00, 247.00),
				NewBidder("Tie", 2000.00, 3000.00, 500.00),
			},
			expected: "Tie",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			winner := DetermineWinner(tt.bidders)
			if winner.Name != tt.expected {
				t.Errorf("expected %s, got %s", tt.expected, winner.Name)
			}
		})
	}
}
