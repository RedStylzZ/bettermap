# BetterMap

## What is it?
BetterMap is a Go implementation of the Python dictionary functions.

## How To
```go
func main() {
	bMap := bettermap.NewBetterMap[string, int]()

	// Add data
	bMap.Set("dollar", 10)
	bMap.Set("euro", 12)

	// Get data
	euro := bMap.Get("euro")
	euro, exist := bMap.GetAndCheck("dollar")

	// Delete data
	bMap.Remove("euro")

	// Get Raw map[T]V
	raw := bMap.Raw()

	// Get all keys
	keys := bMap.Keys()

	// Get all values
	values := bMap.Values()
}
```