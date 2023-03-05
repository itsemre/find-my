package fmy

type userInfo struct {
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

type findMyInfo struct {
	StatusCode string `json:"statusCode"`
	Alert      any    `json:"alert"`

	UserInfo struct {
		FirstName        string `json:"firstName"`
		LastName         string `json:"lastName"`
		AccountFormatter int    `json:"accountFormatter"`
		MembersInfo      any    `json:"membersInfo"`
		HasMembers       bool   `json:"hasMembers"`
	} `json:"userInfo"`

	UserPreferences any `json:"userPreferences"`

	ServerContext struct {
		Info              string `json:"info"`
		ServerTimestamp   int64  `json:"serverTimestamp"`
		ValidRegion       bool   `json:"validRegion"`
		PreferredLanguage string `json:"preferredLanguage"`

		Timezone struct {
			TzName             string `json:"tzName"`
			TzCurrentName      string `json:"tzCurrentName"`
			CurrentOffset      int    `json:"currentOffset"`
			PreviousTransition int64  `json:"previousTransition"`
			PreviousOffset     int    `json:"previousOffset"`
		} `json:"timezone"`

		CallbackIntervalInMS    int `json:"callbackIntervalInMS"`
		MinCallbackIntervalInMS int `json:"minCallbackIntervalInMS"`
		MaxCallbackIntervalInMS int `json:"maxCallbackIntervalInMS"`

		Enable2FAFamilyActions bool `json:"enable2FAFamilyActions"`
		Enable2FAFamilyRemove  bool `json:"enable2FAFamilyRemove"`
		Enable2FAErase         bool `json:"enable2FAErase"`

		SessionLifespan          int `json:"sessionLifespan"`
		LastSessionExtensionTime any `json:"lastSessionExtensionTime"`

		ClassicUser   bool   `json:"classicUser"`
		CloudUser     bool   `json:"cloudUser"`
		ClientID      string `json:"clientId"`
		AuthToken     any    `json:"authToken"`
		UseAuthWidget bool   `json:"useAuthWidget"`

		IsHSA bool `json:"isHSA"`

		PrefsUpdateTime  int    `json:"prefsUpdateTime"`
		ImageBaseURL     string `json:"imageBaseUrl"`
		ItemLearnMoreURL string `json:"itemLearnMoreURL"`

		EnableMapStats               bool `json:"enableMapStats"`
		MaxLocatingTime              int  `json:"maxLocatingTime"`
		TrackInfoCacheDurationInSecs int  `json:"trackInfoCacheDurationInSecs"`
		MinTrackLocThresholdInMts    int  `json:"minTrackLocThresholdInMts"`
		InaccuracyRadiusThreshold    int  `json:"inaccuracyRadiusThreshold"`
		ItemsTabEnabled              bool `json:"itemsTabEnabled"`

		PrsID      int64 `json:"prsId"`
		ShowSllNow bool  `json:"showSllNow"`

		MacCount           int    `json:"macCount"`
		DeviceImageVersion string `json:"deviceImageVersion"`
		DeviceLoadStatus   string `json:"deviceLoadStatus"`
		MaxDeviceLoadTime  int    `json:"maxDeviceLoadTime"`
	} `json:"serverContext"`

	Content []deviceInfo `json:"content"`
}

type deviceInfo struct {
	Name                  string `json:"name"`
	ID                    string `json:"id"`
	ThisDevice            bool   `json:"thisDevice"`
	EncodedDeviceID       any    `json:"encodedDeviceId"`
	DeviceDiscoveryID     string `json:"deviceDiscoveryId"`
	DeviceStatus          string `json:"deviceStatus"`
	DeviceDisplayName     string `json:"deviceDisplayName"`
	DeviceModel           string `json:"deviceModel"`
	ModelDisplayName      string `json:"modelDisplayName"`
	RawDeviceModel        string `json:"rawDeviceModel"`
	DeviceColor           any    `json:"deviceColor"`
	DeviceClass           string `json:"deviceClass"`
	IsMac                 bool   `json:"isMac"`
	IsConsideredAccessory bool   `json:"isConsideredAccessory"`
	DeviceWithYou         bool   `json:"deviceWithYou"`
	FmlyShare             bool   `json:"fmlyShare"`

	PasscodeLength int `json:"passcodeLength"`

	Scd      bool   `json:"scd"`
	Snd      any    `json:"snd"`
	BaUUID   string `json:"baUUID"`
	PrsID    any    `json:"prsId"`
	Rm2State int    `json:"rm2State"`
	ScdPh    string `json:"scdPh"`
	Nwd      bool   `json:"nwd"`

	LowPowerMode  bool    `json:"lowPowerMode"`
	BatteryStatus string  `json:"batteryStatus"`
	BatteryLevel  float32 `json:"batteryLevel"`

	Msg        any `json:"msg"`
	Mesg       any `json:"mesg"`
	MaxMsgChar int `json:"maxMsgChar"`

	DarkWake         bool   `json:"darkWake"`
	IsLocating       bool   `json:"isLocating"`
	TrackingInfo     any    `json:"trackingInfo"`
	LocationCapable  bool   `json:"locationCapable"`
	LocationEnabled  bool   `json:"locationEnabled"`
	LostDevice       any    `json:"lostDevice"`
	LostTimestamp    string `json:"lostTimestamp"`
	LocFoundEnabled  bool   `json:"locFoundEnabled"`
	LostModeCapable  bool   `json:"lostModeCapable"`
	LostModeEnabled  bool   `json:"lostModeEnabled"`
	ActivationLocked bool   `json:"activationLocked"`
	RemoteLock       any    `json:"remoteLock"`
	LockedTimestamp  any    `json:"lockedTimestamp"`
	RemoteWipe       any    `json:"remoteWipe"`
	CanWipeAfterLock bool   `json:"canWipeAfterLock"`
	WipedTimestamp   any    `json:"wipedTimestamp"`
	WipeInProgress   bool   `json:"wipeInProgress"`

	AudioChannels []struct {
		Name      string `json:"name"`
		Available int    `json:"available"`
		Playing   bool   `json:"playing"`
		Muted     bool   `json:"muted"`
	} `json:"audioChannels"`

	Features struct {
		Btr bool `json:"BTR"`
		Llc bool `json:"LLC"`
		Clk bool `json:"CLK"`
		Teu bool `json:"TEU"`
		Snd bool `json:"SND"`
		Als bool `json:"ALS"`
		Clt bool `json:"CLT"`
		Svp bool `json:"SVP"`
		Spn bool `json:"SPN"`
		Xrm bool `json:"XRM"`
		Nwf bool `json:"NWF"`
		Cwp bool `json:"CWP"`
		Msg bool `json:"MSG"`
		Loc bool `json:"LOC"`
		Lme bool `json:"LME"`
		Lmg bool `json:"LMG"`
		Lkl bool `json:"LKL"`
		Lst bool `json:"LST"`
		Lkm bool `json:"LKM"`
		Wmg bool `json:"WMG"`
		Pss bool `json:"PSS"`
		Eal bool `json:"EAL"`
		Lae bool `json:"LAE"`
		Pin bool `json:"PIN"`
		Lck bool `json:"LCK"`
		Rem bool `json:"REM"`
		Mcs bool `json:"MCS"`
		Key bool `json:"KEY"`
		Kpd bool `json:"KPD"`
		Wip bool `json:"WIP"`
	} `json:"features,omitempty"`

	Location deviceLocation `json:"location"`
}

type deviceLocation struct {
	Timestamp          int64   `json:"timeStamp"`
	LocationFinished   bool    `json:"locationFinished"`
	LocationMode       any     `json:"locationMode"`
	LocationType       string  `json:"locationType"`
	PositionType       string  `json:"positionType"`
	SecureLocation     any     `json:"secureLocation"`
	SecureLocationTs   int     `json:"secureLocationTs"`
	IsOld              bool    `json:"isOld"`
	IsInaccurate       bool    `json:"isInaccurate"`
	VerticalAccuracy   float64 `json:"verticalAccuracy"`
	HorizontalAccuracy float64 `json:"horizontalAccuracy"`
	Altitude           float64 `json:"altitude"`
	FloorLevel         int     `json:"floorLevel"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	Address            string  `json:"formatted_address,omitempty"`
}
