type RecentCounter struct {
	data []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		data: []int{},
	}
}

func (this *RecentCounter) Ping(t int) int {
	this.data = append(this.data, t)
	for len(this.data) > 0 && (this.data[0] < (t - 3000)) {
		this.data = this.data[1:]
	}
	return len(this.data)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */