package confformat

import "fmt"

// MarshalAll takes some data stored in structs and convert them to the various data formats
func MarshalAll() error {
	t := TOMLData{
		Name: "name1",
		Age: 20,
	}

	j := JSONData{
		Name: "name2",
		Age: 30,
	}

	y := YAMLData{
		Name: "name3",
		Age: 40,
	}

	tomlRes, err := t.ToTOML()
	if err != nil {
		return err
	}
	fmt.Println("TOML Marshal = ", tomlRes.String())

	jsonRes, err := j.ToJSON()
	if err != nil {
		return err
	}
	fmt.Println("JSON Marshal = ", jsonRes.String())

	yamlRes, err := y.ToYAML()
	if err != nil {
		return err
	}
	fmt.Println("YAML Marshal = ", yamlRes.String())

	return nil
}
