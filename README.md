# Telemetry system in Golang

## Why/what
I always wandered how modern car systems might work, this is not a real attempt to build such systems, but instead, a way to understand how they communicate, and use eachothers data to, maybe, give alerts. In other words, this seems just for learning purposes and might not even be close to what systems like these are, also, most use the CAN bus (Controller Area Network), which i plan to implement as one of the strategies to send messages.

In the end this is just a learning project/idea instead of a tool.

## What it does
Currently not much, the generators simply generate data, and they are using the strategy pattern to decide on which broker might be in use, currently only rabbit is implemented, then it will be MQTT, which will be on standby until i understand it. Moreover, the generated data can be faulty (unhealthy) or not faulty (healthy), which can be changed through the two exiting endpoints '/healthy' and '/unhealthy', there is also a third endpoint to kill the system which is '/kill'.

Now its possible to visualize the data on grafna, and the docker compose has everything to run on independent containers. Now, because i am too lazy, i didn't bother yet to search on how to build grafana dashboards, and so i am using some presets, [6671](https://grafana.com/grafana/dashboards/6671-go-processes/) this allows you to see go processess, not related to the telemetry sent from rabbit and [16966](https://grafana.com/api/dashboards/6671/images/4286/image) which was supposed to show the logs via Loki, which is currently isn't working.

The ./prometheus.yaml, ./grafana/provisioning/datasources/datasources.yaml and ./config.alloy were AI generated as i have no knowledge of it and i didn't have time to learn it yet.

## What will it do
There will be a telemetry server, which will store the data as logs, i believe that loki and grafana can be used here, grafana will have dashboards, and loki will serve as the log aggregation system and why not just throw a time series database, like influxDb or TimescaleDb from Postgres.

## Why so many technologies on something so simple?
*Because i want to learn them*


### Contributions
Any contribution is welcome, just make sure its around this idea. Or go crazy and implmenet a cool feature, there is always space to learn something new.

