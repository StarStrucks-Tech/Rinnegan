#!/bin/bash

psql -f 001_onboarding_stage_enums.down.sql
psql -f 001_onboarding_users.down.sql
