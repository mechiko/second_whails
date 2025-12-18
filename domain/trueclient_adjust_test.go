package domain

import (
	"encoding/json"
	"fmt"
	"testing"
	// "github.com/stretchr/testify/assert"
)

var testsParse = []AdjustBody{
	{
		ParticipantInn: "INN",
		Codes: []AdjustStep{
			{
				Code:           []string{"1", "2"},
				ProductionDate: "date",
			},
		},
	},
	{
		ParticipantInn: "INN",
		Codes: []AdjustStep{
			{
				Code:           []string{"1", "2"},
				ProductionDate: "date",
				ExpirationDate: "date",
			},
		},
	},
	{
		ParticipantInn: "INN",
		Codes: []AdjustStep{
			{
				Code:           []string{"1", "2"},
				ExpirationDate: "date",
			},
		},
	},
	{
		ParticipantInn: "INN",
		Codes: []AdjustStep{
			{
				Code: []string{"1", "2"},
				PermitDocsAdjust: &PermitDocsAdjust{
					PermitDocsOperation: 0,
					PermitDocs: []PermitDoc{
						{PermitDocNumber: "number", PermitDocDate: "date", PermitDocType: 2},
					},
				},
			},
		},
	},
}

func TestAdjustStep(t *testing.T) {
	// The execution loop
	// Capture tt for safety, use NoError, and put expected before actual in Equal
	for _, tt := range testsParse {
		tt := tt
		t.Run(tt.ParticipantInn, func(t *testing.T) {
			str, err := json.Marshal(tt)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(str))
		})
	}
}
