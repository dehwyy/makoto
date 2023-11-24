package utils

import (
	"reflect"
	"testing"

	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
	lib "github.com/dehwyy/makoto/libs/lib/pkg"
)

func Test_FilterItemsByQueryAndTags(t *testing.T) {
	bm := lib.NewBenchmark(lib.Benchmark{
		Name: "FilterItems",
	})

	type S struct {
		items      []*hashmap.Item
		query      string
		tags       []*hashmap.GetItemsPayload_FilterTags
		expect_len int
	}

	items := []*hashmap.Item{
		{
			Id:    1,
			Key:   "key",
			Value: "value",
			Extra: "extra",
			Tags:  []*hashmap.Tag{{Text: "makoto"}},
		},
		{
			Id:  2,
			Key: "zyx",
			Tags: []*hashmap.Tag{
				{Text: "tag"},
				{Text: "not a tag"},
			},
		},
		{Id: 3, Value: "abc"},
		{
			Id:    4,
			Extra: "987v",
		},
	}

	args := []S{
		{
			items:      items,
			query:      "a",
			tags:       []*hashmap.GetItemsPayload_FilterTags{},
			expect_len: 2,
		},
		{
			items:      items,
			query:      "5",
			tags:       []*hashmap.GetItemsPayload_FilterTags{},
			expect_len: 0,
		},
		{
			items:      items,
			query:      "key",
			tags:       []*hashmap.GetItemsPayload_FilterTags{},
			expect_len: 1,
		},
		{
			items:      items,
			query:      "v",
			tags:       []*hashmap.GetItemsPayload_FilterTags{},
			expect_len: 2,
		},
		{
			items:      items,
			query:      "",
			tags:       []*hashmap.GetItemsPayload_FilterTags{},
			expect_len: 4,
		},
		{
			items:      items,
			query:      "1",
			tags:       []*hashmap.GetItemsPayload_FilterTags{},
			expect_len: 1,
		},
		{
			items:      items,
			query:      "",
			tags:       []*hashmap.GetItemsPayload_FilterTags{{Text: "makoto", Include: true}},
			expect_len: 1,
		},
		{
			items:      items,
			query:      "",
			tags:       []*hashmap.GetItemsPayload_FilterTags{{Text: "makoto", Include: false}},
			expect_len: 3,
		},
		{
			items:      items,
			query:      "",
			tags:       []*hashmap.GetItemsPayload_FilterTags{{Text: "tag", Include: true}, {Text: "not a tag", Include: true}},
			expect_len: 1,
		},
	}

	bm.Start()
	defer bm.EndAndSummarize(len(args))

	for i, arg := range args {
		res := FilterItemsByQueryAndTags(arg.items, arg.query, arg.tags)
		if len(res) != arg.expect_len {
			t.Errorf("case %d, expected %v, got %v", i, arg.expect_len, len(res))
		}
	}
}

func Test_GetPart(t *testing.T) {
	bm := lib.NewBenchmark(lib.Benchmark{
		Name: "GetPart",
	})

	type S struct {
		val       []int
		res       []int
		part      int
		part_size int
	}

	args := []S{
		{
			val:       []int{1, 2, 3},
			res:       []int{1},
			part:      0,
			part_size: 1,
		},
		{
			val:       []int{},
			res:       []int{},
			part:      0,
			part_size: 2,
		},
		{
			val:       []int{},
			res:       []int{},
			part:      2,
			part_size: 2,
		},
		{
			val:       []int{1, 2},
			res:       []int{},
			part:      1,
			part_size: 2,
		},
		{
			val:       []int{1, 2, 3},
			res:       []int{3},
			part:      1,
			part_size: 2,
		}, {
			val:       []int{1, 2, 3},
			res:       []int{},
			part:      0,
			part_size: 0,
		},
		{
			val:       []int{1, 2, 3},
			res:       []int{1, 2, 3},
			part:      0,
			part_size: 5,
		},
	}

	bm.Start()
	defer bm.EndAndSummarize(len(args))

	for _, arg := range args {
		res := GetPart(arg.val, arg.part, arg.part_size)
		if !reflect.DeepEqual(res, arg.res) {
			t.Errorf("expected %v, got %v", arg.res, res)
		}
	}

}

func Test_safeSlice(t *testing.T) {
	bm := lib.NewBenchmark(lib.Benchmark{
		Name: "SafeSlice",
	})

	type S struct {
		val    []int
		expect []int
		from   int
		to     int
	}

	args := []S{
		{
			val:    []int{1, 2, 3},
			expect: []int{1},
			from:   0,
			to:     1,
		},
		{
			val:    []int{1, 2, 3},
			expect: []int{2, 3},
			from:   1,
			to:     3,
		},
		{
			val:    []int{1, 2, 3},
			expect: []int{2, 3},
			from:   1,
			to:     4,
		},
		{
			val:    []int{1, 2, 3},
			expect: []int{1, 2, 3},
			from:   0,
			to:     6,
		},
		{
			val:    []int{1, 2, 3},
			expect: []int{},
			from:   4,
			to:     6,
		},
		{
			val:    []int{1, 2, 3},
			expect: []int{},
			from:   2,
			to:     2,
		},
		{
			val:    []int{1, 2, 3},
			expect: []int{},
			from:   0,
			to:     0,
		},
	}

	bm.Start()
	defer bm.EndAndSummarize(len(args))

	for _, arg := range args {
		res := safeSlice(arg.val, arg.from, arg.to)
		if !reflect.DeepEqual(res, arg.expect) {
			t.Errorf("expected %v, got %v", arg.expect, res)
		}
	}
}

func Test_onlyUnique(t *testing.T) {
	bm := lib.NewBenchmark(lib.Benchmark{
		Name: "OnlyUnique",
	})

	type S struct {
		val []*hashmap.Item
		res []*hashmap.Item
	}

	args := []S{
		{
			val: []*hashmap.Item{
				{Id: 1},
				{Id: 2},
				{Id: 3},
			},
			res: []*hashmap.Item{
				{Id: 1},
				{Id: 2},
				{Id: 3},
			},
		},
		{
			val: []*hashmap.Item{
				{Id: 1},
				{Id: 1},
				{Id: 3},
			},
			res: []*hashmap.Item{
				{Id: 1},
				{Id: 3},
			},
		},
		{
			val: []*hashmap.Item{
				{Id: 1},
				{Id: 1},
				{Id: 1},
			},
			res: []*hashmap.Item{
				{Id: 1},
			},
		},
	}

	bm.Start()

	for _, arg := range args {
		res := onlyUniqueItems(arg.val)
		if !reflect.DeepEqual(res, arg.res) {
			t.Errorf("expected %v, got %v", arg.res, res)
		}
	}

	bm.EndAndSummarize(len(args))
}
