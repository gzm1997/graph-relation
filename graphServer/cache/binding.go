package cache


type BindingKey int

const (
	CreateFileLink BindingKey = iota
	ClickFileLink
	CreateGroupShareLink
	ClickGroupShareLink
	File
	User
	Group
)

var AllBindingKeys = [...]BindingKey{CreateFileLink, ClickFileLink, CreateGroupShareLink, ClickGroupShareLink, User, File, Group}

func (this BindingKey) String() string {
	switch this {
	case CreateFileLink:
		return "CreateFileLink"
	case ClickFileLink:
		return "ClickFileLink"
	case CreateGroupShareLink:
		return "CreateGroupShareLink"
	case ClickGroupShareLink:
		return "ClickGroupShareLink"
	case User:
		return "User"
	case File:
		return "File"
	case Group:
		return "Group"
	default:
		return "unknown"
	}
}