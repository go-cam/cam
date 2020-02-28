package camModels

// application's config struct
type AppConfig struct {
	// The component name of the default components.Database
	DefaultDBName string
	// default xorm template's file relative path.
	// Deprecated: remove on v0.3.0
	DefaultTemplatesDir string
}
