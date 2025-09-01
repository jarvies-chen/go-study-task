// 合并区间：以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
// 请你合并所有重叠的区间，并返回一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间。
// 可以先对区间数组按照区间的起始位置进行排序，然后使用一个切片来存储合并后的区间，遍历排序后的区间数组，
// 将当前区间与切片中最后一个区间进行比较，如果有重叠，则合并区间；如果没有重叠，则将当前区间添加到切片中。
package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

func mergeIntervals(intervals []Interval) []Interval {
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

func main() {
	intervals := []Interval{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	result1 := mergeIntervals(intervals)
	fmt.Println("Merged intervals is ", result1)
}
