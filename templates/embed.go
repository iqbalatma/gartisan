package templates

import _ "embed"

//go:embed controller.tmpl
var ControllerTmpl string

//go:embed model.tmpl
var ModelTmpl string

//go:embed repository.tmpl
var RepositoryTmpl string

//go:embed service.tmpl
var ServiceTmpl string
