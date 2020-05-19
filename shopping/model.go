package shopping

import "time"

type NameValueList struct {
	Name  string   `xml:"Name"`
	Value []string `xml:"Value"`
}

type Seller struct {
	UserID                  string  `xml:"UserID"`
	FeedbackRatingStar      string  `xml:"FeedbackRatingStar"`
	FeedbackScore           int     `xml:"FeedbackScore"`
	PositiveFeedbackPercent float64 `xml:"PositiveFeedbackPercent"`
	TopRatedSeller          bool    `xml:"TopRatedSeller"`
}

type SimpleItem struct {
	AutoPay                             bool                      `xml:"AutoPay"`
	AvailableForPickupDropOff           bool                      `xml:"AvailableForPickupDropOff"`
	BestOfferEnabled                    bool                      `xml:"BestOfferEnabled"`
	BidCount                            int                       `xml:"BidCount"`
	BusinessSellerDetails               BusinessSellerDetails     `xml:"BusinessSellerDetails"`
	BuyItNowAvailable                   bool                      `xml:"BuyItNowAvailable"`
	BuyItNowPrice                       float64                   `xml:"BuyItNowPrice"`
	Charity                             Charity                   `xml:"Charity"`
	ConditionDescription                string                    `xml:"ConditionDescription"`
	ConditionDisplayName                string                    `xml:"ConditionDisplayName"`
	ConditionID                         int                       `xml:"ConditionID"`
	ConvertedBuyItNowPrice              float64                   `xml:"ConvertedBuyItNowPrice"`
	ConvertedCurrentPrice               float64                   `xml:"ConvertedCurrentPrice"`
	ConvertedCurrentPriceCurrency       string                    `xml:"ConvertedCurrentPrice>currencyID,attr"`
	Country                             string                    `xml:"Country"`
	CurrentPrice                        float64                   `xml:"CurrentPrice"`
	CurrentPriceCurrency                string                    `xml:"CurrentPrice>currencyID,attr"`
	Description                         string                    `xml:"Description"`
	DiscountPriceInfo                   DiscountPriceInfo         `xml:"DiscountPriceInfo"`
	DistanceFromBuyer                   float64                   `xml:"DistanceFromBuyer"`
	EBayNowAvailable                    bool                      `xml:"eBayNowAvailable"`
	EBayNowEligible                     bool                      `xml:"eBayNowEligible"`
	EligibleForPickupDropOff            bool                      `xml:"EligibleForPickupDropOff"`
	EndTime                             time.Time                 `xml:"EndTime"`
	ExcludeShipToLocation               string                    `xml:"ExcludeShipToLocation"`
	GalleryURL                          string                    `xml:"GalleryURL"`
	GermanMotorsSearchable              bool                      `xml:"GermanMotorsSearchable"`
	GetItFast                           bool                      `xml:"GetItFast"`
	Gift                                bool                      `xml:"Gift"`
	GlobalShipping                      bool                      `xml:"GlobalShipping"`
	HalfItemCondition                   string                    `xml:"HalfItemCondition"`
	HandlingTime                        int                       `xml:"HandlingTime"`
	HighBidder                          SimpleUser                `xml:"HighBidder"`
	HitCount                            int                       `xml:"HitCount"`
	IgnoreQuantity                      bool                      `xml:"IgnoreQuantity"`
	IntegratedMerchantCreditCardEnabled bool                      `xml:"IntegratedMerchantCreditCardEnabled"`
	ItemCompatibilityCount              int                       `xml:"ItemCompatibilityCount"`
	ItemCompatibilityList               []ItemCompatibility       `xml:"ItemCompatibilityList>Compatibility"`
	ItemID                              string                    `xml:"ItemID"`
	ItemSpecifics                       []NameValueList           `xml:"ItemSpecifics>NameValueList"`
	ListingStatus                       ListingStatusCode         `xml:"ListingStatus"`
	ListingType                         ListingTypeCode           `xml:"ListingType"`
	Location                            string                    `xml:"Location"`
	LotSize                             int                       `xml:"LotSize"`
	MinimumToBid                        float64                   `xml:"MinimumToBid"`
	MinimumToBidCurrency                string                    `xml:"MinimumToBid>currencyID,attr"`
	NewBestOffer                        bool                      `xml:"NewBestOffer"`
	PaymentAllowedSite                  SiteCode                  `xml:"PaymentAllowedSite"`
	PaymentMethods                      []BuyerPaymentMethodCode  `xml:"PaymentMethods"`
	PictureExists                       bool                      `xml:"PictureExists"`
	PictureURLs                         []string                  `xml:"PictureURL"`
	PostalCode                          string                    `xml:"PostalCode"`
	PrimaryCategoryID                   string                    `xml:"PrimaryCategoryID"`
	PrimaryCategoryIDPath               string                    `xml:"PrimaryCategoryIDPath"`
	PrimaryCategoryName                 string                    `xml:"PrimaryCategoryName"`
	ProductID                           string                    `xml:"ProductID"`
	Quantity                            int                       `xml:"Quantity"`
	QuantityAvailableHint               QuantityAvailableHintCode `xml:"QuantityAvailableHint"`
	QuantityInfo                        QuantityInfo              `xml:"QuantityInfo"`
	QuantitySold                        int                       `xml:"QuantitySold"`
	QuantitySoldByPickupInStore         int                       `xml:"QuantitySoldByPickupInStore"`
	QuantityThreshold                   int                       `xml:"QuantityThreshold"`
	RecentListing                       bool                      `xml:"RecentListing"`
	ReserveMet                          bool                      `xml:"ReserveMet"`
	ReturnPolicy                        ReturnPolicy              `xml:"ReturnPolicy"`
	SecondaryCategoryID                 string                    `xml:"SecondaryCategoryID"`
	SecondaryCategoryIDPath             string                    `xml:"SecondaryCategoryIDPath"`
	SecondaryCategoryName               string                    `xml:"SecondaryCategoryName"`
	Seller                              SimpleUser                `xml:"Seller"`
	SellerComments                      string                    `xml:"SellerComments"`
	ShippingCostSummary                 ShippingCostSummary       `xml:"ShippingCostSummary"`
	ShipToLocations                     string                    `xml:"ShipToLocations"`
	Site                                SiteCode                  `xml:"Site"`
	SKU                                 string                    `xml:"SKU"`
	StartTime                           time.Time                 `xml:"StartTime"`
	Storefront                          Storefront                `xml:"Storefront"`
	Subtitle                            string                    `xml:"Subtitle"`
	TimeLeft                            time.Duration             `xml:"TimeLeft"`
	Title                               string                    `xml:"Title"`
	TopRatedListing                     bool                      `xml:"TopRatedListing"`
	UnitInfo                            UnitInfo                  `xml:"UnitInfo"`
	Variations                          Variations                `xml:"Variations"`
	VhrAvailable                        bool                      `xml:"VhrAvailable"`
	VhrUrl                              string                    `xml:"VhrUrl"`
	ViewItemURLForNaturalSearch         string                    `xml:"ViewItemURLForNaturalSearch"`
	WatchCount                          int                       `xml:"WatchCount"`
}

