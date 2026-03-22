{{ if ne .Style "go_zero" }}style: {{.Style}}

{{ end }}gen:
    hooks:
        after:
            - pzero gen swagger