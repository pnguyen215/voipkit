package ami

func NewAMIPayloadDb() *AMIPayloadDb {
	d := &AMIPayloadDb{}
	return d
}

func (d *AMIPayloadDb) SetFamily(value string) *AMIPayloadDb {
	d.Family = value
	return d
}

func (d *AMIPayloadDb) SetKey(value string) *AMIPayloadDb {
	d.Key = value
	return d
}

func (d *AMIPayloadDb) SetValue(value string) *AMIPayloadDb {
	d.Value = value
	return d
}
