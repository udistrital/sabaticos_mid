package helpers

// getIntPtr extrae un valor int de un map y retorna *int
func GetIntPtr(m map[string]interface{}, key string) *int {
	if val, ok := m[key]; ok {
		if floatVal, ok := val.(float64); ok {
			intVal := int(floatVal)
			return &intVal
		}
	}
	return nil
}

// InterfaceToInt convierte un interface{} a int. Maneja float64 e int.
// Retorna el valor convertido o defaultValue si no se puede convertir.
func InterfaceToInt(val interface{}, defaultValue int) int {
	if val == nil {
		return defaultValue
	}
	switch v := val.(type) {
	case float64:
		return int(v)
	case int:
		return v
	default:
		return defaultValue
	}
}

// InterfaceToIntPtr convierte un interface{} a *int. Maneja float64 e int.
// Retorna nil si val es nil o no se puede convertir.
func InterfaceToIntPtr(val interface{}) *int {
	if val == nil {
		return nil
	}
	switch v := val.(type) {
	case float64:
		intVal := int(v)
		return &intVal
	case int:
		return &v
	default:
		return nil
	}
}
