#!/bin/sh

bq mk logging_sample
bq mk -t logging_sample.app_log ./schema/app.json
bq mk -t logging_sample.event_log ./schema/event.json
