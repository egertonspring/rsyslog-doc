# Rsyslog rules and templates
This is a documentation of rsyslog rule and templates and what they will actually do.
I will only document the "new" advanced format here because the only legacy that counts is "Hogwarts Legacy"

## Forwarding
The following rule will forward logs which are tagged "hello-go" to the server which is defined in the target-section:

```
if $programname == 'hello-go' then {
   action(type="omfwd"
       protocol="udp"
       template="hello-go-flat"
       target="rsyslog-target.domain.local"
       port="514"
       queue.filename="rsyslog-target.domain.local-queue"
       queue.type="linkedList"
)
}
```

## Templating
I have written a little app in go which will write log messages every 3 seconds. This is to test all this logging.


The following template ...
```
template(name="hello-go-flat" type="list") {
    property(name="timestamp" dateFormat="rfc3339")
    constant(value=" ")
    property(name="hostname")
    constant(value=" ")
    property(name="syslogtag")
    constant(value=" ")
    constant(value="APP_HELLOGO")
    property(name="msg" spifno1stsp="on" )
    property(name="msg" droplastlf="on" )
    constant(value="\n")
    }
```
...will result in such a log message:
```
2024-03-04T16:40:02.226078+01:00 rsyslog-client hello-go[340]: APP_HELLOGO Plate Encourages!
```

### Next

This template ...
```
template(name="hello-go-json" type="list" option.jsonf="on") {
     property(outname="TIME" name="timereported" dateFormat="rfc3339" format="jsonf")
     property(outname="HOST" name="hostname" format="jsonf")
     property(outname="SEVERITY" name="syslogseverity" caseConversion="upper" format="jsonf" datatyp>
     property(outname="FACILITY" name="syslogfacility" caseConversion="upper" format="jsonf" datatyp>
     property(outname="TAG" name="syslogtag" format="jsonf")
     constant(outname="CUSTOM_APP_TAG" value="APP_HELLOGO" format="jsonf")
     property(outname="SRC" name="app-name" format="jsonf" onEmpty="null")
     property(outname="MSG" name="msg" format="jsonf")
     }
```

... end up logged as this:
```
2024-03-04T16:43:20.288697+01:00 rsyslog-client.domain.local  {"TIME":"2024-03-04T16:43:20.227336+01:00", "HOST":"rpi0", "SEVERITY":6, "FACILITY":3, "TAG":"hello-go[340]:", "CUSTOM_APP_TAG": "APP_HELLOGO", "SRC":"hello-go", "MSG":" Canvas Reads!"}
```

### Next

The following template...

```
template(name="hello-go-long-json" type="list" option.jsonf="on") {
     property(outname="TIME" name="timereported" dateFormat="rfc3339" format="jsonf")
     property(outname="HOST" name="hostname" format="jsonf")
     property(outname="SEVERITY_NUM" name="syslogseverity" format="jsonf")
     property(outname="SEVERITY_TXT" name="syslogseverity-text" format="jsonf")
     property(outname="FACILITY_NUM" name="syslogfacility" format="jsonf")
     property(outname="FACILITY_TXT" name="syslogfacility-text" format="jsonf")
     property(outname="TAG" name="syslogtag" format="jsonf")
     constant(outname="CUSTOM_APP_TAG" value="APP_HELLOGO" format="jsonf")
     property(outname="SRC" name="app-name" format="jsonf" onEmpty="null")
     property(outname="MSG" name="msg" format="jsonf")
     }
```

... is logged as...

```
2024-03-04T16:45:41.283912+01:00 rsyslog-client.domain.local  {"TIME":"2024-03-04T16:45:41.227274+01:00", "HOST":"rpi0", "SEVERITY_NUM":"6", "SEVERITY_TXT":"info", "FACILITY_NUM":"3", "FACILITY_TXT":"daemon", "TAG":"hello-go[340]:", "CUSTOM_APP_TAG": "APP_HELLOGO", "SRC":"hello-go", "MSG":" Canvas Forgets!"}
```

### Next

This one ...
```
template(name="hello-go-experimental-1" type="list") {
    property(name="timestamp" dateFormat="rfc3339")
    constant(value=" ")
    property(name="hostname")
    constant(value=" ")
    property(name="syslogtag")
    constant(value=" ")
    property(name="syslogseverity")
    constant(value=" ")
    property(name="syslogfacility")
    constant(value=" ")
    constant(value="APP_HELLOGO")
    property(name="msg" spifno1stsp="on" )
    property(name="msg" droplastlf="on" )
    constant(value="\n")
    }
```

will be logged in this way:

```
2024-03-04T16:47:47.228103+01:00 rsyslog-client hello-go[340]: 6 3 APP_HELLOGO Cat Drives!
```

### Next

This

```
template(name="hello-go-experimental-2" type="list") {
    property(name="timestamp" dateFormat="rfc3339")
    constant(value=" ")
    property(name="hostname")
    constant(value=" ")
    property(name="syslogtag")
    constant(value=" ")
    property(name="syslogseverity")
    constant(value=" ")
    property(name="syslogseverity-text")
    constant(value=" ")
    property(name="syslogfacility")
    constant(value=" ")
    property(name="syslogfacility-text")
    constant(value=" ")
    constant(value="APP_HELLOGO")
    property(name="msg" spifno1stsp="on" )
    property(name="msg" droplastlf="on" )
    constant(value="\n")
    }
```
will be written in such a way:

```
2024-03-04T16:49:35.227293+01:00 rsyslogclient hello-go[340]: 6 info 3 daemon APP_HELLOGO Knife Enjoys!
```
