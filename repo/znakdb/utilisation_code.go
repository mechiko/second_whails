package znakdb

func (z *DbZnak) UtilisationCodes(id int64) (out []string, err error) {
	return z.fetchAndParseCodes("order_mark_utilisation_codes", "id_order_mark_utilisation", id)
}
