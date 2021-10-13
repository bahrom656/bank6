package stats

import (
	"fmt"
	"github.com/bahrom656/bank/pkg/bank/types"
)

func ExampleAvg() {
	payments := []types.Payment{
		{ID: 0,
			Amount:   1,
			Category: "card",
		},
		{ID: 1,
			Amount:   2,
			Category: "card",
		},
		{ID: 2,
			Amount:   3,
			Category: "card",
		},
	}

	fmt.Println(Avg(payments))
	//Output: 2
}

func ExampleTotalInCategory() {
	payments := []types.Payment{
		{ID: 0,
			Amount:   1,
			Category: "card",
		},
		{ID: 1,
			Amount:   2,
			Category: "c",
		},
		{ID: 2,
			Amount:   3,
			Category: "card",
		},
	}
	fmt.Println(TotalInCategory(payments, "card"))
	//Output: 4
}
