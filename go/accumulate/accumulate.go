// Package accumulate provides an accumulate operation that applies an operation
// on each element of a collection.
package accumulate

const testVersion = 1

// Accumulate applies the given operation on each element of the collection and
// return the resulting collection (of the same size).
func Accumulate(collection []string, operation func(string) string) []string {
	var result []string
	for _, element := range collection {
		result = append(result, operation(element))
	}

	return result
}
