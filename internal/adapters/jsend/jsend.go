package jsend

func Success(data interface{}) map[string]interface{} {
    return map[string]interface{}{
        "status": "success",
        "data":   data,
    }
}

func Fail(data interface{}) map[string]interface{} {
    return map[string]interface{}{
        "status": "fail",
        "data":   data,
    }
}

func Error(message string) map[string]interface{} {
    return map[string]interface{}{
        "status":  "error",
        "message": message,
    }
}