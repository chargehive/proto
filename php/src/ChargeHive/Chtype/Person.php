<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: contact.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chtype.Person</code>
 */
class Person extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string title = 1;</code>
     */
    private $title = '';
    /**
     * Generated from protobuf field <code>string first_name = 2;</code>
     */
    private $first_name = '';
    /**
     * Generated from protobuf field <code>string last_name = 3;</code>
     */
    private $last_name = '';
    /**
     * Generated from protobuf field <code>string full_name = 4;</code>
     */
    private $full_name = '';
    /**
     * Generated from protobuf field <code>string email = 5;</code>
     */
    private $email = '';
    /**
     * Generated from protobuf field <code>string phone_number = 6;</code>
     */
    private $phone_number = '';
    /**
     * Generated from protobuf field <code>string language = 7;</code>
     */
    private $language = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $title
     *     @type string $first_name
     *     @type string $last_name
     *     @type string $full_name
     *     @type string $email
     *     @type string $phone_number
     *     @type string $language
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\Contact::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string title = 1;</code>
     * @return string
     */
    public function getTitle()
    {
        return $this->title;
    }

    /**
     * Generated from protobuf field <code>string title = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setTitle($var)
    {
        GPBUtil::checkString($var, True);
        $this->title = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string first_name = 2;</code>
     * @return string
     */
    public function getFirstName()
    {
        return $this->first_name;
    }

    /**
     * Generated from protobuf field <code>string first_name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setFirstName($var)
    {
        GPBUtil::checkString($var, True);
        $this->first_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string last_name = 3;</code>
     * @return string
     */
    public function getLastName()
    {
        return $this->last_name;
    }

    /**
     * Generated from protobuf field <code>string last_name = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setLastName($var)
    {
        GPBUtil::checkString($var, True);
        $this->last_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string full_name = 4;</code>
     * @return string
     */
    public function getFullName()
    {
        return $this->full_name;
    }

    /**
     * Generated from protobuf field <code>string full_name = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setFullName($var)
    {
        GPBUtil::checkString($var, True);
        $this->full_name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string email = 5;</code>
     * @return string
     */
    public function getEmail()
    {
        return $this->email;
    }

    /**
     * Generated from protobuf field <code>string email = 5;</code>
     * @param string $var
     * @return $this
     */
    public function setEmail($var)
    {
        GPBUtil::checkString($var, True);
        $this->email = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string phone_number = 6;</code>
     * @return string
     */
    public function getPhoneNumber()
    {
        return $this->phone_number;
    }

    /**
     * Generated from protobuf field <code>string phone_number = 6;</code>
     * @param string $var
     * @return $this
     */
    public function setPhoneNumber($var)
    {
        GPBUtil::checkString($var, True);
        $this->phone_number = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string language = 7;</code>
     * @return string
     */
    public function getLanguage()
    {
        return $this->language;
    }

    /**
     * Generated from protobuf field <code>string language = 7;</code>
     * @param string $var
     * @return $this
     */
    public function setLanguage($var)
    {
        GPBUtil::checkString($var, True);
        $this->language = $var;

        return $this;
    }

}
