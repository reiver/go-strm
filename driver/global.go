package strmdriver


// Global variable that function as the registries for
// Beginner's, Strmer's and Ender's.
//
// To get things into these registries the RegisterBeginner,
// RegisterStrmer and RegisterEnder funcs are called.
//
// To get things out of these registries the GetBeginner,
// GetStrmer and GetEnder funcs are called.
var (
	globalBeginners map[string]Beginner = map[string]Beginner{}
	globalStrmers   map[string]Strmer   = map[string]Strmer{}
	globalEnders    map[string]Ender    = map[string]Ender{}
)
