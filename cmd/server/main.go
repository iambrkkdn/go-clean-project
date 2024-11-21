package main

import "github.com/iambrkkdn/go-clean-project/internal/config"

func main() {
	c, err := config.Read()
	if err != nil {
		//log.Error(context.TODO()).Err(err).Msg("failed to read config")
		return
	}

	a, err := newApplication(*c)
	if err != nil {
		//log.Fatal().Err(err).Msg("failed to initialize application")
		return
	}
	a.Run()
}
