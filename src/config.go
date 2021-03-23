package src

type AWSConfig struct {
	AccessKeyID     string `mapstructure:"access_key_id"`
	SecretAccessKey string `mapstructure:"secret_access_key"`
	Region          string `mapStructure:"region"`
}
type Configuration struct {
	AWS AWSConfig `mapstructure:"aws"`

	BucketName         string   `mapstructure:"bucket_name"`
	RecursiveSearch    bool     `mapstructure:"recursive"`
	DeleteWhenUploaded bool     `mapstructure:"delete_when_uploaded"`
	DryRun 			   bool     `mapstructure:"dry_run"`
	BasePath           string   `mapstructure:"base_path"`
	Paths              []string `mapstructure:"paths"`
}
