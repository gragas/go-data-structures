package queue

import "testing"

func TestQueue(t *testing.T) {
	const NUM_CASES = 7
	cases_tested := 0
	for c := 1; c <= NUM_CASES; c += 1 {
		switch c {
		case 1:
			var queue Queue
			Push(&queue, 1)
			val, err := Front(&queue)
			if (val != 1) || (err != nil) {
				t.Errorf("Case %d failed!", c)
			}
			cases_tested += 1
		case 2:
			var queue Queue
			Push(&queue, -1)
			val, err := Front(&queue)
			if (val != -1) || (err != nil) {
				t.Errorf("Case %d failed!", c)
			}
			cases_tested += 1
		case 3:
			var queue Queue
			Push(&queue, 0)
			val, err := Front(&queue)
			if (val != 0) || (err != nil) {
				t.Errorf("Case %d failed!", c)
			}
			cases_tested += 1
		case 4:
			var queue Queue
			Push(&queue, 55)
			val, err := Pop(&queue)
			if (val != 55) || (err != nil) || (!IsEmpty(&queue)) {
				t.Errorf("Case %d failed!", c)
			}
			cases_tested += 1
		case 5:
			var queue Queue
			Push(&queue, 55)
			val1, err1 := Pop(&queue)
			if (val1 != 55) || (err1 != nil) || (!IsEmpty(&queue)) {
				t.Errorf("Case %d failed!", c)
			}
			val2, err2 := Pop(&queue)
			if (val2 != 0) || (err2 == nil) {
				t.Errorf("Case %d failed!", c)
			}
			cases_tested += 1
		case 6:
			var queue Queue
			Push(&queue, 55)
			Push(&queue, 77)
			val1, err1 := Pop(&queue)
			if (val1 != 55) || (err1 != nil) {
				t.Errorf("Case %d failed!", c)
				t.Errorf("Wanted 55, got %d.", val1)
				t.Errorf("Wanted nil, got %T.", err1)
			}
			val2, err2 := Pop(&queue)
			o2 := !IsEmpty(&queue)
			if (val2 != 77) || (err2 != nil) || (!IsEmpty(&queue)) {
				t.Errorf("Case %d failed!", c)
				t.Errorf("Wanted 77, got %d.", val2)
				t.Errorf("Wanted nil, got %T.", err2)
				t.Errorf("Wanted an empty, got %T.", o2)
			}
			cases_tested += 1
		case 7:
			var queue Queue
			Push(&queue, 55)
			Push(&queue, 77)
			val1, err1 := Pop(&queue)
			if (val1 != 55) || (err1 != nil) {
				t.Errorf("Case %d failed!", c)
				t.Errorf("Wanted 77, got %d.", val1)
				t.Errorf("Wanted nil, got %T.", err1)
			}
			val2, err2 := Pop(&queue)
			o2 := !IsEmpty(&queue)
			if (val2 != 77) || (err2 != nil) || (!IsEmpty(&queue)) {
				t.Errorf("Case %d failed!", c)
				t.Errorf("Wanted 77, got %d.", val2)
				t.Errorf("Wanted nil, got %T.", err2)
				t.Errorf("Wanted an empty, got %T.", o2)
			}
			val3, err3 := Pop(&queue)
			o3 := !IsEmpty(&queue)
			if (val3 != 0) || (err3 == nil) || (o3) {
				t.Errorf("Case %d failed!", c)
				t.Errorf("Wanted not 0, got %d.", val3)
				t.Errorf("Wanted an error, got %T.", err3)
				t.Errorf("Wanted an empty, got %T.", o3)
			}
			cases_tested += 1
		}
	}
	if (cases_tested != NUM_CASES) {
		t.Errorf("Did not test all cases! Be sure update NUM_CASES.")
	}
}
