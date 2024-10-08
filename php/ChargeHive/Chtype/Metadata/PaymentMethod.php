<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# NO CHECKED-IN PROTOBUF GENCODE
# source: chargehive/chtype/payment_method.proto

namespace ChargeHive\Chtype\Metadata;

class PaymentMethod
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \ChargeHive\Chtype\Metadata\Enum::initOnce();
        \ChargeHive\Chtype\Metadata\PaymentMethodSchema::initOnce();
        \GPBMetadata\Google\Protobuf\Timestamp::initOnce();
        $pool->internalAddGeneratedFile(
            "\x0A\x82\x0F\x0A&chargehive/chtype/payment_method.proto\x12\x11chargehive.chtype\x1A-chargehive/chtype/payment_method_schema.proto\x1A\x1Fgoogle/protobuf/timestamp.proto\"\x83\x02\x0A\x0DPaymentMethod\x126\x0A\x06schema\x18\x01 \x01(\x0E2&.chargehive.chtype.PaymentMethodSchema\x12\x0C\x0A\x04json\x18\x02 \x01(\x0C\x122\x0A\x04type\x18\x03 \x01(\x0E2\$.chargehive.chtype.PaymentMethodType\x12:\x0A\x08provider\x18\x04 \x01(\x0E2(.chargehive.chtype.PaymentMethodProvider\x12<\x0A\x09inputType\x18\x05 \x01(\x0E2).chargehive.chtype.PaymentMethodInputType\"\xAC\x01\x0A\x1DPaymentMethodVerificationItem\x12B\x0A\x04type\x18\x01 \x01(\x0E24.chargehive.chtype.PaymentMethodVerificationItemType\x12\x0D\x0A\x05value\x18\x02 \x01(\x0C\x12\x18\x0A\x10transport_key_id\x18\x04 \x01(\x09\x12\x10\x0A\x08is_error\x18\x05 \x01(\x08\x12\x0C\x0A\x04name\x18\x06 \x01(\x09\"\xEB\x03\x0A\x11PaymentMethodInfo\x12\x10\x0A\x08token_id\x18\x01 \x01(\x09\x12\x0C\x0A\x04name\x18\x02 \x01(\x09\x12\x16\x0A\x0Epayment_scheme\x18\x03 \x01(\x09\x12.\x0A\x0Avalid_from\x18\x04 \x01(\x0B2\x1A.google.protobuf.Timestamp\x12*\x0A\x06expiry\x18\x05 \x01(\x0B2\x1A.google.protobuf.Timestamp\x122\x0A\x04type\x18\x06 \x01(\x0E2\$.chargehive.chtype.PaymentMethodType\x12<\x0A\x04info\x18\x07 \x03(\x0B2..chargehive.chtype.PaymentMethodInfo.InfoEntry\x12:\x0A\x08provider\x18\x08 \x01(\x0E2(.chargehive.chtype.PaymentMethodProvider\x12<\x0A\x09inputType\x18\x09 \x01(\x0E2).chargehive.chtype.PaymentMethodInputType\x12)\x0A\x06status\x18\x0A \x01(\x0E2\x19.chargehive.chtype.Status\x1A+\x0A\x09InfoEntry\x12\x0B\x0A\x03key\x18\x01 \x01(\x09\x12\x0D\x0A\x05value\x18\x02 \x01(\x09:\x028\x01*\xF2\x02\x0A\x19PaymentMethodUpdateReason\x12(\x0A\$PAYMENT_METHOD_UPDATE_REASON_INVALID\x10\x00\x12/\x0A+PAYMENT_METHOD_UPDATE_REASON_ACCOUNT_CLOSED\x10\x01\x120\x0A,PAYMENT_METHOD_UPDATE_REASON_NEW_EXPIRY_DATE\x10\x02\x123\x0A/PAYMENT_METHOD_UPDATE_REASON_NEW_ACCOUNT_NUMBER\x10\x03\x127\x0A3PAYMENT_METHOD_UPDATE_REASON_NEW_ACCOUNT_AND_EXPIRY\x10\x04\x12*\x0A&PAYMENT_METHOD_UPDATE_REASON_NO_CHANGE\x10\x05\x12.\x0A*PAYMENT_METHOD_UPDATE_REASON_NOT_SUPPORTED\x10\x06*\xEB\x02\x0A!PaymentMethodVerificationItemType\x121\x0A-PAYMENT_METHOD_VERIFICATION_ITEM_TYPE_INVALID\x10\x00\x12-\x0A)PAYMENT_METHOD_VERIFICATION_ITEM_TYPE_CVV\x10\x01\x128\x0A4PAYMENT_METHOD_VERIFICATION_ITEM_TYPE_DEVICE_DETAILS\x10\x02\x12;\x0A7PAYMENT_METHOD_VERIFICATION_ITEM_TYPE_IDENTIFY_RESPONSE\x10\x03\x12<\x0A8PAYMENT_METHOD_VERIFICATION_ITEM_TYPE_CHALLENGE_RESPONSE\x10\x04\x12/\x0A+PAYMENT_METHOD_VERIFICATION_ITEM_TYPE_NAMED\x10\x05BgZ4github.com/chargehive/proto/golang/chargehive/chtype\xCA\x02\x11ChargeHive\\Chtype\xE2\x02\x1AChargeHive\\Chtype\\Metadatab\x06proto3"
        , true);

        static::$is_initialized = true;
    }
}

