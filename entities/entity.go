package entities

type Entity struct {
	Id int
}

func CreateNewEntity(id int) Entity {
	return Entity{
		Id: id,
	}
}