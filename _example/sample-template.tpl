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

