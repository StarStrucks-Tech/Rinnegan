-- Create Users table
CREATE TABLE Users (
    actor_id VARCHAR PRIMARY KEY,
    first_name VARCHAR NOT NULL,
    middle_name VARCHAR,
    last_name VARCHAR NOT NULL,
    phone_number VARCHAR,
    email VARCHAR,
    phone_verification_status PHONE_VERIFICATION_STATUS NOT NULL DEFAULT 'unverified',
    pan_verification_status PAN_VERIFICATION_STATUS NOT NULL DEFAULT 'unverified',
    extra_details_collection_status EXTRA_DETAILS_COLLECTION_STATUS NOT NULL DEFAULT 'incomplete',
    consent_stage_status CONSENT_STAGE_STATUS NOT NULL DEFAULT 'CONSENT_NOT_INITIATED',
    liveliness_check_status LIVENESS_CHECK_STATUS NOT NULL DEFAULT 'not_checked'
);
