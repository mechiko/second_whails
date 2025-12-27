package target

import (
	"fmt"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/reductor"
	"korrectkm/trueclient"
	"strings"
	"sync/atomic"
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

	mtcl, err := reductor.Model[*modeltrueclient.TrueClientModel](domain.TrueClient, t)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	tc, err := trueclient.NewFromModelSingle(t, mtcl)
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	// cises := map[string]domain.TargetCis{}
	page := 1
	for {
		dataTarget := domain.TargetResult{}
		if err := tc.TargetFilterPost(&dataTarget, data.Filter); err != nil {
			return fmt.Errorf("%w", err)
		}
		if len(dataTarget.Result) == 0 {
			if dataTarget.IsLastPage {
				break
			}
			return fmt.Errorf("результат 0")
		}
		for _, rec := range dataTarget.Result {
			out := domain.TargetCis{Result: rec}
			// cises[rec.Cis] = out
			data.TargetCis = append(data.TargetCis, out)

		}
		lastcis := dataTarget.Result[len(dataTarget.Result)-1]
		data.Filter.Pagination.Sgtin = lastcis.Cis
		data.Filter.Pagination.LastEmissionDate = lastcis.EmissionDate
		if dataTarget.IsLastPage {
			break
		}
		page++
		data.Progress = page
		t.ModelUpdate(data)
	}
	if len(data.TargetCis) > 0 {
		cisStatus := make(map[string][]domain.TargetCis)
		for _, cisItem := range data.TargetCis {
			// status := domain.DictCisTypes(cisItem.Result.Status)
			producedDate := cisItem.EmissionDate
			if cisItem.EmissionDate == "" {
				producedDate = "отсутствует"
			}
			status := domain.StatusNameByAlias[strings.ToLower(cisItem.Result.Status)]
			if status == "" {
				status = strings.ToLower(cisItem.Result.Status)
			}
			if cisStatus[status] == nil {
				cisStatus[status] = make([]domain.TargetCis, 0)
			}
			newCis := domain.TargetCis{Result: domain.Result{Cis: cisItem.Result.Cis, Status: status, ReceiptDate: producedDate}}
			cisStatus[status] = append(cisStatus[status], newCis)
		}
		data.CisStatus = cisStatus
	}
	t.ModelUpdate(data)
	// if err := reductor.SetModel(data, false); err != nil {
	// 	return fmt.Errorf("set model: %w", err)
	// }
	return nil
}

// func dateString(t time.Time) string {
// 	return t.Format(domain.TargetDateLayout)
// }
