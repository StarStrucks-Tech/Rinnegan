-- Create Users table
CREATE TABLE Users (
    actor_id VARCHAR(30) PRIMARY KEY,
    first_name VARCHAR(40) NOT NULL,
    middle_name VARCHAR(40),
    last_name VARCHAR(20) NOT NULL,
    phone_number VARCHAR(10),
    email VARCHAR(40),
    phone_verification_status PHONE_VERIFICATION_STATUS NOT NULL DEFAULT 'PHONE_NOT_INITIATED',
    pan_verification_status PAN_VERIFICATION_STATUS NOT NULL DEFAULT 'PAN_VERIFICATION_NOT_INITIATED',
    extra_details_collection_status VARCHAR(50),
    consent_stage_status CONSENT_STAGE_STATUS NOT NULL DEFAULT 'CONSENT_NOT_INITIATED',
    liveliness_check_status LIVENESS_CHECK_STATUS NOT NULL DEFAULT 'LIVENESS_NOT_INITIATED'
);
