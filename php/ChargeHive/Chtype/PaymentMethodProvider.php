<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/enum.proto

namespace ChargeHive\Chtype;

use UnexpectedValueException;

/**
 * Protobuf type <code>chargehive.chtype.PaymentMethodProvider</code>
 */
class PaymentMethodProvider
{
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_PROVIDER_INVALID = 0;</code>
     */
    const PAYMENT_METHOD_PROVIDER_INVALID = 0;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_PROVIDER_FORM = 1;</code>
     */
    const PAYMENT_METHOD_PROVIDER_FORM = 1;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_PROVIDER_PAYPAL = 2;</code>
     */
    const PAYMENT_METHOD_PROVIDER_PAYPAL = 2;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_PROVIDER_APPLEPAY = 3;</code>
     */
    const PAYMENT_METHOD_PROVIDER_APPLEPAY = 3;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_PROVIDER_GOOGLEPAY = 4;</code>
     */
    const PAYMENT_METHOD_PROVIDER_GOOGLEPAY = 4;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_PROVIDER_AMAZONPAY = 5;</code>
     */
    const PAYMENT_METHOD_PROVIDER_AMAZONPAY = 5;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_PROVIDER_METHOD_UPDATE = 10;</code>
     */
    const PAYMENT_METHOD_PROVIDER_METHOD_UPDATE = 10;

    private static $valueToName = [
        self::PAYMENT_METHOD_PROVIDER_INVALID => 'PAYMENT_METHOD_PROVIDER_INVALID',
        self::PAYMENT_METHOD_PROVIDER_FORM => 'PAYMENT_METHOD_PROVIDER_FORM',
        self::PAYMENT_METHOD_PROVIDER_PAYPAL => 'PAYMENT_METHOD_PROVIDER_PAYPAL',
        self::PAYMENT_METHOD_PROVIDER_APPLEPAY => 'PAYMENT_METHOD_PROVIDER_APPLEPAY',
        self::PAYMENT_METHOD_PROVIDER_GOOGLEPAY => 'PAYMENT_METHOD_PROVIDER_GOOGLEPAY',
        self::PAYMENT_METHOD_PROVIDER_AMAZONPAY => 'PAYMENT_METHOD_PROVIDER_AMAZONPAY',
        self::PAYMENT_METHOD_PROVIDER_METHOD_UPDATE => 'PAYMENT_METHOD_PROVIDER_METHOD_UPDATE',
    ];

    public static function name($value)
    {
        if (!isset(self::$valueToName[$value])) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no name defined for value %s', __CLASS__, $value));
        }
        return self::$valueToName[$value];
    }


    public static function value($name)
    {
        $const = __CLASS__ . '::' . strtoupper($name);
        if (!defined($const)) {
            throw new UnexpectedValueException(sprintf(
                    'Enum %s has no value defined for name %s', __CLASS__, $name));
        }
        return constant($const);
    }
}

