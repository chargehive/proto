<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/charge.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chargehive.chtype.Dimension</code>
 */
class Dimension extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>int32 width = 1;</code>
     */
    protected $width = 0;
    /**
     * Generated from protobuf field <code>int32 height = 2;</code>
     */
    protected $height = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type int $width
     *     @type int $height
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\Charge::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>int32 width = 1;</code>
     * @return int
     */
    public function getWidth()
    {
        return $this->width;
    }

    /**
     * Generated from protobuf field <code>int32 width = 1;</code>
     * @param int $var
     * @return $this
     */
    public function setWidth($var)
    {
        GPBUtil::checkInt32($var);
        $this->width = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>int32 height = 2;</code>
     * @return int
     */
    public function getHeight()
    {
        return $this->height;
    }

    /**
     * Generated from protobuf field <code>int32 height = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setHeight($var)
    {
        GPBUtil::checkInt32($var);
        $this->height = $var;

        return $this;
    }

}

