package utils

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/dehwyy/makoto/libs/grpc/generated/hashmap"
)

type FilterTags struct {
	Text  string
	State bool // if true -> selected, have to be on item, else -> should not appear in item's tags
}

func FilterItemsByQueryAndTags(items []*hashmap.Item, query string, tags []FilterTags) []*hashmap.Item {
	// initial slice capacity
	initial_slice_cap := int(math.Min(20, float64(len(items))))
	// minimum of 5 and query.length / 2 (ceiled)
	initial_query_cap := int(math.Min(5, float64(len(query))/2))

	// split
	split_query := strings.Split(query, " ")

	// to store matched items
	exact_match := make([]*hashmap.Item, 0, initial_slice_cap)
	strong_filtered_items := make([]*hashmap.Item, 0, initial_slice_cap) // has more priority
	weak_filtered_items := make([]*hashmap.Item, 0, initial_slice_cap)

	// strong match
	queries := make([]string, 0, initial_query_cap)
	// weak match (weak match is less strict)
	// f.e query /a[a-zа-я\\s]*o[a-zа-я\\s]*/ would match `admkvopedpoko` but Strong wouldn't (it would match only words which has `ao` sequentially)
	weak_queries := make([]*regexp.Regexp, 0, initial_query_cap)
	// reserved search params looks like item?(id|key|value|extra)={value}
	reserved_search_params := make(map[string]string, 4)

	reserved_regexp := regexp.MustCompile(`^item\?(id|key|value|extra){1}=`)

	for _, query := range split_query {
		// if query doesn't match special -> common query ->
		if is_reserved := reserved_regexp.MatchString(query); !is_reserved {
			queries = append(queries, query)

			weak_query := make([]string, len(query))
			for _, r := range query {
				weak_query = append(weak_query, string(r)+"[a-zа-я\\s]*")
			}

			regex_weak_query, err := regexp.Compile(strings.Join(weak_query, ""))
			if err != nil { // should not appear, just for safety
				continue
			}

			weak_queries = append(weak_queries, regex_weak_query)

			continue
		}

		split_query = regexp.MustCompile(`[?=]`).Split(query, 2) // would be []string{"item", {`key`}, {`value`}}
		reserved_search_params[split_query[1]] = split_query[2]
	}

	// for near FOR (to remove duplicate declarations)
	const goroutines = 3

ItemsSort:
	for _, item := range items {

		falsy_tags := 0
		requested_tags := make(map[string]bool)
		for _, tag := range tags {
			if !tag.State {
				falsy_tags++
			}
			requested_tags[tag.Text] = tag.State
		}

		for _, item_tag := range item.Tags {
			val, ok := requested_tags[item_tag.Text]
			// if tag wasn't marked as request -> skip
			if !ok {
				continue
			}

			// if val is false -> should not be in item
			if !val {
				continue ItemsSort
			}

			// delete the key
			delete(requested_tags, item_tag.Text)
		}

		// if requested_tags is not empty -> not all selected tags were provided in the item
		if len(requested_tags) > falsy_tags {
			continue
		}

		is_ok := make(chan bool, 1)
		ready := make(chan bool, 1)

		go func() {
			for key, value := range reserved_search_params {

				if key == "id" && fmt.Sprint(item.Id) != value {
					is_ok <- false
				} else if key == "key" && item.Key != value {
					is_ok <- false
				} else if key == "value" && item.Value != value {
					is_ok <- false
				} else if key == "extra" && item.Extra != value {
					is_ok <- false
				}
			}

			ready <- true
		}()

		// strong match -> if any of item's field matches query -> add to strong
		go func() {
			for _, query := range queries {
				// `c` is Contains, while second letter is first letter of query
				ck := strings.Contains(item.Key, query)
				cv := strings.Contains(item.Value, query)
				ce := strings.Contains(item.Extra, query)
				ci := strings.Contains(fmt.Sprint(item.Id), query)

				if ck || cv || ce || ci {
					strong_filtered_items = append(strong_filtered_items, item)
					break // once
				}
			}
			ready <- true
		}()

		// weak match
		go func() {
			for _, query := range weak_queries {
				// `m` is Match
				mk := query.MatchString(item.Key)
				mv := query.MatchString(item.Value)
				me := query.MatchString(item.Extra)
				mi := query.MatchString(fmt.Sprint(item.Id))

				if mk || mv || me || mi {
					weak_filtered_items = append(weak_filtered_items, item)
					break // if once happens -> done
				}
			}

			ready <- true
		}()

		timeout := time.NewTicker(time.Second * 5)

		ready_counter := 0
	WaitInSelect:
		for {
			select {
			// time has passed but not all goroutines were done -> skip them
			case <-timeout.C:
				timeout.Stop()
				break WaitInSelect
			// it gets value from reserved query, if it's not matched -> skip as item doesn't match (strictly)
			case ok := <-is_ok:
				if !ok {
					continue ItemsSort
				}

				exact_match = append(exact_match, item)

			// if ready == 3 -> all goroutines were done -> break and execute next code
			case <-ready:
				ready_counter++
				if ready_counter == goroutines {
					break WaitInSelect
				}
			}
		}
	}

	simple_filtered_items := append(strong_filtered_items, weak_filtered_items...)

	return onlyUniqueItems(append(exact_match, simple_filtered_items...))
}

func GetPart[T any](items []T, part int, part_size ...int) []T {
	if len(items) == 0 {
		return []T{}
	}

	// default value
	part_s := 50
	// if other is provided -> use it
	if len(part_size) > 0 {
		part_s = part_size[0]
	}

	return saveSlice(items, part*part_s, (part+1)*part_s)
}

func saveSlice[T any](arr []T, from, to int) []T {
	if to > len(arr) {
		to = len(arr)
	}

	if from > len(arr)-1 {
		return []T{}
	}

	return arr[from:to]
}

func onlyUniqueItems(items []*hashmap.Item) []*hashmap.Item {
	val := make(map[uint32]bool)
	unique_items := make([]*hashmap.Item, 0, len(items))

	for _, item := range items {
		if _, ok := val[item.Id]; !ok {
			unique_items = append(unique_items, item)
			val[item.Id] = true
		}
	}

	return unique_items
}
