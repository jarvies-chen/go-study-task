package task1

import "sort"

type Interval struct {
	Start int
	End   int
}

func MergeIntervals(intervals []Interval) []Interval {
	if len(intervals) == 0 {
		return nil
	}

	sort.Slice(intervals, func(i, j int) bool { //按照Start 升序排序
		return intervals[i].Start < intervals[j].Start
	})
	result := make([]Interval, 0, len(intervals))
	current := intervals[0]

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]
		if next.Start <= current.End { //start小于上次的end，有重叠
			if next.End > current.End {
				current.End = next.End
			}
		} else { //不重叠
			result = append(result, current)
			current = next
		}
	}
	result = append(result, current)
	return result
}
