package chtype

import "testing"

func TestCountryCodeAlpha2(t *testing.T) {
	for country, alpha2 := range map[CountryCode]string{
		COUNTRY_CODE_AUSTRIA:                  "AT",
		COUNTRY_CODE_ARGENTINA:                "AR",
		COUNTRY_CODE_AUSTRALIA:                "AU",
		COUNTRY_CODE_BAHAMAS:                  "BS",
		COUNTRY_CODE_BELGIUM:                  "BE",
		COUNTRY_CODE_BOSNIA_AND_HERZEGOVINA:   "BA",
		COUNTRY_CODE_BRAZIL:                   "BR",
		COUNTRY_CODE_BRITISH_VIRGIN_ISLANDS:   "VG",
		COUNTRY_CODE_VIRGIN_ISLANDS:           "VI",
		COUNTRY_CODE_UNITED_KINGDOM:           "GB",
		COUNTRY_CODE_UNITED_STATES_OF_AMERICA: "US",
		COUNTRY_CODE_INVALID:                  "",
		COUNTRY_CODE_GLOBAL:                   "",
	} {
		if got := country.Alpha2(); got != alpha2 {
			t.Errorf("%s.Alpha2() = %q, expected %q", country, got, alpha2)
		}
	}
}

func TestCountryCodeFromAlpha2(t *testing.T) {
	for alpha2, country := range map[string]CountryCode{
		"AT":   COUNTRY_CODE_AUSTRIA,
		"at":   COUNTRY_CODE_AUSTRIA,
		" AU ": COUNTRY_CODE_AUSTRALIA,
		"":     COUNTRY_CODE_INVALID,
		"XX":   COUNTRY_CODE_INVALID,
		"XK":   COUNTRY_CODE_INVALID,
	} {
		if got := CountryCodeFromAlpha2(alpha2); got != country {
			t.Errorf("CountryCodeFromAlpha2(%q) = %s, expected %s", alpha2, got, country)
		}
	}
}

// TestEveryCountryCodeHasAnAlpha2 checks each country the enum holds converts to an alpha-2 code and back.
func TestEveryCountryCodeHasAnAlpha2(t *testing.T) {
	for number, name := range CountryCode_name {
		country := CountryCode(number)
		if country == COUNTRY_CODE_INVALID || country == COUNTRY_CODE_GLOBAL {
			continue
		}
		alpha2 := country.Alpha2()
		if len(alpha2) != 2 {
			t.Errorf("%s has alpha-2 code %q", name, alpha2)
			continue
		}
		if back := CountryCodeFromAlpha2(alpha2); back != country {
			t.Errorf("%s is %s, which converts back to %s", name, alpha2, back)
		}
	}
}
