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

//go:embed utils/hashing.tmpl
var HashingTmpl string

//go:embed response_code.tmpl
var ResponseCodeTmpl string

//go:embed utils/http_response.tmpl
var HttpResponseTmpl string

//go:embed utils/paginate.tmpl
var PaginationTmpl string

//go:embed config/config.tmpl
var ConfigTmpl string

//go:embed config/database.tmpl
var DatabaseTmpl string

//go:embed config/logger.tmpl
var LoggerTmpl string
