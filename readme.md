log2team
=========

`log2team` transfers log from STDIN to Microsoft Teams's channel by WebHook.

    log2team WEB-HOOK-URL TITLE-STRING < LOG-FILE

For example:

    cat /var/log/backup.log | log2team "https://outlook.office.com/webhook/.../IncomingWebhook/..." "Backup to /var/log/backup.log"

[Download Binary](https://github.com/zat-kaoru-hayama/log2team/releases)
