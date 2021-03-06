<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: chargehive/chtype/fraud.proto

namespace ChargeHive\Chtype\Metadata;

class Fraud
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Google\Protobuf\Timestamp::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
chargehive/chtype/fraud.protochargehive.chtype"�
FraudResult
fraud_check_id (	4
overall_score (2.chargehive.chtype.FraudScore1

sub_scores (2.chargehive.chtype.FraudScore<
suggested_action (2".chargehive.chtype.SuggestedAction-
	scan_time (2.google.protobuf.Timestamp
connector_library (	K
additional_data (22.chargehive.chtype.FraudResult.AdditionalDataEntryA

info_links (2-.chargehive.chtype.FraudResult.InfoLinksEntry5
AdditionalDataEntry
key (	
value (	:80
InfoLinksEntry
key (	
value (	:8"�

FraudScore
score (0

risk_level (2.chargehive.chtype.RiskLevel.
factor (2.chargehive.chtype.FraudFactor
summary (	5
data (2\'.chargehive.chtype.FraudScore.DataEntry+
	DataEntry
key (	
value (	:8"�
FraudFactor-
on (2!.chargehive.chtype.FraudFactor.On5
factor (2%.chargehive.chtype.FraudFactor.Factor
other_on (	
other_factor (	"�
On
FRAUD_FACTOR_ON_INVALID 
FRAUD_FACTOR_ON_ADDRESS
FRAUD_FACTOR_ON_BROWSER
FRAUD_FACTOR_ON_IP
FRAUD_FACTOR_ON_COUNTRY
FRAUD_FACTOR_ON_STATE
FRAUD_FACTOR_ON_CITY
FRAUD_FACTOR_ON_DEVICE
FRAUD_FACTOR_ON_EMAIL 
FRAUD_FACTOR_ON_EMAIL_DOMAIN	
FRAUD_FACTOR_ON_DATE

FRAUD_FACTOR_ON_MERCHANT
FRAUD_FACTOR_ON_ISSUER
FRAUD_FACTOR_ON_BIN
FRAUD_FACTOR_ON_METHOD
FRAUD_FACTOR_ON_CUSTOMER
FRAUD_FACTOR_ON_AMOUNT
FRAUD_FACTOR_ON_PHONE
FRAUD_FACTOR_ON_LANGUAGE\'
#FRAUD_FACTOR_ON_DISTANCE_TO_BILLING(
$FRAUD_FACTOR_ON_DISTANCE_TO_SHIPPING"
FRAUD_FACTOR_ON_DISTANCE_TO_IP
FRAUD_FACTOR_ON_OTHER"�
Factor
FRAUD_FACTOR_INVALID 
FRAUD_FACTOR_VERIFICATION
FRAUD_FACTOR_VALIDATION
FRAUD_FACTOR_VELOCITY
FRAUD_FACTOR_TRUST
FRAUD_FACTOR_DURATION
FRAUD_FACTOR_OTHER*�
SuggestedAction
SUGGESTED_ACTION_INVALID "
SUGGESTED_ACTION_NO_SUGGESTION
SUGGESTED_ACTION_REVIEW
SUGGESTED_ACTION_ALLOW
SUGGESTED_ACTION_DENY*�
	RiskLevel
RISK_LEVEL_INVALID 
RISK_LEVEL_NOT_RATED
RISK_LEVEL_NEGLIGIBLE

RISK_LEVEL_MINOR
RISK_LEVEL_MODERATE2
RISK_LEVEL_SIGNIFICANTF
RISK_LEVEL_SERVEREZBgZ4github.com/chargehive/proto/golang/chargehive/chtype�ChargeHive\\Chtype�ChargeHive\\Chtype\\Metadatabproto3'
        , true);

        static::$is_initialized = true;
    }
}

