package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"

	"gopkg.in/yaml.v2"
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
	fmt.Println(fmt.Sprintf(`Reading values from %s`, chartPath))

	dat, readErr := ioutil.ReadFile(chartPath)
	check(readErr)

	chart := Chart{}
	yamlErr := yaml.Unmarshal([]byte(dat), &chart)
	check(yamlErr)

	fmt.Println(fmt.Sprintf(`apiVersion=%s >> $GITHUB_OUTPUT`, chart.ApiVersion))
	fmt.Println(fmt.Sprintf(`name=%s >> $GITHUB_OUTPUT`, chart.Name))
	fmt.Println(fmt.Sprintf(`version=%s >> $GITHUB_OUTPUT`, chart.Version))
	fmt.Println(fmt.Sprintf(`kubeVersion=%s >> $GITHUB_OUTPUT`, chart.KubeVersion))
	fmt.Println(fmt.Sprintf(`description=%s >> $GITHUB_OUTPUT`, chart.Description))
	fmt.Println(fmt.Sprintf(`type=%s >> $GITHUB_OUTPUT`, chart.Type))
	fmt.Println(fmt.Sprintf(`keywords=%s >> $GITHUB_OUTPUT`, chart.Keywords))
	fmt.Println(fmt.Sprintf(`home=%s >> $GITHUB_OUTPUT`, chart.Home))
	fmt.Println(fmt.Sprintf(`sources=%s >> $GITHUB_OUTPUT`, strings.Join(chart.Sources[:], ",")))
	fmt.Println(fmt.Sprintf(`repository=%s >> $GITHUB_OUTPUT`, chart.Repository))
	fmt.Println(fmt.Sprintf(`icon=%s >> $GITHUB_OUTPUT`, chart.Icon))
	fmt.Println(fmt.Sprintf(`appVersion=%s >> $GITHUB_OUTPUT`, chart.AppVersion))
	fmt.Println(fmt.Sprintf(`deprecated=%t >> $GITHUB_OUTPUT`, chart.Deprecated))

	for _, dep := range chart.Dependencies {
		fmt.Println(fmt.Sprintf(`dependencies_%s_version=%s`, dep.Name, dep.Version))
		fmt.Println(fmt.Sprintf(`dependencies_%s_repository=%s`, dep.Name, dep.Repository))
	}
}
