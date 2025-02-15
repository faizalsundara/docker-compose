package dto

func ResponseFailed(msg string, err interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "error",
		"message": msg,
		"error":   err,
	}
}

func ResponseSuccesNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func ResponseSuccesWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}
