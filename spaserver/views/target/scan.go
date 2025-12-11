package target

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/reductor"
	"korrectkm/trueclient"
	"sync/atomic"
	"time"
)

var reentranceSearchFlag int64

func (t *page) scan() (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("target scan panic %v", r)
		}
	}()
	// эта часть в фоне выполняется только в одном экземпляре
	if atomic.CompareAndSwapInt64(&reentranceSearchFlag, 0, 1) {
		// defer atomic.StoreInt64(&reentranceSearchFlag, 0)
		defer func() {
			atomic.StoreInt64(&reentranceSearchFlag, 0)
		}()
	} else {
		t.Logger().Errorf("reenter page:target:scan")
		return
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
	cises := map[string]domain.TargetCis{}
	page := 1
	for {
		dataTarget := domain.TargetResult{}
		if err := tc.TargetFilterPost(&dataTarget, data.Filter); err != nil {
			return fmt.Errorf("%w", err)
		}
		for _, rec := range dataTarget.Result {
			out := domain.TargetCis{Result: rec}
			cises[rec.Cis] = out
		}
		lastcis := dataTarget.Result[len(dataTarget.Result)-1]
		data.Filter.Pagination.Sgtin = lastcis.Cis
		data.Filter.Pagination.LastEmissionDate = dateString(lastcis.EmissionDate)
		t.Logger().Debugf("page %d cises %d", page, len(cises))
		if dataTarget.IsLastPage {
			break
		}
		page++
		data.Progress = page

	}
	data.Updated = time.Now()
	if err := reductor.SetModel(data, false); err != nil {
		return fmt.Errorf("set model: %w", err)
	}
	return nil
}

func dateString(t time.Time) string {
	return t.Format(domain.TargetDateLayout)
}

func truncateToStartOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func truncateToEndOfDay(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 0, t.Location())
}
