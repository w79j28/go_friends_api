package output

// Output  JSON : {"success": bool}
type Output struct {
	Success bool `json:"success" description:"desc"`
}

// SUCCESS success
var SUCCESS = Output{true}

// FAILED failed
var FAILED = Output{false}
