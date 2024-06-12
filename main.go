package main

import (
	"embed"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

const (
	appName            = "pggen"
	defaultProjectName = "playground"
	defaultDestPath    = "./playground"
)

//go:embed templates/*
var TemplateFiles embed.FS

//go:embed base_codes/*
var BaseCodes embed.FS

// Config is the configuration for the generator.
type Config struct {
	// ProjectName is the name of the project.
	// used in docker compose container and network names.
	ProjectName string
	// DestPath is the path to the output directory.
	DestPath string
}

func main() {
	printMessage("generator started")
	cfg := parseFlags()

	printMessage(fmt.Sprintf("project name: %s, output directory: %s", cfg.ProjectName, cfg.DestPath))

	err := prepareOutputDirectory(cfg.DestPath)
	if err != nil {
		log.Fatalf("failed to prepare output directory: %v", err)
	}

	printMessage("generating files from templates")
	err = GenerateTemplateFiles(TemplateFiles, cfg)
	if err != nil {
		log.Fatalf("failed to generate from template files: %v", err)
	}

	printMessage("copying base code")
	err = DuplicateFromEmbed(BaseCodes, cfg.DestPath)
	if err != nil {
		log.Fatalf("failed to generate base code: %v", err)
	}

	printMessage("generator finished successfully")
}

func parseFlags() Config {
	projectname := flag.String("n", "", "project name(used in container and network names)")
	destpath := flag.String("o", "", "generator's output directory")
	flag.Parse()

	switch {
	case *projectname == "" && *destpath == "":
		*projectname = defaultProjectName
		*destpath = defaultDestPath
	case *projectname != "" && *destpath == "":
		*destpath = "./" + *projectname
	case *projectname == "" && *destpath != "":
		*projectname = filepath.Base(*destpath)
	}

	return Config{
		ProjectName: *projectname,
		DestPath:    *destpath,
	}
}

func prepareOutputDirectory(destPath string) error {
	destPath = filepath.Clean(destPath)
	if destPath == "" || destPath == "." {
		return nil
	}
	printMessage(fmt.Sprintf("creating directory \"%s\"", destPath))
	err := os.MkdirAll(destPath, 0755)
	if err != nil {
		return err
	}
	return nil
}

func printMessage(msg string) {
	fmt.Printf("%s: %s\n", appName, msg)
}
