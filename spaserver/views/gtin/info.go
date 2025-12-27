package gtin

import (
	"encoding/json"
	"fmt"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/reductor"
	"korrectkm/trueclient"
	"strings"
	"time"
)

func (t *page) info(gtin string) error {
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

	infoResult := domain.ApiGtinInfo{}
	if err := tc.GtinShort(&infoResult, gtin); err != nil {
		return fmt.Errorf("%w", err)
	}
	if len(infoResult.Results) > 0 {
		data.Info = infoResult.Results[0]
	} else {
		return fmt.Errorf("результат пустой %s", infoResult.ErrorCode)
	}
	if data.Info.ProductGroup != "" {
		if name, exist := domain.ProductGroupByAlias[strings.ToLower(data.Info.ProductGroup)]; exist {
			data.Info.ProductGroupName = name.Name
		}
	}
	data.Updated = time.Now()
	prettyJSON, err := json.MarshalIndent(data.Info, "", "    ") // 4 spaces for indent
	if err != nil {
		return fmt.Errorf("%w", err)
	}
	data.Json = string(prettyJSON)
	reductor.SetModel(data, false)
	return nil
}
