package apple

type LoginRequestBody struct {
	AppleID       string `json:"apple_id"`
	Password      string `json:"password"`
	ExtendedLogin bool   `json:"extended_login"`
}

type LoginResponseBody struct {
	RequestInfo struct {
		Country  string `json:"country"`
		TimeZone string `json:"timeZone"`
	} `json:"requestInfo"`

	ICloudInfo struct {
		SafariBookmarksHasMigratedToCloudKit bool `json:"SafariBookmarksHasMigratedToCloudKit"`
	} `json:"iCloudInfo"`

	Version         int  `json:"version"`
	IsExtendedLogin bool `json:"isExtendedLogin"`

	DsInfo struct {
		StatusCode int `json:"statusCode"`

		AppleID   string `json:"appleId"`
		FullName  string `json:"fullName"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`

		Locale       string `json:"locale"`
		CountryCode  string `json:"countryCode"`
		LanguageCode string `json:"languageCode"`

		Locked bool `json:"locked"`

		PrimaryEmail         string `json:"primaryEmail"`
		PrimaryEmailVerified bool   `json:"primaryEmailVerified"`

		AppleIDAliases     []any  `json:"appleIdAliases"`
		AppleIDAlias       string `json:"appleIdAlias"`
		ICloudAppleIDAlias string `json:"iCloudAppleIdAlias"`

		AppleIDEntries []struct {
			IsPrimary bool   `json:"isPrimary"`
			Type      string `json:"type"`
			Value     string `json:"value"`
		} `json:"appleIdEntries"`

		HasPaymentInfo  bool `json:"hasPaymentInfo"`
		BeneficiaryInfo struct {
			IsBeneficiary bool `json:"isBeneficiary"`
		} `json:"beneficiaryInfo"`

		HasICloudQualifyingDevice          bool     `json:"hasICloudQualifyingDevice"`
		ICDRSCapableDeviceCount            int      `json:"ICDRSCapableDeviceCount"`
		ICDRSCapableDeviceList             string   `json:"ICDRSCapableDeviceList"`
		ContinueOnDeviceEligibleDeviceInfo []string `json:"ContinueOnDeviceEligibleDeviceInfo"`

		FamilyEligible bool `json:"familyEligible"`

		ICDPEnabled        bool `json:"iCDPEnabled"`
		HsaEnabled         bool `json:"hsaEnabled"`
		UbiquityEOLEnabled bool `json:"ubiquityEOLEnabled"`
		GilliganEnabled    bool `json:"gilligan-enabled"`

		BrZoneConsolidated bool `json:"brZoneConsolidated"`
		GilliganInvited    bool `json:"gilligan-invited"`

		MailFlags struct {
			IsThreadingAvailable           bool `json:"isThreadingAvailable"`
			IsSearchV2Provisioned          bool `json:"isSearchV2Provisioned"`
			IsCKMail                       bool `json:"isCKMail"`
			IsMppSupportedInCurrentCountry bool `json:"isMppSupportedInCurrentCountry"`
		} `json:"mailFlags"`

		PcsDeleted bool `json:"pcsDeleted"`

		IroncadeMigrated bool `json:"ironcadeMigrated"`
		TantorMigrated   bool `json:"tantorMigrated"`
		NotesMigrated    bool `json:"notesMigrated"`
		BrMigrated       bool `json:"brMigrated"`

		IsHideMyEmailFeatureAvailable   bool `json:"isHideMyEmailFeatureAvailable"`
		IsHideMyEmailSubscriptionActive bool `json:"isHideMyEmailSubscriptionActive"`

		IsCustomDomainsFeatureAvailable          bool `json:"isCustomDomainsFeatureAvailable"`
		IsCustomDomainTransferSubscriptionActive bool `json:"isCustomDomainTransferSubscriptionActive"`

		AnalyticsOptInStatus bool `json:"analyticsOptInStatus"`
		IsManagedAppleID     bool `json:"isManagedAppleID"`
		IsPaidDeveloper      bool `json:"isPaidDeveloper"`
		HasUnreleasedOS      bool `json:"hasUnreleasedOS"`
		IsWebAccessAllowed   bool `json:"isWebAccessAllowed"`

		ADsID          string `json:"aDsID"`
		Dsid           string `json:"dsid"`
		HsaVersion     int    `json:"hsaVersion"`
		NotificationID string `json:"notificationId"`
	} `json:"dsInfo"`

	HasMinimumDeviceForPhotosWeb bool `json:"hasMinimumDeviceForPhotosWeb"`

	ICDPEnabled bool `json:"iCDPEnabled"`
	PcsEnabled  bool `json:"pcsEnabled"`

	ConfigBag struct {
		Urls struct {
			AccountCreateUI     string `json:"accountCreateUI"`
			AccountLoginUI      string `json:"accountLoginUI"`
			AccountLogin        string `json:"accountLogin"`
			AccountRepairUI     string `json:"accountRepairUI"`
			DownloadICloudTerms string `json:"downloadICloudTerms"`
			RepairDone          string `json:"repairDone"`
			AccountAuthorizeUI  string `json:"accountAuthorizeUI"`
			VettingURLForEmail  string `json:"vettingUrlForEmail"`
			AccountCreate       string `json:"accountCreate"`
			GetICloudTerms      string `json:"getICloudTerms"`
			VettingURLForPhone  string `json:"vettingUrlForPhone"`
		} `json:"urls"`
		AccountCreateEnabled bool `json:"accountCreateEnabled"`
	} `json:"configBag"`

	Webservices struct {
		Reminders struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"reminders"`
		Notes struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"notes"`
		Mail struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"mail"`
		Ckdatabasews struct {
			PcsRequired bool   `json:"pcsRequired"`
			URL         string `json:"url"`
			Status      string `json:"status"`
		} `json:"ckdatabasews"`
		Photosupload struct {
			PcsRequired bool   `json:"pcsRequired"`
			URL         string `json:"url"`
			Status      string `json:"status"`
		} `json:"photosupload"`
		Mcc struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"mcc"`
		Photos struct {
			PcsRequired bool   `json:"pcsRequired"`
			UploadURL   string `json:"uploadUrl"`
			URL         string `json:"url"`
			Status      string `json:"status"`
		} `json:"photos"`
		Drivews struct {
			PcsRequired bool   `json:"pcsRequired"`
			URL         string `json:"url"`
			Status      string `json:"status"`
		} `json:"drivews"`
		Uploadimagews struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"uploadimagews"`
		Schoolwork struct {
		} `json:"schoolwork"`
		Cksharews struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"cksharews"`
		Findme struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"findme"`
		Ckdeviceservice struct {
			URL string `json:"url"`
		} `json:"ckdeviceservice"`
		Iworkthumbnailws struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"iworkthumbnailws"`
		Mccgateway struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"mccgateway"`
		Calendar struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"calendar"`
		Docws struct {
			PcsRequired bool   `json:"pcsRequired"`
			URL         string `json:"url"`
			Status      string `json:"status"`
		} `json:"docws"`
		Settings struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"settings"`
		Premiummailsettings struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"premiummailsettings"`
		Ubiquity struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"ubiquity"`
		Streams struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"streams"`
		Keyvalue struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"keyvalue"`
		Mpp struct {
		} `json:"mpp"`
		Archivews struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"archivews"`
		Push struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"push"`
		Iwmb struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"iwmb"`
		Iworkexportws struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"iworkexportws"`
		Sharedlibrary struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"sharedlibrary"`
		Geows struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"geows"`
		Account struct {
			ICloudEnv struct {
				ShortID   string `json:"shortId"`
				VipSuffix string `json:"vipSuffix"`
			} `json:"iCloudEnv"`
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"account"`
		Contacts struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"contacts"`
		Developerapi struct {
			URL    string `json:"url"`
			Status string `json:"status"`
		} `json:"developerapi"`
	} `json:"webservices"`

	AppsOrder []string `json:"appsOrder"`
	Apps      struct {
		Calendar struct {
		} `json:"calendar"`
		Reminders struct {
		} `json:"reminders"`
		Keynote struct {
			IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
		} `json:"keynote"`
		Settings struct {
			CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
		} `json:"settings"`
		Mail struct {
		} `json:"mail"`
		Numbers struct {
			IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
		} `json:"numbers"`
		Photos struct {
		} `json:"photos"`
		Pages struct {
			IsQualifiedForBeta bool `json:"isQualifiedForBeta"`
		} `json:"pages"`
		Notes3 struct {
		} `json:"notes3"`
		Find struct {
			CanLaunchWithOneFactor bool `json:"canLaunchWithOneFactor"`
		} `json:"find"`
		Iclouddrive struct {
		} `json:"iclouddrive"`
		Newspublisher struct {
			IsHidden bool `json:"isHidden"`
		} `json:"newspublisher"`
		Contacts struct {
		} `json:"contacts"`
	} `json:"apps"`

	HsaTrustedBrowser    bool `json:"hsaTrustedBrowser"`
	HsaChallengeRequired bool `json:"hsaChallengeRequired"`

	PcsServiceIdentitiesIncluded bool `json:"pcsServiceIdentitiesIncluded"`
	PcsDeleted                   bool `json:"pcsDeleted"`
}

type Device struct {
	ID               string
	Name             string
	DeviceModel      string
	ModelDisplayName string
	BatteryLevel     string
	IsLocating       bool
	LostModeCapable  bool
	Location         struct {
		// lalala
	}
}
