// Code generated by protoc-gen-go-enums. DO NOT EDIT.
// source: chargehive/chtype/response_codes.proto

package chtype

const (
	RESPONSE_CATEGORY_INVALID       = ResponseCategory_RESPONSE_CATEGORY_INVALID
	RESPONSE_CATEGORY_CHARGE        = ResponseCategory_RESPONSE_CATEGORY_CHARGE
	RESPONSE_CATEGORY_METHOD        = ResponseCategory_RESPONSE_CATEGORY_METHOD
	RESPONSE_CATEGORY_PERSON        = ResponseCategory_RESPONSE_CATEGORY_PERSON
	RESPONSE_CATEGORY_CONFIGURATION = ResponseCategory_RESPONSE_CATEGORY_CONFIGURATION
	RESPONSE_CATEGORY_CONNECTIVITY  = ResponseCategory_RESPONSE_CATEGORY_CONNECTIVITY
	RESPONSE_CATEGORY_FRAUD         = ResponseCategory_RESPONSE_CATEGORY_FRAUD
	RESPONSE_CATEGORY_VERIFICATION  = ResponseCategory_RESPONSE_CATEGORY_VERIFICATION
	RESPONSE_CATEGORY_PROCESSING    = ResponseCategory_RESPONSE_CATEGORY_PROCESSING
	RESPONSE_CATEGORY_UNKNOWN       = ResponseCategory_RESPONSE_CATEGORY_UNKNOWN
	RESPONSE_CATEGORY_REQUEST       = ResponseCategory_RESPONSE_CATEGORY_REQUEST
)

const (
	RESPONSE_ERROR_INVALID         = ResponseError_RESPONSE_ERROR_INVALID
	RESPONSE_ERROR_NONE            = ResponseError_RESPONSE_ERROR_NONE
	RESPONSE_ERROR_AVAILABLE_FUNDS = ResponseError_RESPONSE_ERROR_AVAILABLE_FUNDS
	RESPONSE_ERROR_PAYLOAD         = ResponseError_RESPONSE_ERROR_PAYLOAD
	RESPONSE_ERROR_LIMIT           = ResponseError_RESPONSE_ERROR_LIMIT
	RESPONSE_ERROR_EXPIRED         = ResponseError_RESPONSE_ERROR_EXPIRED
	RESPONSE_ERROR_UNAVAILABLE     = ResponseError_RESPONSE_ERROR_UNAVAILABLE
	RESPONSE_ERROR_UNSUPPORTED     = ResponseError_RESPONSE_ERROR_UNSUPPORTED
	RESPONSE_ERROR_LOST            = ResponseError_RESPONSE_ERROR_LOST
	RESPONSE_ERROR_STOLEN          = ResponseError_RESPONSE_ERROR_STOLEN
	RESPONSE_ERROR_FRAUD           = ResponseError_RESPONSE_ERROR_FRAUD
	RESPONSE_ERROR_PICKUP          = ResponseError_RESPONSE_ERROR_PICKUP
	RESPONSE_ERROR_VELOCITY        = ResponseError_RESPONSE_ERROR_VELOCITY
	RESPONSE_ERROR_ADDRESS         = ResponseError_RESPONSE_ERROR_ADDRESS
	RESPONSE_ERROR_DUPLICATE       = ResponseError_RESPONSE_ERROR_DUPLICATE
	RESPONSE_ERROR_TIMEOUT         = ResponseError_RESPONSE_ERROR_TIMEOUT
	RESPONSE_ERROR_NOT_FOUND       = ResponseError_RESPONSE_ERROR_NOT_FOUND
	RESPONSE_ERROR_DISPUTED        = ResponseError_RESPONSE_ERROR_DISPUTED
	RESPONSE_ERROR_PERMISSION      = ResponseError_RESPONSE_ERROR_PERMISSION
	RESPONSE_ERROR_DECLINE         = ResponseError_RESPONSE_ERROR_DECLINE
	RESPONSE_ERROR_USER_INPUT      = ResponseError_RESPONSE_ERROR_USER_INPUT
	RESPONSE_ERROR_USER_DEVICE     = ResponseError_RESPONSE_ERROR_USER_DEVICE
	RESPONSE_ERROR_ALREADY_DONE    = ResponseError_RESPONSE_ERROR_ALREADY_DONE
	RESPONSE_ERROR_RETRY           = ResponseError_RESPONSE_ERROR_RETRY
	RESPONSE_ERROR_QUEUED          = ResponseError_RESPONSE_ERROR_QUEUED
	RESPONSE_ERROR_SYSTEM          = ResponseError_RESPONSE_ERROR_SYSTEM
	RESPONSE_ERROR_UNKNOWN         = ResponseError_RESPONSE_ERROR_UNKNOWN
	RESPONSE_ERROR_NOT_READY       = ResponseError_RESPONSE_ERROR_NOT_READY
	RESPONSE_ERROR_DISABLED        = ResponseError_RESPONSE_ERROR_DISABLED
	RESPONSE_ERROR_CVV             = ResponseError_RESPONSE_ERROR_CVV
)

