package main

import (
	"bytes"
	"text/template"
)

// Helper function for building hostname.
func hostname(tpl, name string) (string, error) {
	var formatted bytes.Buffer

	vars := output{
		Name: name,
	}

	t := template.Must(template.New("name").Parse(tpl))

	err := t.Execute(&formatted, vars)
	if err != nil {
		return "", err
	}

	return formatted.String(), nil
}
