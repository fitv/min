package logger

type Level int

const (
	DebugLevel = Level(iota)
	InfoLevel
	WarnLevel
	ErrorLevel
)

var LevelMap = map[Level]string{
	DebugLevel: "debug",
	InfoLevel:  "info",
	WarnLevel:  "warn",
	ErrorLevel: "error",
}

func (l Level) String() string {
	return LevelMap[l]
}
