package scheme

// MetaOutput defines the scheme for the data output by a "plugin meta" command.
type MetaOutput struct {
	ID       string
	Type     string
	Model    string
	Protocol string
	Rack     string
	Board    string
}
