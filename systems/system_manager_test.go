package systems

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"BasicECS/core"
)

const testComponentName = "TestComponent"

//test component
type testComponent struct {
	val int
}

func (c testComponent) GetComponentName() string { return getTestComponentName() }

func getTestComponentName() string {return testComponentName }

//test system
type testSystem struct {}

func (s testSystem) Update(dt float64, entityManager *core.EntityManager) {

	entityIds := entityManager.GetAllEntitiesPossessingComponentsOfClass(getTestComponentName())

	for _, entityId := range entityIds {

		testCompp := entityManager.GetComponentOfClass(   // <-- also all of this stuff ? wtf
			getTestComponentName(),
			entityId,
		).(*testComponent)

		testCompp.val++
	}
}

//tests

func TestCreateSystemManager(t *testing.T) {
	systemManager := CreateSystemManager()
	assert.NotNil(t, systemManager)
	assert.Equal(t, 0, len(systemManager.systems))
}

func TestSystemManagerAddSystem(t *testing.T) {
	testSys := testSystem{}
	systemManager := CreateSystemManager()
	systemManager.AddSystem(testSys)

	assert.Equal(t, 1, len(systemManager.systems))

	systemManager.AddSystem(testSys)

	assert.Equal(t, 2, len(systemManager.systems))
}

func TestSystemManagerProcessSystems(t *testing.T) {

	entityManager := core.CreateEntityManager(10)
	testEntity := entityManager.CreateEntity()
	entityManager.AddComponentToEntity(testEntity.Id, &testComponent{val:0})

	testSys := testSystem{}
	systemManager := CreateSystemManager()
	systemManager.AddSystem(testSys)

	systemManager.ProcessSystems(1, &entityManager)

	testCompp := entityManager.GetComponentOfClass(
		getTestComponentName(),
		testEntity.Id,
	).(*testComponent)

	assert.Equal(t, 1, testCompp.val)

	systemManager.ProcessSystems(1, &entityManager)

	assert.Equal(t, 2, testCompp.val)
}