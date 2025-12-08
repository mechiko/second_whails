package znakdb

func (z *DbZnak) OrderCodes(id int64) (out []string, err error) {
	return z.fetchAndParseCodes("order_mark_codes_serial_numbers", "id_order_mark_codes", id)
}
