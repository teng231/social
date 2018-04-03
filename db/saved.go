package db

import (
	m "github.com/my0sot1s/social/mirrors"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) CreateSaved(saved *m.Saved) error {
	collection := db.Db.C(saveCollection)
	saved.ID = bson.NewObjectId()
	err := collection.Insert(&saved)
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) RemoveSaved(sid string) error {
	collection := db.Db.C(saveCollection)
	err := collection.RemoveId(bson.ObjectIdHex(sid))
	if err != nil {
		return err
	}
	return nil
}
func (db *DB) ListSaved(limit int, anchor, uid string) (error, []*m.Saved) {
	saveds := make([]*m.Saved, 0)
	err, ms := db.ReadByIdCondition(saveCollection, anchor, limit, bson.M{"user_id": uid})
	if err != nil {
		return err, nil
	}
	for _, v := range ms {
		a := &m.Saved{}
		a.ToSaved(v)
		saveds = append(saveds, a)
	}
	return nil, saveds
}
