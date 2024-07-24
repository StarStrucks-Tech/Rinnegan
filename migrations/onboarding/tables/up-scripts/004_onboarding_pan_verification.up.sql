-- Create PANVerification table
CREATE TABLE PAN_Verification_Stage (
    pan_verification_id VARCHAR(50) PRIMARY KEY,
    actor_id VARCHAR(50),
    pan_number VARCHAR(20),
    dob DATE,
    pan_check_consent BOOLEAN,
    pan_verification_status PAN_VERIFICATION_STATUS,
    FOREIGN KEY (actor_id) REFERENCES Users(actor_id)
);