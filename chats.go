package main

type Chat struct {
	users  State[int, int]
	search []int
}

func (c *Chat) AddToSearch(id int) {
	c.search = append(c.search, id)
}

func (c *Chat) RemoveFromSearch(id int) {
	// bro im tired... it's 4 am...
}

func (c *Chat) Connect(a int, b int) {
	c.users.Set(a, b)
	c.users.Set(b, a)
}

func (c *Chat) Disconnect(a int, b int) {
	c.users.Set(a, 0)
	c.users.Set(b, 0)
}

func (c *Chat) Get(id int) int {
	return c.users.Get(id)
}
