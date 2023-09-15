package inject

var gDataContainer DataContainer = &nilDataContainer{}

func InjDataContainer(dataContainer DataContainer) {
	gDataContainer = dataContainer
}

func GetDataContainer() DataContainer {
	return gDataContainer
}

var gLinker Linker = &nilLinker{}

func InjLinker(linker Linker) {
	gLinker = linker
}

func GetLinker() Linker {
	return gLinker
}
