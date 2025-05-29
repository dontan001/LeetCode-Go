package leetcode

import (
	"strconv"
	"strings"
)

func RestoreIpAddresses(s string) []string {

	var ips = []string{}
	if len(s) < 4 || len(s) > 12 {
		return ips
	}

	var restore func([]string, int)
	restore = func(parts []string, begin int) {
		if len(parts) == 4 {
			if begin == len(s) {
				var ip = strings.Join(parts, ".")
				ips = append(ips, ip)
			}
			return
		}

		for i := 1; i <= 3; i++ {
			if begin+i <= len(s) {
				partStr := s[begin : begin+i]
				partInt, _ := strconv.Atoi(partStr)
				if i == 1 ||
					(i == 2 && partInt >= 10) ||
					(i == 3 && partInt >= 100 && partInt <= 255) {
					parts = append(parts, partStr)
					restore(parts, begin+i)
					parts = parts[:len(parts)-1]
				}
			}
		}
	}

	restore([]string{}, 0)
	return ips
}
