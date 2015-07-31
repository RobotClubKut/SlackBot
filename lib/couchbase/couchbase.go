package couchbase

import (
	"crypto/md5"
	"encoding/hex"
	"strconv"

	"github.com/RobotClubKut/SlackBot/lib/conf"
	"github.com/RobotClubKut/SlackBot/lib/log"
	"github.com/RobotClubKut/SlackBot/lib/nosub"
	"github.com/couchbase/go-couchbase"
)

func InsertDataBase(noSub []nosub.Data) {
	configure := conf.ReadConfigure()
	c, err := couchbase.Connect("http://" + configure.CB.ServerName + ":" + strconv.Itoa(configure.CB.Port))
	log.Terminate(err)

	pool, err := c.GetPool(configure.CB.Pool)
	log.Terminate(err)

	bucket, err := pool.GetBucket(configure.CB.DBName)
	log.Terminate(err)

	for i, n := range noSub {
		seed := n.Title + n.Time + strconv.Itoa(i)
		hash := md5.New()
		hash.Write([]byte(seed))
		key := hex.EncodeToString(hash.Sum(nil))
		err = bucket.Set(key, 0, n)
		//err = bucket.Set(n.Key, 0, []string{n.Title, n.URL, n.ImageURL, n.Time})
		log.Terminate(err)
	}
}
