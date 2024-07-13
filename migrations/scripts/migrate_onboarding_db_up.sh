#!/bin/bash

psql -f migrations/onboarding/enums/001_onboarding_stage_enums.up.sql
psql -f migrations/onboarding/tables/001_onboarding_users.up.sql
