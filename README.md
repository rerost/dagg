# Dagg 
For dag expression tool

![Screenshot from 2019-07-17 23-00-05](https://user-images.githubusercontent.com/5201588/61381933-c6a29900-a8e6-11e9-968f-105215d714e8.png)

## Install
```
go get github.com/rerost/dagg/cmd/dagg
```

## Example
sample-dag.yaml
```yaml
name: sample
option:
  team: rerost
jobs:
  - name: jobA
    command: "echo A"
    option:
      repo: repo
  - name: jobB
    command: "echo B"
    dependencies: [jobA]
    option:
      repo: repo
  - name: jobC
    command: "echo C"
    dependencies: [jobB]
    option:
      repo: repo
  - name: jobD
    command: "echo D"
    dependencies: [jobB, jobC]
    option:
      repo: repo
```

sample-template.tpl
```yaml
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: {{index .Option "team"}}-{{.Name}}-
spec:
  entrypoint: {{index .Option "team"}}-{{.Name}}-dag
  templates:
  - name: {{index .Option "team"}}-{{.Name}}-dag
    dag:
      tasks:
      {{- range .Jobs}}
      - name: {{.Name}}
        template: template-{{.Name}}
        {{- $deplen := len .Dependencies}}
        {{- if ne 0 $deplen}}
        dependencies: [{{- range $index, $var := .Dependencies}}{{- if ne $index 0}}, {{- end}}{{$var}}{{- end}}]
        {{- end}}
      {{- end}}

  {{- range .Jobs}}
  - name: template-{{.Name}}
    container:
      image: rerost/{{index .Option "repo"}}:latest
      command: ["{{.Command}}"]
  {{- end}}
```

```bash
$ dagg gen sample-dag.yaml sample-template.tpl
apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: rerost-sample-
spec:
  entrypoint: rerost-sample-dag
  templates:
  - name: rerost-sample-dag
    dag:
      tasks:
      - name: jobA
        template: template-jobA
      - name: jobB
        template: template-jobB
        dependencies: [jobA]
      - name: jobC
        template: template-jobC
        dependencies: [jobB]
      - name: jobD
        template: template-jobD
        dependencies: [jobB,jobC]
  - name: template-jobA
    container:
      image: rerost/repo:latest
      command: ["echo A"]
  - name: template-jobB
    container:
      image: rerost/repo:latest
      command: ["echo B"]
  - name: template-jobC
    container:
      image: rerost/repo:latest
      command: ["echo C"]
  - name: template-jobD
    container:
      image: rerost/repo:latest
      command: ["echo D"]
```

## Commands
- [x] `dagg gen`
  - Generate DAG instance by template from DAG file
  - e.g.) DAG file -> Argo manifest
- [ ] `dagg draw`
  - Draw DAG
- [ ] `dagg reverse`
  - Reverse DAG instance to DAG file
  - e.g.) Argo manifest -> DAG file
