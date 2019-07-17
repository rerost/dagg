# Dagg 
For dag expression tool

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
      repository: repo
  - name: jobB
    command: "echo B"
    dependencies: [jobA]
    option:
      repository: repo
  - name: jobC
    command: "echo C"
    dependencies: [jobB]
    option:
      repository: repo
  - name: jobD
    command: "echo D"
    dependencies: [jobB, jobC]
    option:
      repository: repo
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
      {{range .Jobs -}}
      - name: {{.Name}}
        template: template-{{.Name}}
      {{end -}}

  {{range .Jobs -}}
  - name: template-{{.Name}}
    container:
      image: rerost/{{index .Option "repo"}}
      command: ["{{.Command}}"]
  {{end -}}
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
      - name: jobC
        template: template-jobC
      - name: jobD
        template: template-jobD
      - name: template-jobA
    container:
      image: rerost/
      command: ["echo A"]
  - name: template-jobB
    container:
      image: rerost/
      command: ["echo B"]
  - name: template-jobC
    container:
      image: rerost/
      command: ["echo C"]
  - name: template-jobD
    container:
      image: rerost/
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
