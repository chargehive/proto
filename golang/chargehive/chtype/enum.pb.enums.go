// Code generated by protoc-gen-go-enums. DO NOT EDIT.
// source: chargehive/chtype/enum.proto

package chtype

const (
	STATUS_INVALID      = Status_STATUS_INVALID
	STATUS_PENDING      = Status_STATUS_PENDING
	STATUS_SETUP        = Status_STATUS_SETUP
	STATUS_PROVISIONING = Status_STATUS_PROVISIONING
	STATUS_ACTIVE       = Status_STATUS_ACTIVE
	STATUS_SUSPENDED    = Status_STATUS_SUSPENDED
	STATUS_CANCELLED    = Status_STATUS_CANCELLED
	STATUS_ARCHIVED     = Status_STATUS_ARCHIVED
	STATUS_DELETED      = Status_STATUS_DELETED
)

const (
	ACTOR_TYPE_INVALID           = ActorType_ACTOR_TYPE_INVALID
	ACTOR_TYPE_CONSUMER          = ActorType_ACTOR_TYPE_CONSUMER
	ACTOR_TYPE_MERCHANT          = ActorType_ACTOR_TYPE_MERCHANT
	ACTOR_TYPE_CHARGEHIVE        = ActorType_ACTOR_TYPE_CHARGEHIVE
	ACTOR_TYPE_WALLET            = ActorType_ACTOR_TYPE_WALLET
	ACTOR_TYPE_CONNECTOR         = ActorType_ACTOR_TYPE_CONNECTOR
	ACTOR_TYPE_PAYMENT_PROCESSOR = ActorType_ACTOR_TYPE_PAYMENT_PROCESSOR
	ACTOR_TYPE_PAYMENT_NETWORK   = ActorType_ACTOR_TYPE_PAYMENT_NETWORK
	ACTOR_TYPE_ACQUIRING_BANK    = ActorType_ACTOR_TYPE_ACQUIRING_BANK
	ACTOR_TYPE_ISSUING_BANK      = ActorType_ACTOR_TYPE_ISSUING_BANK
	ACTOR_TYPE_RECOVERY_AGENT    = ActorType_ACTOR_TYPE_RECOVERY_AGENT
	ACTOR_TYPE_EXCHANGE          = ActorType_ACTOR_TYPE_EXCHANGE
	ACTOR_TYPE_POLICY            = ActorType_ACTOR_TYPE_POLICY
)

const (
	HANDLER_TYPE_INVALID     = HandlerType_HANDLER_TYPE_INVALID
	HANDLER_TYPE_WEB         = HandlerType_HANDLER_TYPE_WEB
	HANDLER_TYPE_CLI         = HandlerType_HANDLER_TYPE_CLI
	HANDLER_TYPE_API         = HandlerType_HANDLER_TYPE_API
	HANDLER_TYPE_MOBILE      = HandlerType_HANDLER_TYPE_MOBILE
	HANDLER_TYPE_PASSTHROUGH = HandlerType_HANDLER_TYPE_PASSTHROUGH
)

const (
	TRANSACTION_TYPE_INVALID   = TransactionType_TRANSACTION_TYPE_INVALID
	TRANSACTION_TYPE_AUTHORIZE = TransactionType_TRANSACTION_TYPE_AUTHORIZE
	TRANSACTION_TYPE_CAPTURE   = TransactionType_TRANSACTION_TYPE_CAPTURE
	TRANSACTION_TYPE_REFUND    = TransactionType_TRANSACTION_TYPE_REFUND
	TRANSACTION_TYPE_CANCEL    = TransactionType_TRANSACTION_TYPE_CANCEL
	TRANSACTION_TYPE_DISPUTE   = TransactionType_TRANSACTION_TYPE_DISPUTE
	TRANSACTION_TYPE_VERIFY    = TransactionType_TRANSACTION_TYPE_VERIFY
	TRANSACTION_TYPE_EXPIRED   = TransactionType_TRANSACTION_TYPE_EXPIRED
	TRANSACTION_TYPE_INTERNAL  = TransactionType_TRANSACTION_TYPE_INTERNAL
)

