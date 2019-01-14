package maps

// GetMap retorna un hash key -> value
func GetMap(name string) int {
	// maptest := make(map[string]int)
	mapTest := map[string]int{
		"juan":   18,
		"yohan":  30,
		"freddy": 100,
	}
	mapTest["llave1"] = 1
	mapTest["llave2"] = 10
	mapTest["llave3"] = 300

	delete(mapTest, "llave1")
	delete(mapTest, "juan")
	return mapTest[name]
}
