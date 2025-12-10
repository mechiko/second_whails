package target

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/reductor"
	"korrectkm/trueclient"
	"time"
)

func (t *page) balance() error {
	param := domain.TargetFilter{
		Filter: domain.FilterTarget{
			Gtins:         []string{"04603744398059"},
			ProductGroups: []string{"water"},
		},
		Pagination: struct {
			PerPage          int    "json:\"perPage\""
			LastEmissionDate string "json:\"lastEmissionDate\""
			Sgtin            string "json:\"sgtin\""
			Direction        int    "json:\"direction\""
		}{
			PerPage:          1000,
			LastEmissionDate: "2025-12-10T22:00:00.000Z",
			Direction:        0,
			Sgtin:            "000000000000000000000",
		},
	}

	data, err := t.PageModel()
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	mtcl, err := reductor.Model[*modeltrueclient.TrueClientModel](domain.TrueClient)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	tc, err := trueclient.NewFromModelSingle(t, mtcl)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	cises := map[string]string{}
	page := 1
	for {
		dataTarget := domain.TargetResult{}
		if err := tc.TargetFilterPost(&dataTarget, param); err != nil {
			return fmt.Errorf("%w", err)
		}
		for _, rec := range dataTarget.Result {
			cises[rec.Cis] = rec.Cis
		}
		lastcis := dataTarget.Result[len(dataTarget.Result)-1]
		param.Pagination.Sgtin = lastcis.Cis
		param.Pagination.LastEmissionDate = lastcis.EmissionDate
		fmt.Printf("page %d cises %d\n", page, len(cises))
		if dataTarget.IsLastPage {
			break
		}
		page++
	}
	data.Updated = time.Now()
	if err := reductor.SetModel(data, false); err != nil {
		return fmt.Errorf("set model: %w", err)
	}
	return nil
}
