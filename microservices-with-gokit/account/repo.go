package account

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/boltdb/bolt"
	"github.com/go-kit/kit/log"
	"time"
)

type Repoisitory interface {
	Create(context.Context, User) error
}

type repo struct {
	DB     *bolt.DB
	Logger log.Logger
}

func NewRepo(db *bolt.DB, logger log.Logger) Repoisitory {
	return repo{DB: db, Logger: logger}
}

var i = 100

func (r repo) Create(ctx context.Context, user User) error {
	err := r.DB.Update(func(tx *bolt.Tx) error {
		a := tx.Bucket([]byte("Account"))
		if a == nil {
			r.Logger.Log("db error", " Account Bucket Not Found ")
		}
		u := a.Bucket([]byte("User"))
		if u == nil {
			r.Logger.Log("db error", " User Bucket Not Found ")
		}

		k := []byte(string(i))
		i++
		v, _ := json.Marshal(user)

		err := u.Put(k, v)
		if err != nil {
			r.Logger.Log("db error", r)
		}
		return nil
	})
	if err != nil {
		return err
	}
	//Get(i)
	return nil
}

func SetupBolt() (*bolt.DB, error) {
	db, err := bolt.Open("test.db", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte("Account"))
		if err != nil {
			return err
		}
		b.CreateBucketIfNotExists([]byte("User"))
		return nil
	})
	return db, nil
}

func (r repo) Get(id string) error {
	r.DB.View(func(tx *bolt.Tx) error {

		user := tx.Bucket([]byte("Account")).Bucket([]byte("User")).Get([]byte(id))
		var u User
		json.Unmarshal(user, &u)
		tx.Bucket([]byte("Account")).Bucket([]byte("User")).ForEach(func(k []byte, v []byte) error {
			var u SignUp
			json.Unmarshal(v, &u)
			fmt.Println("k = ", string(k), u)
			return nil
		})
		return nil
	})
	return nil
}
