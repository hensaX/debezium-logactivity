package main

import "reflect"

func main() {

}

type Changes map[string]Change
type Change struct {
	Old any `json:"old"`
	New any `json:"new"`
}

func DiffBeforeAfter(before, after map[string]interface{}) Changes {
	changes := Changes{}

	for key, afterVal := range after {
		beforeVal, exists := before[key]

		if !exists {
			// field baru
			changes[key] = Change{
				Old: nil,
				New: afterVal,
			}
			continue
		}

		// JSONB (map[string]interface{})
		if isMap(beforeVal) && isMap(afterVal) {
			subChanges := diffJSON(
				beforeVal.(map[string]interface{}),
				afterVal.(map[string]interface{}),
				key,
			)
			for k, v := range subChanges {
				changes[k] = v
			}
			continue
		}

		// normal field
		if !reflect.DeepEqual(beforeVal, afterVal) {
			changes[key] = Change{
				Old: beforeVal,
				New: afterVal,
			}
		}
	}

	return changes
}

func diffJSON(
	before, after map[string]interface{},
	prefix string,
) Changes {
	changes := Changes{}

	for key, afterVal := range after {
		fullKey := prefix + "." + key
		beforeVal, exists := before[key]

		if !exists {
			changes[fullKey] = Change{
				Old: nil,
				New: afterVal,
			}
			continue
		}

		if isMap(beforeVal) && isMap(afterVal) {
			sub := diffJSON(
				beforeVal.(map[string]interface{}),
				afterVal.(map[string]interface{}),
				fullKey,
			)
			for k, v := range sub {
				changes[k] = v
			}
			continue
		}

		if !reflect.DeepEqual(beforeVal, afterVal) {
			changes[fullKey] = Change{
				Old: beforeVal,
				New: afterVal,
			}
		}
	}

	return changes
}

func isMap(v interface{}) bool {
	_, ok := v.(map[string]interface{})
	return ok
}
