package models

// PropertyType represents the type of a property in the Attom API
type PropertyType string

const (
	PropertyTypeAgriculturalNEC       PropertyType = "AGRICULTURAL (NEC)"
	PropertyTypeApartment             PropertyType = "APARTMENT"
	PropertyTypeCabin                 PropertyType = "CABIN"
	PropertyTypeClub                  PropertyType = "CLUB"
	PropertyTypeCommercialNEC         PropertyType = "COMMERCIAL (NEC)"
	PropertyTypeCommercialBuilding    PropertyType = "COMMERCIAL BUILDING"
	PropertyTypeCommercialCondominium PropertyType = "COMMERCIAL CONDOMINIUM"
	PropertyTypeCommonArea            PropertyType = "COMMON AREA"
	PropertyTypeCondominium           PropertyType = "CONDOMINIUM"
	PropertyTypeConvertedResidence    PropertyType = "CONVERTED RESIDENCE"
	PropertyTypeCountyProperty        PropertyType = "COUNTY PROPERTY"
	PropertyTypeDuplex                PropertyType = "DUPLEX"
	PropertyTypeFarms                 PropertyType = "FARMS"
	PropertyTypeFastFoodFranchise     PropertyType = "FAST FOOD FRANCHISE"
	PropertyTypeFederalProperty       PropertyType = "FEDERAL PROPERTY"
	PropertyTypeForest                PropertyType = "FOREST"
	PropertyTypeGroupQuarters         PropertyType = "GROUP QUARTERS"
	PropertyTypeIndustrialNEC         PropertyType = "INDUSTRIAL (NEC)"
	PropertyTypeIndustrialPlant       PropertyType = "INDUSTRIAL PLANT"
	PropertyTypeManufacturedHome      PropertyType = "MANUFACTURED HOME"
	PropertyTypeMarinaFacility        PropertyType = "MARINA FACILITY"
	PropertyTypeMiscellaneous         PropertyType = "MISCELLANEOUS"
	PropertyTypeMobileHome            PropertyType = "MOBILE HOME"
	PropertyTypeMobileHomeLot         PropertyType = "MOBILE HOME LOT"
	PropertyTypeMobileHomePark        PropertyType = "MOBILE HOME PARK"
	PropertyTypeMultiFamilyDwelling   PropertyType = "MULTI FAMILY DWELLING"
	PropertyTypeNurseryHorticulture   PropertyType = "NURSERY/HORTICULTURE"
	PropertyTypeOfficeBuilding        PropertyType = "OFFICE BUILDING"
	PropertyTypePublicNEC             PropertyType = "PUBLIC (NEC)"
	PropertyTypeReligious             PropertyType = "RELIGIOUS"
	PropertyTypeResidentialNEC        PropertyType = "RESIDENTIAL (NEC)"
	PropertyTypeResidentialAcreage    PropertyType = "RESIDENTIAL ACREAGE"
	PropertyTypeResidentialLot        PropertyType = "RESIDENTIAL LOT"
	PropertyTypeRetailTrade           PropertyType = "RETAIL TRADE"
	PropertyTypeSFR                   PropertyType = "SFR"
	PropertyTypeStateProperty         PropertyType = "STATE PROPERTY"
	PropertyTypeStoresAndOffices      PropertyType = "STORES & OFFICES"
	PropertyTypeStoresAndResidential  PropertyType = "STORES & RESIDENTIAL"
	PropertyTypeTaxExempt             PropertyType = "TAX EXEMPT"
	PropertyTypeTownhouseRowhouse     PropertyType = "TOWNHOUSE/ROWHOUSE"
	PropertyTypeTriplex               PropertyType = "TRIPLEX"
	PropertyTypeUtilities             PropertyType = "UTILITIES"
	PropertyTypeVacantLandNEC         PropertyType = "VACANT LAND (NEC)"
)
