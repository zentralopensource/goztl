package goztl

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/assert"
)

var depEnrollmentGetJsonResponse = `{
	"id": 30408, 
	"enrollment_secret": {
		"id": 62398, 
		"secret": "DKHG8zv5ipO7ve7dtiQpjamS8OjERUiCKL9MXKf73rKkMWrWCchDpwCWHsTvYm7x", 
		"meta_business_unit": 6771, 
		"tags": [], 
		"request_count": 0
	}, 
	"display_name": "Y3XDMSJEc7fF", 
	"use_realm_user": false, 
	"username_pattern": "", 
	"realm_user_is_admin": true, 
	"hidden_admin": false, 
	"admin_password_complexity": 3, 
	"admin_password_rotation_delay": 60, 
	"name": "BvDxOQafViRS", 
	"allow_pairing": false, 
	"auto_advance_setup": false, 
	"await_device_configured": false,
	"department": "",
	"is_mandatory": true,
	"is_mdm_removable": false,
	"is_multi_user": true,
	"is_supervised": true,
	"language": "en",
	"org_magic": "",
	"region": "",
	"skip_setup_items": ["Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"],
	"support_email_address": "",
	"support_phone_number": "",
	"include_tls_certificates": false,
	"ios_max_version": "",
	"ios_min_version": "",
	"macos_max_version": "",
	"macos_min_version": "",
	"push_certificate": 44450,
	"scep_issuer": "bdb05af8-327e-48ce-8774-340cc4cac1ed",
	"virtual_server": 38787
}`

var depEnrollmentListJsonResponse = `{
	"count": 1, 
	"results": [
		{
			"id": 30418, 
			"enrollment_secret": {
				"id": 62408, 
				"secret": "mjZSx7Uy6POMhlv2NxIHtMr1r7RflbPOQqeFteOvaoPZp7bvMdfrbUi0ku9ClPQW", 
				"meta_business_unit": 6771, 
				"tags": [], 
				"request_count": 0
			}, 
			"display_name": "E5nYP0NSQpkO", 
			"use_realm_user": false, 
			"username_pattern": "", 
			"realm_user_is_admin": true, 
			"hidden_admin": false, 
			"admin_password_complexity": 3, 
			"admin_password_rotation_delay": 60, 
			"name": "ClrXuif46jou", 
			"allow_pairing": false, 
			"auto_advance_setup": false, 
			"await_device_configured": false, 
			"department": "", 
			"is_mandatory": true, 
			"is_mdm_removable": false, 
			"is_multi_user": true, 
			"is_supervised": true, 
			"language": "", 
			"org_magic": "", 
			"region": "", 
			"skip_setup_items": ["Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"], 
			"support_email_address": "",
			"support_phone_number": "", 
			"include_tls_certificates": false, 
			"ios_max_version": "", 
			"ios_min_version": "", 
			"macos_max_version": "", 
			"macos_min_version": "", 
			"push_certificate": 44460, 
			"scep_issuer": "b4a91cfd-4afb-4a54-9041-d13c89b393de", 
			"virtual_server": 38797
		}
	]
}`

var depEnrollmentListFirstPageJsonResponse = `{
	"count": 2, 
	"next": "http://example.com/mdm/dep_enrollments/?page=2",
	"results": [
		{
			"id": 30418, 
			"enrollment_secret": {
				"id": 62408, 
				"secret": "mjZSx7Uy6POMhlv2NxIHtMr1r7RflbPOQqeFteOvaoPZp7bvMdfrbUi0ku9ClPQW", 
				"meta_business_unit": 6771, 
				"tags": [], 
				"request_count": 0
			}, 
			"display_name": "E5nYP0NSQpkO", 
			"use_realm_user": false, 
			"username_pattern": "", 
			"realm_user_is_admin": true, 
			"hidden_admin": false, 
			"admin_password_complexity": 3, 
			"admin_password_rotation_delay": 60, 
			"name": "ClrXuif46jou", 
			"allow_pairing": false, 
			"auto_advance_setup": false, 
			"await_device_configured": false, 
			"department": "", 
			"is_mandatory": true, 
			"is_mdm_removable": false, 
			"is_multi_user": true, 
			"is_supervised": true, 
			"language": "", 
			"org_magic": "", 
			"region": "", 
			"skip_setup_items": ["Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"], 
			"support_email_address": "",
			"support_phone_number": "", 
			"include_tls_certificates": false, 
			"ios_max_version": "", 
			"ios_min_version": "", 
			"macos_max_version": "", 
			"macos_min_version": "", 
			"push_certificate": 44460, 
			"scep_issuer": "b4a91cfd-4afb-4a54-9041-d13c89b393de", 
			"virtual_server": 38797
		}
	]
}`

