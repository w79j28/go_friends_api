package output

type Output struct {
	Success bool `json:"success" description:"desc"`
}

var SUCCESS = Output{true}
var FAILED = Output{false}