const (
	TRANSACTION_SUB_TYPE_INVALID      = TransactionSubType_TRANSACTION_SUB_TYPE_INVALID
	TRANSACTION_SUB_TYPE_IDENTIFY     = TransactionSubType_TRANSACTION_SUB_TYPE_IDENTIFY
	TRANSACTION_SUB_TYPE_CHALLENGE    = TransactionSubType_TRANSACTION_SUB_TYPE_CHALLENGE
	TRANSACTION_SUB_TYPE_CAPTURE_AUTH = TransactionSubType_TRANSACTION_SUB_TYPE_CAPTURE_AUTH
	TRANSACTION_SUB_TYPE_NONE         = TransactionSubType_TRANSACTION_SUB_TYPE_NONE
)

const (
	PAYMENT_METHOD_TYPE_INVALID        = PaymentMethodType_PAYMENT_METHOD_TYPE_INVALID
	PAYMENT_METHOD_TYPE_CARD           = PaymentMethodType_PAYMENT_METHOD_TYPE_CARD
	PAYMENT_METHOD_TYPE_DIGITALWALLET  = PaymentMethodType_PAYMENT_METHOD_TYPE_DIGITALWALLET
	PAYMENT_METHOD_TYPE_DIRECTDEBIT    = PaymentMethodType_PAYMENT_METHOD_TYPE_DIRECTDEBIT
	PAYMENT_METHOD_TYPE_CRYPTOCURRENCY = PaymentMethodType_PAYMENT_METHOD_TYPE_CRYPTOCURRENCY
)

const (
	PAYMENT_METHOD_PROVIDER_INVALID       = PaymentMethodProvider_PAYMENT_METHOD_PROVIDER_INVALID
	PAYMENT_METHOD_PROVIDER_FORM          = PaymentMethodProvider_PAYMENT_METHOD_PROVIDER_FORM
	PAYMENT_METHOD_PROVIDER_PAYPAL        = PaymentMethodProvider_PAYMENT_METHOD_PROVIDER_PAYPAL
	PAYMENT_METHOD_PROVIDER_APPLEPAY      = PaymentMethodProvider_PAYMENT_METHOD_PROVIDER_APPLEPAY
	PAYMENT_METHOD_PROVIDER_GOOGLEPAY     = PaymentMethodProvider_PAYMENT_METHOD_PROVIDER_GOOGLEPAY
	PAYMENT_METHOD_PROVIDER_AMAZONPAY     = PaymentMethodProvider_PAYMENT_METHOD_PROVIDER_AMAZONPAY
	PAYMENT_METHOD_PROVIDER_METHOD_UPDATE = PaymentMethodProvider_PAYMENT_METHOD_PROVIDER_METHOD_UPDATE
)

const (
	INPUT_TYPE_INVALID  = PaymentMethodInputType_INPUT_TYPE_INVALID
	INPUT_TYPE_PHYSICAL = PaymentMethodInputType_INPUT_TYPE_PHYSICAL
	INPUT_TYPE_VIRTUAL  = PaymentMethodInputType_INPUT_TYPE_VIRTUAL
	INPUT_TYPE_PROXY    = PaymentMethodInputType_INPUT_TYPE_PROXY
)

const (
	CONTRACT_TYPE_INVALID              = ContractType_CONTRACT_TYPE_INVALID
	CONTRACT_TYPE_NONE                 = ContractType_CONTRACT_TYPE_NONE
	CONTRACT_TYPE_PAYMENT              = ContractType_CONTRACT_TYPE_PAYMENT
	CONTRACT_TYPE_SUBSCRIPTION_INITIAL = ContractType_CONTRACT_TYPE_SUBSCRIPTION_INITIAL
	CONTRACT_TYPE_SUBSCRIPTION_RENEWAL = ContractType_CONTRACT_TYPE_SUBSCRIPTION_RENEWAL
	CONTRACT_TYPE_ONECLICK             = ContractType_CONTRACT_TYPE_ONECLICK
)