type SimpleUser struct {
	UserID                  string                 `xml:"UserID"`
	AboutMeURL              string                 `xml:"AboutMeURL"`
	FeedbackDetailsURL      string                 `xml:"FeedbackDetailsURL"`
	FeedbackPrivate         bool                   `xml:"UserID"`
	FeedbackRatingStar      FeedbackRatingStarCode `xml:"FeedbackRatingStar"`
	FeedbackScore           int                    `xml:"FeedbackScore"`
	MyWorldLargeImage       string                 `xml:"MyWorldLargeImage"`
	MyWorldSmallImage       string                 `xml:"MyWorldSmallImage"`
	MyWorldURL              string                 `xml:"MyWorldURL"`
	NewUser                 bool                   `xml:"NewUser"`
	PositiveFeedbackPercent float64                `xml:"PositiveFeedbackPercent"`
	RegistrationDate        time.Time              `xml:"RegistrationDate"`
	RegistrationSite        SiteCode               `xml:"RegistrationSite"`
	ReviewsAndGuidesURL     string                 `xml:"ReviewsAndGuidesURL"`
	SellerBusinessType      SellerBusinessCode     `xml:"SellerBusinessType"`
	SellerItemsURL          string                 `xml:"SellerItemsURL"`
	SellerLevel             SellerLevelCode        `xml:"SellerLevel"`
	Status                  UserStatusCode         `xml:"Status"`
	StoreName               string                 `xml:"StoreName"`
	StoreURL                string                 `xml:"StoreURL"`
	TopRatedSeller          bool                   `xml:"TopRatedSeller"`
	UserAnonymized          bool                   `xml:"UserAnonymized"`
}

