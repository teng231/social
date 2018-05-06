package db

import (
	"fmt"

	m "github.com/my0sot1s/social/mirrors"
	"gopkg.in/mgo.v2/bson"
)

func (db *DB) GetUserByUname(username string) (error, *m.User) {
	user := &m.User{}
	err, mu := db.ReadOneBasic(userCollection, bson.M{"username": username})
	if err != nil {
		return err, nil
	}
	user.ToUser(mu)
	return nil, user
}
func (db *DB) GetUserById(id string) (error, *m.User) {
	user := &m.User{}
	err, mu := db.ReadById(userCollection, id)
	if err != nil {
		return err, nil
	}
	user.ToUser(mu)
	return nil, user
}

func (db *DB) GetUserByEmail(email string) (error, *m.User) {
	user := &m.User{}
	err, mu := db.ReadOneBasic(userCollection, bson.M{"email": email})
	if err != nil {
		return err, nil
	}
	user.ToUser(mu)
	return nil, user
}

func (db *DB) GetUserByIds(uIDs []string) (error, []*m.User) {
	collection := db.Db.C(userCollection)
	var users []*m.User
	listQueriesID := make([]bson.M, 0)
	for _, p := range uIDs {
		if p == "" {
			continue
		}
		bsonID := bson.M{"_id": bson.ObjectIdHex(p)}
		listQueriesID = append(listQueriesID, bsonID)
	}
	err := collection.Find(bson.M{"$or": listQueriesID}).All(&users)
	if err != nil {
		return err, nil
	}
	return nil, users
}

func (db *DB) CreateUser(u *m.User) error {
	collection := db.Db.C(userCollection)
	u.ID = bson.NewObjectId()
	u.AlbumName = fmt.Sprintf("@%s_%s", u.GetUserName(), u.GetID())
	err := collection.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}

// func (db *DB) GetUsersLikePost(userIDs []string) (error, []*m.User) {
// 	collection := db.Db.C(userCollection)
// 	var users []*m.User
// 	listQueriesID := make([]bson.M, 0)
// 	for _, p := range userIDs {
// 		if p == "" {
// 			continue
// 		}
// 		bsonID := bson.M{"_id": bson.ObjectIdHex(p)}
// 		listQueriesID = append(listQueriesID, bsonID)
// 	}
// 	err := collection.Find(bson.M{"$or": listQueriesID}).All(&users)
// 	if err != nil {
// 		return err, nil
// 	}
// 	return nil, users
// }

func (db *DB) UpdateStateUser(uid, state string) error {
	collection := db.Db.C(userCollection)
	selector := bson.M{"_id": bson.ObjectIdHex(uid), "state": "pendding"}
	update := bson.M{"$set": bson.M{"state": state}}
	err := collection.Update(selector, update)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) UpdateUserPassword(uid, password string) error {
	collection := db.Db.C(userCollection)
	selector := bson.M{"user_id": bson.ObjectIdHex(uid)}
	update := bson.M{"$set": bson.M{"password": password}}
	err := collection.Update(selector, update)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) SearchUser(query string) (error, []*m.User) {
	collection := db.Db.C(userCollection)
	regexQuery := bson.M{"$regex": query}
	allQuery := []bson.M{bson.M{"username": regexQuery}, bson.M{"email": regexQuery}, bson.M{"fullname": regexQuery}}
	var users []*m.User
	err := collection.Find(bson.M{"$or": allQuery}).Limit(20).All(&users)
	if err != nil {
		return err, nil
	}
	return nil, users
}