var depEnrollmentListNextPageJsonResponse = `{
	"count": 2, 
	"results": [
		{
			"id": 30408, 
			"enrollment_secret": {
				"id": 62398, 
				"secret": "DKHG8zv5ipO7ve7dtiQpjamS8OjERUiCKL9MXKf73rKkMWrWCchDpwCWHsTvYm7x", 
				"meta_business_unit": 6771, 
				"tags": [], 
				"request_count": 0
			}, 
			"display_name": "Y3XDMSJEc7fF", 
			"use_realm_user": false, 
			"username_pattern": "", 
			"realm_user_is_admin": true, 
			"hidden_admin": false, 
			"admin_password_complexity": 3, 
			"admin_password_rotation_delay": 60, 
			"name": "BvDxOQafViRS", 
			"allow_pairing": false, 
			"auto_advance_setup": false, 
			"await_device_configured": false,
			"department": "",
			"is_mandatory": true,
			"is_mdm_removable": false,
			"is_multi_user": true,
			"is_supervised": true,
			"language": "en",
			"org_magic": "",
			"region": "",
			"skip_setup_items": ["Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"],
			"support_email_address": "",
			"support_phone_number": "",
			"include_tls_certificates": false,
			"ios_max_version": "",
			"ios_min_version": "",
			"macos_max_version": "",
			"macos_min_version": "",
			"push_certificate": 44450,
			"scep_issuer": "bdb05af8-327e-48ce-8774-340cc4cac1ed",
			"virtual_server": 38787
		}
	]
}`

var depEnrollmentCreateJsonResponse = `{
	"id": 30398, 
	"enrollment_secret": {
		"id": 62387, 
		"secret": "SldZTd7onT2EXgnxZJhe49YtdR3t39eI0L3LOGvfdxRKAMmAuja38eL3ya8t0K5t", 
		"meta_business_unit": 6771, 
		"tags": [15483], 
		"serial_numbers": ["wN7W0ESzy8GV"], 
		"udids": ["9d7be619-c727-48e2-b985-163508cf4a92"], 
		"request_count": 0
	}, 
	"display_name": "7xx3cNH7Q7lK", 
	"use_realm_user": false, 
	"username_pattern": "", 
	"realm_user_is_admin": false, 
	"hidden_admin": false, 
	"admin_password_complexity": 3, 
	"admin_password_rotation_delay": 60, 
	"name": "0gP1SHkU3oEw", 
	"allow_pairing": false, 
	"auto_advance_setup": false, 
	"await_device_configured": true, 
	"department": "", 
	"is_mandatory": true, 
	"is_mdm_removable": true, 
	"is_multi_user": true, 
	"is_supervised": false, 
	"language": "", 
	"org_magic": "", 
	"region": "", 
	"skip_setup_items": ["Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"], 
	"support_email_address": "", 
	"support_phone_number": "", 
	"include_tls_certificates": false, 
	"ios_max_version": "1.2.3", 
	"ios_min_version": "1.2.3", 
	"macos_max_version": "1.2.3", 
	"macos_min_version": "1.2.3", 
	"push_certificate": 44428, 
	"acme_issuer": "c23d9dd9-f1c1-409e-a905-5befebe7edf3", 
	"scep_issuer": "06d33014-7d66-4a81-8d63-2b96ec8948dd", 
	"virtual_server": 38765
}`

