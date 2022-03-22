package internal

import "path"

const ChromePath = "/Applications/Google Chrome.app"
const ChromeTagKey = "KSChannelID"
const ChromeBrandPath = "/Library/Google/Google Chrome Brand.plist"
const ChromeBrandKey = "KSBrandID"
const ChromeVersionKey = "KSVersion"

var ChromeInfoPlistPath = path.Join(ChromePath, "Contents/Info.plist")
