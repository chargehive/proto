<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/fraud.proto

namespace ChargeHive\Chtype\FraudFactor;

use UnexpectedValueException;

/**
 * Protobuf type <code>chargehive.chtype.FraudFactor.On</code>
 */
class On
{
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_INVALID = 0;</code>
     */
    const FRAUD_FACTOR_ON_INVALID = 0;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_ADDRESS = 1;</code>
     */
    const FRAUD_FACTOR_ON_ADDRESS = 1;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_BROWSER = 2;</code>
     */
    const FRAUD_FACTOR_ON_BROWSER = 2;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_IP = 3;</code>
     */
    const FRAUD_FACTOR_ON_IP = 3;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_COUNTRY = 4;</code>
     */
    const FRAUD_FACTOR_ON_COUNTRY = 4;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_STATE = 5;</code>
     */
    const FRAUD_FACTOR_ON_STATE = 5;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_CITY = 6;</code>
     */
    const FRAUD_FACTOR_ON_CITY = 6;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_DEVICE = 7;</code>
     */
    const FRAUD_FACTOR_ON_DEVICE = 7;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_EMAIL = 8;</code>
     */
    const FRAUD_FACTOR_ON_EMAIL = 8;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_EMAIL_DOMAIN = 9;</code>
     */
    const FRAUD_FACTOR_ON_EMAIL_DOMAIN = 9;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_DATE = 10;</code>
     */
    const FRAUD_FACTOR_ON_DATE = 10;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_MERCHANT = 11;</code>
     */
    const FRAUD_FACTOR_ON_MERCHANT = 11;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_ISSUER = 12;</code>
     */
    const FRAUD_FACTOR_ON_ISSUER = 12;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_BIN = 13;</code>
     */
    const FRAUD_FACTOR_ON_BIN = 13;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_METHOD = 14;</code>
     */
    const FRAUD_FACTOR_ON_METHOD = 14;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_CUSTOMER = 15;</code>
     */
    const FRAUD_FACTOR_ON_CUSTOMER = 15;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_AMOUNT = 16;</code>
     */
    const FRAUD_FACTOR_ON_AMOUNT = 16;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_PHONE = 17;</code>
     */
    const FRAUD_FACTOR_ON_PHONE = 17;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_LANGUAGE = 18;</code>
     */
    const FRAUD_FACTOR_ON_LANGUAGE = 18;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_DISTANCE_TO_BILLING = 19;</code>
     */
    const FRAUD_FACTOR_ON_DISTANCE_TO_BILLING = 19;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_DISTANCE_TO_SHIPPING = 20;</code>
     */
    const FRAUD_FACTOR_ON_DISTANCE_TO_SHIPPING = 20;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_DISTANCE_TO_IP = 21;</code>
     */
    const FRAUD_FACTOR_ON_DISTANCE_TO_IP = 21;
    /**
     * Generated from protobuf enum <code>FRAUD_FACTOR_ON_OTHER = 22;</code>
     */
    const FRAUD_FACTOR_ON_OTHER = 22;

    private static $valueToName = [
        self::FRAUD_FACTOR_ON_INVALID => 'FRAUD_FACTOR_ON_INVALID',
        self::FRAUD_FACTOR_ON_ADDRESS => 'FRAUD_FACTOR_ON_ADDRESS',
        self::FRAUD_FACTOR_ON_BROWSER => 'FRAUD_FACTOR_ON_BROWSER',
        self::FRAUD_FACTOR_ON_IP => 'FRAUD_FACTOR_ON_IP',
        self::FRAUD_FACTOR_ON_COUNTRY => 'FRAUD_FACTOR_ON_COUNTRY',
        self::FRAUD_FACTOR_ON_STATE => 'FRAUD_FACTOR_ON_STATE',
        self::FRAUD_FACTOR_ON_CITY => 'FRAUD_FACTOR_ON_CITY',
        self::FRAUD_FACTOR_ON_DEVICE => 'FRAUD_FACTOR_ON_DEVICE',
        self::FRAUD_FACTOR_ON_EMAIL => 'FRAUD_FACTOR_ON_EMAIL',
        self::FRAUD_FACTOR_ON_EMAIL_DOMAIN => 'FRAUD_FACTOR_ON_EMAIL_DOMAIN',
        self::FRAUD_FACTOR_ON_DATE => 'FRAUD_FACTOR_ON_DATE',
        self::FRAUD_FACTOR_ON_MERCHANT => 'FRAUD_FACTOR_ON_MERCHANT',
        self::FRAUD_FACTOR_ON_ISSUER => 'FRAUD_FACTOR_ON_ISSUER',
        self::FRAUD_FACTOR_ON_BIN => 'FRAUD_FACTOR_ON_BIN',
        self::FRAUD_FACTOR_ON_METHOD => 'FRAUD_FACTOR_ON_METHOD',
        self::FRAUD_FACTOR_ON_CUSTOMER => 'FRAUD_FACTOR_ON_CUSTOMER',
        self::FRAUD_FACTOR_ON_AMOUNT => 'FRAUD_FACTOR_ON_AMOUNT',
        self::FRAUD_FACTOR_ON_PHONE => 'FRAUD_FACTOR_ON_PHONE',
        self::FRAUD_FACTOR_ON_LANGUAGE => 'FRAUD_FACTOR_ON_LANGUAGE',
        self::FRAUD_FACTOR_ON_DISTANCE_TO_BILLING => 'FRAUD_FACTOR_ON_DISTANCE_TO_BILLING',
        self::FRAUD_FACTOR_ON_DISTANCE_TO_SHIPPING => 'FRAUD_FACTOR_ON_DISTANCE_TO_SHIPPING',
        self::FRAUD_FACTOR_ON_DISTANCE_TO_IP => 'FRAUD_FACTOR_ON_DISTANCE_TO_IP',
        self::FRAUD_FACTOR_ON_OTHER => 'FRAUD_FACTOR_ON_OTHER',
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

