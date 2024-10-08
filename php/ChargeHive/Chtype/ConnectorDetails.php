<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/generic.proto

namespace ChargeHive\Chtype;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>chargehive.chtype.ConnectorDetails</code>
 */
class ConnectorDetails extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string connector_library = 1;</code>
     */
    protected $connector_library = '';
    /**
     * Generated from protobuf field <code>bytes connector_credentials_json = 2;</code>
     */
    protected $connector_credentials_json = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $connector_library
     *     @type string $connector_credentials_json
     * }
     */
    public function __construct($data = NULL) {
        \ChargeHive\Chtype\Metadata\Generic::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string connector_library = 1;</code>
     * @return string
     */
    public function getConnectorLibrary()
    {
        return $this->connector_library;
    }

    /**
     * Generated from protobuf field <code>string connector_library = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setConnectorLibrary($var)
    {
        GPBUtil::checkString($var, True);
        $this->connector_library = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>bytes connector_credentials_json = 2;</code>
     * @return string
     */
    public function getConnectorCredentialsJson()
    {
        return $this->connector_credentials_json;
    }

    /**
     * Generated from protobuf field <code>bytes connector_credentials_json = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setConnectorCredentialsJson($var)
    {
        GPBUtil::checkString($var, False);
        $this->connector_credentials_json = $var;

        return $this;
    }

}

