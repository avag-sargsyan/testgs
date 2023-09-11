package conf

type App struct {
	HTTPAddress string `env:"GS_HTTP_ADDRESS" envDefault:":8080"`
	PackSizes   []int  `env:"GS_PACK_SIZES" envDefault:"250,500,1000,2000,5000"`
}
