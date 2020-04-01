package models

type Module struct {
	Model
	Name string `json:"name"`
}

func ModuleExistByName(name string) bool {
	return name == ""
}

func AddModule(name string) {

}