const (
	ENVIRONMENT_INVALID  = Environment_ENVIRONMENT_INVALID
	ENVIRONMENT_BLACKBOX = Environment_ENVIRONMENT_BLACKBOX
	ENVIRONMENT_TEST     = Environment_ENVIRONMENT_TEST
	ENVIRONMENT_LIVE     = Environment_ENVIRONMENT_LIVE
	ENVIRONMENT_MOCK     = Environment_ENVIRONMENT_MOCK
)

const (
	PRODUCT_TYPE_INVALID      = ProductType_PRODUCT_TYPE_INVALID
	PRODUCT_TYPE_PRODUCT      = ProductType_PRODUCT_TYPE_PRODUCT
	PRODUCT_TYPE_SERVICE      = ProductType_PRODUCT_TYPE_SERVICE
	PRODUCT_TYPE_SUBSCRIPTION = ProductType_PRODUCT_TYPE_SUBSCRIPTION
)

const (
	SKU_TYPE_INVALID = SKUType_SKU_TYPE_INVALID
	SKU_TYPE_PRIMARY = SKUType_SKU_TYPE_PRIMARY
	SKU_TYPE_ADDON   = SKUType_SKU_TYPE_ADDON
	SKU_TYPE_UPSELL  = SKUType_SKU_TYPE_UPSELL
)

const (
	DELIVERY_TYPE_INVALID  = DeliveryType_DELIVERY_TYPE_INVALID
	DELIVERY_TYPE_NONE     = DeliveryType_DELIVERY_TYPE_NONE
	DELIVERY_TYPE_PHYSICAL = DeliveryType_DELIVERY_TYPE_PHYSICAL
	DELIVERY_TYPE_VIRTUAL  = DeliveryType_DELIVERY_TYPE_VIRTUAL
)

const (
	DELIVERY_STANDARD_INVALID  = DeliveryStandard_DELIVERY_STANDARD_INVALID
	DELIVERY_STANDARD_SAME_DAY = DeliveryStandard_DELIVERY_STANDARD_SAME_DAY
	DELIVERY_STANDARD_NEXT_DAY = DeliveryStandard_DELIVERY_STANDARD_NEXT_DAY
	DELIVERY_STANDARD_TWO_DAY  = DeliveryStandard_DELIVERY_STANDARD_TWO_DAY
	DELIVERY_STANDARD_ECONOMY  = DeliveryStandard_DELIVERY_STANDARD_ECONOMY
	DELIVERY_STANDARD_OTHER    = DeliveryStandard_DELIVERY_STANDARD_OTHER
	DELIVERY_STANDARD_NONE     = DeliveryStandard_DELIVERY_STANDARD_NONE
)

const (
	TERM_TYPE_INVALID  = TermType_TERM_TYPE_INVALID
	TERM_TYPE_ONE_TIME = TermType_TERM_TYPE_ONE_TIME
	TERM_TYPE_MINUTE   = TermType_TERM_TYPE_MINUTE
	TERM_TYPE_DAY      = TermType_TERM_TYPE_DAY
	TERM_TYPE_WEEK     = TermType_TERM_TYPE_WEEK
	TERM_TYPE_MONTH    = TermType_TERM_TYPE_MONTH
	TERM_TYPE_YEAR     = TermType_TERM_TYPE_YEAR
	TERM_TYPE_NONE     = TermType_TERM_TYPE_NONE
)

const (
	FAILURE_TYPE_INVALID  = FailureType_FAILURE_TYPE_INVALID
	FAILURE_TYPE_NONE     = FailureType_FAILURE_TYPE_NONE
	FAILURE_TYPE_SOFT     = FailureType_FAILURE_TYPE_SOFT
	FAILURE_TYPE_HARD     = FailureType_FAILURE_TYPE_HARD
	FAILURE_TYPE_RETRY    = FailureType_FAILURE_TYPE_RETRY
	FAILURE_TYPE_INTERNAL = FailureType_FAILURE_TYPE_INTERNAL
)

