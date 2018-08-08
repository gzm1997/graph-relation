package models

type Result int

const (
	Ok Result = iota
	Error
)

func (this Result) Map() map[string]string {
	switch this {
	case Ok:
		return map[string]string{"result": "ok"}
	case Error:
		return map[string]string{"result": "error"}
	default:
		return map[string]string{"result": "unknown"}
	}
}