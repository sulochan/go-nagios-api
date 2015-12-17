# Nagios JSON API

Intended to provide a JSON API like features for Nagios installations. Its written in golang and therefore systems wanting to run it only need to run a single binary for their distribution.

Installation:
==
1. clone the repo and build for your system.
2. download the binary for your system.

Run:
==
```
$ ./nagios-api --cache-file=/opt/nagios/object.cache --status-file=/opt/nagios/status.dat --command-file=/opt/nagios/nagios.cmd
```
It will start the api service on port 8080. If you wish to change the port simply pass --port=80 to make it run on port 80. For running in production see init scripts.

API Calls
==

#### Hosts and Services
```
 GET /hosts : get all configured hosts
 GET /host/<hostname> : get this host
 GET /hoststatus : get all hoststatus
 GET /hoststatus/<hostname> : get hoststatus for this host
 GET /hostgroups : get all configured hostgroups
 GET /services : get all configured services
 GET /servicestatus : get all servicestatus
 GET /servicestatus/<servicename> : get service status for this service
```

#### External Commands
```
POST /disable_notifications 
POST /enable_notifications
POST /disable_host_check  
POST /enable_host_check   
POST /disable_host_notifications
POST /enable_host_notifications
POST /acknowledge_host_problem
POST /acknowledge_service_problem
...
see code for full list of available commands.
```
