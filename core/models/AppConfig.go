package models

// application's config struct
type AppConfig struct {
	// The component name of the default components.Database
	DefaultDBName string
	// default xorm template's file relative path.
	DefaultTemplatesDir string
}
