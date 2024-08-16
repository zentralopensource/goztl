package goztl

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
)

var rListJSONResponse = `
[
    {
        "uuid": "af751e50-9eae-4fdb-b197-dd5041072a69",
        "name": "Default",
        "backend": "ldap",
        "ldap_config": {
            "host": "ldap.example.com",
            "bind_dn": "uid=zentral,ou=Users,o=yolo,dc=example,dc=com",
            "bind_password": "yolo",
            "users_base_dn": "ou=Users,o=yolo,dc=example,dc=com"
        },
        "openidc_config": null,
        "saml_config": null,
        "enabled_for_login": false,
        "login_session_expiry": 120,
        "username_claim": "username",
        "email_claim": "email",
        "first_name_claim": "first_name",
        "last_name_claim": "last_name",
        "full_name_claim": "full_name",
        "custom_attr_1_claim": "department",
        "custom_attr_2_claim": "branch",
        "scim_enabled": false,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var rGetByNameJSONResponse = `
[
    {
        "uuid": "af751e50-9eae-4fdb-b197-dd5041072a69",
        "name": "Default",
        "backend": "saml",
	"ldap_config": null,
        "openidc_config": null,
        "saml_config": {
	    "default_relay_state": "29eb0205-3572-4901-b773-fc82bef847ef",
            "idp_metadata": "<md></md>"
        },
        "enabled_for_login": true,
        "login_session_expiry": 120,
        "username_claim": "username",
        "email_claim": "email",
        "first_name_claim": "first_name",
        "last_name_claim": "last_name",
        "full_name_claim": "full_name",
        "custom_attr_1_claim": "",
        "custom_attr_2_claim": "",
        "scim_enabled": true,
        "created_at": "2022-07-22T01:02:03.444444",
        "updated_at": "2022-07-22T01:02:03.444444"
    }
]
`

var rGetJSONResponse = `
{
    "uuid": "af751e50-9eae-4fdb-b197-dd5041072a69",
    "name": "Default",
    "backend": "openidc",
    "ldap_config": null,
    "openidc_config": {
        "client_id": "yolo",
        "client_secret": "fomo",
        "discovery_url": "https://zentral.example.com/.well-known/openid-configuration",
        "extra_scopes": ["profile"]
    },
    "saml_config": null,
    "enabled_for_login": true,
    "login_session_expiry": 120,
    "username_claim": "username",
    "email_claim": "email",
    "first_name_claim": "first_name",
    "last_name_claim": "last_name",
    "full_name_claim": "full_name",
    "custom_attr_1_claim": "",
    "custom_attr_2_claim": "",
    "scim_enabled": true,
    "created_at": "2022-07-22T01:02:03.444444",
    "updated_at": "2022-07-22T01:02:03.444444"
}
`

func TestRealmsRealmsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/realms/realms/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, rListJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.RealmsRealms.List(ctx, nil)
	if err != nil {
		t.Errorf("RealmsRealms.List returned error: %v", err)
	}

	want := []RealmsRealm{
		{
			UUID:    "af751e50-9eae-4fdb-b197-dd5041072a69",
			Name:    "Default",
			Backend: "ldap",
			LDAPConfig: &LDAPConfig{
				Host:         "ldap.example.com",
				BindDN:       "uid=zentral,ou=Users,o=yolo,dc=example,dc=com",
				BindPassword: "yolo",
				UsersBaseDN:  "ou=Users,o=yolo,dc=example,dc=com",
			},
			EnabledForLogin:    false,
			LoginSessionExpiry: 120,
			UsernameClaim:      "username",
			EmailClaim:         "email",
			FirstNameClaim:     "first_name",
			LastNameClaim:      "last_name",
			FullNameClaim:      "full_name",
			CustomAttr1Claim:   "department",
			CustomAttr2Claim:   "branch",
			SCIMEnabled:        false,
			Created:            Timestamp{referenceTime},
			Updated:            Timestamp{referenceTime},
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("RealmsRealms.List returned %+v, want %+v", got, want)
	}
}

func TestRealmsRealmsService_GetByUUID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/realms/realms/af751e50-9eae-4fdb-b197-dd5041072a69/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, rGetJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.RealmsRealms.GetByUUID(ctx, "af751e50-9eae-4fdb-b197-dd5041072a69")
	if err != nil {
		t.Errorf("RealmsRealms.GetByUUID returned error: %v", err)
	}

	want := &RealmsRealm{
		UUID:    "af751e50-9eae-4fdb-b197-dd5041072a69",
		Name:    "Default",
		Backend: "openidc",
		OpenIDCConfig: &OpenIDCConfig{
			ClientID:     "yolo",
			ClientSecret: String("fomo"),
			DiscoveryURL: "https://zentral.example.com/.well-known/openid-configuration",
			ExtraScopes:  []string{"profile"},
		},
		EnabledForLogin:    true,
		LoginSessionExpiry: 120,
		UsernameClaim:      "username",
		EmailClaim:         "email",
		FirstNameClaim:     "first_name",
		LastNameClaim:      "last_name",
		FullNameClaim:      "full_name",
		CustomAttr1Claim:   "",
		CustomAttr2Claim:   "",
		SCIMEnabled:        true,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("RealmsRealms.GetByID returned %+v, want %+v", got, want)
	}
}

func TestRealmsRealmsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/realms/realms/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "Default")
		fmt.Fprint(w, rGetByNameJSONResponse)
	})

	ctx := context.Background()
	got, _, err := client.RealmsRealms.GetByName(ctx, "Default")
	if err != nil {
		t.Errorf("RealmsRealms.GetByName returned error: %v", err)
	}

	want := &RealmsRealm{
		UUID:    "af751e50-9eae-4fdb-b197-dd5041072a69",
		Name:    "Default",
		Backend: "saml",
		SAMLConfig: &SAMLConfig{
			DefaultRelayState: "29eb0205-3572-4901-b773-fc82bef847ef",
			IDPMetadata:       "<md></md>",
		},
		EnabledForLogin:    true,
		LoginSessionExpiry: 120,
		UsernameClaim:      "username",
		EmailClaim:         "email",
		FirstNameClaim:     "first_name",
		LastNameClaim:      "last_name",
		FullNameClaim:      "full_name",
		CustomAttr1Claim:   "",
		CustomAttr2Claim:   "",
		SCIMEnabled:        true,
		Created:            Timestamp{referenceTime},
		Updated:            Timestamp{referenceTime},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("RealmsRealms.GetByName returned %+v, want %+v", got, want)
	}
}
