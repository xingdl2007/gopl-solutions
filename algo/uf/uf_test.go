package uf

import (
	"testing"
	"sort"
	"fmt"
)

func TestUnionFind(t *testing.T) {
	var accounts = []string{
		"Hanzo2@m.co", "Hanzo3@m.co",
		"Hanzo4@m.co", "Hanzo5@m.co",
		"Hanzo0@m.co", "Hanzo1@m.co",
		"Hanzo3@m.co", "Hanzo4@m.co",
		"Hanzo7@m.co", "Hanzo8@m.co",
		"Hanzo1@m.co", "Hanzo2@m.co",
		"Hanzo6@m.co", "Hanzo7@m.co",
		"Hanzo5@m.co", "Hanzo6@m.co",
	}

	var uf = NewUnionFind()
	for i := 1; i < len(accounts); i++ {
		uf.Union(accounts[i-1], accounts[i])
	}

	id := uf.Find(accounts[0])
	for i := 1; i < len(accounts); i++ {
		if id != uf.Find(accounts[i]) {
			t.Errorf("%s ID error.", accounts[i])
		}
	}
}

// 49 / 49 test cases passed.
// 489ms, 100%, union-find good job!
func accountsMerge(accounts [][]string) [][]string {
	var uf = NewUnionFind()
	for _, row := range accounts {
		// skip user names
		if len(row) <= 2 {
			uf.Insert(row[1])
		} else {
			for i := 2; i < len(row); i++ {
				uf.Union(row[i-1], row[i])
			}
		}
	}
	var res [][]string
	idxMapping := make(map[int]int)
	seen := make(map[string]bool)

	for _, row := range accounts {
		mail := row[1]
		root := uf.Find(mail)

		if _, ok := idxMapping[root]; !ok {
			idxMapping[root] = len(res)
			// user name
			res = append(res, []string{row[0]})
		}
		idx := idxMapping[root]
		for _, mail := range row[1:] {
			if !seen[mail] {
				seen[mail] = true
				res[idx] = append(res[idx], mail)
			}
		}
	}
	for _, mails := range res {
		sort.Strings(mails[1:])
	}
	return res
}

func TestAccountsMerge(t *testing.T) {
	var accounts = [][]string{
		{"Lily", "Lily3@m.co", "Lily4@m.co", "Lily17@m.co"},
		{"Lily", "Lily5@m.co", "Lily3@m.co", "Lily23@m.co"},
		{"Lily", "Lily0@m.co", "Lily1@m.co", "Lily8@m.co"},
		{"Lily", "Lily14@m.co", "Lily23@m.co"},
		{"Lily", "Lily4@m.co", "Lily5@m.co", "Lily20@m.co"},
		{"Lily", "Lily1@m.co", "Lily2@m.co", "Lily11@m.co"},
		{"Lily", "Lily2@m.co", "Lily0@m.co", "Lily14@m.co"},
	}

	fmt.Println(accountsMerge(accounts))
}

func TestAccountsMerge2(t *testing.T) {
	accounts := [][]string{{"John", "johnsmith@mail.com", "john00@mail.com"},
		{"John", "johnnybravo@mail.com"}, {"John", "johnsmith@mail.com", "john_newyork@mail.com"}, {"Mary", "mary@mail.com"}}
	fmt.Println(accountsMerge(accounts))
}

func TestAccountsMerge3(t *testing.T) {
	accounts := [][]string{
		{"Alex", "Alex5@m.co", "Alex4@m.co", "Alex0@m.co"},
		{"Ethan", "Ethan3@m.co", "Ethan3@m.co", "Ethan0@m.co"},
		{"Kevin", "Kevin4@m.co", "Kevin2@m.co", "Kevin2@m.co"},
		{"Gabe", "Gabe0@m.co", "Gabe3@m.co", "Gabe2@m.co"},
		{"Gabe", "Gabe3@m.co", "Gabe4@m.co", "Gabe2@m.co"},
	}
	fmt.Println(accountsMerge(accounts))
}
