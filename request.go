package cocwrapper

// ReqClans is the struct which you need to send a request
type ReqClans struct {
	Name      string
	MinLevel  int
	MaxLevel  int
	Limit     int
	MinPoints int
}
