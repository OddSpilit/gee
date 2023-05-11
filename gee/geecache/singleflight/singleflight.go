package singleflight

import "sync"

type call struct {
	wg  sync.WaitGroup
	val interface{}
	err error
}

type Group struct {
	mu sync.Mutex
	m  map[string]*call
}

/**
每次大量请求， 控制fn函数执行次数
*/
func (g *Group) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	g.mu.Lock()
	if g.m[key] == nil {
		g.m = make(map[string]*call)
	}

	if c, ok := g.m[key]; ok {
		g.mu.Unlock()
		c.wg.Wait()
		return c.val, c.err
	}

	c := new(call)
	c.wg.Add(1)
	g.m[key] = c
	g.mu.Unlock() // 已经设置了key => c, 没必要锁

	c.val, c.err = fn()
	c.wg.Done() // 告知已完成值跟错误信息的设置

	// 貌似这块可以不加锁操作
	g.mu.Lock()
	delete(g.m, key)
	g.mu.Unlock()

	return c.val, c.err
}
