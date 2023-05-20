package model

type Provider string

const (
	ProviderS3    Provider = "S3"
	ProviderGCS   Provider = "GCS"
	ProviderAZURE Provider = "AZURE"
)

type S3 struct {
	Bucket   string
	Region   string
	Endpoint string
	Prefix   string
}
type GCS struct {
	Bucket string
	Prefix string
}
type AZURE struct {
	Container string
	Prefix    string
}
type Storage struct {
	Provider Provider
	S3       *S3
	Gcs      *GCS
	Azure    *AZURE
}
type BackupStorage struct {
	Storage Storage
}
