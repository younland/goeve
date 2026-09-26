package goeve

// Languages supported by the Accept-Language header and the `language` query parameter.
// Accept-Language 请求头与 `language` 查询参数支持的语言。
const (
	// LanguageEnglish renders responses in English. / LanguageEnglish 使用英语返回。
	LanguageEnglish = "en"
	// LanguageEnglishUS renders responses in American English. / LanguageEnglishUS 使用美式英语返回。
	LanguageEnglishUS = "en-us"
	// LanguageChinese renders responses in Simplified Chinese. / LanguageChinese 使用简体中文返回。
	LanguageChinese = "zh"
)

// DefaultDatasource is the only datasource of the NetEase EVE server (equivalent
// to "tranquility" on the international server).
//
// DefaultDatasource 是网易 EVE 服务器的唯一数据源（相当于国际服的 "tranquility"）。
const DefaultDatasource = "tranquility"
