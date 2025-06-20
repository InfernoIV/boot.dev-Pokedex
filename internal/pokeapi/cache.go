package pokeapi

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time // - A time.Time that represents when the entry was created.
	val       []byte    //- A []byte that represents the raw data we're caching.
}

type cache struct {
	CacheEntry map[string]cacheEntry
	Add        func(key string, val []byte)
	Get        func(key string) (data []byte, found bool)
	ReapLoop   func()
	Mu         *sync.Mutex
}

func NewCache(interval time.Duration) cache {
	//create cache
	Cache := cache{}

	//set mutex
	Cache.Mu = &sync.Mutex{}

	//create empty mapping
	Cache.CacheEntry = make(map[string]cacheEntry)

	//Create a cache.Add() method that adds a new entry to the cache. It should take a key (a string) and a val (a []byte).
	Cache.Add = func(key string, val []byte) {
		//lock cache
		Cache.Mu.Lock()
		//create an empty entry
		entry := cacheEntry{}
		//set the values
		entry.createdAt = time.Now()
		//add the value
		entry.val = val
		//save to cache
		Cache.CacheEntry[key] = entry
		//unlock cache
		Cache.Mu.Unlock()
	}

	//Create a cache.Get() method that gets an entry from the cache. It should take a key (a string) and return a []byte and a bool. The bool should be true if the entry was found and false if it wasn't.
	Cache.Get = func(key string) (data []byte, found bool) {
		//get the clicommand and check if it is ok
		cacheEntry, ok := Cache.CacheEntry[key]
		//if command is in the list
		if ok {
			return cacheEntry.val, true
		}
		return nil, false
	}

	//Create a cache.reapLoop() method that is called when the cache is created (by the NewCache function).
	// Each time an interval (the time.Duration passed to NewCache) passes it should remove any entries that are older than the interval. This makes sure that the cache doesn't grow too large over time. F
	// or example, if the interval is 5 seconds, and an entry was added 7 seconds ago, that entry should be removed.
	Cache.ReapLoop = func() {
		//create a ticker
		ticker := time.NewTicker(interval)
		//make sure it is stopped when the program ends
		defer ticker.Stop()
		//if ticked
		for range ticker.C {
			//lock cache
			Cache.Mu.Lock()
			//get the time of now
			now := time.Now()
			//create a list of entries to be deleted
			var cache_to_be_deleted []string
			//go over all entries
			for k, v := range Cache.CacheEntry {
				//if they have expired
				if now.Sub(v.createdAt) > interval {
					//add to the list to be deleted
					cache_to_be_deleted = append(cache_to_be_deleted, k)
				}
			}

			//for every that needs to be deleted
			for _, v := range cache_to_be_deleted {
				//delete the entry
				delete(Cache.CacheEntry, v)
			}
			//unlock cache
			Cache.Mu.Unlock()
		}
	}


		
	
	//call the function
	go Cache.ReapLoop()

	return Cache
}
