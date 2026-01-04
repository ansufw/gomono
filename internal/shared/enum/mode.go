package enum

type Mode string

const (
	ModeDev  Mode = "dev"
	ModeProd Mode = "prod"
)

type Service string

const (
	ServiceAPI Service = "api"
	ServiceWeb Service = "web"
)