type ItemCompatibility struct {
	Compatibility []NameValueList `xml:"NameValueList"`
	Notes         string          `xml:"CompatibilityNotes"`
}

type Category struct {
	Id           string `xml:"CategoryID,omitempty"`
	IdPath       string `xml:"CategoryIDPath,omitempty"`
	Name         string `xml:"CategoryName,omitempty"`
	NamePath     string `xml:"CategoryNamePath,omitempty"`
	Level        uint   `xml:"CategoryLevel,omitempty"`
	LeafCategory bool   `xml:"LeafCategory,omitempty"`
	ParentID     string `xml:"CategoryParentID,omitempty"`
}

type BusinessSellerDetails struct {
	AdditionalContactInformation string     `xml:"AdditionalContactInformation"`
	Address                      Address    `xml:"Address"`
	Email                        string     `xml:"Email"`
	Fax                          string     `xml:"Fax"`
	LegalInvoice                 bool       `xml:"LegalInvoice"`
	TermsAndConditions           string     `xml:"TermsAndConditions"`
	TradeRegistrationNumber      string     `xml:"TradeRegistrationNumber"`
	VATDetails                   VATDetails `xml:"VATDetails"`
}

type Address struct {
	AddressID                 string `xml:"AddressID,omitempty"`
	CityName                  string `xml:"CityName,omitempty"`
	CompanyName               string `xml:"CompanyName,omitempty"`
	CountryName               string `xml:"CountryName,omitempty"`
	County                    string `xml:"County,omitempty"`
	ExternalAddressID         string `xml:"ExternalAddressID,omitempty"`
	FirstName                 string `xml:"FirstName,omitempty"`
	InternationalName         string `xml:"InternationalName,omitempty"`
	InternationalStateAndCity string `xml:"InternationalStateAndCity,omitempty"`
	InternationalStreet       string `xml:"InternationalStreet,omitempty"`
	LastName                  string `xml:"LastName,omitempty"`
	Name                      string `xml:"Name,omitempty"`
	Phone                     string `xml:"Phone,omitempty"`
	Phone2AreaOrCityCode      string `xml:"Phone2AreaOrCityCode,omitempty"`
	Phone2CountryPrefix       string `xml:"Phone2CountryPrefix,omitempty"`
	Phone2LocalNumber         string `xml:"Phone2LocalNumber,omitempty"`
	PhoneAreaOrCityCode       string `xml:"PhoneAreaOrCityCode,omitempty"`
	PhoneCountryPrefix        string `xml:"PhoneCountryPrefix,omitempty"`
	PhoneLocalNumber          string `xml:"PhoneLocalNumber,omitempty"`
	PostalCode                string `xml:"PostalCode,omitempty"`
	StateOrProvince           string `xml:"StateOrProvince,omitempty"`
	Street                    string `xml:"Street,omitempty"`
	Street1                   string `xml:"Street1,omitempty"`
	Street2                   string `xml:"Street2,omitempty"`
}

