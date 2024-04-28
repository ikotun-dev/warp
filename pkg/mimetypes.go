package main

import "strings"

func getMimeType(extension string) string {
	DefaultMimeType := "text/plain"
	ext := strings.ToLower(extension)
	switch ext {
	case "aac":
		return "audio/aac"
	case "abw":
		return "application/x-abiword"
	case "apng":
		return "image/apng"
	case "arc":
		return "application/x-freearc"
	case "avif":
		return "image/avif"
	case "avi":
		return "video/x-msvideo"
	case "azw":
		return "application/vnd.amazon.ebook"
	case "bmp":
		return "image/bmp"
	case "bz":
		return "application/x-bzip"
	case "bz2":
		return "application/x-bzip2"
	case "cda":
		return "application/x-cdf"
	case "csh":
		return "application/x-csh"
	case "css":
		return "text/css"
	case "csv":
		return "text/csv"
	case "doc":
		return "application/msword"
	case "docx":
		return "application/vnd.openxmlformats-officedocument.wordprocessingml.document"
	case "eot":
		return "application/vnd.ms-fontobject"
	case "epub":
		return "application/epub+zip"
	case "gz":
		return "application/gzip"
	case "gif":
		return "image/gif"
	case "htc", "htm", "html", "stm":
		return "text/html; charset=utf-8"
	case "htt":
		return "text/webviewhtml"
	case "ico":
		return "image/vnd.microsoft.icon"
	case "ics":
		return "text/calendar"
	case "jar":
		return "application/java-archive"
	case "jpeg", "jpg":
		return "image/jpeg"
	case "js", "mjs":
		return "text/javascript"
	case "json":
		return "application/json"
	case "jsonld":
		return "application/ld+json"
	case "mid", "midi":
		return "audio/midi"
	case "mht", "mhtml", "nws":
		return "message/rfc822"
	case "mp3":
		return "audio/mpeg"
	case "mp4":
		return "video/mp4"
	case "mpeg", "mpg", "mpa", "mpe", "mp2", "mpv2":
		return "video/mpeg"
	case "mpkg":
		return "application/vnd.apple.installer+xml"
	case "mov", "qt":
		return "video/quicktime"
	case "odp":
		return "application/vnd.oasis.opendocument.presentation"
	case "ods":
		return "application/vnd.oasis.opendocument.spreadsheet"
	case "odt":
		return "application/vnd.oasis.opendocument.text"
	case "oga":
		return "audio/ogg"
	case "ogv":
		return "video/ogg"
	case "ogx":
		return "application/ogg"
	case "opus":
		return "audio/opus"
	case "otf":
		return "font/otf"
	case "png":
		return "image/png"
	case "pdf":
		return "application/pdf"
	case "php":
		return "application/x-httpd-php"
	case "ppt":
		return "application/vnd.ms-powerpoint"
	case "pptx":
		return "application/vnd.openxmlformats-officedocument.presentationml.presentation"
	case "rgb":
		return "image/x-rgb"
	case "rar":
		return "application/vnd.rar"
	case "rtf":
		return "application/rtf"
	case "rtx":
		return "text/richtext"
	case "sh":
		return "application/x-sh"
	case "svg":
		return "image/svg+xml"
	case "tar":
		return "application/x-tar"
	case "tif", "tiff":
		return "image/tiff"
	case "ts":
		return "video/mp2t"
	case "ttf":
		return "font/ttf"
	case "txt", "c", "h", "bas":
		return "text/plain"
	case "vcf":
		return "text/vcard"
	case "vsd":
		return "application/vnd.visio"
	case "wasm":
		return "application/wasm"
	case "wav":
		return "audio/wav"
	case "weba":
		return "audio/webm"
	case "webm":
		return "video/webm"
	case "webp":
		return "image/webp"
	case "woff":
		return "font/woff"
	case "woff2":
		return "font/woff2"
	case "xhtml":
		return "application/xhtml+xml"
	case "xls":
		return "application/vnd.ms-excel"
	case "xlsx":
		return "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
	case "xml":
		return "text/xml"
	case "xul":
		return "application/vnd.mozilla.xul+xml"
	case "zip":
		return "application/zip"
	case "3gp":
		return "video/3gpp"
	case "3g2":
		return "video/3gpp2"
	case "7z":
		return "application/x-7z-compressed"
	default:
		return DefaultMimeType
	}
}
