package server

type MDXFile struct {
	Meta Meta
	Body string
	Err  error
}

type Meta struct {
	Slug        string
	Institution string
	Location    string
	Country     string
	StartDate   string
	EndDate     string
	Position    string
	Occupation  string
	Err         error
}
