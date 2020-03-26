package config

const (
	MODE_ALLOCATE = "ALLOCATE"
	MODE_POOL     = "POOL"
)

type Config struct {
	Replicas      int
	Limit         int
	Ip            string
	Port          string
	Mode          string
	Redis_ADDRESS string
	Redis_PWD     string
}
