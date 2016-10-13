package dbhwin32

import (
	"syscall"

	"github.com/dbheise/dbhwin32/wrapper"
	"github.com/winlabs/gowin32"
)

type NumberString struct {
	Number uint32
	String string
}
type OfficeDetails struct {
	ReleaseVersion string
	ReleaseType    string
	MajorVersion   string
	MinorVersion   string
	ProductID      string
	LanguageID     string
	is64           bool
	isDebug        bool
}
type InstalledProduct struct {
	ProductGUID          string
	InstallState         NumberString
	PackageName          string
	Transforms           string
	Language             string
	ProductName          string
	AssignmentType       string
	InstanceType         string
	AuthorizedLUAApp     string
	PackageCode          string
	Version              string
	ProductIcon          string
	InstalledProductName string
	VersionString        string
	HelpLink             string
	HelpTelephone        string
	InstallLocation      string
	InstallSource        string
	InstallDate          string
	Publisher            string
	LocalPackage         string
	URLInfoAbout         string
	URLUpdateInfo        string
	VersionMinor         string
	VersionMajor         string
	ProductID            string
	RegCompany           string
	RegOwner             string
	InstalledLanguage    string
}
type ProductPatch struct {
	PatchGUID string
	Transform string
}

func GetAllComponentConsumers(componentID string) []string {
	products := []string{}
	buf := make([]uint16, 39)
	compID := syscall.StringToUTF16Ptr(componentID)
	var i uint32
	for i = 0; ; i++ {
		r := wrapper.MsiEnumClients(compID, i, &buf[0])
		if r == nil {
			products = append(products, syscall.UTF16ToString(buf))
		} else if r == syscall.Errno(259) {
			break
		} else {
			panic(r)
		}
	}
	return products
}

func GetAllProducts() []string {
	products := []string{}
	buf := make([]uint16, 39)
	var i uint32
	for i = 0; ; i++ {
		r := wrapper.MsiEnumProducts(i, &buf[0])
		if r == nil {
			products = append(products, syscall.UTF16ToString(buf))
		} else if r == syscall.Errno(259) {
			break
		} else {
			panic(r)
		}
	}
	return products
}

func GetAllPatches(productID string) []ProductPatch {
	patches := []ProductPatch{}
	buf := make([]uint16, 39)
	prodID := syscall.StringToUTF16Ptr(productID)

	var i, transformBufSize uint32
	for i = 0; ; i++ {
		transformBufSize = 2048
		transformBuf := make([]uint16, transformBufSize)
		err := wrapper.MsiEnumPatches(prodID, i, &buf[0], &transformBuf[0], &transformBufSize)
		if err == nil {
			var patch ProductPatch
			patch.PatchGUID = syscall.UTF16ToString(buf)
			patch.Transform = syscall.UTF16ToString(transformBuf)
			patches = append(patches, patch)
		} else if err == syscall.Errno(259) {
			break
		} else {
			panic(err)
		}
	}
	return patches
}

func convertInstallStateToString(state gowin32.InstallState) (ans string) {
	switch state {
	case gowin32.InstallStateBadConfig:
		ans = "Bad Config"
	case gowin32.InstallStateIncomplete:
		ans = "Incomplete"
	case gowin32.InstallStateSourceAbsent:
		ans = "Source Absent"
	case gowin32.InstallStateMoreData:
		ans = "More Data"
	case gowin32.InstallStateInvalidArg:
		ans = "Invalid Arg"
	case gowin32.InstallStateUnknown:
		ans = "Unknown"
	case gowin32.InstallStateBroken:
		ans = "Broken"
	case gowin32.InstallStateAdvertised:
		ans = "Advertised"
	case gowin32.InstallStateAbsent:
		ans = "Absent"
	case gowin32.InstallStateLocal:
		ans = "Local"
	case gowin32.InstallStateSource:
		ans = "Source"
	case gowin32.InstallStateDefault:
		ans = "Default"
	}
	return ans
}

func convertReleaseVersionToString(release string) (ans string) {
	switch release {
	case "0":
		ans = "Pre Beta 1"
	}
	return ans
}

func parseOfficeProductGUID(productID string) OfficeDetails {
	var details OfficeDetails
	details.ReleaseVersion = convertReleaseVersionToString(productID[1:2])
	details.ReleaseType = productID[2:3]
	details.MajorVersion = productID[3:5]
	details.MinorVersion = productID[5:9]
	details.ProductID = productID[10:14]
	details.LanguageID = productID[15:19]
	details.is64 = bool(productID[20] == 1)
	details.isDebug = bool(productID[25] == 1)
	return details
}

func GetProductDetails(productGUID string) InstalledProduct {
	var d InstalledProduct
	d.ProductGUID = productGUID
	d.InstallState.Number = uint32(gowin32.GetInstalledProductState(d.ProductGUID))
	d.InstallState.String = convertInstallStateToString(gowin32.InstallState(d.InstallState.Number))
	d.PackageName, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyPackageName)
	d.Transforms, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyTransforms)
	d.Language, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyLanguage)
	d.ProductName, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyProductName)
	d.AssignmentType, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyAssignmentType)
	d.InstanceType, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyInstanceType)
	d.AuthorizedLUAApp, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyAuthorizedLUAApp)
	d.PackageCode, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyPackageCode)
	d.Version, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyVersion)
	d.ProductIcon, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyProductIcon)
	d.InstalledProductName, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyInstalledProductName)
	d.VersionString, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyVersionString)
	d.HelpLink, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyHelpLink)
	d.HelpTelephone, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyHelpTelephone)
	d.InstallLocation, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyInstallLocation)
	d.InstallSource, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyInstallSource)
	d.InstallDate, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyInstallDate)
	d.Publisher, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyPublisher)
	d.LocalPackage, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyLocalPackage)
	d.URLInfoAbout, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyURLInfoAbout)
	d.URLUpdateInfo, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyURLUpdateInfo)
	d.VersionMinor, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyVersionMinor)
	d.VersionMajor, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyVersionMajor)
	d.ProductID, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyProductID)
	d.RegCompany, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyRegCompany)
	d.RegOwner, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyRegOwner)
	d.InstalledLanguage, _ = gowin32.GetInstalledProductProperty(d.ProductGUID, gowin32.InstallPropertyInstalledLanguage)
	return d
}

func GetOfficeDetails() {
	/*	if strings.HasSuffix(guids[i], "000000FF1CE}") { //Office 2007+
			details := parseOfficeProductGUID(guids[i])
		} else if strings.HasSuffix(guids[i], "78E1-11D2-B60F-006097C998E7}") { //Office 2000

		} else if strings.HasSuffix(guids[i], "6000-11D3-8CFE-0050048383C9}") { //Office XP

		} else if strings.HasSuffix(guids[i], "6000-11D3-8CFE-0150048383C9}") { //Office 2003

		}

		//m, _ := json.Marshal(product)
		//fmt.Printf("%s\n", m)*/
}

func GetInstalledProducts() (products []InstalledProduct) {
	guids := GetAllProducts()
	for i := 0; i < len(guids); i++ {
		product := GetProductDetails(guids[i])
		products = append(products, product)
	}
	return
}
