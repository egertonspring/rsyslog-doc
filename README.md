# Rsyslog rules and templates
This is a documentation of rsyslog rule and templates and what they will actually do.
I will only document the "new" advanced format here because the only legacy that counts is "Hogwarts Legacy".

## Forwarding
### Forwarding Example 1
The following rule will forward logs which are tagged "hello-go" to the server which is defined in the target-section. If you have multiple templates defined within the rsyslog.conf.d-File make sure you select the right one in your forwarding rule in the template-section:
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
### Forwarding Example 2
You can also forward the log messages to multiple destinations just by defining multiple forwarding rules:
```
if $programname == 'hello-go' then {
   action(type="omfwd"
       protocol="udp"
       template="hello-go-flat"
       target="rsyslog-target1.domain.local"
       port="514"
       queue.filename="rsyslog-target.domain.local-queue"
       queue.type="linkedList"
)
   action(type="omfwd"
       protocol="udp"
       template="hello-go-flat"
       target="rsyslog-target2.domain.local"
       port="514"
       queue.filename="rsyslog-target.domain.local-queue"
       queue.type="linkedList"
)
}
```
### Forwarding Example 3
The following example forwards every log message from the program 'hello-go' via TCP to rsyslog-target.domain.local on port 11611.
```
if $programname == 'hello-go' then {
   action(type="omfwd"
       protocol="tcp"
       template="hello-go-flat"
       target="rsyslog-target.domain.local"
       port="11611"
       queue.filename="rsyslog-target.domain.local-queue"
       queue.type="linkedList"
)
}
```
### Forwarding Example 4
If you want to forward logs to a specific file, the following rule could help you:
```
if $programname == 'kubelet' then {
   action(type="omfile"
       file="/var/log/kubernetes/kubelet.log"
       template="kubelet-json"
)
}
```
The events of kublet will be logged like this (I have used the template from example 2):
```
{"@TIMESTAMP":"2024-03-06T15:10:07.743775+01:00", "HOST":"rsyslogclient", "SEVERITY":6, "FACILITY":3, "TAG":"kubelet[10332]:", "APP": "APP_KUBELET", "SRC":"kubelet", "MSG":" I0306 15:10:07.743564   10332 status_manager.go:809] \"Failed to get status for pod\" podUID=d8bda651380477fe3a0683c878987dc9 pod=\"kube-system\/kube-scheduler-rsyslogclient\" err=\"Get \\\"https:\/\/192.168.1.199:6443\/api\/v1\/namespaces\/kube-system\/pods\/kube-scheduler-rsyslogclient\\\": dial tcp 192.168.1.199:6443: connect: connection refused\""}
```

## Templating
I have written a little app in go which runs as sstemd services and will write log messages every 3 seconds. This is to test all this logging.

### Templating Example 1
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
### Templating Example 2
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
... ends up logged as this:
```
2024-03-04T16:43:20.288697+01:00 rsyslog-client.domain.local  {"TIME":"2024-03-04T16:43:20.227336+01:00", "HOST":"rpi0", "SEVERITY":6, "FACILITY":3, "TAG":"hello-go[340]:", "CUSTOM_APP_TAG": "APP_HELLOGO", "SRC":"hello-go", "MSG":" Canvas Reads!"}
```
### Templating Example 3
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
### Templating Example 4
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
### Templating Example 5
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
### Templating Example 6
This one
```
template(name="hello-go-experimental-3" type="list") {
    property(name="timestamp" dateFormat="rfc3339")
    constant(value=" ")
    property(name="hostname")
    constant(value=" ")
    property(name="syslogtag")
    constant(value=" ")
    property(name="syslogseverity-text")
    constant(value=":")
    property(name="syslogfacility-text")
    constant(value=" ")
    constant(value="APP_HELLOGO")
    property(name="msg" spifno1stsp="on" )
    property(name="msg" droplastlf="on" )
    constant(value="\n")
    }
```
results in this log line:
```
2024-03-05T13:48:25.123630+01:00 rsyslogclient hello-go[341]: info:daemon APP_HELLOGO Microphone Builds!
```
