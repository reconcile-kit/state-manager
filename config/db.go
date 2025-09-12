package config

import (
	"fmt"
	"net/url"
)

func BuildPostgresDSN() (string, error) {
	host := getEnv("DATABASE_HOST", "localhost")
	port := getEnv("DATABASE_PORT", "5432")
	user, err := mustEnv("DATABASE_USERNAME")
	if err != nil {
		return "", err
	}
	pass, err := mustEnv("DATABASE_PASSWORD")
	if err != nil {
		return "", err
	}
	name, err := mustEnv("DATABASE_NAME")
	if err != nil {
		return "", err
	}
	sslmode := getEnv("DATABASE_SSLMODE", "disable")

	u := &url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(user, pass),
		Host:   fmt.Sprintf("%s:%s", host, port),
		Path:   name,
	}

	q := u.Query()
	q.Set("sslmode", sslmode)
	u.RawQuery = q.Encode()
	return u.String(), nil
}
