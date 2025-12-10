package adjust

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/reductor"
	"korrectkm/trueclient"
	"time"
)

func (t *page) balance() error {
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
	data.Balance = &domain.Balance{}
	if err := tc.Balance(data.Balance, 0); err != nil {
		return fmt.Errorf("%w", err)
	}
	data.Updated = time.Now()
	if err := reductor.SetModel(data, false); err != nil {
		return fmt.Errorf("set model: %w", err)
	}
	return nil
}
