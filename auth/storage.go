package auth

import (
	"fmt"
	"github.com/AlexKLWS/lws-blog-server/config"
	"github.com/prologic/bitcask"
	"github.com/spf13/viper"
	"log"
	"sync"
	"time"
)

var (
	activeTokens map[string]time.Time
	tokenDB      *bitcask.Bitcask
	mutex        sync.Mutex
)

var tokenLifetime = time.Duration(viper.GetInt(config.TokenLifetime)) * time.Hour

func InitializeTokenStorage() {
	db, dbOpenError := bitcask.Open("./token_DB")
	if dbOpenError != nil || db == nil {
		panic(fmt.Errorf("Could not load persisted tokens db: %s \n", dbOpenError))
	}
	tokenDB = db
	values := tokenDB.Keys()

	now := time.Now()
	for v := range values {
		t, err := tokenDB.Get(v)
		if err != nil {
			continue
		}
		parsedTime, timeErr := time.Parse(time.StampMilli, string(t))
		if timeErr != nil {
			continue
		}
		if parsedTime.Before(now) {
			tokenDB.Delete(v)
		} else {
			go func() {
				go expirationJob(v, now, parsedTime)
				mutex.Lock()
				activeTokens[string(v)] = parsedTime
				mutex.Unlock()
			}()
		}
	}
}

func AddTokenToStorage(token string) {
	now := time.Now()
	expirationTime := now.Add(tokenLifetime)
	go expirationJob([]byte(token), now, expirationTime)
	mutex.Lock()
	if activeTokens == nil {
		activeTokens = make(map[string]time.Time)
	}
	activeTokens[token] = expirationTime
	mutex.Unlock()
	if tokenDB == nil {
		log.Printf("Token %s could not be persisted cause token db is unavailable!", token)
		return
	}
	tokenDB.Put([]byte(token), []byte(expirationTime.Format(time.StampMilli)))
}

func TokenExistsInStorage(token string) bool {
	_, ok := activeTokens[token]
	return ok
}

func expirationJob(key []byte, now time.Time, expirationTimestamp time.Time) {
	diff := expirationTimestamp.Sub(now)
	time.Sleep(diff)
	if tokenDB == nil {
		log.Printf("Expiration job for key %s could not complete cause token db is unavailable!", string(key))
		return
	}

	if tokenDB.Has(key) {
		err := tokenDB.Delete(key)
		if err != nil {
			log.Printf("Could not delete key %s! Error: %d", string(key), err)
			return
		}
	}

	sk := string(key)
	if _, ok := activeTokens[sk]; ok {
		mutex.Lock()
		delete(activeTokens, sk)
		mutex.Unlock()
	}
}
