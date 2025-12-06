package innfias

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/reductor"
	"korrectkm/trueclient"
	"time"
)

func (t *page) info() error {
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

	data.InnInfo = &domain.ModFiasInfo{}
	if err := tc.ModsList(data.InnInfo, []string{"7811755774"}); err != nil {
		return fmt.Errorf("%w", err)
	}
	data.Updated = time.Now()
	reductor.SetModel(data, false)
	return nil
}
