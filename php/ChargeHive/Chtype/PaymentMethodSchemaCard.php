<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/payment_method_schema.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chargehive.chtype.PaymentMethodSchemaCard</code>
 */
class PaymentMethodSchemaCard extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string name_on_card = 1;</code>
     */
    protected $name_on_card = '';
    /**
     * Generated from protobuf field <code>int32 valid_from_month = 2;</code>
     */
    protected $valid_from_month = 0;
    /**
     * Generated from protobuf field <code>int32 valid_from_year = 3;</code>
     */
    protected $valid_from_year = 0;
    /**
     * Generated from protobuf field <code>int32 expiry_month = 4;</code>
     */
    protected $expiry_month = 0;
    /**
     * Generated from protobuf field <code>int32 expiry_year = 5;</code>
     */
    protected $expiry_year = 0;
    /**
     * Generated from protobuf field <code>int32 issue_number = 6;</code>
     */
    protected $issue_number = 0;
    /**
     * Generated from protobuf field <code>string number = 7;</code>
     */
    protected $number = '';
    /**
     * Generated from protobuf field <code>string scheme = 8;</code>
     */
    protected $scheme = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $name_on_card
     *     @type int $valid_from_month
     *     @type int $valid_from_year
     *     @type int $expiry_month
     *     @type int $expiry_year
     *     @type int $issue_number
     *     @type string $number
     *     @type string $scheme
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\PaymentMethodSchema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string name_on_card = 1;</code>
     * @return string
     */
    public function getNameOnCard()
    {
        return $this->name_on_card;
    }

    /**
     * Generated from protobuf field <code>string name_on_card = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setNameOnCard($var)
    {
        GPBUtil::checkString($var, True);
        $this->name_on_card = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 valid_from_month = 2;</code>
     * @return int
     */
    public function getValidFromMonth()
    {
        return $this->valid_from_month;
    }

    /**
     * Generated from protobuf field <code>int32 valid_from_month = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setValidFromMonth($var)
    {
        GPBUtil::checkInt32($var);
        $this->valid_from_month = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 valid_from_year = 3;</code>
     * @return int
     */
    public function getValidFromYear()
    {
        return $this->valid_from_year;
    }

    /**
     * Generated from protobuf field <code>int32 valid_from_year = 3;</code>
     * @param int $var
     * @return $this
     */
    public function setValidFromYear($var)
    {
        GPBUtil::checkInt32($var);
        $this->valid_from_year = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 expiry_month = 4;</code>
     * @return int
     */
    public function getExpiryMonth()
    {
        return $this->expiry_month;
    }

    /**
     * Generated from protobuf field <code>int32 expiry_month = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setExpiryMonth($var)
    {
        GPBUtil::checkInt32($var);
        $this->expiry_month = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 expiry_year = 5;</code>
     * @return int
     */
    public function getExpiryYear()
    {
        return $this->expiry_year;
    }

    /**
     * Generated from protobuf field <code>int32 expiry_year = 5;</code>
     * @param int $var
     * @return $this
     */
    public function setExpiryYear($var)
    {
        GPBUtil::checkInt32($var);
        $this->expiry_year = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 issue_number = 6;</code>
     * @return int
     */
    public function getIssueNumber()
    {
        return $this->issue_number;
    }

    /**
     * Generated from protobuf field <code>int32 issue_number = 6;</code>
     * @param int $var
     * @return $this
     */
    public function setIssueNumber($var)
    {
        GPBUtil::checkInt32($var);
        $this->issue_number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string number = 7;</code>
     * @return string
     */
    public function getNumber()
    {
        return $this->number;
    }

    /**
     * Generated from protobuf field <code>string number = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string scheme = 8;</code>
     * @return string
     */
    public function getScheme()
    {
        return $this->scheme;
    }

    /**
     * Generated from protobuf field <code>string scheme = 8;</code>
     * @param string $var
     * @return $this
     */
    public function setScheme($var)
    {
        GPBUtil::checkString($var, True);
        $this->scheme = $var;

        return $this;
    }

}

