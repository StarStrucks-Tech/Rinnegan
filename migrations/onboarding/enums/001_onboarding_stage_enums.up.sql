-- Create enum types if they don't exist
CREATE TYPE PHONE_VERIFICATION_STATUS AS ENUM ('unverified', 'verified');
CREATE TYPE PAN_VERIFICATION_STATUS AS ENUM ('unverified', 'verified');
CREATE TYPE EXTRA_DETAILS_COLLECTION_STATUS AS ENUM ('incomplete', 'complete');
CREATE TYPE CONSENT_STAGE_STATUS AS ENUM ('CONSENT_NOT_INITIATED', 'CONSENT_INITIATED');
CREATE TYPE LIVENESS_CHECK_STATUS AS ENUM ('not_checked', 'checked');