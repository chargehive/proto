<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chargehive/chtype/enum.proto

namespace ChargeHive\Chtype;

use UnexpectedValueException;

/**
 * Protobuf type <code>chargehive.chtype.InputType</code>
 */
class InputType
{
    /**
     * Generated from protobuf enum <code>INPUT_TYPE_INVALID = 0;</code>
     */
    const INPUT_TYPE_INVALID = 0;
    /**
     * Generated from protobuf enum <code>INPUT_TYPE_PHYSICAL = 1;</code>
     */
    const INPUT_TYPE_PHYSICAL = 1;
    /**
     * Generated from protobuf enum <code>INPUT_TYPE_VIRTUAL = 2;</code>
     */
    const INPUT_TYPE_VIRTUAL = 2;
    /**
     * Generated from protobuf enum <code>INPUT_TYPE_PROXY = 3;</code>
     */
    const INPUT_TYPE_PROXY = 3;

    private static $valueToName = [
        self::INPUT_TYPE_INVALID => 'INPUT_TYPE_INVALID',
        self::INPUT_TYPE_PHYSICAL => 'INPUT_TYPE_PHYSICAL',
        self::INPUT_TYPE_VIRTUAL => 'INPUT_TYPE_VIRTUAL',
        self::INPUT_TYPE_PROXY => 'INPUT_TYPE_PROXY',
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
