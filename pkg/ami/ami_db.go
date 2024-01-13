package ami

import (
	"context"

	"github.com/pnguyen215/voipkit/pkg/ami/config"
)

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

// DBDel Delete DB entry.
func DBDel(ctx context.Context, s AMISocket, family, key string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDBDel)
	db := NewAMIPayloadDb().SetFamily(family).SetKey(key)
	c.SetV(db)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DBDelTree delete DB tree.
func DBDelTree(ctx context.Context, s AMISocket, family, key string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDBDelTree)
	db := NewAMIPayloadDb().SetFamily(family).SetKey(key)
	c.SetV(db)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DBPut puts DB entry.
func DBPut(ctx context.Context, s AMISocket, family, key, value string) (AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDBPut)
	db := NewAMIPayloadDb().SetFamily(family).SetKey(key).SetValue(value)
	c.SetV(db)
	callback := NewAmiCallbackService(ctx, s, c, []string{}, []string{})
	return callback.Send()
}

// DBGet gets DB Entry.
func DBGet(ctx context.Context, s AMISocket, family, key string) ([]AmiReply, error) {
	c := NewCommand().SetId(s.UUID).SetAction(config.AmiActionDBGet)
	db := NewAMIPayloadDb().SetFamily(family).SetKey(key)
	c.SetV(db)
	callback := NewAmiCallbackService(ctx, s, c, []string{config.AmiListenerEventDbGetResponse}, []string{config.AmiListenerEventDBGetComplete})
	return callback.SendSuperLevel()
}
