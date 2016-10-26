package cli

import (
	"io/ioutil"
	"github.com/ghodss/yaml"
	"github.com/tidwall/gjson"
)

func ParseFile(path string) (Project, error) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return Project{}, err
	} else {
		return ParseYaml(bytes)
	}
}

func ParseYaml(bytes []byte) (Project, error) {
	jsonbytes, err := yaml.YAMLToJSON(bytes)
	if err != nil {
		return Project{}, err
	}

	json := gjson.Parse(string(jsonbytes))
	return ParseProject(json), nil
}

func ParseProject(json gjson.Result) (Project) {
	p := Project{}
	if j := json.Get("project"); j.Exists() {
		p.Name = j.String()
	}
	if j := json.Get("desc"); j.Exists() {
		p.Desc = j.String()
	}
	if j := json.Get("includes"); j.Exists() {
		for _, s := range j.Array() {
			p.Includes = append(p.Includes, s.String())
		}
	}
	if j := json.Get("extends"); j.Exists() {
		for _, s := range j.Array() {
			p.Extends = append(p.Extends, s.String())
		}
	}
	if j := json.Get("env"); j.Exists() {
		p.Env = make(map[string]string)
		for k, v := range j.Map() {
			p.Env[k] = v.String()
		}
	}
	if j := json.Get("env_files"); j.Exists() {
		for _, s := range j.Array() {
			p.EnvFiles = append(p.EnvFiles, s.String())
		}
	}
	if j := json.Get("tags"); j.Exists() {
		for _, s := range j.Array() {
			p.Tags = append(p.Tags, s.String())
		}
	}
	if j := json.Get("tasks"); j.Exists() {
		for k, v := range j.Map() {
			p.Tasks = append(p.Tasks, ParseTask(k, v))
		}
	}
	return p
}

func ParseTask(name string, json gjson.Result) (Task) {
	t := Task{}
	t.Name = name

	if j := json.Get("desc"); j.Exists() {
		t.Desc = j.String()
	}
	if j := json.Get("cmd"); j.Exists() {
		t.Cmd = j.String()
	}
	if j := json.Get("before"); j.Exists() {
		for _, s := range j.Array() {
			t.Before = append(t.Before, s.String())
		}
	}
	if j := json.Get("after"); j.Exists() {
		for _, s := range j.Array() {
			t.After = append(t.After, s.String())
		}
	}
	return t
}
