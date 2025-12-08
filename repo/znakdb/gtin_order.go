package znakdb

import (
	"fmt"

	"github.com/mechiko/utility"
)

func (z *DbZnak) GtinCodes(gtin string) (out []string, err error) {
	return z.fetchAndParseCodes("order_mark_codes_serial_numbers", "gtin", gtin)
}

func (z *DbZnak) fetchAndParseCodes(collection string, field string, value interface{}) (out []string, err error) {
	sess := z.dbSession
	codes := make([]map[string]interface{}, 0)
	res := sess.Collection(collection).Find(field, value)
	if err := res.All(&codes); err != nil {
		return nil, err
	}
	mpCheck := map[string]bool{}
	out = make([]string, len(codes))
	for i, code := range codes {
		c, ok := code["code"].(string)
		if !ok {
			return nil, fmt.Errorf("%v not string %T", code["code"], code["code"])
		}
		cis, err := utility.ParseCisInfo(c)
		if err != nil {
			return nil, fmt.Errorf("parse cis error %w", err)
		}
		if _, exist := mpCheck[cis.Cis]; exist {
			return nil, fmt.Errorf("дубль %s заказ %v", cis.Cis, code["id_order_mark_codes"])
		}
		mpCheck[cis.Cis] = true
		out[i] = cis.Cis
	}
	return out, nil
}
