package logger

type Level int

const (
	DebugLevel = Level(iota)
	InfoLevel
	WarnLevel
	ErrorLevel
)

var LevelMap = map[Level]string{
	DebugLevel: "DEBUG",
	InfoLevel:  "INFO",
	WarnLevel:  "WARN",
	ErrorLevel: "ERROR",
}

func (l Level) String() string {
	return LevelMap[l]
}
