<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chargehive/chtype/payment_method_schema.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chargehive.chtype.PaymentMethodSchemaRaw</code>
 */
class PaymentMethodSchemaRaw extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>bytes raw = 1;</code>
     */
    protected $raw = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $raw
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\PaymentMethodSchema::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>bytes raw = 1;</code>
     * @return string
     */
    public function getRaw()
    {
        return $this->raw;
    }

    /**
     * Generated from protobuf field <code>bytes raw = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setRaw($var)
    {
        GPBUtil::checkString($var, False);
        $this->raw = $var;

        return $this;
    }

}
