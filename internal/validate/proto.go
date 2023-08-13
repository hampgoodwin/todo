package validate

var (
	toDo = map[string]string{
		"Id":        "required,KSUID",
		"Message":   "required",
		"CreatedAt": "required",
	}
	listToDoRequest = map[string]string{
		"Ids":      "omitempty,dive,KSUID",
		"PageSize": "omitempty,number",
		// "PageToken": "omitempty,number",
	}
)
