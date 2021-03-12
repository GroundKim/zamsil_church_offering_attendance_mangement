package models

type Config struct {
	DB_USERNAME string `yaml:"DB_USERNAME"`
	DB_PASSWORD string `yaml:"DB_PASSWORD"`
	DB_PORT     string `yaml:"DB_PORT"`
	DB_HOST     string `yaml:"DB_HOST"`
	DB_NAME     string `yaml:"DB_NAME"`
}

type Student struct {
	id            int
	name          string
	department_id int
	class_id      int
}
