<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/payment_method_schema.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chargehive.chtype.PaymentMethodSchemaApplePay</code>
 */
class PaymentMethodSchemaApplePay extends \Google\Protobuf\Internal\Message
{
    /**
     * apple pay session url issued by applePayJS used to identify merchant
     *
     * Generated from protobuf field <code>string session_url = 1;</code>
     */
    protected $session_url = '';
    /**
     *domain that applePay is used on, must match registered apple pay domain
     *
     * Generated from protobuf field <code>string initiative_context = 2;</code>
     */
    protected $initiative_context = '';
    /**
     * token response from applePayJS
     *
     * Generated from protobuf field <code>string token = 3;</code>
     */
    protected $token = '';
    /**
     * session response from apple
     *
     * Generated from protobuf field <code>string session = 4;</code>
     */
    protected $session = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $session_url
     *           apple pay session url issued by applePayJS used to identify merchant
     *     @type string $initiative_context
     *          domain that applePay is used on, must match registered apple pay domain
     *     @type string $token
     *           token response from applePayJS
     *     @type string $session
     *           session response from apple
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\PaymentMethodSchema::initOnce();
        parent::__construct($data);
    }

    /**
     * apple pay session url issued by applePayJS used to identify merchant
     *
     * Generated from protobuf field <code>string session_url = 1;</code>
     * @return string
     */
    public function getSessionUrl()
    {
        return $this->session_url;
    }

    /**
     * apple pay session url issued by applePayJS used to identify merchant
     *
     * Generated from protobuf field <code>string session_url = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setSessionUrl($var)
    {
        GPBUtil::checkString($var, True);
        $this->session_url = $var;

        return $this;
    }

    /**
     *domain that applePay is used on, must match registered apple pay domain
     *
     * Generated from protobuf field <code>string initiative_context = 2;</code>
     * @return string
     */
    public function getInitiativeContext()
    {
        return $this->initiative_context;
    }

    /**
     *domain that applePay is used on, must match registered apple pay domain
     *
     * Generated from protobuf field <code>string initiative_context = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setInitiativeContext($var)
    {
        GPBUtil::checkString($var, True);
        $this->initiative_context = $var;

        return $this;
    }

    /**
     * token response from applePayJS
     *
     * Generated from protobuf field <code>string token = 3;</code>
     * @return string
     */
    public function getToken()
    {
        return $this->token;
    }

    /**
     * token response from applePayJS
     *
     * Generated from protobuf field <code>string token = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setToken($var)
    {
        GPBUtil::checkString($var, True);
        $this->token = $var;

        return $this;
    }

    /**
     * session response from apple
     *
     * Generated from protobuf field <code>string session = 4;</code>
     * @return string
     */
    public function getSession()
    {
        return $this->session;
    }

    /**
     * session response from apple
     *
     * Generated from protobuf field <code>string session = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setSession($var)
    {
        GPBUtil::checkString($var, True);
        $this->session = $var;

        return $this;
    }

}

