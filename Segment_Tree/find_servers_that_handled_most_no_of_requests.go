/*
{
    "problem_name": "Find Servers That Handled Most Number of Requests",
    "category": "Segment Tree / Priority Queue",
    "time_complexity": "O(N log K) -> Custom slice-based Min-Heap for busy servers and Segment Tree for O(log K) next-available server queries.",
    "space_complexity": "O(K) -> Strictly pre-allocated arrays. Zero runtime GC spikes."
}
*/
package main

// --- BARE-METAL HEAP FOR BUSY SERVERS ---
// Tracks {endTime, serverId}
type BusyNode struct { t, id int }
type BusyHeap []BusyNode

func (h *BusyHeap) push(n BusyNode) {
	*h = append(*h, n)
	h.up(len(*h) - 1)
}
func (h *BusyHeap) pop() BusyNode {
	n := len(*h) - 1
	res := (*h)[0]
	(*h)[0] = (*h)[n]
	*h = (*h)[:n]
	if n > 0 { h.down(0) }
	return res
}
func (h BusyHeap) up(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if h[i].t < h[p].t {
			h[i], h[p] = h[p], h[i]
			i = p
		} else { break }
	}
}
func (h BusyHeap) down(i int) {
	n := len(h)
	for {
		min := i
		l, r := 2*i+1, 2*i+2
		if l < n && h[l].t < h[min].t { min = l }
		if r < n && h[r].t < h[min].t { min = r }
		if min == i { break }
		h[i], h[min] = h[min], h[i]
		i = min
	}
}

// --- FLAT SEGMENT TREE FOR O(log K) FREE SERVER LOOKUP ---
// tree[node] stores 1 if AT LEAST ONE server is free in its range, else 0
type SegTree struct {
	tree []int
	size int
}

func newSegTree(n int) *SegTree {
	st := &SegTree{tree: make([]int, 4*n), size: n}
	st.build(1, 0, n-1)
	return st
}
func (st *SegTree) build(node, l, r int) {
	if l == r {
		st.tree[node] = 1 // All servers initially free
		return
	}
	mid := (l + r) / 2
	st.build(2*node, l, mid)
	st.build(2*node+1, mid+1, r)
	st.tree[node] = st.tree[2*node] | st.tree[2*node+1]
}
func (st *SegTree) update(node, l, r, idx, val int) {
	if l == r {
		st.tree[node] = val
		return
	}
	mid := (l + r) / 2
	if idx <= mid {
		st.update(2*node, l, mid, idx, val)
	} else {
		st.update(2*node+1, mid+1, r, idx, val)
	}
	st.tree[node] = st.tree[2*node] | st.tree[2*node+1]
}

// Finds the smallest available index >= startIdx
func (st *SegTree) queryFirstAvailable(node, l, r, startIdx int) int {
	if st.tree[node] == 0 || r < startIdx {
		return -1 // No free server in this range
	}
	if l == r {
		return l // Found the exact server
	}
	mid := (l + r) / 2
	// Always try left branch first (if it overlaps with startIdx) to get the smallest index
	res := st.queryFirstAvailable(2*node, l, mid, startIdx)
	if res != -1 {
		return res
	}
	// If left fails, try right branch
	return st.queryFirstAvailable(2*node+1, mid+1, r, startIdx)
}

// --- MAIN EXECUTION ENGINE ---
func busiestServers(k int, arrival []int, load []int) []int {
	st := newSegTree(k)
	busyQ := make(BusyHeap, 0, k)
	reqCount := make([]int, k)
	maxReq := 0

	for i, arrTime := range arrival {
		// 1. Free up servers whose tasks are done by 'arrTime'
		for len(busyQ) > 0 && busyQ[0].t <= arrTime {
			done := busyQ.pop()
			st.update(1, 0, k-1, done.id, 1) // Mark as free in Segment Tree
		}

		// 2. Find next available server starting from i % k
		target := i % k
		assignId := st.queryFirstAvailable(1, 0, k-1, target)

		// 3. If none found to the right, wrap around and search from 0
		if assignId == -1 {
			assignId = st.queryFirstAvailable(1, 0, k-1, 0)
		}

		// 4. If a server is found, deploy it!
		if assignId != -1 {
			st.update(1, 0, k-1, assignId, 0) // Mark as busy
			busyQ.push(BusyNode{t: arrTime + load[i], id: assignId})
			reqCount[assignId]++
			if reqCount[assignId] > maxReq {
				maxReq = reqCount[assignId]
			}
		}
	}

	// 5. Collect the elite tier servers
	var ans []int
	for i, c := range reqCount {
		if c == maxReq {
			ans = append(ans, i)
		}
	}
	return ans
}
