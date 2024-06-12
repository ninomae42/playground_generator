package main

import (
	"embed"
	"fmt"
	"html/template"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

const (
	taskFileTemplatePath    = "templates/taskfile.yaml.tmpl"
	composeFileTemplatePath = "templates/compose.yaml.tmpl"
)

// GenerateTemplateFiles generates the template files for the project.
func GenerateTemplateFiles(embed embed.FS, cfg Config) error {
	tmpl, err := template.ParseFS(embed, taskFileTemplatePath, composeFileTemplatePath)
	if err != nil {
		return fmt.Errorf("failed to parse templates: %v", err)
	}

	outFiles, err := createTemplateOutFiles(cfg.DestPath, taskFileTemplatePath, composeFileTemplatePath)
	if err != nil {
		return fmt.Errorf("failed to create output files: %v", err)
	}

	err = executeTemplates(tmpl, cfg.ProjectName, outFiles)
	if err != nil {
		return fmt.Errorf("failed to execute templates: %v", err)
	}
	return nil
}

type outFile map[string]*os.File

func createTemplateOutFiles(destPath string, paths ...string) (outFile, error) {
	files := make(outFile, len(paths))

	for _, p := range paths {
		baseName := filepath.Base(p)
		extension := filepath.Ext(baseName)

		filename := strings.TrimSuffix(baseName, extension)
		filename = filepath.Join(destPath, filename)

		f, err := os.Create(filename)
		if err != nil {
			return nil, err
		}
		files[baseName] = f
	}

	return files, nil
}

func executeTemplates(tmpl *template.Template, projectName string, files outFile) error {
	data := struct {
		ProjectName string
	}{
		ProjectName: projectName,
	}

	for name, file := range files {
		err := tmpl.ExecuteTemplate(file, name, data)
		if err != nil {
			return err
		}
	}
	return nil
}

// DuplicateFromEmbed duplicates the files from the embed.FS to the destination directory.
func DuplicateFromEmbed(embed embed.FS, destDir string) error {
	err := fs.WalkDir(embed, "base_codes", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		relativePath, err := filepath.Rel("base_codes", path)
		if err != nil {
			return fmt.Errorf("error getting relative path: %w", err)
		}

		newPath := filepath.Join(destDir, relativePath)
		if d.IsDir() {
			err := os.MkdirAll(newPath, 0755)
			if err != nil {
				return fmt.Errorf("error creating directory %s: %w", newPath, err)
			}
		} else {
			data, err := embed.ReadFile(path)
			if err != nil {
				return fmt.Errorf("error reading embedded file %s: %w", path, err)
			}
			err = os.WriteFile(newPath, data, 0644)
			if err != nil {
				return fmt.Errorf("error writing file %s: %w", newPath, err)
			}

		}
		return nil
	})
	return err
}
