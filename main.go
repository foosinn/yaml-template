package main

import (
	"flag"
	"html/template"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type config struct {
	datafile string
	template string
}

func main() {
	c := config{}
	flag.StringVar(&c.datafile, "datafile", "./vars.yml", "variables for the template")
	flag.StringVar(&c.template, "template", "template.tpl", "go template file")

	dataFile, err := os.Open(c.datafile)
	if err != nil {
		log.Fatalf("unable to open vars file: %s", err)
	}
	var data interface{}
	if err := yaml.NewDecoder(dataFile).Decode(&data); err != nil {
		log.Fatalf("unable to load toml file: %s", err)
	}

	t, err := template.ParseFiles(c.template)
	if err != nil {
		log.Fatalf("unable to load template: %s", err)
	}

	err = t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalf("unable to render template: %s", err)
	}
}
