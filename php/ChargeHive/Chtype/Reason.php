<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/generic.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chargehive.chtype.Reason</code>
 */
class Reason extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string description = 1;</code>
     */
    protected $description = '';
    /**
     * Generated from protobuf field <code>.chargehive.chtype.Reason.ReasonType reason_type = 2;</code>
     */
    protected $reason_type = 0;
    /**
     * Generated from protobuf field <code>string requestor_comment = 3;</code>
     */
    protected $requestor_comment = '';
    /**
     * Generated from protobuf field <code>.chargehive.chtype.ActorType requested_by = 4;</code>
     */
    protected $requested_by = 0;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $description
     *     @type int $reason_type
     *     @type string $requestor_comment
     *     @type int $requested_by
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\Generic::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string description = 1;</code>
     * @return string
     */
    public function getDescription()
    {
        return $this->description;
    }

    /**
     * Generated from protobuf field <code>string description = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setDescription($var)
    {
        GPBUtil::checkString($var, True);
        $this->description = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.Reason.ReasonType reason_type = 2;</code>
     * @return int
     */
    public function getReasonType()
    {
        return $this->reason_type;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.Reason.ReasonType reason_type = 2;</code>
     * @param int $var
     * @return $this
     */
    public function setReasonType($var)
    {
        GPBUtil::checkEnum($var, \ChargeHive\Chtype\Reason\ReasonType::class);
        $this->reason_type = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string requestor_comment = 3;</code>
     * @return string
     */
    public function getRequestorComment()
    {
        return $this->requestor_comment;
    }

    /**
     * Generated from protobuf field <code>string requestor_comment = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setRequestorComment($var)
    {
        GPBUtil::checkString($var, True);
        $this->requestor_comment = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.ActorType requested_by = 4;</code>
     * @return int
     */
    public function getRequestedBy()
    {
        return $this->requested_by;
    }

    /**
     * Generated from protobuf field <code>.chargehive.chtype.ActorType requested_by = 4;</code>
     * @param int $var
     * @return $this
     */
    public function setRequestedBy($var)
    {
        GPBUtil::checkEnum($var, \ChargeHive\Chtype\ActorType::class);
        $this->requested_by = $var;

        return $this;
    }

}

