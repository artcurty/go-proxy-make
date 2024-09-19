package internal

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
	"strings"
)

func loadOpenAPIYAML(filename string) (map[string]OpenAPIPath, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var openAPI struct {
		Paths map[string]OpenAPIPath `yaml:"paths"`
	}

	err = yaml.Unmarshal(data, &openAPI)
	if err != nil {
		return nil, err
	}

	return openAPI.Paths, nil
}

func GenerateProxyFunctionForInput(inputFile, outputDir string) error {
	inputPaths, err := loadOpenAPIYAML(inputFile)
	if err != nil {
		return fmt.Errorf("erro ao carregar arquivo de input OpenAPI: %v", err)
	}

	inputBaseName := strings.TrimSuffix(filepath.Base(inputFile), ".yaml")
	outputFileName := filepath.Join(outputDir, fmt.Sprintf("generated_proxies_%s.go", inputBaseName))
	file, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de mapeamento: %v", err)
	}
	defer file.Close()

	file.WriteString("package generated\n\n")
	file.WriteString("import (\n\t\"github.com/artcurty/go-proxy-make/internal\"\n\t\"github.com/gorilla/mux\"\n\t\"net/http\"\n )\n\n")

	file.WriteString("func init() {\n")
	file.WriteString("\tinternal.AddRouteRegistration(registerUserRoutes)\n")
	file.WriteString("}\n\n")

	file.WriteString("func registerUserRoutes(r *mux.Router) {\n")

	for path, openAPIPath := range inputPaths {
		for method, _ := range openAPIPath.Methods {

			functionName := fmt.Sprintf("Proxy%s%s", strings.ToUpper(method), strings.ReplaceAll(path, "/", ""))
			file.WriteString(fmt.Sprintf("\tr.HandleFunc(\"%s\", %s).Methods(\"%s\")\n", path, functionName, strings.ToUpper(method)))
		}
	}
	file.WriteString("}\n\n")

	for path, openAPIPath := range inputPaths {
		for method, apiMethod := range openAPIPath.Methods {
			generateProxyFunction(file, path, strings.ToUpper(method), apiMethod)
		}
	}

	fmt.Println("Arquivo de funções de proxy gerado:", outputFileName)
	return nil
}

func generateProxyFunction(file *os.File, path string, method string, apiMethod *OpenAPIMethod) string {
	functionName := fmt.Sprintf("Proxy%s%s", strings.Title(method), strings.ReplaceAll(path, "/", ""))

	file.WriteString(fmt.Sprintf("func %s(w http.ResponseWriter, r *http.Request) {\n", functionName))
	file.WriteString(fmt.Sprintf("\tinternal.ProxyRequest(w, r, \"%s%s\", \"%s\", map[string]string{\n", apiMethod.ProxyMapping.ProxyHost, apiMethod.ProxyMapping.ProxyEndpoint, strings.ToUpper(apiMethod.ProxyMapping.ProxyMethod)))
	for inputField, outputField := range apiMethod.ProxyMapping.FieldMappings {
		file.WriteString(fmt.Sprintf("\t\t\"%s\": \"%s\",\n", inputField, outputField))
	}
	file.WriteString("\t})\n")
	file.WriteString("}\n\n")

	return functionName
}