var depEnrollmentUpdateJsonResponse = `{
	"id": 30423, 
	"enrollment_secret": {"id": 62413, "secret": "xQ9xCZ7tluOTm7Ys06VudXjtvXMaaV3Utg3VCFyo6DR2NvrC8PJvDMKjZLlbF43N", "meta_business_unit": 6771, "tags": [15484], "serial_numbers": ["QliTndr8J0Cb"], "udids": ["d3502b02-fd1a-4ded-aa92-7da4304bbbb6"], "request_count": 0}, 
	"display_name": "ZYG9MeCjfIIw", 
	"use_realm_user": false, 
	"username_pattern": "", 
	"realm_user_is_admin": false, 
	"hidden_admin": false, 
	"admin_password_complexity": 3, 
	"admin_password_rotation_delay": 60, 
	"name": "G4wpXGEkPQCC", 
	"allow_pairing": false, 
	"auto_advance_setup": false, 
	"await_device_configured": true, 
	"department": "", 
	"is_mandatory": true, 
	"is_mdm_removable": true, 
	"is_multi_user": true, 
	"is_supervised": false, 
	"language": "", 
	"org_magic": "", 
	"region": "", 
	"skip_setup_items": ["Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"], 
	"support_email_address": "", 
	"support_phone_number": "", 
	"include_tls_certificates": false, 
	"ios_max_version": "1.2.3", 
	"ios_min_version": "1.2.3", 
	"macos_max_version": "1.2.3", 
	"macos_min_version": "1.2.3", 
	"push_certificate": 44466, 
	"acme_issuer": "9f68bf4f-5790-4fee-9e8c-ea63f55aaae4", 
	"scep_issuer": "fbda373b-c498-447b-a0ff-5f6e4e75e3fa", 
	"virtual_server": 38802
}`

