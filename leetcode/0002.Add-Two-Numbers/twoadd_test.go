package leetcode

import "testing"

func TestTwoAdd(t *testing.T) {

	t.Run("case 1", func(t *testing.T) {
		one := &Node{2, &Node{4, &Node{3, nil}}}
		another := &Node{5, &Node{6, &Node{4, nil}}}

		result := TwoAddV1(one, another)
		for cur := result; cur != nil; cur = cur.next {
			t.Logf("%d", cur.digit)
		}
	})

	t.Run("case 2", func(t *testing.T) {
		one := &Node{2, &Node{4, &Node{3, nil}}}
		another := &Node{5, &Node{6, &Node{4, &Node{9, nil}}}}

		result := TwoAddV1(one, another)
		for cur := result; cur != nil; cur = cur.next {
			t.Logf("%d", cur.digit)
		}
	})

	t.Run("case 3", func(t *testing.T) {
		one := &Node{5, nil}
		another := &Node{7, nil}

		result := TwoAddV1(one, another)
		for cur := result; cur != nil; cur = cur.next {
			t.Logf("%d", cur.digit)
		}
	})
}
