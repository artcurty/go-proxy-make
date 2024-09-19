package internal

type OpenAPIPath struct {
	Methods map[string]*OpenAPIMethod `yaml:",inline"`
}

type OpenAPIMethod struct {
	Summary      string       `yaml:"summary"`
	RequestBody  RequestBody  `yaml:"requestBody"`
	ProxyMapping ProxyMapping `yaml:"proxy_mapping"`
}

type ProxyMapping struct {
	ProxyHost     string            `yaml:"proxy_host"`
	ProxyEndpoint string            `yaml:"proxy_endpoint"`
	ProxyMethod   string            `yaml:"proxy_method"`
	FieldMappings map[string]string `yaml:"field_mappings"`
}

type RequestBody struct {
	Content Content `yaml:"content"`
}

type Content struct {
	ApplicationJSON ApplicationJSON `yaml:"application/json"`
}

type ApplicationJSON struct {
	Schema Schema `yaml:"schema"`
}

type Schema struct {
	Type       string                 `yaml:"type"`
	Properties map[string]*Properties `yaml:"properties"`
}

type Properties struct {
	Type string `yaml:"type"`
}
