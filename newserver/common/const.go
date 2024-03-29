package common

const (
	SUCCESS_CODE = 0
	SUCCESS_MESS = "Success"

	DB_ERR_CODE = 1
	DB_ERR_MESS = "DB Error"

	INVALID_INPUT_ERR_CODE = 2
	INVALID_INPUT_ERR_MESS = "Invalid Input Error"

	STATUS_ERR_CODE = 3
	STATUS_ERR_MESS = "Invalid Status Error"

	STAGE_NOT_ALLOWED_ERR_CODE = 4
	STAGE_NOT_ALLOWED_ERR_MESS = "Invalid Stage"
)

// User error code 100-199
const (
	LOGIN_ERR_CODE = 100
	LOGIn_ERR_MESS = "Login Error"
)

// TOPIC error code 200-299
const (
	TOPIC_NOT_EXIST_CODE = 200
	TOPIC_NOT_EXIST_MESS = "Topic Not Exist"

	STUDENT_INVALID_CODE = 201
	STUDENT_INVALID_MESS = "Student Invalid"
)

// Stage error code
const (
	STAGE_NOT_EXISTS_ERR_CODE = 300
	STAGE_NOT_EXISTS_ERR_MESS = "Stage Not Exists"
)

// event error code
const (
	CURRENT_EVENT_NOT_FOUND_CODE = 400
	CURRENT_EVENT_NOT_FOUND_MESS = "Current Event Not Found"
)

// JWT key
const (
	JWT_IAT_KEY        = "iat"
	JWT_EXP_KEY        = "exp"
	JWT_USER_ID_KEY    = "user_id"
	JWT_SESSION_ID_KEY = "session_id"
)
