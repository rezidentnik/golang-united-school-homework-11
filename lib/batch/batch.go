package batch

import (
	"time"
)

type user struct {
	ID int64
}

func getOne(id int64) user {
	time.Sleep(time.Millisecond * 100)
	return user{ID: id}
}

func userLoader(userIdsPool <-chan int64, loadedUsers chan<- user) {
	for id := range userIdsPool {
		loadedUsers <- getOne(id)
	}
}

func getBatch(n int64, pool int64) (res []user) {

	userIdsPool := make(chan int64, n)
	loadedUsers := make(chan user, n)

	if n < pool {
		pool = n
	}

	var i int64 = 0
	for i = 0; i < pool; i++ {
		go userLoader(userIdsPool, loadedUsers)
	}

	for id := int64(0); id < n; id++ {
		userIdsPool <- id
	}
	close(userIdsPool)

	for i = 0; i < n; i++ {
		res = append(res, <-loadedUsers)
	}
	close(loadedUsers)

	return
}
