package kvdb

// DB - БД "ключ-значение"
type DB map[string]string

func New() DB {
	return make(map[string]string)
}

func (db DB) GET(key string) string {
	if val, ok := db[key]; ok {
		return val
	}
	return ""
}

func (db DB) SET(key string, val string) {
	db[key] = val
}
