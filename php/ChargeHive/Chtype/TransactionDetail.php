<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/transaction.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chargehive.chtype.TransactionDetail</code>
 */
class TransactionDetail extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string actor_id = 1;</code>
     */
    protected $actor_id = '';
    /**
     * Generated from protobuf field <code>string actor_version = 2;</code>
     */
    protected $actor_version = '';
    /**
     * Generated from protobuf field <code>.chargehive.chtype.ActorType actor_type = 3;</code>
     */
    protected $actor_type = 0;
    /**
     * Generated from protobuf field <code>string transaction_id = 4;</code>
     */
    protected $transaction_id = '';
    /**
     *Amount sent in the request
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount requested_amount = 5;</code>
     */
    protected $requested_amount = null;
    /**
     *Amount received in the result
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount processed_amount = 6;</code>
     */
    protected $processed_amount = null;
    /**
     *Amount charged for the transaction
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount fee_estimate = 7;</code>
     */
    protected $fee_estimate = null;
    /**
     *Amount charged for the transaction
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount fee_actual = 8;</code>
     */
    protected $fee_actual = null;
    /**
     *Request Send Time
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp start_time = 9;</code>
     */
    protected $start_time = null;
    /**
     *Response Receive Time
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp end_time = 10;</code>
     */
    protected $end_time = null;
    /**
     * Generated from protobuf field <code>bool was_successful = 11;</code>
     */
    protected $was_successful = false;
    /**
     * Generated from protobuf field <code>.chargehive.chtype.Environment environment = 13;</code>
     */
    protected $environment = 0;
    /**
     * Generated from protobuf field <code>.chargehive.chtype.ResponseDetail response = 14;</code>
     */
    protected $response = null;
    /**
     * Generated from protobuf field <code>.chargehive.chtype.VerificationResult verification_result = 15;</code>
     */
    protected $verification_result = null;
    /**
     * Generated from protobuf field <code>map<string, string> additional_data = 16;</code>
     */
    private $additional_data;
    /**
     * Generated from protobuf field <code>string authorization_code = 17;</code>
     */
    protected $authorization_code = '';
    /**
     * Generated from protobuf field <code>.chargehive.chtype.Liability liability = 18;</code>
     */
    protected $liability = 0;
    /**
     * Generated from protobuf field <code>string connector_library = 19;</code>
     */
    protected $connector_library = '';
    /**
     * Generated from protobuf field <code>string connector_id = 20;</code>
     */
    protected $connector_id = '';
    /**
     * Generated from protobuf field <code>bool primary = 21;</code>
     */
    protected $primary = false;
    /**
     * Network ID returned by the PSP
     *
     * Generated from protobuf field <code>string network_id = 22;</code>
     */
    protected $network_id = '';
    /**
     * Transaction ID returned by the PSP to replace transaction_id which is ambiguous
     *
     * Generated from protobuf field <code>string psp_transaction_id = 23;</code>
     */
    protected $psp_transaction_id = '';
    /**
     * Request ID returned by PSP
     *
     * Generated from protobuf field <code>string psp_request_id = 24;</code>
     */
    protected $psp_request_id = '';
    /**
     * Generated from protobuf field <code>string issuer_name = 25;</code>
     */
    protected $issuer_name = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $actor_id
     *     @type string $actor_version
     *     @type int $actor_type
     *     @type string $transaction_id
     *     @type \ChargeHive\Chtype\Amount $requested_amount
     *          Amount sent in the request
     *     @type \ChargeHive\Chtype\Amount $processed_amount
     *          Amount received in the result
     *     @type \ChargeHive\Chtype\Amount $fee_estimate
     *          Amount charged for the transaction
     *     @type \ChargeHive\Chtype\Amount $fee_actual
     *          Amount charged for the transaction
     *     @type \Google\Protobuf\Timestamp $start_time
     *          Request Send Time
     *     @type \Google\Protobuf\Timestamp $end_time
     *          Response Receive Time
     *     @type bool $was_successful
     *     @type int $environment
     *     @type \ChargeHive\Chtype\ResponseDetail $response
     *     @type \ChargeHive\Chtype\VerificationResult $verification_result
     *     @type array|\Google\Protobuf\Internal\MapField $additional_data
     *     @type string $authorization_code
     *     @type int $liability
     *     @type string $connector_library
     *     @type string $connector_id
     *     @type bool $primary
     *     @type string $network_id
     *           Network ID returned by the PSP
     *     @type string $psp_transaction_id
     *           Transaction ID returned by the PSP to replace transaction_id which is ambiguous
     *     @type string $psp_request_id
     *           Request ID returned by PSP
     *     @type string $issuer_name
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\Transaction::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string actor_id = 1;</code>
     * @return string
     */
    public function getActorId()
    {
        return $this->actor_id;
    }

    /**
     * Generated from protobuf field <code>string actor_id = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setActorId($var)
    {
        GPBUtil::checkString($var, True);
        $this->actor_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string actor_version = 2;</code>
     * @return string
     */
    public function getActorVersion()
    {
        return $this->actor_version;
    }

    /**
     * Generated from protobuf field <code>string actor_version = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setActorVersion($var)
    {
        GPBUtil::checkString($var, True);
        $this->actor_version = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.ActorType actor_type = 3;</code>
     * @return int
     */
    public function getActorType()
    {
        return $this->actor_type;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.ActorType actor_type = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setActorType($var)
    {
        GPBUtil::checkEnum($var, \ChargeHive\Chtype\ActorType::class);
        $this->actor_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string transaction_id = 4;</code>
     * @return string
     */
    public function getTransactionId()
    {
        return $this->transaction_id;
    }

    /**
     * Generated from protobuf field <code>string transaction_id = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setTransactionId($var)
    {
        GPBUtil::checkString($var, True);
        $this->transaction_id = $var;

        return $this;
    }

    /**
     *Amount sent in the request
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount requested_amount = 5;</code>
     * @return \ChargeHive\Chtype\Amount|null
     */
    public function getRequestedAmount()
    {
        return $this->requested_amount;
    }

    public function hasRequestedAmount()
    {
        return isset($this->requested_amount);
    }

    public function clearRequestedAmount()
    {
        unset($this->requested_amount);
    }

    /**
     *Amount sent in the request
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount requested_amount = 5;</code>
     * @param \ChargeHive\Chtype\Amount $var
     * @return $this
     */
    public function setRequestedAmount($var)
    {
        GPBUtil::checkMessage($var, \ChargeHive\Chtype\Amount::class);
        $this->requested_amount = $var;

        return $this;
    }

    /**
     *Amount received in the result
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount processed_amount = 6;</code>
     * @return \ChargeHive\Chtype\Amount|null
     */
    public function getProcessedAmount()
    {
        return $this->processed_amount;
    }

    public function hasProcessedAmount()
    {
        return isset($this->processed_amount);
    }

    public function clearProcessedAmount()
    {
        unset($this->processed_amount);
    }

    /**
     *Amount received in the result
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount processed_amount = 6;</code>
     * @param \ChargeHive\Chtype\Amount $var
     * @return $this
     */
    public function setProcessedAmount($var)
    {
        GPBUtil::checkMessage($var, \ChargeHive\Chtype\Amount::class);
        $this->processed_amount = $var;

        return $this;
    }

    /**
     *Amount charged for the transaction
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount fee_estimate = 7;</code>
     * @return \ChargeHive\Chtype\Amount|null
     */
    public function getFeeEstimate()
    {
        return $this->fee_estimate;
    }

    public function hasFeeEstimate()
    {
        return isset($this->fee_estimate);
    }

    public function clearFeeEstimate()
    {
        unset($this->fee_estimate);
    }

    /**
     *Amount charged for the transaction
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount fee_estimate = 7;</code>
     * @param \ChargeHive\Chtype\Amount $var
     * @return $this
     */
    public function setFeeEstimate($var)
    {
        GPBUtil::checkMessage($var, \ChargeHive\Chtype\Amount::class);
        $this->fee_estimate = $var;

        return $this;
    }

    /**
     *Amount charged for the transaction
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount fee_actual = 8;</code>
     * @return \ChargeHive\Chtype\Amount|null
     */
    public function getFeeActual()
    {
        return $this->fee_actual;
    }

    public function hasFeeActual()
    {
        return isset($this->fee_actual);
    }

    public function clearFeeActual()
    {
        unset($this->fee_actual);
    }

    /**
     *Amount charged for the transaction
     *
     * Generated from protobuf field <code>.chargehive.chtype.Amount fee_actual = 8;</code>
     * @param \ChargeHive\Chtype\Amount $var
     * @return $this
     */
    public function setFeeActual($var)
    {
        GPBUtil::checkMessage($var, \ChargeHive\Chtype\Amount::class);
        $this->fee_actual = $var;

        return $this;
    }

    /**
     *Request Send Time
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp start_time = 9;</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getStartTime()
    {
        return $this->start_time;
    }

    public function hasStartTime()
    {
        return isset($this->start_time);
    }

    public function clearStartTime()
    {
        unset($this->start_time);
    }

    /**
     *Request Send Time
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp start_time = 9;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setStartTime($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->start_time = $var;

        return $this;
    }

    /**
     *Response Receive Time
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp end_time = 10;</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getEndTime()
    {
        return $this->end_time;
    }

    public function hasEndTime()
    {
        return isset($this->end_time);
    }

    public function clearEndTime()
    {
        unset($this->end_time);
    }

    /**
     *Response Receive Time
     *
     * Generated from protobuf field <code>.google.protobuf.Timestamp end_time = 10;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setEndTime($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->end_time = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool was_successful = 11;</code>
     * @return bool
     */
    public function getWasSuccessful()
    {
        return $this->was_successful;
    }

    /**
     * Generated from protobuf field <code>bool was_successful = 11;</code>
     * @param bool $var
     * @return $this
     */
    public function setWasSuccessful($var)
    {
        GPBUtil::checkBool($var);
        $this->was_successful = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.Environment environment = 13;</code>
     * @return int
     */
    public function getEnvironment()
    {
        return $this->environment;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.Environment environment = 13;</code>
     * @param int $var
     * @return $this
     */
    public function setEnvironment($var)
    {
        GPBUtil::checkEnum($var, \ChargeHive\Chtype\Environment::class);
        $this->environment = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.ResponseDetail response = 14;</code>
     * @return \ChargeHive\Chtype\ResponseDetail|null
     */
    public function getResponse()
    {
        return $this->response;
    }

    public function hasResponse()
    {
        return isset($this->response);
    }

    public function clearResponse()
    {
        unset($this->response);
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.ResponseDetail response = 14;</code>
     * @param \ChargeHive\Chtype\ResponseDetail $var
     * @return $this
     */
    public function setResponse($var)
    {
        GPBUtil::checkMessage($var, \ChargeHive\Chtype\ResponseDetail::class);
        $this->response = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.VerificationResult verification_result = 15;</code>
     * @return \ChargeHive\Chtype\VerificationResult|null
     */
    public function getVerificationResult()
    {
        return $this->verification_result;
    }

    public function hasVerificationResult()
    {
        return isset($this->verification_result);
    }

    public function clearVerificationResult()
    {
        unset($this->verification_result);
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.VerificationResult verification_result = 15;</code>
     * @param \ChargeHive\Chtype\VerificationResult $var
     * @return $this
     */
    public function setVerificationResult($var)
    {
        GPBUtil::checkMessage($var, \ChargeHive\Chtype\VerificationResult::class);
        $this->verification_result = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>map<string, string> additional_data = 16;</code>
     * @return \Google\Protobuf\Internal\MapField
     */
    public function getAdditionalData()
    {
        return $this->additional_data;
    }

    /**
     * Generated from protobuf field <code>map<string, string> additional_data = 16;</code>
     * @param array|\Google\Protobuf\Internal\MapField $var
     * @return $this
     */
    public function setAdditionalData($var)
    {
        $arr = GPBUtil::checkMapField($var, \Google\Protobuf\Internal\GPBType::STRING, \Google\Protobuf\Internal\GPBType::STRING);
        $this->additional_data = $arr;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string authorization_code = 17;</code>
     * @return string
     */
    public function getAuthorizationCode()
    {
        return $this->authorization_code;
    }

    /**
     * Generated from protobuf field <code>string authorization_code = 17;</code>
     * @param string $var
     * @return $this
     */
    public function setAuthorizationCode($var)
    {
        GPBUtil::checkString($var, True);
        $this->authorization_code = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.Liability liability = 18;</code>
     * @return int
     */
    public function getLiability()
    {
        return $this->liability;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.Liability liability = 18;</code>
     * @param int $var
     * @return $this
     */
    public function setLiability($var)
    {
        GPBUtil::checkEnum($var, \ChargeHive\Chtype\Liability::class);
        $this->liability = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string connector_library = 19;</code>
     * @return string
     */
    public function getConnectorLibrary()
    {
        return $this->connector_library;
    }

    /**
     * Generated from protobuf field <code>string connector_library = 19;</code>
     * @param string $var
     * @return $this
     */
    public function setConnectorLibrary($var)
    {
        GPBUtil::checkString($var, True);
        $this->connector_library = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string connector_id = 20;</code>
     * @return string
     */
    public function getConnectorId()
    {
        return $this->connector_id;
    }

    /**
     * Generated from protobuf field <code>string connector_id = 20;</code>
     * @param string $var
     * @return $this
     */
    public function setConnectorId($var)
    {
        GPBUtil::checkString($var, True);
        $this->connector_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bool primary = 21;</code>
     * @return bool
     */
    public function getPrimary()
    {
        return $this->primary;
    }

    /**
     * Generated from protobuf field <code>bool primary = 21;</code>
     * @param bool $var
     * @return $this
     */
    public function setPrimary($var)
    {
        GPBUtil::checkBool($var);
        $this->primary = $var;

        return $this;
    }

    /**
     * Network ID returned by the PSP
     *
     * Generated from protobuf field <code>string network_id = 22;</code>
     * @return string
     */
    public function getNetworkId()
    {
        return $this->network_id;
    }

    /**
     * Network ID returned by the PSP
     *
     * Generated from protobuf field <code>string network_id = 22;</code>
     * @param string $var
     * @return $this
     */
    public function setNetworkId($var)
    {
        GPBUtil::checkString($var, True);
        $this->network_id = $var;

        return $this;
    }

    /**
     * Transaction ID returned by the PSP to replace transaction_id which is ambiguous
     *
     * Generated from protobuf field <code>string psp_transaction_id = 23;</code>
     * @return string
     */
    public function getPspTransactionId()
    {
        return $this->psp_transaction_id;
    }

    /**
     * Transaction ID returned by the PSP to replace transaction_id which is ambiguous
     *
     * Generated from protobuf field <code>string psp_transaction_id = 23;</code>
     * @param string $var
     * @return $this
     */
    public function setPspTransactionId($var)
    {
        GPBUtil::checkString($var, True);
        $this->psp_transaction_id = $var;

        return $this;
    }

    /**
     * Request ID returned by PSP
     *
     * Generated from protobuf field <code>string psp_request_id = 24;</code>
     * @return string
     */
    public function getPspRequestId()
    {
        return $this->psp_request_id;
    }

    /**
     * Request ID returned by PSP
     *
     * Generated from protobuf field <code>string psp_request_id = 24;</code>
     * @param string $var
     * @return $this
     */
    public function setPspRequestId($var)
    {
        GPBUtil::checkString($var, True);
        $this->psp_request_id = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string issuer_name = 25;</code>
     * @return string
     */
    public function getIssuerName()
    {
        return $this->issuer_name;
    }

    /**
     * Generated from protobuf field <code>string issuer_name = 25;</code>
     * @param string $var
     * @return $this
     */
    public function setIssuerName($var)
    {
        GPBUtil::checkString($var, True);
        $this->issuer_name = $var;

        return $this;
    }

}

