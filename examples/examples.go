package examples

// Examples - consists of the name of the example and the demonstration code
type Examples struct {
	Example []Example
}

// Example - Contains the basic structure for an example
type Example struct {
	Name        string
	SourceFiles []Source
}

// Source - An array of source files that are added to an example
type Source struct {
	Filename string
	Code     string
}

var examples Examples

func (examples *Examples) lastExample() int {
	return len(examples.Example)
}

func (example *Example) addFile(filename string, code string) {
	example.SourceFiles = append(example.SourceFiles, Source{filename, code})
}

func (examples *Examples) addExample(newExample Example) {
	examples.Example = append(examples.Example, newExample)
}

func init() {
	examples.addWebServer()
	examples.addHello()
}

// GetAllExamples returns all of the examples
func GetAllExamples() []Example {
	return examples.Example
}

// GetExample - Returns a specific example if found
func GetExample(exampleName string) *Example {
	for _, e := range examples.Example {
		if e.Name == exampleName {
			return &e
		}
	}
	return nil
}
