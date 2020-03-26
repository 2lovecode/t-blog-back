package models

type Module struct {
	Name string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	CreatedBy string `json:"created_by"`
	UpdatedBy string `json:"updated_by"`
}

func ModuleExistByName(name string) bool {
	return name == ""
}

func AddModule(name string) {

}
