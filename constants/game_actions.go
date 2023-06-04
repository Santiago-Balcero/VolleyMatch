package constants

const (
	Attack          string = "a"
	Block           string = "b"
	Service         string = "s"
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
	Block,
	Service,
	OpponentError,
	OpponentAttack,
	OpponentBlock,
	OpponentService,
	Error,
	Exit,
	RollBack,
}
