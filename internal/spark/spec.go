package spark

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

// Node represents the structure of your n8n custom node as defined in the YAML file.
//
//	Adjust this struct to match your YAML file's structure.
type Node struct {
	Name        string `yaml:"name"`
	Description string `yaml:"description"`
	// Add other fields as needed...
}

// ReadNodeSpec reads the YAML configuration file and returns a Node struct.
func ReadNodeSpec(filePath string) (*Node, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read YAML file: %w", err)
	}

	var node Node
	err = yaml.Unmarshal(data, &node)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal YAML: %w", err)
	}

	return &node, nil
}

// Example usage
func main() {
	node, err := ReadNodeSpec("path/to/your/config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Node Name: %s\n", node.Name)
	fmt.Printf("Node Description: %s\n", node.Description)
}
