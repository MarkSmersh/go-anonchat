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
	newSearch := []int{}

	for i := range c.search {
		if c.search[i] != id {
			newSearch = append(newSearch, c.search[i])
		}
	}

	c.search = newSearch
}

func (c *Chat) GetFirstCompanion(id int) int {
	for i := 0; i < len(c.search); i++ {
		if c.search[i] != id {
			return c.search[0]
		}
	}

	return 0
}

func (c *Chat) Connect(a int, b int) {
	c.users.Set(a, b)
	c.users.Set(b, a)

	c.RemoveFromSearch(a)
	c.RemoveFromSearch(b)
}

func (c *Chat) Disconnect(a int, b int) {
	c.users.Set(a, 0)
	c.users.Set(b, 0)
}

func (c *Chat) Get(id int) int {
	return c.users.Get(id)
}