func TestMDMDEPEnrollmentsService_List(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep_enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")

		if r.URL.Query().Get("page") == "" {
			fmt.Fprint(w, depEnrollmentListFirstPageJsonResponse)
			return
		}

		testQueryArg(t, r, "page", "2")
		fmt.Fprint(w, depEnrollmentListNextPageJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollments.List(ctx, nil)
	if err != nil {
		t.Errorf("MDMDEPEnrollments.List returned error: %v", err)
	}

	want := []MDMDEPEnrollment{
		{
			ID:                         30418,
			Name:                       "ClrXuif46jou",
			DisplayName:                "E5nYP0NSQpkO",
			Secret:                     EnrollmentSecret{ID: 62408, Secret: "mjZSx7Uy6POMhlv2NxIHtMr1r7RflbPOQqeFteOvaoPZp7bvMdfrbUi0ku9ClPQW", MetaBusinessUnitID: 6771, TagIDs: []int{}, Quota: nil, RequestCount: 0},
			UseRealmUser:               false,
			UsernamePattern:            "",
			RealmUserIsAdmin:           true,
			AdminFullName:              nil,
			AdminShortName:             nil,
			HiddenAdmin:                false,
			AdminPasswordComplexity:    3,
			AdminPasswordRotationDelay: 60,
			AllowPairing:               false,
			AutoAdvanceSetup:           false,
			AwaitDeviceConfigured:      false,
			Department:                 "",
			IsMandatory:                true,
			IsMDMRemovable:             false,
			IsMultiUser:                true,
			IsSupervised:               true,
			Language:                   "",
			OrgMagic:                   "",
			Region:                     "",
			SkipSetupItems:             []string{"Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"},
			SupportEmailAddress:        "",
			SupportPhoneNumber:         "",
			IncludeTLSCertificates:     false,
			IOSMaxVersion:              "",
			IOSMinVersion:              "",
			MacOSMaxVersion:            "",
			MacOSMinVersion:            "",
			PushCertificateID:          44460,
			ACMEIssuerUUID:             nil,
			SCEPIssuerUUID:             "b4a91cfd-4afb-4a54-9041-d13c89b393de",
			BlueprintID:                nil,
			RealmUUID:                  nil,
			VirtualServerID:            38797,
		},
		{
			ID:                         30408,
			Name:                       "BvDxOQafViRS",
			DisplayName:                "Y3XDMSJEc7fF",
			Secret:                     EnrollmentSecret{ID: 62398, Secret: "DKHG8zv5ipO7ve7dtiQpjamS8OjERUiCKL9MXKf73rKkMWrWCchDpwCWHsTvYm7x", MetaBusinessUnitID: 6771, TagIDs: []int{}, Quota: nil, RequestCount: 0},
			UseRealmUser:               false,
			UsernamePattern:            "",
			RealmUserIsAdmin:           true,
			AdminFullName:              nil,
			AdminShortName:             nil,
			HiddenAdmin:                false,
			AdminPasswordComplexity:    3,
			AdminPasswordRotationDelay: 60,
			AllowPairing:               false,
			AutoAdvanceSetup:           false,
			AwaitDeviceConfigured:      false,
			Department:                 "",
			IsMandatory:                true,
			IsMDMRemovable:             false,
			IsMultiUser:                true,
			IsSupervised:               true,
			Language:                   "en",
			OrgMagic:                   "",
			Region:                     "",
			SkipSetupItems:             []string{"Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"},
			SupportEmailAddress:        "",
			SupportPhoneNumber:         "",
			IncludeTLSCertificates:     false,
			IOSMaxVersion:              "",
			IOSMinVersion:              "",
			MacOSMaxVersion:            "",
			MacOSMinVersion:            "",
			PushCertificateID:          44450,
			SCEPIssuerUUID:             "bdb05af8-327e-48ce-8774-340cc4cac1ed",
			VirtualServerID:            38787,
		},
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollment.List returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentsService_GetByID(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep_enrollments/30408/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		fmt.Fprint(w, depEnrollmentGetJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollments.GetByID(ctx, 30408)
	if err != nil {
		t.Errorf("MDMDEPEnrollments.GetByID returned error: %v", err)
	}

	want := &MDMDEPEnrollment{
		ID:                         30408,
		Name:                       "BvDxOQafViRS",
		DisplayName:                "Y3XDMSJEc7fF",
		Secret:                     EnrollmentSecret{ID: 62398, Secret: "DKHG8zv5ipO7ve7dtiQpjamS8OjERUiCKL9MXKf73rKkMWrWCchDpwCWHsTvYm7x", MetaBusinessUnitID: 6771, TagIDs: []int{}, Quota: nil, RequestCount: 0},
		UseRealmUser:               false,
		UsernamePattern:            "",
		RealmUserIsAdmin:           true,
		AdminFullName:              nil,
		AdminShortName:             nil,
		HiddenAdmin:                false,
		AdminPasswordComplexity:    3,
		AdminPasswordRotationDelay: 60,
		AllowPairing:               false,
		AutoAdvanceSetup:           false,
		AwaitDeviceConfigured:      false,
		Department:                 "",
		IsMandatory:                true,
		IsMDMRemovable:             false,
		IsMultiUser:                true,
		IsSupervised:               true,
		Language:                   "en",
		OrgMagic:                   "",
		Region:                     "",
		SkipSetupItems:             []string{"Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"},
		SupportEmailAddress:        "",
		SupportPhoneNumber:         "",
		IncludeTLSCertificates:     false,
		IOSMaxVersion:              "",
		IOSMinVersion:              "",
		MacOSMaxVersion:            "",
		MacOSMinVersion:            "",
		PushCertificateID:          44450,
		SCEPIssuerUUID:             "bdb05af8-327e-48ce-8774-340cc4cac1ed",
		VirtualServerID:            38787,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollments.GetByID returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentsService_GetByName(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep_enrollments/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, "Accept", "application/json")
		testQueryArg(t, r, "name", "ClrXuif46jou")
		fmt.Fprint(w, depEnrollmentListJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollments.GetByName(ctx, "ClrXuif46jou")
	if err != nil {
		t.Errorf("MDMDEPEnrollments.GetByName returned error: %v", err)
	}

	want := &MDMDEPEnrollment{
		ID:                         30418,
		Name:                       "ClrXuif46jou",
		DisplayName:                "E5nYP0NSQpkO",
		Secret:                     EnrollmentSecret{ID: 62408, Secret: "mjZSx7Uy6POMhlv2NxIHtMr1r7RflbPOQqeFteOvaoPZp7bvMdfrbUi0ku9ClPQW", MetaBusinessUnitID: 6771, TagIDs: []int{}, Quota: nil, RequestCount: 0},
		UseRealmUser:               false,
		UsernamePattern:            "",
		RealmUserIsAdmin:           true,
		AdminFullName:              nil,
		AdminShortName:             nil,
		HiddenAdmin:                false,
		AdminPasswordComplexity:    3,
		AdminPasswordRotationDelay: 60,
		AllowPairing:               false,
		AutoAdvanceSetup:           false,
		AwaitDeviceConfigured:      false,
		Department:                 "",
		IsMandatory:                true,
		IsMDMRemovable:             false,
		IsMultiUser:                true,
		IsSupervised:               true,
		Language:                   "",
		OrgMagic:                   "",
		Region:                     "",
		SkipSetupItems:             []string{"Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"},
		SupportEmailAddress:        "",
		SupportPhoneNumber:         "",
		IncludeTLSCertificates:     false,
		IOSMaxVersion:              "",
		IOSMinVersion:              "",
		MacOSMaxVersion:            "",
		MacOSMinVersion:            "",
		PushCertificateID:          44460,
		ACMEIssuerUUID:             nil,
		SCEPIssuerUUID:             "b4a91cfd-4afb-4a54-9041-d13c89b393de",
		BlueprintID:                nil,
		RealmUUID:                  nil,
		VirtualServerID:            38797,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollments.GetByName returned %+v, want %+v", got, want)
	}

}

func TestMDMDEPEnrollmentsService_Create(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	createRequest := &MDMDEPEnrollmentCreationRequest{
		Name:                       "0gP1SHkU3oEw",
		DisplayName:                "7xx3cNH7Q7lK",
		Secret:                     EnrollmentSecretRequest{MetaBusinessUnitID: 6771, TagIDs: []int{15483}, SerialNumbers: []string{"wN7W0ESzy8GV"}, UDIDs: []string{"9d7be619-c727-48e2-b985-163508cf4a92"}, Quota: nil},
		UseRealmUser:               false,
		UsernamePattern:            "",
		RealmUserIsAdmin:           false,
		HiddenAdmin:                false,
		AdminPasswordComplexity:    3,
		AdminPasswordRotationDelay: 60,
		AllowPairing:               false,
		AutoAdvanceSetup:           false,
		AwaitDeviceConfigured:      true,
		IsMandatory:                true,
		IsMDMRemovable:             true,
		IsMultiUser:                true,
		IsSupervised:               false,
		SkipSetupItems:             []string{"Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"},
		IncludeTLSCertificates:     false,
		IOSMaxVersion:              "1.2.3",
		IOSMinVersion:              "1.2.3",
		MacOSMaxVersion:            "1.2.3",
		MacOSMinVersion:            "1.2.3",
		PushCertificateID:          44428,
		ACMEIssuerUUID:             String("c23d9dd9-f1c1-409e-a905-5befebe7edf3"),
		SCEPIssuerUUID:             "06d33014-7d66-4a81-8d63-2b96ec8948dd",
		VirtualServerID:            38765,
	}

	mux.HandleFunc("/mdm/dep_enrollments/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDEPEnrollmentCreationRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "POST")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, createRequest, v)

		fmt.Fprint(w, depEnrollmentCreateJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollments.Create(ctx, createRequest)
	if err != nil {
		t.Errorf("MDMDEPEnrollments.Create returned error: %v", err)
	}

	want := &MDMDEPEnrollment{
		ID:                         30398,
		Name:                       "0gP1SHkU3oEw",
		DisplayName:                "7xx3cNH7Q7lK",
		Secret:                     EnrollmentSecret{ID: 62387, Secret: "SldZTd7onT2EXgnxZJhe49YtdR3t39eI0L3LOGvfdxRKAMmAuja38eL3ya8t0K5t", MetaBusinessUnitID: 6771, TagIDs: []int{15483}, SerialNumbers: []string{"wN7W0ESzy8GV"}, UDIDs: []string{"9d7be619-c727-48e2-b985-163508cf4a92"}, Quota: nil, RequestCount: 0},
		UseRealmUser:               false,
		UsernamePattern:            "",
		RealmUserIsAdmin:           false,
		HiddenAdmin:                false,
		AdminPasswordComplexity:    3,
		AdminPasswordRotationDelay: 60,
		AllowPairing:               false,
		AutoAdvanceSetup:           false,
		AwaitDeviceConfigured:      true,
		Department:                 "",
		IsMandatory:                true,
		IsMDMRemovable:             true,
		IsMultiUser:                true,
		IsSupervised:               false,
		Language:                   "",
		OrgMagic:                   "",
		Region:                     "",
		SkipSetupItems:             []string{"Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"},
		SupportEmailAddress:        "",
		SupportPhoneNumber:         "",
		IncludeTLSCertificates:     false,
		IOSMaxVersion:              "1.2.3",
		IOSMinVersion:              "1.2.3",
		MacOSMaxVersion:            "1.2.3",
		MacOSMinVersion:            "1.2.3",
		PushCertificateID:          44428,
		ACMEIssuerUUID:             String("c23d9dd9-f1c1-409e-a905-5befebe7edf3"),
		SCEPIssuerUUID:             "06d33014-7d66-4a81-8d63-2b96ec8948dd",
		VirtualServerID:            38765,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMDEPEnrollments.Create returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentsService_Update(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	updateRequest := &MDMDEPEnrollmentUpdateRequest{
		DisplayName: "ZYG9MeCjfIIw",
	}

	mux.HandleFunc("/mdm/dep_enrollments/30423/", func(w http.ResponseWriter, r *http.Request) {
		v := new(MDMDEPEnrollmentUpdateRequest)
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil {
			t.Fatal(err)
		}
		testMethod(t, r, "PUT")
		testHeader(t, r, "Accept", "application/json")
		testHeader(t, r, "Content-Type", "application/json")
		assert.Equal(t, updateRequest, v)
		fmt.Fprint(w, depEnrollmentUpdateJsonResponse)
	})

	ctx := context.Background()
	got, _, err := client.MDMDEPEnrollments.Update(ctx, 30423, updateRequest)
	if err != nil {
		t.Errorf("MDMDEPEnrollments.Update returned error: %v", err)
	}

	want := &MDMDEPEnrollment{
		ID:                         30423,
		Name:                       "G4wpXGEkPQCC",
		DisplayName:                "ZYG9MeCjfIIw",
		Secret:                     EnrollmentSecret{ID: 62413, Secret: "xQ9xCZ7tluOTm7Ys06VudXjtvXMaaV3Utg3VCFyo6DR2NvrC8PJvDMKjZLlbF43N", MetaBusinessUnitID: 6771, TagIDs: []int{15484}, SerialNumbers: []string{"QliTndr8J0Cb"}, UDIDs: []string{"d3502b02-fd1a-4ded-aa92-7da4304bbbb6"}, Quota: nil, RequestCount: 0},
		UseRealmUser:               false,
		UsernamePattern:            "",
		RealmUserIsAdmin:           false,
		HiddenAdmin:                false,
		AdminPasswordComplexity:    3,
		AdminPasswordRotationDelay: 60,
		AllowPairing:               false,
		AutoAdvanceSetup:           false,
		AwaitDeviceConfigured:      true,
		Department:                 "",
		IsMandatory:                true,
		IsMDMRemovable:             true,
		IsMultiUser:                true,
		IsSupervised:               false,
		Language:                   "",
		OrgMagic:                   "",
		Region:                     "",
		SkipSetupItems:             []string{"Accessibility", "ActionButton", "AgeBasedSafetySettings", "Android", "Appearance", "AppleID", "AppStore", "Biometric", "CameraButton", "DeviceToDeviceMigration", "Diagnostics", "DisplayTone", "EnableLockdownMode", "FileVault", "HomeButtonSensitivity", "iCloudDiagnostics", "iCloudStorage", "iMessageAndFaceTime", "Intelligence", "Keyboard", "Location", "MessagingActivationUsingPhoneNumber", "Multitasking", "OnBoarding", "OSShowcase", "Passcode", "Payment", "Privacy", "Restore", "RestoreCompleted", "ScreenSaver", "Safety", "SafetyAndHandling", "ScreenTime", "SIMSetup", "Siri", "SoftwareUpdate", "SpokenLanguage", "TapToSetup", "TermsOfAddress", "Tips", "TOS", "TVHomeScreenSync", "TVProviderSignIn", "TVRoom", "UnlockWithWatch", "UpdateCompleted", "Wallpaper", "WatchMigration", "WebContentFiltering", "Welcome", "Zoom"},
		SupportEmailAddress:        "",
		SupportPhoneNumber:         "",
		IncludeTLSCertificates:     false,
		IOSMaxVersion:              "1.2.3",
		IOSMinVersion:              "1.2.3",
		MacOSMaxVersion:            "1.2.3",
		MacOSMinVersion:            "1.2.3",
		PushCertificateID:          44466,
		ACMEIssuerUUID:             String("9f68bf4f-5790-4fee-9e8c-ea63f55aaae4"),
		SCEPIssuerUUID:             "fbda373b-c498-447b-a0ff-5f6e4e75e3fa",
		VirtualServerID:            38802,
	}
	if !cmp.Equal(got, want) {
		t.Errorf("MDMOTAEnrollments.Update returned %+v, want %+v", got, want)
	}
}

func TestMDMDEPEnrollmentsService_Delete(t *testing.T) {
	client, mux, teardown := setup()
	defer teardown()

	mux.HandleFunc("/mdm/dep_enrollments/1/", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "DELETE")
		w.WriteHeader(http.StatusNoContent)
	})

	ctx := context.Background()
	_, err := client.MDMDEPEnrollments.Delete(ctx, 1)
	if err != nil {
		t.Errorf("MDMOTAEnrollments.Delete returned error: %v", err)
	}
}
