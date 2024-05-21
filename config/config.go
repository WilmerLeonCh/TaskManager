package config

type Config struct {
	TaskManagerPostgresHost     string `env:"TASK_MANAGER_POSTGRES_HOST,required=true"`
	TaskManagerPostgresPort     string `env:"TASK_MANAGER_POSTGRES_PORT,required=true"`
	TaskManagerPostgresUser     string `env:"TASK_MANAGER_POSTGRES_USER,required=true"`
	TaskManagerPostgresPassword string `env:"TASK_MANAGER_POSTGRES_PASSWORD,required=true"`
	TaskManagerPostgresDB       string `env:"TASK_MANAGER_POSTGRES_DB,required=true"`
}
