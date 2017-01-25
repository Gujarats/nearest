package config

type (
	Config struct {
		Database map[string]*struct {
			Master string
			Slave  string
		}

		Redis map[string]*struct {
			Master string
			Slave  string
		}
	}
)