const (
	TRANSACTION_RESULT_INVALID         = TransactionResult_TRANSACTION_RESULT_INVALID
	TRANSACTION_RESULT_SUCCESS         = TransactionResult_TRANSACTION_RESULT_SUCCESS
	TRANSACTION_RESULT_DECLINED        = TransactionResult_TRANSACTION_RESULT_DECLINED
	TRANSACTION_RESULT_VERIFY          = TransactionResult_TRANSACTION_RESULT_VERIFY
	TRANSACTION_RESULT_PENDING         = TransactionResult_TRANSACTION_RESULT_PENDING
	TRANSACTION_RESULT_PENDING_SUCCESS = TransactionResult_TRANSACTION_RESULT_PENDING_SUCCESS
	TRANSACTION_RESULT_RETRY           = TransactionResult_TRANSACTION_RESULT_RETRY
	TRANSACTION_RESULT_EXPIRED         = TransactionResult_TRANSACTION_RESULT_EXPIRED
	TRANSACTION_RESULT_PARTIAL_SUCCESS = TransactionResult_TRANSACTION_RESULT_PARTIAL_SUCCESS
	TRANSACTION_RESULT_FAILURE         = TransactionResult_TRANSACTION_RESULT_FAILURE
	TRANSACTION_RESULT_POLLING         = TransactionResult_TRANSACTION_RESULT_POLLING
)

const (
	VERIFICATION_STATUS_INVALID     = VerificationStatus_VERIFICATION_STATUS_INVALID
	VERIFICATION_STATUS_NOT_CHECKED = VerificationStatus_VERIFICATION_STATUS_NOT_CHECKED
	VERIFICATION_STATUS_PASSED      = VerificationStatus_VERIFICATION_STATUS_PASSED
	VERIFICATION_STATUS_FAILED      = VerificationStatus_VERIFICATION_STATUS_FAILED
	VERIFICATION_STATUS_CONTINUE    = VerificationStatus_VERIFICATION_STATUS_CONTINUE
)

const (
	VERIFY_REQUEST_TYPE_INVALID   = VerifyRequestType_VERIFY_REQUEST_TYPE_INVALID
	VERIFY_REQUEST_TYPE_IDENTIFY  = VerifyRequestType_VERIFY_REQUEST_TYPE_IDENTIFY
	VERIFY_REQUEST_TYPE_CHALLENGE = VerifyRequestType_VERIFY_REQUEST_TYPE_CHALLENGE
	VERIFY_REQUEST_TYPE_NONE      = VerifyRequestType_VERIFY_REQUEST_TYPE_NONE
)

const (
	LIABILITY_INVALID  = Liability_LIABILITY_INVALID
	LIABILITY_UNKNOWN  = Liability_LIABILITY_UNKNOWN
	LIABILITY_MERCHANT = Liability_LIABILITY_MERCHANT
	LIABILITY_ISSUER   = Liability_LIABILITY_ISSUER
)

const (
	PLACEMENT_CAPABILITY_INVALID          = PlacementCapability_PLACEMENT_CAPABILITY_INVALID
	PLACEMENT_CAPABILITY_CARD_FORM        = PlacementCapability_PLACEMENT_CAPABILITY_CARD_FORM
	PLACEMENT_CAPABILITY_APPLE_PAY        = PlacementCapability_PLACEMENT_CAPABILITY_APPLE_PAY
	PLACEMENT_CAPABILITY_GOOGLE_PAY       = PlacementCapability_PLACEMENT_CAPABILITY_GOOGLE_PAY
	PLACEMENT_CAPABILITY_PAYPAL           = PlacementCapability_PLACEMENT_CAPABILITY_PAYPAL
	PLACEMENT_CAPABILITY_TOKEN            = PlacementCapability_PLACEMENT_CAPABILITY_TOKEN
	PLACEMENT_CAPABILITY_DIRECTDEBIT_FORM = PlacementCapability_PLACEMENT_CAPABILITY_DIRECTDEBIT_FORM
)