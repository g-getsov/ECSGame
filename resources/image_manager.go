package resources

import (
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type ImageManager struct {
	images map[string]*ebiten.Image
}

func CreateImageManager() ImageManager {
	return ImageManager{
		images: make(map[string]*ebiten.Image, 0),
	}
}

func (m *ImageManager) LoadImage(path string) *ebiten.Image {

	loadedImage, _, err := ebitenutil.NewImageFromFile(path, ebiten.FilterDefault)

	if err != nil {
		fmt.Println("Could not load image at path %s", path)
		return nil
	}

	m.images[path] = loadedImage

	fmt.Println("Successfully loaded image %s", path)

	return loadedImage
}

func (m *ImageManager) GetImageOrLoad(path string) *ebiten.Image {

	image := m.images[path]

	if image == nil {
		return m.LoadImage(path)
	}

	return image
}