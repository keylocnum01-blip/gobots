$content = Get-Content main.go -Raw
$content = $content -replace '(?<!func\s)CheckUnbanBots\(', 'handler.CheckUnbanBots('
$content = $content -replace '(?<!func\s)Purgesip\(', 'handler.Purgesip('
$content = $content -replace '(?<!func\s)KickBl\(', 'handler.KickBl('
$content = $content -replace '(?<!func\s)CheckExprd\(', 'handler.CheckExprd('
$content = $content -replace '(?<!func\s)CekDuedate\(', 'handler.CekDuedate('
$content = $content -replace '(?<!func\s)CheckLastActive\(', 'handler.CheckLastActive('
$content = $content -replace '(?<!func\s)Checklistaccess\(', 'handler.CheckListAccess('
$content = $content -replace '(?<!func\s)AutojoinQr22\(', 'handler.AutojoinQr22('
Set-Content main.go $content
