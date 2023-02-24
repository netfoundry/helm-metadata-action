package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v3"
)

const ChartFile = "Chart.yaml"

func check(e error) {
	if e != nil {
		log.Fatalf("error: %v", e)
	}
}

type Chart struct {
	ApiVersion   string   `yaml:"apiVersion"`
	Name         string   `yaml:"name"`
	Version      string   `yaml:"version"`
	KubeVersion  string   `yaml:"kubeVersion"`
	Description  string   `yaml:"description"`
	Type         string   `yaml:"type"`
	Keywords     []string `yaml:"keywords"`
	Home         string   `yaml:"home"`
	Sources      []string `yaml:"sources"`
	Dependencies []*Chart `yaml:"dependencies"`
	Repository   string   `yaml:"repository"`
	Icon         string   `yaml:"icon"`
	AppVersion   string   `yaml:"appVersion"`
	Deprecated   bool     `yaml:"deprecated"`
}

func main() {
	chartPath := path.Join(os.Getenv("INPUT_PATH"), ChartFile)
	fmt.Printf("Reading values from %s\n", chartPath)

	dat, readErr := ioutil.ReadFile(chartPath)
	check(readErr)

	chart := Chart{}
	yamlErr := yaml.Unmarshal([]byte(dat), &chart)
	check(yamlErr)

	var output_lines []byte

	output_lines = append(output_lines, fmt.Sprintf("apiVersion=%s\n", chart.ApiVersion)...)
	output_lines = append(output_lines, fmt.Sprintf("name=%s\n", chart.Name)...)
	output_lines = append(output_lines, fmt.Sprintf("version=%s\n", chart.Version)...)
	output_lines = append(output_lines, fmt.Sprintf("kubeVersion=%s\n", chart.KubeVersion)...)
	output_lines = append(output_lines, fmt.Sprintf("description=%s\n", chart.Description)...)
	output_lines = append(output_lines, fmt.Sprintf("type=%s\n", chart.Type)...)
	output_lines = append(output_lines, fmt.Sprintf("keywords=%s\n", chart.Keywords)...)
	output_lines = append(output_lines, fmt.Sprintf("home=%s\n", chart.Home)...)
	output_lines = append(output_lines, fmt.Sprintf("sources=%s\n", strings.Join(chart.Sources[:], ","))...)
	output_lines = append(output_lines, fmt.Sprintf("repository=%s\n", chart.Repository)...)
	output_lines = append(output_lines, fmt.Sprintf("icon=%s\n", chart.Icon)...)
	output_lines = append(output_lines, fmt.Sprintf("appVersion=%s\n", chart.AppVersion)...)
	output_lines = append(output_lines, fmt.Sprintf("deprecated=%t\n", chart.Deprecated)...)

	for _, dep := range chart.Dependencies {
		fmt.Printf("dependencies_%s_version=%s\n", dep.Name, dep.Version)
		fmt.Printf("dependencies_%s_repository=%s\n", dep.Name, dep.Repository)
	}

	err := os.WriteFile(os.Getenv("GITHUB_OUTPUT"), output_lines, 0640)
	check(err)

}