const (
	RESPONSE_CODE_INVALID                             = ResponseCode_RESPONSE_CODE_INVALID
	RESPONSE_CODE_INFO                                = ResponseCode_RESPONSE_CODE_INFO
	RESPONSE_CODE_OK                                  = ResponseCode_RESPONSE_CODE_OK
	RESPONSE_CODE_OK_PARTIAL                          = ResponseCode_RESPONSE_CODE_OK_PARTIAL
	RESPONSE_CODE_REDIRECT                            = ResponseCode_RESPONSE_CODE_REDIRECT
	RESPONSE_CODE_REDIRECT_STILL_PROCESSING           = ResponseCode_RESPONSE_CODE_REDIRECT_STILL_PROCESSING
	RESPONSE_CODE_CLIENT                              = ResponseCode_RESPONSE_CODE_CLIENT
	RESPONSE_CODE_CLIENT_PAYLOAD                      = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD
	RESPONSE_CODE_CLIENT_PAYLOAD_AMOUNT               = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_AMOUNT
	RESPONSE_CODE_CLIENT_PAYLOAD_CARD                 = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_CARD
	RESPONSE_CODE_CLIENT_PAYLOAD_CARD_NUMBER          = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_CARD_NUMBER
	RESPONSE_CODE_CLIENT_PAYLOAD_CARD_DATE            = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_CARD_DATE
	RESPONSE_CODE_CLIENT_PAYLOAD_CARD_CVV             = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_CARD_CVV
	RESPONSE_CODE_CLIENT_PAYLOAD_CARD_TYPE            = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_CARD_TYPE
	RESPONSE_CODE_CLIENT_PAYLOAD_ADDRESS              = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_ADDRESS
	RESPONSE_CODE_CLIENT_PAYLOAD_ZIP                  = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_ZIP
	RESPONSE_CODE_CLIENT_PAYLOAD_CURRENCY             = ResponseCode_RESPONSE_CODE_CLIENT_PAYLOAD_CURRENCY
	RESPONSE_CODE_CLIENT_PAYMENT_METHOD               = ResponseCode_RESPONSE_CODE_CLIENT_PAYMENT_METHOD
	RESPONSE_CODE_CLIENT_PAYMENT_METHOD_FUNDS         = ResponseCode_RESPONSE_CODE_CLIENT_PAYMENT_METHOD_FUNDS
	RESPONSE_CODE_CLIENT_PAYMENT_METHOD_LIMIT         = ResponseCode_RESPONSE_CODE_CLIENT_PAYMENT_METHOD_LIMIT
	RESPONSE_CODE_CLIENT_PAYMENT_METHOD_NOT_PERMITTED = ResponseCode_RESPONSE_CODE_CLIENT_PAYMENT_METHOD_NOT_PERMITTED
	RESPONSE_CODE_CLIENT_SECURITY                     = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY
	RESPONSE_CODE_CLIENT_SECURITY_FRAUD               = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_FRAUD
	RESPONSE_CODE_CLIENT_SECURITY_LOST                = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_LOST
	RESPONSE_CODE_CLIENT_SECURITY_STOLEN              = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_STOLEN
	RESPONSE_CODE_CLIENT_SECURITY_PICKUP              = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_PICKUP
	RESPONSE_CODE_CLIENT_SECURITY_CONTACT_ISSUER      = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_CONTACT_ISSUER
	RESPONSE_CODE_CLIENT_SECURITY_SUSPECTED_FRAUD     = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_SUSPECTED_FRAUD
	RESPONSE_CODE_CLIENT_SECURITY_VELOCITY            = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_VELOCITY
	RESPONSE_CODE_CLIENT_SECURITY_AVS                 = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_AVS
	RESPONSE_CODE_CLIENT_SECURITY_CVV                 = ResponseCode_RESPONSE_CODE_CLIENT_SECURITY_CVV
	RESPONSE_CODE_CLIENT_TRANS                        = ResponseCode_RESPONSE_CODE_CLIENT_TRANS
	RESPONSE_CODE_CLIENT_TRANS_DECLINED               = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_DECLINED
	RESPONSE_CODE_CLIENT_TRANS_FAILED                 = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_FAILED
	RESPONSE_CODE_CLIENT_TRANS_PERMISSION             = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_PERMISSION
	RESPONSE_CODE_CLIENT_TRANS_NOT_FOUND              = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_NOT_FOUND
	RESPONSE_CODE_CLIENT_TRANS_DUPLICATE              = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_DUPLICATE
	RESPONSE_CODE_CLIENT_TRANS_RETRY                  = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_RETRY
	RESPONSE_CODE_CLIENT_TRANS_EXPIRED                = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_EXPIRED
	RESPONSE_CODE_CLIENT_TRANS_DISPUTE                = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_DISPUTE
	RESPONSE_CODE_CLIENT_TRANS_TIMEOUT                = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_TIMEOUT
	RESPONSE_CODE_CLIENT_TRANS_ALREADY_DONE           = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_ALREADY_DONE
	RESPONSE_CODE_CLIENT_TRANS_NOT_READY              = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_NOT_READY
	RESPONSE_CODE_CLIENT_TRANS_PAYMENT_METHOD         = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_PAYMENT_METHOD
	RESPONSE_CODE_CLIENT_TRANS_EXCEEDS_LIMIT          = ResponseCode_RESPONSE_CODE_CLIENT_TRANS_EXCEEDS_LIMIT
	RESPONSE_CODE_CLIENT_3DSEC                        = ResponseCode_RESPONSE_CODE_CLIENT_3DSEC
	RESPONSE_CODE_CLIENT_3DSEC_AUTHENTICATION         = ResponseCode_RESPONSE_CODE_CLIENT_3DSEC_AUTHENTICATION
	RESPONSE_CODE_CONN                                = ResponseCode_RESPONSE_CODE_CONN
	RESPONSE_CODE_CONN_CHARGEHIVE                     = ResponseCode_RESPONSE_CODE_CONN_CHARGEHIVE
	RESPONSE_CODE_CONN_CONFIG_AUTH                    = ResponseCode_RESPONSE_CODE_CONN_CONFIG_AUTH
	RESPONSE_CODE_CONN_INVALID                        = ResponseCode_RESPONSE_CODE_CONN_INVALID
	RESPONSE_CODE_CONN_UNAVAILABLE                    = ResponseCode_RESPONSE_CODE_CONN_UNAVAILABLE
	RESPONSE_CODE_CONN_ERROR                          = ResponseCode_RESPONSE_CODE_CONN_ERROR
	RESPONSE_CODE_CONN_LIMIT_EXCEEDED                 = ResponseCode_RESPONSE_CODE_CONN_LIMIT_EXCEEDED
	RESPONSE_CODE_CONN_UNSUPPORTED_METHOD             = ResponseCode_RESPONSE_CODE_CONN_UNSUPPORTED_METHOD
	RESPONSE_CODE_CONN_TIMEOUT                        = ResponseCode_RESPONSE_CODE_CONN_TIMEOUT
	RESPONSE_CODE_UNKNOWN                             = ResponseCode_RESPONSE_CODE_UNKNOWN
)