type VATDetails struct {
	BusinessSeller       bool    `xml:"BusinessSeller,omitempty"`
	RestrictedToBusiness bool    `xml:"RestrictedToBusiness,omitempty"`
	VATID                string  `xml:"VATID,omitempty"`
	VATPercent           float64 `xml:"VATPercent,omitempty"`
	VATSite              string  `xml:"VATSite,omitempty"`
}

type Charity struct {
	CharityID       string            `xml:"CharityID"`
	CharityListing  bool              `xml:"CharityListing"`
	CharityName     string            `xml:"CharityName"`
	CharityNumber   int               `xml:"CharityNumber"`
	DonationPercent float64           `xml:"DonationPercent"`
	LogoURL         string            `xml:"LogoURL"`
	Mission         string            `xml:"Mission"`
	Status          CharityStatusCode `xml:"Status"`
}

type DiscountPriceInfo struct {
	MinimumAdvertisedPrice         float64              `xml:"MinimumAdvertisedPrice,omitempty"`
	MinimumAdvertisedPriceCurrency string               `xml:"MinimumAdvertisedPriceCurrency>currencyId,attr"`
	MinimumAdvertisedPriceExposure MAPExposureCode      `xml:"MinimumAdvertisedPriceExposure"`
	OriginalRetailPrice            float64              `xml:"OriginalRetailPrice"`
	OriginalRetailPriceCurrency    float64              `xml:"OriginalRetailPrice>currencyID,attr"`
	PricingTreatment               PricingTreatmentCode `xml:"PricingTreatment"`
	SoldOffEBay                    bool                 `xml:"SoldOffeBay"`
	SoldOnEBay                     bool                 `xml:"SoldOneBay"`
}

type QuantityInfo struct {
	MinimumRemnantSet int `xml:"MinimumRemnantSet"`
}

type ReturnPolicy struct {
	Description                     string `xml:"Description"`
	EAN                             string `xml:"EAN"`
	InternationalRefund             string `xml:"InternationalRefund"`
	InternationalReturnsAccepted    string `xml:"InternationalReturnsAccepted"`
	InternationalReturnsWithin      string `xml:"InternationalReturnsWithin"`
	InternationalShippingCostPaidBy string `xml:"InternationalShippingCostPaidBy"`
	Refund                          string `xml:"Refund"`
	RestockingFeeValue              string `xml:"RestockingFeeValue"`
	RestockingFeeValueOption        string `xml:"RestockingFeeValueOption"`
	ReturnsAccepted                 string `xml:"ReturnsAccepted"`
	ReturnsWithin                   string `xml:"ReturnsWithin"`
	ShippingCostPaidBy              string `xml:"ShippingCostPaidBy"`
	WarrantyDuration                string `xml:"WarrantyDuration"`
	WarrantyOffered                 string `xml:"WarrantyOffered"`
	WarrantyType                    string `xml:"WarrantyType"`
}

type ShippingCostSummary struct {
	ListedShippingServiceCost         float64          `xml:"ListedShippingServiceCost"`
	ListedShippingServiceCostCurrency string           `xml:"ListedShippingServiceCost>currencyID,attr"`
	LocalPickup                       bool             `xml:"LocalPickup"`
	LogisticPlanType                  string           `xml:"LogisticPlanType"`
	ShippingServiceCost               float64          `xml:"ShippingServiceCost"`
	ShippingServiceCostCurrency       string           `xml:"ShippingServiceCostCurrency>currencyID,attr"`
	ShippingServiceName               string           `xml:"ShippingServiceName"`
	ShippingType                      ShippingTypeCode `xml:"ShippingType"`
}

type Storefront struct {
	StoreName string `xml:"StoreName"`
	StoreURL  string `xml:"StoreURL"`
}

type UnitInfo struct {
	UnitQuantity float64 `xml:"UnitQuantity"`
	UnitType     string  `xml:"UnitType"`
}

type Variations struct {
	Pictures []Pictures `xml:"Pictures"`
}

