package save

type SaveData struct {
	UnlockedLevel int
	LevelStars    map[int]int

	// Settings SettingsData

	TotalDeaths int
	TotalStars  int

	CompletedTutorial bool
}

func SaveGame(data SaveData) {

}

func LoadGame() SaveData {
	return SaveData{}
}
