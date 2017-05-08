#!/bin/sh

bq mk logging_sample
bq mk -t logging_sample.app ./schema/app.json
bq mk -t logging_sample.event ./schema/event.json