type Pictures struct {
	VariationSpecificName       string                        `xml:"VariationSpecificName"`
	VariationSpecificPictureSet []VariationSpecificPictureSet `xml:"VariationSpecificPictureSet"`
}

type VariationSpecificPictureSet struct {
	PictureURL             []string `xml:"PictureURL"`
	VariationSpecificValue string   `xml:"VariationSpecificValue"`
}

type Variation struct {
	DiscountPriceInfo  DiscountPriceInfo `xml:"DiscountPriceInfo"`
	ProductID          string            `xml:"ProductID"`
	Quantity           int               `xml:"Quantity"`
	QuantitySold       int               `xml:"QuantitySold"`
	SellingStatus      SellingStatus     `xml:"SellingStatus"`
	SKU                string            `xml:"SKU"`
	StartPrice         float64           `xml:"StartPrice"`
	StartPriceCurrency string            `xml:"StartPriceCurrency"`
	VariationSpecifics NameValueList     `xml:"VariationSpecifics"`
}

type SellingStatus struct {
	QuantitySold                int `xml:"QuantitySold"`
	QuantitySoldByPickupInStore int `xml:"QuantitySoldByPickupInStore"`
}

type FeedbackDetail struct {
	CommentingUser      string                 `xml:"CommentingUser"`
	CommentingUserScore int                    `xml:"CommentingUserScore"`
	CommentReplaced     bool                   `xml:"CommentReplaced"`
	CommentText         string                 `xml:"CommentText"`
	CommentTime         time.Time              `xml:"CommentTime"`
	CommentType         CommentTypeCode        `xml:"CommentType"`
	Countable           bool                   `xml:"Countable"`
	FeedbackID          string                 `xml:"FeedbackID"`
	FeedbackRatingStar  FeedbackRatingStarCode `xml:"FeedbackRatingStar"`
	FeedbackResponse    string                 `xml:"FeedbackResponse"`
	FollowUp            string                 `xml:"FollowUp"`
	FollowUpReplaced    bool                   `xml:"FollowUpReplaced"`
	ItemID              string                 `xml:"ItemID"`
	ItemPrice           float64                `xml:"ItemPrice"`
	ItemPriceCurrency   string                 `xml:"ItemPrice>currencyID,attr"`
	ItemTitle           string                 `xml:"ItemTitle"`
	ResponseReplaced    bool                   `xml:"ResponseReplaced"`
	Role                TradingRoleCode        `xml:"Role"`
	TransactionID       string                 `xml:"TransactionID"`
}

type FeedbackHistory struct {
	AverageRatingDetails                  []AverageRatingDetails `xml:"AverageRatingDetails"`
	BidRetractionFeedbackPeriods          []FeedbackPeriod       `xml:"BidRetractionFeedbackPeriods"`
	NegativeFeedbackPeriods               []FeedbackPeriod       `xml:"NegativeFeedbackPeriods"`
	NeutralCommentCountFromSuspendedUsers int                    `xml:"NeutralCommentCountFromSuspendedUsers"`
	NeutralFeedbackPeriods                []FeedbackPeriod       `xml:"NeutralFeedbackPeriods"`
	PositiveFeedbackPeriods               []FeedbackPeriod       `xml:"PositiveFeedbackPeriods"`
	TotalFeedbackPeriods                  []FeedbackPeriod       `xml:"TotalFeedbackPeriods"`
	UniqueNegativeFeedbackCount           int                    `xml:"UniqueNegativeFeedbackCount"`
	UniqueNeutralFeedbackCount            int                    `xml:"UniqueNeutralFeedbackCount"`
	UniquePositiveFeedbackCount           int                    `xml:"UniquePositiveFeedbackCount"`
}

type AverageRatingDetails struct {
	Rating       float64                  `xml:"Rating"`
	RatingCount  int                      `xml:"RatingCount"`
	RatingDetail FeedbackRatingDetailCode `xml:"RatingDetail"`
}

type FeedbackPeriod struct {
	Count        int `xml:"Count"`
	PeriodInDays int `xml:"PeriodInDays"`
}
