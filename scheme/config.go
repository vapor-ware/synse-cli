package scheme

type Config struct {
	Locale     string                 `json:"locale",yaml:"locale"`
	PrettyJSON bool                   `json:"pretty_json",yaml:"pretty_json"`
	Logging    string                 `json:"logging",yaml:"logging"`
	Cache      map[string]interface{} `json:"cache",yaml:"cache"`
	GRPC       map[string]interface{} `json:"grpc",yaml:"grpc"`
}
