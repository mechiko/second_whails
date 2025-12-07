package cisinfo

import (
	"encoding/json"
	"fmt"
	"korrectkm/domain"
	"korrectkm/domain/models/modeltrueclient"
	"korrectkm/reductor"
	"korrectkm/trueclient"
	"time"
)

func (t *page) info(cis string) error {
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

	resp := []domain.CisResult{}
	data.CisInfo = domain.CisResult{}
	data.Json = ""
	data.MapInfo = make(map[string]interface{})
	if jsonRaw, err := tc.CisesList(&resp, []string{cis}); err != nil {
		return fmt.Errorf("%w", err)
	} else {
		if len(resp) > 0 {
			data.CisInfo = resp[0]
		}
		respMap := make([]map[string]interface{}, 0)
		err := json.Unmarshal([]byte(jsonRaw), &respMap)
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		if len(respMap) > 0 {
			data.MapInfo = respMap[0]
		}
		prettyJSON, err := json.MarshalIndent(data.MapInfo, "", "    ") // 4 spaces for indent
		if err != nil {
			return fmt.Errorf("%w", err)
		}
		data.Json = string(prettyJSON)
	}
	data.Updated = time.Now()
	reductor.SetModel(data, false)
	return nil
}
