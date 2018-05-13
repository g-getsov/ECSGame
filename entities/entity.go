package entities

type Entity struct {
	Id string
}

func CreateNewEntity(id string) Entity {
	return Entity{
		Id: id,
	}
}