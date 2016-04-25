package data

const (
	JARVIS_USER_ID = "jarvis-user-id"
)

func JarvisUserId() string {
	_, id := Get(JARVIS_USER_ID)
	return id
}
