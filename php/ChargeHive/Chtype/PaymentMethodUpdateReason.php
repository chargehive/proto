<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chargehive/chtype/payment_method.proto

namespace ChargeHive\Chtype;

use UnexpectedValueException;

/**
 * Protobuf type <code>chargehive.chtype.PaymentMethodUpdateReason</code>
 */
class PaymentMethodUpdateReason
{
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_UPDATE_REASON_INVALID = 0;</code>
     */
    const PAYMENT_METHOD_UPDATE_REASON_INVALID = 0;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_UPDATE_REASON_ACCOUNT_CLOSED = 1;</code>
     */
    const PAYMENT_METHOD_UPDATE_REASON_ACCOUNT_CLOSED = 1;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_UPDATE_REASON_NEW_EXPIRY_DATE = 2;</code>
     */
    const PAYMENT_METHOD_UPDATE_REASON_NEW_EXPIRY_DATE = 2;
    /**
     * Generated from protobuf enum <code>PAYMENT_METHOD_UPDATE_REASON_NEW_ACCOUNT_NUMBER = 3;</code>
     */
    const PAYMENT_METHOD_UPDATE_REASON_NEW_ACCOUNT_NUMBER = 3;

    private static $valueToName = [
        self::PAYMENT_METHOD_UPDATE_REASON_INVALID => 'PAYMENT_METHOD_UPDATE_REASON_INVALID',
        self::PAYMENT_METHOD_UPDATE_REASON_ACCOUNT_CLOSED => 'PAYMENT_METHOD_UPDATE_REASON_ACCOUNT_CLOSED',
        self::PAYMENT_METHOD_UPDATE_REASON_NEW_EXPIRY_DATE => 'PAYMENT_METHOD_UPDATE_REASON_NEW_EXPIRY_DATE',
        self::PAYMENT_METHOD_UPDATE_REASON_NEW_ACCOUNT_NUMBER => 'PAYMENT_METHOD_UPDATE_REASON_NEW_ACCOUNT_NUMBER',
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
