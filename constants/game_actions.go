package constants

const (
	Attack          string = "a"
	AttackNeutral   string = "an"
	AttackError     string = "ae"
	Block           string = "b"
	BlockNeutral    string = "bn"
	BlockError      string = "be"
	Serve           string = "s"
	ServeNeutral    string = "sn"
	ServeError      string = "se"
	OpponentError   string = "oe"
	OpponentAttack  string = "oa"
	OpponentBlock   string = "ob"
	OpponentService string = "os"
	Error           string = "e"
	Exit            string = "x"
	PlaySet         string = "play"
	RollBack        string = "rb"
)

var SetActions = []string{
	Attack,
	AttackNeutral,
	AttackError,
	Block,
	BlockNeutral,
	BlockError,
	Serve,
	ServeNeutral,
	ServeError,
	OpponentError,
	OpponentAttack,
	OpponentBlock,
	OpponentService,
	Error,
	Exit,
	RollBack,
}
