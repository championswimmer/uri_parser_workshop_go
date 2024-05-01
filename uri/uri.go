package uri

type Authority struct {
	UserInfo string
	Host     string
	Port     string
}

type URI struct {
	Scheme    string
	Authority Authority
	Path      string
	Query     map[string]string
	Fragment  string
}

func (u URI) ToString() string {
	return ""
}
