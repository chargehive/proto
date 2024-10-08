<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/enum.proto

namespace ChargeHive\Chtype;

use UnexpectedValueException;

/**
 * Protobuf type <code>chargehive.chtype.FailureType</code>
 */
class FailureType
{
    /**
     * Generated from protobuf enum <code>FAILURE_TYPE_INVALID = 0;</code>
     */
    const FAILURE_TYPE_INVALID = 0;
    /**
     *Nothing went wrong
     *
     * Generated from protobuf enum <code>FAILURE_TYPE_NONE = 1;</code>
     */
    const FAILURE_TYPE_NONE = 1;
    /**
     *Should be able to retry and may work
     *
     * Generated from protobuf enum <code>FAILURE_TYPE_SOFT = 2;</code>
     */
    const FAILURE_TYPE_SOFT = 2;
    /**
     *Retry without changing the payload will continue to fail
     *
     * Generated from protobuf enum <code>FAILURE_TYPE_HARD = 3;</code>
     */
    const FAILURE_TYPE_HARD = 3;
    /**
     *Must Retry
     *
     * Generated from protobuf enum <code>FAILURE_TYPE_RETRY = 4;</code>
     */
    const FAILURE_TYPE_RETRY = 4;
    /**
     *Server Error
     *
     * Generated from protobuf enum <code>FAILURE_TYPE_INTERNAL = 5;</code>
     */
    const FAILURE_TYPE_INTERNAL = 5;

    private static $valueToName = [
        self::FAILURE_TYPE_INVALID => 'FAILURE_TYPE_INVALID',
        self::FAILURE_TYPE_NONE => 'FAILURE_TYPE_NONE',
        self::FAILURE_TYPE_SOFT => 'FAILURE_TYPE_SOFT',
        self::FAILURE_TYPE_HARD => 'FAILURE_TYPE_HARD',
        self::FAILURE_TYPE_RETRY => 'FAILURE_TYPE_RETRY',
        self::FAILURE_TYPE_INTERNAL => 'FAILURE_TYPE_INTERNAL',
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

