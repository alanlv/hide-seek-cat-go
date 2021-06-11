package service

import (
	"HideSeekCatGo/model"
	"sync"
)

/**
handler主要做解析参数、返回数据操作及简单的逻辑。
对于业务复杂、代码量大的逻辑放在service中处理
*/
func ListUser(username string, offset, limit int) ([]*model.UserInfo, int64, error) {
	/**
	并发处理和锁的处理
	*/
	infos := make([]*model.UserInfo, 0)
	users, count, err := model.ListUser(username, offset, limit)
	if err != nil {
		return nil, count, err
	}
	ids := []uint{}
	for _, user := range users {
		ids = append(ids, user.ID)
	}
	wg := sync.WaitGroup{}
	userList := model.UserList{
		Lock:  new(sync.Mutex),
		IdMap: make(map[uint]*model.UserInfo, len(users)),
	}
	errChan := make(chan error, 1)
	finished := make(chan bool, 1)

	// 并发处理
	for _, u := range users {
		wg.Add(1)
		go func(u *model.User) {
			defer wg.Done()

			userList.Lock.Lock()
			defer userList.Lock.Unlock()
			userList.IdMap[u.ID] = &model.UserInfo{
				ID:        u.ID,
				UserName:  u.UserName,
				CreatedAt: u.CreatedAt,
				UpdatedAt: u.UpdatedAt,
			}
		}(u)
	}
	go func() {
		wg.Wait()
		close(finished)
	}()

	select {
	case <-finished:
	case err := <-errChan:
		return nil, count, err
	}

	// 并发会打乱排序，所以用map来重新复位
	for _, id := range ids {
		infos = append(infos, userList.IdMap[id])
	}
	return infos, count, nil
}
