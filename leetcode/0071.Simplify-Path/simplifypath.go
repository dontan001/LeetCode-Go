package leetcode

import "strings"

type StackString struct {
	items []string
}

func (ss *StackString) pop() string {
	var last string
	if len(ss.items) != 0 {
		last = ss.items[len(ss.items)-1]
		ss.items = ss.items[:len(ss.items)-1]
	}

	return last
}

func (ss *StackString) push(item string) {
	ss.items = append(ss.items, item)
}

func (ss *StackString) length() int {
	if ss.items == nil {
		return 0
	} else {
		return len(ss.items)
	}
}

func SimplifyPath(path string) string {
	parts := strings.Split(path, "/")
	output := StackString{}

	for _, part := range parts {
		if part == "" || part == "." {
			continue
		}

		if part == ".." {
			output.pop()
		} else {
			output.push(part)
		}
	}

	if output.length() != 0 {
		return "/" + strings.Join(output.items, "/")
	} else {
		return "/"
	}
}
