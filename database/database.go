// Package database contains all the database low-level functionality
package database

import (
	"encoding/json"
	"fmt"
	"hash/fnv"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/boltdb/bolt"
	"github.com/paulidealiste/ErroneusDilettante/models"
)

var namesBucket = []byte("NAMES")
var surnamesBucket = []byte("SURNAMES")
var reviewsBucket = []byte("REVIEWS")

//Store is the main database type
type Store struct {
	db  *bolt.DB
	fnp string
}

//KickstartDB creates or sets the existing database as a BoltDB instance
func (s *Store) KickstartDB(fnp string) error {
	s.fnp = fnp
	ldb, _ := bolt.Open(s.fnp, 0600, nil)
	s.db = ldb
	defer s.db.Close()
	_ = s.db.Update(func(tx *bolt.Tx) error {
		tx.CreateBucketIfNotExists(namesBucket)
		tx.CreateBucketIfNotExists(surnamesBucket)
		tx.CreateBucketIfNotExists(reviewsBucket)
		return nil
	})
	fmt.Println("...")
	fmt.Println("Kickstarting the database and bucket init successfull.")
	fmt.Println("---")
	return nil
}

//HoopEntities adds the name key/value pairs to the corresponding bucket
func (s *Store) HoopEntities(entities []string, bucket models.TargetBucket) error {
	ldb, _ := bolt.Open(s.fnp, 0600, nil)
	s.db = ldb
	defer s.db.Close()
	return s.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(whichBucket(bucket))
		for _, entity := range entities {
			id := strconv.FormatUint(uint64(hash(entity)), 10)
			complete := models.EntityModel{ID: id, Content: entity}
			buf, _ := json.Marshal(complete)
			b.Put([]byte(id), buf)
			fmt.Print(".")
		}
		fmt.Println("")
		fmt.Println("---")
		fmt.Println("Finished adding entities.")
		fmt.Println("---")
		return nil
	})
}

func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}

func whichBucket(bucket models.TargetBucket) []byte {
	var b []byte
	switch bucket {
	case models.Names:
		b = namesBucket
	case models.Surnames:
		b = surnamesBucket
	case models.Reviews:
		b = reviewsBucket
	}
	return b
}

//CheerEntities prints all of the entities from one bucket
func (s *Store) CheerEntities(bucket models.TargetBucket) error {
	ldb, _ := bolt.Open(s.fnp, 0600, nil)
	s.db = ldb
	defer s.db.Close()
	return s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(whichBucket(bucket))
		c := b.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			ment := models.EntityModel{}
			json.Unmarshal(v, &ment)
			fmt.Println(ment.Content)
		}
		return nil
	})
}

//CrunchEntities prints randomly concatenated entities from the database
func (s *Store) CrunchEntities() (string, error) {
	ldb, _ := bolt.Open(s.fnp, 0600, nil)
	s.db = ldb
	defer s.db.Close()
	var crucn string
	err := s.db.View(func(tx *bolt.Tx) error {
		fmt.Println("---")
		fmt.Println("Crunch start")
		fmt.Println("---")
		//BUCKETS//
		bn := tx.Bucket(namesBucket)
		bs := tx.Bucket(surnamesBucket)
		br := tx.Bucket(reviewsBucket)
		//CURSORS AND RANDOM LENGTH//
		rln := randomWhole(1, bn.Stats().KeyN+1)
		rls := randomWhole(1, bs.Stats().KeyN+1)
		rlr := randomWhole(1, br.Stats().KeyN+1)
		cn := bn.Cursor()
		cs := bs.Cursor()
		cr := br.Cursor()
		//CRUNCH LOOPS WITH RANDOM COUNTERS//
		crunched := []string{}
		cursorTraverser(cn, rln, &crunched)
		cursorTraverser(cs, rls, &crunched)
		cursorTraverser(cr, rlr, &crunched)
		crucn = strings.Join(crunched, " ")
		return nil
	})
	return crucn, err
}

func cursorTraverser(crsr *bolt.Cursor, bound int, crunched *[]string) {
	rancan := 1
	for k, v := crsr.First(); k != nil; k, v = crsr.Next() {
		if rancan == bound {
			ment := models.EntityModel{}
			json.Unmarshal(v, &ment)
			*crunched = append(*crunched, ment.Content)
			break
		}
		rancan = rancan + 1
	}
}

func randomWhole(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
