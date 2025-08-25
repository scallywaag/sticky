package flags

type Flags struct {
	// operations
	List string
	Add  string
	Del  int
	Mut  int

	// mutations
	Pin    bool
	Cross  bool
	Red    bool
	Green  bool
	Yellow bool
	Blue   bool

	// lists
	ListLists bool
	AddList   string
	DelList   int
}
